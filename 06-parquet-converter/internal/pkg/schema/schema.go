package schema

import "github.com/apache/arrow/go/v16/arrow"

func GetRecordSchema() *arrow.Schema {
	schemaRecord := arrow.NewSchema(Fields, nil)
	return schemaRecord
}

func GetStructSchema() *arrow.StructType {

	// Schema Struct
	schemaStruct := arrow.StructOf(Fields...)
	return schemaStruct
}

var Fields []arrow.Field = []arrow.Field{
	{Name: "id", Type: arrow.PrimitiveTypes.Int64},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "city", Type: arrow.BinaryTypes.String},
	{Name: "review", Type: arrow.PrimitiveTypes.Float64},
}
