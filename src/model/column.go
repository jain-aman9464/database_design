package model

type ColumnType int

const (
	INT ColumnType = iota
	STRING
)

type Column struct {
	columnName string
	columnType ColumnType
}

func NewColumn(columnName string, columnType ColumnType) Column {
	return Column{
		columnName: columnName,
		columnType: columnType,
	}
}

func (c Column) GetColumnName() string {
	return c.columnName
}
