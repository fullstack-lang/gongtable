import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest, BehaviorSubject } from 'rxjs';

// insertion point sub template for services imports 
import { CellDB } from './cell-db'
import { CellService } from './cell.service'

import { CellStringDB } from './cellstring-db'
import { CellStringService } from './cellstring.service'

import { RowDB } from './row-db'
import { RowService } from './row.service'

import { TableDB } from './table-db'
import { TableService } from './table.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  Cells_array = new Array<CellDB>(); // array of repo instances
  Cells = new Map<number, CellDB>(); // map of repo instances
  Cells_batch = new Map<number, CellDB>(); // same but only in last GET (for finding repo instances to delete)
  CellStrings_array = new Array<CellStringDB>(); // array of repo instances
  CellStrings = new Map<number, CellStringDB>(); // map of repo instances
  CellStrings_batch = new Map<number, CellStringDB>(); // same but only in last GET (for finding repo instances to delete)
  Rows_array = new Array<RowDB>(); // array of repo instances
  Rows = new Map<number, RowDB>(); // map of repo instances
  Rows_batch = new Map<number, RowDB>(); // same but only in last GET (for finding repo instances to delete)
  Tables_array = new Array<TableDB>(); // array of repo instances
  Tables = new Map<number, TableDB>(); // map of repo instances
  Tables_batch = new Map<number, TableDB>(); // same but only in last GET (for finding repo instances to delete)
}

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
  ID: number = 0 // ID of the calling instance

  // the reverse pointer is the name of the generated field on the destination
  // struct of the ONE-MANY association
  ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
  OrderingMode: boolean = false // if true, this is for ordering items

  // there are different selection mode : ONE_MANY or MANY_MANY
  SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

  // used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
  //
  // In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
  // 
  // in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
  // at the end of the ONE-MANY association
  SourceStruct: string = ""  // The "Aclass"
  SourceField: string = "" // the "AnarrayofbUse"
  IntermediateStruct: string = "" // the "AclassBclassUse" 
  IntermediateStructField: string = "" // the "Bclass" as field
  NextAssociationStruct: string = "" // the "Bclass"

  GONG__StackPath: string = ""
}

export enum SelectionMode {
  ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
  MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
  providedIn: 'root'
})
export class FrontRepoService {

  GONG__StackPath: string = ""

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  //
  // Store of all instances of the stack
  //
  frontRepo = new (FrontRepo)

  constructor(
    private http: HttpClient, // insertion point sub template 
    private cellService: CellService,
    private cellstringService: CellStringService,
    private rowService: RowService,
    private tableService: TableService,
  ) { }

  // postService provides a post function for each struct name
  postService(structName: string, instanceToBePosted: any) {
    let service = this[structName.toLowerCase() + "Service" + "Service" as keyof FrontRepoService]
    let servicePostFunction = service[("post" + structName) as keyof typeof service] as (instance: typeof instanceToBePosted) => Observable<typeof instanceToBePosted>

    servicePostFunction(instanceToBePosted).subscribe(
      instance => {
        let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged") as keyof typeof instanceToBePosted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("post")
      }
    );
  }

  // deleteService provides a delete function for each struct name
  deleteService(structName: string, instanceToBeDeleted: any) {
    let service = this[structName.toLowerCase() + "Service" as keyof FrontRepoService]
    let serviceDeleteFunction = service["delete" + structName as keyof typeof service] as (instance: typeof instanceToBeDeleted) => Observable<typeof instanceToBeDeleted>

    serviceDeleteFunction(instanceToBeDeleted).subscribe(
      instance => {
        let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged") as keyof typeof instanceToBeDeleted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("delete")
      }
    );
  }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [ // insertion point sub template 
    Observable<CellDB[]>,
    Observable<CellStringDB[]>,
    Observable<RowDB[]>,
    Observable<TableDB[]>,
  ] = [ // insertion point sub template
      this.cellService.getCells(this.GONG__StackPath),
      this.cellstringService.getCellStrings(this.GONG__StackPath),
      this.rowService.getRows(this.GONG__StackPath),
      this.tableService.getTables(this.GONG__StackPath),
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(GONG__StackPath: string = ""): Observable<FrontRepo> {

    this.GONG__StackPath = GONG__StackPath

    this.observableFrontRepo = [ // insertion point sub template
      this.cellService.getCells(this.GONG__StackPath),
      this.cellstringService.getCellStrings(this.GONG__StackPath),
      this.rowService.getRows(this.GONG__StackPath),
      this.tableService.getTables(this.GONG__StackPath),
    ]

    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([ // insertion point sub template for declarations 
            cells_,
            cellstrings_,
            rows_,
            tables_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var cells: CellDB[]
            cells = cells_ as CellDB[]
            var cellstrings: CellStringDB[]
            cellstrings = cellstrings_ as CellStringDB[]
            var rows: RowDB[]
            rows = rows_ as RowDB[]
            var tables: TableDB[]
            tables = tables_ as TableDB[]

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            this.frontRepo.Cells_array = cells

            // clear the map that counts Cell in the GET
            this.frontRepo.Cells_batch.clear()

            cells.forEach(
              cell => {
                this.frontRepo.Cells.set(cell.ID, cell)
                this.frontRepo.Cells_batch.set(cell.ID, cell)
              }
            )

            // clear cells that are absent from the batch
            this.frontRepo.Cells.forEach(
              cell => {
                if (this.frontRepo.Cells_batch.get(cell.ID) == undefined) {
                  this.frontRepo.Cells.delete(cell.ID)
                }
              }
            )

            // sort Cells_array array
            this.frontRepo.Cells_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.CellStrings_array = cellstrings

            // clear the map that counts CellString in the GET
            this.frontRepo.CellStrings_batch.clear()

            cellstrings.forEach(
              cellstring => {
                this.frontRepo.CellStrings.set(cellstring.ID, cellstring)
                this.frontRepo.CellStrings_batch.set(cellstring.ID, cellstring)
              }
            )

            // clear cellstrings that are absent from the batch
            this.frontRepo.CellStrings.forEach(
              cellstring => {
                if (this.frontRepo.CellStrings_batch.get(cellstring.ID) == undefined) {
                  this.frontRepo.CellStrings.delete(cellstring.ID)
                }
              }
            )

            // sort CellStrings_array array
            this.frontRepo.CellStrings_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Rows_array = rows

            // clear the map that counts Row in the GET
            this.frontRepo.Rows_batch.clear()

            rows.forEach(
              row => {
                this.frontRepo.Rows.set(row.ID, row)
                this.frontRepo.Rows_batch.set(row.ID, row)
              }
            )

            // clear rows that are absent from the batch
            this.frontRepo.Rows.forEach(
              row => {
                if (this.frontRepo.Rows_batch.get(row.ID) == undefined) {
                  this.frontRepo.Rows.delete(row.ID)
                }
              }
            )

            // sort Rows_array array
            this.frontRepo.Rows_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Tables_array = tables

            // clear the map that counts Table in the GET
            this.frontRepo.Tables_batch.clear()

            tables.forEach(
              table => {
                this.frontRepo.Tables.set(table.ID, table)
                this.frontRepo.Tables_batch.set(table.ID, table)
              }
            )

            // clear tables that are absent from the batch
            this.frontRepo.Tables.forEach(
              table => {
                if (this.frontRepo.Tables_batch.get(table.ID) == undefined) {
                  this.frontRepo.Tables.delete(table.ID)
                }
              }
            )

            // sort Tables_array array
            this.frontRepo.Tables_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });


            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template for redeem 
            cells.forEach(
              cell => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field CellString redeeming
                {
                  let _cellstring = this.frontRepo.CellStrings.get(cell.CellStringID.Int64)
                  if (_cellstring) {
                    cell.CellString = _cellstring
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Row.Cells redeeming
                {
                  let _row = this.frontRepo.Rows.get(cell.Row_CellsDBID.Int64)
                  if (_row) {
                    if (_row.Cells == undefined) {
                      _row.Cells = new Array<CellDB>()
                    }
                    _row.Cells.push(cell)
                    if (cell.Row_Cells_reverse == undefined) {
                      cell.Row_Cells_reverse = _row
                    }
                  }
                }
              }
            )
            cellstrings.forEach(
              cellstring => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            rows.forEach(
              row => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Table.Rows redeeming
                {
                  let _table = this.frontRepo.Tables.get(row.Table_RowsDBID.Int64)
                  if (_table) {
                    if (_table.Rows == undefined) {
                      _table.Rows = new Array<RowDB>()
                    }
                    _table.Rows.push(row)
                    if (row.Table_Rows_reverse == undefined) {
                      row.Table_Rows_reverse = _table
                    }
                  }
                }
              }
            )
            tables.forEach(
              table => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // insertion point for pull per struct 

  // CellPull performs a GET on Cell of the stack and redeem association pointers 
  CellPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.cellService.getCells(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            cells,
          ]) => {
            // init the array
            this.frontRepo.Cells_array = cells

            // clear the map that counts Cell in the GET
            this.frontRepo.Cells_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            cells.forEach(
              cell => {
                this.frontRepo.Cells.set(cell.ID, cell)
                this.frontRepo.Cells_batch.set(cell.ID, cell)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field CellString redeeming
                {
                  let _cellstring = this.frontRepo.CellStrings.get(cell.CellStringID.Int64)
                  if (_cellstring) {
                    cell.CellString = _cellstring
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Row.Cells redeeming
                {
                  let _row = this.frontRepo.Rows.get(cell.Row_CellsDBID.Int64)
                  if (_row) {
                    if (_row.Cells == undefined) {
                      _row.Cells = new Array<CellDB>()
                    }
                    _row.Cells.push(cell)
                    if (cell.Row_Cells_reverse == undefined) {
                      cell.Row_Cells_reverse = _row
                    }
                  }
                }
              }
            )

            // clear cells that are absent from the GET
            this.frontRepo.Cells.forEach(
              cell => {
                if (this.frontRepo.Cells_batch.get(cell.ID) == undefined) {
                  this.frontRepo.Cells.delete(cell.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // CellStringPull performs a GET on CellString of the stack and redeem association pointers 
  CellStringPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.cellstringService.getCellStrings(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            cellstrings,
          ]) => {
            // init the array
            this.frontRepo.CellStrings_array = cellstrings

            // clear the map that counts CellString in the GET
            this.frontRepo.CellStrings_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            cellstrings.forEach(
              cellstring => {
                this.frontRepo.CellStrings.set(cellstring.ID, cellstring)
                this.frontRepo.CellStrings_batch.set(cellstring.ID, cellstring)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear cellstrings that are absent from the GET
            this.frontRepo.CellStrings.forEach(
              cellstring => {
                if (this.frontRepo.CellStrings_batch.get(cellstring.ID) == undefined) {
                  this.frontRepo.CellStrings.delete(cellstring.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // RowPull performs a GET on Row of the stack and redeem association pointers 
  RowPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.rowService.getRows(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            rows,
          ]) => {
            // init the array
            this.frontRepo.Rows_array = rows

            // clear the map that counts Row in the GET
            this.frontRepo.Rows_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            rows.forEach(
              row => {
                this.frontRepo.Rows.set(row.ID, row)
                this.frontRepo.Rows_batch.set(row.ID, row)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Table.Rows redeeming
                {
                  let _table = this.frontRepo.Tables.get(row.Table_RowsDBID.Int64)
                  if (_table) {
                    if (_table.Rows == undefined) {
                      _table.Rows = new Array<RowDB>()
                    }
                    _table.Rows.push(row)
                    if (row.Table_Rows_reverse == undefined) {
                      row.Table_Rows_reverse = _table
                    }
                  }
                }
              }
            )

            // clear rows that are absent from the GET
            this.frontRepo.Rows.forEach(
              row => {
                if (this.frontRepo.Rows_batch.get(row.ID) == undefined) {
                  this.frontRepo.Rows.delete(row.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // TablePull performs a GET on Table of the stack and redeem association pointers 
  TablePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.tableService.getTables(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            tables,
          ]) => {
            // init the array
            this.frontRepo.Tables_array = tables

            // clear the map that counts Table in the GET
            this.frontRepo.Tables_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            tables.forEach(
              table => {
                this.frontRepo.Tables.set(table.ID, table)
                this.frontRepo.Tables_batch.set(table.ID, table)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear tables that are absent from the GET
            this.frontRepo.Tables.forEach(
              table => {
                if (this.frontRepo.Tables_batch.get(table.ID) == undefined) {
                  this.frontRepo.Tables.delete(table.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }
}

// insertion point for get unique ID per struct 
export function getCellUniqueID(id: number): number {
  return 31 * id
}
export function getCellStringUniqueID(id: number): number {
  return 37 * id
}
export function getRowUniqueID(id: number): number {
  return 41 * id
}
export function getTableUniqueID(id: number): number {
  return 43 * id
}
