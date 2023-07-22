import { NgModule } from '@angular/core';
import { GithubComFullstackLangGongtableGoDataModelSpecificComponent } from './github-com-fullstack-lang-gongtable-go-data-model-specific/github-com-fullstack-lang-gongtable-go-data-model-specific.component';

import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { MatTableModule } from '@angular/material/table';


@NgModule({
  declarations: [
    GithubComFullstackLangGongtableGoDataModelSpecificComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    MatTableModule,
  ],
  exports: [
    GithubComFullstackLangGongtableGoDataModelSpecificComponent
  ]
})
export class GongtablespecificModule { }
