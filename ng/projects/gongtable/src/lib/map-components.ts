// insertion point sub template for components imports 
  import { CellsTableComponent } from './cells-table/cells-table.component'
  import { CellSortingComponent } from './cell-sorting/cell-sorting.component'
  import { CellStringsTableComponent } from './cellstrings-table/cellstrings-table.component'
  import { CellStringSortingComponent } from './cellstring-sorting/cellstring-sorting.component'
  import { DisplayedColumnsTableComponent } from './displayedcolumns-table/displayedcolumns-table.component'
  import { DisplayedColumnSortingComponent } from './displayedcolumn-sorting/displayedcolumn-sorting.component'
  import { RowsTableComponent } from './rows-table/rows-table.component'
  import { RowSortingComponent } from './row-sorting/row-sorting.component'
  import { TablesTableComponent } from './tables-table/tables-table.component'
  import { TableSortingComponent } from './table-sorting/table-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfCellsComponents: Map<string, any> = new Map([["CellsTableComponent", CellsTableComponent],])
  export const MapOfCellSortingComponents: Map<string, any> = new Map([["CellSortingComponent", CellSortingComponent],])
  export const MapOfCellStringsComponents: Map<string, any> = new Map([["CellStringsTableComponent", CellStringsTableComponent],])
  export const MapOfCellStringSortingComponents: Map<string, any> = new Map([["CellStringSortingComponent", CellStringSortingComponent],])
  export const MapOfDisplayedColumnsComponents: Map<string, any> = new Map([["DisplayedColumnsTableComponent", DisplayedColumnsTableComponent],])
  export const MapOfDisplayedColumnSortingComponents: Map<string, any> = new Map([["DisplayedColumnSortingComponent", DisplayedColumnSortingComponent],])
  export const MapOfRowsComponents: Map<string, any> = new Map([["RowsTableComponent", RowsTableComponent],])
  export const MapOfRowSortingComponents: Map<string, any> = new Map([["RowSortingComponent", RowSortingComponent],])
  export const MapOfTablesComponents: Map<string, any> = new Map([["TablesTableComponent", TablesTableComponent],])
  export const MapOfTableSortingComponents: Map<string, any> = new Map([["TableSortingComponent", TableSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Cell", MapOfCellsComponents],
      ["CellString", MapOfCellStringsComponents],
      ["DisplayedColumn", MapOfDisplayedColumnsComponents],
      ["Row", MapOfRowsComponents],
      ["Table", MapOfTablesComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Cell", MapOfCellSortingComponents],
      ["CellString", MapOfCellStringSortingComponents],
      ["DisplayedColumn", MapOfDisplayedColumnSortingComponents],
      ["Row", MapOfRowSortingComponents],
      ["Table", MapOfTableSortingComponents],
    ]
  )
