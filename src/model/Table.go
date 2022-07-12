package model

import (
	"fmt"
	"sort"
)

type Table struct {
	autoIncID int
	tableName string
	columnMap map[string]Column
	rows      []Row
}

func NewTable(tableName string, columns []Column) *Table {
	tableObj := Table{}
	tableObj.columnMap = make(map[string]Column)
	for _, column := range columns {
		tableObj.columnMap[column.GetColumnName()] = column
	}

	tableObj.autoIncID = 0
	tableObj.tableName = tableName
	tableObj.rows = make([]Row, 0)

	return &tableObj
}

func (t *Table) InsertRow(columnValues map[Column]interface{}) {
	for column, _ := range columnValues {
		if !t.checkIfColumnExists(column.GetColumnName()) {
			return
		}
	}

	rowID := t.getAutoIncrementId()
	row := NewRow(rowID, columnValues)
	t.rows = append(t.rows, row)
}

func (t Table) GetRows() []Row {
	return t.rows
}
func (t Table) PrintRows() {
	fmt.Printf("Printing all rows for table: %s \n", t.tableName)
	t.printRecords(t.rows)
}

func (t Table) GetRecordsByColumnValue(column Column, value interface{}) {
	rows := make([]Row, 0)

	for _, row := range t.rows {
		columnVal := row.GetColumnData()[column]
		if columnVal == value {
			rows = append(rows, value.(Row))
		}
	}

	t.printRecords(rows)
}

func (t *Table) TruncateRows() {
	t.rows = nil
}

func (t *Table) getAutoIncrementId() int {
	t.autoIncID++
	return t.autoIncID
}

func (t Table) checkIfColumnExists(columnName string) bool {
	if _, ok := t.columnMap[columnName]; !ok {
		fmt.Printf("Tablename: %s does not contain column %s\n", t.tableName, columnName)
		return false
	}

	return true
}

func (t Table) printRecords(rows []Row) {
	keys := []string{}

	for k, _ := range t.columnMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("\t %s \t", key)
	}

	//for column, _ := range t.columnMap {
	//	t.columnMap[]
	//	fmt.Printf("\t %s \t", column)
	//}

	for _, row := range rows {
		fmt.Printf("\n\t %d \t", row.GetRowID())
		for _, val := range row.GetColumnData() {
			fmt.Printf("\t %v \t", val)
		}
	}

	fmt.Println("")
}
