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
var map_DocLink_Identifier_Form map[string]any = map[string]any{}

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
	__Field__000000_InputTypeEnum := (&models.Field{Name: `InputTypeEnum`}).Stage(stage)
	__Field__000001_Label := (&models.Field{Name: `Label`}).Stage(stage)
	__Field__000002_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000003_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000004_Placeholder := (&models.Field{Name: `Placeholder`}).Stage(stage)
	__Field__000005_Step := (&models.Field{Name: `Step`}).Stage(stage)
	__Field__000006_Value := (&models.Field{Name: `Value`}).Stage(stage)
	__Field__000007_Value := (&models.Field{Name: `Value`}).Stage(stage)
	__Field__000008_Value := (&models.Field{Name: `Value`}).Stage(stage)
	__Field__000009_Value := (&models.Field{Name: `Value`}).Stage(stage)
	__Field__000010_Value := (&models.Field{Name: `Value`}).Stage(stage)
	__Field__000011_Value := (&models.Field{Name: `Value`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape
	__GongEnumShape__000000_Form_InputTypeEnum := (&models.GongEnumShape{Name: `Form-InputTypeEnum`}).Stage(stage)

	// Declarations of staged instances of GongEnumValueEntry
	__GongEnumValueEntry__000000_Color := (&models.GongEnumValueEntry{Name: `Color`}).Stage(stage)
	__GongEnumValueEntry__000001_Date := (&models.GongEnumValueEntry{Name: `Date`}).Stage(stage)
	__GongEnumValueEntry__000002_Datetime := (&models.GongEnumValueEntry{Name: `Datetime`}).Stage(stage)
	__GongEnumValueEntry__000003_Email := (&models.GongEnumValueEntry{Name: `Email`}).Stage(stage)
	__GongEnumValueEntry__000004_File := (&models.GongEnumValueEntry{Name: `File`}).Stage(stage)
	__GongEnumValueEntry__000005_Hidden := (&models.GongEnumValueEntry{Name: `Hidden`}).Stage(stage)
	__GongEnumValueEntry__000006_Month := (&models.GongEnumValueEntry{Name: `Month`}).Stage(stage)
	__GongEnumValueEntry__000007_Number := (&models.GongEnumValueEntry{Name: `Number`}).Stage(stage)
	__GongEnumValueEntry__000008_Password := (&models.GongEnumValueEntry{Name: `Password`}).Stage(stage)
	__GongEnumValueEntry__000009_Range := (&models.GongEnumValueEntry{Name: `Range`}).Stage(stage)
	__GongEnumValueEntry__000010_Search := (&models.GongEnumValueEntry{Name: `Search`}).Stage(stage)
	__GongEnumValueEntry__000011_Tel := (&models.GongEnumValueEntry{Name: `Tel`}).Stage(stage)
	__GongEnumValueEntry__000012_Text := (&models.GongEnumValueEntry{Name: `Text`}).Stage(stage)
	__GongEnumValueEntry__000013_Time := (&models.GongEnumValueEntry{Name: `Time`}).Stage(stage)
	__GongEnumValueEntry__000014_URL := (&models.GongEnumValueEntry{Name: `URL`}).Stage(stage)
	__GongEnumValueEntry__000015_Week := (&models.GongEnumValueEntry{Name: `Week`}).Stage(stage)

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_Form_Form := (&models.GongStructShape{Name: `Form-Form`}).Stage(stage)
	__GongStructShape__000001_Form_FormCell := (&models.GongStructShape{Name: `Form-FormCell`}).Stage(stage)
	__GongStructShape__000002_Form_FormCellBoolean := (&models.GongStructShape{Name: `Form-FormCellBoolean`}).Stage(stage)
	__GongStructShape__000003_Form_FormCellFloat64 := (&models.GongStructShape{Name: `Form-FormCellFloat64`}).Stage(stage)
	__GongStructShape__000004_Form_FormCellInt := (&models.GongStructShape{Name: `Form-FormCellInt`}).Stage(stage)
	__GongStructShape__000005_Form_FormCellString := (&models.GongStructShape{Name: `Form-FormCellString`}).Stage(stage)
	__GongStructShape__000006_Form_FormFieldDate := (&models.GongStructShape{Name: `Form-FormFieldDate`}).Stage(stage)
	__GongStructShape__000007_Form_FormFieldTime := (&models.GongStructShape{Name: `Form-FormFieldTime`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_FormCellBool := (&models.Link{Name: `FormCellBool`}).Stage(stage)
	__Link__000001_FormCellFloat64 := (&models.Link{Name: `FormCellFloat64`}).Stage(stage)
	__Link__000002_FormCellInt := (&models.Link{Name: `FormCellInt`}).Stage(stage)
	__Link__000003_FormCellString := (&models.Link{Name: `FormCellString`}).Stage(stage)
	__Link__000004_FormCells := (&models.Link{Name: `FormCells`}).Stage(stage)
	__Link__000005_FormFieldDate := (&models.Link{Name: `FormFieldDate`}).Stage(stage)
	__Link__000006_FormFieldTime := (&models.Link{Name: `FormFieldTime`}).Stage(stage)

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_Form_Form := (&models.Position{Name: `Pos-Form-Form`}).Stage(stage)
	__Position__000001_Pos_Form_FormCell := (&models.Position{Name: `Pos-Form-FormCell`}).Stage(stage)
	__Position__000002_Pos_Form_FormCellBoolean := (&models.Position{Name: `Pos-Form-FormCellBoolean`}).Stage(stage)
	__Position__000003_Pos_Form_FormCellFloat64 := (&models.Position{Name: `Pos-Form-FormCellFloat64`}).Stage(stage)
	__Position__000004_Pos_Form_FormCellInt := (&models.Position{Name: `Pos-Form-FormCellInt`}).Stage(stage)
	__Position__000005_Pos_Form_FormCellString := (&models.Position{Name: `Pos-Form-FormCellString`}).Stage(stage)
	__Position__000006_Pos_Form_FormFieldDate := (&models.Position{Name: `Pos-Form-FormFieldDate`}).Stage(stage)
	__Position__000007_Pos_Form_FormFieldTime := (&models.Position{Name: `Pos-Form-FormFieldTime`}).Stage(stage)
	__Position__000008_Pos_Form_InputTypeEnum := (&models.Position{Name: `Pos-Form-InputTypeEnum`}).Stage(stage)

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_Form_in_middle_between_Form_Form_and_Form_FormCell := (&models.Vertice{Name: `Verticle in class diagram Form in middle between Form-Form and Form-FormCell`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellBoolean := (&models.Vertice{Name: `Verticle in class diagram Form in middle between Form-FormCell and Form-FormCellBoolean`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellFloat64 := (&models.Vertice{Name: `Verticle in class diagram Form in middle between Form-FormCell and Form-FormCellFloat64`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellInt := (&models.Vertice{Name: `Verticle in class diagram Form in middle between Form-FormCell and Form-FormCellInt`}).Stage(stage)
	__Vertice__000004_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellString := (&models.Vertice{Name: `Verticle in class diagram Form in middle between Form-FormCell and Form-FormCellString`}).Stage(stage)
	__Vertice__000005_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldDate := (&models.Vertice{Name: `Verticle in class diagram Form in middle between Form-FormCell and Form-FormFieldDate`}).Stage(stage)
	__Vertice__000006_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldTime := (&models.Vertice{Name: `Verticle in class diagram Form in middle between Form-FormCell and Form-FormFieldTime`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Form.Name = `Form`
	__Classdiagram__000000_Form.IsInDrawMode = true

	// Field values setup
	__Field__000000_InputTypeEnum.Name = `InputTypeEnum`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormField.InputTypeEnum]
	__Field__000000_InputTypeEnum.Identifier = `ref_models.FormField.InputTypeEnum`
	__Field__000000_InputTypeEnum.FieldTypeAsString = ``
	__Field__000000_InputTypeEnum.Structname = `FormField`
	__Field__000000_InputTypeEnum.Fieldtypename = `InputTypeEnum`

	// Field values setup
	__Field__000001_Label.Name = `Label`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormField.Label]
	__Field__000001_Label.Identifier = `ref_models.FormField.Label`
	__Field__000001_Label.FieldTypeAsString = ``
	__Field__000001_Label.Structname = `FormField`
	__Field__000001_Label.Fieldtypename = `string`

	// Field values setup
	__Field__000002_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormGroup.Name]
	__Field__000002_Name.Identifier = `ref_models.FormGroup.Name`
	__Field__000002_Name.FieldTypeAsString = ``
	__Field__000002_Name.Structname = `Form`
	__Field__000002_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000003_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormField.Name]
	__Field__000003_Name.Identifier = `ref_models.FormField.Name`
	__Field__000003_Name.FieldTypeAsString = ``
	__Field__000003_Name.Structname = `FormField`
	__Field__000003_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000004_Placeholder.Name = `Placeholder`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormField.Placeholder]
	__Field__000004_Placeholder.Identifier = `ref_models.FormField.Placeholder`
	__Field__000004_Placeholder.FieldTypeAsString = ``
	__Field__000004_Placeholder.Structname = `FormField`
	__Field__000004_Placeholder.Fieldtypename = `string`

	// Field values setup
	__Field__000005_Step.Name = `Step`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldTime.Step]
	__Field__000005_Step.Identifier = `ref_models.FormFieldTime.Step`
	__Field__000005_Step.FieldTypeAsString = ``
	__Field__000005_Step.Structname = `FormFieldTime`
	__Field__000005_Step.Fieldtypename = `float64`

	// Field values setup
	__Field__000006_Value.Name = `Value`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldInt.Value]
	__Field__000006_Value.Identifier = `ref_models.FormFieldInt.Value`
	__Field__000006_Value.FieldTypeAsString = ``
	__Field__000006_Value.Structname = `FormCellInt`
	__Field__000006_Value.Fieldtypename = `int`

	// Field values setup
	__Field__000007_Value.Name = `Value`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldString.Value]
	__Field__000007_Value.Identifier = `ref_models.FormFieldString.Value`
	__Field__000007_Value.FieldTypeAsString = ``
	__Field__000007_Value.Structname = `FormCellString`
	__Field__000007_Value.Fieldtypename = `string`

	// Field values setup
	__Field__000008_Value.Name = `Value`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldTime.Value]
	__Field__000008_Value.Identifier = `ref_models.FormFieldTime.Value`
	__Field__000008_Value.FieldTypeAsString = ``
	__Field__000008_Value.Structname = `FormFieldTime`
	__Field__000008_Value.Fieldtypename = `Time`

	// Field values setup
	__Field__000009_Value.Name = `Value`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldFloat64.Value]
	__Field__000009_Value.Identifier = `ref_models.FormFieldFloat64.Value`
	__Field__000009_Value.FieldTypeAsString = ``
	__Field__000009_Value.Structname = `FormCellFloat64`
	__Field__000009_Value.Fieldtypename = `float64`

	// Field values setup
	__Field__000010_Value.Name = `Value`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldBoolean.Value]
	__Field__000010_Value.Identifier = `ref_models.FormFieldBoolean.Value`
	__Field__000010_Value.FieldTypeAsString = ``
	__Field__000010_Value.Structname = `FormCellBoolean`
	__Field__000010_Value.Fieldtypename = `bool`

	// Field values setup
	__Field__000011_Value.Name = `Value`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldDate.Value]
	__Field__000011_Value.Identifier = `ref_models.FormFieldDate.Value`
	__Field__000011_Value.FieldTypeAsString = ``
	__Field__000011_Value.Structname = `FormFieldDate`
	__Field__000011_Value.Fieldtypename = `Time`

	// GongEnumShape values setup
	__GongEnumShape__000000_Form_InputTypeEnum.Name = `Form-InputTypeEnum`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum]
	__GongEnumShape__000000_Form_InputTypeEnum.Identifier = `ref_models.InputTypeEnum`
	__GongEnumShape__000000_Form_InputTypeEnum.Width = 240.000000
	__GongEnumShape__000000_Form_InputTypeEnum.Heigth = 303.000000

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000000_Color.Name = `Color`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Color]
	__GongEnumValueEntry__000000_Color.Identifier = `ref_models.InputTypeEnum.Color`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000001_Date.Name = `Date`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Date]
	__GongEnumValueEntry__000001_Date.Identifier = `ref_models.InputTypeEnum.Date`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000002_Datetime.Name = `Datetime`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Datetime]
	__GongEnumValueEntry__000002_Datetime.Identifier = `ref_models.InputTypeEnum.Datetime`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000003_Email.Name = `Email`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Email]
	__GongEnumValueEntry__000003_Email.Identifier = `ref_models.InputTypeEnum.Email`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000004_File.Name = `File`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.File]
	__GongEnumValueEntry__000004_File.Identifier = `ref_models.InputTypeEnum.File`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000005_Hidden.Name = `Hidden`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Hidden]
	__GongEnumValueEntry__000005_Hidden.Identifier = `ref_models.InputTypeEnum.Hidden`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000006_Month.Name = `Month`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Month]
	__GongEnumValueEntry__000006_Month.Identifier = `ref_models.InputTypeEnum.Month`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000007_Number.Name = `Number`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Number]
	__GongEnumValueEntry__000007_Number.Identifier = `ref_models.InputTypeEnum.Number`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000008_Password.Name = `Password`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Password]
	__GongEnumValueEntry__000008_Password.Identifier = `ref_models.InputTypeEnum.Password`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000009_Range.Name = `Range`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Range]
	__GongEnumValueEntry__000009_Range.Identifier = `ref_models.InputTypeEnum.Range`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000010_Search.Name = `Search`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Search]
	__GongEnumValueEntry__000010_Search.Identifier = `ref_models.InputTypeEnum.Search`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000011_Tel.Name = `Tel`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Tel]
	__GongEnumValueEntry__000011_Tel.Identifier = `ref_models.InputTypeEnum.Tel`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000012_Text.Name = `Text`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Text]
	__GongEnumValueEntry__000012_Text.Identifier = `ref_models.InputTypeEnum.Text`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000013_Time.Name = `Time`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Time]
	__GongEnumValueEntry__000013_Time.Identifier = `ref_models.InputTypeEnum.Time`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000014_URL.Name = `URL`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.URL]
	__GongEnumValueEntry__000014_URL.Identifier = `ref_models.InputTypeEnum.URL`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000015_Week.Name = `Week`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Week]
	__GongEnumValueEntry__000015_Week.Identifier = `ref_models.InputTypeEnum.Week`

	// GongStructShape values setup
	__GongStructShape__000000_Form_Form.Name = `Form-Form`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormGroup]
	__GongStructShape__000000_Form_Form.Identifier = `ref_models.FormGroup`
	__GongStructShape__000000_Form_Form.ShowNbInstances = true
	__GongStructShape__000000_Form_Form.NbInstances = 0
	__GongStructShape__000000_Form_Form.Width = 276.000000
	__GongStructShape__000000_Form_Form.Heigth = 78.000000
	__GongStructShape__000000_Form_Form.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_Form_FormCell.Name = `Form-FormCell`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormField]
	__GongStructShape__000001_Form_FormCell.Identifier = `ref_models.FormField`
	__GongStructShape__000001_Form_FormCell.ShowNbInstances = true
	__GongStructShape__000001_Form_FormCell.NbInstances = 0
	__GongStructShape__000001_Form_FormCell.Width = 263.000000
	__GongStructShape__000001_Form_FormCell.Heigth = 123.000000
	__GongStructShape__000001_Form_FormCell.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_Form_FormCellBoolean.Name = `Form-FormCellBoolean`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldBoolean]
	__GongStructShape__000002_Form_FormCellBoolean.Identifier = `ref_models.FormFieldBoolean`
	__GongStructShape__000002_Form_FormCellBoolean.ShowNbInstances = true
	__GongStructShape__000002_Form_FormCellBoolean.NbInstances = 0
	__GongStructShape__000002_Form_FormCellBoolean.Width = 240.000000
	__GongStructShape__000002_Form_FormCellBoolean.Heigth = 78.000000
	__GongStructShape__000002_Form_FormCellBoolean.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000003_Form_FormCellFloat64.Name = `Form-FormCellFloat64`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldFloat64]
	__GongStructShape__000003_Form_FormCellFloat64.Identifier = `ref_models.FormFieldFloat64`
	__GongStructShape__000003_Form_FormCellFloat64.ShowNbInstances = true
	__GongStructShape__000003_Form_FormCellFloat64.NbInstances = 0
	__GongStructShape__000003_Form_FormCellFloat64.Width = 240.000000
	__GongStructShape__000003_Form_FormCellFloat64.Heigth = 78.000000
	__GongStructShape__000003_Form_FormCellFloat64.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000004_Form_FormCellInt.Name = `Form-FormCellInt`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldInt]
	__GongStructShape__000004_Form_FormCellInt.Identifier = `ref_models.FormFieldInt`
	__GongStructShape__000004_Form_FormCellInt.ShowNbInstances = true
	__GongStructShape__000004_Form_FormCellInt.NbInstances = 0
	__GongStructShape__000004_Form_FormCellInt.Width = 240.000000
	__GongStructShape__000004_Form_FormCellInt.Heigth = 78.000000
	__GongStructShape__000004_Form_FormCellInt.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000005_Form_FormCellString.Name = `Form-FormCellString`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldString]
	__GongStructShape__000005_Form_FormCellString.Identifier = `ref_models.FormFieldString`
	__GongStructShape__000005_Form_FormCellString.ShowNbInstances = true
	__GongStructShape__000005_Form_FormCellString.NbInstances = 0
	__GongStructShape__000005_Form_FormCellString.Width = 240.000000
	__GongStructShape__000005_Form_FormCellString.Heigth = 78.000000
	__GongStructShape__000005_Form_FormCellString.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000006_Form_FormFieldDate.Name = `Form-FormFieldDate`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldDate]
	__GongStructShape__000006_Form_FormFieldDate.Identifier = `ref_models.FormFieldDate`
	__GongStructShape__000006_Form_FormFieldDate.ShowNbInstances = false
	__GongStructShape__000006_Form_FormFieldDate.NbInstances = 0
	__GongStructShape__000006_Form_FormFieldDate.Width = 240.000000
	__GongStructShape__000006_Form_FormFieldDate.Heigth = 78.000000
	__GongStructShape__000006_Form_FormFieldDate.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000007_Form_FormFieldTime.Name = `Form-FormFieldTime`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldTime]
	__GongStructShape__000007_Form_FormFieldTime.Identifier = `ref_models.FormFieldTime`
	__GongStructShape__000007_Form_FormFieldTime.ShowNbInstances = false
	__GongStructShape__000007_Form_FormFieldTime.NbInstances = 0
	__GongStructShape__000007_Form_FormFieldTime.Width = 240.000000
	__GongStructShape__000007_Form_FormFieldTime.Heigth = 93.000000
	__GongStructShape__000007_Form_FormFieldTime.IsSelected = false

	// Link values setup
	__Link__000000_FormCellBool.Name = `FormCellBool`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormField.FormFieldBool]
	__Link__000000_FormCellBool.Identifier = `ref_models.FormField.FormFieldBool`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldBoolean]
	__Link__000000_FormCellBool.Fieldtypename = `ref_models.FormFieldBoolean`
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

	//gong:ident [ref_models.FormField.FormFieldFloat64]
	__Link__000001_FormCellFloat64.Identifier = `ref_models.FormField.FormFieldFloat64`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldFloat64]
	__Link__000001_FormCellFloat64.Fieldtypename = `ref_models.FormFieldFloat64`
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

	//gong:ident [ref_models.FormField.FormFieldInt]
	__Link__000002_FormCellInt.Identifier = `ref_models.FormField.FormFieldInt`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldInt]
	__Link__000002_FormCellInt.Fieldtypename = `ref_models.FormFieldInt`
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

	//gong:ident [ref_models.FormField.FormFieldString]
	__Link__000003_FormCellString.Identifier = `ref_models.FormField.FormFieldString`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldString]
	__Link__000003_FormCellString.Fieldtypename = `ref_models.FormFieldString`
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

	//gong:ident [ref_models.FormGroup.FormDivs]
	__Link__000004_FormCells.Identifier = `ref_models.FormGroup.FormDivs`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormField]
	__Link__000004_FormCells.Fieldtypename = `ref_models.FormField`
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

	// Link values setup
	__Link__000005_FormFieldDate.Name = `FormFieldDate`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormField.FormFieldDate]
	__Link__000005_FormFieldDate.Identifier = `ref_models.FormField.FormFieldDate`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldDate]
	__Link__000005_FormFieldDate.Fieldtypename = `ref_models.FormFieldDate`
	__Link__000005_FormFieldDate.FieldOffsetX = -116.000000
	__Link__000005_FormFieldDate.FieldOffsetY = -22.000000
	__Link__000005_FormFieldDate.TargetMultiplicity = models.ZERO_ONE
	__Link__000005_FormFieldDate.TargetMultiplicityOffsetX = -49.000000
	__Link__000005_FormFieldDate.TargetMultiplicityOffsetY = 19.000000
	__Link__000005_FormFieldDate.SourceMultiplicity = models.MANY
	__Link__000005_FormFieldDate.SourceMultiplicityOffsetX = 10.000000
	__Link__000005_FormFieldDate.SourceMultiplicityOffsetY = -50.000000
	__Link__000005_FormFieldDate.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_FormFieldDate.StartRatio = 0.500000
	__Link__000005_FormFieldDate.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_FormFieldDate.EndRatio = 0.500000
	__Link__000005_FormFieldDate.CornerOffsetRatio = 1.731673

	// Link values setup
	__Link__000006_FormFieldTime.Name = `FormFieldTime`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormField.FormFieldTime]
	__Link__000006_FormFieldTime.Identifier = `ref_models.FormField.FormFieldTime`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldTime]
	__Link__000006_FormFieldTime.Fieldtypename = `ref_models.FormFieldTime`
	__Link__000006_FormFieldTime.FieldOffsetX = -123.000000
	__Link__000006_FormFieldTime.FieldOffsetY = -25.000000
	__Link__000006_FormFieldTime.TargetMultiplicity = models.ZERO_ONE
	__Link__000006_FormFieldTime.TargetMultiplicityOffsetX = -46.000000
	__Link__000006_FormFieldTime.TargetMultiplicityOffsetY = 24.000000
	__Link__000006_FormFieldTime.SourceMultiplicity = models.MANY
	__Link__000006_FormFieldTime.SourceMultiplicityOffsetX = 41.000000
	__Link__000006_FormFieldTime.SourceMultiplicityOffsetY = 26.000000
	__Link__000006_FormFieldTime.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_FormFieldTime.StartRatio = 0.500000
	__Link__000006_FormFieldTime.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_FormFieldTime.EndRatio = 0.500000
	__Link__000006_FormFieldTime.CornerOffsetRatio = 1.731673

	// Position values setup
	__Position__000000_Pos_Form_Form.X = 18.000000
	__Position__000000_Pos_Form_Form.Y = 50.000000
	__Position__000000_Pos_Form_Form.Name = `Pos-Form-Form`

	// Position values setup
	__Position__000001_Pos_Form_FormCell.X = 25.000000
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

	// Position values setup
	__Position__000006_Pos_Form_FormFieldDate.X = 639.000000
	__Position__000006_Pos_Form_FormFieldDate.Y = 515.000000
	__Position__000006_Pos_Form_FormFieldDate.Name = `Pos-Form-FormFieldDate`

	// Position values setup
	__Position__000007_Pos_Form_FormFieldTime.X = 647.000000
	__Position__000007_Pos_Form_FormFieldTime.Y = 610.000000
	__Position__000007_Pos_Form_FormFieldTime.Name = `Pos-Form-FormFieldTime`

	// Position values setup
	__Position__000008_Pos_Form_InputTypeEnum.X = 40.000000
	__Position__000008_Pos_Form_InputTypeEnum.Y = 429.000000
	__Position__000008_Pos_Form_InputTypeEnum.Name = `Pos-Form-InputTypeEnum`

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

	// Vertice values setup
	__Vertice__000005_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldDate.X = 707.000000
	__Vertice__000005_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldDate.Y = 413.500000
	__Vertice__000005_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldDate.Name = `Verticle in class diagram Form in middle between Form-FormCell and Form-FormFieldDate`

	// Vertice values setup
	__Vertice__000006_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldTime.X = 711.000000
	__Vertice__000006_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldTime.Y = 461.000000
	__Vertice__000006_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldTime.Name = `Verticle in class diagram Form in middle between Form-FormCell and Form-FormFieldTime`

	// Setup of pointers
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000000_Form_Form)
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000001_Form_FormCell)
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000002_Form_FormCellBoolean)
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000003_Form_FormCellFloat64)
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000004_Form_FormCellInt)
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000005_Form_FormCellString)
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000007_Form_FormFieldTime)
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000006_Form_FormFieldDate)
	__Classdiagram__000000_Form.GongEnumShapes = append(__Classdiagram__000000_Form.GongEnumShapes, __GongEnumShape__000000_Form_InputTypeEnum)
	__GongEnumShape__000000_Form_InputTypeEnum.Position = __Position__000008_Pos_Form_InputTypeEnum
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000012_Text)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000008_Password)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000007_Number)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000003_Email)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000011_Tel)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000001_Date)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000002_Datetime)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000013_Time)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000014_URL)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000010_Search)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000009_Range)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000000_Color)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000004_File)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000005_Hidden)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000006_Month)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000015_Week)
	__GongStructShape__000000_Form_Form.Position = __Position__000000_Pos_Form_Form
	__GongStructShape__000000_Form_Form.Fields = append(__GongStructShape__000000_Form_Form.Fields, __Field__000002_Name)
	__GongStructShape__000000_Form_Form.Links = append(__GongStructShape__000000_Form_Form.Links, __Link__000004_FormCells)
	__GongStructShape__000001_Form_FormCell.Position = __Position__000001_Pos_Form_FormCell
	__GongStructShape__000001_Form_FormCell.Fields = append(__GongStructShape__000001_Form_FormCell.Fields, __Field__000003_Name)
	__GongStructShape__000001_Form_FormCell.Fields = append(__GongStructShape__000001_Form_FormCell.Fields, __Field__000000_InputTypeEnum)
	__GongStructShape__000001_Form_FormCell.Fields = append(__GongStructShape__000001_Form_FormCell.Fields, __Field__000001_Label)
	__GongStructShape__000001_Form_FormCell.Fields = append(__GongStructShape__000001_Form_FormCell.Fields, __Field__000004_Placeholder)
	__GongStructShape__000001_Form_FormCell.Links = append(__GongStructShape__000001_Form_FormCell.Links, __Link__000000_FormCellBool)
	__GongStructShape__000001_Form_FormCell.Links = append(__GongStructShape__000001_Form_FormCell.Links, __Link__000001_FormCellFloat64)
	__GongStructShape__000001_Form_FormCell.Links = append(__GongStructShape__000001_Form_FormCell.Links, __Link__000002_FormCellInt)
	__GongStructShape__000001_Form_FormCell.Links = append(__GongStructShape__000001_Form_FormCell.Links, __Link__000003_FormCellString)
	__GongStructShape__000001_Form_FormCell.Links = append(__GongStructShape__000001_Form_FormCell.Links, __Link__000006_FormFieldTime)
	__GongStructShape__000001_Form_FormCell.Links = append(__GongStructShape__000001_Form_FormCell.Links, __Link__000005_FormFieldDate)
	__GongStructShape__000002_Form_FormCellBoolean.Position = __Position__000002_Pos_Form_FormCellBoolean
	__GongStructShape__000002_Form_FormCellBoolean.Fields = append(__GongStructShape__000002_Form_FormCellBoolean.Fields, __Field__000010_Value)
	__GongStructShape__000003_Form_FormCellFloat64.Position = __Position__000003_Pos_Form_FormCellFloat64
	__GongStructShape__000003_Form_FormCellFloat64.Fields = append(__GongStructShape__000003_Form_FormCellFloat64.Fields, __Field__000009_Value)
	__GongStructShape__000004_Form_FormCellInt.Position = __Position__000004_Pos_Form_FormCellInt
	__GongStructShape__000004_Form_FormCellInt.Fields = append(__GongStructShape__000004_Form_FormCellInt.Fields, __Field__000006_Value)
	__GongStructShape__000005_Form_FormCellString.Position = __Position__000005_Pos_Form_FormCellString
	__GongStructShape__000005_Form_FormCellString.Fields = append(__GongStructShape__000005_Form_FormCellString.Fields, __Field__000007_Value)
	__GongStructShape__000006_Form_FormFieldDate.Position = __Position__000006_Pos_Form_FormFieldDate
	__GongStructShape__000006_Form_FormFieldDate.Fields = append(__GongStructShape__000006_Form_FormFieldDate.Fields, __Field__000011_Value)
	__GongStructShape__000007_Form_FormFieldTime.Position = __Position__000007_Pos_Form_FormFieldTime
	__GongStructShape__000007_Form_FormFieldTime.Fields = append(__GongStructShape__000007_Form_FormFieldTime.Fields, __Field__000008_Value)
	__GongStructShape__000007_Form_FormFieldTime.Fields = append(__GongStructShape__000007_Form_FormFieldTime.Fields, __Field__000005_Step)
	__Link__000000_FormCellBool.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellBoolean
	__Link__000001_FormCellFloat64.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellFloat64
	__Link__000002_FormCellInt.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellInt
	__Link__000003_FormCellString.Middlevertice = __Vertice__000004_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellString
	__Link__000004_FormCells.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Form_in_middle_between_Form_Form_and_Form_FormCell
	__Link__000005_FormFieldDate.Middlevertice = __Vertice__000005_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldDate
	__Link__000006_FormFieldTime.Middlevertice = __Vertice__000006_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldTime
}
