package models

type FormField struct {
	Name string

	FormFieldString  *FormFieldString
	FormFieldFloat64 *FormFieldFloat64
	FormFieldInt     *FormFieldInt
	FormFieldBool    *FormFieldBoolean
}
