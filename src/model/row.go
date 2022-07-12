package model

type Row struct {
	RowID      int
	columnData map[Column]interface{}
}

func NewRow(rowID int, columnData map[Column]interface{}) Row {
	return Row{
		RowID:      rowID,
		columnData: columnData,
	}
}

func (r Row) GetRowID() int {
	return r.RowID
}

func (r Row) GetColumnData() map[Column]interface{} {
	return r.columnData
}
