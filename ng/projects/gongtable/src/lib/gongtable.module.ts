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

import { CellBooleansTableComponent } from './cellbooleans-table/cellbooleans-table.component'
import { CellBooleanSortingComponent } from './cellboolean-sorting/cellboolean-sorting.component'
import { CellBooleanDetailComponent } from './cellboolean-detail/cellboolean-detail.component'

import { CellFloat64sTableComponent } from './cellfloat64s-table/cellfloat64s-table.component'
import { CellFloat64SortingComponent } from './cellfloat64-sorting/cellfloat64-sorting.component'
import { CellFloat64DetailComponent } from './cellfloat64-detail/cellfloat64-detail.component'

import { CellIconsTableComponent } from './cellicons-table/cellicons-table.component'
import { CellIconSortingComponent } from './cellicon-sorting/cellicon-sorting.component'
import { CellIconDetailComponent } from './cellicon-detail/cellicon-detail.component'

import { CellIntsTableComponent } from './cellints-table/cellints-table.component'
import { CellIntSortingComponent } from './cellint-sorting/cellint-sorting.component'
import { CellIntDetailComponent } from './cellint-detail/cellint-detail.component'

import { CellStringsTableComponent } from './cellstrings-table/cellstrings-table.component'
import { CellStringSortingComponent } from './cellstring-sorting/cellstring-sorting.component'
import { CellStringDetailComponent } from './cellstring-detail/cellstring-detail.component'

import { DisplayedColumnsTableComponent } from './displayedcolumns-table/displayedcolumns-table.component'
import { DisplayedColumnSortingComponent } from './displayedcolumn-sorting/displayedcolumn-sorting.component'
import { DisplayedColumnDetailComponent } from './displayedcolumn-detail/displayedcolumn-detail.component'

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

		CellBooleansTableComponent,
		CellBooleanSortingComponent,
		CellBooleanDetailComponent,

		CellFloat64sTableComponent,
		CellFloat64SortingComponent,
		CellFloat64DetailComponent,

		CellIconsTableComponent,
		CellIconSortingComponent,
		CellIconDetailComponent,

		CellIntsTableComponent,
		CellIntSortingComponent,
		CellIntDetailComponent,

		CellStringsTableComponent,
		CellStringSortingComponent,
		CellStringDetailComponent,

		DisplayedColumnsTableComponent,
		DisplayedColumnSortingComponent,
		DisplayedColumnDetailComponent,

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

		CellBooleansTableComponent,
		CellBooleanSortingComponent,
		CellBooleanDetailComponent,

		CellFloat64sTableComponent,
		CellFloat64SortingComponent,
		CellFloat64DetailComponent,

		CellIconsTableComponent,
		CellIconSortingComponent,
		CellIconDetailComponent,

		CellIntsTableComponent,
		CellIntSortingComponent,
		CellIntDetailComponent,

		CellStringsTableComponent,
		CellStringSortingComponent,
		CellStringDetailComponent,

		DisplayedColumnsTableComponent,
		DisplayedColumnSortingComponent,
		DisplayedColumnDetailComponent,

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
