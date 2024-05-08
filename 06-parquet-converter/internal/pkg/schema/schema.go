package schema

import "github.com/apache/arrow/go/v16/arrow"

func GetRecordSchema() *arrow.Schema {
	schemaRecord := arrow.NewSchema(Fields, nil)
	return schemaRecord
}

func GetStructSchema() *arrow.StructType {
	schemaStruct := arrow.StructOf(Fields...)
	return schemaStruct
}

var Fields []arrow.Field = []arrow.Field{
	{Name: "id", Type: arrow.PrimitiveTypes.Int64},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "city", Type: arrow.BinaryTypes.String},
	{Name: "review", Type: arrow.PrimitiveTypes.Float64},
}

// var Fields []arrow.Field = []arrow.Field{
// 	{Name: "id", Nullable: true, Type: arrow.PrimitiveTypes.Int64},
// 	{Name: "name", Nullable: true, Type: arrow.BinaryTypes.String},
// 	{Name: "city", Nullable: true, Type: arrow.BinaryTypes.String},
// 	{Name: "review", Nullable: true, Type: arrow.PrimitiveTypes.Float64},
// 	{Name: "created_at", Nullable: true, Type: arrow.BinaryTypes.String},
// 	{Name: "average_price", Nullable: true, Type: arrow.PrimitiveTypes.Int64},
// 	{Name: "is_active", Nullable: true, Type: arrow.PrimitiveTypes.Int64},
// }
