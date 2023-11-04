// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongtable/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
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
			OnSave: __gong__New__CellFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		cell := new(models.Cell)
		FillUpForm(cell, formGroup, probe)
	case "CellBoolean":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " CellBoolean Form",
			OnSave: __gong__New__CellBooleanFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		cellboolean := new(models.CellBoolean)
		FillUpForm(cellboolean, formGroup, probe)
	case "CellFloat64":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " CellFloat64 Form",
			OnSave: __gong__New__CellFloat64FormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		cellfloat64 := new(models.CellFloat64)
		FillUpForm(cellfloat64, formGroup, probe)
	case "CellIcon":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " CellIcon Form",
			OnSave: __gong__New__CellIconFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		cellicon := new(models.CellIcon)
		FillUpForm(cellicon, formGroup, probe)
	case "CellInt":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " CellInt Form",
			OnSave: __gong__New__CellIntFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		cellint := new(models.CellInt)
		FillUpForm(cellint, formGroup, probe)
	case "CellString":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " CellString Form",
			OnSave: __gong__New__CellStringFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		cellstring := new(models.CellString)
		FillUpForm(cellstring, formGroup, probe)
	case "CheckBox":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " CheckBox Form",
			OnSave: __gong__New__CheckBoxFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		checkbox := new(models.CheckBox)
		FillUpForm(checkbox, formGroup, probe)
	case "DisplayedColumn":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " DisplayedColumn Form",
			OnSave: __gong__New__DisplayedColumnFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		displayedcolumn := new(models.DisplayedColumn)
		FillUpForm(displayedcolumn, formGroup, probe)
	case "FormDiv":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormDiv Form",
			OnSave: __gong__New__FormDivFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		formdiv := new(models.FormDiv)
		FillUpForm(formdiv, formGroup, probe)
	case "FormEditAssocButton":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormEditAssocButton Form",
			OnSave: __gong__New__FormEditAssocButtonFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		formeditassocbutton := new(models.FormEditAssocButton)
		FillUpForm(formeditassocbutton, formGroup, probe)
	case "FormField":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormField Form",
			OnSave: __gong__New__FormFieldFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		formfield := new(models.FormField)
		FillUpForm(formfield, formGroup, probe)
	case "FormFieldDate":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormFieldDate Form",
			OnSave: __gong__New__FormFieldDateFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		formfielddate := new(models.FormFieldDate)
		FillUpForm(formfielddate, formGroup, probe)
	case "FormFieldDateTime":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormFieldDateTime Form",
			OnSave: __gong__New__FormFieldDateTimeFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		formfielddatetime := new(models.FormFieldDateTime)
		FillUpForm(formfielddatetime, formGroup, probe)
	case "FormFieldFloat64":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormFieldFloat64 Form",
			OnSave: __gong__New__FormFieldFloat64FormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		formfieldfloat64 := new(models.FormFieldFloat64)
		FillUpForm(formfieldfloat64, formGroup, probe)
	case "FormFieldInt":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormFieldInt Form",
			OnSave: __gong__New__FormFieldIntFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		formfieldint := new(models.FormFieldInt)
		FillUpForm(formfieldint, formGroup, probe)
	case "FormFieldSelect":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormFieldSelect Form",
			OnSave: __gong__New__FormFieldSelectFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		formfieldselect := new(models.FormFieldSelect)
		FillUpForm(formfieldselect, formGroup, probe)
	case "FormFieldString":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormFieldString Form",
			OnSave: __gong__New__FormFieldStringFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		formfieldstring := new(models.FormFieldString)
		FillUpForm(formfieldstring, formGroup, probe)
	case "FormFieldTime":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormFieldTime Form",
			OnSave: __gong__New__FormFieldTimeFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		formfieldtime := new(models.FormFieldTime)
		FillUpForm(formfieldtime, formGroup, probe)
	case "FormGroup":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormGroup Form",
			OnSave: __gong__New__FormGroupFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		formgroup := new(models.FormGroup)
		FillUpForm(formgroup, formGroup, probe)
	case "FormSortAssocButton":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormSortAssocButton Form",
			OnSave: __gong__New__FormSortAssocButtonFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		formsortassocbutton := new(models.FormSortAssocButton)
		FillUpForm(formsortassocbutton, formGroup, probe)
	case "Option":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Option Form",
			OnSave: __gong__New__OptionFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		option := new(models.Option)
		FillUpForm(option, formGroup, probe)
	case "Row":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Row Form",
			OnSave: __gong__New__RowFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		row := new(models.Row)
		FillUpForm(row, formGroup, probe)
	case "Table":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Table Form",
			OnSave: __gong__New__TableFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		table := new(models.Table)
		FillUpForm(table, formGroup, probe)
	}
	formStage.Commit()
}
