package models

import "log"

// FormEditAssocButton is a button on the
// front end to edit a 0..1-N association
// when submitted, it will
type FormEditAssocButton struct {
	Name string

	Label string

	// OnEditMode is true when the user
	OnEditMode bool

	OnAssocEditon FormEditAssocButtonInterface
}

// OnAfterUpdate is called when the button is pressed
func (formEditAssocButton *FormEditAssocButton) OnAfterUpdate(
	stage *StageStruct,
	stagedInstance, frontInstance *FormEditAssocButton) {

	log.Println("OnAfterUpdate")

	// start the stack and reset the field
	if frontInstance.OnEditMode == true {
		stagedInstance.Commit(stage)

		if stagedInstance.OnAssocEditon != nil {
			stagedInstance.OnAssocEditon.OnButtonPressed()
		}
	}
}

type FormEditAssocButtonInterface interface {
	OnButtonPressed()
}
