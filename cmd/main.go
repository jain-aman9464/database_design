package main

import (
	"github.com/tokopedia/test/database_design/src/model"
)

func main() {
	tableName := "employee"
	id := model.NewColumn("id", model.INT)
	name := model.NewColumn("name", model.STRING)
	age := model.NewColumn("age", model.INT)
	salary := model.NewColumn("salary", model.INT)

	db := model.NewDatabase("myDB")
	columns := make([]model.Column, 0)
	columns = append(columns, id, name, age, salary)
	db.CreateTable(tableName, columns)
	columnValues := map[model.Column]interface{}{
		name:   "John",
		age:    25,
		salary: 10000,
	}

	columnValues2 := map[model.Column]interface{}{
		name:   "Kim",
		age:    24,
		salary: 20000,
	}

	db.InsertTableRows(tableName, columnValues)
	db.InsertTableRows(tableName, columnValues2)
	db.PrintTableAllRows(tableName)
	//db.FilterTableRecordsByColumnValue(tableName, age, 24)
	//db.Truncate(tableName)
	//db.DropTable(tableName)
	//db.PrintTableAllRows(tableName)

}
