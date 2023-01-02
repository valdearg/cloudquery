package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/mssql/client"
	"github.com/cloudquery/cloudquery/plugins/destination/mssql/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/serve"
)

const (
	sentryDSN = "TODO"
)

func main() {
	serve.Destination(
		destination.NewPlugin(
			"mssql",
			plugin.Version,
			client.New,
			destination.WithDefaultBatchSize(1000),
		),
		serve.WithDestinationSentryDSN(sentryDSN),
	)
}
