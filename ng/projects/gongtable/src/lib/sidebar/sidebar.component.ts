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
    private cellfloat64Service: CellFloat64Service,
    private celliconService: CellIconService,
    private cellintService: CellIntService,
    private cellstringService: CellStringService,
    private displayedcolumnService: DisplayedColumnService,
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
