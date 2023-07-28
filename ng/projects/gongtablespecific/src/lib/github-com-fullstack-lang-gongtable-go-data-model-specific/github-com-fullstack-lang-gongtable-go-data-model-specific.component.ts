import { Component, Input, OnInit, ViewChild } from '@angular/core';
import { Subscription, debounceTime, distinctUntilChanged } from 'rxjs';

import * as gongtable from 'gongtable'

import { MatTableDataSource } from '@angular/material/table';
import { FormControl } from '@angular/forms';
import { MatSort } from '@angular/material/sort';

@Component({
  selector: 'lib-github-com-fullstack-lang-gongtable-go-data-model-specific',
  templateUrl: './github-com-fullstack-lang-gongtable-go-data-model-specific.component.html',
  styleUrls: ['./github-com-fullstack-lang-gongtable-go-data-model-specific.component.css']
})
export class GithubComFullstackLangGongtableGoDataModelSpecificComponent implements OnInit {
  displayedColumns: string[] = []
  mapHeaderIdIndex = new Map<string, number>()

  dataSource = new MatTableDataSource<gongtable.RowDB>()


  selectedTable: gongtable.TableDB | undefined = undefined

  @Input() DataStack: string = ""
  @Input() TableName: string = ""

  // for filtering
  filterControl = new FormControl()

  // for sorting
  @ViewChild(MatSort)
  sort: MatSort | undefined
  matSortDirective : string = ""

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

        this.dataSource = new MatTableDataSource(this.selectedTable.Rows!)

  
        // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)

        if (this.selectedTable.DisplayedColumns == undefined) {
          return
        }

        this.mapHeaderIdIndex = new Map<string, number>()
        let index = 0
        for (let column of this.selectedTable.DisplayedColumns) {
          this.mapHeaderIdIndex.set(column.Name, index)
          this.displayedColumns.push(column.Name)
          index++
        }

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

            let cell : gongtable.CellDB = rowDB.Cells[index]
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

      }
    )
  }

  ngAfterViewInit() {
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.dataSource.filter = filterValue.trim().toLowerCase();
  }
}
