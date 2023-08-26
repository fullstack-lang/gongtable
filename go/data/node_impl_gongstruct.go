// generated code - do not edit
package data

import (
	"fmt"
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtable_models "github.com/fullstack-lang/gongtable/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/maticons/maticons"

	"github.com/fullstack-lang/gongtable/go/models"
	"github.com/fullstack-lang/gongtable/go/orm"
)

type NodeImplGongstruct struct {
	gongStruct         *gong_models.GongStruct
	gongtableStage     *gongtable_models.StageStruct
	stageOfInterest    *models.StageStruct
	backRepoOfInterest *orm.BackRepoStruct
}

func NewNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	gongtableStage *gongtable_models.StageStruct,
	stageOfInterest *models.StageStruct,
	backRepoOfInterest *orm.BackRepoStruct,
) (nodeImplGongstruct *NodeImplGongstruct) {

	nodeImplGongstruct = new(NodeImplGongstruct)
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.gongtableStage = gongtableStage
	nodeImplGongstruct.stageOfInterest = stageOfInterest
	nodeImplGongstruct.backRepoOfInterest = backRepoOfInterest
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

	tableStage := nodeImplGongstruct.gongtableStage
	tableStage.Reset()
	tableStage.Commit()

	table := new(gongtable_models.Table).Stage(tableStage)
	table.Name = "Table"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true
	table.HasCheckableRows = false
	table.HasSaveButton = false

	// insertion point
	if nodeImplGongstruct.gongStruct.GetName() == "Cell" {
		fillUpTable[models.Cell](nodeImplGongstruct, tableStage, table)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CellBoolean" {
		fillUpTable[models.CellBoolean](nodeImplGongstruct, tableStage, table)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CellFloat64" {
		fillUpTable[models.CellFloat64](nodeImplGongstruct, tableStage, table)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CellIcon" {
		fillUpTable[models.CellIcon](nodeImplGongstruct, tableStage, table)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CellInt" {
		fillUpTable[models.CellInt](nodeImplGongstruct, tableStage, table)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CellString" {
		fillUpTable[models.CellString](nodeImplGongstruct, tableStage, table)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CheckBox" {
		fillUpTable[models.CheckBox](nodeImplGongstruct, tableStage, table)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DisplayedColumn" {
		fillUpTable[models.DisplayedColumn](nodeImplGongstruct, tableStage, table)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormDiv" {
		fillUpTable[models.FormDiv](nodeImplGongstruct, tableStage, table)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormEditAssocButton" {
		fillUpTable[models.FormEditAssocButton](nodeImplGongstruct, tableStage, table)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormField" {
		fillUpTable[models.FormField](nodeImplGongstruct, tableStage, table)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldDate" {
		fillUpTable[models.FormFieldDate](nodeImplGongstruct, tableStage, table)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldDateTime" {
		fillUpTable[models.FormFieldDateTime](nodeImplGongstruct, tableStage, table)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldFloat64" {
		fillUpTable[models.FormFieldFloat64](nodeImplGongstruct, tableStage, table)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldInt" {
		fillUpTable[models.FormFieldInt](nodeImplGongstruct, tableStage, table)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldSelect" {
		fillUpTable[models.FormFieldSelect](nodeImplGongstruct, tableStage, table)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldString" {
		fillUpTable[models.FormFieldString](nodeImplGongstruct, tableStage, table)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldTime" {
		fillUpTable[models.FormFieldTime](nodeImplGongstruct, tableStage, table)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormGroup" {
		fillUpTable[models.FormGroup](nodeImplGongstruct, tableStage, table)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormSortAssocButton" {
		fillUpTable[models.FormSortAssocButton](nodeImplGongstruct, tableStage, table)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Option" {
		fillUpTable[models.Option](nodeImplGongstruct, tableStage, table)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Row" {
		fillUpTable[models.Row](nodeImplGongstruct, tableStage, table)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Table" {
		fillUpTable[models.Table](nodeImplGongstruct, tableStage, table)
	}

	tableStage.Commit()
}

func fillUpTable[T models.Gongstruct](
	nodeImplGongstruct *NodeImplGongstruct,
	tableStage *gongtable_models.StageStruct,
	table *gongtable_models.Table,
) {

	fields := models.GetFields[T]()
	table.NbOfStickyColumns = 3

	setOfStructs := (*models.GetGongstructInstancesSet[T](nodeImplGongstruct.stageOfInterest))

	column := new(gongtable_models.DisplayedColumn).Stage(tableStage)
	column.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	column = new(gongtable_models.DisplayedColumn).Stage(tableStage)
	column.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	for _, fieldName := range fields {
		column := new(gongtable_models.DisplayedColumn).Stage(tableStage)
		column.Name = fieldName
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	fieldIndex := 0
	for structInstance := range setOfStructs {
		row := new(gongtable_models.Row).Stage(tableStage)
		row.Name = models.GetFieldStringValue[T](*structInstance, "Name")
		table.Rows = append(table.Rows, row)

		cell := (&gongtable_models.Cell{
			Name: "ID",
		}).Stage(tableStage)
		row.Cells = append(row.Cells, cell)
		cellInt := (&gongtable_models.CellInt{
			Name: "ID",
			Value: orm.GetID(
				nodeImplGongstruct.stageOfInterest,
				nodeImplGongstruct.backRepoOfInterest,
				structInstance,
			),
		}).Stage(tableStage)
		cell.CellInt = cellInt

		cell = (&gongtable_models.Cell{
			Name: "Delete Icon",
		}).Stage(tableStage)
		row.Cells = append(row.Cells, cell)
		cellIcon := (&gongtable_models.CellIcon{
			Name: "Delete Icon",
			Icon: string(maticons.BUTTON_delete),
		}).Stage(tableStage)
		cell.CellIcon = cellIcon

		for _, fieldName := range fields {
			value := models.GetFieldStringValue[T](*structInstance, fieldName)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value
			fieldIndex++
			// log.Println(fieldName, value)
			cell := (&gongtable_models.Cell{
				Name: name,
			}).Stage(tableStage)
			row.Cells = append(row.Cells, cell)

			cellString := (&gongtable_models.CellString{
				Name:  name,
				Value: value,
			}).Stage(tableStage)
			cell.CellString = cellString

		}
	}
}
