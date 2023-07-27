package main

import (
	"time"

	"github.com/fullstack-lang/gongtable/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_stage models.StageStruct
var ___dummy__Time_stage time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_stage map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["stage"] = stageInjection
// }

// stageInjection will stage objects of database "stage"
func stageInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Cell
	__Cell__000000_Row_1_Cell_1 := (&models.Cell{Name: `Row 1 - Cell 1`}).Stage(stage)
	__Cell__000001_Row_1_Cell_2 := (&models.Cell{Name: `Row 1 - Cell 2`}).Stage(stage)
	__Cell__000002_Row_1_Cell_3 := (&models.Cell{Name: `Row 1 - Cell 3`}).Stage(stage)
	__Cell__000003_Row_1_Cell_4 := (&models.Cell{Name: `Row 1 - Cell 4`}).Stage(stage)
	__Cell__000004_Row_2_Cell_1 := (&models.Cell{Name: `Row 2 - Cell 1`}).Stage(stage)
	__Cell__000005_Row_2_Cell_2 := (&models.Cell{Name: `Row 2 - Cell 2`}).Stage(stage)
	__Cell__000006_Row_2_Cell_3 := (&models.Cell{Name: `Row 2 - Cell 3`}).Stage(stage)
	__Cell__000007_Row_2_Cell_4 := (&models.Cell{Name: `Row 2 - Cell 4`}).Stage(stage)

	// Declarations of staged instances of CellFloat64
	__CellFloat64__000000_Row_2_Cell_2 := (&models.CellFloat64{Name: `Row 2 - Cell 2`}).Stage(stage)

	// Declarations of staged instances of CellIcon
	__CellIcon__000000_Row_1_Cell_3 := (&models.CellIcon{Name: `Row 1 - Cell 3`}).Stage(stage)
	__CellIcon__000001_Row_2_Cell_3 := (&models.CellIcon{Name: `Row 2 - Cell 3`}).Stage(stage)

	// Declarations of staged instances of CellInt
	__CellInt__000000_Row_1_Cell_2 := (&models.CellInt{Name: `Row 1 - Cell 2`}).Stage(stage)

	// Declarations of staged instances of CellString
	__CellString__000000_Row_1_Cell_1_ := (&models.CellString{Name: `Row 1 - Cell 1 `}).Stage(stage)
	__CellString__000001_Row_1_Cell_4 := (&models.CellString{Name: `Row 1 - Cell 4`}).Stage(stage)
	__CellString__000002_Row_2_Cell_1_ := (&models.CellString{Name: `Row 2 - Cell 1 `}).Stage(stage)
	__CellString__000003_Row_2_Cell_4 := (&models.CellString{Name: `Row 2 - Cell 4`}).Stage(stage)

	// Declarations of staged instances of DisplayedColumn
	__DisplayedColumn__000000_Column_1_String := (&models.DisplayedColumn{Name: `Column 1 - String`}).Stage(stage)
	__DisplayedColumn__000001_Column_2_Int := (&models.DisplayedColumn{Name: `Column 2 - Int`}).Stage(stage)
	__DisplayedColumn__000002_Column_3_Icon := (&models.DisplayedColumn{Name: `Column 3 - Icon`}).Stage(stage)
	__DisplayedColumn__000003_Column_4 := (&models.DisplayedColumn{Name: `Column 4`}).Stage(stage)

	// Declarations of staged instances of Row
	__Row__000000_Row_1 := (&models.Row{Name: `Row 1`}).Stage(stage)
	__Row__000001_Row_2 := (&models.Row{Name: `Row 2`}).Stage(stage)

	// Declarations of staged instances of Table
	__Table__000000_Table := (&models.Table{Name: `Table`}).Stage(stage)

	// Setup of values

	// Cell values setup
	__Cell__000000_Row_1_Cell_1.Name = `Row 1 - Cell 1`

	// Cell values setup
	__Cell__000001_Row_1_Cell_2.Name = `Row 1 - Cell 2`

	// Cell values setup
	__Cell__000002_Row_1_Cell_3.Name = `Row 1 - Cell 3`

	// Cell values setup
	__Cell__000003_Row_1_Cell_4.Name = `Row 1 - Cell 4`

	// Cell values setup
	__Cell__000004_Row_2_Cell_1.Name = `Row 2 - Cell 1`

	// Cell values setup
	__Cell__000005_Row_2_Cell_2.Name = `Row 2 - Cell 2`

	// Cell values setup
	__Cell__000006_Row_2_Cell_3.Name = `Row 2 - Cell 3`

	// Cell values setup
	__Cell__000007_Row_2_Cell_4.Name = `Row 2 - Cell 4`

	// CellFloat64 values setup
	__CellFloat64__000000_Row_2_Cell_2.Name = `Row 2 - Cell 2`
	__CellFloat64__000000_Row_2_Cell_2.Value = 500.433333

	// CellIcon values setup
	__CellIcon__000000_Row_1_Cell_3.Name = `Row 1 - Cell 3`
	__CellIcon__000000_Row_1_Cell_3.Icon = `delete`

	// CellIcon values setup
	__CellIcon__000001_Row_2_Cell_3.Name = `Row 2 - Cell 3`
	__CellIcon__000001_Row_2_Cell_3.Icon = `home`

	// CellInt values setup
	__CellInt__000000_Row_1_Cell_2.Name = `Row 1 - Cell 2`
	__CellInt__000000_Row_1_Cell_2.Value = 10

	// CellString values setup
	__CellString__000000_Row_1_Cell_1_.Name = `Row 1 - Cell 1 `

	// CellString values setup
	__CellString__000001_Row_1_Cell_4.Name = `Row 1 - Cell 4`

	// CellString values setup
	__CellString__000002_Row_2_Cell_1_.Name = `Row 2 - Cell 1 `

	// CellString values setup
	__CellString__000003_Row_2_Cell_4.Name = `Row 2 - Cell 4`

	// DisplayedColumn values setup
	__DisplayedColumn__000000_Column_1_String.Name = `Column 1 - String`

	// DisplayedColumn values setup
	__DisplayedColumn__000001_Column_2_Int.Name = `Column 2 - Int`

	// DisplayedColumn values setup
	__DisplayedColumn__000002_Column_3_Icon.Name = `Column 3 - Icon`

	// DisplayedColumn values setup
	__DisplayedColumn__000003_Column_4.Name = `Column 4`

	// Row values setup
	__Row__000000_Row_1.Name = `Row 1`

	// Row values setup
	__Row__000001_Row_2.Name = `Row 2`

	// Table values setup
	__Table__000000_Table.Name = `Table`

	// Setup of pointers
	__Cell__000000_Row_1_Cell_1.CellString = __CellString__000000_Row_1_Cell_1_
	__Cell__000001_Row_1_Cell_2.CellInt = __CellInt__000000_Row_1_Cell_2
	__Cell__000002_Row_1_Cell_3.CellIcon = __CellIcon__000000_Row_1_Cell_3
	__Cell__000003_Row_1_Cell_4.CellString = __CellString__000001_Row_1_Cell_4
	__Cell__000004_Row_2_Cell_1.CellString = __CellString__000002_Row_2_Cell_1_
	__Cell__000005_Row_2_Cell_2.CellFloat64 = __CellFloat64__000000_Row_2_Cell_2
	__Cell__000006_Row_2_Cell_3.CellIcon = __CellIcon__000001_Row_2_Cell_3
	__Cell__000007_Row_2_Cell_4.CellString = __CellString__000003_Row_2_Cell_4
	__Row__000000_Row_1.Cells = append(__Row__000000_Row_1.Cells, __Cell__000000_Row_1_Cell_1)
	__Row__000000_Row_1.Cells = append(__Row__000000_Row_1.Cells, __Cell__000001_Row_1_Cell_2)
	__Row__000000_Row_1.Cells = append(__Row__000000_Row_1.Cells, __Cell__000002_Row_1_Cell_3)
	__Row__000000_Row_1.Cells = append(__Row__000000_Row_1.Cells, __Cell__000003_Row_1_Cell_4)
	__Row__000001_Row_2.Cells = append(__Row__000001_Row_2.Cells, __Cell__000004_Row_2_Cell_1)
	__Row__000001_Row_2.Cells = append(__Row__000001_Row_2.Cells, __Cell__000005_Row_2_Cell_2)
	__Row__000001_Row_2.Cells = append(__Row__000001_Row_2.Cells, __Cell__000006_Row_2_Cell_3)
	__Row__000001_Row_2.Cells = append(__Row__000001_Row_2.Cells, __Cell__000007_Row_2_Cell_4)
	__Table__000000_Table.DisplayedColumns = append(__Table__000000_Table.DisplayedColumns, __DisplayedColumn__000002_Column_3_Icon)
	__Table__000000_Table.DisplayedColumns = append(__Table__000000_Table.DisplayedColumns, __DisplayedColumn__000003_Column_4)
	__Table__000000_Table.DisplayedColumns = append(__Table__000000_Table.DisplayedColumns, __DisplayedColumn__000001_Column_2_Int)
	__Table__000000_Table.DisplayedColumns = append(__Table__000000_Table.DisplayedColumns, __DisplayedColumn__000000_Column_1_String)
	__Table__000000_Table.Rows = append(__Table__000000_Table.Rows, __Row__000000_Row_1)
	__Table__000000_Table.Rows = append(__Table__000000_Table.Rows, __Row__000001_Row_2)
}


