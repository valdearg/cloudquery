package client

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/xanzy/go-gitlab"
)

func SharedTransformers() []transformers.StructTransformerOption {
	return []transformers.StructTransformerOption{
		transformers.WithUnwrapAllEmbeddedStructs(),
		transformers.WithUnwrapStructFields("Spec", "Status"),
		transformers.WithTypeTransformer(typeTransformer),
	}
}

func typeTransformer(field reflect.StructField) (schema.ValueType, error) {
	timestamp := gitlab.ISOTime{}
	switch field.Type {
	case reflect.TypeOf(timestamp), reflect.TypeOf(&timestamp):
		return schema.TypeTimestamp, nil
	default:
		return schema.TypeInvalid, nil
	}
}
