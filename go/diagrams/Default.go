package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongtable/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_Default models.StageStruct
var ___dummy__Time_Default time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_Default ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_Default map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"ref_models.Cell": &(ref_models.Cell{}),

	"ref_models.Cell.CellString": (ref_models.Cell{}).CellString,

	"ref_models.Cell.Name": (ref_models.Cell{}).Name,

	"ref_models.CellString": &(ref_models.CellString{}),

	"ref_models.CellString.Name": (ref_models.CellString{}).Name,

	"ref_models.DisplayedColumn": &(ref_models.DisplayedColumn{}),

	"ref_models.DisplayedColumn.Name": (ref_models.DisplayedColumn{}).Name,

	"ref_models.Row": &(ref_models.Row{}),

	"ref_models.Row.Cells": (ref_models.Row{}).Cells,

	"ref_models.Row.Name": (ref_models.Row{}).Name,

	"ref_models.Table": &(ref_models.Table{}),

	"ref_models.Table.DisplayedColumns": (ref_models.Table{}).DisplayedColumns,

	"ref_models.Table.Name": (ref_models.Table{}).Name,

	"ref_models.Table.Rows": (ref_models.Table{}).Rows,
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["Default"] = DefaultInjection
// }

// DefaultInjection will stage objects of database "Default"
func DefaultInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Default := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_Default_Cell := (&models.GongStructShape{Name: `Default-Cell`}).Stage(stage)
	__GongStructShape__000001_Default_CellString := (&models.GongStructShape{Name: `Default-CellString`}).Stage(stage)
	__GongStructShape__000002_Default_DisplayedColumn := (&models.GongStructShape{Name: `Default-DisplayedColumn`}).Stage(stage)
	__GongStructShape__000003_Default_Row := (&models.GongStructShape{Name: `Default-Row`}).Stage(stage)
	__GongStructShape__000004_Default_Table := (&models.GongStructShape{Name: `Default-Table`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_CellString := (&models.Link{Name: `CellString`}).Stage(stage)
	__Link__000001_Cells := (&models.Link{Name: `Cells`}).Stage(stage)
	__Link__000002_DisplayedColumns := (&models.Link{Name: `DisplayedColumns`}).Stage(stage)
	__Link__000003_Rows := (&models.Link{Name: `Rows`}).Stage(stage)

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_Default_Cell := (&models.Position{Name: `Pos-Default-Cell`}).Stage(stage)
	__Position__000001_Pos_Default_CellString := (&models.Position{Name: `Pos-Default-CellString`}).Stage(stage)
	__Position__000002_Pos_Default_DisplayedColumn := (&models.Position{Name: `Pos-Default-DisplayedColumn`}).Stage(stage)
	__Position__000003_Pos_Default_Row := (&models.Position{Name: `Pos-Default-Row`}).Stage(stage)
	__Position__000004_Pos_Default_Table := (&models.Position{Name: `Pos-Default-Table`}).Stage(stage)

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellString := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Cell and Default-CellString`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Row_and_Default_Cell := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Row and Default-Cell`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Table_and_Default_DisplayedColumn := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Table and Default-DisplayedColumn`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Table_and_Default_Row := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Table and Default-Row`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = true

	// GongStructShape values setup
	__GongStructShape__000000_Default_Cell.Name = `Default-Cell`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Cell]
	__GongStructShape__000000_Default_Cell.Identifier = `ref_models.Cell`
	__GongStructShape__000000_Default_Cell.ShowNbInstances = true
	__GongStructShape__000000_Default_Cell.NbInstances = 4
	__GongStructShape__000000_Default_Cell.Width = 240.000000
	__GongStructShape__000000_Default_Cell.Heigth = 63.000000
	__GongStructShape__000000_Default_Cell.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_Default_CellString.Name = `Default-CellString`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CellString]
	__GongStructShape__000001_Default_CellString.Identifier = `ref_models.CellString`
	__GongStructShape__000001_Default_CellString.ShowNbInstances = true
	__GongStructShape__000001_Default_CellString.NbInstances = 4
	__GongStructShape__000001_Default_CellString.Width = 240.000000
	__GongStructShape__000001_Default_CellString.Heigth = 63.000000
	__GongStructShape__000001_Default_CellString.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_Default_DisplayedColumn.Name = `Default-DisplayedColumn`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DisplayedColumn]
	__GongStructShape__000002_Default_DisplayedColumn.Identifier = `ref_models.DisplayedColumn`
	__GongStructShape__000002_Default_DisplayedColumn.ShowNbInstances = false
	__GongStructShape__000002_Default_DisplayedColumn.NbInstances = 0
	__GongStructShape__000002_Default_DisplayedColumn.Width = 240.000000
	__GongStructShape__000002_Default_DisplayedColumn.Heigth = 63.000000
	__GongStructShape__000002_Default_DisplayedColumn.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000003_Default_Row.Name = `Default-Row`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Row]
	__GongStructShape__000003_Default_Row.Identifier = `ref_models.Row`
	__GongStructShape__000003_Default_Row.ShowNbInstances = true
	__GongStructShape__000003_Default_Row.NbInstances = 2
	__GongStructShape__000003_Default_Row.Width = 240.000000
	__GongStructShape__000003_Default_Row.Heigth = 63.000000
	__GongStructShape__000003_Default_Row.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000004_Default_Table.Name = `Default-Table`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Table]
	__GongStructShape__000004_Default_Table.Identifier = `ref_models.Table`
	__GongStructShape__000004_Default_Table.ShowNbInstances = true
	__GongStructShape__000004_Default_Table.NbInstances = 1
	__GongStructShape__000004_Default_Table.Width = 240.000000
	__GongStructShape__000004_Default_Table.Heigth = 63.000000
	__GongStructShape__000004_Default_Table.IsSelected = false

	// Link values setup
	__Link__000000_CellString.Name = `CellString`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Cell.CellString]
	__Link__000000_CellString.Identifier = `ref_models.Cell.CellString`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CellString]
	__Link__000000_CellString.Fieldtypename = `ref_models.CellString`
	__Link__000000_CellString.FieldOffsetX = -84.000000
	__Link__000000_CellString.FieldOffsetY = -22.000000
	__Link__000000_CellString.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_CellString.TargetMultiplicityOffsetX = -49.000000
	__Link__000000_CellString.TargetMultiplicityOffsetY = 26.000000
	__Link__000000_CellString.SourceMultiplicity = models.MANY
	__Link__000000_CellString.SourceMultiplicityOffsetX = 9.000000
	__Link__000000_CellString.SourceMultiplicityOffsetY = -13.000000
	__Link__000000_CellString.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_CellString.StartRatio = 0.500000
	__Link__000000_CellString.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_CellString.EndRatio = 0.500000
	__Link__000000_CellString.CornerOffsetRatio = 1.205278

	// Link values setup
	__Link__000001_Cells.Name = `Cells`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Row.Cells]
	__Link__000001_Cells.Identifier = `ref_models.Row.Cells`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Cell]
	__Link__000001_Cells.Fieldtypename = `ref_models.Cell`
	__Link__000001_Cells.FieldOffsetX = -50.000000
	__Link__000001_Cells.FieldOffsetY = -16.000000
	__Link__000001_Cells.TargetMultiplicity = models.MANY
	__Link__000001_Cells.TargetMultiplicityOffsetX = -33.000000
	__Link__000001_Cells.TargetMultiplicityOffsetY = 23.000000
	__Link__000001_Cells.SourceMultiplicity = models.ZERO_ONE
	__Link__000001_Cells.SourceMultiplicityOffsetX = 15.000000
	__Link__000001_Cells.SourceMultiplicityOffsetY = -11.000000
	__Link__000001_Cells.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Cells.StartRatio = 0.500000
	__Link__000001_Cells.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Cells.EndRatio = 0.500000
	__Link__000001_Cells.CornerOffsetRatio = 1.121944

	// Link values setup
	__Link__000002_DisplayedColumns.Name = `DisplayedColumns`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Table.DisplayedColumns]
	__Link__000002_DisplayedColumns.Identifier = `ref_models.Table.DisplayedColumns`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DisplayedColumn]
	__Link__000002_DisplayedColumns.Fieldtypename = `ref_models.DisplayedColumn`
	__Link__000002_DisplayedColumns.FieldOffsetX = -153.000000
	__Link__000002_DisplayedColumns.FieldOffsetY = -15.000000
	__Link__000002_DisplayedColumns.TargetMultiplicity = models.MANY
	__Link__000002_DisplayedColumns.TargetMultiplicityOffsetX = -29.000000
	__Link__000002_DisplayedColumns.TargetMultiplicityOffsetY = 21.000000
	__Link__000002_DisplayedColumns.SourceMultiplicity = models.ZERO_ONE
	__Link__000002_DisplayedColumns.SourceMultiplicityOffsetX = 8.000000
	__Link__000002_DisplayedColumns.SourceMultiplicityOffsetY = 24.000000
	__Link__000002_DisplayedColumns.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_DisplayedColumns.StartRatio = 0.412698
	__Link__000002_DisplayedColumns.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_DisplayedColumns.EndRatio = 0.492063
	__Link__000002_DisplayedColumns.CornerOffsetRatio = 1.195833

	// Link values setup
	__Link__000003_Rows.Name = `Rows`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Table.Rows]
	__Link__000003_Rows.Identifier = `ref_models.Table.Rows`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Row]
	__Link__000003_Rows.Fieldtypename = `ref_models.Row`
	__Link__000003_Rows.FieldOffsetX = 19.000000
	__Link__000003_Rows.FieldOffsetY = -17.000000
	__Link__000003_Rows.TargetMultiplicity = models.MANY
	__Link__000003_Rows.TargetMultiplicityOffsetX = -24.000000
	__Link__000003_Rows.TargetMultiplicityOffsetY = -9.000000
	__Link__000003_Rows.SourceMultiplicity = models.ZERO_ONE
	__Link__000003_Rows.SourceMultiplicityOffsetX = -43.000000
	__Link__000003_Rows.SourceMultiplicityOffsetY = 16.000000
	__Link__000003_Rows.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000003_Rows.StartRatio = 0.377507
	__Link__000003_Rows.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000003_Rows.EndRatio = 0.356673
	__Link__000003_Rows.CornerOffsetRatio = 1.746032

	// Position values setup
	__Position__000000_Pos_Default_Cell.X = 538.000000
	__Position__000000_Pos_Default_Cell.Y = 269.000000
	__Position__000000_Pos_Default_Cell.Name = `Pos-Default-Cell`

	// Position values setup
	__Position__000001_Pos_Default_CellString.X = 953.999939
	__Position__000001_Pos_Default_CellString.Y = 317.000000
	__Position__000001_Pos_Default_CellString.Name = `Pos-Default-CellString`

	// Position values setup
	__Position__000002_Pos_Default_DisplayedColumn.X = 536.000000
	__Position__000002_Pos_Default_DisplayedColumn.Y = 90.000000
	__Position__000002_Pos_Default_DisplayedColumn.Name = `Pos-Default-DisplayedColumn`

	// Position values setup
	__Position__000003_Pos_Default_Row.X = 21.000000
	__Position__000003_Pos_Default_Row.Y = 271.000000
	__Position__000003_Pos_Default_Row.Name = `Pos-Default-Row`

	// Position values setup
	__Position__000004_Pos_Default_Table.X = 17.000000
	__Position__000004_Pos_Default_Table.Y = 95.000000
	__Position__000004_Pos_Default_Table.Name = `Pos-Default-Table`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellString.X = 1144.499969
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellString.Y = 287.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellString.Name = `Verticle in class diagram Default in middle between Default-Cell and Default-CellString`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Row_and_Default_Cell.X = 831.000000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Row_and_Default_Cell.Y = 232.500000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Row_and_Default_Cell.Name = `Verticle in class diagram Default in middle between Default-Row and Default-Cell`

	// Vertice values setup
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Table_and_Default_DisplayedColumn.X = 373.500000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Table_and_Default_DisplayedColumn.Y = 237.000000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Table_and_Default_DisplayedColumn.Name = `Verticle in class diagram Default in middle between Default-Table and Default-DisplayedColumn`

	// Vertice values setup
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Table_and_Default_Row.X = 568.000000
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Table_and_Default_Row.Y = 182.000000
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Table_and_Default_Row.Name = `Verticle in class diagram Default in middle between Default-Table and Default-Row`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000004_Default_Table)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000003_Default_Row)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_CellString)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Cell)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_DisplayedColumn)
	__GongStructShape__000000_Default_Cell.Position = __Position__000000_Pos_Default_Cell
	__GongStructShape__000000_Default_Cell.Links = append(__GongStructShape__000000_Default_Cell.Links, __Link__000000_CellString)
	__GongStructShape__000001_Default_CellString.Position = __Position__000001_Pos_Default_CellString
	__GongStructShape__000002_Default_DisplayedColumn.Position = __Position__000002_Pos_Default_DisplayedColumn
	__GongStructShape__000003_Default_Row.Position = __Position__000003_Pos_Default_Row
	__GongStructShape__000003_Default_Row.Links = append(__GongStructShape__000003_Default_Row.Links, __Link__000001_Cells)
	__GongStructShape__000004_Default_Table.Position = __Position__000004_Pos_Default_Table
	__GongStructShape__000004_Default_Table.Links = append(__GongStructShape__000004_Default_Table.Links, __Link__000003_Rows)
	__GongStructShape__000004_Default_Table.Links = append(__GongStructShape__000004_Default_Table.Links, __Link__000002_DisplayedColumns)
	__Link__000000_CellString.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellString
	__Link__000001_Cells.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Row_and_Default_Cell
	__Link__000002_DisplayedColumns.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Table_and_Default_DisplayedColumn
	__Link__000003_Rows.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Table_and_Default_Row
}


