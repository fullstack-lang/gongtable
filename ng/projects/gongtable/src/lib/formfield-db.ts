// insertion point for imports
import { FormFieldStringDB } from './formfieldstring-db'
import { FormFieldFloat64DB } from './formfieldfloat64-db'
import { FormFieldIntDB } from './formfieldint-db'
import { FormFieldBooleanDB } from './formfieldboolean-db'
import { FormDB } from './form-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormFieldDB {

	static GONGSTRUCT_NAME = "FormField"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	FormFieldString?: FormFieldStringDB
	FormFieldStringID: NullInt64 = new NullInt64 // if pointer is null, FormFieldString.ID = 0

	FormFieldFloat64?: FormFieldFloat64DB
	FormFieldFloat64ID: NullInt64 = new NullInt64 // if pointer is null, FormFieldFloat64.ID = 0

	FormFieldInt?: FormFieldIntDB
	FormFieldIntID: NullInt64 = new NullInt64 // if pointer is null, FormFieldInt.ID = 0

	FormFieldBool?: FormFieldBooleanDB
	FormFieldBoolID: NullInt64 = new NullInt64 // if pointer is null, FormFieldBool.ID = 0

	Form_FormCellsDBID: NullInt64 = new NullInt64
	Form_FormCellsDBID_Index: NullInt64  = new NullInt64 // store the index of the formfield instance in Form.FormCells
	Form_FormCells_reverse?: FormDB 

}
