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
	__Cell__000004_Row_1_Cell_5 := (&models.Cell{Name: `Row 1 - Cell 5`}).Stage(stage)
	__Cell__000005_Row_2_Cell_1 := (&models.Cell{Name: `Row 2 - Cell 1`}).Stage(stage)
	__Cell__000006_Row_2_Cell_2 := (&models.Cell{Name: `Row 2 - Cell 2`}).Stage(stage)
	__Cell__000007_Row_2_Cell_3 := (&models.Cell{Name: `Row 2 - Cell 3`}).Stage(stage)
	__Cell__000008_Row_2_Cell_4 := (&models.Cell{Name: `Row 2 - Cell 4`}).Stage(stage)
	__Cell__000009_Row_2_Cell_5 := (&models.Cell{Name: `Row 2 - Cell 5`}).Stage(stage)

	// Declarations of staged instances of CellBoolean
	__CellBoolean__000000_Row_1_Cell_1_Cell_False := (&models.CellBoolean{Name: `Row 1 -Cell 1 - Cell False`}).Stage(stage)
	__CellBoolean__000001_Row_2_Cell_1_Cell_true := (&models.CellBoolean{Name: `Row 2 - Cell 1 - Cell true`}).Stage(stage)

	// Declarations of staged instances of CellFloat64
	__CellFloat64__000000_Row_1_Cell2_Float := (&models.CellFloat64{Name: `Row 1 - Cell2 - Float`}).Stage(stage)
	__CellFloat64__000001_Row_2_Cell_2_Float := (&models.CellFloat64{Name: `Row 2 - Cell 2 - Float`}).Stage(stage)

	// Declarations of staged instances of CellIcon
	__CellIcon__000000_Row_1_Cell_3_Delete := (&models.CellIcon{Name: `Row 1 - Cell 3 - Delete`}).Stage(stage)
	__CellIcon__000001_Row_2_Cell_3_home := (&models.CellIcon{Name: `Row 2 - Cell 3 - home`}).Stage(stage)

	// Declarations of staged instances of CellInt
	__CellInt__000000_Row_1_Cell_4_Int := (&models.CellInt{Name: `Row 1 - Cell 4 - Int`}).Stage(stage)
	__CellInt__000001_Row_2_Cell_4_Int := (&models.CellInt{Name: `Row 2 - Cell 4 - Int`}).Stage(stage)

	// Declarations of staged instances of CellString
	__CellString__000000_Row_1_Cell_5 := (&models.CellString{Name: `Row 1 - Cell 5`}).Stage(stage)
	__CellString__000001_Row_2_Cell_5 := (&models.CellString{Name: `Row 2 - Cell 5`}).Stage(stage)

	// Declarations of staged instances of DisplayedColumn
	__DisplayedColumn__000000_Column_1_Boolean := (&models.DisplayedColumn{Name: `Column 1 - Boolean`}).Stage(stage)
	__DisplayedColumn__000001_Column_2_Float64 := (&models.DisplayedColumn{Name: `Column 2 - Float64`}).Stage(stage)
	__DisplayedColumn__000002_Column_3_Icon := (&models.DisplayedColumn{Name: `Column 3 - Icon`}).Stage(stage)
	__DisplayedColumn__000003_Column_4_Int := (&models.DisplayedColumn{Name: `Column 4 - Int`}).Stage(stage)
	__DisplayedColumn__000004_Column_5_String := (&models.DisplayedColumn{Name: `Column 5 - String`}).Stage(stage)

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
	__Cell__000004_Row_1_Cell_5.Name = `Row 1 - Cell 5`

	// Cell values setup
	__Cell__000005_Row_2_Cell_1.Name = `Row 2 - Cell 1`

	// Cell values setup
	__Cell__000006_Row_2_Cell_2.Name = `Row 2 - Cell 2`

	// Cell values setup
	__Cell__000007_Row_2_Cell_3.Name = `Row 2 - Cell 3`

	// Cell values setup
	__Cell__000008_Row_2_Cell_4.Name = `Row 2 - Cell 4`

	// Cell values setup
	__Cell__000009_Row_2_Cell_5.Name = `Row 2 - Cell 5`

	// CellBoolean values setup
	__CellBoolean__000000_Row_1_Cell_1_Cell_False.Name = `Row 1 -Cell 1 - Cell False`
	__CellBoolean__000000_Row_1_Cell_1_Cell_False.Value = false

	// CellBoolean values setup
	__CellBoolean__000001_Row_2_Cell_1_Cell_true.Name = `Row 2 - Cell 1 - Cell true`
	__CellBoolean__000001_Row_2_Cell_1_Cell_true.Value = true

	// CellFloat64 values setup
	__CellFloat64__000000_Row_1_Cell2_Float.Name = `Row 1 - Cell2 - Float`
	__CellFloat64__000000_Row_1_Cell2_Float.Value = 20.433333

	// CellFloat64 values setup
	__CellFloat64__000001_Row_2_Cell_2_Float.Name = `Row 2 - Cell 2 - Float`
	__CellFloat64__000001_Row_2_Cell_2_Float.Value = 18.550000

	// CellIcon values setup
	__CellIcon__000000_Row_1_Cell_3_Delete.Name = `Row 1 - Cell 3 - Delete`
	__CellIcon__000000_Row_1_Cell_3_Delete.Icon = `delete`

	// CellIcon values setup
	__CellIcon__000001_Row_2_Cell_3_home.Name = `Row 2 - Cell 3 - home`
	__CellIcon__000001_Row_2_Cell_3_home.Icon = `home`

	// CellInt values setup
	__CellInt__000000_Row_1_Cell_4_Int.Name = `Row 1 - Cell 4 - Int`
	__CellInt__000000_Row_1_Cell_4_Int.Value = 10

	// CellInt values setup
	__CellInt__000001_Row_2_Cell_4_Int.Name = `Row 2 - Cell 4 - Int`
	__CellInt__000001_Row_2_Cell_4_Int.Value = 288

	// CellString values setup
	__CellString__000000_Row_1_Cell_5.Name = `Row 1 - Cell 5`
	__CellString__000000_Row_1_Cell_5.Value = `Je ferais le métier`

	// CellString values setup
	__CellString__000001_Row_2_Cell_5.Name = `Row 2 - Cell 5`
	__CellString__000001_Row_2_Cell_5.Value = `des idoles antiques`

	// DisplayedColumn values setup
	__DisplayedColumn__000000_Column_1_Boolean.Name = `Column 1 - Boolean`

	// DisplayedColumn values setup
	__DisplayedColumn__000001_Column_2_Float64.Name = `Column 2 - Float64`

	// DisplayedColumn values setup
	__DisplayedColumn__000002_Column_3_Icon.Name = `Column 3 - Icon`

	// DisplayedColumn values setup
	__DisplayedColumn__000003_Column_4_Int.Name = `Column 4 - Int`

	// DisplayedColumn values setup
	__DisplayedColumn__000004_Column_5_String.Name = `Column 5 - String`

	// Row values setup
	__Row__000000_Row_1.Name = `Row 1`

	// Row values setup
	__Row__000001_Row_2.Name = `Row 2`

	// Table values setup
	__Table__000000_Table.Name = `Table`

	// Setup of pointers
	__Cell__000000_Row_1_Cell_1.CellBool = __CellBoolean__000000_Row_1_Cell_1_Cell_False
	__Cell__000001_Row_1_Cell_2.CellFloat64 = __CellFloat64__000000_Row_1_Cell2_Float
	__Cell__000002_Row_1_Cell_3.CellIcon = __CellIcon__000000_Row_1_Cell_3_Delete
	__Cell__000003_Row_1_Cell_4.CellInt = __CellInt__000000_Row_1_Cell_4_Int
	__Cell__000004_Row_1_Cell_5.CellString = __CellString__000000_Row_1_Cell_5
	__Cell__000005_Row_2_Cell_1.CellBool = __CellBoolean__000001_Row_2_Cell_1_Cell_true
	__Cell__000006_Row_2_Cell_2.CellFloat64 = __CellFloat64__000001_Row_2_Cell_2_Float
	__Cell__000007_Row_2_Cell_3.CellIcon = __CellIcon__000001_Row_2_Cell_3_home
	__Cell__000008_Row_2_Cell_4.CellInt = __CellInt__000001_Row_2_Cell_4_Int
	__Cell__000009_Row_2_Cell_5.CellString = __CellString__000001_Row_2_Cell_5
	__Row__000000_Row_1.Cells = append(__Row__000000_Row_1.Cells, __Cell__000000_Row_1_Cell_1)
	__Row__000000_Row_1.Cells = append(__Row__000000_Row_1.Cells, __Cell__000001_Row_1_Cell_2)
	__Row__000000_Row_1.Cells = append(__Row__000000_Row_1.Cells, __Cell__000002_Row_1_Cell_3)
	__Row__000000_Row_1.Cells = append(__Row__000000_Row_1.Cells, __Cell__000003_Row_1_Cell_4)
	__Row__000000_Row_1.Cells = append(__Row__000000_Row_1.Cells, __Cell__000004_Row_1_Cell_5)
	__Row__000001_Row_2.Cells = append(__Row__000001_Row_2.Cells, __Cell__000005_Row_2_Cell_1)
	__Row__000001_Row_2.Cells = append(__Row__000001_Row_2.Cells, __Cell__000006_Row_2_Cell_2)
	__Row__000001_Row_2.Cells = append(__Row__000001_Row_2.Cells, __Cell__000007_Row_2_Cell_3)
	__Row__000001_Row_2.Cells = append(__Row__000001_Row_2.Cells, __Cell__000008_Row_2_Cell_4)
	__Row__000001_Row_2.Cells = append(__Row__000001_Row_2.Cells, __Cell__000009_Row_2_Cell_5)
	__Table__000000_Table.DisplayedColumns = append(__Table__000000_Table.DisplayedColumns, __DisplayedColumn__000004_Column_5_String)
	__Table__000000_Table.DisplayedColumns = append(__Table__000000_Table.DisplayedColumns, __DisplayedColumn__000002_Column_3_Icon)
	__Table__000000_Table.DisplayedColumns = append(__Table__000000_Table.DisplayedColumns, __DisplayedColumn__000003_Column_4_Int)
	__Table__000000_Table.DisplayedColumns = append(__Table__000000_Table.DisplayedColumns, __DisplayedColumn__000001_Column_2_Float64)
	__Table__000000_Table.DisplayedColumns = append(__Table__000000_Table.DisplayedColumns, __DisplayedColumn__000000_Column_1_Boolean)
	__Table__000000_Table.Rows = append(__Table__000000_Table.Rows, __Row__000000_Row_1)
	__Table__000000_Table.Rows = append(__Table__000000_Table.Rows, __Row__000001_Row_2)
}


