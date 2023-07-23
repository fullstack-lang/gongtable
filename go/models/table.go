package models

type Table struct {
	Name             string
	DisplayedColumns []*DisplayedColumn
	Rows             []*Row
}
