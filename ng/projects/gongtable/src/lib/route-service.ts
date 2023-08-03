import { Injectable } from '@angular/core';
import { Route, Router, Routes } from '@angular/router';

// insertion point for imports
import { CellsTableComponent } from './cells-table/cells-table.component'
import { CellDetailComponent } from './cell-detail/cell-detail.component'

import { CellBooleansTableComponent } from './cellbooleans-table/cellbooleans-table.component'
import { CellBooleanDetailComponent } from './cellboolean-detail/cellboolean-detail.component'

import { CellFloat64sTableComponent } from './cellfloat64s-table/cellfloat64s-table.component'
import { CellFloat64DetailComponent } from './cellfloat64-detail/cellfloat64-detail.component'

import { CellIconsTableComponent } from './cellicons-table/cellicons-table.component'
import { CellIconDetailComponent } from './cellicon-detail/cellicon-detail.component'

import { CellIntsTableComponent } from './cellints-table/cellints-table.component'
import { CellIntDetailComponent } from './cellint-detail/cellint-detail.component'

import { CellStringsTableComponent } from './cellstrings-table/cellstrings-table.component'
import { CellStringDetailComponent } from './cellstring-detail/cellstring-detail.component'

import { DisplayedColumnsTableComponent } from './displayedcolumns-table/displayedcolumns-table.component'
import { DisplayedColumnDetailComponent } from './displayedcolumn-detail/displayedcolumn-detail.component'

import { FormsTableComponent } from './forms-table/forms-table.component'
import { FormDetailComponent } from './form-detail/form-detail.component'

import { FormCellsTableComponent } from './formcells-table/formcells-table.component'
import { FormCellDetailComponent } from './formcell-detail/formcell-detail.component'

import { FormCellBooleansTableComponent } from './formcellbooleans-table/formcellbooleans-table.component'
import { FormCellBooleanDetailComponent } from './formcellboolean-detail/formcellboolean-detail.component'

import { FormCellFloat64sTableComponent } from './formcellfloat64s-table/formcellfloat64s-table.component'
import { FormCellFloat64DetailComponent } from './formcellfloat64-detail/formcellfloat64-detail.component'

import { FormCellIntsTableComponent } from './formcellints-table/formcellints-table.component'
import { FormCellIntDetailComponent } from './formcellint-detail/formcellint-detail.component'

import { FormCellStringsTableComponent } from './formcellstrings-table/formcellstrings-table.component'
import { FormCellStringDetailComponent } from './formcellstring-detail/formcellstring-detail.component'

import { RowsTableComponent } from './rows-table/rows-table.component'
import { RowDetailComponent } from './row-detail/row-detail.component'

import { TablesTableComponent } from './tables-table/tables-table.component'
import { TableDetailComponent } from './table-detail/table-detail.component'


@Injectable({
    providedIn: 'root'
})
export class RouteService {
    private routes: Routes = []

    constructor(private router: Router) { }

    public addRoutes(newRoutes: Routes): void {
        const existingRoutes = this.router.config
        this.routes = this.router.config

        for (let newRoute of newRoutes) {
            if (!existingRoutes.includes(newRoute)) {
                this.routes.push(newRoute)
            }
        }
        this.router.resetConfig(this.routes)
    }

    getPathRoot(): string {
        return 'github_com_fullstack_lang_gongtable_go'
    }
    getTableOutlet(stackPath: string): string {
        return this.getPathRoot() + '_table' + '_' + stackPath
    }
    getEditorOutlet(stackPath: string): string {
        return this.getPathRoot() + '_editor' + '_' + stackPath
    }
    // insertion point for per gongstruct route/path getters
    getCellTablePath(): string {
        return this.getPathRoot() + '-cells/:GONG__StackPath'
    }
    getCellTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCellTablePath(), component: CellsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getCellAdderPath(): string {
        return this.getPathRoot() + '-cell-adder/:GONG__StackPath'
    }
    getCellAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCellAdderPath(), component: CellDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getCellAdderForUsePath(): string {
        return this.getPathRoot() + '-cell-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getCellAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCellAdderForUsePath(), component: CellDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getCellDetailPath(): string {
        return this.getPathRoot() + '-cell-detail/:id/:GONG__StackPath'
    }
    getCellDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCellDetailPath(), component: CellDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getCellBooleanTablePath(): string {
        return this.getPathRoot() + '-cellbooleans/:GONG__StackPath'
    }
    getCellBooleanTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCellBooleanTablePath(), component: CellBooleansTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getCellBooleanAdderPath(): string {
        return this.getPathRoot() + '-cellboolean-adder/:GONG__StackPath'
    }
    getCellBooleanAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCellBooleanAdderPath(), component: CellBooleanDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getCellBooleanAdderForUsePath(): string {
        return this.getPathRoot() + '-cellboolean-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getCellBooleanAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCellBooleanAdderForUsePath(), component: CellBooleanDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getCellBooleanDetailPath(): string {
        return this.getPathRoot() + '-cellboolean-detail/:id/:GONG__StackPath'
    }
    getCellBooleanDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCellBooleanDetailPath(), component: CellBooleanDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getCellFloat64TablePath(): string {
        return this.getPathRoot() + '-cellfloat64s/:GONG__StackPath'
    }
    getCellFloat64TableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCellFloat64TablePath(), component: CellFloat64sTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getCellFloat64AdderPath(): string {
        return this.getPathRoot() + '-cellfloat64-adder/:GONG__StackPath'
    }
    getCellFloat64AdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCellFloat64AdderPath(), component: CellFloat64DetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getCellFloat64AdderForUsePath(): string {
        return this.getPathRoot() + '-cellfloat64-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getCellFloat64AdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCellFloat64AdderForUsePath(), component: CellFloat64DetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getCellFloat64DetailPath(): string {
        return this.getPathRoot() + '-cellfloat64-detail/:id/:GONG__StackPath'
    }
    getCellFloat64DetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCellFloat64DetailPath(), component: CellFloat64DetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getCellIconTablePath(): string {
        return this.getPathRoot() + '-cellicons/:GONG__StackPath'
    }
    getCellIconTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCellIconTablePath(), component: CellIconsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getCellIconAdderPath(): string {
        return this.getPathRoot() + '-cellicon-adder/:GONG__StackPath'
    }
    getCellIconAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCellIconAdderPath(), component: CellIconDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getCellIconAdderForUsePath(): string {
        return this.getPathRoot() + '-cellicon-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getCellIconAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCellIconAdderForUsePath(), component: CellIconDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getCellIconDetailPath(): string {
        return this.getPathRoot() + '-cellicon-detail/:id/:GONG__StackPath'
    }
    getCellIconDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCellIconDetailPath(), component: CellIconDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getCellIntTablePath(): string {
        return this.getPathRoot() + '-cellints/:GONG__StackPath'
    }
    getCellIntTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCellIntTablePath(), component: CellIntsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getCellIntAdderPath(): string {
        return this.getPathRoot() + '-cellint-adder/:GONG__StackPath'
    }
    getCellIntAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCellIntAdderPath(), component: CellIntDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getCellIntAdderForUsePath(): string {
        return this.getPathRoot() + '-cellint-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getCellIntAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCellIntAdderForUsePath(), component: CellIntDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getCellIntDetailPath(): string {
        return this.getPathRoot() + '-cellint-detail/:id/:GONG__StackPath'
    }
    getCellIntDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCellIntDetailPath(), component: CellIntDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getCellStringTablePath(): string {
        return this.getPathRoot() + '-cellstrings/:GONG__StackPath'
    }
    getCellStringTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCellStringTablePath(), component: CellStringsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getCellStringAdderPath(): string {
        return this.getPathRoot() + '-cellstring-adder/:GONG__StackPath'
    }
    getCellStringAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCellStringAdderPath(), component: CellStringDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getCellStringAdderForUsePath(): string {
        return this.getPathRoot() + '-cellstring-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getCellStringAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCellStringAdderForUsePath(), component: CellStringDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getCellStringDetailPath(): string {
        return this.getPathRoot() + '-cellstring-detail/:id/:GONG__StackPath'
    }
    getCellStringDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCellStringDetailPath(), component: CellStringDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getDisplayedColumnTablePath(): string {
        return this.getPathRoot() + '-displayedcolumns/:GONG__StackPath'
    }
    getDisplayedColumnTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDisplayedColumnTablePath(), component: DisplayedColumnsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getDisplayedColumnAdderPath(): string {
        return this.getPathRoot() + '-displayedcolumn-adder/:GONG__StackPath'
    }
    getDisplayedColumnAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDisplayedColumnAdderPath(), component: DisplayedColumnDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getDisplayedColumnAdderForUsePath(): string {
        return this.getPathRoot() + '-displayedcolumn-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getDisplayedColumnAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDisplayedColumnAdderForUsePath(), component: DisplayedColumnDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getDisplayedColumnDetailPath(): string {
        return this.getPathRoot() + '-displayedcolumn-detail/:id/:GONG__StackPath'
    }
    getDisplayedColumnDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDisplayedColumnDetailPath(), component: DisplayedColumnDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getFormTablePath(): string {
        return this.getPathRoot() + '-forms/:GONG__StackPath'
    }
    getFormTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormTablePath(), component: FormsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getFormAdderPath(): string {
        return this.getPathRoot() + '-form-adder/:GONG__StackPath'
    }
    getFormAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormAdderPath(), component: FormDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormAdderForUsePath(): string {
        return this.getPathRoot() + '-form-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getFormAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormAdderForUsePath(), component: FormDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormDetailPath(): string {
        return this.getPathRoot() + '-form-detail/:id/:GONG__StackPath'
    }
    getFormDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormDetailPath(), component: FormDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getFormCellTablePath(): string {
        return this.getPathRoot() + '-formcells/:GONG__StackPath'
    }
    getFormCellTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormCellTablePath(), component: FormCellsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getFormCellAdderPath(): string {
        return this.getPathRoot() + '-formcell-adder/:GONG__StackPath'
    }
    getFormCellAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormCellAdderPath(), component: FormCellDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormCellAdderForUsePath(): string {
        return this.getPathRoot() + '-formcell-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getFormCellAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormCellAdderForUsePath(), component: FormCellDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormCellDetailPath(): string {
        return this.getPathRoot() + '-formcell-detail/:id/:GONG__StackPath'
    }
    getFormCellDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormCellDetailPath(), component: FormCellDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getFormCellBooleanTablePath(): string {
        return this.getPathRoot() + '-formcellbooleans/:GONG__StackPath'
    }
    getFormCellBooleanTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormCellBooleanTablePath(), component: FormCellBooleansTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getFormCellBooleanAdderPath(): string {
        return this.getPathRoot() + '-formcellboolean-adder/:GONG__StackPath'
    }
    getFormCellBooleanAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormCellBooleanAdderPath(), component: FormCellBooleanDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormCellBooleanAdderForUsePath(): string {
        return this.getPathRoot() + '-formcellboolean-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getFormCellBooleanAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormCellBooleanAdderForUsePath(), component: FormCellBooleanDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormCellBooleanDetailPath(): string {
        return this.getPathRoot() + '-formcellboolean-detail/:id/:GONG__StackPath'
    }
    getFormCellBooleanDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormCellBooleanDetailPath(), component: FormCellBooleanDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getFormCellFloat64TablePath(): string {
        return this.getPathRoot() + '-formcellfloat64s/:GONG__StackPath'
    }
    getFormCellFloat64TableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormCellFloat64TablePath(), component: FormCellFloat64sTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getFormCellFloat64AdderPath(): string {
        return this.getPathRoot() + '-formcellfloat64-adder/:GONG__StackPath'
    }
    getFormCellFloat64AdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormCellFloat64AdderPath(), component: FormCellFloat64DetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormCellFloat64AdderForUsePath(): string {
        return this.getPathRoot() + '-formcellfloat64-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getFormCellFloat64AdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormCellFloat64AdderForUsePath(), component: FormCellFloat64DetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormCellFloat64DetailPath(): string {
        return this.getPathRoot() + '-formcellfloat64-detail/:id/:GONG__StackPath'
    }
    getFormCellFloat64DetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormCellFloat64DetailPath(), component: FormCellFloat64DetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getFormCellIntTablePath(): string {
        return this.getPathRoot() + '-formcellints/:GONG__StackPath'
    }
    getFormCellIntTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormCellIntTablePath(), component: FormCellIntsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getFormCellIntAdderPath(): string {
        return this.getPathRoot() + '-formcellint-adder/:GONG__StackPath'
    }
    getFormCellIntAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormCellIntAdderPath(), component: FormCellIntDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormCellIntAdderForUsePath(): string {
        return this.getPathRoot() + '-formcellint-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getFormCellIntAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormCellIntAdderForUsePath(), component: FormCellIntDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormCellIntDetailPath(): string {
        return this.getPathRoot() + '-formcellint-detail/:id/:GONG__StackPath'
    }
    getFormCellIntDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormCellIntDetailPath(), component: FormCellIntDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getFormCellStringTablePath(): string {
        return this.getPathRoot() + '-formcellstrings/:GONG__StackPath'
    }
    getFormCellStringTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormCellStringTablePath(), component: FormCellStringsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getFormCellStringAdderPath(): string {
        return this.getPathRoot() + '-formcellstring-adder/:GONG__StackPath'
    }
    getFormCellStringAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormCellStringAdderPath(), component: FormCellStringDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormCellStringAdderForUsePath(): string {
        return this.getPathRoot() + '-formcellstring-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getFormCellStringAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormCellStringAdderForUsePath(), component: FormCellStringDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormCellStringDetailPath(): string {
        return this.getPathRoot() + '-formcellstring-detail/:id/:GONG__StackPath'
    }
    getFormCellStringDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormCellStringDetailPath(), component: FormCellStringDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getRowTablePath(): string {
        return this.getPathRoot() + '-rows/:GONG__StackPath'
    }
    getRowTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getRowTablePath(), component: RowsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getRowAdderPath(): string {
        return this.getPathRoot() + '-row-adder/:GONG__StackPath'
    }
    getRowAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getRowAdderPath(), component: RowDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getRowAdderForUsePath(): string {
        return this.getPathRoot() + '-row-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getRowAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getRowAdderForUsePath(), component: RowDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getRowDetailPath(): string {
        return this.getPathRoot() + '-row-detail/:id/:GONG__StackPath'
    }
    getRowDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getRowDetailPath(), component: RowDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getTableTablePath(): string {
        return this.getPathRoot() + '-tables/:GONG__StackPath'
    }
    getTableTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTableTablePath(), component: TablesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getTableAdderPath(): string {
        return this.getPathRoot() + '-table-adder/:GONG__StackPath'
    }
    getTableAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTableAdderPath(), component: TableDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getTableAdderForUsePath(): string {
        return this.getPathRoot() + '-table-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getTableAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTableAdderForUsePath(), component: TableDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getTableDetailPath(): string {
        return this.getPathRoot() + '-table-detail/:id/:GONG__StackPath'
    }
    getTableDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTableDetailPath(), component: TableDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }



    addDataPanelRoutes(stackPath: string) {

        this.addRoutes([
            // insertion point for all routes getter
            this.getCellTableRoute(stackPath),
            this.getCellAdderRoute(stackPath),
            this.getCellAdderForUseRoute(stackPath),
            this.getCellDetailRoute(stackPath),

            this.getCellBooleanTableRoute(stackPath),
            this.getCellBooleanAdderRoute(stackPath),
            this.getCellBooleanAdderForUseRoute(stackPath),
            this.getCellBooleanDetailRoute(stackPath),

            this.getCellFloat64TableRoute(stackPath),
            this.getCellFloat64AdderRoute(stackPath),
            this.getCellFloat64AdderForUseRoute(stackPath),
            this.getCellFloat64DetailRoute(stackPath),

            this.getCellIconTableRoute(stackPath),
            this.getCellIconAdderRoute(stackPath),
            this.getCellIconAdderForUseRoute(stackPath),
            this.getCellIconDetailRoute(stackPath),

            this.getCellIntTableRoute(stackPath),
            this.getCellIntAdderRoute(stackPath),
            this.getCellIntAdderForUseRoute(stackPath),
            this.getCellIntDetailRoute(stackPath),

            this.getCellStringTableRoute(stackPath),
            this.getCellStringAdderRoute(stackPath),
            this.getCellStringAdderForUseRoute(stackPath),
            this.getCellStringDetailRoute(stackPath),

            this.getDisplayedColumnTableRoute(stackPath),
            this.getDisplayedColumnAdderRoute(stackPath),
            this.getDisplayedColumnAdderForUseRoute(stackPath),
            this.getDisplayedColumnDetailRoute(stackPath),

            this.getFormTableRoute(stackPath),
            this.getFormAdderRoute(stackPath),
            this.getFormAdderForUseRoute(stackPath),
            this.getFormDetailRoute(stackPath),

            this.getFormCellTableRoute(stackPath),
            this.getFormCellAdderRoute(stackPath),
            this.getFormCellAdderForUseRoute(stackPath),
            this.getFormCellDetailRoute(stackPath),

            this.getFormCellBooleanTableRoute(stackPath),
            this.getFormCellBooleanAdderRoute(stackPath),
            this.getFormCellBooleanAdderForUseRoute(stackPath),
            this.getFormCellBooleanDetailRoute(stackPath),

            this.getFormCellFloat64TableRoute(stackPath),
            this.getFormCellFloat64AdderRoute(stackPath),
            this.getFormCellFloat64AdderForUseRoute(stackPath),
            this.getFormCellFloat64DetailRoute(stackPath),

            this.getFormCellIntTableRoute(stackPath),
            this.getFormCellIntAdderRoute(stackPath),
            this.getFormCellIntAdderForUseRoute(stackPath),
            this.getFormCellIntDetailRoute(stackPath),

            this.getFormCellStringTableRoute(stackPath),
            this.getFormCellStringAdderRoute(stackPath),
            this.getFormCellStringAdderForUseRoute(stackPath),
            this.getFormCellStringDetailRoute(stackPath),

            this.getRowTableRoute(stackPath),
            this.getRowAdderRoute(stackPath),
            this.getRowAdderForUseRoute(stackPath),
            this.getRowDetailRoute(stackPath),

            this.getTableTableRoute(stackPath),
            this.getTableAdderRoute(stackPath),
            this.getTableAdderForUseRoute(stackPath),
            this.getTableDetailRoute(stackPath),

        ])
    }
}
