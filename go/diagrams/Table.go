package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongtable/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_Table models.StageStruct
var ___dummy__Time_Table time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_Table ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_Table map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"ref_models.Cell": &(ref_models.Cell{}),

	"ref_models.Cell.CellBool": (ref_models.Cell{}).CellBool,

	"ref_models.Cell.CellFloat64": (ref_models.Cell{}).CellFloat64,

	"ref_models.Cell.CellIcon": (ref_models.Cell{}).CellIcon,

	"ref_models.Cell.CellInt": (ref_models.Cell{}).CellInt,

	"ref_models.Cell.CellString": (ref_models.Cell{}).CellString,

	"ref_models.Cell.Name": (ref_models.Cell{}).Name,

	"ref_models.CellBoolean": &(ref_models.CellBoolean{}),

	"ref_models.CellBoolean.Name": (ref_models.CellBoolean{}).Name,

	"ref_models.CellBoolean.Value": (ref_models.CellBoolean{}).Value,

	"ref_models.CellFloat64": &(ref_models.CellFloat64{}),

	"ref_models.CellFloat64.Name": (ref_models.CellFloat64{}).Name,

	"ref_models.CellFloat64.Value": (ref_models.CellFloat64{}).Value,

	"ref_models.CellIcon": &(ref_models.CellIcon{}),

	"ref_models.CellIcon.Icon": (ref_models.CellIcon{}).Icon,

	"ref_models.CellIcon.Name": (ref_models.CellIcon{}).Name,

	"ref_models.CellInt": &(ref_models.CellInt{}),

	"ref_models.CellInt.Name": (ref_models.CellInt{}).Name,

	"ref_models.CellInt.Value": (ref_models.CellInt{}).Value,

	"ref_models.CellString": &(ref_models.CellString{}),

	"ref_models.CellString.Name": (ref_models.CellString{}).Name,

	"ref_models.CellString.Value": (ref_models.CellString{}).Value,

	"ref_models.DisplayedColumn": &(ref_models.DisplayedColumn{}),

	"ref_models.DisplayedColumn.Name": (ref_models.DisplayedColumn{}).Name,

	"ref_models.Form": &(ref_models.FormGroup{}),

	"ref_models.Form.FormCells": (ref_models.FormGroup{}).FormFields,

	"ref_models.Form.Name": (ref_models.FormGroup{}).Name,

	"ref_models.FormCell": &(ref_models.FormField{}),

	"ref_models.FormCell.FormCellBool": (ref_models.FormField{}).FormFieldBool,

	"ref_models.FormCell.FormCellFloat64": (ref_models.FormField{}).FormFieldFloat64,

	"ref_models.FormCell.FormCellInt": (ref_models.FormField{}).FormFieldInt,

	"ref_models.FormCell.FormCellString": (ref_models.FormField{}).FormFieldString,

	"ref_models.FormCell.Name": (ref_models.FormField{}).Name,

	"ref_models.FormCellBoolean": &(ref_models.FormFieldBoolean{}),

	"ref_models.FormCellBoolean.Name": (ref_models.FormFieldBoolean{}).Name,

	"ref_models.FormCellBoolean.Value": (ref_models.FormFieldBoolean{}).Value,

	"ref_models.FormCellFloat64": &(ref_models.FormFieldFloat64{}),

	"ref_models.FormCellFloat64.Name": (ref_models.FormFieldFloat64{}).Name,

	"ref_models.FormCellFloat64.Value": (ref_models.FormFieldFloat64{}).Value,

	"ref_models.FormCellInt": &(ref_models.FormFieldInt{}),

	"ref_models.FormCellInt.Name": (ref_models.FormFieldInt{}).Name,

	"ref_models.FormCellInt.Value": (ref_models.FormFieldInt{}).Value,

	"ref_models.FormCellString": &(ref_models.FormFieldString{}),

	"ref_models.FormCellString.Name": (ref_models.FormFieldString{}).Name,

	"ref_models.FormCellString.Value": (ref_models.FormFieldString{}).Value,

	"ref_models.Row": &(ref_models.Row{}),

	"ref_models.Row.Cells": (ref_models.Row{}).Cells,

	"ref_models.Row.IsChecked": (ref_models.Row{}).IsChecked,

	"ref_models.Row.Name": (ref_models.Row{}).Name,

	"ref_models.Table": &(ref_models.Table{}),

	"ref_models.Table.CanDragDropRows": (ref_models.Table{}).CanDragDropRows,

	"ref_models.Table.DisplayedColumns": (ref_models.Table{}).DisplayedColumns,

	"ref_models.Table.HasCheckableRows": (ref_models.Table{}).HasCheckableRows,

	"ref_models.Table.HasColumnSorting": (ref_models.Table{}).HasColumnSorting,

	"ref_models.Table.HasFiltering": (ref_models.Table{}).HasFiltering,

	"ref_models.Table.HasPaginator": (ref_models.Table{}).HasPaginator,

	"ref_models.Table.HasSaveButton": (ref_models.Table{}).HasSaveButton,

	"ref_models.Table.Name": (ref_models.Table{}).Name,

	"ref_models.Table.Rows": (ref_models.Table{}).Rows,
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["Table"] = TableInjection
// }

// TableInjection will stage objects of database "Table"
func TableInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Table := (&models.Classdiagram{Name: `Table`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_CanDragDropRows := (&models.Field{Name: `CanDragDropRows`}).Stage(stage)
	__Field__000001_HasCheckableRows := (&models.Field{Name: `HasCheckableRows`}).Stage(stage)
	__Field__000002_HasColumnSorting := (&models.Field{Name: `HasColumnSorting`}).Stage(stage)
	__Field__000003_HasFiltering := (&models.Field{Name: `HasFiltering`}).Stage(stage)
	__Field__000004_HasPaginator := (&models.Field{Name: `HasPaginator`}).Stage(stage)
	__Field__000005_HasSaveButton := (&models.Field{Name: `HasSaveButton`}).Stage(stage)
	__Field__000006_Icon := (&models.Field{Name: `Icon`}).Stage(stage)
	__Field__000007_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000008_Value := (&models.Field{Name: `Value`}).Stage(stage)
	__Field__000009_Value := (&models.Field{Name: `Value`}).Stage(stage)
	__Field__000010_Value := (&models.Field{Name: `Value`}).Stage(stage)
	__Field__000011_Value := (&models.Field{Name: `Value`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_Default_Cell := (&models.GongStructShape{Name: `Default-Cell`}).Stage(stage)
	__GongStructShape__000001_Default_CellBoolean := (&models.GongStructShape{Name: `Default-CellBoolean`}).Stage(stage)
	__GongStructShape__000002_Default_CellFloat64 := (&models.GongStructShape{Name: `Default-CellFloat64`}).Stage(stage)
	__GongStructShape__000003_Default_CellIcon := (&models.GongStructShape{Name: `Default-CellIcon`}).Stage(stage)
	__GongStructShape__000004_Default_CellInt := (&models.GongStructShape{Name: `Default-CellInt`}).Stage(stage)
	__GongStructShape__000005_Default_CellString := (&models.GongStructShape{Name: `Default-CellString`}).Stage(stage)
	__GongStructShape__000006_Default_DisplayedColumn := (&models.GongStructShape{Name: `Default-DisplayedColumn`}).Stage(stage)
	__GongStructShape__000007_Default_Row := (&models.GongStructShape{Name: `Default-Row`}).Stage(stage)
	__GongStructShape__000008_Default_Table := (&models.GongStructShape{Name: `Default-Table`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_CellBool := (&models.Link{Name: `CellBool`}).Stage(stage)
	__Link__000001_CellFloat64 := (&models.Link{Name: `CellFloat64`}).Stage(stage)
	__Link__000002_CellIcon := (&models.Link{Name: `CellIcon`}).Stage(stage)
	__Link__000003_CellInt := (&models.Link{Name: `CellInt`}).Stage(stage)
	__Link__000004_CellString := (&models.Link{Name: `CellString`}).Stage(stage)
	__Link__000005_Cells := (&models.Link{Name: `Cells`}).Stage(stage)
	__Link__000006_DisplayedColumns := (&models.Link{Name: `DisplayedColumns`}).Stage(stage)
	__Link__000007_Rows := (&models.Link{Name: `Rows`}).Stage(stage)

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_Default_Cell := (&models.Position{Name: `Pos-Default-Cell`}).Stage(stage)
	__Position__000001_Pos_Default_CellBoolean := (&models.Position{Name: `Pos-Default-CellBoolean`}).Stage(stage)
	__Position__000002_Pos_Default_CellFloat64 := (&models.Position{Name: `Pos-Default-CellFloat64`}).Stage(stage)
	__Position__000003_Pos_Default_CellIcon := (&models.Position{Name: `Pos-Default-CellIcon`}).Stage(stage)
	__Position__000004_Pos_Default_CellInt := (&models.Position{Name: `Pos-Default-CellInt`}).Stage(stage)
	__Position__000005_Pos_Default_CellString := (&models.Position{Name: `Pos-Default-CellString`}).Stage(stage)
	__Position__000006_Pos_Default_DisplayedColumn := (&models.Position{Name: `Pos-Default-DisplayedColumn`}).Stage(stage)
	__Position__000007_Pos_Default_Row := (&models.Position{Name: `Pos-Default-Row`}).Stage(stage)
	__Position__000008_Pos_Default_Table := (&models.Position{Name: `Pos-Default-Table`}).Stage(stage)

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellBoolean := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Cell and Default-CellBoolean`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellFloat64 := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Cell and Default-CellFloat64`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellIcon := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Cell and Default-CellIcon`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellInt := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Cell and Default-CellInt`}).Stage(stage)
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellString := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Cell and Default-CellString`}).Stage(stage)
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_Row_and_Default_Cell := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Row and Default-Cell`}).Stage(stage)
	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_Table_and_Default_DisplayedColumn := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Table and Default-DisplayedColumn`}).Stage(stage)
	__Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_Table_and_Default_Row := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Table and Default-Row`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Table.Name = `Table`
	__Classdiagram__000000_Table.IsInDrawMode = true

	// Field values setup
	__Field__000000_CanDragDropRows.Name = `CanDragDropRows`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Table.CanDragDropRows]
	__Field__000000_CanDragDropRows.Identifier = `ref_models.Table.CanDragDropRows`
	__Field__000000_CanDragDropRows.FieldTypeAsString = ``
	__Field__000000_CanDragDropRows.Structname = `Table`
	__Field__000000_CanDragDropRows.Fieldtypename = `bool`

	// Field values setup
	__Field__000001_HasCheckableRows.Name = `HasCheckableRows`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Table.HasCheckableRows]
	__Field__000001_HasCheckableRows.Identifier = `ref_models.Table.HasCheckableRows`
	__Field__000001_HasCheckableRows.FieldTypeAsString = ``
	__Field__000001_HasCheckableRows.Structname = `Table`
	__Field__000001_HasCheckableRows.Fieldtypename = `bool`

	// Field values setup
	__Field__000002_HasColumnSorting.Name = `HasColumnSorting`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Table.HasColumnSorting]
	__Field__000002_HasColumnSorting.Identifier = `ref_models.Table.HasColumnSorting`
	__Field__000002_HasColumnSorting.FieldTypeAsString = ``
	__Field__000002_HasColumnSorting.Structname = `Table`
	__Field__000002_HasColumnSorting.Fieldtypename = `bool`

	// Field values setup
	__Field__000003_HasFiltering.Name = `HasFiltering`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Table.HasFiltering]
	__Field__000003_HasFiltering.Identifier = `ref_models.Table.HasFiltering`
	__Field__000003_HasFiltering.FieldTypeAsString = ``
	__Field__000003_HasFiltering.Structname = `Table`
	__Field__000003_HasFiltering.Fieldtypename = `bool`

	// Field values setup
	__Field__000004_HasPaginator.Name = `HasPaginator`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Table.HasPaginator]
	__Field__000004_HasPaginator.Identifier = `ref_models.Table.HasPaginator`
	__Field__000004_HasPaginator.FieldTypeAsString = ``
	__Field__000004_HasPaginator.Structname = `Table`
	__Field__000004_HasPaginator.Fieldtypename = `bool`

	// Field values setup
	__Field__000005_HasSaveButton.Name = `HasSaveButton`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Table.HasSaveButton]
	__Field__000005_HasSaveButton.Identifier = `ref_models.Table.HasSaveButton`
	__Field__000005_HasSaveButton.FieldTypeAsString = ``
	__Field__000005_HasSaveButton.Structname = `Table`
	__Field__000005_HasSaveButton.Fieldtypename = `bool`

	// Field values setup
	__Field__000006_Icon.Name = `Icon`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CellIcon.Icon]
	__Field__000006_Icon.Identifier = `ref_models.CellIcon.Icon`
	__Field__000006_Icon.FieldTypeAsString = ``
	__Field__000006_Icon.Structname = `CellIcon`
	__Field__000006_Icon.Fieldtypename = `string`

	// Field values setup
	__Field__000007_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Table.Name]
	__Field__000007_Name.Identifier = `ref_models.Table.Name`
	__Field__000007_Name.FieldTypeAsString = ``
	__Field__000007_Name.Structname = `Table`
	__Field__000007_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000008_Value.Name = `Value`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CellInt.Value]
	__Field__000008_Value.Identifier = `ref_models.CellInt.Value`
	__Field__000008_Value.FieldTypeAsString = ``
	__Field__000008_Value.Structname = `CellInt`
	__Field__000008_Value.Fieldtypename = `int`

	// Field values setup
	__Field__000009_Value.Name = `Value`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CellString.Value]
	__Field__000009_Value.Identifier = `ref_models.CellString.Value`
	__Field__000009_Value.FieldTypeAsString = ``
	__Field__000009_Value.Structname = `CellString`
	__Field__000009_Value.Fieldtypename = `string`

	// Field values setup
	__Field__000010_Value.Name = `Value`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CellBoolean.Value]
	__Field__000010_Value.Identifier = `ref_models.CellBoolean.Value`
	__Field__000010_Value.FieldTypeAsString = ``
	__Field__000010_Value.Structname = `CellBoolean`
	__Field__000010_Value.Fieldtypename = `bool`

	// Field values setup
	__Field__000011_Value.Name = `Value`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CellFloat64.Value]
	__Field__000011_Value.Identifier = `ref_models.CellFloat64.Value`
	__Field__000011_Value.FieldTypeAsString = ``
	__Field__000011_Value.Structname = `CellFloat64`
	__Field__000011_Value.Fieldtypename = `float64`

	// GongStructShape values setup
	__GongStructShape__000000_Default_Cell.Name = `Default-Cell`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Cell]
	__GongStructShape__000000_Default_Cell.Identifier = `ref_models.Cell`
	__GongStructShape__000000_Default_Cell.ShowNbInstances = true
	__GongStructShape__000000_Default_Cell.NbInstances = 16
	__GongStructShape__000000_Default_Cell.Width = 240.000000
	__GongStructShape__000000_Default_Cell.Heigth = 63.000000
	__GongStructShape__000000_Default_Cell.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_Default_CellBoolean.Name = `Default-CellBoolean`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CellBoolean]
	__GongStructShape__000001_Default_CellBoolean.Identifier = `ref_models.CellBoolean`
	__GongStructShape__000001_Default_CellBoolean.ShowNbInstances = true
	__GongStructShape__000001_Default_CellBoolean.NbInstances = 2
	__GongStructShape__000001_Default_CellBoolean.Width = 240.000000
	__GongStructShape__000001_Default_CellBoolean.Heigth = 78.000000
	__GongStructShape__000001_Default_CellBoolean.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_Default_CellFloat64.Name = `Default-CellFloat64`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CellFloat64]
	__GongStructShape__000002_Default_CellFloat64.Identifier = `ref_models.CellFloat64`
	__GongStructShape__000002_Default_CellFloat64.ShowNbInstances = true
	__GongStructShape__000002_Default_CellFloat64.NbInstances = 2
	__GongStructShape__000002_Default_CellFloat64.Width = 240.000000
	__GongStructShape__000002_Default_CellFloat64.Heigth = 78.000000
	__GongStructShape__000002_Default_CellFloat64.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000003_Default_CellIcon.Name = `Default-CellIcon`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CellIcon]
	__GongStructShape__000003_Default_CellIcon.Identifier = `ref_models.CellIcon`
	__GongStructShape__000003_Default_CellIcon.ShowNbInstances = true
	__GongStructShape__000003_Default_CellIcon.NbInstances = 2
	__GongStructShape__000003_Default_CellIcon.Width = 240.000000
	__GongStructShape__000003_Default_CellIcon.Heigth = 78.000000
	__GongStructShape__000003_Default_CellIcon.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000004_Default_CellInt.Name = `Default-CellInt`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CellInt]
	__GongStructShape__000004_Default_CellInt.Identifier = `ref_models.CellInt`
	__GongStructShape__000004_Default_CellInt.ShowNbInstances = true
	__GongStructShape__000004_Default_CellInt.NbInstances = 2
	__GongStructShape__000004_Default_CellInt.Width = 240.000000
	__GongStructShape__000004_Default_CellInt.Heigth = 78.000000
	__GongStructShape__000004_Default_CellInt.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000005_Default_CellString.Name = `Default-CellString`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CellString]
	__GongStructShape__000005_Default_CellString.Identifier = `ref_models.CellString`
	__GongStructShape__000005_Default_CellString.ShowNbInstances = true
	__GongStructShape__000005_Default_CellString.NbInstances = 8
	__GongStructShape__000005_Default_CellString.Width = 240.000000
	__GongStructShape__000005_Default_CellString.Heigth = 78.000000
	__GongStructShape__000005_Default_CellString.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000006_Default_DisplayedColumn.Name = `Default-DisplayedColumn`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DisplayedColumn]
	__GongStructShape__000006_Default_DisplayedColumn.Identifier = `ref_models.DisplayedColumn`
	__GongStructShape__000006_Default_DisplayedColumn.ShowNbInstances = true
	__GongStructShape__000006_Default_DisplayedColumn.NbInstances = 7
	__GongStructShape__000006_Default_DisplayedColumn.Width = 240.000000
	__GongStructShape__000006_Default_DisplayedColumn.Heigth = 63.000000
	__GongStructShape__000006_Default_DisplayedColumn.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000007_Default_Row.Name = `Default-Row`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Row]
	__GongStructShape__000007_Default_Row.Identifier = `ref_models.Row`
	__GongStructShape__000007_Default_Row.ShowNbInstances = true
	__GongStructShape__000007_Default_Row.NbInstances = 5
	__GongStructShape__000007_Default_Row.Width = 240.000000
	__GongStructShape__000007_Default_Row.Heigth = 63.000000
	__GongStructShape__000007_Default_Row.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000008_Default_Table.Name = `Default-Table`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Table]
	__GongStructShape__000008_Default_Table.Identifier = `ref_models.Table`
	__GongStructShape__000008_Default_Table.ShowNbInstances = true
	__GongStructShape__000008_Default_Table.NbInstances = 2
	__GongStructShape__000008_Default_Table.Width = 240.000000
	__GongStructShape__000008_Default_Table.Heigth = 168.000000
	__GongStructShape__000008_Default_Table.IsSelected = false

	// Link values setup
	__Link__000000_CellBool.Name = `CellBool`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Cell.CellBool]
	__Link__000000_CellBool.Identifier = `ref_models.Cell.CellBool`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CellBoolean]
	__Link__000000_CellBool.Fieldtypename = `ref_models.CellBoolean`
	__Link__000000_CellBool.FieldOffsetX = -73.000000
	__Link__000000_CellBool.FieldOffsetY = -16.000000
	__Link__000000_CellBool.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_CellBool.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_CellBool.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_CellBool.SourceMultiplicity = models.MANY
	__Link__000000_CellBool.SourceMultiplicityOffsetX = 10.000000
	__Link__000000_CellBool.SourceMultiplicityOffsetY = -50.000000
	__Link__000000_CellBool.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_CellBool.StartRatio = 0.500000
	__Link__000000_CellBool.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_CellBool.EndRatio = 0.500000
	__Link__000000_CellBool.CornerOffsetRatio = 1.190007

	// Link values setup
	__Link__000001_CellFloat64.Name = `CellFloat64`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Cell.CellFloat64]
	__Link__000001_CellFloat64.Identifier = `ref_models.Cell.CellFloat64`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CellFloat64]
	__Link__000001_CellFloat64.Fieldtypename = `ref_models.CellFloat64`
	__Link__000001_CellFloat64.FieldOffsetX = -97.000000
	__Link__000001_CellFloat64.FieldOffsetY = -14.000000
	__Link__000001_CellFloat64.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_CellFloat64.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_CellFloat64.TargetMultiplicityOffsetY = 16.000000
	__Link__000001_CellFloat64.SourceMultiplicity = models.MANY
	__Link__000001_CellFloat64.SourceMultiplicityOffsetX = 10.000000
	__Link__000001_CellFloat64.SourceMultiplicityOffsetY = -50.000000
	__Link__000001_CellFloat64.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_CellFloat64.StartRatio = 0.500000
	__Link__000001_CellFloat64.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_CellFloat64.EndRatio = 0.500000
	__Link__000001_CellFloat64.CornerOffsetRatio = 1.194173

	// Link values setup
	__Link__000002_CellIcon.Name = `CellIcon`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Cell.CellIcon]
	__Link__000002_CellIcon.Identifier = `ref_models.Cell.CellIcon`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CellIcon]
	__Link__000002_CellIcon.Fieldtypename = `ref_models.CellIcon`
	__Link__000002_CellIcon.FieldOffsetX = -70.000000
	__Link__000002_CellIcon.FieldOffsetY = -17.000000
	__Link__000002_CellIcon.TargetMultiplicity = models.ZERO_ONE
	__Link__000002_CellIcon.TargetMultiplicityOffsetX = -50.000000
	__Link__000002_CellIcon.TargetMultiplicityOffsetY = 16.000000
	__Link__000002_CellIcon.SourceMultiplicity = models.MANY
	__Link__000002_CellIcon.SourceMultiplicityOffsetX = 10.000000
	__Link__000002_CellIcon.SourceMultiplicityOffsetY = -50.000000
	__Link__000002_CellIcon.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_CellIcon.StartRatio = 0.500000
	__Link__000002_CellIcon.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_CellIcon.EndRatio = 0.500000
	__Link__000002_CellIcon.CornerOffsetRatio = 1.177507

	// Link values setup
	__Link__000003_CellInt.Name = `CellInt`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Cell.CellInt]
	__Link__000003_CellInt.Identifier = `ref_models.Cell.CellInt`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CellInt]
	__Link__000003_CellInt.Fieldtypename = `ref_models.CellInt`
	__Link__000003_CellInt.FieldOffsetX = -58.000000
	__Link__000003_CellInt.FieldOffsetY = -17.000000
	__Link__000003_CellInt.TargetMultiplicity = models.ZERO_ONE
	__Link__000003_CellInt.TargetMultiplicityOffsetX = -50.000000
	__Link__000003_CellInt.TargetMultiplicityOffsetY = 16.000000
	__Link__000003_CellInt.SourceMultiplicity = models.MANY
	__Link__000003_CellInt.SourceMultiplicityOffsetX = 10.000000
	__Link__000003_CellInt.SourceMultiplicityOffsetY = -50.000000
	__Link__000003_CellInt.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_CellInt.StartRatio = 0.500000
	__Link__000003_CellInt.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_CellInt.EndRatio = 0.500000
	__Link__000003_CellInt.CornerOffsetRatio = 1.194173

	// Link values setup
	__Link__000004_CellString.Name = `CellString`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Cell.CellString]
	__Link__000004_CellString.Identifier = `ref_models.Cell.CellString`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CellString]
	__Link__000004_CellString.Fieldtypename = `ref_models.CellString`
	__Link__000004_CellString.FieldOffsetX = -89.000000
	__Link__000004_CellString.FieldOffsetY = -17.000000
	__Link__000004_CellString.TargetMultiplicity = models.ZERO_ONE
	__Link__000004_CellString.TargetMultiplicityOffsetX = -50.000000
	__Link__000004_CellString.TargetMultiplicityOffsetY = 16.000000
	__Link__000004_CellString.SourceMultiplicity = models.MANY
	__Link__000004_CellString.SourceMultiplicityOffsetX = 10.000000
	__Link__000004_CellString.SourceMultiplicityOffsetY = -50.000000
	__Link__000004_CellString.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_CellString.StartRatio = 0.500000
	__Link__000004_CellString.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_CellString.EndRatio = 0.500000
	__Link__000004_CellString.CornerOffsetRatio = 1.190007

	// Link values setup
	__Link__000005_Cells.Name = `Cells`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Row.Cells]
	__Link__000005_Cells.Identifier = `ref_models.Row.Cells`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Cell]
	__Link__000005_Cells.Fieldtypename = `ref_models.Cell`
	__Link__000005_Cells.FieldOffsetX = -50.000000
	__Link__000005_Cells.FieldOffsetY = -16.000000
	__Link__000005_Cells.TargetMultiplicity = models.MANY
	__Link__000005_Cells.TargetMultiplicityOffsetX = -33.000000
	__Link__000005_Cells.TargetMultiplicityOffsetY = 23.000000
	__Link__000005_Cells.SourceMultiplicity = models.ZERO_ONE
	__Link__000005_Cells.SourceMultiplicityOffsetX = 15.000000
	__Link__000005_Cells.SourceMultiplicityOffsetY = -11.000000
	__Link__000005_Cells.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_Cells.StartRatio = 0.500000
	__Link__000005_Cells.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_Cells.EndRatio = 0.500000
	__Link__000005_Cells.CornerOffsetRatio = 1.121944

	// Link values setup
	__Link__000006_DisplayedColumns.Name = `DisplayedColumns`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Table.DisplayedColumns]
	__Link__000006_DisplayedColumns.Identifier = `ref_models.Table.DisplayedColumns`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DisplayedColumn]
	__Link__000006_DisplayedColumns.Fieldtypename = `ref_models.DisplayedColumn`
	__Link__000006_DisplayedColumns.FieldOffsetX = -153.000000
	__Link__000006_DisplayedColumns.FieldOffsetY = -15.000000
	__Link__000006_DisplayedColumns.TargetMultiplicity = models.MANY
	__Link__000006_DisplayedColumns.TargetMultiplicityOffsetX = -29.000000
	__Link__000006_DisplayedColumns.TargetMultiplicityOffsetY = 21.000000
	__Link__000006_DisplayedColumns.SourceMultiplicity = models.ZERO_ONE
	__Link__000006_DisplayedColumns.SourceMultiplicityOffsetX = 8.000000
	__Link__000006_DisplayedColumns.SourceMultiplicityOffsetY = 24.000000
	__Link__000006_DisplayedColumns.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_DisplayedColumns.StartRatio = 0.412698
	__Link__000006_DisplayedColumns.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_DisplayedColumns.EndRatio = 0.492063
	__Link__000006_DisplayedColumns.CornerOffsetRatio = 1.195833

	// Link values setup
	__Link__000007_Rows.Name = `Rows`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Table.Rows]
	__Link__000007_Rows.Identifier = `ref_models.Table.Rows`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Row]
	__Link__000007_Rows.Fieldtypename = `ref_models.Row`
	__Link__000007_Rows.FieldOffsetX = 19.000000
	__Link__000007_Rows.FieldOffsetY = -17.000000
	__Link__000007_Rows.TargetMultiplicity = models.MANY
	__Link__000007_Rows.TargetMultiplicityOffsetX = -24.000000
	__Link__000007_Rows.TargetMultiplicityOffsetY = -9.000000
	__Link__000007_Rows.SourceMultiplicity = models.ZERO_ONE
	__Link__000007_Rows.SourceMultiplicityOffsetX = -43.000000
	__Link__000007_Rows.SourceMultiplicityOffsetY = 16.000000
	__Link__000007_Rows.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000007_Rows.StartRatio = 0.376111
	__Link__000007_Rows.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000007_Rows.EndRatio = 0.338611
	__Link__000007_Rows.CornerOffsetRatio = 1.315476

	// Position values setup
	__Position__000000_Pos_Default_Cell.X = 535.000000
	__Position__000000_Pos_Default_Cell.Y = 369.000000
	__Position__000000_Pos_Default_Cell.Name = `Pos-Default-Cell`

	// Position values setup
	__Position__000001_Pos_Default_CellBoolean.X = 953.000000
	__Position__000001_Pos_Default_CellBoolean.Y = 56.000000
	__Position__000001_Pos_Default_CellBoolean.Name = `Pos-Default-CellBoolean`

	// Position values setup
	__Position__000002_Pos_Default_CellFloat64.X = 965.000000
	__Position__000002_Pos_Default_CellFloat64.Y = 583.000000
	__Position__000002_Pos_Default_CellFloat64.Name = `Pos-Default-CellFloat64`

	// Position values setup
	__Position__000003_Pos_Default_CellIcon.X = 957.000000
	__Position__000003_Pos_Default_CellIcon.Y = 344.000000
	__Position__000003_Pos_Default_CellIcon.Name = `Pos-Default-CellIcon`

	// Position values setup
	__Position__000004_Pos_Default_CellInt.X = 963.000000
	__Position__000004_Pos_Default_CellInt.Y = 494.000000
	__Position__000004_Pos_Default_CellInt.Name = `Pos-Default-CellInt`

	// Position values setup
	__Position__000005_Pos_Default_CellString.X = 951.999939
	__Position__000005_Pos_Default_CellString.Y = 162.000000
	__Position__000005_Pos_Default_CellString.Name = `Pos-Default-CellString`

	// Position values setup
	__Position__000006_Pos_Default_DisplayedColumn.X = 535.000000
	__Position__000006_Pos_Default_DisplayedColumn.Y = 130.000000
	__Position__000006_Pos_Default_DisplayedColumn.Name = `Pos-Default-DisplayedColumn`

	// Position values setup
	__Position__000007_Pos_Default_Row.X = 20.000000
	__Position__000007_Pos_Default_Row.Y = 364.000000
	__Position__000007_Pos_Default_Row.Name = `Pos-Default-Row`

	// Position values setup
	__Position__000008_Pos_Default_Table.X = 17.000000
	__Position__000008_Pos_Default_Table.Y = 95.000000
	__Position__000008_Pos_Default_Table.Name = `Pos-Default-Table`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellBoolean.X = 1108.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellBoolean.Y = 406.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellBoolean.Name = `Verticle in class diagram Default in middle between Default-Cell and Default-CellBoolean`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellFloat64.X = 1108.500000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellFloat64.Y = 260.500000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellFloat64.Name = `Verticle in class diagram Default in middle between Default-Cell and Default-CellFloat64`

	// Vertice values setup
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellIcon.X = 1109.000000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellIcon.Y = 306.500000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellIcon.Name = `Verticle in class diagram Default in middle between Default-Cell and Default-CellIcon`

	// Vertice values setup
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellInt.X = 1109.500000
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellInt.Y = 353.000000
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellInt.Name = `Verticle in class diagram Default in middle between Default-Cell and Default-CellInt`

	// Vertice values setup
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellString.X = 1105.999970
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellString.Y = 215.500000
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellString.Name = `Verticle in class diagram Default in middle between Default-Cell and Default-CellString`

	// Vertice values setup
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_Row_and_Default_Cell.X = 831.000000
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_Row_and_Default_Cell.Y = 232.500000
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_Row_and_Default_Cell.Name = `Verticle in class diagram Default in middle between Default-Row and Default-Cell`

	// Vertice values setup
	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_Table_and_Default_DisplayedColumn.X = 373.500000
	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_Table_and_Default_DisplayedColumn.Y = 237.000000
	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_Table_and_Default_DisplayedColumn.Name = `Verticle in class diagram Default in middle between Default-Table and Default-DisplayedColumn`

	// Vertice values setup
	__Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_Table_and_Default_Row.X = 568.000000
	__Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_Table_and_Default_Row.Y = 182.000000
	__Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_Table_and_Default_Row.Name = `Verticle in class diagram Default in middle between Default-Table and Default-Row`

	// Setup of pointers
	__Classdiagram__000000_Table.GongStructShapes = append(__Classdiagram__000000_Table.GongStructShapes, __GongStructShape__000008_Default_Table)
	__Classdiagram__000000_Table.GongStructShapes = append(__Classdiagram__000000_Table.GongStructShapes, __GongStructShape__000007_Default_Row)
	__Classdiagram__000000_Table.GongStructShapes = append(__Classdiagram__000000_Table.GongStructShapes, __GongStructShape__000005_Default_CellString)
	__Classdiagram__000000_Table.GongStructShapes = append(__Classdiagram__000000_Table.GongStructShapes, __GongStructShape__000000_Default_Cell)
	__Classdiagram__000000_Table.GongStructShapes = append(__Classdiagram__000000_Table.GongStructShapes, __GongStructShape__000006_Default_DisplayedColumn)
	__Classdiagram__000000_Table.GongStructShapes = append(__Classdiagram__000000_Table.GongStructShapes, __GongStructShape__000001_Default_CellBoolean)
	__Classdiagram__000000_Table.GongStructShapes = append(__Classdiagram__000000_Table.GongStructShapes, __GongStructShape__000002_Default_CellFloat64)
	__Classdiagram__000000_Table.GongStructShapes = append(__Classdiagram__000000_Table.GongStructShapes, __GongStructShape__000003_Default_CellIcon)
	__Classdiagram__000000_Table.GongStructShapes = append(__Classdiagram__000000_Table.GongStructShapes, __GongStructShape__000004_Default_CellInt)
	__GongStructShape__000000_Default_Cell.Position = __Position__000000_Pos_Default_Cell
	__GongStructShape__000000_Default_Cell.Links = append(__GongStructShape__000000_Default_Cell.Links, __Link__000000_CellBool)
	__GongStructShape__000000_Default_Cell.Links = append(__GongStructShape__000000_Default_Cell.Links, __Link__000001_CellFloat64)
	__GongStructShape__000000_Default_Cell.Links = append(__GongStructShape__000000_Default_Cell.Links, __Link__000002_CellIcon)
	__GongStructShape__000000_Default_Cell.Links = append(__GongStructShape__000000_Default_Cell.Links, __Link__000003_CellInt)
	__GongStructShape__000000_Default_Cell.Links = append(__GongStructShape__000000_Default_Cell.Links, __Link__000004_CellString)
	__GongStructShape__000001_Default_CellBoolean.Position = __Position__000001_Pos_Default_CellBoolean
	__GongStructShape__000001_Default_CellBoolean.Fields = append(__GongStructShape__000001_Default_CellBoolean.Fields, __Field__000010_Value)
	__GongStructShape__000002_Default_CellFloat64.Position = __Position__000002_Pos_Default_CellFloat64
	__GongStructShape__000002_Default_CellFloat64.Fields = append(__GongStructShape__000002_Default_CellFloat64.Fields, __Field__000011_Value)
	__GongStructShape__000003_Default_CellIcon.Position = __Position__000003_Pos_Default_CellIcon
	__GongStructShape__000003_Default_CellIcon.Fields = append(__GongStructShape__000003_Default_CellIcon.Fields, __Field__000006_Icon)
	__GongStructShape__000004_Default_CellInt.Position = __Position__000004_Pos_Default_CellInt
	__GongStructShape__000004_Default_CellInt.Fields = append(__GongStructShape__000004_Default_CellInt.Fields, __Field__000008_Value)
	__GongStructShape__000005_Default_CellString.Position = __Position__000005_Pos_Default_CellString
	__GongStructShape__000005_Default_CellString.Fields = append(__GongStructShape__000005_Default_CellString.Fields, __Field__000009_Value)
	__GongStructShape__000006_Default_DisplayedColumn.Position = __Position__000006_Pos_Default_DisplayedColumn
	__GongStructShape__000007_Default_Row.Position = __Position__000007_Pos_Default_Row
	__GongStructShape__000007_Default_Row.Links = append(__GongStructShape__000007_Default_Row.Links, __Link__000005_Cells)
	__GongStructShape__000008_Default_Table.Position = __Position__000008_Pos_Default_Table
	__GongStructShape__000008_Default_Table.Fields = append(__GongStructShape__000008_Default_Table.Fields, __Field__000007_Name)
	__GongStructShape__000008_Default_Table.Fields = append(__GongStructShape__000008_Default_Table.Fields, __Field__000003_HasFiltering)
	__GongStructShape__000008_Default_Table.Fields = append(__GongStructShape__000008_Default_Table.Fields, __Field__000002_HasColumnSorting)
	__GongStructShape__000008_Default_Table.Fields = append(__GongStructShape__000008_Default_Table.Fields, __Field__000004_HasPaginator)
	__GongStructShape__000008_Default_Table.Fields = append(__GongStructShape__000008_Default_Table.Fields, __Field__000001_HasCheckableRows)
	__GongStructShape__000008_Default_Table.Fields = append(__GongStructShape__000008_Default_Table.Fields, __Field__000005_HasSaveButton)
	__GongStructShape__000008_Default_Table.Fields = append(__GongStructShape__000008_Default_Table.Fields, __Field__000000_CanDragDropRows)
	__GongStructShape__000008_Default_Table.Links = append(__GongStructShape__000008_Default_Table.Links, __Link__000007_Rows)
	__GongStructShape__000008_Default_Table.Links = append(__GongStructShape__000008_Default_Table.Links, __Link__000006_DisplayedColumns)
	__Link__000000_CellBool.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellBoolean
	__Link__000001_CellFloat64.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellFloat64
	__Link__000002_CellIcon.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellIcon
	__Link__000003_CellInt.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellInt
	__Link__000004_CellString.Middlevertice = __Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Cell_and_Default_CellString
	__Link__000005_Cells.Middlevertice = __Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_Row_and_Default_Cell
	__Link__000006_DisplayedColumns.Middlevertice = __Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_Table_and_Default_DisplayedColumn
	__Link__000007_Rows.Middlevertice = __Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_Table_and_Default_Row
}
