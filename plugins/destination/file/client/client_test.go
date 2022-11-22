package client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

func getTestLogger(t *testing.T) zerolog.Logger {
	t.Helper()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	return zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
}

func TestClient(t *testing.T) {
	ctx := context.Background()
	client, err := New(ctx, getTestLogger(t), specs.Destination{
		WriteMode: specs.WriteModeAppend,
		Spec:      Spec{},
	})
	if err != nil {
		t.Fatal(err)
	}
	if client == nil {
		t.Fatal("client is nil")
	}
	if err := client.Close(ctx); err != nil {
		t.Fatal(err)
	}

	_, err = New(ctx, getTestLogger(t), specs.Destination{})
	if err.Error() != "file destination only supports append mode" {
		t.Fatal("expected error: 'file destination only supports append mode'")
	}
}

func newTestMode(ctx context.Context, logger zerolog.Logger, spec specs.Destination) (plugins.DestinationClient, error) {
	c, err := New(ctx, logger, spec)
	if err != nil {
		return nil, err
	}
	c.(*Client).testMode = true
	return c, nil
}

func TestPluginCSV(t *testing.T) {
	p := plugins.NewDestinationPlugin("file", "development", newTestMode)
	// plugins.DestinationPluginTestSuiteRunner(t, p,
	// 	Spec{
	// 		Directory: t.TempDir(),
	// 		Backend:   BackendTypeLocal,
	// 		Format:    FormatTypeCSV,
	// 	},
	// 	plugins.DestinationTestSuiteTests{
	// 		SkipOverwrite:   true,
	// 		SkipDeleteStale: true,
	// 	},
	// )

	plugins.DestinationPluginTestSuiteRunner(t, p,
		Spec{
			Directory: "cq-yev-test/dest-plugin-file",
			Backend:   BackendTypeGCS,
			Format:    FormatTypeCSV,
		},
		plugins.DestinationTestSuiteTests{
			SkipOverwrite:   true,
			SkipDeleteStale: true,
		},
	)
}

func TestPluginJSON(t *testing.T) {
	p := plugins.NewDestinationPlugin("file", "development", newTestMode)
	// plugins.DestinationPluginTestSuiteRunner(t, p,
	// 	Spec{
	// 		Directory: t.TempDir(),
	// 		Backend:   BackendTypeLocal,
	// 		Format:    FormatTypeJSON,
	// 	},
	// 	plugins.DestinationTestSuiteTests{
	// 		SkipOverwrite:   true,
	// 		SkipDeleteStale: true,
	// 	},
	// )

	plugins.DestinationPluginTestSuiteRunner(t, p,
		Spec{
			Directory: "cq-yev-test/dest-plugin-file",
			Backend:   BackendTypeGCS,
			Format:    FormatTypeJSON,
		},
		plugins.DestinationTestSuiteTests{
			SkipOverwrite:   true,
			SkipDeleteStale: true,
		},
	)
}
