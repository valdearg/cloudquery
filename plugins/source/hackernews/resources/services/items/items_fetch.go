package items

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/cloudquery/cloudquery/plugins/source/hackernews/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/hermanschaaf/hackernews"
	"golang.org/x/sync/errgroup"
)

// fetchItems gets called by the CloudQuery SDK to fetch the data for the `hackernews_items` table. It is passed a context and a
// schema.ClientMeta interface which is used to access the plugin configuration and the state backend. The
// resource argument is not used here, as it refers to the parent table, which is not applicable in this case.
// The res channel is used to send the fetched data back to the CloudQuery SDK.
//
// Because this is an incremental table, we must first load the cursor from the state backend, and then update
// it whenever we've successfully fetched a new item. The cursor is a string, and in this case it is the ID of
// the last item we've fetched.
//
// Incremental tables should guarantee at-least-once delivery. The Hacker News API gives us the current maximum item ID,
// and we need to fetch all items up to that ID. Once we've fetched all items up to a certain ID, we can safely
// update the cursor. If we crash before updating the cursor, we will fetch the same items again on the next run.
//
// This resolver is different from most in that it needs to concurrently fetch items using its own goroutines. Usually,
// a resolver with pagination will not have this ability and will simply iterate over all pages.
func fetchItems(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	tableName := Items().Name
	value, err := c.Backend.Get(ctx, tableName, c.ID())
	if err != nil {
		return fmt.Errorf("failed to retrieve state from backend: %w", err)
	}

	// read the cursor from the state, or default to 0 if it's not set
	cursor := 0
	if value == "" {
		c.Logger().Info().
			Str("table", tableName).Str("client_id", c.ID()).Msgf("No previous cursor found")
	} else {
		cursor, err = strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("failed to convert cursor to int: %w", err)
		}
		c.Logger().Info().
			Str("table", tableName).Str("client_id", c.ID()).Msg("Found previous cursor with value " + strconv.Itoa(cursor))
	}

	// find the max item ID from the Hacker News API
	maxID, err := c.HackerNews.MaxItemID(ctx)
	if err != nil {
		return err
	}
	c.Logger().Info().Msg("Found max ID, reading up to " + strconv.Itoa(maxID))

	// Fetch items in batches of (max) 1000.
	// This is not necessarily the most efficient way of doing it, but this code
	// is meant to be for instructional purposes as an example of updating cursors,
	// so we're keeping the logic relatively simple.
	// The important thing is that the state backend does not ensure that the cursor
	// is strictly increasing--it is the responsibility of the resolver to ensure this.
	for cursor < maxID {
		endID := cursor + 1000
		if endID > maxID {
			endID = maxID
		}
		err := fetchBatch(ctx, c, tableName, cursor+1, endID, res)
		if err != nil {
			return err
		}
		// save the new cursor position after a batch has been successfully fetched
		cursor = endID
		err = c.Backend.Set(ctx, tableName, c.ID(), strconv.Itoa(cursor))
		if err != nil {
			return fmt.Errorf("failed to save state to backend: %w", err)
		}
	}
	return nil
}

// fetchBatch fetches the items in the inclusive range [startID, endID] and sends them to the res channel. It blocks
// until the entire batch has either succeeded or failed.
func fetchBatch(ctx context.Context, c *client.Client, tableName string, startID, endID int, res chan<- any) error {
	g, gctx := errgroup.WithContext(ctx)
	g.SetLimit(c.Spec.ItemConcurrency)
	for i := startID; i <= endID; i++ {
		i := i
		g.Go(func() error {
			return c.RetryOnError(gctx, tableName, func() error {
				return fetchItem(gctx, c, i, res)
			})
		})
	}
	return g.Wait()
}

// fetchItem fetches a single item from the Hacker News API and sends it to the CloudQuery SDK
// to dispatch to destinations via the res channel.
func fetchItem(ctx context.Context, c *client.Client, itemID int, res chan<- any) error {
	item, err := c.HackerNews.GetItem(ctx, itemID)
	if err != nil {
		var httpErr hackernews.HTTPError
		if errors.As(err, &httpErr) && httpErr.Code == 404 {
			// Assume item was deleted
			return nil
		}
		return err
	}
	res <- item
	return nil
}
