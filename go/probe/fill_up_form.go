// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongtable/go/models"
	"github.com/fullstack-lang/gongtable/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	playground *Playground,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Cell:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("CellString", instanceWithInferedType.CellString, formGroup, playground)
		AssociationFieldToForm("CellFloat64", instanceWithInferedType.CellFloat64, formGroup, playground)
		AssociationFieldToForm("CellInt", instanceWithInferedType.CellInt, formGroup, playground)
		AssociationFieldToForm("CellBool", instanceWithInferedType.CellBool, formGroup, playground)
		AssociationFieldToForm("CellIcon", instanceWithInferedType.CellIcon, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Row"
			rf.Fieldname = "Cells"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Row),
					"Cells",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Row, *models.Cell](
					nil,
					"Cells",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.CellBoolean:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.CellFloat64:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.CellIcon:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Icon", instanceWithInferedType.Icon, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.CellInt:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.CellString:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.CheckBox:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "FormDiv"
			rf.Fieldname = "CheckBoxs"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.FormDiv),
					"CheckBoxs",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.FormDiv, *models.CheckBox](
					nil,
					"CheckBoxs",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.DisplayedColumn:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Table"
			rf.Fieldname = "DisplayedColumns"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Table),
					"DisplayedColumns",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Table, *models.DisplayedColumn](
					nil,
					"DisplayedColumns",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.FormDiv:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("FormFields", instanceWithInferedType, &instanceWithInferedType.FormFields, formGroup, playground)
		AssociationSliceToForm("CheckBoxs", instanceWithInferedType, &instanceWithInferedType.CheckBoxs, formGroup, playground)
		AssociationFieldToForm("FormEditAssocButton", instanceWithInferedType.FormEditAssocButton, formGroup, playground)
		AssociationFieldToForm("FormSortAssocButton", instanceWithInferedType.FormSortAssocButton, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "FormGroup"
			rf.Fieldname = "FormDivs"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.FormGroup),
					"FormDivs",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.FormGroup, *models.FormDiv](
					nil,
					"FormDivs",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.FormEditAssocButton:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.FormField:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		EnumTypeStringToForm("InputTypeEnum", instanceWithInferedType.InputTypeEnum, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Placeholder", instanceWithInferedType.Placeholder, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("FormFieldString", instanceWithInferedType.FormFieldString, formGroup, playground)
		AssociationFieldToForm("FormFieldFloat64", instanceWithInferedType.FormFieldFloat64, formGroup, playground)
		AssociationFieldToForm("FormFieldInt", instanceWithInferedType.FormFieldInt, formGroup, playground)
		AssociationFieldToForm("FormFieldDate", instanceWithInferedType.FormFieldDate, formGroup, playground)
		AssociationFieldToForm("FormFieldTime", instanceWithInferedType.FormFieldTime, formGroup, playground)
		AssociationFieldToForm("FormFieldDateTime", instanceWithInferedType.FormFieldDateTime, formGroup, playground)
		AssociationFieldToForm("FormFieldSelect", instanceWithInferedType.FormFieldSelect, formGroup, playground)
		BasicFieldtoForm("HasBespokeWidth", instanceWithInferedType.HasBespokeWidth, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("BespokeWidthPx", instanceWithInferedType.BespokeWidthPx, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "FormDiv"
			rf.Fieldname = "FormFields"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.FormDiv),
					"FormFields",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.FormDiv, *models.FormField](
					nil,
					"FormFields",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.FormFieldDate:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.FormFieldDateTime:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.FormFieldFloat64:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("HasMinValidator", instanceWithInferedType.HasMinValidator, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("MinValue", instanceWithInferedType.MinValue, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("HasMaxValidator", instanceWithInferedType.HasMaxValidator, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("MaxValue", instanceWithInferedType.MaxValue, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.FormFieldInt:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("HasMinValidator", instanceWithInferedType.HasMinValidator, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("MinValue", instanceWithInferedType.MinValue, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("HasMaxValidator", instanceWithInferedType.HasMaxValidator, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("MaxValue", instanceWithInferedType.MaxValue, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.FormFieldSelect:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("Value", instanceWithInferedType.Value, formGroup, playground)
		AssociationSliceToForm("Options", instanceWithInferedType, &instanceWithInferedType.Options, formGroup, playground)
		BasicFieldtoForm("CanBeEmpty", instanceWithInferedType.CanBeEmpty, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.FormFieldString:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("IsTextArea", instanceWithInferedType.IsTextArea, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.FormFieldTime:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Step", instanceWithInferedType.Step, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.FormGroup:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("FormDivs", instanceWithInferedType, &instanceWithInferedType.FormDivs, formGroup, playground)

	case *models.FormSortAssocButton:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.Option:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "FormFieldSelect"
			rf.Fieldname = "Options"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.FormFieldSelect),
					"Options",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.FormFieldSelect, *models.Option](
					nil,
					"Options",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Row:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Cells", instanceWithInferedType, &instanceWithInferedType.Cells, formGroup, playground)
		BasicFieldtoForm("IsChecked", instanceWithInferedType.IsChecked, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Table"
			rf.Fieldname = "Rows"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Table),
					"Rows",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Table, *models.Row](
					nil,
					"Rows",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Table:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("DisplayedColumns", instanceWithInferedType, &instanceWithInferedType.DisplayedColumns, formGroup, playground)
		AssociationSliceToForm("Rows", instanceWithInferedType, &instanceWithInferedType.Rows, formGroup, playground)
		BasicFieldtoForm("HasFiltering", instanceWithInferedType.HasFiltering, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("HasColumnSorting", instanceWithInferedType.HasColumnSorting, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("HasPaginator", instanceWithInferedType.HasPaginator, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("HasCheckableRows", instanceWithInferedType.HasCheckableRows, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("HasSaveButton", instanceWithInferedType.HasSaveButton, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("CanDragDropRows", instanceWithInferedType.CanDragDropRows, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("HasCloseButton", instanceWithInferedType.HasCloseButton, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("SavingInProgress", instanceWithInferedType.SavingInProgress, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("NbOfStickyColumns", instanceWithInferedType.NbOfStickyColumns, instanceWithInferedType, playground.formStage, formGroup, false)

	default:
		_ = instanceWithInferedType
	}
}
