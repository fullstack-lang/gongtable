import { Component, Input, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

import { BehaviorSubject, Subscription } from 'rxjs';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { CommitNbFromBackService } from '../commitnbfromback.service'
import { GongstructSelectionService } from '../gongstruct-selection.service'

// insertion point for per struct import code
import { CellService } from '../cell.service'
import { getCellUniqueID } from '../front-repo.service'
import { CellBooleanService } from '../cellboolean.service'
import { getCellBooleanUniqueID } from '../front-repo.service'
import { CellFloat64Service } from '../cellfloat64.service'
import { getCellFloat64UniqueID } from '../front-repo.service'
import { CellIconService } from '../cellicon.service'
import { getCellIconUniqueID } from '../front-repo.service'
import { CellIntService } from '../cellint.service'
import { getCellIntUniqueID } from '../front-repo.service'
import { CellStringService } from '../cellstring.service'
import { getCellStringUniqueID } from '../front-repo.service'
import { DisplayedColumnService } from '../displayedcolumn.service'
import { getDisplayedColumnUniqueID } from '../front-repo.service'
import { FormDivService } from '../formdiv.service'
import { getFormDivUniqueID } from '../front-repo.service'
import { FormFieldService } from '../formfield.service'
import { getFormFieldUniqueID } from '../front-repo.service'
import { FormFieldBooleanService } from '../formfieldboolean.service'
import { getFormFieldBooleanUniqueID } from '../front-repo.service'
import { FormFieldDateService } from '../formfielddate.service'
import { getFormFieldDateUniqueID } from '../front-repo.service'
import { FormFieldFloat64Service } from '../formfieldfloat64.service'
import { getFormFieldFloat64UniqueID } from '../front-repo.service'
import { FormFieldIntService } from '../formfieldint.service'
import { getFormFieldIntUniqueID } from '../front-repo.service'
import { FormFieldStringService } from '../formfieldstring.service'
import { getFormFieldStringUniqueID } from '../front-repo.service'
import { FormFieldTimeService } from '../formfieldtime.service'
import { getFormFieldTimeUniqueID } from '../front-repo.service'
import { FormGroupService } from '../formgroup.service'
import { getFormGroupUniqueID } from '../front-repo.service'
import { RowService } from '../row.service'
import { getRowUniqueID } from '../front-repo.service'
import { TableService } from '../table.service'
import { getTableUniqueID } from '../front-repo.service'

import { RouteService } from '../route-service';

/**
 * Types of a GongNode / GongFlatNode
 */
export enum GongNodeType {
  STRUCT = "STRUCT",
  INSTANCE = "INSTANCE",
  ONE__ZERO_ONE_ASSOCIATION = 'ONE__ZERO_ONE_ASSOCIATION',
  ONE__ZERO_MANY_ASSOCIATION = 'ONE__ZERO_MANY_ASSOCIATION',
}

/**
 * GongNode is the "data" node
 */
interface GongNode {
  name: string; // if STRUCT, the name of the struct, if INSTANCE the name of the instance
  children: GongNode[];
  type: GongNodeType;
  structName: string;
  associationField: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


/** 
 * GongFlatNode is the dynamic visual node with expandable and level information
 * */
interface GongFlatNode {
  expandable: boolean;
  name: string;
  level: number;
  type: GongNodeType;
  structName: string;
  associationField: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


@Component({
  selector: 'app-gongtable-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css'],
})
export class SidebarComponent implements OnInit {

  /**
  * _transformer generated a displayed node from a data node
  *
  * @param node input data noe
  * @param level input level
  *
  * @returns an ExampleFlatNode
  */
  private _transformer = (node: GongNode, level: number) => {
    return {

      /**
      * in javascript, The !! ensures the resulting type is a boolean (true or false).
      *
      * !!node.children will evaluate to true is the variable is defined
      */
      expandable: !!node.children && node.children.length > 0,
      name: node.name,
      level: level,
      type: node.type,
      structName: node.structName,
      associationField: node.associationField,
      associatedStructName: node.associatedStructName,
      id: node.id,
      uniqueIdPerStack: node.uniqueIdPerStack,
    }
  }

  /**
   * treeControl is passed as the paramter treeControl in the "mat-tree" selector
   *
   * Flat tree control. Able to expand/collapse a subtree recursively for flattened tree.
   *
   * Construct with flat tree data node functions getLevel and isExpandable.
  constructor(
    getLevel: (dataNode: T) => number,
    isExpandable: (dataNode: T) => boolean, 
    options?: FlatTreeControlOptions<T, K> | undefined);
   */
  treeControl = new FlatTreeControl<GongFlatNode>(
    node => node.level,
    node => node.expandable
  );

  /**
   * from mat-tree documentation
   *
   * Tree flattener to convert a normal type of node to node with children & level information.
   */
  treeFlattener = new MatTreeFlattener(
    this._transformer,
    node => node.level,
    node => node.expandable,
    node => node.children
  );

  /**
   * data is the other paramter to the "mat-tree" selector
   * 
   * strangely, the dataSource declaration has to follow the treeFlattener declaration
   */
  dataSource = new MatTreeFlatDataSource(this.treeControl, this.treeFlattener);

  /**
   * hasChild is used by the selector for expandable nodes
   * 
   *  <mat-tree-node *matTreeNodeDef="let node;when: hasChild" matTreeNodePadding>
   * 
   * @param _ 
   * @param node 
   */
  hasChild = (_: number, node: GongFlatNode) => node.expandable;

  // front repo
  frontRepo: FrontRepo = new (FrontRepo)
  commitNbFromBack: number = 0

  // "data" tree that is constructed during NgInit and is passed to the mat-tree component
  gongNodeTree = new Array<GongNode>();

  // SelectedStructChanged is the behavior subject that will emit
  // the selected gong struct whose table has to be displayed in the table outlet
  SelectedStructChanged: BehaviorSubject<string> = new BehaviorSubject("");

  subscription: Subscription = new Subscription

  @Input() GONG__StackPath: string = ""

  constructor(
    private router: Router,
    private frontRepoService: FrontRepoService,
    private gongstructSelectionService: GongstructSelectionService,

    // insertion point for per struct service declaration
    private cellService: CellService,
    private cellbooleanService: CellBooleanService,
    private cellfloat64Service: CellFloat64Service,
    private celliconService: CellIconService,
    private cellintService: CellIntService,
    private cellstringService: CellStringService,
    private displayedcolumnService: DisplayedColumnService,
    private formdivService: FormDivService,
    private formfieldService: FormFieldService,
    private formfieldbooleanService: FormFieldBooleanService,
    private formfielddateService: FormFieldDateService,
    private formfieldfloat64Service: FormFieldFloat64Service,
    private formfieldintService: FormFieldIntService,
    private formfieldstringService: FormFieldStringService,
    private formfieldtimeService: FormFieldTimeService,
    private formgroupService: FormGroupService,
    private rowService: RowService,
    private tableService: TableService,

    private routeService: RouteService,
  ) { }

  ngOnDestroy() {
    // prevent memory leak when component destroyed
    this.subscription.unsubscribe();
  }

  ngOnInit(): void {

    console.log("Sidebar init: " + this.GONG__StackPath)

    // add the routes that will used by this side panel component and
    // by the component that are called from this component
    this.routeService.addDataPanelRoutes(this.GONG__StackPath)

    this.subscription = this.gongstructSelectionService.gongtructSelected$.subscribe(
      gongstructName => {
        // console.log("sidebar gongstruct selected " + gongstructName)

        this.setTableRouterOutlet(gongstructName.toLowerCase() + "s")
      });

    this.refresh()

    this.SelectedStructChanged.subscribe(
      selectedStruct => {
        if (selectedStruct != "") {
          this.setTableRouterOutlet(selectedStruct.toLowerCase() + "s")
        }
      }
    )

    // insertion point for per struct observable for refresh trigger
    // observable for changes in structs
    this.cellService.CellServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.cellbooleanService.CellBooleanServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.cellfloat64Service.CellFloat64ServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.celliconService.CellIconServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.cellintService.CellIntServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.cellstringService.CellStringServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.displayedcolumnService.DisplayedColumnServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.formdivService.FormDivServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.formfieldService.FormFieldServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.formfieldbooleanService.FormFieldBooleanServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.formfielddateService.FormFieldDateServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.formfieldfloat64Service.FormFieldFloat64ServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.formfieldintService.FormFieldIntServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.formfieldstringService.FormFieldStringServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.formfieldtimeService.FormFieldTimeServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.formgroupService.FormGroupServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.rowService.RowServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.tableService.TableServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
  }

  refresh(): void {
    this.frontRepoService.pull(this.GONG__StackPath).subscribe(frontRepo => {
      this.frontRepo = frontRepo

      // use of a Gödel number to uniquely identfy nodes : 2 * node.id + 3 * node.level
      let memoryOfExpandedNodes = new Map<number, boolean>()
      let nonInstanceNodeId = 1

      this.treeControl.dataNodes?.forEach(
        node => {
          if (this.treeControl.isExpanded(node)) {
            memoryOfExpandedNodes.set(node.uniqueIdPerStack, true)
          } else {
            memoryOfExpandedNodes.set(node.uniqueIdPerStack, false)
          }
        }
      )

      // reset the gong node tree
      this.gongNodeTree = new Array<GongNode>();

      // insertion point for per struct tree construction
      /**
      * fill up the Cell part of the mat tree
      */
      let cellGongNodeStruct: GongNode = {
        name: "Cell",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Cell",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(cellGongNodeStruct)

      this.frontRepo.Cells_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Cells_array.forEach(
        cellDB => {
          let cellGongNodeInstance: GongNode = {
            name: cellDB.Name,
            type: GongNodeType.INSTANCE,
            id: cellDB.ID,
            uniqueIdPerStack: getCellUniqueID(cellDB.ID),
            structName: "Cell",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          cellGongNodeStruct.children!.push(cellGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association CellString
          */
          let CellStringGongNodeAssociation: GongNode = {
            name: "(CellString) CellString",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: cellDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Cell",
            associationField: "CellString",
            associatedStructName: "CellString",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          cellGongNodeInstance.children!.push(CellStringGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation CellString
            */
          if (cellDB.CellString != undefined) {
            let cellGongNodeInstance_CellString: GongNode = {
              name: cellDB.CellString.Name,
              type: GongNodeType.INSTANCE,
              id: cellDB.CellString.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getCellUniqueID(cellDB.ID)
                + 5 * getCellStringUniqueID(cellDB.CellString.ID),
              structName: "CellString",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            CellStringGongNodeAssociation.children.push(cellGongNodeInstance_CellString)
          }

          /**
          * let append a node for the association CellFloat64
          */
          let CellFloat64GongNodeAssociation: GongNode = {
            name: "(CellFloat64) CellFloat64",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: cellDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Cell",
            associationField: "CellFloat64",
            associatedStructName: "CellFloat64",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          cellGongNodeInstance.children!.push(CellFloat64GongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation CellFloat64
            */
          if (cellDB.CellFloat64 != undefined) {
            let cellGongNodeInstance_CellFloat64: GongNode = {
              name: cellDB.CellFloat64.Name,
              type: GongNodeType.INSTANCE,
              id: cellDB.CellFloat64.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getCellUniqueID(cellDB.ID)
                + 5 * getCellFloat64UniqueID(cellDB.CellFloat64.ID),
              structName: "CellFloat64",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            CellFloat64GongNodeAssociation.children.push(cellGongNodeInstance_CellFloat64)
          }

          /**
          * let append a node for the association CellInt
          */
          let CellIntGongNodeAssociation: GongNode = {
            name: "(CellInt) CellInt",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: cellDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Cell",
            associationField: "CellInt",
            associatedStructName: "CellInt",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          cellGongNodeInstance.children!.push(CellIntGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation CellInt
            */
          if (cellDB.CellInt != undefined) {
            let cellGongNodeInstance_CellInt: GongNode = {
              name: cellDB.CellInt.Name,
              type: GongNodeType.INSTANCE,
              id: cellDB.CellInt.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getCellUniqueID(cellDB.ID)
                + 5 * getCellIntUniqueID(cellDB.CellInt.ID),
              structName: "CellInt",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            CellIntGongNodeAssociation.children.push(cellGongNodeInstance_CellInt)
          }

          /**
          * let append a node for the association CellBool
          */
          let CellBoolGongNodeAssociation: GongNode = {
            name: "(CellBoolean) CellBool",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: cellDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Cell",
            associationField: "CellBool",
            associatedStructName: "CellBoolean",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          cellGongNodeInstance.children!.push(CellBoolGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation CellBool
            */
          if (cellDB.CellBool != undefined) {
            let cellGongNodeInstance_CellBool: GongNode = {
              name: cellDB.CellBool.Name,
              type: GongNodeType.INSTANCE,
              id: cellDB.CellBool.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getCellUniqueID(cellDB.ID)
                + 5 * getCellBooleanUniqueID(cellDB.CellBool.ID),
              structName: "CellBoolean",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            CellBoolGongNodeAssociation.children.push(cellGongNodeInstance_CellBool)
          }

          /**
          * let append a node for the association CellIcon
          */
          let CellIconGongNodeAssociation: GongNode = {
            name: "(CellIcon) CellIcon",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: cellDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Cell",
            associationField: "CellIcon",
            associatedStructName: "CellIcon",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          cellGongNodeInstance.children!.push(CellIconGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation CellIcon
            */
          if (cellDB.CellIcon != undefined) {
            let cellGongNodeInstance_CellIcon: GongNode = {
              name: cellDB.CellIcon.Name,
              type: GongNodeType.INSTANCE,
              id: cellDB.CellIcon.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getCellUniqueID(cellDB.ID)
                + 5 * getCellIconUniqueID(cellDB.CellIcon.ID),
              structName: "CellIcon",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            CellIconGongNodeAssociation.children.push(cellGongNodeInstance_CellIcon)
          }

        }
      )

      /**
      * fill up the CellBoolean part of the mat tree
      */
      let cellbooleanGongNodeStruct: GongNode = {
        name: "CellBoolean",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "CellBoolean",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(cellbooleanGongNodeStruct)

      this.frontRepo.CellBooleans_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.CellBooleans_array.forEach(
        cellbooleanDB => {
          let cellbooleanGongNodeInstance: GongNode = {
            name: cellbooleanDB.Name,
            type: GongNodeType.INSTANCE,
            id: cellbooleanDB.ID,
            uniqueIdPerStack: getCellBooleanUniqueID(cellbooleanDB.ID),
            structName: "CellBoolean",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          cellbooleanGongNodeStruct.children!.push(cellbooleanGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the CellFloat64 part of the mat tree
      */
      let cellfloat64GongNodeStruct: GongNode = {
        name: "CellFloat64",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "CellFloat64",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(cellfloat64GongNodeStruct)

      this.frontRepo.CellFloat64s_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.CellFloat64s_array.forEach(
        cellfloat64DB => {
          let cellfloat64GongNodeInstance: GongNode = {
            name: cellfloat64DB.Name,
            type: GongNodeType.INSTANCE,
            id: cellfloat64DB.ID,
            uniqueIdPerStack: getCellFloat64UniqueID(cellfloat64DB.ID),
            structName: "CellFloat64",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          cellfloat64GongNodeStruct.children!.push(cellfloat64GongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the CellIcon part of the mat tree
      */
      let celliconGongNodeStruct: GongNode = {
        name: "CellIcon",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "CellIcon",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(celliconGongNodeStruct)

      this.frontRepo.CellIcons_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.CellIcons_array.forEach(
        celliconDB => {
          let celliconGongNodeInstance: GongNode = {
            name: celliconDB.Name,
            type: GongNodeType.INSTANCE,
            id: celliconDB.ID,
            uniqueIdPerStack: getCellIconUniqueID(celliconDB.ID),
            structName: "CellIcon",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          celliconGongNodeStruct.children!.push(celliconGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the CellInt part of the mat tree
      */
      let cellintGongNodeStruct: GongNode = {
        name: "CellInt",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "CellInt",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(cellintGongNodeStruct)

      this.frontRepo.CellInts_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.CellInts_array.forEach(
        cellintDB => {
          let cellintGongNodeInstance: GongNode = {
            name: cellintDB.Name,
            type: GongNodeType.INSTANCE,
            id: cellintDB.ID,
            uniqueIdPerStack: getCellIntUniqueID(cellintDB.ID),
            structName: "CellInt",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          cellintGongNodeStruct.children!.push(cellintGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the CellString part of the mat tree
      */
      let cellstringGongNodeStruct: GongNode = {
        name: "CellString",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "CellString",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(cellstringGongNodeStruct)

      this.frontRepo.CellStrings_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.CellStrings_array.forEach(
        cellstringDB => {
          let cellstringGongNodeInstance: GongNode = {
            name: cellstringDB.Name,
            type: GongNodeType.INSTANCE,
            id: cellstringDB.ID,
            uniqueIdPerStack: getCellStringUniqueID(cellstringDB.ID),
            structName: "CellString",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          cellstringGongNodeStruct.children!.push(cellstringGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the DisplayedColumn part of the mat tree
      */
      let displayedcolumnGongNodeStruct: GongNode = {
        name: "DisplayedColumn",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "DisplayedColumn",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(displayedcolumnGongNodeStruct)

      this.frontRepo.DisplayedColumns_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.DisplayedColumns_array.forEach(
        displayedcolumnDB => {
          let displayedcolumnGongNodeInstance: GongNode = {
            name: displayedcolumnDB.Name,
            type: GongNodeType.INSTANCE,
            id: displayedcolumnDB.ID,
            uniqueIdPerStack: getDisplayedColumnUniqueID(displayedcolumnDB.ID),
            structName: "DisplayedColumn",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          displayedcolumnGongNodeStruct.children!.push(displayedcolumnGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the FormDiv part of the mat tree
      */
      let formdivGongNodeStruct: GongNode = {
        name: "FormDiv",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "FormDiv",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(formdivGongNodeStruct)

      this.frontRepo.FormDivs_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.FormDivs_array.forEach(
        formdivDB => {
          let formdivGongNodeInstance: GongNode = {
            name: formdivDB.Name,
            type: GongNodeType.INSTANCE,
            id: formdivDB.ID,
            uniqueIdPerStack: getFormDivUniqueID(formdivDB.ID),
            structName: "FormDiv",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          formdivGongNodeStruct.children!.push(formdivGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer FormFields
          */
          let FormFieldsGongNodeAssociation: GongNode = {
            name: "(FormField) FormFields",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: formdivDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "FormDiv",
            associationField: "FormFields",
            associatedStructName: "FormField",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          formdivGongNodeInstance.children.push(FormFieldsGongNodeAssociation)

          formdivDB.FormFields?.forEach(formfieldDB => {
            let formfieldNode: GongNode = {
              name: formfieldDB.Name,
              type: GongNodeType.INSTANCE,
              id: formfieldDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getFormDivUniqueID(formdivDB.ID)
                + 11 * getFormFieldUniqueID(formfieldDB.ID),
              structName: "FormField",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            FormFieldsGongNodeAssociation.children.push(formfieldNode)
          })

        }
      )

      /**
      * fill up the FormField part of the mat tree
      */
      let formfieldGongNodeStruct: GongNode = {
        name: "FormField",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "FormField",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(formfieldGongNodeStruct)

      this.frontRepo.FormFields_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.FormFields_array.forEach(
        formfieldDB => {
          let formfieldGongNodeInstance: GongNode = {
            name: formfieldDB.Name,
            type: GongNodeType.INSTANCE,
            id: formfieldDB.ID,
            uniqueIdPerStack: getFormFieldUniqueID(formfieldDB.ID),
            structName: "FormField",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          formfieldGongNodeStruct.children!.push(formfieldGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association FormFieldString
          */
          let FormFieldStringGongNodeAssociation: GongNode = {
            name: "(FormFieldString) FormFieldString",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: formfieldDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "FormField",
            associationField: "FormFieldString",
            associatedStructName: "FormFieldString",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          formfieldGongNodeInstance.children!.push(FormFieldStringGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation FormFieldString
            */
          if (formfieldDB.FormFieldString != undefined) {
            let formfieldGongNodeInstance_FormFieldString: GongNode = {
              name: formfieldDB.FormFieldString.Name,
              type: GongNodeType.INSTANCE,
              id: formfieldDB.FormFieldString.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getFormFieldUniqueID(formfieldDB.ID)
                + 5 * getFormFieldStringUniqueID(formfieldDB.FormFieldString.ID),
              structName: "FormFieldString",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            FormFieldStringGongNodeAssociation.children.push(formfieldGongNodeInstance_FormFieldString)
          }

          /**
          * let append a node for the association FormFieldFloat64
          */
          let FormFieldFloat64GongNodeAssociation: GongNode = {
            name: "(FormFieldFloat64) FormFieldFloat64",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: formfieldDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "FormField",
            associationField: "FormFieldFloat64",
            associatedStructName: "FormFieldFloat64",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          formfieldGongNodeInstance.children!.push(FormFieldFloat64GongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation FormFieldFloat64
            */
          if (formfieldDB.FormFieldFloat64 != undefined) {
            let formfieldGongNodeInstance_FormFieldFloat64: GongNode = {
              name: formfieldDB.FormFieldFloat64.Name,
              type: GongNodeType.INSTANCE,
              id: formfieldDB.FormFieldFloat64.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getFormFieldUniqueID(formfieldDB.ID)
                + 5 * getFormFieldFloat64UniqueID(formfieldDB.FormFieldFloat64.ID),
              structName: "FormFieldFloat64",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            FormFieldFloat64GongNodeAssociation.children.push(formfieldGongNodeInstance_FormFieldFloat64)
          }

          /**
          * let append a node for the association FormFieldInt
          */
          let FormFieldIntGongNodeAssociation: GongNode = {
            name: "(FormFieldInt) FormFieldInt",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: formfieldDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "FormField",
            associationField: "FormFieldInt",
            associatedStructName: "FormFieldInt",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          formfieldGongNodeInstance.children!.push(FormFieldIntGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation FormFieldInt
            */
          if (formfieldDB.FormFieldInt != undefined) {
            let formfieldGongNodeInstance_FormFieldInt: GongNode = {
              name: formfieldDB.FormFieldInt.Name,
              type: GongNodeType.INSTANCE,
              id: formfieldDB.FormFieldInt.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getFormFieldUniqueID(formfieldDB.ID)
                + 5 * getFormFieldIntUniqueID(formfieldDB.FormFieldInt.ID),
              structName: "FormFieldInt",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            FormFieldIntGongNodeAssociation.children.push(formfieldGongNodeInstance_FormFieldInt)
          }

          /**
          * let append a node for the association FormFieldBool
          */
          let FormFieldBoolGongNodeAssociation: GongNode = {
            name: "(FormFieldBoolean) FormFieldBool",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: formfieldDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "FormField",
            associationField: "FormFieldBool",
            associatedStructName: "FormFieldBoolean",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          formfieldGongNodeInstance.children!.push(FormFieldBoolGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation FormFieldBool
            */
          if (formfieldDB.FormFieldBool != undefined) {
            let formfieldGongNodeInstance_FormFieldBool: GongNode = {
              name: formfieldDB.FormFieldBool.Name,
              type: GongNodeType.INSTANCE,
              id: formfieldDB.FormFieldBool.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getFormFieldUniqueID(formfieldDB.ID)
                + 5 * getFormFieldBooleanUniqueID(formfieldDB.FormFieldBool.ID),
              structName: "FormFieldBoolean",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            FormFieldBoolGongNodeAssociation.children.push(formfieldGongNodeInstance_FormFieldBool)
          }

          /**
          * let append a node for the association FormFieldDate
          */
          let FormFieldDateGongNodeAssociation: GongNode = {
            name: "(FormFieldDate) FormFieldDate",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: formfieldDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "FormField",
            associationField: "FormFieldDate",
            associatedStructName: "FormFieldDate",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          formfieldGongNodeInstance.children!.push(FormFieldDateGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation FormFieldDate
            */
          if (formfieldDB.FormFieldDate != undefined) {
            let formfieldGongNodeInstance_FormFieldDate: GongNode = {
              name: formfieldDB.FormFieldDate.Name,
              type: GongNodeType.INSTANCE,
              id: formfieldDB.FormFieldDate.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getFormFieldUniqueID(formfieldDB.ID)
                + 5 * getFormFieldDateUniqueID(formfieldDB.FormFieldDate.ID),
              structName: "FormFieldDate",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            FormFieldDateGongNodeAssociation.children.push(formfieldGongNodeInstance_FormFieldDate)
          }

          /**
          * let append a node for the association FormFieldTime
          */
          let FormFieldTimeGongNodeAssociation: GongNode = {
            name: "(FormFieldTime) FormFieldTime",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: formfieldDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "FormField",
            associationField: "FormFieldTime",
            associatedStructName: "FormFieldTime",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          formfieldGongNodeInstance.children!.push(FormFieldTimeGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation FormFieldTime
            */
          if (formfieldDB.FormFieldTime != undefined) {
            let formfieldGongNodeInstance_FormFieldTime: GongNode = {
              name: formfieldDB.FormFieldTime.Name,
              type: GongNodeType.INSTANCE,
              id: formfieldDB.FormFieldTime.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getFormFieldUniqueID(formfieldDB.ID)
                + 5 * getFormFieldTimeUniqueID(formfieldDB.FormFieldTime.ID),
              structName: "FormFieldTime",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            FormFieldTimeGongNodeAssociation.children.push(formfieldGongNodeInstance_FormFieldTime)
          }

        }
      )

      /**
      * fill up the FormFieldBoolean part of the mat tree
      */
      let formfieldbooleanGongNodeStruct: GongNode = {
        name: "FormFieldBoolean",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "FormFieldBoolean",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(formfieldbooleanGongNodeStruct)

      this.frontRepo.FormFieldBooleans_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.FormFieldBooleans_array.forEach(
        formfieldbooleanDB => {
          let formfieldbooleanGongNodeInstance: GongNode = {
            name: formfieldbooleanDB.Name,
            type: GongNodeType.INSTANCE,
            id: formfieldbooleanDB.ID,
            uniqueIdPerStack: getFormFieldBooleanUniqueID(formfieldbooleanDB.ID),
            structName: "FormFieldBoolean",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          formfieldbooleanGongNodeStruct.children!.push(formfieldbooleanGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the FormFieldDate part of the mat tree
      */
      let formfielddateGongNodeStruct: GongNode = {
        name: "FormFieldDate",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "FormFieldDate",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(formfielddateGongNodeStruct)

      this.frontRepo.FormFieldDates_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.FormFieldDates_array.forEach(
        formfielddateDB => {
          let formfielddateGongNodeInstance: GongNode = {
            name: formfielddateDB.Name,
            type: GongNodeType.INSTANCE,
            id: formfielddateDB.ID,
            uniqueIdPerStack: getFormFieldDateUniqueID(formfielddateDB.ID),
            structName: "FormFieldDate",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          formfielddateGongNodeStruct.children!.push(formfielddateGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the FormFieldFloat64 part of the mat tree
      */
      let formfieldfloat64GongNodeStruct: GongNode = {
        name: "FormFieldFloat64",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "FormFieldFloat64",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(formfieldfloat64GongNodeStruct)

      this.frontRepo.FormFieldFloat64s_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.FormFieldFloat64s_array.forEach(
        formfieldfloat64DB => {
          let formfieldfloat64GongNodeInstance: GongNode = {
            name: formfieldfloat64DB.Name,
            type: GongNodeType.INSTANCE,
            id: formfieldfloat64DB.ID,
            uniqueIdPerStack: getFormFieldFloat64UniqueID(formfieldfloat64DB.ID),
            structName: "FormFieldFloat64",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          formfieldfloat64GongNodeStruct.children!.push(formfieldfloat64GongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the FormFieldInt part of the mat tree
      */
      let formfieldintGongNodeStruct: GongNode = {
        name: "FormFieldInt",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "FormFieldInt",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(formfieldintGongNodeStruct)

      this.frontRepo.FormFieldInts_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.FormFieldInts_array.forEach(
        formfieldintDB => {
          let formfieldintGongNodeInstance: GongNode = {
            name: formfieldintDB.Name,
            type: GongNodeType.INSTANCE,
            id: formfieldintDB.ID,
            uniqueIdPerStack: getFormFieldIntUniqueID(formfieldintDB.ID),
            structName: "FormFieldInt",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          formfieldintGongNodeStruct.children!.push(formfieldintGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the FormFieldString part of the mat tree
      */
      let formfieldstringGongNodeStruct: GongNode = {
        name: "FormFieldString",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "FormFieldString",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(formfieldstringGongNodeStruct)

      this.frontRepo.FormFieldStrings_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.FormFieldStrings_array.forEach(
        formfieldstringDB => {
          let formfieldstringGongNodeInstance: GongNode = {
            name: formfieldstringDB.Name,
            type: GongNodeType.INSTANCE,
            id: formfieldstringDB.ID,
            uniqueIdPerStack: getFormFieldStringUniqueID(formfieldstringDB.ID),
            structName: "FormFieldString",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          formfieldstringGongNodeStruct.children!.push(formfieldstringGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the FormFieldTime part of the mat tree
      */
      let formfieldtimeGongNodeStruct: GongNode = {
        name: "FormFieldTime",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "FormFieldTime",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(formfieldtimeGongNodeStruct)

      this.frontRepo.FormFieldTimes_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.FormFieldTimes_array.forEach(
        formfieldtimeDB => {
          let formfieldtimeGongNodeInstance: GongNode = {
            name: formfieldtimeDB.Name,
            type: GongNodeType.INSTANCE,
            id: formfieldtimeDB.ID,
            uniqueIdPerStack: getFormFieldTimeUniqueID(formfieldtimeDB.ID),
            structName: "FormFieldTime",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          formfieldtimeGongNodeStruct.children!.push(formfieldtimeGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the FormGroup part of the mat tree
      */
      let formgroupGongNodeStruct: GongNode = {
        name: "FormGroup",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "FormGroup",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(formgroupGongNodeStruct)

      this.frontRepo.FormGroups_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.FormGroups_array.forEach(
        formgroupDB => {
          let formgroupGongNodeInstance: GongNode = {
            name: formgroupDB.Name,
            type: GongNodeType.INSTANCE,
            id: formgroupDB.ID,
            uniqueIdPerStack: getFormGroupUniqueID(formgroupDB.ID),
            structName: "FormGroup",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          formgroupGongNodeStruct.children!.push(formgroupGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer FormDivs
          */
          let FormDivsGongNodeAssociation: GongNode = {
            name: "(FormDiv) FormDivs",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: formgroupDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "FormGroup",
            associationField: "FormDivs",
            associatedStructName: "FormDiv",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          formgroupGongNodeInstance.children.push(FormDivsGongNodeAssociation)

          formgroupDB.FormDivs?.forEach(formdivDB => {
            let formdivNode: GongNode = {
              name: formdivDB.Name,
              type: GongNodeType.INSTANCE,
              id: formdivDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getFormGroupUniqueID(formgroupDB.ID)
                + 11 * getFormDivUniqueID(formdivDB.ID),
              structName: "FormDiv",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            FormDivsGongNodeAssociation.children.push(formdivNode)
          })

        }
      )

      /**
      * fill up the Row part of the mat tree
      */
      let rowGongNodeStruct: GongNode = {
        name: "Row",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Row",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(rowGongNodeStruct)

      this.frontRepo.Rows_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Rows_array.forEach(
        rowDB => {
          let rowGongNodeInstance: GongNode = {
            name: rowDB.Name,
            type: GongNodeType.INSTANCE,
            id: rowDB.ID,
            uniqueIdPerStack: getRowUniqueID(rowDB.ID),
            structName: "Row",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          rowGongNodeStruct.children!.push(rowGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Cells
          */
          let CellsGongNodeAssociation: GongNode = {
            name: "(Cell) Cells",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: rowDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Row",
            associationField: "Cells",
            associatedStructName: "Cell",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          rowGongNodeInstance.children.push(CellsGongNodeAssociation)

          rowDB.Cells?.forEach(cellDB => {
            let cellNode: GongNode = {
              name: cellDB.Name,
              type: GongNodeType.INSTANCE,
              id: cellDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getRowUniqueID(rowDB.ID)
                + 11 * getCellUniqueID(cellDB.ID),
              structName: "Cell",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            CellsGongNodeAssociation.children.push(cellNode)
          })

        }
      )

      /**
      * fill up the Table part of the mat tree
      */
      let tableGongNodeStruct: GongNode = {
        name: "Table",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Table",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(tableGongNodeStruct)

      this.frontRepo.Tables_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Tables_array.forEach(
        tableDB => {
          let tableGongNodeInstance: GongNode = {
            name: tableDB.Name,
            type: GongNodeType.INSTANCE,
            id: tableDB.ID,
            uniqueIdPerStack: getTableUniqueID(tableDB.ID),
            structName: "Table",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          tableGongNodeStruct.children!.push(tableGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer DisplayedColumns
          */
          let DisplayedColumnsGongNodeAssociation: GongNode = {
            name: "(DisplayedColumn) DisplayedColumns",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: tableDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Table",
            associationField: "DisplayedColumns",
            associatedStructName: "DisplayedColumn",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          tableGongNodeInstance.children.push(DisplayedColumnsGongNodeAssociation)

          tableDB.DisplayedColumns?.forEach(displayedcolumnDB => {
            let displayedcolumnNode: GongNode = {
              name: displayedcolumnDB.Name,
              type: GongNodeType.INSTANCE,
              id: displayedcolumnDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getTableUniqueID(tableDB.ID)
                + 11 * getDisplayedColumnUniqueID(displayedcolumnDB.ID),
              structName: "DisplayedColumn",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            DisplayedColumnsGongNodeAssociation.children.push(displayedcolumnNode)
          })

          /**
          * let append a node for the slide of pointer Rows
          */
          let RowsGongNodeAssociation: GongNode = {
            name: "(Row) Rows",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: tableDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Table",
            associationField: "Rows",
            associatedStructName: "Row",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          tableGongNodeInstance.children.push(RowsGongNodeAssociation)

          tableDB.Rows?.forEach(rowDB => {
            let rowNode: GongNode = {
              name: rowDB.Name,
              type: GongNodeType.INSTANCE,
              id: rowDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getTableUniqueID(tableDB.ID)
                + 11 * getRowUniqueID(rowDB.ID),
              structName: "Row",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            RowsGongNodeAssociation.children.push(rowNode)
          })

        }
      )


      this.dataSource.data = this.gongNodeTree

      // expand nodes that were exapanded before
      this.treeControl.dataNodes?.forEach(
        node => {
          if (memoryOfExpandedNodes.get(node.uniqueIdPerStack)) {
            this.treeControl.expand(node)
          }
        }
      )
    })
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutlet(path: string) {
    let outletName = this.routeService.getTableOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + path
    this.router.navigate([{
      outlets: {
        outletName: [fullPath]
      }
    }]);
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutletFromTree(path: string, type: GongNodeType, structName: string, id: number) {

    if (type == GongNodeType.STRUCT) {
      let outletName = this.routeService.getTableOutlet(this.GONG__StackPath)
      let fullPath = this.routeService.getPathRoot() + "-" + path.toLowerCase()
      let outletConf: any = {}
      outletConf[outletName] = [fullPath, this.GONG__StackPath]

      this.router.navigate([{ outlets: outletConf }])
    }

    if (type == GongNodeType.INSTANCE) {
      let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
      let fullPath = this.routeService.getPathRoot() + "-" + structName.toLowerCase() + "-detail"

      let outletConf: any = {}
      outletConf[outletName] = [fullPath, id, this.GONG__StackPath]

      this.router.navigate([{ outlets: outletConf }])
    }
  }

  setEditorRouterOutlet(path: string) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + path.toLowerCase()
    let outletConf : any = {}
    outletConf[outletName] = [fullPath, this.GONG__StackPath]
    this.router.navigate([{ outlets: outletConf }]);
  }

  setEditorSpecialRouterOutlet(node: GongFlatNode) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + node.associatedStructName.toLowerCase() + "-adder"
    this.router.navigate([{
      outlets: {
        outletName: [fullPath, node.id, node.structName, node.associationField, this.GONG__StackPath]
      }
    }]);
  }
}
