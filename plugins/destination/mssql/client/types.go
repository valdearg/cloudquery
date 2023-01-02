package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func msSQLType(t schema.ValueType) string {
	switch t {
	case schema.TypeBool:
		return "bit"
	case schema.TypeInt:
		return "bigint"
	case schema.TypeFloat:
		return "float(53)"
	case schema.TypeUUID:
		return "uniqueidentifier"
	case schema.TypeByteArray:
		return "varbinary(max)"
	case schema.TypeTimestamp:
		return "datetime2(7)"
	case schema.TypeString,
		schema.TypeStringArray,
		schema.TypeJSON,
		schema.TypeUUIDArray,
		schema.TypeCIDR,
		schema.TypeCIDRArray,
		schema.TypeMacAddr,
		schema.TypeMacAddrArray,
		schema.TypeInet,
		schema.TypeInetArray,
		schema.TypeIntArray:
		return "nvarchar(max)"
	default:
		panic("unknown type " + t.String())
	}
}
