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

	// SavingInProgress is true when rows are being saved
	// it is set to true by the front at the begining and set back
	// to false by the front
	// This information is used by the back to compute when all rows to be updated
	// have been updated
	SavingInProgress bool
}
