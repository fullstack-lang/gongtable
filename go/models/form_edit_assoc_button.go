package models

// FormEditAssocButton is a button on the
// front end to edit a 0..1-N association
// when submitted, it will
type FormEditAssocButton struct {
	Name string

	// OnEditMode is true when the user
	OnEditMode bool

	OnAssocEditon FormEditAssocButtonInterface
}

// OnAfterUpdate is called when the button is pressed
func (formEditAssocButton *FormEditAssocButton) OnAfterUpdate(
	stage *StageStruct,
	stagedInstance, frontInstance *FormEditAssocButton) {
}

type FormEditAssocButtonInterface interface {
	OnButtonPressed()
}
