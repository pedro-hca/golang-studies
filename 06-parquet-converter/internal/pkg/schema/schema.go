package schema

import "github.com/apache/arrow/go/v16/arrow"

func GetRecordSchema() *arrow.Schema {
	schemaRecord := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
			{Name: "city", Type: arrow.BinaryTypes.String},
			{Name: "review", Type: arrow.PrimitiveTypes.Float64},
		}, nil,
	)
	return schemaRecord
}

func GetStructSchema() *arrow.StructType {

	// Schema Struct
	schemaStruct := arrow.StructOf(
		arrow.Field{Name: "id", Type: arrow.PrimitiveTypes.Int64},
		arrow.Field{Name: "name", Type: arrow.BinaryTypes.String},
		arrow.Field{Name: "city", Type: arrow.BinaryTypes.String},
		arrow.Field{Name: "review", Type: arrow.PrimitiveTypes.Float64},
	)
	return schemaStruct
}
