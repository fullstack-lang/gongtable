// generated by gong
import { Component, OnInit, Inject, Optional } from '@angular/core';
import { TypeofExpr } from '@angular/compiler';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { FormCellDB } from '../formcell-db'
import { FormCellService } from '../formcell.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { NullInt64 } from '../null-int64'

@Component({
  selector: 'lib-formcell-sorting',
  templateUrl: './formcell-sorting.component.html',
  styleUrls: ['./formcell-sorting.component.css']
})
export class FormCellSortingComponent implements OnInit {

  frontRepo: FrontRepo = new (FrontRepo)

  // array of FormCell instances that are in the association
  associatedFormCells = new Array<FormCellDB>();

  constructor(
    private formcellService: FormCellService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of formcell instances
    public dialogRef: MatDialogRef<FormCellSortingComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    this.getFormCells()
  }

  getFormCells(): void {
    this.frontRepoService.pull(this.dialogData.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        let index = 0
        for (let formcell of this.frontRepo.FormCells_array) {
          let ID = this.dialogData.ID
          let revPointerID = formcell[this.dialogData.ReversePointer as keyof FormCellDB] as unknown as NullInt64
          let revPointerID_Index = formcell[this.dialogData.ReversePointer + "_Index" as keyof FormCellDB] as unknown as NullInt64
          if (revPointerID.Int64 == ID) {
            if (revPointerID_Index == undefined) {
              revPointerID_Index = new NullInt64
              revPointerID_Index.Valid = true
              revPointerID_Index.Int64 = index++
            }
            this.associatedFormCells.push(formcell)
          }
        }

        // sort associated formcell according to order
        this.associatedFormCells.sort((t1, t2) => {
          let t1_revPointerID_Index = t1[this.dialogData.ReversePointer + "_Index" as keyof typeof t1] as unknown as NullInt64
          let t2_revPointerID_Index = t2[this.dialogData.ReversePointer + "_Index" as keyof typeof t2] as unknown as NullInt64
          if (t1_revPointerID_Index && t2_revPointerID_Index) {
            if (t1_revPointerID_Index.Int64 > t2_revPointerID_Index.Int64) {
              return 1;
            }
            if (t1_revPointerID_Index.Int64 < t2_revPointerID_Index.Int64) {
              return -1;
            }
          }
          return 0;
        });
      }
    )
  }

  drop(event: CdkDragDrop<string[]>) {
    moveItemInArray(this.associatedFormCells, event.previousIndex, event.currentIndex);

    // set the order of FormCell instances
    let index = 0

    for (let formcell of this.associatedFormCells) {
      let revPointerID_Index = formcell[this.dialogData.ReversePointer + "_Index" as keyof FormCellDB] as unknown as NullInt64
      revPointerID_Index.Valid = true
      revPointerID_Index.Int64 = index++
    }
  }

  save() {

    this.associatedFormCells.forEach(
      formcell => {
        this.formcellService.updateFormCell(formcell, this.dialogData.GONG__StackPath)
          .subscribe(formcell => {
            this.formcellService.FormCellServiceChanged.next("update")
          });
      }
    )

    this.dialogRef.close('Sorting of ' + this.dialogData.ReversePointer + ' done');
  }
}