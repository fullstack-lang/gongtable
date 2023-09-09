// generated code - do not edit
package probe

import (
	"log"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongtable/go/models"
)

const __dummmy__time = time.Nanosecond

// insertion point
func NewCellFormCallback(
	cell *models.Cell,
	playground *Playground,
) (cellFormCallback *CellFormCallback) {
	cellFormCallback = new(CellFormCallback)
	cellFormCallback.playground = playground
	cellFormCallback.cell = cell
	return
}

type CellFormCallback struct {
	cell *models.Cell

	playground *Playground
}

func (cellFormCallback *CellFormCallback) OnSave() {

	log.Println("CellFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	cellFormCallback.playground.formStage.Checkout()

	if cellFormCallback.cell == nil {
		cellFormCallback.cell = new(models.Cell).Stage(cellFormCallback.playground.stageOfInterest)
	}
	cell_ := cellFormCallback.cell
	_ = cell_

	// get the formGroup
	formGroup := cellFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(cell_.Name), formDiv)
		case "CellString":
			FormDivSelectFieldToField(&(cell_.CellString), cellFormCallback.playground.stageOfInterest, formDiv)
		case "CellFloat64":
			FormDivSelectFieldToField(&(cell_.CellFloat64), cellFormCallback.playground.stageOfInterest, formDiv)
		case "CellInt":
			FormDivSelectFieldToField(&(cell_.CellInt), cellFormCallback.playground.stageOfInterest, formDiv)
		case "CellBool":
			FormDivSelectFieldToField(&(cell_.CellBool), cellFormCallback.playground.stageOfInterest, formDiv)
		case "CellIcon":
			FormDivSelectFieldToField(&(cell_.CellIcon), cellFormCallback.playground.stageOfInterest, formDiv)
		}
	}

	cellFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Cell](
		cellFormCallback.playground,
	)
	cellFormCallback.playground.tableStage.Commit()
}
func NewCellBooleanFormCallback(
	cellboolean *models.CellBoolean,
	playground *Playground,
) (cellbooleanFormCallback *CellBooleanFormCallback) {
	cellbooleanFormCallback = new(CellBooleanFormCallback)
	cellbooleanFormCallback.playground = playground
	cellbooleanFormCallback.cellboolean = cellboolean
	return
}

type CellBooleanFormCallback struct {
	cellboolean *models.CellBoolean

	playground *Playground
}

func (cellbooleanFormCallback *CellBooleanFormCallback) OnSave() {

	log.Println("CellBooleanFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	cellbooleanFormCallback.playground.formStage.Checkout()

	if cellbooleanFormCallback.cellboolean == nil {
		cellbooleanFormCallback.cellboolean = new(models.CellBoolean).Stage(cellbooleanFormCallback.playground.stageOfInterest)
	}
	cellboolean_ := cellbooleanFormCallback.cellboolean
	_ = cellboolean_

	// get the formGroup
	formGroup := cellbooleanFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(cellboolean_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(cellboolean_.Value), formDiv)
		}
	}

	cellbooleanFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.CellBoolean](
		cellbooleanFormCallback.playground,
	)
	cellbooleanFormCallback.playground.tableStage.Commit()
}
func NewCellFloat64FormCallback(
	cellfloat64 *models.CellFloat64,
	playground *Playground,
) (cellfloat64FormCallback *CellFloat64FormCallback) {
	cellfloat64FormCallback = new(CellFloat64FormCallback)
	cellfloat64FormCallback.playground = playground
	cellfloat64FormCallback.cellfloat64 = cellfloat64
	return
}

type CellFloat64FormCallback struct {
	cellfloat64 *models.CellFloat64

	playground *Playground
}

func (cellfloat64FormCallback *CellFloat64FormCallback) OnSave() {

	log.Println("CellFloat64FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	cellfloat64FormCallback.playground.formStage.Checkout()

	if cellfloat64FormCallback.cellfloat64 == nil {
		cellfloat64FormCallback.cellfloat64 = new(models.CellFloat64).Stage(cellfloat64FormCallback.playground.stageOfInterest)
	}
	cellfloat64_ := cellfloat64FormCallback.cellfloat64
	_ = cellfloat64_

	// get the formGroup
	formGroup := cellfloat64FormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(cellfloat64_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(cellfloat64_.Value), formDiv)
		}
	}

	cellfloat64FormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.CellFloat64](
		cellfloat64FormCallback.playground,
	)
	cellfloat64FormCallback.playground.tableStage.Commit()
}
func NewCellIconFormCallback(
	cellicon *models.CellIcon,
	playground *Playground,
) (celliconFormCallback *CellIconFormCallback) {
	celliconFormCallback = new(CellIconFormCallback)
	celliconFormCallback.playground = playground
	celliconFormCallback.cellicon = cellicon
	return
}

type CellIconFormCallback struct {
	cellicon *models.CellIcon

	playground *Playground
}

func (celliconFormCallback *CellIconFormCallback) OnSave() {

	log.Println("CellIconFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	celliconFormCallback.playground.formStage.Checkout()

	if celliconFormCallback.cellicon == nil {
		celliconFormCallback.cellicon = new(models.CellIcon).Stage(celliconFormCallback.playground.stageOfInterest)
	}
	cellicon_ := celliconFormCallback.cellicon
	_ = cellicon_

	// get the formGroup
	formGroup := celliconFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(cellicon_.Name), formDiv)
		case "Icon":
			FormDivBasicFieldToField(&(cellicon_.Icon), formDiv)
		}
	}

	celliconFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.CellIcon](
		celliconFormCallback.playground,
	)
	celliconFormCallback.playground.tableStage.Commit()
}
func NewCellIntFormCallback(
	cellint *models.CellInt,
	playground *Playground,
) (cellintFormCallback *CellIntFormCallback) {
	cellintFormCallback = new(CellIntFormCallback)
	cellintFormCallback.playground = playground
	cellintFormCallback.cellint = cellint
	return
}

type CellIntFormCallback struct {
	cellint *models.CellInt

	playground *Playground
}

func (cellintFormCallback *CellIntFormCallback) OnSave() {

	log.Println("CellIntFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	cellintFormCallback.playground.formStage.Checkout()

	if cellintFormCallback.cellint == nil {
		cellintFormCallback.cellint = new(models.CellInt).Stage(cellintFormCallback.playground.stageOfInterest)
	}
	cellint_ := cellintFormCallback.cellint
	_ = cellint_

	// get the formGroup
	formGroup := cellintFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(cellint_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(cellint_.Value), formDiv)
		}
	}

	cellintFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.CellInt](
		cellintFormCallback.playground,
	)
	cellintFormCallback.playground.tableStage.Commit()
}
func NewCellStringFormCallback(
	cellstring *models.CellString,
	playground *Playground,
) (cellstringFormCallback *CellStringFormCallback) {
	cellstringFormCallback = new(CellStringFormCallback)
	cellstringFormCallback.playground = playground
	cellstringFormCallback.cellstring = cellstring
	return
}

type CellStringFormCallback struct {
	cellstring *models.CellString

	playground *Playground
}

func (cellstringFormCallback *CellStringFormCallback) OnSave() {

	log.Println("CellStringFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	cellstringFormCallback.playground.formStage.Checkout()

	if cellstringFormCallback.cellstring == nil {
		cellstringFormCallback.cellstring = new(models.CellString).Stage(cellstringFormCallback.playground.stageOfInterest)
	}
	cellstring_ := cellstringFormCallback.cellstring
	_ = cellstring_

	// get the formGroup
	formGroup := cellstringFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(cellstring_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(cellstring_.Value), formDiv)
		}
	}

	cellstringFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.CellString](
		cellstringFormCallback.playground,
	)
	cellstringFormCallback.playground.tableStage.Commit()
}
func NewCheckBoxFormCallback(
	checkbox *models.CheckBox,
	playground *Playground,
) (checkboxFormCallback *CheckBoxFormCallback) {
	checkboxFormCallback = new(CheckBoxFormCallback)
	checkboxFormCallback.playground = playground
	checkboxFormCallback.checkbox = checkbox
	return
}

type CheckBoxFormCallback struct {
	checkbox *models.CheckBox

	playground *Playground
}

func (checkboxFormCallback *CheckBoxFormCallback) OnSave() {

	log.Println("CheckBoxFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	checkboxFormCallback.playground.formStage.Checkout()

	if checkboxFormCallback.checkbox == nil {
		checkboxFormCallback.checkbox = new(models.CheckBox).Stage(checkboxFormCallback.playground.stageOfInterest)
	}
	checkbox_ := checkboxFormCallback.checkbox
	_ = checkbox_

	// get the formGroup
	formGroup := checkboxFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(checkbox_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(checkbox_.Value), formDiv)
		}
	}

	checkboxFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.CheckBox](
		checkboxFormCallback.playground,
	)
	checkboxFormCallback.playground.tableStage.Commit()
}
func NewDisplayedColumnFormCallback(
	displayedcolumn *models.DisplayedColumn,
	playground *Playground,
) (displayedcolumnFormCallback *DisplayedColumnFormCallback) {
	displayedcolumnFormCallback = new(DisplayedColumnFormCallback)
	displayedcolumnFormCallback.playground = playground
	displayedcolumnFormCallback.displayedcolumn = displayedcolumn
	return
}

type DisplayedColumnFormCallback struct {
	displayedcolumn *models.DisplayedColumn

	playground *Playground
}

func (displayedcolumnFormCallback *DisplayedColumnFormCallback) OnSave() {

	log.Println("DisplayedColumnFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	displayedcolumnFormCallback.playground.formStage.Checkout()

	if displayedcolumnFormCallback.displayedcolumn == nil {
		displayedcolumnFormCallback.displayedcolumn = new(models.DisplayedColumn).Stage(displayedcolumnFormCallback.playground.stageOfInterest)
	}
	displayedcolumn_ := displayedcolumnFormCallback.displayedcolumn
	_ = displayedcolumn_

	// get the formGroup
	formGroup := displayedcolumnFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(displayedcolumn_.Name), formDiv)
		}
	}

	displayedcolumnFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.DisplayedColumn](
		displayedcolumnFormCallback.playground,
	)
	displayedcolumnFormCallback.playground.tableStage.Commit()
}
func NewFormDivFormCallback(
	formdiv *models.FormDiv,
	playground *Playground,
) (formdivFormCallback *FormDivFormCallback) {
	formdivFormCallback = new(FormDivFormCallback)
	formdivFormCallback.playground = playground
	formdivFormCallback.formdiv = formdiv
	return
}

type FormDivFormCallback struct {
	formdiv *models.FormDiv

	playground *Playground
}

func (formdivFormCallback *FormDivFormCallback) OnSave() {

	log.Println("FormDivFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formdivFormCallback.playground.formStage.Checkout()

	if formdivFormCallback.formdiv == nil {
		formdivFormCallback.formdiv = new(models.FormDiv).Stage(formdivFormCallback.playground.stageOfInterest)
	}
	formdiv_ := formdivFormCallback.formdiv
	_ = formdiv_

	// get the formGroup
	formGroup := formdivFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formdiv_.Name), formDiv)
		case "FormEditAssocButton":
			FormDivSelectFieldToField(&(formdiv_.FormEditAssocButton), formdivFormCallback.playground.stageOfInterest, formDiv)
		case "FormSortAssocButton":
			FormDivSelectFieldToField(&(formdiv_.FormSortAssocButton), formdivFormCallback.playground.stageOfInterest, formDiv)
		}
	}

	formdivFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormDiv](
		formdivFormCallback.playground,
	)
	formdivFormCallback.playground.tableStage.Commit()
}
func NewFormEditAssocButtonFormCallback(
	formeditassocbutton *models.FormEditAssocButton,
	playground *Playground,
) (formeditassocbuttonFormCallback *FormEditAssocButtonFormCallback) {
	formeditassocbuttonFormCallback = new(FormEditAssocButtonFormCallback)
	formeditassocbuttonFormCallback.playground = playground
	formeditassocbuttonFormCallback.formeditassocbutton = formeditassocbutton
	return
}

type FormEditAssocButtonFormCallback struct {
	formeditassocbutton *models.FormEditAssocButton

	playground *Playground
}

func (formeditassocbuttonFormCallback *FormEditAssocButtonFormCallback) OnSave() {

	log.Println("FormEditAssocButtonFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formeditassocbuttonFormCallback.playground.formStage.Checkout()

	if formeditassocbuttonFormCallback.formeditassocbutton == nil {
		formeditassocbuttonFormCallback.formeditassocbutton = new(models.FormEditAssocButton).Stage(formeditassocbuttonFormCallback.playground.stageOfInterest)
	}
	formeditassocbutton_ := formeditassocbuttonFormCallback.formeditassocbutton
	_ = formeditassocbutton_

	// get the formGroup
	formGroup := formeditassocbuttonFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formeditassocbutton_.Name), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(formeditassocbutton_.Label), formDiv)
		}
	}

	formeditassocbuttonFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormEditAssocButton](
		formeditassocbuttonFormCallback.playground,
	)
	formeditassocbuttonFormCallback.playground.tableStage.Commit()
}
func NewFormFieldFormCallback(
	formfield *models.FormField,
	playground *Playground,
) (formfieldFormCallback *FormFieldFormCallback) {
	formfieldFormCallback = new(FormFieldFormCallback)
	formfieldFormCallback.playground = playground
	formfieldFormCallback.formfield = formfield
	return
}

type FormFieldFormCallback struct {
	formfield *models.FormField

	playground *Playground
}

func (formfieldFormCallback *FormFieldFormCallback) OnSave() {

	log.Println("FormFieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldFormCallback.playground.formStage.Checkout()

	if formfieldFormCallback.formfield == nil {
		formfieldFormCallback.formfield = new(models.FormField).Stage(formfieldFormCallback.playground.stageOfInterest)
	}
	formfield_ := formfieldFormCallback.formfield
	_ = formfield_

	// get the formGroup
	formGroup := formfieldFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfield_.Name), formDiv)
		case "InputTypeEnum":
			FormDivEnumStringFieldToField(&(formfield_.InputTypeEnum), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(formfield_.Label), formDiv)
		case "Placeholder":
			FormDivBasicFieldToField(&(formfield_.Placeholder), formDiv)
		case "FormFieldString":
			FormDivSelectFieldToField(&(formfield_.FormFieldString), formfieldFormCallback.playground.stageOfInterest, formDiv)
		case "FormFieldFloat64":
			FormDivSelectFieldToField(&(formfield_.FormFieldFloat64), formfieldFormCallback.playground.stageOfInterest, formDiv)
		case "FormFieldInt":
			FormDivSelectFieldToField(&(formfield_.FormFieldInt), formfieldFormCallback.playground.stageOfInterest, formDiv)
		case "FormFieldDate":
			FormDivSelectFieldToField(&(formfield_.FormFieldDate), formfieldFormCallback.playground.stageOfInterest, formDiv)
		case "FormFieldTime":
			FormDivSelectFieldToField(&(formfield_.FormFieldTime), formfieldFormCallback.playground.stageOfInterest, formDiv)
		case "FormFieldDateTime":
			FormDivSelectFieldToField(&(formfield_.FormFieldDateTime), formfieldFormCallback.playground.stageOfInterest, formDiv)
		case "FormFieldSelect":
			FormDivSelectFieldToField(&(formfield_.FormFieldSelect), formfieldFormCallback.playground.stageOfInterest, formDiv)
		case "HasBespokeWidth":
			FormDivBasicFieldToField(&(formfield_.HasBespokeWidth), formDiv)
		case "BespokeWidthPx":
			FormDivBasicFieldToField(&(formfield_.BespokeWidthPx), formDiv)
		}
	}

	formfieldFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormField](
		formfieldFormCallback.playground,
	)
	formfieldFormCallback.playground.tableStage.Commit()
}
func NewFormFieldDateFormCallback(
	formfielddate *models.FormFieldDate,
	playground *Playground,
) (formfielddateFormCallback *FormFieldDateFormCallback) {
	formfielddateFormCallback = new(FormFieldDateFormCallback)
	formfielddateFormCallback.playground = playground
	formfielddateFormCallback.formfielddate = formfielddate
	return
}

type FormFieldDateFormCallback struct {
	formfielddate *models.FormFieldDate

	playground *Playground
}

func (formfielddateFormCallback *FormFieldDateFormCallback) OnSave() {

	log.Println("FormFieldDateFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfielddateFormCallback.playground.formStage.Checkout()

	if formfielddateFormCallback.formfielddate == nil {
		formfielddateFormCallback.formfielddate = new(models.FormFieldDate).Stage(formfielddateFormCallback.playground.stageOfInterest)
	}
	formfielddate_ := formfielddateFormCallback.formfielddate
	_ = formfielddate_

	// get the formGroup
	formGroup := formfielddateFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfielddate_.Name), formDiv)
		}
	}

	formfielddateFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormFieldDate](
		formfielddateFormCallback.playground,
	)
	formfielddateFormCallback.playground.tableStage.Commit()
}
func NewFormFieldDateTimeFormCallback(
	formfielddatetime *models.FormFieldDateTime,
	playground *Playground,
) (formfielddatetimeFormCallback *FormFieldDateTimeFormCallback) {
	formfielddatetimeFormCallback = new(FormFieldDateTimeFormCallback)
	formfielddatetimeFormCallback.playground = playground
	formfielddatetimeFormCallback.formfielddatetime = formfielddatetime
	return
}

type FormFieldDateTimeFormCallback struct {
	formfielddatetime *models.FormFieldDateTime

	playground *Playground
}

func (formfielddatetimeFormCallback *FormFieldDateTimeFormCallback) OnSave() {

	log.Println("FormFieldDateTimeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfielddatetimeFormCallback.playground.formStage.Checkout()

	if formfielddatetimeFormCallback.formfielddatetime == nil {
		formfielddatetimeFormCallback.formfielddatetime = new(models.FormFieldDateTime).Stage(formfielddatetimeFormCallback.playground.stageOfInterest)
	}
	formfielddatetime_ := formfielddatetimeFormCallback.formfielddatetime
	_ = formfielddatetime_

	// get the formGroup
	formGroup := formfielddatetimeFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfielddatetime_.Name), formDiv)
		}
	}

	formfielddatetimeFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormFieldDateTime](
		formfielddatetimeFormCallback.playground,
	)
	formfielddatetimeFormCallback.playground.tableStage.Commit()
}
func NewFormFieldFloat64FormCallback(
	formfieldfloat64 *models.FormFieldFloat64,
	playground *Playground,
) (formfieldfloat64FormCallback *FormFieldFloat64FormCallback) {
	formfieldfloat64FormCallback = new(FormFieldFloat64FormCallback)
	formfieldfloat64FormCallback.playground = playground
	formfieldfloat64FormCallback.formfieldfloat64 = formfieldfloat64
	return
}

type FormFieldFloat64FormCallback struct {
	formfieldfloat64 *models.FormFieldFloat64

	playground *Playground
}

func (formfieldfloat64FormCallback *FormFieldFloat64FormCallback) OnSave() {

	log.Println("FormFieldFloat64FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldfloat64FormCallback.playground.formStage.Checkout()

	if formfieldfloat64FormCallback.formfieldfloat64 == nil {
		formfieldfloat64FormCallback.formfieldfloat64 = new(models.FormFieldFloat64).Stage(formfieldfloat64FormCallback.playground.stageOfInterest)
	}
	formfieldfloat64_ := formfieldfloat64FormCallback.formfieldfloat64
	_ = formfieldfloat64_

	// get the formGroup
	formGroup := formfieldfloat64FormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfieldfloat64_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(formfieldfloat64_.Value), formDiv)
		case "HasMinValidator":
			FormDivBasicFieldToField(&(formfieldfloat64_.HasMinValidator), formDiv)
		case "MinValue":
			FormDivBasicFieldToField(&(formfieldfloat64_.MinValue), formDiv)
		case "HasMaxValidator":
			FormDivBasicFieldToField(&(formfieldfloat64_.HasMaxValidator), formDiv)
		case "MaxValue":
			FormDivBasicFieldToField(&(formfieldfloat64_.MaxValue), formDiv)
		}
	}

	formfieldfloat64FormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormFieldFloat64](
		formfieldfloat64FormCallback.playground,
	)
	formfieldfloat64FormCallback.playground.tableStage.Commit()
}
func NewFormFieldIntFormCallback(
	formfieldint *models.FormFieldInt,
	playground *Playground,
) (formfieldintFormCallback *FormFieldIntFormCallback) {
	formfieldintFormCallback = new(FormFieldIntFormCallback)
	formfieldintFormCallback.playground = playground
	formfieldintFormCallback.formfieldint = formfieldint
	return
}

type FormFieldIntFormCallback struct {
	formfieldint *models.FormFieldInt

	playground *Playground
}

func (formfieldintFormCallback *FormFieldIntFormCallback) OnSave() {

	log.Println("FormFieldIntFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldintFormCallback.playground.formStage.Checkout()

	if formfieldintFormCallback.formfieldint == nil {
		formfieldintFormCallback.formfieldint = new(models.FormFieldInt).Stage(formfieldintFormCallback.playground.stageOfInterest)
	}
	formfieldint_ := formfieldintFormCallback.formfieldint
	_ = formfieldint_

	// get the formGroup
	formGroup := formfieldintFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfieldint_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(formfieldint_.Value), formDiv)
		case "HasMinValidator":
			FormDivBasicFieldToField(&(formfieldint_.HasMinValidator), formDiv)
		case "MinValue":
			FormDivBasicFieldToField(&(formfieldint_.MinValue), formDiv)
		case "HasMaxValidator":
			FormDivBasicFieldToField(&(formfieldint_.HasMaxValidator), formDiv)
		case "MaxValue":
			FormDivBasicFieldToField(&(formfieldint_.MaxValue), formDiv)
		}
	}

	formfieldintFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormFieldInt](
		formfieldintFormCallback.playground,
	)
	formfieldintFormCallback.playground.tableStage.Commit()
}
func NewFormFieldSelectFormCallback(
	formfieldselect *models.FormFieldSelect,
	playground *Playground,
) (formfieldselectFormCallback *FormFieldSelectFormCallback) {
	formfieldselectFormCallback = new(FormFieldSelectFormCallback)
	formfieldselectFormCallback.playground = playground
	formfieldselectFormCallback.formfieldselect = formfieldselect
	return
}

type FormFieldSelectFormCallback struct {
	formfieldselect *models.FormFieldSelect

	playground *Playground
}

func (formfieldselectFormCallback *FormFieldSelectFormCallback) OnSave() {

	log.Println("FormFieldSelectFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldselectFormCallback.playground.formStage.Checkout()

	if formfieldselectFormCallback.formfieldselect == nil {
		formfieldselectFormCallback.formfieldselect = new(models.FormFieldSelect).Stage(formfieldselectFormCallback.playground.stageOfInterest)
	}
	formfieldselect_ := formfieldselectFormCallback.formfieldselect
	_ = formfieldselect_

	// get the formGroup
	formGroup := formfieldselectFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfieldselect_.Name), formDiv)
		case "Value":
			FormDivSelectFieldToField(&(formfieldselect_.Value), formfieldselectFormCallback.playground.stageOfInterest, formDiv)
		case "CanBeEmpty":
			FormDivBasicFieldToField(&(formfieldselect_.CanBeEmpty), formDiv)
		}
	}

	formfieldselectFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormFieldSelect](
		formfieldselectFormCallback.playground,
	)
	formfieldselectFormCallback.playground.tableStage.Commit()
}
func NewFormFieldStringFormCallback(
	formfieldstring *models.FormFieldString,
	playground *Playground,
) (formfieldstringFormCallback *FormFieldStringFormCallback) {
	formfieldstringFormCallback = new(FormFieldStringFormCallback)
	formfieldstringFormCallback.playground = playground
	formfieldstringFormCallback.formfieldstring = formfieldstring
	return
}

type FormFieldStringFormCallback struct {
	formfieldstring *models.FormFieldString

	playground *Playground
}

func (formfieldstringFormCallback *FormFieldStringFormCallback) OnSave() {

	log.Println("FormFieldStringFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldstringFormCallback.playground.formStage.Checkout()

	if formfieldstringFormCallback.formfieldstring == nil {
		formfieldstringFormCallback.formfieldstring = new(models.FormFieldString).Stage(formfieldstringFormCallback.playground.stageOfInterest)
	}
	formfieldstring_ := formfieldstringFormCallback.formfieldstring
	_ = formfieldstring_

	// get the formGroup
	formGroup := formfieldstringFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfieldstring_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(formfieldstring_.Value), formDiv)
		}
	}

	formfieldstringFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormFieldString](
		formfieldstringFormCallback.playground,
	)
	formfieldstringFormCallback.playground.tableStage.Commit()
}
func NewFormFieldTimeFormCallback(
	formfieldtime *models.FormFieldTime,
	playground *Playground,
) (formfieldtimeFormCallback *FormFieldTimeFormCallback) {
	formfieldtimeFormCallback = new(FormFieldTimeFormCallback)
	formfieldtimeFormCallback.playground = playground
	formfieldtimeFormCallback.formfieldtime = formfieldtime
	return
}

type FormFieldTimeFormCallback struct {
	formfieldtime *models.FormFieldTime

	playground *Playground
}

func (formfieldtimeFormCallback *FormFieldTimeFormCallback) OnSave() {

	log.Println("FormFieldTimeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldtimeFormCallback.playground.formStage.Checkout()

	if formfieldtimeFormCallback.formfieldtime == nil {
		formfieldtimeFormCallback.formfieldtime = new(models.FormFieldTime).Stage(formfieldtimeFormCallback.playground.stageOfInterest)
	}
	formfieldtime_ := formfieldtimeFormCallback.formfieldtime
	_ = formfieldtime_

	// get the formGroup
	formGroup := formfieldtimeFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfieldtime_.Name), formDiv)
		case "Step":
			FormDivBasicFieldToField(&(formfieldtime_.Step), formDiv)
		}
	}

	formfieldtimeFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormFieldTime](
		formfieldtimeFormCallback.playground,
	)
	formfieldtimeFormCallback.playground.tableStage.Commit()
}
func NewFormGroupFormCallback(
	formgroup *models.FormGroup,
	playground *Playground,
) (formgroupFormCallback *FormGroupFormCallback) {
	formgroupFormCallback = new(FormGroupFormCallback)
	formgroupFormCallback.playground = playground
	formgroupFormCallback.formgroup = formgroup
	return
}

type FormGroupFormCallback struct {
	formgroup *models.FormGroup

	playground *Playground
}

func (formgroupFormCallback *FormGroupFormCallback) OnSave() {

	log.Println("FormGroupFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formgroupFormCallback.playground.formStage.Checkout()

	if formgroupFormCallback.formgroup == nil {
		formgroupFormCallback.formgroup = new(models.FormGroup).Stage(formgroupFormCallback.playground.stageOfInterest)
	}
	formgroup_ := formgroupFormCallback.formgroup
	_ = formgroup_

	// get the formGroup
	formGroup := formgroupFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formgroup_.Name), formDiv)
		}
	}

	formgroupFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormGroup](
		formgroupFormCallback.playground,
	)
	formgroupFormCallback.playground.tableStage.Commit()
}
func NewFormSortAssocButtonFormCallback(
	formsortassocbutton *models.FormSortAssocButton,
	playground *Playground,
) (formsortassocbuttonFormCallback *FormSortAssocButtonFormCallback) {
	formsortassocbuttonFormCallback = new(FormSortAssocButtonFormCallback)
	formsortassocbuttonFormCallback.playground = playground
	formsortassocbuttonFormCallback.formsortassocbutton = formsortassocbutton
	return
}

type FormSortAssocButtonFormCallback struct {
	formsortassocbutton *models.FormSortAssocButton

	playground *Playground
}

func (formsortassocbuttonFormCallback *FormSortAssocButtonFormCallback) OnSave() {

	log.Println("FormSortAssocButtonFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formsortassocbuttonFormCallback.playground.formStage.Checkout()

	if formsortassocbuttonFormCallback.formsortassocbutton == nil {
		formsortassocbuttonFormCallback.formsortassocbutton = new(models.FormSortAssocButton).Stage(formsortassocbuttonFormCallback.playground.stageOfInterest)
	}
	formsortassocbutton_ := formsortassocbuttonFormCallback.formsortassocbutton
	_ = formsortassocbutton_

	// get the formGroup
	formGroup := formsortassocbuttonFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formsortassocbutton_.Name), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(formsortassocbutton_.Label), formDiv)
		}
	}

	formsortassocbuttonFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormSortAssocButton](
		formsortassocbuttonFormCallback.playground,
	)
	formsortassocbuttonFormCallback.playground.tableStage.Commit()
}
func NewOptionFormCallback(
	option *models.Option,
	playground *Playground,
) (optionFormCallback *OptionFormCallback) {
	optionFormCallback = new(OptionFormCallback)
	optionFormCallback.playground = playground
	optionFormCallback.option = option
	return
}

type OptionFormCallback struct {
	option *models.Option

	playground *Playground
}

func (optionFormCallback *OptionFormCallback) OnSave() {

	log.Println("OptionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	optionFormCallback.playground.formStage.Checkout()

	if optionFormCallback.option == nil {
		optionFormCallback.option = new(models.Option).Stage(optionFormCallback.playground.stageOfInterest)
	}
	option_ := optionFormCallback.option
	_ = option_

	// get the formGroup
	formGroup := optionFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(option_.Name), formDiv)
		}
	}

	optionFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Option](
		optionFormCallback.playground,
	)
	optionFormCallback.playground.tableStage.Commit()
}
func NewRowFormCallback(
	row *models.Row,
	playground *Playground,
) (rowFormCallback *RowFormCallback) {
	rowFormCallback = new(RowFormCallback)
	rowFormCallback.playground = playground
	rowFormCallback.row = row
	return
}

type RowFormCallback struct {
	row *models.Row

	playground *Playground
}

func (rowFormCallback *RowFormCallback) OnSave() {

	log.Println("RowFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rowFormCallback.playground.formStage.Checkout()

	if rowFormCallback.row == nil {
		rowFormCallback.row = new(models.Row).Stage(rowFormCallback.playground.stageOfInterest)
	}
	row_ := rowFormCallback.row
	_ = row_

	// get the formGroup
	formGroup := rowFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(row_.Name), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(row_.IsChecked), formDiv)
		}
	}

	rowFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Row](
		rowFormCallback.playground,
	)
	rowFormCallback.playground.tableStage.Commit()
}
func NewTableFormCallback(
	table *models.Table,
	playground *Playground,
) (tableFormCallback *TableFormCallback) {
	tableFormCallback = new(TableFormCallback)
	tableFormCallback.playground = playground
	tableFormCallback.table = table
	return
}

type TableFormCallback struct {
	table *models.Table

	playground *Playground
}

func (tableFormCallback *TableFormCallback) OnSave() {

	log.Println("TableFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tableFormCallback.playground.formStage.Checkout()

	if tableFormCallback.table == nil {
		tableFormCallback.table = new(models.Table).Stage(tableFormCallback.playground.stageOfInterest)
	}
	table_ := tableFormCallback.table
	_ = table_

	// get the formGroup
	formGroup := tableFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(table_.Name), formDiv)
		case "HasFiltering":
			FormDivBasicFieldToField(&(table_.HasFiltering), formDiv)
		case "HasColumnSorting":
			FormDivBasicFieldToField(&(table_.HasColumnSorting), formDiv)
		case "HasPaginator":
			FormDivBasicFieldToField(&(table_.HasPaginator), formDiv)
		case "HasCheckableRows":
			FormDivBasicFieldToField(&(table_.HasCheckableRows), formDiv)
		case "HasSaveButton":
			FormDivBasicFieldToField(&(table_.HasSaveButton), formDiv)
		case "CanDragDropRows":
			FormDivBasicFieldToField(&(table_.CanDragDropRows), formDiv)
		case "HasCloseButton":
			FormDivBasicFieldToField(&(table_.HasCloseButton), formDiv)
		case "SavingInProgress":
			FormDivBasicFieldToField(&(table_.SavingInProgress), formDiv)
		case "NbOfStickyColumns":
			FormDivBasicFieldToField(&(table_.NbOfStickyColumns), formDiv)
		}
	}

	tableFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Table](
		tableFormCallback.playground,
	)
	tableFormCallback.playground.tableStage.Commit()
}
