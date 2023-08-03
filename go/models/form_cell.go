package models

type FormCell struct {
	Name string

	FormCellString  *FormCellString
	FormCellFloat64 *FormCellFloat64
	FormCellInt     *FormCellInt
	FormCellBool    *FormCellBoolean
}
