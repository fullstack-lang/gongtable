import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import { MatRadioModule } from '@angular/material/radio';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';

import { AngularSplitModule } from 'angular-split';

import * as gongtable from '../../projects/gongtable/src/public-api'

import { MaterialTableComponent } from '../../projects/gongtablespecific/src/lib/material-table/material-table.component';
import { MaterialFormComponent } from '../../projects/gongtablespecific/src/lib/material-form/material-form.component';

// import { TreeComponent } from '@vendored_components/github.com/fullstack-lang/gongtree/ng-github.com-fullstack-lang-gongtree/projects/gongtreespecific/src/public-api'
// import { PanelComponent } from '@vendored_components/github.com/fullstack-lang/gongdoc/ng-github.com-fullstack-lang-gongdoc/projects/gongdocspecific/src/public-api'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  standalone: true,
  imports: [
    MatRadioModule,
    FormsModule,
    CommonModule,
    MatButtonModule,
    MatIconModule,

    AngularSplitModule,
    MaterialTableComponent,
    MaterialFormComponent,

    // TreeComponent,
    // PanelComponent,
  ],
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