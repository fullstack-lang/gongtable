// generated by gong
import { Component, OnInit, AfterViewInit, ViewChild, Inject, Optional, Input } from '@angular/core';
import { BehaviorSubject } from 'rxjs'
import { MatSort } from '@angular/material/sort';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData, FrontRepoService, FrontRepo, SelectionMode } from '../front-repo.service'
import { NullInt64 } from '../null-int64'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { ActivatedRoute, Router, RouterState } from '@angular/router';
import { FormDB } from '../form-db'
import { FormService } from '../form.service'

// insertion point for additional imports

import { RouteService } from '../route-service';

// TableComponent is initilizaed from different routes
// TableComponentMode detail different cases 
enum TableComponentMode {
  DISPLAY_MODE,
  ONE_MANY_ASSOCIATION_MODE,
  MANY_MANY_ASSOCIATION_MODE,
}

// generated table component
@Component({
  selector: 'app-formstable',
  templateUrl: './forms-table.component.html',
  styleUrls: ['./forms-table.component.css'],
})
export class FormsTableComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of Form instances
  selection: SelectionModel<FormDB> = new (SelectionModel)
  initialSelection = new Array<FormDB>()

  // the data source for the table
  forms: FormDB[] = []
  matTableDataSource: MatTableDataSource<FormDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.forms
  frontRepo: FrontRepo = new (FrontRepo)

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  // for sorting & pagination
  @ViewChild(MatSort)
  sort: MatSort | undefined
  @ViewChild(MatPaginator)
  paginator: MatPaginator | undefined;

  ngAfterViewInit() {

    // enable sorting on all fields (including pointers and reverse pointer)
    this.matTableDataSource.sortingDataAccessor = (formDB: FormDB, property: string) => {
      switch (property) {
        case 'ID':
          return formDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return formDB.Name;

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (formDB: FormDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the formDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += formDB.Name.toLowerCase()

      let isSelected = mergedContent.includes(filter.toLowerCase())
      return isSelected
    };

    this.matTableDataSource.sort = this.sort!
    this.matTableDataSource.paginator = this.paginator!
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.matTableDataSource.filter = filterValue.trim().toLowerCase();
  }

  constructor(
    private formService: FormService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of form instances
    public dialogRef: MatDialogRef<FormsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
    private activatedRoute: ActivatedRoute,

    private routeService: RouteService,
  ) {

    // compute mode
    if (dialogData == undefined) {
      this.mode = TableComponentMode.DISPLAY_MODE
    } else {
      this.GONG__StackPath = dialogData.GONG__StackPath
      switch (dialogData.SelectionMode) {
        case SelectionMode.ONE_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.ONE_MANY_ASSOCIATION_MODE
          break
        case SelectionMode.MANY_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.MANY_MANY_ASSOCIATION_MODE
          break
        default:
      }
    }

    // observable for changes in structs
    this.formService.FormServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getForms()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Delete', // insertion point for columns to display
        "Name",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
      ]
      this.selection = new SelectionModel<FormDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    let stackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')
    if (stackPath != undefined) {
      this.GONG__StackPath = stackPath
    }

    this.getForms()

    this.matTableDataSource = new MatTableDataSource(this.forms)
  }

  getForms(): void {
    this.frontRepoService.pull(this.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.forms = this.frontRepo.Forms_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let form of this.forms) {
            let ID = this.dialogData.ID
            let revPointer = form[this.dialogData.ReversePointer as keyof FormDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(form)
            }
            this.selection = new SelectionModel<FormDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, FormDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          // we associates on sourceInstance of type SourceStruct with a MANY MANY associations to FormDB
          // the field name is sourceField
          let sourceFieldArray = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as FormDB[]
          if (sourceFieldArray != null) {
            for (let associationInstance of sourceFieldArray) {
              let form = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as FormDB
              this.initialSelection.push(form)
            }
          }

          this.selection = new SelectionModel<FormDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.forms
      }
    )
  }

  // newForm initiate a new form
  // create a new Form objet
  newForm() {
  }

  deleteForm(formID: number, form: FormDB) {
    // list of forms is truncated of form before the delete
    this.forms = this.forms.filter(h => h !== form);

    this.formService.deleteForm(formID, this.GONG__StackPath).subscribe(
      form => {
        this.formService.FormServiceChanged.next("delete")
      }
    );
  }

  editForm(formID: number, form: FormDB) {

  }

  // set editor outlet
  setEditorRouterOutlet(formID: number) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + "form" + "-detail"

    let outletConf: any = {}
    outletConf[outletName] = [fullPath, formID, this.GONG__StackPath]

    this.router.navigate([{ outlets: outletConf }])
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.forms.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.forms.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<FormDB>()

      // reset all initial selection of form that belong to form
      for (let form of this.initialSelection) {
        let index = form[this.dialogData.ReversePointer as keyof FormDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(form)

      }

      // from selection, set form that belong to form
      for (let form of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = form[this.dialogData.ReversePointer as keyof FormDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(form)
      }


      // update all form (only update selection & initial selection)
      for (let form of toUpdate) {
        this.formService.updateForm(form, this.GONG__StackPath)
          .subscribe(form => {
            this.formService.FormServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, FormDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedForm = new Set<number>()
      for (let form of this.initialSelection) {
        if (this.selection.selected.includes(form)) {
          // console.log("form " + form.Name + " is still selected")
        } else {
          console.log("form " + form.Name + " has been unselected")
          unselectedForm.add(form.ID)
          console.log("is unselected " + unselectedForm.has(form.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let form = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as FormDB
      if (unselectedForm.has(form.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<FormDB>) = new Array<FormDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          form => {
            if (!this.initialSelection.includes(form)) {
              // console.log("form " + form.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + form.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = form.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = form.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("form " + form.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<FormDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}