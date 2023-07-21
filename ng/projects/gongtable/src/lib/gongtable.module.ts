import { NgModule } from '@angular/core';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Routes, RouterModule } from '@angular/router';

// for angular material
import { MatSliderModule } from '@angular/material/slider';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatTableModule } from '@angular/material/table'
import { MatSortModule } from '@angular/material/sort'
import { MatPaginatorModule } from '@angular/material/paginator'
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar'
import { MatListModule } from '@angular/material/list'
import { MatExpansionModule } from '@angular/material/expansion';
import { MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatTreeModule } from '@angular/material/tree';
import { DragDropModule } from '@angular/cdk/drag-drop';

import { AngularSplitModule, SplitComponent } from 'angular-split';

import {
	NgxMatDatetimePickerModule,
	NgxMatNativeDateModule,
	NgxMatTimepickerModule
} from '@angular-material-components/datetime-picker';

import { AppRoutingModule } from './app-routing.module';

import { SplitterComponent } from './splitter/splitter.component'
import { SidebarComponent } from './sidebar/sidebar.component';
import { GongstructSelectionService } from './gongstruct-selection.service'

// insertion point for imports 
import { CellsTableComponent } from './cells-table/cells-table.component'
import { CellSortingComponent } from './cell-sorting/cell-sorting.component'
import { CellDetailComponent } from './cell-detail/cell-detail.component'

import { CellStringsTableComponent } from './cellstrings-table/cellstrings-table.component'
import { CellStringSortingComponent } from './cellstring-sorting/cellstring-sorting.component'
import { CellStringDetailComponent } from './cellstring-detail/cellstring-detail.component'

import { RowsTableComponent } from './rows-table/rows-table.component'
import { RowSortingComponent } from './row-sorting/row-sorting.component'
import { RowDetailComponent } from './row-detail/row-detail.component'

import { TablesTableComponent } from './tables-table/tables-table.component'
import { TableSortingComponent } from './table-sorting/table-sorting.component'
import { TableDetailComponent } from './table-detail/table-detail.component'


@NgModule({
	declarations: [
		// insertion point for declarations 
		CellsTableComponent,
		CellSortingComponent,
		CellDetailComponent,

		CellStringsTableComponent,
		CellStringSortingComponent,
		CellStringDetailComponent,

		RowsTableComponent,
		RowSortingComponent,
		RowDetailComponent,

		TablesTableComponent,
		TableSortingComponent,
		TableDetailComponent,


		SplitterComponent,
		SidebarComponent
	],
	imports: [
		FormsModule,
		ReactiveFormsModule,
		CommonModule,
		RouterModule,

		AppRoutingModule,

		MatSliderModule,
		MatSelectModule,
		MatFormFieldModule,
		MatInputModule,
		MatDatepickerModule,
		MatTableModule,
		MatSortModule,
		MatPaginatorModule,
		MatCheckboxModule,
		MatButtonModule,
		MatIconModule,
		MatToolbarModule,
		MatExpansionModule,
		MatListModule,
		MatDialogModule,
		MatGridListModule,
		MatTreeModule,
		DragDropModule,

		NgxMatDatetimePickerModule,
		NgxMatNativeDateModule,
		NgxMatTimepickerModule,

		AngularSplitModule,
	],
	exports: [
		// insertion point for declarations 
		CellsTableComponent,
		CellSortingComponent,
		CellDetailComponent,

		CellStringsTableComponent,
		CellStringSortingComponent,
		CellStringDetailComponent,

		RowsTableComponent,
		RowSortingComponent,
		RowDetailComponent,

		TablesTableComponent,
		TableSortingComponent,
		TableDetailComponent,


		SplitterComponent,
		SidebarComponent,

	],
	providers: [
		GongstructSelectionService,
		{
			provide: MatDialogRef,
			useValue: {}
		},
	],
})
export class GongtableModule { }
