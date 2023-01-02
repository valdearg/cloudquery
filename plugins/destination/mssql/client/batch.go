package client

import (
	"context"
	"sync/atomic"
)

type stmt struct {
	query  string
	params []any
}

// Here send will try sending all queries.
// It fails if any of requests returns non-nil err.
// The returned error will be the first error encountered.
func (c *Client) send(ctx context.Context, batch []*stmt) error {
	if len(batch) == 0 {
		return nil
	}

	var errs []error
	for _, s := range batch {
		if _, err := c.db.ExecContext(ctx, s.query, s.params...); err != nil {
			errs = append(errs, err)
		}
	}

	atomic.AddUint64(&c.metrics.Errors, uint64(len(errs)))
	atomic.AddUint64(&c.metrics.Writes, uint64(len(batch)-len(errs)))

	if len(errs) > 0 {
		return errs[0]
	}
	return nil
}
