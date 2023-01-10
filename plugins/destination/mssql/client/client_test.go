package client

import (
	"os"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/microsoft/go-mssqldb/msdsn"
)

func getTestConnection() string {
	if testConn := os.Getenv("CQ_DEST_MS_SQL_TEST_CONN"); len(testConn) > 0 {
		return testConn
	}

	return msdsn.Config{
		Port:       1433,
		Host:       "localhost",
		Database:   "cloudquery",
		User:       "SA",
		Password:   "yourStrongP@ssword",
		Encryption: msdsn.EncryptionDisabled,
		LogFlags:   (msdsn.LogRetries << 1) - 1,
	}.URL().String()
}

func TestPgPlugin(t *testing.T) {
	p := destination.NewPlugin("mssql", plugin.Version, New)
	destination.PluginTestSuiteRunner(t, p,
		Spec{
			ConnectionString: getTestConnection(),
		},
		destination.PluginTestSuiteTests{},
	)
}
