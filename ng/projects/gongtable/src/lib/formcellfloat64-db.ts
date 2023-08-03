// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormCellFloat64DB {

	static GONGSTRUCT_NAME = "FormCellFloat64"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: number = 0

	// insertion point for other declarations
}
