import { Injectable } from '@angular/core';
import { Route, Router, Routes } from '@angular/router';

// insertion point for imports
import { CellsTableComponent } from './cells-table/cells-table.component'
import { CellDetailComponent } from './cell-detail/cell-detail.component'

import { CellStringsTableComponent } from './cellstrings-table/cellstrings-table.component'
import { CellStringDetailComponent } from './cellstring-detail/cellstring-detail.component'

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

            this.getCellStringTableRoute(stackPath),
            this.getCellStringAdderRoute(stackPath),
            this.getCellStringAdderForUseRoute(stackPath),
            this.getCellStringDetailRoute(stackPath),

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
