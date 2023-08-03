package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongtable/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_Form models.StageStruct
var ___dummy__Time_Form time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_Form ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_Form map[string]any = map[string]any{
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

	"ref_models.Form": &(ref_models.Form{}),

	"ref_models.Form.FormCells": (ref_models.Form{}).FormCells,

	"ref_models.Form.Name": (ref_models.Form{}).Name,

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
// 	InjectionGateway["Form"] = FormInjection
// }

// FormInjection will stage objects of database "Form"
func FormInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Form := (&models.Classdiagram{Name: `Form`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000001_Value := (&models.Field{Name: `Value`}).Stage(stage)
	__Field__000002_Value := (&models.Field{Name: `Value`}).Stage(stage)
	__Field__000003_Value := (&models.Field{Name: `Value`}).Stage(stage)
	__Field__000004_Value := (&models.Field{Name: `Value`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_Form_Form := (&models.GongStructShape{Name: `Form-Form`}).Stage(stage)
	__GongStructShape__000001_Form_FormCell := (&models.GongStructShape{Name: `Form-FormCell`}).Stage(stage)
	__GongStructShape__000002_Form_FormCellBoolean := (&models.GongStructShape{Name: `Form-FormCellBoolean`}).Stage(stage)
	__GongStructShape__000003_Form_FormCellFloat64 := (&models.GongStructShape{Name: `Form-FormCellFloat64`}).Stage(stage)
	__GongStructShape__000004_Form_FormCellInt := (&models.GongStructShape{Name: `Form-FormCellInt`}).Stage(stage)
	__GongStructShape__000005_Form_FormCellString := (&models.GongStructShape{Name: `Form-FormCellString`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_FormCellBool := (&models.Link{Name: `FormCellBool`}).Stage(stage)
	__Link__000001_FormCellFloat64 := (&models.Link{Name: `FormCellFloat64`}).Stage(stage)
	__Link__000002_FormCellInt := (&models.Link{Name: `FormCellInt`}).Stage(stage)
	__Link__000003_FormCellString := (&models.Link{Name: `FormCellString`}).Stage(stage)
	__Link__000004_FormCells := (&models.Link{Name: `FormCells`}).Stage(stage)

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_Form_Form := (&models.Position{Name: `Pos-Form-Form`}).Stage(stage)
	__Position__000001_Pos_Form_FormCell := (&models.Position{Name: `Pos-Form-FormCell`}).Stage(stage)
	__Position__000002_Pos_Form_FormCellBoolean := (&models.Position{Name: `Pos-Form-FormCellBoolean`}).Stage(stage)
	__Position__000003_Pos_Form_FormCellFloat64 := (&models.Position{Name: `Pos-Form-FormCellFloat64`}).Stage(stage)
	__Position__000004_Pos_Form_FormCellInt := (&models.Position{Name: `Pos-Form-FormCellInt`}).Stage(stage)
	__Position__000005_Pos_Form_FormCellString := (&models.Position{Name: `Pos-Form-FormCellString`}).Stage(stage)

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_Form_in_middle_between_Form_Form_and_Form_FormCell := (&models.Vertice{Name: `Verticle in class diagram Form in middle between Form-Form and Form-FormCell`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellBoolean := (&models.Vertice{Name: `Verticle in class diagram Form in middle between Form-FormCell and Form-FormCellBoolean`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellFloat64 := (&models.Vertice{Name: `Verticle in class diagram Form in middle between Form-FormCell and Form-FormCellFloat64`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellInt := (&models.Vertice{Name: `Verticle in class diagram Form in middle between Form-FormCell and Form-FormCellInt`}).Stage(stage)
	__Vertice__000004_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellString := (&models.Vertice{Name: `Verticle in class diagram Form in middle between Form-FormCell and Form-FormCellString`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Form.Name = `Form`
	__Classdiagram__000000_Form.IsInDrawMode = true

	// Field values setup
	__Field__000000_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Form.Name]
	__Field__000000_Name.Identifier = `ref_models.Form.Name`
	__Field__000000_Name.FieldTypeAsString = ``
	__Field__000000_Name.Structname = `Form`
	__Field__000000_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000001_Value.Name = `Value`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormCellFloat64.Value]
	__Field__000001_Value.Identifier = `ref_models.FormCellFloat64.Value`
	__Field__000001_Value.FieldTypeAsString = ``
	__Field__000001_Value.Structname = `FormCellFloat64`
	__Field__000001_Value.Fieldtypename = `float64`

	// Field values setup
	__Field__000002_Value.Name = `Value`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormCellBoolean.Value]
	__Field__000002_Value.Identifier = `ref_models.FormCellBoolean.Value`
	__Field__000002_Value.FieldTypeAsString = ``
	__Field__000002_Value.Structname = `FormCellBoolean`
	__Field__000002_Value.Fieldtypename = `bool`

	// Field values setup
	__Field__000003_Value.Name = `Value`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormCellString.Value]
	__Field__000003_Value.Identifier = `ref_models.FormCellString.Value`
	__Field__000003_Value.FieldTypeAsString = ``
	__Field__000003_Value.Structname = `FormCellString`
	__Field__000003_Value.Fieldtypename = `string`

	// Field values setup
	__Field__000004_Value.Name = `Value`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormCellInt.Value]
	__Field__000004_Value.Identifier = `ref_models.FormCellInt.Value`
	__Field__000004_Value.FieldTypeAsString = ``
	__Field__000004_Value.Structname = `FormCellInt`
	__Field__000004_Value.Fieldtypename = `int`

	// GongStructShape values setup
	__GongStructShape__000000_Form_Form.Name = `Form-Form`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Form]
	__GongStructShape__000000_Form_Form.Identifier = `ref_models.Form`
	__GongStructShape__000000_Form_Form.ShowNbInstances = false
	__GongStructShape__000000_Form_Form.NbInstances = 0
	__GongStructShape__000000_Form_Form.Width = 240.000000
	__GongStructShape__000000_Form_Form.Heigth = 78.000000
	__GongStructShape__000000_Form_Form.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_Form_FormCell.Name = `Form-FormCell`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormCell]
	__GongStructShape__000001_Form_FormCell.Identifier = `ref_models.FormCell`
	__GongStructShape__000001_Form_FormCell.ShowNbInstances = false
	__GongStructShape__000001_Form_FormCell.NbInstances = 0
	__GongStructShape__000001_Form_FormCell.Width = 240.000000
	__GongStructShape__000001_Form_FormCell.Heigth = 63.000000
	__GongStructShape__000001_Form_FormCell.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_Form_FormCellBoolean.Name = `Form-FormCellBoolean`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormCellBoolean]
	__GongStructShape__000002_Form_FormCellBoolean.Identifier = `ref_models.FormCellBoolean`
	__GongStructShape__000002_Form_FormCellBoolean.ShowNbInstances = false
	__GongStructShape__000002_Form_FormCellBoolean.NbInstances = 0
	__GongStructShape__000002_Form_FormCellBoolean.Width = 240.000000
	__GongStructShape__000002_Form_FormCellBoolean.Heigth = 78.000000
	__GongStructShape__000002_Form_FormCellBoolean.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000003_Form_FormCellFloat64.Name = `Form-FormCellFloat64`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormCellFloat64]
	__GongStructShape__000003_Form_FormCellFloat64.Identifier = `ref_models.FormCellFloat64`
	__GongStructShape__000003_Form_FormCellFloat64.ShowNbInstances = false
	__GongStructShape__000003_Form_FormCellFloat64.NbInstances = 0
	__GongStructShape__000003_Form_FormCellFloat64.Width = 240.000000
	__GongStructShape__000003_Form_FormCellFloat64.Heigth = 78.000000
	__GongStructShape__000003_Form_FormCellFloat64.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000004_Form_FormCellInt.Name = `Form-FormCellInt`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormCellInt]
	__GongStructShape__000004_Form_FormCellInt.Identifier = `ref_models.FormCellInt`
	__GongStructShape__000004_Form_FormCellInt.ShowNbInstances = false
	__GongStructShape__000004_Form_FormCellInt.NbInstances = 0
	__GongStructShape__000004_Form_FormCellInt.Width = 240.000000
	__GongStructShape__000004_Form_FormCellInt.Heigth = 78.000000
	__GongStructShape__000004_Form_FormCellInt.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000005_Form_FormCellString.Name = `Form-FormCellString`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormCellString]
	__GongStructShape__000005_Form_FormCellString.Identifier = `ref_models.FormCellString`
	__GongStructShape__000005_Form_FormCellString.ShowNbInstances = false
	__GongStructShape__000005_Form_FormCellString.NbInstances = 0
	__GongStructShape__000005_Form_FormCellString.Width = 240.000000
	__GongStructShape__000005_Form_FormCellString.Heigth = 78.000000
	__GongStructShape__000005_Form_FormCellString.IsSelected = false

	// Link values setup
	__Link__000000_FormCellBool.Name = `FormCellBool`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormCell.FormCellBool]
	__Link__000000_FormCellBool.Identifier = `ref_models.FormCell.FormCellBool`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormCellBoolean]
	__Link__000000_FormCellBool.Fieldtypename = `ref_models.FormCellBoolean`
	__Link__000000_FormCellBool.FieldOffsetX = -116.000000
	__Link__000000_FormCellBool.FieldOffsetY = -18.000000
	__Link__000000_FormCellBool.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_FormCellBool.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_FormCellBool.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_FormCellBool.SourceMultiplicity = models.MANY
	__Link__000000_FormCellBool.SourceMultiplicityOffsetX = 27.000000
	__Link__000000_FormCellBool.SourceMultiplicityOffsetY = 41.000000
	__Link__000000_FormCellBool.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_FormCellBool.StartRatio = 0.500000
	__Link__000000_FormCellBool.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_FormCellBool.EndRatio = 0.500000
	__Link__000000_FormCellBool.CornerOffsetRatio = 1.731641

	// Link values setup
	__Link__000001_FormCellFloat64.Name = `FormCellFloat64`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormCell.FormCellFloat64]
	__Link__000001_FormCellFloat64.Identifier = `ref_models.FormCell.FormCellFloat64`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormCellFloat64]
	__Link__000001_FormCellFloat64.Fieldtypename = `ref_models.FormCellFloat64`
	__Link__000001_FormCellFloat64.FieldOffsetX = -131.000000
	__Link__000001_FormCellFloat64.FieldOffsetY = -21.000000
	__Link__000001_FormCellFloat64.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_FormCellFloat64.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_FormCellFloat64.TargetMultiplicityOffsetY = 16.000000
	__Link__000001_FormCellFloat64.SourceMultiplicity = models.MANY
	__Link__000001_FormCellFloat64.SourceMultiplicityOffsetX = 27.000000
	__Link__000001_FormCellFloat64.SourceMultiplicityOffsetY = 28.000000
	__Link__000001_FormCellFloat64.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_FormCellFloat64.StartRatio = 0.500000
	__Link__000001_FormCellFloat64.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_FormCellFloat64.EndRatio = 0.500000
	__Link__000001_FormCellFloat64.CornerOffsetRatio = 1.731641

	// Link values setup
	__Link__000002_FormCellInt.Name = `FormCellInt`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormCell.FormCellInt]
	__Link__000002_FormCellInt.Identifier = `ref_models.FormCell.FormCellInt`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormCellInt]
	__Link__000002_FormCellInt.Fieldtypename = `ref_models.FormCellInt`
	__Link__000002_FormCellInt.FieldOffsetX = -93.000000
	__Link__000002_FormCellInt.FieldOffsetY = -16.000000
	__Link__000002_FormCellInt.TargetMultiplicity = models.ZERO_ONE
	__Link__000002_FormCellInt.TargetMultiplicityOffsetX = -50.000000
	__Link__000002_FormCellInt.TargetMultiplicityOffsetY = 16.000000
	__Link__000002_FormCellInt.SourceMultiplicity = models.MANY
	__Link__000002_FormCellInt.SourceMultiplicityOffsetX = 14.000000
	__Link__000002_FormCellInt.SourceMultiplicityOffsetY = 41.000000
	__Link__000002_FormCellInt.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_FormCellInt.StartRatio = 0.500000
	__Link__000002_FormCellInt.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_FormCellInt.EndRatio = 0.500000
	__Link__000002_FormCellInt.CornerOffsetRatio = 1.731641

	// Link values setup
	__Link__000003_FormCellString.Name = `FormCellString`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormCell.FormCellString]
	__Link__000003_FormCellString.Identifier = `ref_models.FormCell.FormCellString`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormCellString]
	__Link__000003_FormCellString.Fieldtypename = `ref_models.FormCellString`
	__Link__000003_FormCellString.FieldOffsetX = -119.000000
	__Link__000003_FormCellString.FieldOffsetY = -22.000000
	__Link__000003_FormCellString.TargetMultiplicity = models.ZERO_ONE
	__Link__000003_FormCellString.TargetMultiplicityOffsetX = -50.000000
	__Link__000003_FormCellString.TargetMultiplicityOffsetY = 16.000000
	__Link__000003_FormCellString.SourceMultiplicity = models.MANY
	__Link__000003_FormCellString.SourceMultiplicityOffsetX = 16.000000
	__Link__000003_FormCellString.SourceMultiplicityOffsetY = 28.000000
	__Link__000003_FormCellString.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_FormCellString.StartRatio = 0.500000
	__Link__000003_FormCellString.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_FormCellString.EndRatio = 0.500000
	__Link__000003_FormCellString.CornerOffsetRatio = 1.731641

	// Link values setup
	__Link__000004_FormCells.Name = `FormCells`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Form.FormCells]
	__Link__000004_FormCells.Identifier = `ref_models.Form.FormCells`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormCell]
	__Link__000004_FormCells.Fieldtypename = `ref_models.FormCell`
	__Link__000004_FormCells.FieldOffsetX = -83.000000
	__Link__000004_FormCells.FieldOffsetY = -12.000000
	__Link__000004_FormCells.TargetMultiplicity = models.MANY
	__Link__000004_FormCells.TargetMultiplicityOffsetX = 19.000000
	__Link__000004_FormCells.TargetMultiplicityOffsetY = -10.000000
	__Link__000004_FormCells.SourceMultiplicity = models.ZERO_ONE
	__Link__000004_FormCells.SourceMultiplicityOffsetX = 14.000000
	__Link__000004_FormCells.SourceMultiplicityOffsetY = 21.000000
	__Link__000004_FormCells.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000004_FormCells.StartRatio = 0.310807
	__Link__000004_FormCells.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000004_FormCells.EndRatio = 0.298307
	__Link__000004_FormCells.CornerOffsetRatio = 1.307692

	// Position values setup
	__Position__000000_Pos_Form_Form.X = 54.000000
	__Position__000000_Pos_Form_Form.Y = 50.000000
	__Position__000000_Pos_Form_Form.Name = `Pos-Form-Form`

	// Position values setup
	__Position__000001_Pos_Form_FormCell.X = 55.000000
	__Position__000001_Pos_Form_FormCell.Y = 249.000000
	__Position__000001_Pos_Form_FormCell.Name = `Pos-Form-FormCell`

	// Position values setup
	__Position__000002_Pos_Form_FormCellBoolean.X = 645.000000
	__Position__000002_Pos_Form_FormCellBoolean.Y = 39.000000
	__Position__000002_Pos_Form_FormCellBoolean.Name = `Pos-Form-FormCellBoolean`

	// Position values setup
	__Position__000003_Pos_Form_FormCellFloat64.X = 641.000000
	__Position__000003_Pos_Form_FormCellFloat64.Y = 389.000000
	__Position__000003_Pos_Form_FormCellFloat64.Name = `Pos-Form-FormCellFloat64`

	// Position values setup
	__Position__000004_Pos_Form_FormCellInt.X = 641.000000
	__Position__000004_Pos_Form_FormCellInt.Y = 284.000000
	__Position__000004_Pos_Form_FormCellInt.Name = `Pos-Form-FormCellInt`

	// Position values setup
	__Position__000005_Pos_Form_FormCellString.X = 643.000000
	__Position__000005_Pos_Form_FormCellString.Y = 156.000000
	__Position__000005_Pos_Form_FormCellString.Name = `Pos-Form-FormCellString`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_Form_in_middle_between_Form_Form_and_Form_FormCell.X = 443.500000
	__Vertice__000000_Verticle_in_class_diagram_Form_in_middle_between_Form_Form_and_Form_FormCell.Y = 132.500000
	__Vertice__000000_Verticle_in_class_diagram_Form_in_middle_between_Form_Form_and_Form_FormCell.Name = `Verticle in class diagram Form in middle between Form-Form and Form-FormCell`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellBoolean.X = 404.500000
	__Vertice__000001_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellBoolean.Y = 113.500000
	__Vertice__000001_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellBoolean.Name = `Verticle in class diagram Form in middle between Form-FormCell and Form-FormCellBoolean`

	// Vertice values setup
	__Vertice__000002_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellFloat64.X = 409.500000
	__Vertice__000002_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellFloat64.Y = 128.500000
	__Vertice__000002_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellFloat64.Name = `Verticle in class diagram Form in middle between Form-FormCell and Form-FormCellFloat64`

	// Vertice values setup
	__Vertice__000003_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellInt.X = 413.000000
	__Vertice__000003_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellInt.Y = 108.500000
	__Vertice__000003_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellInt.Name = `Verticle in class diagram Form in middle between Form-FormCell and Form-FormCellInt`

	// Vertice values setup
	__Vertice__000004_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellString.X = 435.500000
	__Vertice__000004_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellString.Y = 126.500000
	__Vertice__000004_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellString.Name = `Verticle in class diagram Form in middle between Form-FormCell and Form-FormCellString`

	// Setup of pointers
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000000_Form_Form)
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000001_Form_FormCell)
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000002_Form_FormCellBoolean)
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000003_Form_FormCellFloat64)
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000004_Form_FormCellInt)
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000005_Form_FormCellString)
	__GongStructShape__000000_Form_Form.Position = __Position__000000_Pos_Form_Form
	__GongStructShape__000000_Form_Form.Fields = append(__GongStructShape__000000_Form_Form.Fields, __Field__000000_Name)
	__GongStructShape__000000_Form_Form.Links = append(__GongStructShape__000000_Form_Form.Links, __Link__000004_FormCells)
	__GongStructShape__000001_Form_FormCell.Position = __Position__000001_Pos_Form_FormCell
	__GongStructShape__000001_Form_FormCell.Links = append(__GongStructShape__000001_Form_FormCell.Links, __Link__000000_FormCellBool)
	__GongStructShape__000001_Form_FormCell.Links = append(__GongStructShape__000001_Form_FormCell.Links, __Link__000001_FormCellFloat64)
	__GongStructShape__000001_Form_FormCell.Links = append(__GongStructShape__000001_Form_FormCell.Links, __Link__000002_FormCellInt)
	__GongStructShape__000001_Form_FormCell.Links = append(__GongStructShape__000001_Form_FormCell.Links, __Link__000003_FormCellString)
	__GongStructShape__000002_Form_FormCellBoolean.Position = __Position__000002_Pos_Form_FormCellBoolean
	__GongStructShape__000002_Form_FormCellBoolean.Fields = append(__GongStructShape__000002_Form_FormCellBoolean.Fields, __Field__000002_Value)
	__GongStructShape__000003_Form_FormCellFloat64.Position = __Position__000003_Pos_Form_FormCellFloat64
	__GongStructShape__000003_Form_FormCellFloat64.Fields = append(__GongStructShape__000003_Form_FormCellFloat64.Fields, __Field__000001_Value)
	__GongStructShape__000004_Form_FormCellInt.Position = __Position__000004_Pos_Form_FormCellInt
	__GongStructShape__000004_Form_FormCellInt.Fields = append(__GongStructShape__000004_Form_FormCellInt.Fields, __Field__000004_Value)
	__GongStructShape__000005_Form_FormCellString.Position = __Position__000005_Pos_Form_FormCellString
	__GongStructShape__000005_Form_FormCellString.Fields = append(__GongStructShape__000005_Form_FormCellString.Fields, __Field__000003_Value)
	__Link__000000_FormCellBool.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellBoolean
	__Link__000001_FormCellFloat64.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellFloat64
	__Link__000002_FormCellInt.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellInt
	__Link__000003_FormCellString.Middlevertice = __Vertice__000004_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellString
	__Link__000004_FormCells.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Form_in_middle_between_Form_Form_and_Form_FormCell
}
