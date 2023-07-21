// insertion point sub template for components imports 
  import { TablesTableComponent } from './tables-table/tables-table.component'
  import { TableSortingComponent } from './table-sorting/table-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfTablesComponents: Map<string, any> = new Map([["TablesTableComponent", TablesTableComponent],])
  export const MapOfTableSortingComponents: Map<string, any> = new Map([["TableSortingComponent", TableSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Table", MapOfTablesComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Table", MapOfTableSortingComponents],
    ]
  )
