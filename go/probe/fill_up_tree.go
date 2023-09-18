package probe

import (
	"fmt"
	"sort"

	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gongtable/go/models"
)

func fillUpTree(
	playground *Playground,
) {
	playground.treeStage.Reset()

	// create tree
	sidebar := (&gongtree_models.Tree{Name: "gong"}).Stage(playground.treeStage)

	// collect all gong struct to construe the true
	setOfGongStructs := *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](playground.gongStage)

	sliceOfGongStructsSorted := make([]*gong_models.GongStruct, len(setOfGongStructs))
	i := 0
	for k := range setOfGongStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return sliceOfGongStructsSorted[i].Name < sliceOfGongStructsSorted[j].Name
	})

	for _, gongStruct := range sliceOfGongStructsSorted {

		name := gongStruct.Name + " (" +
			fmt.Sprintf("%d", playground.stageOfInterest.Map_GongStructName_InstancesNb[gongStruct.Name]) + ")"

		nodeGongstruct := (&gongtree_models.Node{Name: name}).Stage(playground.treeStage)


		nodeGongstruct.IsExpanded = false
		switch gongStruct.Name {
		// insertion point
		case "Cell":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Cell](playground.stageOfInterest)
			for cell := range set {
				nodeInstance := (&gongtree_models.Node{Name: cell.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(cell, "Cell", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "CellBoolean":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.CellBoolean](playground.stageOfInterest)
			for cellboolean := range set {
				nodeInstance := (&gongtree_models.Node{Name: cellboolean.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(cellboolean, "CellBoolean", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "CellFloat64":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.CellFloat64](playground.stageOfInterest)
			for cellfloat64 := range set {
				nodeInstance := (&gongtree_models.Node{Name: cellfloat64.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(cellfloat64, "CellFloat64", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "CellIcon":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.CellIcon](playground.stageOfInterest)
			for cellicon := range set {
				nodeInstance := (&gongtree_models.Node{Name: cellicon.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(cellicon, "CellIcon", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "CellInt":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.CellInt](playground.stageOfInterest)
			for cellint := range set {
				nodeInstance := (&gongtree_models.Node{Name: cellint.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(cellint, "CellInt", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "CellString":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.CellString](playground.stageOfInterest)
			for cellstring := range set {
				nodeInstance := (&gongtree_models.Node{Name: cellstring.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(cellstring, "CellString", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "CheckBox":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.CheckBox](playground.stageOfInterest)
			for checkbox := range set {
				nodeInstance := (&gongtree_models.Node{Name: checkbox.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(checkbox, "CheckBox", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DisplayedColumn":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DisplayedColumn](playground.stageOfInterest)
			for displayedcolumn := range set {
				nodeInstance := (&gongtree_models.Node{Name: displayedcolumn.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(displayedcolumn, "DisplayedColumn", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FormDiv":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FormDiv](playground.stageOfInterest)
			for formdiv := range set {
				nodeInstance := (&gongtree_models.Node{Name: formdiv.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(formdiv, "FormDiv", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FormEditAssocButton":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FormEditAssocButton](playground.stageOfInterest)
			for formeditassocbutton := range set {
				nodeInstance := (&gongtree_models.Node{Name: formeditassocbutton.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(formeditassocbutton, "FormEditAssocButton", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FormField":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FormField](playground.stageOfInterest)
			for formfield := range set {
				nodeInstance := (&gongtree_models.Node{Name: formfield.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(formfield, "FormField", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FormFieldDate":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FormFieldDate](playground.stageOfInterest)
			for formfielddate := range set {
				nodeInstance := (&gongtree_models.Node{Name: formfielddate.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(formfielddate, "FormFieldDate", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FormFieldDateTime":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FormFieldDateTime](playground.stageOfInterest)
			for formfielddatetime := range set {
				nodeInstance := (&gongtree_models.Node{Name: formfielddatetime.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(formfielddatetime, "FormFieldDateTime", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FormFieldFloat64":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FormFieldFloat64](playground.stageOfInterest)
			for formfieldfloat64 := range set {
				nodeInstance := (&gongtree_models.Node{Name: formfieldfloat64.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(formfieldfloat64, "FormFieldFloat64", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FormFieldInt":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FormFieldInt](playground.stageOfInterest)
			for formfieldint := range set {
				nodeInstance := (&gongtree_models.Node{Name: formfieldint.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(formfieldint, "FormFieldInt", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FormFieldSelect":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FormFieldSelect](playground.stageOfInterest)
			for formfieldselect := range set {
				nodeInstance := (&gongtree_models.Node{Name: formfieldselect.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(formfieldselect, "FormFieldSelect", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FormFieldString":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FormFieldString](playground.stageOfInterest)
			for formfieldstring := range set {
				nodeInstance := (&gongtree_models.Node{Name: formfieldstring.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(formfieldstring, "FormFieldString", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FormFieldTime":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FormFieldTime](playground.stageOfInterest)
			for formfieldtime := range set {
				nodeInstance := (&gongtree_models.Node{Name: formfieldtime.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(formfieldtime, "FormFieldTime", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FormGroup":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FormGroup](playground.stageOfInterest)
			for formgroup := range set {
				nodeInstance := (&gongtree_models.Node{Name: formgroup.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(formgroup, "FormGroup", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FormSortAssocButton":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FormSortAssocButton](playground.stageOfInterest)
			for formsortassocbutton := range set {
				nodeInstance := (&gongtree_models.Node{Name: formsortassocbutton.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(formsortassocbutton, "FormSortAssocButton", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Option":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Option](playground.stageOfInterest)
			for option := range set {
				nodeInstance := (&gongtree_models.Node{Name: option.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(option, "Option", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Row":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Row](playground.stageOfInterest)
			for row := range set {
				nodeInstance := (&gongtree_models.Node{Name: row.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(row, "Row", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Table":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Table](playground.stageOfInterest)
			for table := range set {
				nodeInstance := (&gongtree_models.Node{Name: table.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(table, "Table", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}	
		}

		nodeGongstruct.IsNodeClickable = true
		nodeGongstruct.Impl = NewTreeNodeImplGongstruct(gongStruct, playground)

		// add add button
		addButton := (&gongtree_models.Button{
			Name: gongStruct.Name + " " + string(gongtree_buttons.BUTTON_add),
			Icon: string(gongtree_buttons.BUTTON_add)}).Stage(playground.treeStage)
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)
		addButton.Impl = NewButtonImplGongstruct(
			gongStruct,
			gongtree_buttons.BUTTON_add,
			playground,
		)

		sidebar.RootNodes = append(sidebar.RootNodes, nodeGongstruct)
	}
	playground.treeStage.Commit()
}

type InstanceNodeCallback[T models.Gongstruct] struct {
	Instance       *T
	gongstructName string
	playground     *Playground
}

func NewInstanceNodeCallback[T models.Gongstruct](
	instance *T,
	gongstructName string,
	playground *Playground) (
	instanceNodeCallback *InstanceNodeCallback[T],
) {
	instanceNodeCallback = new(InstanceNodeCallback[T])

	instanceNodeCallback.playground = playground
	instanceNodeCallback.gongstructName = gongstructName
	instanceNodeCallback.Instance = instance

	return
}

func (instanceNodeCallback *InstanceNodeCallback[T]) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	FillUpFormFromGongstruct(
		instanceNodeCallback.Instance,
		instanceNodeCallback.playground,
	)
}
