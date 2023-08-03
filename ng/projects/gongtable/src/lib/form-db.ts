// insertion point for imports
import { FormCellDB } from './formcell-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormDB {

	static GONGSTRUCT_NAME = "Form"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	FormCells?: Array<FormCellDB>
}
