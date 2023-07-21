// insertion point for imports
import { CellStringDB } from './cellstring-db'
import { RowDB } from './row-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CellDB {

	static GONGSTRUCT_NAME = "Cell"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	CellString?: CellStringDB
	CellStringID: NullInt64 = new NullInt64 // if pointer is null, CellString.ID = 0

	Row_CellsDBID: NullInt64 = new NullInt64
	Row_CellsDBID_Index: NullInt64  = new NullInt64 // store the index of the cell instance in Row.Cells
	Row_Cells_reverse?: RowDB 

}
