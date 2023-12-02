import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
import * as gongtable from 'gongtable'

import { GongdocModule } from 'gongdoc'
import { GongdocspecificModule } from 'gongdocspecific'

import { GongtreeModule } from 'gongtree'
import { GongtreespecificModule } from 'gongtreespecific'

import { GongtableModule } from 'gongtable'
import { GongtablespecificModule } from 'gongtablespecific'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  table_view = gongtable.TableName.TableDefaultName.toString()
  manualy_edited_table_probe = 'Manual Edited Table Stack Probe'

  form_view = gongtable.FormGroupName.FormGroupDefaultName.toString()
  manualy_edited_form_probe = 'Manual Edited Form Probe'

  generated_table_probe_stack = 'Generated Table Stack Probe'

  view = this.form_view

  TableTestNameEnum = gongtable.TableTestNameEnum

  views: string[] = [
    this.table_view,
    this.form_view,
    this.manualy_edited_table_probe,
    this.manualy_edited_form_probe,
    this.generated_table_probe_stack];

  FormName = "Form 1"

  scrollStyle = {
    'overflow- x': 'auto',
    'width': '100%',  // Ensure the div takes the full width of its parent container
  }

  StackType = "github.com/fullstack-lang/gongtable/go/models"

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
