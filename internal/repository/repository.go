package repository

import "database/sql"

type Repository struct {
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{}
}

// spec := generator.ABTestSpec{
// 		ServiceName: "Service1",
// 		RouteName:   "service1",
// 		StructName:  "Service1Out",
// 		Fields: []generator.Field{
// 			{Name: "Service1Param", Type: "int", Tag: "`json:\"service1Param\"`"},
// 			{Name: "ABType", Type: "string", Tag: "`json:\"ABType\"`"},
// 			{Name: "ABExists", Type: "bool", Tag: "`json:\"AbExists\"`"},
// 			{Name: "MyArray", Type: "[]string", Tag: "`json:\"MyArray\"`"},
// 		},
// 		Modifications: []generator.Modification{
// 			{
// 				Name:  "Service1Param",
// 				Type:  "multiply",
// 				Value: "in.Service1Param * 15",
// 			},
// 			{
// 				Name:  "ABType",
// 				Type:  "set",
// 				Value: `"MyCustomABTest"`,
// 			},
// 			{
// 				Name:  "ABExists",
// 				Type:  "set",
// 				Value: "true",
// 			},
// 			{
// 				Name:  "MyArray",
// 				Type:  "append",
// 				Value: "[]string{\"hello1\", \"hello2\"}",
// 			},
// 		},
// 	}
