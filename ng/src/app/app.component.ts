import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
import * as gongtable from 'gongtable'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  default = 'Gongtable Data/Model'
  specific = 'Gongtable'
  view = this.specific

  views: string[] = [this.specific, this.default];

  DataStack = "gongtable"
  GeneratedStack = "generated"
  ModelStacks = "github.com/fullstack-lang/gongtable/go/models"

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
