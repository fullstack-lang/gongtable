package models

type Table struct {
	Name             string
	DisplayedColumns []*DisplayedColumn
	Rows             []*Row

	HasFiltering     bool
	HasColumnSorting bool
	HasPaginator     bool

	HasCheckableRows bool

	HasSaveButton bool

	CanDragDropRows bool
}
