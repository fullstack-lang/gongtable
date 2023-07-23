import { Component, Input, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';

import * as gongtable from 'gongtable'




@Component({
  selector: 'lib-github-com-fullstack-lang-gongtable-go-data-model-specific',
  templateUrl: './github-com-fullstack-lang-gongtable-go-data-model-specific.component.html',
  styleUrls: ['./github-com-fullstack-lang-gongtable-go-data-model-specific.component.css']
})
export class GithubComFullstackLangGongtableGoDataModelSpecificComponent implements OnInit {
  displayedColumns: string[] = [];
  dataSource: string[][] = [[]];

  @Input() DataStack: string = ""

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

        this.dataSource = []

        for (let column of this.gongtableFrontRepo.DisplayedColumns_array) {
          this.displayedColumns.push(column.Name)
        }

        for (let row of this.gongtableFrontRepo.Rows_array) {
          let rowOfStrings: string[] = []

          if (row.Cells == undefined) {
            return
          }
          for (let cell of row.Cells) {
            rowOfStrings.push(cell.Name)
          }

          this.dataSource.push(rowOfStrings)
        }

      }
    )
  }
}
