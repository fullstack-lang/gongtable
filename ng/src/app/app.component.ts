import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
import * as gongtable from 'gongtable'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  table = 'Table'
  table_data_model = 'Table Data/Model'

  form_view = "Form"
  form_data_model = 'Form Data/Model'

  view = this.table

  views: string[] = [this.table, this.form_view, this.table_data_model, this.form_data_model];

  TableStack = "manualy filled table"

  FormStack = "manualy filled form"
  FormName = "Form 1"

  TableGeneratedStack = "generated table"
  ModelStacks = "github.com/fullstack-lang/gongtable/go/models"

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
