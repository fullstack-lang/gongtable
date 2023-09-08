// generated code - do not edit
package data

import (
	"log"

	"github.com/gin-gonic/gin"

	gong_models "github.com/fullstack-lang/gong/go/models"
	form "github.com/fullstack-lang/gongtable/go/models"
	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/gongtable/go/models"
	"github.com/fullstack-lang/gongtable/go/orm"
)

type ButtonImplGongstruct struct {
	gongStruct         *gong_models.GongStruct
	Icon               gongtree_buttons.ButtonType
	tableStage         *form.StageStruct
	formStage          *form.StageStruct
	stageOfInterest    *models.StageStruct
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func NewButtonImplGongstruct(
	gongStruct *gong_models.GongStruct,
	icon gongtree_buttons.ButtonType,
	tableStage *form.StageStruct,
	formStage *form.StageStruct,
	stageOfInterest *models.StageStruct,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (buttonImplGongstruct *ButtonImplGongstruct) {

	buttonImplGongstruct = new(ButtonImplGongstruct)
	buttonImplGongstruct.Icon = icon
	buttonImplGongstruct.gongStruct = gongStruct
	buttonImplGongstruct.tableStage = tableStage
	buttonImplGongstruct.formStage = formStage
	buttonImplGongstruct.stageOfInterest = stageOfInterest
	buttonImplGongstruct.r = r
	buttonImplGongstruct.backRepoOfInterest = backRepoOfInterest

	return
}

func (buttonImpl *ButtonImplGongstruct) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	log.Println("ButtonImplGongstruct: ButtonUpdated")

	formStage := buttonImpl.formStage
	formStage.Reset()
	formStage.Commit()

	switch buttonImpl.gongStruct.Name {
	// insertion point
	case "Cell":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewCellFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		cell := new(models.Cell)
		FillUpForm(cell, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "CellBoolean":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewCellBooleanFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		cellboolean := new(models.CellBoolean)
		FillUpForm(cellboolean, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "CellFloat64":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewCellFloat64FormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		cellfloat64 := new(models.CellFloat64)
		FillUpForm(cellfloat64, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "CellIcon":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewCellIconFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		cellicon := new(models.CellIcon)
		FillUpForm(cellicon, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "CellInt":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewCellIntFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		cellint := new(models.CellInt)
		FillUpForm(cellint, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "CellString":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewCellStringFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		cellstring := new(models.CellString)
		FillUpForm(cellstring, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "CheckBox":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewCheckBoxFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		checkbox := new(models.CheckBox)
		FillUpForm(checkbox, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "DisplayedColumn":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewDisplayedColumnFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		displayedcolumn := new(models.DisplayedColumn)
		FillUpForm(displayedcolumn, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "FormDiv":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewFormDivFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		formdiv := new(models.FormDiv)
		FillUpForm(formdiv, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "FormEditAssocButton":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewFormEditAssocButtonFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		formeditassocbutton := new(models.FormEditAssocButton)
		FillUpForm(formeditassocbutton, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "FormField":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		formfield := new(models.FormField)
		FillUpForm(formfield, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "FormFieldDate":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldDateFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		formfielddate := new(models.FormFieldDate)
		FillUpForm(formfielddate, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "FormFieldDateTime":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldDateTimeFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		formfielddatetime := new(models.FormFieldDateTime)
		FillUpForm(formfielddatetime, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "FormFieldFloat64":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldFloat64FormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		formfieldfloat64 := new(models.FormFieldFloat64)
		FillUpForm(formfieldfloat64, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "FormFieldInt":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldIntFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		formfieldint := new(models.FormFieldInt)
		FillUpForm(formfieldint, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "FormFieldSelect":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldSelectFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		formfieldselect := new(models.FormFieldSelect)
		FillUpForm(formfieldselect, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "FormFieldString":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldStringFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		formfieldstring := new(models.FormFieldString)
		FillUpForm(formfieldstring, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "FormFieldTime":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldTimeFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		formfieldtime := new(models.FormFieldTime)
		FillUpForm(formfieldtime, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "FormGroup":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewFormGroupFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		formgroup := new(models.FormGroup)
		FillUpForm(formgroup, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "FormSortAssocButton":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewFormSortAssocButtonFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		formsortassocbutton := new(models.FormSortAssocButton)
		FillUpForm(formsortassocbutton, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "Option":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewOptionFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		option := new(models.Option)
		FillUpForm(option, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "Row":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewRowFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		row := new(models.Row)
		FillUpForm(row, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "Table":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewTableFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		table := new(models.Table)
		FillUpForm(table, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	}
	formStage.Commit()
}

func FillUpForm[T models.Gongstruct](
	instance *T,
	stageOfInterest *models.StageStruct,
	formStage *form.StageStruct,
	formGroup *form.FormGroup,
	r *gin.Engine,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Cell:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		AssociationFieldToForm("CellString", instanceWithInferedType.CellString, stageOfInterest, formStage, formGroup)
		AssociationFieldToForm("CellFloat64", instanceWithInferedType.CellFloat64, stageOfInterest, formStage, formGroup)
		AssociationFieldToForm("CellInt", instanceWithInferedType.CellInt, stageOfInterest, formStage, formGroup)
		AssociationFieldToForm("CellBool", instanceWithInferedType.CellBool, stageOfInterest, formStage, formGroup)
		AssociationFieldToForm("CellIcon", instanceWithInferedType.CellIcon, stageOfInterest, formStage, formGroup)

	case *models.CellBoolean:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, formStage, formGroup)

	case *models.CellFloat64:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, formStage, formGroup)

	case *models.CellIcon:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Icon", instanceWithInferedType.Icon, instanceWithInferedType, formStage, formGroup)

	case *models.CellInt:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, formStage, formGroup)

	case *models.CellString:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, formStage, formGroup)

	case *models.CheckBox:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, formStage, formGroup)

	case *models.DisplayedColumn:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)

	case *models.FormDiv:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("FormFields", instanceWithInferedType, &instanceWithInferedType.FormFields, stageOfInterest, formStage, formGroup, r)
		AssociationSliceToForm("CheckBoxs", instanceWithInferedType, &instanceWithInferedType.CheckBoxs, stageOfInterest, formStage, formGroup, r)
		AssociationFieldToForm("FormEditAssocButton", instanceWithInferedType.FormEditAssocButton, stageOfInterest, formStage, formGroup)
		AssociationFieldToForm("FormSortAssocButton", instanceWithInferedType.FormSortAssocButton, stageOfInterest, formStage, formGroup)

	case *models.FormEditAssocButton:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, formStage, formGroup)

	case *models.FormField:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		EnumTypeStringToForm("InputTypeEnum", instanceWithInferedType.InputTypeEnum, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Placeholder", instanceWithInferedType.Placeholder, instanceWithInferedType, formStage, formGroup)
		AssociationFieldToForm("FormFieldString", instanceWithInferedType.FormFieldString, stageOfInterest, formStage, formGroup)
		AssociationFieldToForm("FormFieldFloat64", instanceWithInferedType.FormFieldFloat64, stageOfInterest, formStage, formGroup)
		AssociationFieldToForm("FormFieldInt", instanceWithInferedType.FormFieldInt, stageOfInterest, formStage, formGroup)
		AssociationFieldToForm("FormFieldDate", instanceWithInferedType.FormFieldDate, stageOfInterest, formStage, formGroup)
		AssociationFieldToForm("FormFieldTime", instanceWithInferedType.FormFieldTime, stageOfInterest, formStage, formGroup)
		AssociationFieldToForm("FormFieldDateTime", instanceWithInferedType.FormFieldDateTime, stageOfInterest, formStage, formGroup)
		AssociationFieldToForm("FormFieldSelect", instanceWithInferedType.FormFieldSelect, stageOfInterest, formStage, formGroup)
		BasicFieldtoForm("HasBespokeWidth", instanceWithInferedType.HasBespokeWidth, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("BespokeWidthPx", instanceWithInferedType.BespokeWidthPx, instanceWithInferedType, formStage, formGroup)

	case *models.FormFieldDate:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)

	case *models.FormFieldDateTime:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)

	case *models.FormFieldFloat64:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("HasMinValidator", instanceWithInferedType.HasMinValidator, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("MinValue", instanceWithInferedType.MinValue, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("HasMaxValidator", instanceWithInferedType.HasMaxValidator, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("MaxValue", instanceWithInferedType.MaxValue, instanceWithInferedType, formStage, formGroup)

	case *models.FormFieldInt:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("HasMinValidator", instanceWithInferedType.HasMinValidator, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("MinValue", instanceWithInferedType.MinValue, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("HasMaxValidator", instanceWithInferedType.HasMaxValidator, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("MaxValue", instanceWithInferedType.MaxValue, instanceWithInferedType, formStage, formGroup)

	case *models.FormFieldSelect:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		AssociationFieldToForm("Value", instanceWithInferedType.Value, stageOfInterest, formStage, formGroup)
		AssociationSliceToForm("Options", instanceWithInferedType, &instanceWithInferedType.Options, stageOfInterest, formStage, formGroup, r)
		BasicFieldtoForm("CanBeEmpty", instanceWithInferedType.CanBeEmpty, instanceWithInferedType, formStage, formGroup)

	case *models.FormFieldString:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, formStage, formGroup)

	case *models.FormFieldTime:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Step", instanceWithInferedType.Step, instanceWithInferedType, formStage, formGroup)

	case *models.FormGroup:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("FormDivs", instanceWithInferedType, &instanceWithInferedType.FormDivs, stageOfInterest, formStage, formGroup, r)

	case *models.FormSortAssocButton:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, formStage, formGroup)

	case *models.Option:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)

	case *models.Row:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("Cells", instanceWithInferedType, &instanceWithInferedType.Cells, stageOfInterest, formStage, formGroup, r)
		BasicFieldtoForm("IsChecked", instanceWithInferedType.IsChecked, instanceWithInferedType, formStage, formGroup)

	case *models.Table:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("DisplayedColumns", instanceWithInferedType, &instanceWithInferedType.DisplayedColumns, stageOfInterest, formStage, formGroup, r)
		AssociationSliceToForm("Rows", instanceWithInferedType, &instanceWithInferedType.Rows, stageOfInterest, formStage, formGroup, r)
		BasicFieldtoForm("HasFiltering", instanceWithInferedType.HasFiltering, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("HasColumnSorting", instanceWithInferedType.HasColumnSorting, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("HasPaginator", instanceWithInferedType.HasPaginator, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("HasCheckableRows", instanceWithInferedType.HasCheckableRows, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("HasSaveButton", instanceWithInferedType.HasSaveButton, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("CanDragDropRows", instanceWithInferedType.CanDragDropRows, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("HasCloseButton", instanceWithInferedType.HasCloseButton, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("SavingInProgress", instanceWithInferedType.SavingInProgress, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("NbOfStickyColumns", instanceWithInferedType.NbOfStickyColumns, instanceWithInferedType, formStage, formGroup)

	default:
		_ = instanceWithInferedType
	}
}

