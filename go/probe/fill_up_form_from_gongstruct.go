// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongtable/go/models"
)

func FillUpFormFromGongstruct[T models.Gongstruct](instance *T, playground *Playground) {
	formStage := playground.formStage
	formStage.Reset()
	formStage.Commit()

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Cell:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Cell Form",
			OnSave: __gong__New__CellFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.CellBoolean:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update CellBoolean Form",
			OnSave: __gong__New__CellBooleanFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.CellFloat64:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update CellFloat64 Form",
			OnSave: __gong__New__CellFloat64FormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.CellIcon:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update CellIcon Form",
			OnSave: __gong__New__CellIconFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.CellInt:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update CellInt Form",
			OnSave: __gong__New__CellIntFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.CellString:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update CellString Form",
			OnSave: __gong__New__CellStringFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.CheckBox:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update CheckBox Form",
			OnSave: __gong__New__CheckBoxFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.DisplayedColumn:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update DisplayedColumn Form",
			OnSave: __gong__New__DisplayedColumnFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.FormDiv:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update FormDiv Form",
			OnSave: __gong__New__FormDivFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.FormEditAssocButton:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update FormEditAssocButton Form",
			OnSave: __gong__New__FormEditAssocButtonFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.FormField:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update FormField Form",
			OnSave: __gong__New__FormFieldFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.FormFieldDate:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update FormFieldDate Form",
			OnSave: __gong__New__FormFieldDateFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.FormFieldDateTime:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update FormFieldDateTime Form",
			OnSave: __gong__New__FormFieldDateTimeFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.FormFieldFloat64:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update FormFieldFloat64 Form",
			OnSave: __gong__New__FormFieldFloat64FormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.FormFieldInt:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update FormFieldInt Form",
			OnSave: __gong__New__FormFieldIntFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.FormFieldSelect:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update FormFieldSelect Form",
			OnSave: __gong__New__FormFieldSelectFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.FormFieldString:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update FormFieldString Form",
			OnSave: __gong__New__FormFieldStringFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.FormFieldTime:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update FormFieldTime Form",
			OnSave: __gong__New__FormFieldTimeFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.FormGroup:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update FormGroup Form",
			OnSave: __gong__New__FormGroupFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.FormSortAssocButton:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update FormSortAssocButton Form",
			OnSave: __gong__New__FormSortAssocButtonFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.Option:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Option Form",
			OnSave: __gong__New__OptionFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.Row:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Row Form",
			OnSave: __gong__New__RowFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.Table:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Table Form",
			OnSave: __gong__New__TableFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
