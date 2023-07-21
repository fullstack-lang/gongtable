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
	__Cell__000002_Row_2_Cell_1 := (&models.Cell{Name: `Row 2 - Cell 1`}).Stage(stage)
	__Cell__000003_Row_2_Cell_2 := (&models.Cell{Name: `Row 2 - Cell 2`}).Stage(stage)

	// Declarations of staged instances of CellString
	__CellString__000000_Row_1_Cell_1_ := (&models.CellString{Name: `Row 1 - Cell 1 `}).Stage(stage)
	__CellString__000001_Row_1_Cell_2 := (&models.CellString{Name: `Row 1 - Cell 2`}).Stage(stage)
	__CellString__000002_Row_2_Cell_1_ := (&models.CellString{Name: `Row 2 - Cell 1 `}).Stage(stage)
	__CellString__000003_Row_2_Cell_2 := (&models.CellString{Name: `Row 2 - Cell 2`}).Stage(stage)

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
	__Cell__000002_Row_2_Cell_1.Name = `Row 2 - Cell 1`

	// Cell values setup
	__Cell__000003_Row_2_Cell_2.Name = `Row 2 - Cell 2`

	// CellString values setup
	__CellString__000000_Row_1_Cell_1_.Name = `Row 1 - Cell 1 `

	// CellString values setup
	__CellString__000001_Row_1_Cell_2.Name = `Row 1 - Cell 2`

	// CellString values setup
	__CellString__000002_Row_2_Cell_1_.Name = `Row 2 - Cell 1 `

	// CellString values setup
	__CellString__000003_Row_2_Cell_2.Name = `Row 2 - Cell 2`

	// Row values setup
	__Row__000000_Row_1.Name = `Row 1`

	// Row values setup
	__Row__000001_Row_2.Name = `Row 2`

	// Table values setup
	__Table__000000_Table.Name = `Table`

	// Setup of pointers
	__Cell__000000_Row_1_Cell_1.CellString = __CellString__000000_Row_1_Cell_1_
	__Cell__000001_Row_1_Cell_2.CellString = __CellString__000001_Row_1_Cell_2
	__Cell__000002_Row_2_Cell_1.CellString = __CellString__000002_Row_2_Cell_1_
	__Cell__000003_Row_2_Cell_2.CellString = __CellString__000003_Row_2_Cell_2
	__Row__000000_Row_1.Cells = append(__Row__000000_Row_1.Cells, __Cell__000000_Row_1_Cell_1)
	__Row__000000_Row_1.Cells = append(__Row__000000_Row_1.Cells, __Cell__000001_Row_1_Cell_2)
	__Row__000001_Row_2.Cells = append(__Row__000001_Row_2.Cells, __Cell__000002_Row_2_Cell_1)
	__Row__000001_Row_2.Cells = append(__Row__000001_Row_2.Cells, __Cell__000003_Row_2_Cell_2)
	__Table__000000_Table.Rows = append(__Table__000000_Table.Rows, __Row__000000_Row_1)
	__Table__000000_Table.Rows = append(__Table__000000_Table.Rows, __Row__000001_Row_2)
}


