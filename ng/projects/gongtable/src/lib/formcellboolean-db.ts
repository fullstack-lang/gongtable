// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormCellBooleanDB {

	static GONGSTRUCT_NAME = "FormCellBoolean"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: boolean = false

	// insertion point for other declarations
}
