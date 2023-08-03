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

  form = "Form"
  form_data_model = 'Form Data/Model'

  view = this.table

  views: string[] = [this.table, this.form, this.table_data_model, this.form_data_model];

  TableStack = "manualy edited table"
  FormStack = "form"
  TableGeneratedStack = "generated table"
  ModelStacks = "github.com/fullstack-lang/gongtable/go/models"

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
