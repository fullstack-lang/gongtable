package models

type FormField struct {
	Name          string
	InputTypeEnum InputTypeEnum

	// label for the input field.
	// for instance "First Name" in <mat-label>First Name</mat-label>
	Label string

	FormFieldString  *FormFieldString
	FormFieldFloat64 *FormFieldFloat64
	FormFieldInt     *FormFieldInt
	FormFieldBool    *FormFieldBoolean
}
