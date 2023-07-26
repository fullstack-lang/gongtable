import { Injectable } from '@angular/core';
import { Route, Router, Routes } from '@angular/router';

// insertion point for imports
import { CellsTableComponent } from './cells-table/cells-table.component'
import { CellDetailComponent } from './cell-detail/cell-detail.component'

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
