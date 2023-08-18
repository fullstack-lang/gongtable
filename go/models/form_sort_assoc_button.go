package models

import "log"

// FormSortAssocButton is a button on the
// front end to sort a 0..1-N association
// when submitted, it will update the data
type FormSortAssocButton struct {
	Name string

	Label string

	OnAssocEditon FormSortAssocButtonInterface
}

// OnAfterUpdate is called when the button is pressed
func (formSortAssocButton *FormSortAssocButton) OnAfterUpdate(
	stage *StageStruct,
	stagedInstance, frontInstance *FormSortAssocButton) {

	log.Println("OnAfterUpdate")

	if stagedInstance.OnAssocEditon != nil {
		stagedInstance.OnAssocEditon.OnButtonPressed()
	}

}

type FormSortAssocButtonInterface interface {
	OnButtonPressed()
}
