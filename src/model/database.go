package model

import (
	"fmt"
)

type Database struct {
	DatabaseName string
	tableMap     map[string]*Table
}

func NewDatabase(name string) *Database {
	return &Database{
		DatabaseName: name,
		tableMap:     make(map[string]*Table, 0),
	}
}

func (d *Database) CreateTable(tableName string, columns []Column) {
	if d.checkIfTableExists(tableName) {
		fmt.Printf("Tablename: %s already exist", tableName)
		return
	}

	d.tableMap[tableName] = NewTable(tableName, columns)
	fmt.Printf("\nTable %s created\n", tableName)
}

func (d *Database) DropTable(tableName string) {
	if !d.checkIfTableExists(tableName) {
		fmt.Printf("Tablename: %s does not exist", tableName)
	}

	delete(d.tableMap, tableName)
	fmt.Printf("Tablename: %s dropped!!", tableName)
}

func (d *Database) Truncate(tableName string) {
	if !d.checkIfTableExists(tableName) {
		fmt.Printf("Tablename: %s does not exist", tableName)
	}

	table := d.tableMap[tableName]
	table.TruncateRows()

	fmt.Printf("Tablename: %s truncated!!", tableName)
}

func (d *Database) InsertTableRows(tableName string, columnValues map[Column]interface{}) {
	if !d.checkIfTableExists(tableName) {
		fmt.Printf("Tablename: %s does not exist", tableName)
	}

	table := d.tableMap[tableName]
	table.InsertRow(columnValues)
	fmt.Println(table.GetRows())
	fmt.Println("Rows Inserted")
}

func (d Database) PrintTableAllRows(tableName string) {
	if !d.checkIfTableExists(tableName) {
		return
	}

	table := d.tableMap[tableName]
	table.PrintRows()
}

func (d Database) FilterTableRecordsByColumnValue(tableName string, column Column, value interface{}) {
	if !d.checkIfTableExists(tableName) {
		return
	}

	table := d.tableMap[tableName]
	table.GetRecordsByColumnValue(column, value)
}

func (d Database) checkIfTableExists(tableName string) bool {
	if _, ok := d.tableMap[tableName]; !ok {
		fmt.Printf("TableName %s does not exist", tableName)
		return false
	}

	return true
}
