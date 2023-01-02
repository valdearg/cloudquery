package client

import (
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
)

func getTestConnection() string {
	if testConn := os.Getenv("CQ_DEST_MS_SQL_TEST_CONN"); len(testConn) > 0 {
		return testConn
	}

	query := url.Values{
		"encrypt": []string{"disable"},
	}

	return (&url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword("mssql", "<YourStrong@Passw0rd>"),
		Host:     fmt.Sprintf("%s:%d", "localhost", 1433),
		RawQuery: query.Encode(),
	}).String()
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
