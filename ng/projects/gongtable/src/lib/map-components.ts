// insertion point sub template for components imports 
  import { CellsTableComponent } from './cells-table/cells-table.component'
  import { CellSortingComponent } from './cell-sorting/cell-sorting.component'
  import { CellFloat64sTableComponent } from './cellfloat64s-table/cellfloat64s-table.component'
  import { CellFloat64SortingComponent } from './cellfloat64-sorting/cellfloat64-sorting.component'
  import { CellIconsTableComponent } from './cellicons-table/cellicons-table.component'
  import { CellIconSortingComponent } from './cellicon-sorting/cellicon-sorting.component'
  import { CellIntsTableComponent } from './cellints-table/cellints-table.component'
  import { CellIntSortingComponent } from './cellint-sorting/cellint-sorting.component'
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
  export const MapOfCellFloat64sComponents: Map<string, any> = new Map([["CellFloat64sTableComponent", CellFloat64sTableComponent],])
  export const MapOfCellFloat64SortingComponents: Map<string, any> = new Map([["CellFloat64SortingComponent", CellFloat64SortingComponent],])
  export const MapOfCellIconsComponents: Map<string, any> = new Map([["CellIconsTableComponent", CellIconsTableComponent],])
  export const MapOfCellIconSortingComponents: Map<string, any> = new Map([["CellIconSortingComponent", CellIconSortingComponent],])
  export const MapOfCellIntsComponents: Map<string, any> = new Map([["CellIntsTableComponent", CellIntsTableComponent],])
  export const MapOfCellIntSortingComponents: Map<string, any> = new Map([["CellIntSortingComponent", CellIntSortingComponent],])
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
      ["CellFloat64", MapOfCellFloat64sComponents],
      ["CellIcon", MapOfCellIconsComponents],
      ["CellInt", MapOfCellIntsComponents],
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
      ["CellFloat64", MapOfCellFloat64SortingComponents],
      ["CellIcon", MapOfCellIconSortingComponents],
      ["CellInt", MapOfCellIntSortingComponents],
      ["CellString", MapOfCellStringSortingComponents],
      ["DisplayedColumn", MapOfDisplayedColumnSortingComponents],
      ["Row", MapOfRowSortingComponents],
      ["Table", MapOfTableSortingComponents],
    ]
  )
