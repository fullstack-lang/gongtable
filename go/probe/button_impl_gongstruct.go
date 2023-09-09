// generated code - do not edit
package probe

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	form "github.com/fullstack-lang/gongtable/go/models"
	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/gongtable/go/models"
)

type ButtonImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	Icon       gongtree_buttons.ButtonType
	playground *Playground
}

func NewButtonImplGongstruct(
	gongStruct *gong_models.GongStruct,
	icon gongtree_buttons.ButtonType,
	playground *Playground,
) (buttonImplGongstruct *ButtonImplGongstruct) {

	buttonImplGongstruct = new(ButtonImplGongstruct)
	buttonImplGongstruct.Icon = icon
	buttonImplGongstruct.gongStruct = gongStruct
	buttonImplGongstruct.playground = playground

	return
}

func (buttonImpl *ButtonImplGongstruct) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	log.Println("ButtonImplGongstruct: ButtonUpdated")

	formStage := buttonImpl.playground.formStage
	formStage.Reset()
	formStage.Commit()

	switch buttonImpl.gongStruct.Name {
	// insertion point
	case "Cell":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewCellFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		cell := new(models.Cell)
		FillUpForm(cell, formGroup, buttonImpl.playground)
	case "CellBoolean":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewCellBooleanFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		cellboolean := new(models.CellBoolean)
		FillUpForm(cellboolean, formGroup, buttonImpl.playground)
	case "CellFloat64":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewCellFloat64FormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		cellfloat64 := new(models.CellFloat64)
		FillUpForm(cellfloat64, formGroup, buttonImpl.playground)
	case "CellIcon":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewCellIconFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		cellicon := new(models.CellIcon)
		FillUpForm(cellicon, formGroup, buttonImpl.playground)
	case "CellInt":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewCellIntFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		cellint := new(models.CellInt)
		FillUpForm(cellint, formGroup, buttonImpl.playground)
	case "CellString":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewCellStringFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		cellstring := new(models.CellString)
		FillUpForm(cellstring, formGroup, buttonImpl.playground)
	case "CheckBox":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewCheckBoxFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		checkbox := new(models.CheckBox)
		FillUpForm(checkbox, formGroup, buttonImpl.playground)
	case "DisplayedColumn":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewDisplayedColumnFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		displayedcolumn := new(models.DisplayedColumn)
		FillUpForm(displayedcolumn, formGroup, buttonImpl.playground)
	case "FormDiv":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewFormDivFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		formdiv := new(models.FormDiv)
		FillUpForm(formdiv, formGroup, buttonImpl.playground)
	case "FormEditAssocButton":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewFormEditAssocButtonFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		formeditassocbutton := new(models.FormEditAssocButton)
		FillUpForm(formeditassocbutton, formGroup, buttonImpl.playground)
	case "FormField":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		formfield := new(models.FormField)
		FillUpForm(formfield, formGroup, buttonImpl.playground)
	case "FormFieldDate":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldDateFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		formfielddate := new(models.FormFieldDate)
		FillUpForm(formfielddate, formGroup, buttonImpl.playground)
	case "FormFieldDateTime":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldDateTimeFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		formfielddatetime := new(models.FormFieldDateTime)
		FillUpForm(formfielddatetime, formGroup, buttonImpl.playground)
	case "FormFieldFloat64":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldFloat64FormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		formfieldfloat64 := new(models.FormFieldFloat64)
		FillUpForm(formfieldfloat64, formGroup, buttonImpl.playground)
	case "FormFieldInt":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldIntFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		formfieldint := new(models.FormFieldInt)
		FillUpForm(formfieldint, formGroup, buttonImpl.playground)
	case "FormFieldSelect":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldSelectFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		formfieldselect := new(models.FormFieldSelect)
		FillUpForm(formfieldselect, formGroup, buttonImpl.playground)
	case "FormFieldString":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldStringFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		formfieldstring := new(models.FormFieldString)
		FillUpForm(formfieldstring, formGroup, buttonImpl.playground)
	case "FormFieldTime":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldTimeFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		formfieldtime := new(models.FormFieldTime)
		FillUpForm(formfieldtime, formGroup, buttonImpl.playground)
	case "FormGroup":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewFormGroupFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		formgroup := new(models.FormGroup)
		FillUpForm(formgroup, formGroup, buttonImpl.playground)
	case "FormSortAssocButton":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewFormSortAssocButtonFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		formsortassocbutton := new(models.FormSortAssocButton)
		FillUpForm(formsortassocbutton, formGroup, buttonImpl.playground)
	case "Option":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewOptionFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		option := new(models.Option)
		FillUpForm(option, formGroup, buttonImpl.playground)
	case "Row":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewRowFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		row := new(models.Row)
		FillUpForm(row, formGroup, buttonImpl.playground)
	case "Table":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewTableFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		table := new(models.Table)
		FillUpForm(table, formGroup, buttonImpl.playground)
	}
	formStage.Commit()
}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	playground *Playground,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Cell:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("CellString", instanceWithInferedType.CellString, formGroup, playground)
		AssociationFieldToForm("CellFloat64", instanceWithInferedType.CellFloat64, formGroup, playground)
		AssociationFieldToForm("CellInt", instanceWithInferedType.CellInt, formGroup, playground)
		AssociationFieldToForm("CellBool", instanceWithInferedType.CellBool, formGroup, playground)
		AssociationFieldToForm("CellIcon", instanceWithInferedType.CellIcon, formGroup, playground)

	case *models.CellBoolean:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup)

	case *models.CellFloat64:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup)

	case *models.CellIcon:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Icon", instanceWithInferedType.Icon, instanceWithInferedType, playground.formStage, formGroup)

	case *models.CellInt:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup)

	case *models.CellString:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup)

	case *models.CheckBox:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup)

	case *models.DisplayedColumn:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)

	case *models.FormDiv:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("FormFields", instanceWithInferedType, &instanceWithInferedType.FormFields, formGroup, playground)
		AssociationSliceToForm("CheckBoxs", instanceWithInferedType, &instanceWithInferedType.CheckBoxs, formGroup, playground)
		AssociationFieldToForm("FormEditAssocButton", instanceWithInferedType.FormEditAssocButton, formGroup, playground)
		AssociationFieldToForm("FormSortAssocButton", instanceWithInferedType.FormSortAssocButton, formGroup, playground)

	case *models.FormEditAssocButton:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, playground.formStage, formGroup)

	case *models.FormField:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeStringToForm("InputTypeEnum", instanceWithInferedType.InputTypeEnum, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Placeholder", instanceWithInferedType.Placeholder, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("FormFieldString", instanceWithInferedType.FormFieldString, formGroup, playground)
		AssociationFieldToForm("FormFieldFloat64", instanceWithInferedType.FormFieldFloat64, formGroup, playground)
		AssociationFieldToForm("FormFieldInt", instanceWithInferedType.FormFieldInt, formGroup, playground)
		AssociationFieldToForm("FormFieldDate", instanceWithInferedType.FormFieldDate, formGroup, playground)
		AssociationFieldToForm("FormFieldTime", instanceWithInferedType.FormFieldTime, formGroup, playground)
		AssociationFieldToForm("FormFieldDateTime", instanceWithInferedType.FormFieldDateTime, formGroup, playground)
		AssociationFieldToForm("FormFieldSelect", instanceWithInferedType.FormFieldSelect, formGroup, playground)
		BasicFieldtoForm("HasBespokeWidth", instanceWithInferedType.HasBespokeWidth, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("BespokeWidthPx", instanceWithInferedType.BespokeWidthPx, instanceWithInferedType, playground.formStage, formGroup)

	case *models.FormFieldDate:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)

	case *models.FormFieldDateTime:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)

	case *models.FormFieldFloat64:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("HasMinValidator", instanceWithInferedType.HasMinValidator, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("MinValue", instanceWithInferedType.MinValue, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("HasMaxValidator", instanceWithInferedType.HasMaxValidator, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("MaxValue", instanceWithInferedType.MaxValue, instanceWithInferedType, playground.formStage, formGroup)

	case *models.FormFieldInt:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("HasMinValidator", instanceWithInferedType.HasMinValidator, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("MinValue", instanceWithInferedType.MinValue, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("HasMaxValidator", instanceWithInferedType.HasMaxValidator, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("MaxValue", instanceWithInferedType.MaxValue, instanceWithInferedType, playground.formStage, formGroup)

	case *models.FormFieldSelect:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("Value", instanceWithInferedType.Value, formGroup, playground)
		AssociationSliceToForm("Options", instanceWithInferedType, &instanceWithInferedType.Options, formGroup, playground)
		BasicFieldtoForm("CanBeEmpty", instanceWithInferedType.CanBeEmpty, instanceWithInferedType, playground.formStage, formGroup)

	case *models.FormFieldString:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup)

	case *models.FormFieldTime:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Step", instanceWithInferedType.Step, instanceWithInferedType, playground.formStage, formGroup)

	case *models.FormGroup:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("FormDivs", instanceWithInferedType, &instanceWithInferedType.FormDivs, formGroup, playground)

	case *models.FormSortAssocButton:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, playground.formStage, formGroup)

	case *models.Option:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)

	case *models.Row:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("Cells", instanceWithInferedType, &instanceWithInferedType.Cells, formGroup, playground)
		BasicFieldtoForm("IsChecked", instanceWithInferedType.IsChecked, instanceWithInferedType, playground.formStage, formGroup)

	case *models.Table:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("DisplayedColumns", instanceWithInferedType, &instanceWithInferedType.DisplayedColumns, formGroup, playground)
		AssociationSliceToForm("Rows", instanceWithInferedType, &instanceWithInferedType.Rows, formGroup, playground)
		BasicFieldtoForm("HasFiltering", instanceWithInferedType.HasFiltering, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("HasColumnSorting", instanceWithInferedType.HasColumnSorting, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("HasPaginator", instanceWithInferedType.HasPaginator, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("HasCheckableRows", instanceWithInferedType.HasCheckableRows, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("HasSaveButton", instanceWithInferedType.HasSaveButton, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("CanDragDropRows", instanceWithInferedType.CanDragDropRows, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("HasCloseButton", instanceWithInferedType.HasCloseButton, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("SavingInProgress", instanceWithInferedType.SavingInProgress, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("NbOfStickyColumns", instanceWithInferedType.NbOfStickyColumns, instanceWithInferedType, playground.formStage, formGroup)

	default:
		_ = instanceWithInferedType
	}
}

