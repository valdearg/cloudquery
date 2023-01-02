package client

import (
	mssql "github.com/microsoft/go-mssqldb"
	"github.com/microsoft/go-mssqldb/azuread"
)

type Flavor string

const (
	FlavorAzure = "azure"
	FlavorMS    = "ms"
)

type Spec struct {
	ConnectionString string `json:"connection_string,omitempty"`
	Flavor           Flavor `json:"flavor,omitempty"`
}

func (s *Spec) Connector() (*mssql.Connector, error) {
	if s.Flavor == FlavorAzure {
		return azuread.NewConnector(s.ConnectionString)
	}
	return mssql.NewConnector(s.ConnectionString)
}
