import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest, BehaviorSubject } from 'rxjs';

// insertion point sub template for services imports 
import { CellDB } from './cell-db'
import { CellService } from './cell.service'

import { CellBooleanDB } from './cellboolean-db'
import { CellBooleanService } from './cellboolean.service'

import { CellFloat64DB } from './cellfloat64-db'
import { CellFloat64Service } from './cellfloat64.service'

import { CellIconDB } from './cellicon-db'
import { CellIconService } from './cellicon.service'

import { CellIntDB } from './cellint-db'
import { CellIntService } from './cellint.service'

import { CellStringDB } from './cellstring-db'
import { CellStringService } from './cellstring.service'

import { DisplayedColumnDB } from './displayedcolumn-db'
import { DisplayedColumnService } from './displayedcolumn.service'

import { FormDB } from './form-db'
import { FormService } from './form.service'

import { FormCellDB } from './formcell-db'
import { FormCellService } from './formcell.service'

import { FormCellBooleanDB } from './formcellboolean-db'
import { FormCellBooleanService } from './formcellboolean.service'

import { FormCellFloat64DB } from './formcellfloat64-db'
import { FormCellFloat64Service } from './formcellfloat64.service'

import { FormCellIntDB } from './formcellint-db'
import { FormCellIntService } from './formcellint.service'

import { FormCellStringDB } from './formcellstring-db'
import { FormCellStringService } from './formcellstring.service'

import { RowDB } from './row-db'
import { RowService } from './row.service'

import { TableDB } from './table-db'
import { TableService } from './table.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  Cells_array = new Array<CellDB>(); // array of repo instances
  Cells = new Map<number, CellDB>(); // map of repo instances
  Cells_batch = new Map<number, CellDB>(); // same but only in last GET (for finding repo instances to delete)
  CellBooleans_array = new Array<CellBooleanDB>(); // array of repo instances
  CellBooleans = new Map<number, CellBooleanDB>(); // map of repo instances
  CellBooleans_batch = new Map<number, CellBooleanDB>(); // same but only in last GET (for finding repo instances to delete)
  CellFloat64s_array = new Array<CellFloat64DB>(); // array of repo instances
  CellFloat64s = new Map<number, CellFloat64DB>(); // map of repo instances
  CellFloat64s_batch = new Map<number, CellFloat64DB>(); // same but only in last GET (for finding repo instances to delete)
  CellIcons_array = new Array<CellIconDB>(); // array of repo instances
  CellIcons = new Map<number, CellIconDB>(); // map of repo instances
  CellIcons_batch = new Map<number, CellIconDB>(); // same but only in last GET (for finding repo instances to delete)
  CellInts_array = new Array<CellIntDB>(); // array of repo instances
  CellInts = new Map<number, CellIntDB>(); // map of repo instances
  CellInts_batch = new Map<number, CellIntDB>(); // same but only in last GET (for finding repo instances to delete)
  CellStrings_array = new Array<CellStringDB>(); // array of repo instances
  CellStrings = new Map<number, CellStringDB>(); // map of repo instances
  CellStrings_batch = new Map<number, CellStringDB>(); // same but only in last GET (for finding repo instances to delete)
  DisplayedColumns_array = new Array<DisplayedColumnDB>(); // array of repo instances
  DisplayedColumns = new Map<number, DisplayedColumnDB>(); // map of repo instances
  DisplayedColumns_batch = new Map<number, DisplayedColumnDB>(); // same but only in last GET (for finding repo instances to delete)
  Forms_array = new Array<FormDB>(); // array of repo instances
  Forms = new Map<number, FormDB>(); // map of repo instances
  Forms_batch = new Map<number, FormDB>(); // same but only in last GET (for finding repo instances to delete)
  FormCells_array = new Array<FormCellDB>(); // array of repo instances
  FormCells = new Map<number, FormCellDB>(); // map of repo instances
  FormCells_batch = new Map<number, FormCellDB>(); // same but only in last GET (for finding repo instances to delete)
  FormCellBooleans_array = new Array<FormCellBooleanDB>(); // array of repo instances
  FormCellBooleans = new Map<number, FormCellBooleanDB>(); // map of repo instances
  FormCellBooleans_batch = new Map<number, FormCellBooleanDB>(); // same but only in last GET (for finding repo instances to delete)
  FormCellFloat64s_array = new Array<FormCellFloat64DB>(); // array of repo instances
  FormCellFloat64s = new Map<number, FormCellFloat64DB>(); // map of repo instances
  FormCellFloat64s_batch = new Map<number, FormCellFloat64DB>(); // same but only in last GET (for finding repo instances to delete)
  FormCellInts_array = new Array<FormCellIntDB>(); // array of repo instances
  FormCellInts = new Map<number, FormCellIntDB>(); // map of repo instances
  FormCellInts_batch = new Map<number, FormCellIntDB>(); // same but only in last GET (for finding repo instances to delete)
  FormCellStrings_array = new Array<FormCellStringDB>(); // array of repo instances
  FormCellStrings = new Map<number, FormCellStringDB>(); // map of repo instances
  FormCellStrings_batch = new Map<number, FormCellStringDB>(); // same but only in last GET (for finding repo instances to delete)
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
    private cellbooleanService: CellBooleanService,
    private cellfloat64Service: CellFloat64Service,
    private celliconService: CellIconService,
    private cellintService: CellIntService,
    private cellstringService: CellStringService,
    private displayedcolumnService: DisplayedColumnService,
    private formService: FormService,
    private formcellService: FormCellService,
    private formcellbooleanService: FormCellBooleanService,
    private formcellfloat64Service: FormCellFloat64Service,
    private formcellintService: FormCellIntService,
    private formcellstringService: FormCellStringService,
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
    Observable<CellBooleanDB[]>,
    Observable<CellFloat64DB[]>,
    Observable<CellIconDB[]>,
    Observable<CellIntDB[]>,
    Observable<CellStringDB[]>,
    Observable<DisplayedColumnDB[]>,
    Observable<FormDB[]>,
    Observable<FormCellDB[]>,
    Observable<FormCellBooleanDB[]>,
    Observable<FormCellFloat64DB[]>,
    Observable<FormCellIntDB[]>,
    Observable<FormCellStringDB[]>,
    Observable<RowDB[]>,
    Observable<TableDB[]>,
  ] = [ // insertion point sub template
      this.cellService.getCells(this.GONG__StackPath),
      this.cellbooleanService.getCellBooleans(this.GONG__StackPath),
      this.cellfloat64Service.getCellFloat64s(this.GONG__StackPath),
      this.celliconService.getCellIcons(this.GONG__StackPath),
      this.cellintService.getCellInts(this.GONG__StackPath),
      this.cellstringService.getCellStrings(this.GONG__StackPath),
      this.displayedcolumnService.getDisplayedColumns(this.GONG__StackPath),
      this.formService.getForms(this.GONG__StackPath),
      this.formcellService.getFormCells(this.GONG__StackPath),
      this.formcellbooleanService.getFormCellBooleans(this.GONG__StackPath),
      this.formcellfloat64Service.getFormCellFloat64s(this.GONG__StackPath),
      this.formcellintService.getFormCellInts(this.GONG__StackPath),
      this.formcellstringService.getFormCellStrings(this.GONG__StackPath),
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
      this.cellbooleanService.getCellBooleans(this.GONG__StackPath),
      this.cellfloat64Service.getCellFloat64s(this.GONG__StackPath),
      this.celliconService.getCellIcons(this.GONG__StackPath),
      this.cellintService.getCellInts(this.GONG__StackPath),
      this.cellstringService.getCellStrings(this.GONG__StackPath),
      this.displayedcolumnService.getDisplayedColumns(this.GONG__StackPath),
      this.formService.getForms(this.GONG__StackPath),
      this.formcellService.getFormCells(this.GONG__StackPath),
      this.formcellbooleanService.getFormCellBooleans(this.GONG__StackPath),
      this.formcellfloat64Service.getFormCellFloat64s(this.GONG__StackPath),
      this.formcellintService.getFormCellInts(this.GONG__StackPath),
      this.formcellstringService.getFormCellStrings(this.GONG__StackPath),
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
            cellbooleans_,
            cellfloat64s_,
            cellicons_,
            cellints_,
            cellstrings_,
            displayedcolumns_,
            forms_,
            formcells_,
            formcellbooleans_,
            formcellfloat64s_,
            formcellints_,
            formcellstrings_,
            rows_,
            tables_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var cells: CellDB[]
            cells = cells_ as CellDB[]
            var cellbooleans: CellBooleanDB[]
            cellbooleans = cellbooleans_ as CellBooleanDB[]
            var cellfloat64s: CellFloat64DB[]
            cellfloat64s = cellfloat64s_ as CellFloat64DB[]
            var cellicons: CellIconDB[]
            cellicons = cellicons_ as CellIconDB[]
            var cellints: CellIntDB[]
            cellints = cellints_ as CellIntDB[]
            var cellstrings: CellStringDB[]
            cellstrings = cellstrings_ as CellStringDB[]
            var displayedcolumns: DisplayedColumnDB[]
            displayedcolumns = displayedcolumns_ as DisplayedColumnDB[]
            var forms: FormDB[]
            forms = forms_ as FormDB[]
            var formcells: FormCellDB[]
            formcells = formcells_ as FormCellDB[]
            var formcellbooleans: FormCellBooleanDB[]
            formcellbooleans = formcellbooleans_ as FormCellBooleanDB[]
            var formcellfloat64s: FormCellFloat64DB[]
            formcellfloat64s = formcellfloat64s_ as FormCellFloat64DB[]
            var formcellints: FormCellIntDB[]
            formcellints = formcellints_ as FormCellIntDB[]
            var formcellstrings: FormCellStringDB[]
            formcellstrings = formcellstrings_ as FormCellStringDB[]
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
            this.frontRepo.CellBooleans_array = cellbooleans

            // clear the map that counts CellBoolean in the GET
            this.frontRepo.CellBooleans_batch.clear()

            cellbooleans.forEach(
              cellboolean => {
                this.frontRepo.CellBooleans.set(cellboolean.ID, cellboolean)
                this.frontRepo.CellBooleans_batch.set(cellboolean.ID, cellboolean)
              }
            )

            // clear cellbooleans that are absent from the batch
            this.frontRepo.CellBooleans.forEach(
              cellboolean => {
                if (this.frontRepo.CellBooleans_batch.get(cellboolean.ID) == undefined) {
                  this.frontRepo.CellBooleans.delete(cellboolean.ID)
                }
              }
            )

            // sort CellBooleans_array array
            this.frontRepo.CellBooleans_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.CellFloat64s_array = cellfloat64s

            // clear the map that counts CellFloat64 in the GET
            this.frontRepo.CellFloat64s_batch.clear()

            cellfloat64s.forEach(
              cellfloat64 => {
                this.frontRepo.CellFloat64s.set(cellfloat64.ID, cellfloat64)
                this.frontRepo.CellFloat64s_batch.set(cellfloat64.ID, cellfloat64)
              }
            )

            // clear cellfloat64s that are absent from the batch
            this.frontRepo.CellFloat64s.forEach(
              cellfloat64 => {
                if (this.frontRepo.CellFloat64s_batch.get(cellfloat64.ID) == undefined) {
                  this.frontRepo.CellFloat64s.delete(cellfloat64.ID)
                }
              }
            )

            // sort CellFloat64s_array array
            this.frontRepo.CellFloat64s_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.CellIcons_array = cellicons

            // clear the map that counts CellIcon in the GET
            this.frontRepo.CellIcons_batch.clear()

            cellicons.forEach(
              cellicon => {
                this.frontRepo.CellIcons.set(cellicon.ID, cellicon)
                this.frontRepo.CellIcons_batch.set(cellicon.ID, cellicon)
              }
            )

            // clear cellicons that are absent from the batch
            this.frontRepo.CellIcons.forEach(
              cellicon => {
                if (this.frontRepo.CellIcons_batch.get(cellicon.ID) == undefined) {
                  this.frontRepo.CellIcons.delete(cellicon.ID)
                }
              }
            )

            // sort CellIcons_array array
            this.frontRepo.CellIcons_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.CellInts_array = cellints

            // clear the map that counts CellInt in the GET
            this.frontRepo.CellInts_batch.clear()

            cellints.forEach(
              cellint => {
                this.frontRepo.CellInts.set(cellint.ID, cellint)
                this.frontRepo.CellInts_batch.set(cellint.ID, cellint)
              }
            )

            // clear cellints that are absent from the batch
            this.frontRepo.CellInts.forEach(
              cellint => {
                if (this.frontRepo.CellInts_batch.get(cellint.ID) == undefined) {
                  this.frontRepo.CellInts.delete(cellint.ID)
                }
              }
            )

            // sort CellInts_array array
            this.frontRepo.CellInts_array.sort((t1, t2) => {
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
            this.frontRepo.DisplayedColumns_array = displayedcolumns

            // clear the map that counts DisplayedColumn in the GET
            this.frontRepo.DisplayedColumns_batch.clear()

            displayedcolumns.forEach(
              displayedcolumn => {
                this.frontRepo.DisplayedColumns.set(displayedcolumn.ID, displayedcolumn)
                this.frontRepo.DisplayedColumns_batch.set(displayedcolumn.ID, displayedcolumn)
              }
            )

            // clear displayedcolumns that are absent from the batch
            this.frontRepo.DisplayedColumns.forEach(
              displayedcolumn => {
                if (this.frontRepo.DisplayedColumns_batch.get(displayedcolumn.ID) == undefined) {
                  this.frontRepo.DisplayedColumns.delete(displayedcolumn.ID)
                }
              }
            )

            // sort DisplayedColumns_array array
            this.frontRepo.DisplayedColumns_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Forms_array = forms

            // clear the map that counts Form in the GET
            this.frontRepo.Forms_batch.clear()

            forms.forEach(
              form => {
                this.frontRepo.Forms.set(form.ID, form)
                this.frontRepo.Forms_batch.set(form.ID, form)
              }
            )

            // clear forms that are absent from the batch
            this.frontRepo.Forms.forEach(
              form => {
                if (this.frontRepo.Forms_batch.get(form.ID) == undefined) {
                  this.frontRepo.Forms.delete(form.ID)
                }
              }
            )

            // sort Forms_array array
            this.frontRepo.Forms_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.FormCells_array = formcells

            // clear the map that counts FormCell in the GET
            this.frontRepo.FormCells_batch.clear()

            formcells.forEach(
              formcell => {
                this.frontRepo.FormCells.set(formcell.ID, formcell)
                this.frontRepo.FormCells_batch.set(formcell.ID, formcell)
              }
            )

            // clear formcells that are absent from the batch
            this.frontRepo.FormCells.forEach(
              formcell => {
                if (this.frontRepo.FormCells_batch.get(formcell.ID) == undefined) {
                  this.frontRepo.FormCells.delete(formcell.ID)
                }
              }
            )

            // sort FormCells_array array
            this.frontRepo.FormCells_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.FormCellBooleans_array = formcellbooleans

            // clear the map that counts FormCellBoolean in the GET
            this.frontRepo.FormCellBooleans_batch.clear()

            formcellbooleans.forEach(
              formcellboolean => {
                this.frontRepo.FormCellBooleans.set(formcellboolean.ID, formcellboolean)
                this.frontRepo.FormCellBooleans_batch.set(formcellboolean.ID, formcellboolean)
              }
            )

            // clear formcellbooleans that are absent from the batch
            this.frontRepo.FormCellBooleans.forEach(
              formcellboolean => {
                if (this.frontRepo.FormCellBooleans_batch.get(formcellboolean.ID) == undefined) {
                  this.frontRepo.FormCellBooleans.delete(formcellboolean.ID)
                }
              }
            )

            // sort FormCellBooleans_array array
            this.frontRepo.FormCellBooleans_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.FormCellFloat64s_array = formcellfloat64s

            // clear the map that counts FormCellFloat64 in the GET
            this.frontRepo.FormCellFloat64s_batch.clear()

            formcellfloat64s.forEach(
              formcellfloat64 => {
                this.frontRepo.FormCellFloat64s.set(formcellfloat64.ID, formcellfloat64)
                this.frontRepo.FormCellFloat64s_batch.set(formcellfloat64.ID, formcellfloat64)
              }
            )

            // clear formcellfloat64s that are absent from the batch
            this.frontRepo.FormCellFloat64s.forEach(
              formcellfloat64 => {
                if (this.frontRepo.FormCellFloat64s_batch.get(formcellfloat64.ID) == undefined) {
                  this.frontRepo.FormCellFloat64s.delete(formcellfloat64.ID)
                }
              }
            )

            // sort FormCellFloat64s_array array
            this.frontRepo.FormCellFloat64s_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.FormCellInts_array = formcellints

            // clear the map that counts FormCellInt in the GET
            this.frontRepo.FormCellInts_batch.clear()

            formcellints.forEach(
              formcellint => {
                this.frontRepo.FormCellInts.set(formcellint.ID, formcellint)
                this.frontRepo.FormCellInts_batch.set(formcellint.ID, formcellint)
              }
            )

            // clear formcellints that are absent from the batch
            this.frontRepo.FormCellInts.forEach(
              formcellint => {
                if (this.frontRepo.FormCellInts_batch.get(formcellint.ID) == undefined) {
                  this.frontRepo.FormCellInts.delete(formcellint.ID)
                }
              }
            )

            // sort FormCellInts_array array
            this.frontRepo.FormCellInts_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.FormCellStrings_array = formcellstrings

            // clear the map that counts FormCellString in the GET
            this.frontRepo.FormCellStrings_batch.clear()

            formcellstrings.forEach(
              formcellstring => {
                this.frontRepo.FormCellStrings.set(formcellstring.ID, formcellstring)
                this.frontRepo.FormCellStrings_batch.set(formcellstring.ID, formcellstring)
              }
            )

            // clear formcellstrings that are absent from the batch
            this.frontRepo.FormCellStrings.forEach(
              formcellstring => {
                if (this.frontRepo.FormCellStrings_batch.get(formcellstring.ID) == undefined) {
                  this.frontRepo.FormCellStrings.delete(formcellstring.ID)
                }
              }
            )

            // sort FormCellStrings_array array
            this.frontRepo.FormCellStrings_array.sort((t1, t2) => {
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
                // insertion point for pointer field CellFloat64 redeeming
                {
                  let _cellfloat64 = this.frontRepo.CellFloat64s.get(cell.CellFloat64ID.Int64)
                  if (_cellfloat64) {
                    cell.CellFloat64 = _cellfloat64
                  }
                }
                // insertion point for pointer field CellInt redeeming
                {
                  let _cellint = this.frontRepo.CellInts.get(cell.CellIntID.Int64)
                  if (_cellint) {
                    cell.CellInt = _cellint
                  }
                }
                // insertion point for pointer field CellBool redeeming
                {
                  let _cellboolean = this.frontRepo.CellBooleans.get(cell.CellBoolID.Int64)
                  if (_cellboolean) {
                    cell.CellBool = _cellboolean
                  }
                }
                // insertion point for pointer field CellIcon redeeming
                {
                  let _cellicon = this.frontRepo.CellIcons.get(cell.CellIconID.Int64)
                  if (_cellicon) {
                    cell.CellIcon = _cellicon
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
            cellbooleans.forEach(
              cellboolean => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            cellfloat64s.forEach(
              cellfloat64 => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            cellicons.forEach(
              cellicon => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            cellints.forEach(
              cellint => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            cellstrings.forEach(
              cellstring => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            displayedcolumns.forEach(
              displayedcolumn => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Table.DisplayedColumns redeeming
                {
                  let _table = this.frontRepo.Tables.get(displayedcolumn.Table_DisplayedColumnsDBID.Int64)
                  if (_table) {
                    if (_table.DisplayedColumns == undefined) {
                      _table.DisplayedColumns = new Array<DisplayedColumnDB>()
                    }
                    _table.DisplayedColumns.push(displayedcolumn)
                    if (displayedcolumn.Table_DisplayedColumns_reverse == undefined) {
                      displayedcolumn.Table_DisplayedColumns_reverse = _table
                    }
                  }
                }
              }
            )
            forms.forEach(
              form => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            formcells.forEach(
              formcell => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field FormCellString redeeming
                {
                  let _formcellstring = this.frontRepo.FormCellStrings.get(formcell.FormCellStringID.Int64)
                  if (_formcellstring) {
                    formcell.FormCellString = _formcellstring
                  }
                }
                // insertion point for pointer field FormCellFloat64 redeeming
                {
                  let _formcellfloat64 = this.frontRepo.FormCellFloat64s.get(formcell.FormCellFloat64ID.Int64)
                  if (_formcellfloat64) {
                    formcell.FormCellFloat64 = _formcellfloat64
                  }
                }
                // insertion point for pointer field FormCellInt redeeming
                {
                  let _formcellint = this.frontRepo.FormCellInts.get(formcell.FormCellIntID.Int64)
                  if (_formcellint) {
                    formcell.FormCellInt = _formcellint
                  }
                }
                // insertion point for pointer field FormCellBool redeeming
                {
                  let _formcellboolean = this.frontRepo.FormCellBooleans.get(formcell.FormCellBoolID.Int64)
                  if (_formcellboolean) {
                    formcell.FormCellBool = _formcellboolean
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Form.FormCells redeeming
                {
                  let _form = this.frontRepo.Forms.get(formcell.Form_FormCellsDBID.Int64)
                  if (_form) {
                    if (_form.FormCells == undefined) {
                      _form.FormCells = new Array<FormCellDB>()
                    }
                    _form.FormCells.push(formcell)
                    if (formcell.Form_FormCells_reverse == undefined) {
                      formcell.Form_FormCells_reverse = _form
                    }
                  }
                }
              }
            )
            formcellbooleans.forEach(
              formcellboolean => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            formcellfloat64s.forEach(
              formcellfloat64 => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            formcellints.forEach(
              formcellint => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            formcellstrings.forEach(
              formcellstring => {
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
                // insertion point for pointer field CellFloat64 redeeming
                {
                  let _cellfloat64 = this.frontRepo.CellFloat64s.get(cell.CellFloat64ID.Int64)
                  if (_cellfloat64) {
                    cell.CellFloat64 = _cellfloat64
                  }
                }
                // insertion point for pointer field CellInt redeeming
                {
                  let _cellint = this.frontRepo.CellInts.get(cell.CellIntID.Int64)
                  if (_cellint) {
                    cell.CellInt = _cellint
                  }
                }
                // insertion point for pointer field CellBool redeeming
                {
                  let _cellboolean = this.frontRepo.CellBooleans.get(cell.CellBoolID.Int64)
                  if (_cellboolean) {
                    cell.CellBool = _cellboolean
                  }
                }
                // insertion point for pointer field CellIcon redeeming
                {
                  let _cellicon = this.frontRepo.CellIcons.get(cell.CellIconID.Int64)
                  if (_cellicon) {
                    cell.CellIcon = _cellicon
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

  // CellBooleanPull performs a GET on CellBoolean of the stack and redeem association pointers 
  CellBooleanPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.cellbooleanService.getCellBooleans(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            cellbooleans,
          ]) => {
            // init the array
            this.frontRepo.CellBooleans_array = cellbooleans

            // clear the map that counts CellBoolean in the GET
            this.frontRepo.CellBooleans_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            cellbooleans.forEach(
              cellboolean => {
                this.frontRepo.CellBooleans.set(cellboolean.ID, cellboolean)
                this.frontRepo.CellBooleans_batch.set(cellboolean.ID, cellboolean)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear cellbooleans that are absent from the GET
            this.frontRepo.CellBooleans.forEach(
              cellboolean => {
                if (this.frontRepo.CellBooleans_batch.get(cellboolean.ID) == undefined) {
                  this.frontRepo.CellBooleans.delete(cellboolean.ID)
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

  // CellFloat64Pull performs a GET on CellFloat64 of the stack and redeem association pointers 
  CellFloat64Pull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.cellfloat64Service.getCellFloat64s(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            cellfloat64s,
          ]) => {
            // init the array
            this.frontRepo.CellFloat64s_array = cellfloat64s

            // clear the map that counts CellFloat64 in the GET
            this.frontRepo.CellFloat64s_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            cellfloat64s.forEach(
              cellfloat64 => {
                this.frontRepo.CellFloat64s.set(cellfloat64.ID, cellfloat64)
                this.frontRepo.CellFloat64s_batch.set(cellfloat64.ID, cellfloat64)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear cellfloat64s that are absent from the GET
            this.frontRepo.CellFloat64s.forEach(
              cellfloat64 => {
                if (this.frontRepo.CellFloat64s_batch.get(cellfloat64.ID) == undefined) {
                  this.frontRepo.CellFloat64s.delete(cellfloat64.ID)
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

  // CellIconPull performs a GET on CellIcon of the stack and redeem association pointers 
  CellIconPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.celliconService.getCellIcons(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            cellicons,
          ]) => {
            // init the array
            this.frontRepo.CellIcons_array = cellicons

            // clear the map that counts CellIcon in the GET
            this.frontRepo.CellIcons_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            cellicons.forEach(
              cellicon => {
                this.frontRepo.CellIcons.set(cellicon.ID, cellicon)
                this.frontRepo.CellIcons_batch.set(cellicon.ID, cellicon)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear cellicons that are absent from the GET
            this.frontRepo.CellIcons.forEach(
              cellicon => {
                if (this.frontRepo.CellIcons_batch.get(cellicon.ID) == undefined) {
                  this.frontRepo.CellIcons.delete(cellicon.ID)
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

  // CellIntPull performs a GET on CellInt of the stack and redeem association pointers 
  CellIntPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.cellintService.getCellInts(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            cellints,
          ]) => {
            // init the array
            this.frontRepo.CellInts_array = cellints

            // clear the map that counts CellInt in the GET
            this.frontRepo.CellInts_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            cellints.forEach(
              cellint => {
                this.frontRepo.CellInts.set(cellint.ID, cellint)
                this.frontRepo.CellInts_batch.set(cellint.ID, cellint)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear cellints that are absent from the GET
            this.frontRepo.CellInts.forEach(
              cellint => {
                if (this.frontRepo.CellInts_batch.get(cellint.ID) == undefined) {
                  this.frontRepo.CellInts.delete(cellint.ID)
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

  // DisplayedColumnPull performs a GET on DisplayedColumn of the stack and redeem association pointers 
  DisplayedColumnPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.displayedcolumnService.getDisplayedColumns(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            displayedcolumns,
          ]) => {
            // init the array
            this.frontRepo.DisplayedColumns_array = displayedcolumns

            // clear the map that counts DisplayedColumn in the GET
            this.frontRepo.DisplayedColumns_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            displayedcolumns.forEach(
              displayedcolumn => {
                this.frontRepo.DisplayedColumns.set(displayedcolumn.ID, displayedcolumn)
                this.frontRepo.DisplayedColumns_batch.set(displayedcolumn.ID, displayedcolumn)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Table.DisplayedColumns redeeming
                {
                  let _table = this.frontRepo.Tables.get(displayedcolumn.Table_DisplayedColumnsDBID.Int64)
                  if (_table) {
                    if (_table.DisplayedColumns == undefined) {
                      _table.DisplayedColumns = new Array<DisplayedColumnDB>()
                    }
                    _table.DisplayedColumns.push(displayedcolumn)
                    if (displayedcolumn.Table_DisplayedColumns_reverse == undefined) {
                      displayedcolumn.Table_DisplayedColumns_reverse = _table
                    }
                  }
                }
              }
            )

            // clear displayedcolumns that are absent from the GET
            this.frontRepo.DisplayedColumns.forEach(
              displayedcolumn => {
                if (this.frontRepo.DisplayedColumns_batch.get(displayedcolumn.ID) == undefined) {
                  this.frontRepo.DisplayedColumns.delete(displayedcolumn.ID)
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

  // FormPull performs a GET on Form of the stack and redeem association pointers 
  FormPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formService.getForms(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            forms,
          ]) => {
            // init the array
            this.frontRepo.Forms_array = forms

            // clear the map that counts Form in the GET
            this.frontRepo.Forms_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            forms.forEach(
              form => {
                this.frontRepo.Forms.set(form.ID, form)
                this.frontRepo.Forms_batch.set(form.ID, form)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear forms that are absent from the GET
            this.frontRepo.Forms.forEach(
              form => {
                if (this.frontRepo.Forms_batch.get(form.ID) == undefined) {
                  this.frontRepo.Forms.delete(form.ID)
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

  // FormCellPull performs a GET on FormCell of the stack and redeem association pointers 
  FormCellPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formcellService.getFormCells(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            formcells,
          ]) => {
            // init the array
            this.frontRepo.FormCells_array = formcells

            // clear the map that counts FormCell in the GET
            this.frontRepo.FormCells_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            formcells.forEach(
              formcell => {
                this.frontRepo.FormCells.set(formcell.ID, formcell)
                this.frontRepo.FormCells_batch.set(formcell.ID, formcell)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field FormCellString redeeming
                {
                  let _formcellstring = this.frontRepo.FormCellStrings.get(formcell.FormCellStringID.Int64)
                  if (_formcellstring) {
                    formcell.FormCellString = _formcellstring
                  }
                }
                // insertion point for pointer field FormCellFloat64 redeeming
                {
                  let _formcellfloat64 = this.frontRepo.FormCellFloat64s.get(formcell.FormCellFloat64ID.Int64)
                  if (_formcellfloat64) {
                    formcell.FormCellFloat64 = _formcellfloat64
                  }
                }
                // insertion point for pointer field FormCellInt redeeming
                {
                  let _formcellint = this.frontRepo.FormCellInts.get(formcell.FormCellIntID.Int64)
                  if (_formcellint) {
                    formcell.FormCellInt = _formcellint
                  }
                }
                // insertion point for pointer field FormCellBool redeeming
                {
                  let _formcellboolean = this.frontRepo.FormCellBooleans.get(formcell.FormCellBoolID.Int64)
                  if (_formcellboolean) {
                    formcell.FormCellBool = _formcellboolean
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Form.FormCells redeeming
                {
                  let _form = this.frontRepo.Forms.get(formcell.Form_FormCellsDBID.Int64)
                  if (_form) {
                    if (_form.FormCells == undefined) {
                      _form.FormCells = new Array<FormCellDB>()
                    }
                    _form.FormCells.push(formcell)
                    if (formcell.Form_FormCells_reverse == undefined) {
                      formcell.Form_FormCells_reverse = _form
                    }
                  }
                }
              }
            )

            // clear formcells that are absent from the GET
            this.frontRepo.FormCells.forEach(
              formcell => {
                if (this.frontRepo.FormCells_batch.get(formcell.ID) == undefined) {
                  this.frontRepo.FormCells.delete(formcell.ID)
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

  // FormCellBooleanPull performs a GET on FormCellBoolean of the stack and redeem association pointers 
  FormCellBooleanPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formcellbooleanService.getFormCellBooleans(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            formcellbooleans,
          ]) => {
            // init the array
            this.frontRepo.FormCellBooleans_array = formcellbooleans

            // clear the map that counts FormCellBoolean in the GET
            this.frontRepo.FormCellBooleans_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            formcellbooleans.forEach(
              formcellboolean => {
                this.frontRepo.FormCellBooleans.set(formcellboolean.ID, formcellboolean)
                this.frontRepo.FormCellBooleans_batch.set(formcellboolean.ID, formcellboolean)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear formcellbooleans that are absent from the GET
            this.frontRepo.FormCellBooleans.forEach(
              formcellboolean => {
                if (this.frontRepo.FormCellBooleans_batch.get(formcellboolean.ID) == undefined) {
                  this.frontRepo.FormCellBooleans.delete(formcellboolean.ID)
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

  // FormCellFloat64Pull performs a GET on FormCellFloat64 of the stack and redeem association pointers 
  FormCellFloat64Pull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formcellfloat64Service.getFormCellFloat64s(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            formcellfloat64s,
          ]) => {
            // init the array
            this.frontRepo.FormCellFloat64s_array = formcellfloat64s

            // clear the map that counts FormCellFloat64 in the GET
            this.frontRepo.FormCellFloat64s_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            formcellfloat64s.forEach(
              formcellfloat64 => {
                this.frontRepo.FormCellFloat64s.set(formcellfloat64.ID, formcellfloat64)
                this.frontRepo.FormCellFloat64s_batch.set(formcellfloat64.ID, formcellfloat64)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear formcellfloat64s that are absent from the GET
            this.frontRepo.FormCellFloat64s.forEach(
              formcellfloat64 => {
                if (this.frontRepo.FormCellFloat64s_batch.get(formcellfloat64.ID) == undefined) {
                  this.frontRepo.FormCellFloat64s.delete(formcellfloat64.ID)
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

  // FormCellIntPull performs a GET on FormCellInt of the stack and redeem association pointers 
  FormCellIntPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formcellintService.getFormCellInts(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            formcellints,
          ]) => {
            // init the array
            this.frontRepo.FormCellInts_array = formcellints

            // clear the map that counts FormCellInt in the GET
            this.frontRepo.FormCellInts_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            formcellints.forEach(
              formcellint => {
                this.frontRepo.FormCellInts.set(formcellint.ID, formcellint)
                this.frontRepo.FormCellInts_batch.set(formcellint.ID, formcellint)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear formcellints that are absent from the GET
            this.frontRepo.FormCellInts.forEach(
              formcellint => {
                if (this.frontRepo.FormCellInts_batch.get(formcellint.ID) == undefined) {
                  this.frontRepo.FormCellInts.delete(formcellint.ID)
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

  // FormCellStringPull performs a GET on FormCellString of the stack and redeem association pointers 
  FormCellStringPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formcellstringService.getFormCellStrings(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            formcellstrings,
          ]) => {
            // init the array
            this.frontRepo.FormCellStrings_array = formcellstrings

            // clear the map that counts FormCellString in the GET
            this.frontRepo.FormCellStrings_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            formcellstrings.forEach(
              formcellstring => {
                this.frontRepo.FormCellStrings.set(formcellstring.ID, formcellstring)
                this.frontRepo.FormCellStrings_batch.set(formcellstring.ID, formcellstring)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear formcellstrings that are absent from the GET
            this.frontRepo.FormCellStrings.forEach(
              formcellstring => {
                if (this.frontRepo.FormCellStrings_batch.get(formcellstring.ID) == undefined) {
                  this.frontRepo.FormCellStrings.delete(formcellstring.ID)
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
export function getCellBooleanUniqueID(id: number): number {
  return 37 * id
}
export function getCellFloat64UniqueID(id: number): number {
  return 41 * id
}
export function getCellIconUniqueID(id: number): number {
  return 43 * id
}
export function getCellIntUniqueID(id: number): number {
  return 47 * id
}
export function getCellStringUniqueID(id: number): number {
  return 53 * id
}
export function getDisplayedColumnUniqueID(id: number): number {
  return 59 * id
}
export function getFormUniqueID(id: number): number {
  return 61 * id
}
export function getFormCellUniqueID(id: number): number {
  return 67 * id
}
export function getFormCellBooleanUniqueID(id: number): number {
  return 71 * id
}
export function getFormCellFloat64UniqueID(id: number): number {
  return 73 * id
}
export function getFormCellIntUniqueID(id: number): number {
  return 79 * id
}
export function getFormCellStringUniqueID(id: number): number {
  return 83 * id
}
export function getRowUniqueID(id: number): number {
  return 89 * id
}
export function getTableUniqueID(id: number): number {
  return 97 * id
}
