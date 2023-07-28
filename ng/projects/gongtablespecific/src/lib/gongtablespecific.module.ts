import { NgModule } from '@angular/core';
import { GithubComFullstackLangGongtableGoDataModelSpecificComponent } from './github-com-fullstack-lang-gongtable-go-data-model-specific/github-com-fullstack-lang-gongtable-go-data-model-specific.component';

import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { MatTableModule } from '@angular/material/table';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatFormFieldModule } from '@angular/material/form-field';
import { ReactiveFormsModule } from '@angular/forms';
import { MatInputModule } from '@angular/material/input';
import { MatSortModule } from '@angular/material/sort';
import { MatPaginatorModule } from '@angular/material/paginator'


@NgModule({
  declarations: [
    GithubComFullstackLangGongtableGoDataModelSpecificComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    MatTableModule,
    MatCheckboxModule,
    MatButtonModule,
    MatIconModule,
    MatFormFieldModule,
    ReactiveFormsModule,
    MatInputModule,
    MatSortModule,
    MatPaginatorModule,
  ],
  exports: [
    GithubComFullstackLangGongtableGoDataModelSpecificComponent
  ]
})
export class GongtablespecificModule { }
