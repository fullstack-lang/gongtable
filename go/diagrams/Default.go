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

	"ref_models.CheckBox": &(ref_models.CheckBox{}),

	"ref_models.CheckBox.Name": (ref_models.CheckBox{}).Name,

	"ref_models.CheckBox.Value": (ref_models.CheckBox{}).Value,

	"ref_models.Color": ref_models.Color,

	"ref_models.Date": ref_models.Date,

	"ref_models.Datetime": ref_models.Datetime,

	"ref_models.DisplayedColumn": &(ref_models.DisplayedColumn{}),

	"ref_models.DisplayedColumn.Name": (ref_models.DisplayedColumn{}).Name,

	"ref_models.Email": ref_models.Email,

	"ref_models.File": ref_models.File,

	"ref_models.FormDiv": &(ref_models.FormDiv{}),

	"ref_models.FormDiv.CheckBoxs": (ref_models.FormDiv{}).CheckBoxs,

	"ref_models.FormDiv.FormEditAssocButton": (ref_models.FormDiv{}).FormEditAssocButton,

	"ref_models.FormDiv.FormFields": (ref_models.FormDiv{}).FormFields,

	"ref_models.FormDiv.FormSortAssocButton": (ref_models.FormDiv{}).FormSortAssocButton,

	"ref_models.FormDiv.Name": (ref_models.FormDiv{}).Name,

	"ref_models.FormEditAssocButton": &(ref_models.FormEditAssocButton{}),

	"ref_models.FormEditAssocButton.Label": (ref_models.FormEditAssocButton{}).Label,

	"ref_models.FormEditAssocButton.Name": (ref_models.FormEditAssocButton{}).Name,

	"ref_models.FormField": &(ref_models.FormField{}),

	"ref_models.FormField.BespokeWidthPx": (ref_models.FormField{}).BespokeWidthPx,

	"ref_models.FormField.FormFieldDate": (ref_models.FormField{}).FormFieldDate,

	"ref_models.FormField.FormFieldDateTime": (ref_models.FormField{}).FormFieldDateTime,

	"ref_models.FormField.FormFieldFloat64": (ref_models.FormField{}).FormFieldFloat64,

	"ref_models.FormField.FormFieldInt": (ref_models.FormField{}).FormFieldInt,

	"ref_models.FormField.FormFieldSelect": (ref_models.FormField{}).FormFieldSelect,

	"ref_models.FormField.FormFieldString": (ref_models.FormField{}).FormFieldString,

	"ref_models.FormField.FormFieldTime": (ref_models.FormField{}).FormFieldTime,

	"ref_models.FormField.HasBespokeWidth": (ref_models.FormField{}).HasBespokeWidth,

	"ref_models.FormField.InputTypeEnum": (ref_models.FormField{}).InputTypeEnum,

	"ref_models.FormField.Label": (ref_models.FormField{}).Label,

	"ref_models.FormField.Name": (ref_models.FormField{}).Name,

	"ref_models.FormField.Placeholder": (ref_models.FormField{}).Placeholder,

	"ref_models.FormFieldDate": &(ref_models.FormFieldDate{}),

	"ref_models.FormFieldDate.Name": (ref_models.FormFieldDate{}).Name,

	"ref_models.FormFieldDate.Value": (ref_models.FormFieldDate{}).Value,

	"ref_models.FormFieldDateTime": &(ref_models.FormFieldDateTime{}),

	"ref_models.FormFieldDateTime.Name": (ref_models.FormFieldDateTime{}).Name,

	"ref_models.FormFieldDateTime.Value": (ref_models.FormFieldDateTime{}).Value,

	"ref_models.FormFieldFloat64": &(ref_models.FormFieldFloat64{}),

	"ref_models.FormFieldFloat64.HasMaxValidator": (ref_models.FormFieldFloat64{}).HasMaxValidator,

	"ref_models.FormFieldFloat64.HasMinValidator": (ref_models.FormFieldFloat64{}).HasMinValidator,

	"ref_models.FormFieldFloat64.MaxValue": (ref_models.FormFieldFloat64{}).MaxValue,

	"ref_models.FormFieldFloat64.MinValue": (ref_models.FormFieldFloat64{}).MinValue,

	"ref_models.FormFieldFloat64.Name": (ref_models.FormFieldFloat64{}).Name,

	"ref_models.FormFieldFloat64.Value": (ref_models.FormFieldFloat64{}).Value,

	"ref_models.FormFieldInt": &(ref_models.FormFieldInt{}),

	"ref_models.FormFieldInt.HasMaxValidator": (ref_models.FormFieldInt{}).HasMaxValidator,

	"ref_models.FormFieldInt.HasMinValidator": (ref_models.FormFieldInt{}).HasMinValidator,

	"ref_models.FormFieldInt.MaxValue": (ref_models.FormFieldInt{}).MaxValue,

	"ref_models.FormFieldInt.MinValue": (ref_models.FormFieldInt{}).MinValue,

	"ref_models.FormFieldInt.Name": (ref_models.FormFieldInt{}).Name,

	"ref_models.FormFieldInt.Value": (ref_models.FormFieldInt{}).Value,

	"ref_models.FormFieldSelect": &(ref_models.FormFieldSelect{}),

	"ref_models.FormFieldSelect.CanBeEmpty": (ref_models.FormFieldSelect{}).CanBeEmpty,

	"ref_models.FormFieldSelect.Name": (ref_models.FormFieldSelect{}).Name,

	"ref_models.FormFieldSelect.Options": (ref_models.FormFieldSelect{}).Options,

	"ref_models.FormFieldSelect.Value": (ref_models.FormFieldSelect{}).Value,

	"ref_models.FormFieldString": &(ref_models.FormFieldString{}),

	"ref_models.FormFieldString.IsTextArea": (ref_models.FormFieldString{}).IsTextArea,

	"ref_models.FormFieldString.Name": (ref_models.FormFieldString{}).Name,

	"ref_models.FormFieldString.Value": (ref_models.FormFieldString{}).Value,

	"ref_models.FormFieldTime": &(ref_models.FormFieldTime{}),

	"ref_models.FormFieldTime.Name": (ref_models.FormFieldTime{}).Name,

	"ref_models.FormFieldTime.Step": (ref_models.FormFieldTime{}).Step,

	"ref_models.FormFieldTime.Value": (ref_models.FormFieldTime{}).Value,

	"ref_models.FormGroup": &(ref_models.FormGroup{}),

	"ref_models.FormGroup.FormDivs": (ref_models.FormGroup{}).FormDivs,

	"ref_models.FormGroup.Label": (ref_models.FormGroup{}).Label,

	"ref_models.FormGroup.Name": (ref_models.FormGroup{}).Name,

	"ref_models.FormGroupDefaultName": ref_models.FormGroupDefaultName,

	"ref_models.FormGroupName": ref_models.FormGroupName(""),

	"ref_models.FormSortAssocButton": &(ref_models.FormSortAssocButton{}),

	"ref_models.FormSortAssocButton.Label": (ref_models.FormSortAssocButton{}).Label,

	"ref_models.FormSortAssocButton.Name": (ref_models.FormSortAssocButton{}).Name,

	"ref_models.GeneratedTableStackName": ref_models.GeneratedTableStackName,

	"ref_models.Hidden": ref_models.Hidden,

	"ref_models.InputTypeEnum": ref_models.InputTypeEnum(""),

	"ref_models.ManualyEditedFormStackName": ref_models.ManualyEditedFormStackName,

	"ref_models.ManualyEditedTableStackName": ref_models.ManualyEditedTableStackName,

	"ref_models.Month": ref_models.Month,

	"ref_models.Number": ref_models.Number,

	"ref_models.Option": &(ref_models.Option{}),

	"ref_models.Option.Name": (ref_models.Option{}).Name,

	"ref_models.Password": ref_models.Password,

	"ref_models.Range": ref_models.Range,

	"ref_models.Row": &(ref_models.Row{}),

	"ref_models.Row.Cells": (ref_models.Row{}).Cells,

	"ref_models.Row.IsChecked": (ref_models.Row{}).IsChecked,

	"ref_models.Row.Name": (ref_models.Row{}).Name,

	"ref_models.Search": ref_models.Search,

	"ref_models.StackNamePostFixForTableForAssociation": ref_models.StackNamePostFixForTableForAssociation,

	"ref_models.StackNamePostFixForTableForAssociationSorting": ref_models.StackNamePostFixForTableForAssociationSorting,

	"ref_models.StackNamePostFixForTableForMainForm": ref_models.StackNamePostFixForTableForMainForm,

	"ref_models.StackNamePostFixForTableForMainTable": ref_models.StackNamePostFixForTableForMainTable,

	"ref_models.StackNamePostFixForTableForMainTree": ref_models.StackNamePostFixForTableForMainTree,

	"ref_models.Table": &(ref_models.Table{}),

	"ref_models.Table.CanDragDropRows": (ref_models.Table{}).CanDragDropRows,

	"ref_models.Table.DisplayedColumns": (ref_models.Table{}).DisplayedColumns,

	"ref_models.Table.HasCheckableRows": (ref_models.Table{}).HasCheckableRows,

	"ref_models.Table.HasCloseButton": (ref_models.Table{}).HasCloseButton,

	"ref_models.Table.HasColumnSorting": (ref_models.Table{}).HasColumnSorting,

	"ref_models.Table.HasFiltering": (ref_models.Table{}).HasFiltering,

	"ref_models.Table.HasPaginator": (ref_models.Table{}).HasPaginator,

	"ref_models.Table.HasSaveButton": (ref_models.Table{}).HasSaveButton,

	"ref_models.Table.Name": (ref_models.Table{}).Name,

	"ref_models.Table.NbOfStickyColumns": (ref_models.Table{}).NbOfStickyColumns,

	"ref_models.Table.Rows": (ref_models.Table{}).Rows,

	"ref_models.Table.SavingInProgress": (ref_models.Table{}).SavingInProgress,

	"ref_models.TableDefaultName": ref_models.TableDefaultName,

	"ref_models.TableExtraNameEnum": ref_models.TableExtraNameEnum(""),

	"ref_models.TableExtraPathEnum": ref_models.TableExtraPathEnum(""),

	"ref_models.TableName": ref_models.TableName(""),

	"ref_models.TableSelectExtraName": ref_models.TableSelectExtraName,

	"ref_models.TableSortExtraName": ref_models.TableSortExtraName,

	"ref_models.TableTestNameEnum": ref_models.TableTestNameEnum(""),

	"ref_models.Tel": ref_models.Tel,

	"ref_models.Text": ref_models.Text,

	"ref_models.Time": ref_models.Time,

	"ref_models.URL": ref_models.URL,

	"ref_models.Week": ref_models.Week,
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

	// Declarations of staged instances of Link

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	// Setup of pointers
}
