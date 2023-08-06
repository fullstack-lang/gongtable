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

import { FormDivDB } from './formdiv-db'
import { FormDivService } from './formdiv.service'

import { FormFieldDB } from './formfield-db'
import { FormFieldService } from './formfield.service'

import { FormFieldBooleanDB } from './formfieldboolean-db'
import { FormFieldBooleanService } from './formfieldboolean.service'

import { FormFieldDateDB } from './formfielddate-db'
import { FormFieldDateService } from './formfielddate.service'

import { FormFieldFloat64DB } from './formfieldfloat64-db'
import { FormFieldFloat64Service } from './formfieldfloat64.service'

import { FormFieldIntDB } from './formfieldint-db'
import { FormFieldIntService } from './formfieldint.service'

import { FormFieldStringDB } from './formfieldstring-db'
import { FormFieldStringService } from './formfieldstring.service'

import { FormFieldTimeDB } from './formfieldtime-db'
import { FormFieldTimeService } from './formfieldtime.service'

import { FormGroupDB } from './formgroup-db'
import { FormGroupService } from './formgroup.service'

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
  FormDivs_array = new Array<FormDivDB>(); // array of repo instances
  FormDivs = new Map<number, FormDivDB>(); // map of repo instances
  FormDivs_batch = new Map<number, FormDivDB>(); // same but only in last GET (for finding repo instances to delete)
  FormFields_array = new Array<FormFieldDB>(); // array of repo instances
  FormFields = new Map<number, FormFieldDB>(); // map of repo instances
  FormFields_batch = new Map<number, FormFieldDB>(); // same but only in last GET (for finding repo instances to delete)
  FormFieldBooleans_array = new Array<FormFieldBooleanDB>(); // array of repo instances
  FormFieldBooleans = new Map<number, FormFieldBooleanDB>(); // map of repo instances
  FormFieldBooleans_batch = new Map<number, FormFieldBooleanDB>(); // same but only in last GET (for finding repo instances to delete)
  FormFieldDates_array = new Array<FormFieldDateDB>(); // array of repo instances
  FormFieldDates = new Map<number, FormFieldDateDB>(); // map of repo instances
  FormFieldDates_batch = new Map<number, FormFieldDateDB>(); // same but only in last GET (for finding repo instances to delete)
  FormFieldFloat64s_array = new Array<FormFieldFloat64DB>(); // array of repo instances
  FormFieldFloat64s = new Map<number, FormFieldFloat64DB>(); // map of repo instances
  FormFieldFloat64s_batch = new Map<number, FormFieldFloat64DB>(); // same but only in last GET (for finding repo instances to delete)
  FormFieldInts_array = new Array<FormFieldIntDB>(); // array of repo instances
  FormFieldInts = new Map<number, FormFieldIntDB>(); // map of repo instances
  FormFieldInts_batch = new Map<number, FormFieldIntDB>(); // same but only in last GET (for finding repo instances to delete)
  FormFieldStrings_array = new Array<FormFieldStringDB>(); // array of repo instances
  FormFieldStrings = new Map<number, FormFieldStringDB>(); // map of repo instances
  FormFieldStrings_batch = new Map<number, FormFieldStringDB>(); // same but only in last GET (for finding repo instances to delete)
  FormFieldTimes_array = new Array<FormFieldTimeDB>(); // array of repo instances
  FormFieldTimes = new Map<number, FormFieldTimeDB>(); // map of repo instances
  FormFieldTimes_batch = new Map<number, FormFieldTimeDB>(); // same but only in last GET (for finding repo instances to delete)
  FormGroups_array = new Array<FormGroupDB>(); // array of repo instances
  FormGroups = new Map<number, FormGroupDB>(); // map of repo instances
  FormGroups_batch = new Map<number, FormGroupDB>(); // same but only in last GET (for finding repo instances to delete)
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
    Observable<FormDivDB[]>,
    Observable<FormFieldDB[]>,
    Observable<FormFieldBooleanDB[]>,
    Observable<FormFieldDateDB[]>,
    Observable<FormFieldFloat64DB[]>,
    Observable<FormFieldIntDB[]>,
    Observable<FormFieldStringDB[]>,
    Observable<FormFieldTimeDB[]>,
    Observable<FormGroupDB[]>,
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
      this.formdivService.getFormDivs(this.GONG__StackPath),
      this.formfieldService.getFormFields(this.GONG__StackPath),
      this.formfieldbooleanService.getFormFieldBooleans(this.GONG__StackPath),
      this.formfielddateService.getFormFieldDates(this.GONG__StackPath),
      this.formfieldfloat64Service.getFormFieldFloat64s(this.GONG__StackPath),
      this.formfieldintService.getFormFieldInts(this.GONG__StackPath),
      this.formfieldstringService.getFormFieldStrings(this.GONG__StackPath),
      this.formfieldtimeService.getFormFieldTimes(this.GONG__StackPath),
      this.formgroupService.getFormGroups(this.GONG__StackPath),
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
      this.formdivService.getFormDivs(this.GONG__StackPath),
      this.formfieldService.getFormFields(this.GONG__StackPath),
      this.formfieldbooleanService.getFormFieldBooleans(this.GONG__StackPath),
      this.formfielddateService.getFormFieldDates(this.GONG__StackPath),
      this.formfieldfloat64Service.getFormFieldFloat64s(this.GONG__StackPath),
      this.formfieldintService.getFormFieldInts(this.GONG__StackPath),
      this.formfieldstringService.getFormFieldStrings(this.GONG__StackPath),
      this.formfieldtimeService.getFormFieldTimes(this.GONG__StackPath),
      this.formgroupService.getFormGroups(this.GONG__StackPath),
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
            formdivs_,
            formfields_,
            formfieldbooleans_,
            formfielddates_,
            formfieldfloat64s_,
            formfieldints_,
            formfieldstrings_,
            formfieldtimes_,
            formgroups_,
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
            var formdivs: FormDivDB[]
            formdivs = formdivs_ as FormDivDB[]
            var formfields: FormFieldDB[]
            formfields = formfields_ as FormFieldDB[]
            var formfieldbooleans: FormFieldBooleanDB[]
            formfieldbooleans = formfieldbooleans_ as FormFieldBooleanDB[]
            var formfielddates: FormFieldDateDB[]
            formfielddates = formfielddates_ as FormFieldDateDB[]
            var formfieldfloat64s: FormFieldFloat64DB[]
            formfieldfloat64s = formfieldfloat64s_ as FormFieldFloat64DB[]
            var formfieldints: FormFieldIntDB[]
            formfieldints = formfieldints_ as FormFieldIntDB[]
            var formfieldstrings: FormFieldStringDB[]
            formfieldstrings = formfieldstrings_ as FormFieldStringDB[]
            var formfieldtimes: FormFieldTimeDB[]
            formfieldtimes = formfieldtimes_ as FormFieldTimeDB[]
            var formgroups: FormGroupDB[]
            formgroups = formgroups_ as FormGroupDB[]
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
            this.frontRepo.FormDivs_array = formdivs

            // clear the map that counts FormDiv in the GET
            this.frontRepo.FormDivs_batch.clear()

            formdivs.forEach(
              formdiv => {
                this.frontRepo.FormDivs.set(formdiv.ID, formdiv)
                this.frontRepo.FormDivs_batch.set(formdiv.ID, formdiv)
              }
            )

            // clear formdivs that are absent from the batch
            this.frontRepo.FormDivs.forEach(
              formdiv => {
                if (this.frontRepo.FormDivs_batch.get(formdiv.ID) == undefined) {
                  this.frontRepo.FormDivs.delete(formdiv.ID)
                }
              }
            )

            // sort FormDivs_array array
            this.frontRepo.FormDivs_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.FormFields_array = formfields

            // clear the map that counts FormField in the GET
            this.frontRepo.FormFields_batch.clear()

            formfields.forEach(
              formfield => {
                this.frontRepo.FormFields.set(formfield.ID, formfield)
                this.frontRepo.FormFields_batch.set(formfield.ID, formfield)
              }
            )

            // clear formfields that are absent from the batch
            this.frontRepo.FormFields.forEach(
              formfield => {
                if (this.frontRepo.FormFields_batch.get(formfield.ID) == undefined) {
                  this.frontRepo.FormFields.delete(formfield.ID)
                }
              }
            )

            // sort FormFields_array array
            this.frontRepo.FormFields_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.FormFieldBooleans_array = formfieldbooleans

            // clear the map that counts FormFieldBoolean in the GET
            this.frontRepo.FormFieldBooleans_batch.clear()

            formfieldbooleans.forEach(
              formfieldboolean => {
                this.frontRepo.FormFieldBooleans.set(formfieldboolean.ID, formfieldboolean)
                this.frontRepo.FormFieldBooleans_batch.set(formfieldboolean.ID, formfieldboolean)
              }
            )

            // clear formfieldbooleans that are absent from the batch
            this.frontRepo.FormFieldBooleans.forEach(
              formfieldboolean => {
                if (this.frontRepo.FormFieldBooleans_batch.get(formfieldboolean.ID) == undefined) {
                  this.frontRepo.FormFieldBooleans.delete(formfieldboolean.ID)
                }
              }
            )

            // sort FormFieldBooleans_array array
            this.frontRepo.FormFieldBooleans_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.FormFieldDates_array = formfielddates

            // clear the map that counts FormFieldDate in the GET
            this.frontRepo.FormFieldDates_batch.clear()

            formfielddates.forEach(
              formfielddate => {
                this.frontRepo.FormFieldDates.set(formfielddate.ID, formfielddate)
                this.frontRepo.FormFieldDates_batch.set(formfielddate.ID, formfielddate)
              }
            )

            // clear formfielddates that are absent from the batch
            this.frontRepo.FormFieldDates.forEach(
              formfielddate => {
                if (this.frontRepo.FormFieldDates_batch.get(formfielddate.ID) == undefined) {
                  this.frontRepo.FormFieldDates.delete(formfielddate.ID)
                }
              }
            )

            // sort FormFieldDates_array array
            this.frontRepo.FormFieldDates_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.FormFieldFloat64s_array = formfieldfloat64s

            // clear the map that counts FormFieldFloat64 in the GET
            this.frontRepo.FormFieldFloat64s_batch.clear()

            formfieldfloat64s.forEach(
              formfieldfloat64 => {
                this.frontRepo.FormFieldFloat64s.set(formfieldfloat64.ID, formfieldfloat64)
                this.frontRepo.FormFieldFloat64s_batch.set(formfieldfloat64.ID, formfieldfloat64)
              }
            )

            // clear formfieldfloat64s that are absent from the batch
            this.frontRepo.FormFieldFloat64s.forEach(
              formfieldfloat64 => {
                if (this.frontRepo.FormFieldFloat64s_batch.get(formfieldfloat64.ID) == undefined) {
                  this.frontRepo.FormFieldFloat64s.delete(formfieldfloat64.ID)
                }
              }
            )

            // sort FormFieldFloat64s_array array
            this.frontRepo.FormFieldFloat64s_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.FormFieldInts_array = formfieldints

            // clear the map that counts FormFieldInt in the GET
            this.frontRepo.FormFieldInts_batch.clear()

            formfieldints.forEach(
              formfieldint => {
                this.frontRepo.FormFieldInts.set(formfieldint.ID, formfieldint)
                this.frontRepo.FormFieldInts_batch.set(formfieldint.ID, formfieldint)
              }
            )

            // clear formfieldints that are absent from the batch
            this.frontRepo.FormFieldInts.forEach(
              formfieldint => {
                if (this.frontRepo.FormFieldInts_batch.get(formfieldint.ID) == undefined) {
                  this.frontRepo.FormFieldInts.delete(formfieldint.ID)
                }
              }
            )

            // sort FormFieldInts_array array
            this.frontRepo.FormFieldInts_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.FormFieldStrings_array = formfieldstrings

            // clear the map that counts FormFieldString in the GET
            this.frontRepo.FormFieldStrings_batch.clear()

            formfieldstrings.forEach(
              formfieldstring => {
                this.frontRepo.FormFieldStrings.set(formfieldstring.ID, formfieldstring)
                this.frontRepo.FormFieldStrings_batch.set(formfieldstring.ID, formfieldstring)
              }
            )

            // clear formfieldstrings that are absent from the batch
            this.frontRepo.FormFieldStrings.forEach(
              formfieldstring => {
                if (this.frontRepo.FormFieldStrings_batch.get(formfieldstring.ID) == undefined) {
                  this.frontRepo.FormFieldStrings.delete(formfieldstring.ID)
                }
              }
            )

            // sort FormFieldStrings_array array
            this.frontRepo.FormFieldStrings_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.FormFieldTimes_array = formfieldtimes

            // clear the map that counts FormFieldTime in the GET
            this.frontRepo.FormFieldTimes_batch.clear()

            formfieldtimes.forEach(
              formfieldtime => {
                this.frontRepo.FormFieldTimes.set(formfieldtime.ID, formfieldtime)
                this.frontRepo.FormFieldTimes_batch.set(formfieldtime.ID, formfieldtime)
              }
            )

            // clear formfieldtimes that are absent from the batch
            this.frontRepo.FormFieldTimes.forEach(
              formfieldtime => {
                if (this.frontRepo.FormFieldTimes_batch.get(formfieldtime.ID) == undefined) {
                  this.frontRepo.FormFieldTimes.delete(formfieldtime.ID)
                }
              }
            )

            // sort FormFieldTimes_array array
            this.frontRepo.FormFieldTimes_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.FormGroups_array = formgroups

            // clear the map that counts FormGroup in the GET
            this.frontRepo.FormGroups_batch.clear()

            formgroups.forEach(
              formgroup => {
                this.frontRepo.FormGroups.set(formgroup.ID, formgroup)
                this.frontRepo.FormGroups_batch.set(formgroup.ID, formgroup)
              }
            )

            // clear formgroups that are absent from the batch
            this.frontRepo.FormGroups.forEach(
              formgroup => {
                if (this.frontRepo.FormGroups_batch.get(formgroup.ID) == undefined) {
                  this.frontRepo.FormGroups.delete(formgroup.ID)
                }
              }
            )

            // sort FormGroups_array array
            this.frontRepo.FormGroups_array.sort((t1, t2) => {
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
            formdivs.forEach(
              formdiv => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field FormGroup.FormDivs redeeming
                {
                  let _formgroup = this.frontRepo.FormGroups.get(formdiv.FormGroup_FormDivsDBID.Int64)
                  if (_formgroup) {
                    if (_formgroup.FormDivs == undefined) {
                      _formgroup.FormDivs = new Array<FormDivDB>()
                    }
                    _formgroup.FormDivs.push(formdiv)
                    if (formdiv.FormGroup_FormDivs_reverse == undefined) {
                      formdiv.FormGroup_FormDivs_reverse = _formgroup
                    }
                  }
                }
              }
            )
            formfields.forEach(
              formfield => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field FormFieldString redeeming
                {
                  let _formfieldstring = this.frontRepo.FormFieldStrings.get(formfield.FormFieldStringID.Int64)
                  if (_formfieldstring) {
                    formfield.FormFieldString = _formfieldstring
                  }
                }
                // insertion point for pointer field FormFieldFloat64 redeeming
                {
                  let _formfieldfloat64 = this.frontRepo.FormFieldFloat64s.get(formfield.FormFieldFloat64ID.Int64)
                  if (_formfieldfloat64) {
                    formfield.FormFieldFloat64 = _formfieldfloat64
                  }
                }
                // insertion point for pointer field FormFieldInt redeeming
                {
                  let _formfieldint = this.frontRepo.FormFieldInts.get(formfield.FormFieldIntID.Int64)
                  if (_formfieldint) {
                    formfield.FormFieldInt = _formfieldint
                  }
                }
                // insertion point for pointer field FormFieldBool redeeming
                {
                  let _formfieldboolean = this.frontRepo.FormFieldBooleans.get(formfield.FormFieldBoolID.Int64)
                  if (_formfieldboolean) {
                    formfield.FormFieldBool = _formfieldboolean
                  }
                }
                // insertion point for pointer field FormFieldDate redeeming
                {
                  let _formfielddate = this.frontRepo.FormFieldDates.get(formfield.FormFieldDateID.Int64)
                  if (_formfielddate) {
                    formfield.FormFieldDate = _formfielddate
                  }
                }
                // insertion point for pointer field FormFieldTime redeeming
                {
                  let _formfieldtime = this.frontRepo.FormFieldTimes.get(formfield.FormFieldTimeID.Int64)
                  if (_formfieldtime) {
                    formfield.FormFieldTime = _formfieldtime
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field FormDiv.FormFields redeeming
                {
                  let _formdiv = this.frontRepo.FormDivs.get(formfield.FormDiv_FormFieldsDBID.Int64)
                  if (_formdiv) {
                    if (_formdiv.FormFields == undefined) {
                      _formdiv.FormFields = new Array<FormFieldDB>()
                    }
                    _formdiv.FormFields.push(formfield)
                    if (formfield.FormDiv_FormFields_reverse == undefined) {
                      formfield.FormDiv_FormFields_reverse = _formdiv
                    }
                  }
                }
              }
            )
            formfieldbooleans.forEach(
              formfieldboolean => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            formfielddates.forEach(
              formfielddate => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            formfieldfloat64s.forEach(
              formfieldfloat64 => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            formfieldints.forEach(
              formfieldint => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            formfieldstrings.forEach(
              formfieldstring => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            formfieldtimes.forEach(
              formfieldtime => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            formgroups.forEach(
              formgroup => {
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

  // FormDivPull performs a GET on FormDiv of the stack and redeem association pointers 
  FormDivPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formdivService.getFormDivs(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            formdivs,
          ]) => {
            // init the array
            this.frontRepo.FormDivs_array = formdivs

            // clear the map that counts FormDiv in the GET
            this.frontRepo.FormDivs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            formdivs.forEach(
              formdiv => {
                this.frontRepo.FormDivs.set(formdiv.ID, formdiv)
                this.frontRepo.FormDivs_batch.set(formdiv.ID, formdiv)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field FormGroup.FormDivs redeeming
                {
                  let _formgroup = this.frontRepo.FormGroups.get(formdiv.FormGroup_FormDivsDBID.Int64)
                  if (_formgroup) {
                    if (_formgroup.FormDivs == undefined) {
                      _formgroup.FormDivs = new Array<FormDivDB>()
                    }
                    _formgroup.FormDivs.push(formdiv)
                    if (formdiv.FormGroup_FormDivs_reverse == undefined) {
                      formdiv.FormGroup_FormDivs_reverse = _formgroup
                    }
                  }
                }
              }
            )

            // clear formdivs that are absent from the GET
            this.frontRepo.FormDivs.forEach(
              formdiv => {
                if (this.frontRepo.FormDivs_batch.get(formdiv.ID) == undefined) {
                  this.frontRepo.FormDivs.delete(formdiv.ID)
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

  // FormFieldPull performs a GET on FormField of the stack and redeem association pointers 
  FormFieldPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formfieldService.getFormFields(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            formfields,
          ]) => {
            // init the array
            this.frontRepo.FormFields_array = formfields

            // clear the map that counts FormField in the GET
            this.frontRepo.FormFields_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            formfields.forEach(
              formfield => {
                this.frontRepo.FormFields.set(formfield.ID, formfield)
                this.frontRepo.FormFields_batch.set(formfield.ID, formfield)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field FormFieldString redeeming
                {
                  let _formfieldstring = this.frontRepo.FormFieldStrings.get(formfield.FormFieldStringID.Int64)
                  if (_formfieldstring) {
                    formfield.FormFieldString = _formfieldstring
                  }
                }
                // insertion point for pointer field FormFieldFloat64 redeeming
                {
                  let _formfieldfloat64 = this.frontRepo.FormFieldFloat64s.get(formfield.FormFieldFloat64ID.Int64)
                  if (_formfieldfloat64) {
                    formfield.FormFieldFloat64 = _formfieldfloat64
                  }
                }
                // insertion point for pointer field FormFieldInt redeeming
                {
                  let _formfieldint = this.frontRepo.FormFieldInts.get(formfield.FormFieldIntID.Int64)
                  if (_formfieldint) {
                    formfield.FormFieldInt = _formfieldint
                  }
                }
                // insertion point for pointer field FormFieldBool redeeming
                {
                  let _formfieldboolean = this.frontRepo.FormFieldBooleans.get(formfield.FormFieldBoolID.Int64)
                  if (_formfieldboolean) {
                    formfield.FormFieldBool = _formfieldboolean
                  }
                }
                // insertion point for pointer field FormFieldDate redeeming
                {
                  let _formfielddate = this.frontRepo.FormFieldDates.get(formfield.FormFieldDateID.Int64)
                  if (_formfielddate) {
                    formfield.FormFieldDate = _formfielddate
                  }
                }
                // insertion point for pointer field FormFieldTime redeeming
                {
                  let _formfieldtime = this.frontRepo.FormFieldTimes.get(formfield.FormFieldTimeID.Int64)
                  if (_formfieldtime) {
                    formfield.FormFieldTime = _formfieldtime
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field FormDiv.FormFields redeeming
                {
                  let _formdiv = this.frontRepo.FormDivs.get(formfield.FormDiv_FormFieldsDBID.Int64)
                  if (_formdiv) {
                    if (_formdiv.FormFields == undefined) {
                      _formdiv.FormFields = new Array<FormFieldDB>()
                    }
                    _formdiv.FormFields.push(formfield)
                    if (formfield.FormDiv_FormFields_reverse == undefined) {
                      formfield.FormDiv_FormFields_reverse = _formdiv
                    }
                  }
                }
              }
            )

            // clear formfields that are absent from the GET
            this.frontRepo.FormFields.forEach(
              formfield => {
                if (this.frontRepo.FormFields_batch.get(formfield.ID) == undefined) {
                  this.frontRepo.FormFields.delete(formfield.ID)
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

  // FormFieldBooleanPull performs a GET on FormFieldBoolean of the stack and redeem association pointers 
  FormFieldBooleanPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formfieldbooleanService.getFormFieldBooleans(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            formfieldbooleans,
          ]) => {
            // init the array
            this.frontRepo.FormFieldBooleans_array = formfieldbooleans

            // clear the map that counts FormFieldBoolean in the GET
            this.frontRepo.FormFieldBooleans_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            formfieldbooleans.forEach(
              formfieldboolean => {
                this.frontRepo.FormFieldBooleans.set(formfieldboolean.ID, formfieldboolean)
                this.frontRepo.FormFieldBooleans_batch.set(formfieldboolean.ID, formfieldboolean)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear formfieldbooleans that are absent from the GET
            this.frontRepo.FormFieldBooleans.forEach(
              formfieldboolean => {
                if (this.frontRepo.FormFieldBooleans_batch.get(formfieldboolean.ID) == undefined) {
                  this.frontRepo.FormFieldBooleans.delete(formfieldboolean.ID)
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

  // FormFieldDatePull performs a GET on FormFieldDate of the stack and redeem association pointers 
  FormFieldDatePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formfielddateService.getFormFieldDates(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            formfielddates,
          ]) => {
            // init the array
            this.frontRepo.FormFieldDates_array = formfielddates

            // clear the map that counts FormFieldDate in the GET
            this.frontRepo.FormFieldDates_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            formfielddates.forEach(
              formfielddate => {
                this.frontRepo.FormFieldDates.set(formfielddate.ID, formfielddate)
                this.frontRepo.FormFieldDates_batch.set(formfielddate.ID, formfielddate)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear formfielddates that are absent from the GET
            this.frontRepo.FormFieldDates.forEach(
              formfielddate => {
                if (this.frontRepo.FormFieldDates_batch.get(formfielddate.ID) == undefined) {
                  this.frontRepo.FormFieldDates.delete(formfielddate.ID)
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

  // FormFieldFloat64Pull performs a GET on FormFieldFloat64 of the stack and redeem association pointers 
  FormFieldFloat64Pull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formfieldfloat64Service.getFormFieldFloat64s(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            formfieldfloat64s,
          ]) => {
            // init the array
            this.frontRepo.FormFieldFloat64s_array = formfieldfloat64s

            // clear the map that counts FormFieldFloat64 in the GET
            this.frontRepo.FormFieldFloat64s_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            formfieldfloat64s.forEach(
              formfieldfloat64 => {
                this.frontRepo.FormFieldFloat64s.set(formfieldfloat64.ID, formfieldfloat64)
                this.frontRepo.FormFieldFloat64s_batch.set(formfieldfloat64.ID, formfieldfloat64)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear formfieldfloat64s that are absent from the GET
            this.frontRepo.FormFieldFloat64s.forEach(
              formfieldfloat64 => {
                if (this.frontRepo.FormFieldFloat64s_batch.get(formfieldfloat64.ID) == undefined) {
                  this.frontRepo.FormFieldFloat64s.delete(formfieldfloat64.ID)
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

  // FormFieldIntPull performs a GET on FormFieldInt of the stack and redeem association pointers 
  FormFieldIntPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formfieldintService.getFormFieldInts(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            formfieldints,
          ]) => {
            // init the array
            this.frontRepo.FormFieldInts_array = formfieldints

            // clear the map that counts FormFieldInt in the GET
            this.frontRepo.FormFieldInts_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            formfieldints.forEach(
              formfieldint => {
                this.frontRepo.FormFieldInts.set(formfieldint.ID, formfieldint)
                this.frontRepo.FormFieldInts_batch.set(formfieldint.ID, formfieldint)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear formfieldints that are absent from the GET
            this.frontRepo.FormFieldInts.forEach(
              formfieldint => {
                if (this.frontRepo.FormFieldInts_batch.get(formfieldint.ID) == undefined) {
                  this.frontRepo.FormFieldInts.delete(formfieldint.ID)
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

  // FormFieldStringPull performs a GET on FormFieldString of the stack and redeem association pointers 
  FormFieldStringPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formfieldstringService.getFormFieldStrings(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            formfieldstrings,
          ]) => {
            // init the array
            this.frontRepo.FormFieldStrings_array = formfieldstrings

            // clear the map that counts FormFieldString in the GET
            this.frontRepo.FormFieldStrings_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            formfieldstrings.forEach(
              formfieldstring => {
                this.frontRepo.FormFieldStrings.set(formfieldstring.ID, formfieldstring)
                this.frontRepo.FormFieldStrings_batch.set(formfieldstring.ID, formfieldstring)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear formfieldstrings that are absent from the GET
            this.frontRepo.FormFieldStrings.forEach(
              formfieldstring => {
                if (this.frontRepo.FormFieldStrings_batch.get(formfieldstring.ID) == undefined) {
                  this.frontRepo.FormFieldStrings.delete(formfieldstring.ID)
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

  // FormFieldTimePull performs a GET on FormFieldTime of the stack and redeem association pointers 
  FormFieldTimePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formfieldtimeService.getFormFieldTimes(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            formfieldtimes,
          ]) => {
            // init the array
            this.frontRepo.FormFieldTimes_array = formfieldtimes

            // clear the map that counts FormFieldTime in the GET
            this.frontRepo.FormFieldTimes_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            formfieldtimes.forEach(
              formfieldtime => {
                this.frontRepo.FormFieldTimes.set(formfieldtime.ID, formfieldtime)
                this.frontRepo.FormFieldTimes_batch.set(formfieldtime.ID, formfieldtime)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear formfieldtimes that are absent from the GET
            this.frontRepo.FormFieldTimes.forEach(
              formfieldtime => {
                if (this.frontRepo.FormFieldTimes_batch.get(formfieldtime.ID) == undefined) {
                  this.frontRepo.FormFieldTimes.delete(formfieldtime.ID)
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

  // FormGroupPull performs a GET on FormGroup of the stack and redeem association pointers 
  FormGroupPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formgroupService.getFormGroups(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            formgroups,
          ]) => {
            // init the array
            this.frontRepo.FormGroups_array = formgroups

            // clear the map that counts FormGroup in the GET
            this.frontRepo.FormGroups_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            formgroups.forEach(
              formgroup => {
                this.frontRepo.FormGroups.set(formgroup.ID, formgroup)
                this.frontRepo.FormGroups_batch.set(formgroup.ID, formgroup)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear formgroups that are absent from the GET
            this.frontRepo.FormGroups.forEach(
              formgroup => {
                if (this.frontRepo.FormGroups_batch.get(formgroup.ID) == undefined) {
                  this.frontRepo.FormGroups.delete(formgroup.ID)
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
export function getFormDivUniqueID(id: number): number {
  return 61 * id
}
export function getFormFieldUniqueID(id: number): number {
  return 67 * id
}
export function getFormFieldBooleanUniqueID(id: number): number {
  return 71 * id
}
export function getFormFieldDateUniqueID(id: number): number {
  return 73 * id
}
export function getFormFieldFloat64UniqueID(id: number): number {
  return 79 * id
}
export function getFormFieldIntUniqueID(id: number): number {
  return 83 * id
}
export function getFormFieldStringUniqueID(id: number): number {
  return 89 * id
}
export function getFormFieldTimeUniqueID(id: number): number {
  return 97 * id
}
export function getFormGroupUniqueID(id: number): number {
  return 101 * id
}
export function getRowUniqueID(id: number): number {
  return 103 * id
}
export function getTableUniqueID(id: number): number {
  return 107 * id
}
