import { Component, Input, OnInit, ViewChild } from '@angular/core';
import { Subscription, debounceTime, distinctUntilChanged, forkJoin } from 'rxjs';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import * as gongtable from 'gongtable'

import { MatTableDataSource } from '@angular/material/table';
import { FormControl } from '@angular/forms';
import { MatSort } from '@angular/material/sort';
import { MatPaginator } from '@angular/material/paginator';
import { SelectionModel } from '@angular/cdk/collections';
import { RowDB } from 'projects/gongtable/src/lib/row-db';

const allowMultiSelect = true

@Component({
  selector: 'lib-material-table',
  templateUrl: './material-table.component.html',
  styleUrls: ['./material-table.component.css']
})
export class MaterialTableComponent implements OnInit {

  displayedColumns: string[] = []
  allDisplayedColumns: string[] = [] // in case there is a checkbox

  mapHeaderIdIndex = new Map<string, number>()

  dataSource = new MatTableDataSource<gongtable.RowDB>()


  // for selection
  selectedTable: gongtable.TableDB | undefined = undefined;

  @Input() DataStack: string = ""
  @Input() TableName: string = ""

  // for filtering
  filterControl = new FormControl()

  // for sorting
  @ViewChild(MatSort)
  sort: MatSort | undefined
  matSortDirective: string = ""

  // for pagination
  @ViewChild(MatPaginator)
  paginator: MatPaginator | undefined;

  // the component is refreshed when modification are performed in the back repo 
  // 
  // the checkCommitNbFromBackTimer polls the commit number of the back repo
  // if the commit number has increased, it pulls the front repo and redraw the diagram
  private commutNbFromBackSubscription: Subscription = new Subscription
  lastCommitNbFromBack = -1
  lastPushFromFrontNb = -1
  currTime: number = 0
  dateOfLastTimerEmission: Date = new Date


  public gongtableFrontRepo?: gongtable.FrontRepo

  constructor(
    private gongtableFrontRepoService: gongtable.FrontRepoService,
    private gongtableCommitNbFromBackService: gongtable.CommitNbFromBackService,
    private rowService: gongtable.RowService,
  ) {

  }

  ngOnInit(): void {
    this.startAutoRefresh(500); // Refresh every 500 ms (half second)

    this.filterControl.valueChanges
      .pipe(
        debounceTime(200), // Optional. To reduce number of requests.
        distinctUntilChanged() // Optional. To prevent same filter fire multiple times.
      )
      .subscribe(value => {
        this.dataSource.filter = value;
      })

    this.dataSource.filterPredicate = (data: any, filter: string) => {
      const dataStr = JSON.stringify(data).toLowerCase(); // Convert the data to a lower case string.
      return dataStr.indexOf(filter) !== -1;
    }
  }

  ngOnDestroy(): void {
    this.stopAutoRefresh();
  }


  stopAutoRefresh(): void {
    if (this.commutNbFromBackSubscription) {
      this.commutNbFromBackSubscription.unsubscribe();
    }
  }

  startAutoRefresh(intervalMs: number): void {
    this.commutNbFromBackSubscription = this.gongtableCommitNbFromBackService
      .getCommitNbFromBack(intervalMs, this.DataStack)
      .subscribe((commitNbFromBack: number) => {
        // console.log("OutletComponent, last commit nb " + this.lastCommitNbFromBack + " new: " + commitNbFromBack)

        if (this.lastCommitNbFromBack < commitNbFromBack) {
          const d = new Date()
          console.log("OutletComponent:", this.DataStack, d.toLocaleTimeString() + `.${d.getMilliseconds()}` +
            ", last commit increased nb " + this.lastCommitNbFromBack + " new: " + commitNbFromBack)
          this.lastCommitNbFromBack = commitNbFromBack
          this.refresh()
        }
      }
      )
  }

  refresh(): void {

    this.gongtableFrontRepoService.pull(this.DataStack).subscribe(
      gongtablesFrontRepo => {
        this.gongtableFrontRepo = gongtablesFrontRepo

        this.selectedTable = undefined

        for (let table of this.gongtableFrontRepo.Tables_array) {
          if (table.Name == this.TableName) {
            this.selectedTable = table
          }
        }

        if (this.selectedTable == undefined) {
          return
        }

        // sort rows according to their index
        this.selectedTable.Rows?.sort((t1, t2) => {
          let t1_revPointerID_Index = t1.Table_RowsDBID_Index
          let t2_revPointerID_Index = t2.Table_RowsDBID_Index
          if (t1_revPointerID_Index && t2_revPointerID_Index) {
            if (t1_revPointerID_Index.Int64 > t2_revPointerID_Index.Int64) {
              return 1;
            }
            if (t1_revPointerID_Index.Int64 < t2_revPointerID_Index.Int64) {
              return -1;
            }
          }
          return 0;
        })

        this.dataSource = new MatTableDataSource(this.selectedTable.Rows!)

        if (this.selectedTable.HasCheckableRows) {

        }

        // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)

        if (this.selectedTable.DisplayedColumns == undefined) {
          return
        }

        this.mapHeaderIdIndex = new Map<string, number>()
        let index = 0

        this.displayedColumns = []
        for (let column of this.selectedTable.DisplayedColumns) {
          this.mapHeaderIdIndex.set(column.Name, index)
          this.displayedColumns.push(column.Name)
          index++
        }
        this.allDisplayedColumns = []
        if (this.selectedTable.HasCheckableRows) {
          this.allDisplayedColumns = ['select']

          if (this.selectedTable.Rows != undefined) {
            this.initialSelection = []
            for (let rowDB of this.selectedTable.Rows) {
              if (rowDB.IsChecked) {
                this.initialSelection.push(rowDB)
              }
            }
            this.selection = new SelectionModel<gongtable.RowDB>(allowMultiSelect, this.initialSelection)
          }

        }
        this.allDisplayedColumns = this.allDisplayedColumns.concat(this.displayedColumns)

        if (this.selectedTable.HasFiltering) {
          this.dataSource.filterPredicate = (rowDB: gongtable.RowDB, filter: string) => {

            // filtering is based on finding a lower case filter into a concatenated string
            // the cellDB properties
            let mergedContent = ""

            for (let cell of rowDB.Cells!) {
              if (cell.CellInt) {
                mergedContent += cell.CellInt.Value
              }
              if (cell.CellFloat64) {
                mergedContent += cell.CellFloat64.Value
              }
              if (cell.CellString) {
                mergedContent += cell.CellString.Value
              }
            }

            mergedContent = mergedContent.toLowerCase()
            let isSelected = mergedContent.includes(filter.toLowerCase())
            return isSelected
          }
        }

        this.matSortDirective = ""
        if (this.selectedTable.HasColumnSorting) {
          this.dataSource.sort = this.sort!
          this.matSortDirective = "mat-sort"

          // enable sorting on all fields (including pointers and reverse pointer)
          this.dataSource.sortingDataAccessor = (rowDB: gongtable.RowDB, sortHeaderId: string) => {

            if (rowDB.Cells == undefined) {
              return ""
            }
            let index = this.mapHeaderIdIndex.get(sortHeaderId)
            if (index == undefined) {
              return ""
            }

            let cell: gongtable.CellDB = rowDB.Cells[index]
            if (cell.CellInt) {
              return cell.CellInt.Value
            }
            if (cell.CellFloat64) {
              return cell.CellFloat64.Value
            }
            if (cell.CellString) {
              return cell.CellString.Value
            }
            if (cell.CellIcon) {
              return cell.CellIcon.Icon
            }
            if (cell.CellBool) {
              if (cell.CellBool.Value) {
                return "true"
              } else {
                return "false"
              }
            }

            return "";
          };
        }

        if (this.selectedTable.HasPaginator) {
          this.dataSource.paginator = this.paginator!
        }
      }
    )
  }

  ngAfterViewInit() {
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.dataSource.filter = filterValue.trim().toLowerCase();
  }

  selection: SelectionModel<gongtable.RowDB> = new (SelectionModel)
  initialSelection = new Array<gongtable.RowDB>()

  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.selectedTable?.Rows?.length
    return numSelected === numRows;
  }


  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.selectedTable?.Rows?.forEach(row => this.selection.select(row));
  }

  save() {

    this.selectedTable?.Rows?.forEach(row => row.IsChecked = false)

    for (let row of this.selection.selected) {
      row.IsChecked = true
    }

    const promises = []
    for (let row of this.selectedTable?.Rows!) {
      promises.push(this.rowService.updateRow(row, this.DataStack))
    }

    forkJoin(promises).subscribe(
      () => this.refresh()
    )
  }

  drop(event: CdkDragDrop<string[]>) {
    moveItemInArray(this.selectedTable?.Rows!, event.previousIndex, event.currentIndex)

    const promises = []
    let index = 0
    for (let row of this.selectedTable?.Rows!) {
      row.Table_RowsDBID_Index.Int64 = index++
      promises.push(this.rowService.updateRow(row, this.DataStack))
    }

    forkJoin(promises).subscribe(
      () => this.refresh()
    )
  }

  isDraggableRow = (index: number, item: gongtable.RowDB) => this.selectedTable?.CanDragDropRows
}