// generated code - do not edit
package data

import (
	"fmt"
	"log"
	"sort"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtable "github.com/fullstack-lang/gongtable/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/maticons/maticons"

	"github.com/fullstack-lang/gongtable/go/models"
	"github.com/fullstack-lang/gongtable/go/orm"
)

type NodeImplGongstruct struct {
	gongStruct         *gong_models.GongStruct
	tableStage         *gongtable.StageStruct
	formStage          *gongtable.StageStruct
	stageOfInterest    *models.StageStruct
	backRepoOfInterest *orm.BackRepoStruct
	r                  *gin.Engine
}

func NewNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	tableStage *gongtable.StageStruct,
	formStage *gongtable.StageStruct,
	stageOfInterest *models.StageStruct,
	backRepoOfInterest *orm.BackRepoStruct,
	r *gin.Engine,
) (nodeImplGongstruct *NodeImplGongstruct) {

	nodeImplGongstruct = new(NodeImplGongstruct)
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.tableStage = tableStage
	nodeImplGongstruct.formStage = formStage
	nodeImplGongstruct.stageOfInterest = stageOfInterest
	nodeImplGongstruct.backRepoOfInterest = backRepoOfInterest
	nodeImplGongstruct.r = r
	return
}

func (nodeImplGongstruct *NodeImplGongstruct) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	// setting the value of the staged node	to the new value
	// otherwise, the expansion memory is lost
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

	}

	// if node is checked, add gongstructshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

	}

	// the node was selected. Therefore, one request the
	// table to route to the table
	log.Println("NodeImplGongstruct:OnAfterUpdate with: ", nodeImplGongstruct.gongStruct.GetName())

	tableStage := nodeImplGongstruct.tableStage

	// insertion point
	if nodeImplGongstruct.gongStruct.GetName() == "Cell" {
		fillUpTable[models.Cell](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CellBoolean" {
		fillUpTable[models.CellBoolean](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CellFloat64" {
		fillUpTable[models.CellFloat64](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CellIcon" {
		fillUpTable[models.CellIcon](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CellInt" {
		fillUpTable[models.CellInt](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CellString" {
		fillUpTable[models.CellString](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CheckBox" {
		fillUpTable[models.CheckBox](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DisplayedColumn" {
		fillUpTable[models.DisplayedColumn](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormDiv" {
		fillUpTable[models.FormDiv](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormEditAssocButton" {
		fillUpTable[models.FormEditAssocButton](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormField" {
		fillUpTable[models.FormField](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldDate" {
		fillUpTable[models.FormFieldDate](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldDateTime" {
		fillUpTable[models.FormFieldDateTime](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldFloat64" {
		fillUpTable[models.FormFieldFloat64](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldInt" {
		fillUpTable[models.FormFieldInt](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldSelect" {
		fillUpTable[models.FormFieldSelect](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldString" {
		fillUpTable[models.FormFieldString](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldTime" {
		fillUpTable[models.FormFieldTime](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormGroup" {
		fillUpTable[models.FormGroup](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormSortAssocButton" {
		fillUpTable[models.FormSortAssocButton](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Option" {
		fillUpTable[models.Option](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Row" {
		fillUpTable[models.Row](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Table" {
		fillUpTable[models.Table](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}

	tableStage.Commit()
}

func fillUpTable[T models.Gongstruct](
	stageOfInterest *models.StageStruct,
	tableStage *gongtable.StageStruct,
	formStage *gongtable.StageStruct,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) {

	tableStage.Reset()
	tableStage.Commit()

	table := new(gongtable.Table).Stage(tableStage)
	table.Name = "Table"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true
	table.HasCheckableRows = false
	table.HasSaveButton = false

	fields := models.GetFields[T]()
	table.NbOfStickyColumns = 3

	// refresh the stage of interest
	stageOfInterest.Checkout()

	setOfStructs := (*models.GetGongstructInstancesSet[T](stageOfInterest))
	sliceOfGongStructsSorted := make([]*T, len(setOfStructs))
	i := 0
	for k := range setOfStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return orm.GetID(
			stageOfInterest,
			backRepoOfInterest,
			sliceOfGongStructsSorted[i],
		) <
			orm.GetID(
				stageOfInterest,
				backRepoOfInterest,
				sliceOfGongStructsSorted[j],
			)
	})

	column := new(gongtable.DisplayedColumn).Stage(tableStage)
	column.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	column = new(gongtable.DisplayedColumn).Stage(tableStage)
	column.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	for _, fieldName := range fields {
		column := new(gongtable.DisplayedColumn).Stage(tableStage)
		column.Name = fieldName
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	fieldIndex := 0
	for _, structInstance := range sliceOfGongStructsSorted {
		row := new(gongtable.Row).Stage(tableStage)
		row.Name = models.GetFieldStringValue[T](*structInstance, "Name")

		updater := NewRowUpdate[T](structInstance,
			tableStage,
			formStage,
			stageOfInterest,
			r,
			backRepoOfInterest)
		updater.Instance = structInstance
		row.Impl = updater

		table.Rows = append(table.Rows, row)

		cell := (&gongtable.Cell{
			Name: "ID",
		}).Stage(tableStage)
		row.Cells = append(row.Cells, cell)
		cellInt := (&gongtable.CellInt{
			Name: "ID",
			Value: orm.GetID(
				stageOfInterest,
				backRepoOfInterest,
				structInstance,
			),
		}).Stage(tableStage)
		cell.CellInt = cellInt

		cell = (&gongtable.Cell{
			Name: "Delete Icon",
		}).Stage(tableStage)
		row.Cells = append(row.Cells, cell)
		cellIcon := (&gongtable.CellIcon{
			Name: "Delete Icon",
			Icon: string(maticons.BUTTON_delete),
		}).Stage(tableStage)
		cell.CellIcon = cellIcon

		for _, fieldName := range fields {
			value := models.GetFieldStringValue[T](*structInstance, fieldName)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value
			fieldIndex++
			// log.Println(fieldName, value)
			cell := (&gongtable.Cell{
				Name: name,
			}).Stage(tableStage)
			row.Cells = append(row.Cells, cell)

			cellString := (&gongtable.CellString{
				Name:  name,
				Value: value,
			}).Stage(tableStage)
			cell.CellString = cellString
		}
	}
}

func NewRowUpdate[T models.Gongstruct](
	Instance *T,
	tableStage *gongtable.StageStruct,
	formStage *gongtable.StageStruct,
	stageOfInterest *models.StageStruct,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (rowUpdate *RowUpdate[T]) {
	rowUpdate = new(RowUpdate[T])
	rowUpdate.Instance = Instance
	rowUpdate.tableStage = tableStage
	rowUpdate.formStage = formStage
	rowUpdate.stageOfInterest = stageOfInterest
	rowUpdate.r = r
	rowUpdate.backRepoOfInterest = backRepoOfInterest
	return
}

type RowUpdate[T models.Gongstruct] struct {
	Instance           *T
	tableStage         *gongtable.StageStruct
	formStage          *gongtable.StageStruct
	stageOfInterest    *models.StageStruct
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (rowUpdate *RowUpdate[T]) RowUpdated(stage *gongtable.StageStruct, row, updatedRow *gongtable.Row) {
	log.Println("RowUpdate: RowUpdated", updatedRow.Name)

	formStage := rowUpdate.formStage
	formStage.Reset()
	formStage.Commit()

	switch instancesTyped := any(rowUpdate.Instance).(type) {
	// insertion point
	case *models.Cell:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewCellFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.CellBoolean:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewCellBooleanFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.CellFloat64:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewCellFloat64FormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.CellIcon:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewCellIconFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.CellInt:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewCellIntFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.CellString:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewCellStringFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.CheckBox:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewCheckBoxFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.DisplayedColumn:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewDisplayedColumnFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.FormDiv:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewFormDivFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.FormEditAssocButton:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewFormEditAssocButtonFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.FormField:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.FormFieldDate:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldDateFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.FormFieldDateTime:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldDateTimeFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.FormFieldFloat64:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldFloat64FormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.FormFieldInt:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldIntFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.FormFieldSelect:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldSelectFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.FormFieldString:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldStringFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.FormFieldTime:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldTimeFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.FormGroup:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewFormGroupFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.FormSortAssocButton:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewFormSortAssocButtonFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.Option:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewOptionFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.Row:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewRowFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.Table:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewTableFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	}
	formStage.Commit()

}
