// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongtable/go/models"
)

func FillUpFormFromGongstructName(
	playground *Playground,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := playground.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = "New"
	} else {
		prefix = "Update"
	}

	switch gongstructName {
	// insertion point
	case "Cell":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Cell Form",
			OnSave: NewCellFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		cell := new(models.Cell)
		FillUpForm(cell, formGroup, playground)
	case "CellBoolean":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " CellBoolean Form",
			OnSave: NewCellBooleanFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		cellboolean := new(models.CellBoolean)
		FillUpForm(cellboolean, formGroup, playground)
	case "CellFloat64":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " CellFloat64 Form",
			OnSave: NewCellFloat64FormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		cellfloat64 := new(models.CellFloat64)
		FillUpForm(cellfloat64, formGroup, playground)
	case "CellIcon":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " CellIcon Form",
			OnSave: NewCellIconFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		cellicon := new(models.CellIcon)
		FillUpForm(cellicon, formGroup, playground)
	case "CellInt":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " CellInt Form",
			OnSave: NewCellIntFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		cellint := new(models.CellInt)
		FillUpForm(cellint, formGroup, playground)
	case "CellString":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " CellString Form",
			OnSave: NewCellStringFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		cellstring := new(models.CellString)
		FillUpForm(cellstring, formGroup, playground)
	case "CheckBox":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " CheckBox Form",
			OnSave: NewCheckBoxFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		checkbox := new(models.CheckBox)
		FillUpForm(checkbox, formGroup, playground)
	case "DisplayedColumn":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " DisplayedColumn Form",
			OnSave: NewDisplayedColumnFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		displayedcolumn := new(models.DisplayedColumn)
		FillUpForm(displayedcolumn, formGroup, playground)
	case "FormDiv":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormDiv Form",
			OnSave: NewFormDivFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		formdiv := new(models.FormDiv)
		FillUpForm(formdiv, formGroup, playground)
	case "FormEditAssocButton":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormEditAssocButton Form",
			OnSave: NewFormEditAssocButtonFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		formeditassocbutton := new(models.FormEditAssocButton)
		FillUpForm(formeditassocbutton, formGroup, playground)
	case "FormField":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormField Form",
			OnSave: NewFormFieldFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		formfield := new(models.FormField)
		FillUpForm(formfield, formGroup, playground)
	case "FormFieldDate":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormFieldDate Form",
			OnSave: NewFormFieldDateFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		formfielddate := new(models.FormFieldDate)
		FillUpForm(formfielddate, formGroup, playground)
	case "FormFieldDateTime":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormFieldDateTime Form",
			OnSave: NewFormFieldDateTimeFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		formfielddatetime := new(models.FormFieldDateTime)
		FillUpForm(formfielddatetime, formGroup, playground)
	case "FormFieldFloat64":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormFieldFloat64 Form",
			OnSave: NewFormFieldFloat64FormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		formfieldfloat64 := new(models.FormFieldFloat64)
		FillUpForm(formfieldfloat64, formGroup, playground)
	case "FormFieldInt":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormFieldInt Form",
			OnSave: NewFormFieldIntFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		formfieldint := new(models.FormFieldInt)
		FillUpForm(formfieldint, formGroup, playground)
	case "FormFieldSelect":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormFieldSelect Form",
			OnSave: NewFormFieldSelectFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		formfieldselect := new(models.FormFieldSelect)
		FillUpForm(formfieldselect, formGroup, playground)
	case "FormFieldString":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormFieldString Form",
			OnSave: NewFormFieldStringFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		formfieldstring := new(models.FormFieldString)
		FillUpForm(formfieldstring, formGroup, playground)
	case "FormFieldTime":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormFieldTime Form",
			OnSave: NewFormFieldTimeFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		formfieldtime := new(models.FormFieldTime)
		FillUpForm(formfieldtime, formGroup, playground)
	case "FormGroup":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormGroup Form",
			OnSave: NewFormGroupFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		formgroup := new(models.FormGroup)
		FillUpForm(formgroup, formGroup, playground)
	case "FormSortAssocButton":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormSortAssocButton Form",
			OnSave: NewFormSortAssocButtonFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		formsortassocbutton := new(models.FormSortAssocButton)
		FillUpForm(formsortassocbutton, formGroup, playground)
	case "Option":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Option Form",
			OnSave: NewOptionFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		option := new(models.Option)
		FillUpForm(option, formGroup, playground)
	case "Row":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Row Form",
			OnSave: NewRowFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		row := new(models.Row)
		FillUpForm(row, formGroup, playground)
	case "Table":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Table Form",
			OnSave: NewTableFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		table := new(models.Table)
		FillUpForm(table, formGroup, playground)
	}
	formStage.Commit()
}
