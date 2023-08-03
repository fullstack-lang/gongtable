// insertion point for imports
import { FormCellStringDB } from './formcellstring-db'
import { FormCellFloat64DB } from './formcellfloat64-db'
import { FormCellIntDB } from './formcellint-db'
import { FormCellBooleanDB } from './formcellboolean-db'
import { FormDB } from './form-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormCellDB {

	static GONGSTRUCT_NAME = "FormCell"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	FormCellString?: FormCellStringDB
	FormCellStringID: NullInt64 = new NullInt64 // if pointer is null, FormCellString.ID = 0

	FormCellFloat64?: FormCellFloat64DB
	FormCellFloat64ID: NullInt64 = new NullInt64 // if pointer is null, FormCellFloat64.ID = 0

	FormCellInt?: FormCellIntDB
	FormCellIntID: NullInt64 = new NullInt64 // if pointer is null, FormCellInt.ID = 0

	FormCellBool?: FormCellBooleanDB
	FormCellBoolID: NullInt64 = new NullInt64 // if pointer is null, FormCellBool.ID = 0

	Form_FormCellsDBID: NullInt64 = new NullInt64
	Form_FormCellsDBID_Index: NullInt64  = new NullInt64 // store the index of the formcell instance in Form.FormCells
	Form_FormCells_reverse?: FormDB 

}
