package probe

import (
	"fmt"
	"sort"
	"strings"

	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	tree "github.com/fullstack-lang/gongtree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gongtable/go/models"
)

func fillUpTree(
	playground *Playground,
) {
	// keep in memory which nodes have been unfolded / folded
	expandedNodesSet := make(map[string]any, 0)
	var _sidebar *tree.Tree
	for __sidebar := range playground.treeStage.Trees {
		_sidebar = __sidebar
	}
	if _sidebar != nil {
		for _, node := range _sidebar.RootNodes {
			if node.IsExpanded {
				expandedNodesSet[strings.Fields(node.Name)[0]] = true
			}
		}
	}

	playground.treeStage.Reset()

	// create tree
	sidebar := (&tree.Tree{Name: "gong"}).Stage(playground.treeStage)

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

		nodeGongstruct := (&tree.Node{Name: name}).Stage(playground.treeStage)


		nodeGongstruct.IsExpanded = false
		if _, ok := expandedNodesSet[strings.Fields(name)[0]]; ok {
			nodeGongstruct.IsExpanded = true
		}
		
		switch gongStruct.Name {
		// insertion point
		case "Cell":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Cell](playground.stageOfInterest)
			for _cell := range set {
				nodeInstance := (&tree.Node{Name: _cell.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_cell, "Cell", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "CellBoolean":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.CellBoolean](playground.stageOfInterest)
			for _cellboolean := range set {
				nodeInstance := (&tree.Node{Name: _cellboolean.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_cellboolean, "CellBoolean", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "CellFloat64":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.CellFloat64](playground.stageOfInterest)
			for _cellfloat64 := range set {
				nodeInstance := (&tree.Node{Name: _cellfloat64.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_cellfloat64, "CellFloat64", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "CellIcon":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.CellIcon](playground.stageOfInterest)
			for _cellicon := range set {
				nodeInstance := (&tree.Node{Name: _cellicon.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_cellicon, "CellIcon", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "CellInt":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.CellInt](playground.stageOfInterest)
			for _cellint := range set {
				nodeInstance := (&tree.Node{Name: _cellint.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_cellint, "CellInt", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "CellString":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.CellString](playground.stageOfInterest)
			for _cellstring := range set {
				nodeInstance := (&tree.Node{Name: _cellstring.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_cellstring, "CellString", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "CheckBox":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.CheckBox](playground.stageOfInterest)
			for _checkbox := range set {
				nodeInstance := (&tree.Node{Name: _checkbox.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_checkbox, "CheckBox", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DisplayedColumn":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DisplayedColumn](playground.stageOfInterest)
			for _displayedcolumn := range set {
				nodeInstance := (&tree.Node{Name: _displayedcolumn.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_displayedcolumn, "DisplayedColumn", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FormDiv":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FormDiv](playground.stageOfInterest)
			for _formdiv := range set {
				nodeInstance := (&tree.Node{Name: _formdiv.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_formdiv, "FormDiv", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FormEditAssocButton":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FormEditAssocButton](playground.stageOfInterest)
			for _formeditassocbutton := range set {
				nodeInstance := (&tree.Node{Name: _formeditassocbutton.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_formeditassocbutton, "FormEditAssocButton", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FormField":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FormField](playground.stageOfInterest)
			for _formfield := range set {
				nodeInstance := (&tree.Node{Name: _formfield.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_formfield, "FormField", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FormFieldDate":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FormFieldDate](playground.stageOfInterest)
			for _formfielddate := range set {
				nodeInstance := (&tree.Node{Name: _formfielddate.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_formfielddate, "FormFieldDate", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FormFieldDateTime":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FormFieldDateTime](playground.stageOfInterest)
			for _formfielddatetime := range set {
				nodeInstance := (&tree.Node{Name: _formfielddatetime.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_formfielddatetime, "FormFieldDateTime", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FormFieldFloat64":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FormFieldFloat64](playground.stageOfInterest)
			for _formfieldfloat64 := range set {
				nodeInstance := (&tree.Node{Name: _formfieldfloat64.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_formfieldfloat64, "FormFieldFloat64", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FormFieldInt":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FormFieldInt](playground.stageOfInterest)
			for _formfieldint := range set {
				nodeInstance := (&tree.Node{Name: _formfieldint.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_formfieldint, "FormFieldInt", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FormFieldSelect":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FormFieldSelect](playground.stageOfInterest)
			for _formfieldselect := range set {
				nodeInstance := (&tree.Node{Name: _formfieldselect.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_formfieldselect, "FormFieldSelect", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FormFieldString":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FormFieldString](playground.stageOfInterest)
			for _formfieldstring := range set {
				nodeInstance := (&tree.Node{Name: _formfieldstring.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_formfieldstring, "FormFieldString", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FormFieldTime":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FormFieldTime](playground.stageOfInterest)
			for _formfieldtime := range set {
				nodeInstance := (&tree.Node{Name: _formfieldtime.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_formfieldtime, "FormFieldTime", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FormGroup":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FormGroup](playground.stageOfInterest)
			for _formgroup := range set {
				nodeInstance := (&tree.Node{Name: _formgroup.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_formgroup, "FormGroup", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FormSortAssocButton":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FormSortAssocButton](playground.stageOfInterest)
			for _formsortassocbutton := range set {
				nodeInstance := (&tree.Node{Name: _formsortassocbutton.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_formsortassocbutton, "FormSortAssocButton", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Option":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Option](playground.stageOfInterest)
			for _option := range set {
				nodeInstance := (&tree.Node{Name: _option.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_option, "Option", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Row":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Row](playground.stageOfInterest)
			for _row := range set {
				nodeInstance := (&tree.Node{Name: _row.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_row, "Row", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Table":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Table](playground.stageOfInterest)
			for _table := range set {
				nodeInstance := (&tree.Node{Name: _table.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_table, "Table", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}	
		}

		nodeGongstruct.IsNodeClickable = true
		nodeGongstruct.Impl = NewTreeNodeImplGongstruct(gongStruct, playground)

		// add add button
		addButton := (&tree.Button{
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
	gongtreeStage *tree.StageStruct,
	stagedNode, frontNode *tree.Node) {

	FillUpFormFromGongstruct(
		instanceNodeCallback.Instance,
		instanceNodeCallback.playground,
	)
}
