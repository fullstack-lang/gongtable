// generated code - do not edit
package data

import (
	"log"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"
	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gongtable/go/models"
	"github.com/fullstack-lang/gongtable/go/orm"
)

const __dummmy__time = time.Nanosecond

// insertion point
func NewCellFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	cell *models.Cell,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (cellFormCallback *CellFormCallback) {
	cellFormCallback = new(CellFormCallback)
	cellFormCallback.stageOfInterest = stageOfInterest
	cellFormCallback.tableStage = tableStage
	cellFormCallback.formStage = formStage
	cellFormCallback.cell = cell
	cellFormCallback.r = r
	cellFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type CellFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	cell            *models.Cell
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (cellFormCallback *CellFormCallback) OnSave() {

	log.Println("CellFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	cellFormCallback.formStage.Checkout()

	if cellFormCallback.cell == nil {
		cellFormCallback.cell = new(models.Cell).Stage(cellFormCallback.stageOfInterest)
	}
	cell_ := cellFormCallback.cell
	_ = cell_

	// get the formGroup
	formGroup := cellFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(cell_.Name), formDiv)
		case "CellString":
			FormDivSelectFieldToField(&(cell_.CellString), cellFormCallback.stageOfInterest, formDiv)
		case "CellFloat64":
			FormDivSelectFieldToField(&(cell_.CellFloat64), cellFormCallback.stageOfInterest, formDiv)
		case "CellInt":
			FormDivSelectFieldToField(&(cell_.CellInt), cellFormCallback.stageOfInterest, formDiv)
		case "CellBool":
			FormDivSelectFieldToField(&(cell_.CellBool), cellFormCallback.stageOfInterest, formDiv)
		case "CellIcon":
			FormDivSelectFieldToField(&(cell_.CellIcon), cellFormCallback.stageOfInterest, formDiv)
		}
	}

	cellFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Cell](
		cellFormCallback.stageOfInterest,
		cellFormCallback.tableStage,
		cellFormCallback.formStage,
		cellFormCallback.r,
		cellFormCallback.backRepoOfInterest,
	)
	cellFormCallback.tableStage.Commit()
}
func NewCellBooleanFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	cellboolean *models.CellBoolean,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (cellbooleanFormCallback *CellBooleanFormCallback) {
	cellbooleanFormCallback = new(CellBooleanFormCallback)
	cellbooleanFormCallback.stageOfInterest = stageOfInterest
	cellbooleanFormCallback.tableStage = tableStage
	cellbooleanFormCallback.formStage = formStage
	cellbooleanFormCallback.cellboolean = cellboolean
	cellbooleanFormCallback.r = r
	cellbooleanFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type CellBooleanFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	cellboolean            *models.CellBoolean
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (cellbooleanFormCallback *CellBooleanFormCallback) OnSave() {

	log.Println("CellBooleanFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	cellbooleanFormCallback.formStage.Checkout()

	if cellbooleanFormCallback.cellboolean == nil {
		cellbooleanFormCallback.cellboolean = new(models.CellBoolean).Stage(cellbooleanFormCallback.stageOfInterest)
	}
	cellboolean_ := cellbooleanFormCallback.cellboolean
	_ = cellboolean_

	// get the formGroup
	formGroup := cellbooleanFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(cellboolean_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(cellboolean_.Value), formDiv)
		}
	}

	cellbooleanFormCallback.stageOfInterest.Commit()
	fillUpTable[models.CellBoolean](
		cellbooleanFormCallback.stageOfInterest,
		cellbooleanFormCallback.tableStage,
		cellbooleanFormCallback.formStage,
		cellbooleanFormCallback.r,
		cellbooleanFormCallback.backRepoOfInterest,
	)
	cellbooleanFormCallback.tableStage.Commit()
}
func NewCellFloat64FormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	cellfloat64 *models.CellFloat64,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (cellfloat64FormCallback *CellFloat64FormCallback) {
	cellfloat64FormCallback = new(CellFloat64FormCallback)
	cellfloat64FormCallback.stageOfInterest = stageOfInterest
	cellfloat64FormCallback.tableStage = tableStage
	cellfloat64FormCallback.formStage = formStage
	cellfloat64FormCallback.cellfloat64 = cellfloat64
	cellfloat64FormCallback.r = r
	cellfloat64FormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type CellFloat64FormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	cellfloat64            *models.CellFloat64
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (cellfloat64FormCallback *CellFloat64FormCallback) OnSave() {

	log.Println("CellFloat64FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	cellfloat64FormCallback.formStage.Checkout()

	if cellfloat64FormCallback.cellfloat64 == nil {
		cellfloat64FormCallback.cellfloat64 = new(models.CellFloat64).Stage(cellfloat64FormCallback.stageOfInterest)
	}
	cellfloat64_ := cellfloat64FormCallback.cellfloat64
	_ = cellfloat64_

	// get the formGroup
	formGroup := cellfloat64FormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(cellfloat64_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(cellfloat64_.Value), formDiv)
		}
	}

	cellfloat64FormCallback.stageOfInterest.Commit()
	fillUpTable[models.CellFloat64](
		cellfloat64FormCallback.stageOfInterest,
		cellfloat64FormCallback.tableStage,
		cellfloat64FormCallback.formStage,
		cellfloat64FormCallback.r,
		cellfloat64FormCallback.backRepoOfInterest,
	)
	cellfloat64FormCallback.tableStage.Commit()
}
func NewCellIconFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	cellicon *models.CellIcon,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (celliconFormCallback *CellIconFormCallback) {
	celliconFormCallback = new(CellIconFormCallback)
	celliconFormCallback.stageOfInterest = stageOfInterest
	celliconFormCallback.tableStage = tableStage
	celliconFormCallback.formStage = formStage
	celliconFormCallback.cellicon = cellicon
	celliconFormCallback.r = r
	celliconFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type CellIconFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	cellicon            *models.CellIcon
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (celliconFormCallback *CellIconFormCallback) OnSave() {

	log.Println("CellIconFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	celliconFormCallback.formStage.Checkout()

	if celliconFormCallback.cellicon == nil {
		celliconFormCallback.cellicon = new(models.CellIcon).Stage(celliconFormCallback.stageOfInterest)
	}
	cellicon_ := celliconFormCallback.cellicon
	_ = cellicon_

	// get the formGroup
	formGroup := celliconFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(cellicon_.Name), formDiv)
		case "Icon":
			FormDivBasicFieldToField(&(cellicon_.Icon), formDiv)
		}
	}

	celliconFormCallback.stageOfInterest.Commit()
	fillUpTable[models.CellIcon](
		celliconFormCallback.stageOfInterest,
		celliconFormCallback.tableStage,
		celliconFormCallback.formStage,
		celliconFormCallback.r,
		celliconFormCallback.backRepoOfInterest,
	)
	celliconFormCallback.tableStage.Commit()
}
func NewCellIntFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	cellint *models.CellInt,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (cellintFormCallback *CellIntFormCallback) {
	cellintFormCallback = new(CellIntFormCallback)
	cellintFormCallback.stageOfInterest = stageOfInterest
	cellintFormCallback.tableStage = tableStage
	cellintFormCallback.formStage = formStage
	cellintFormCallback.cellint = cellint
	cellintFormCallback.r = r
	cellintFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type CellIntFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	cellint            *models.CellInt
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (cellintFormCallback *CellIntFormCallback) OnSave() {

	log.Println("CellIntFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	cellintFormCallback.formStage.Checkout()

	if cellintFormCallback.cellint == nil {
		cellintFormCallback.cellint = new(models.CellInt).Stage(cellintFormCallback.stageOfInterest)
	}
	cellint_ := cellintFormCallback.cellint
	_ = cellint_

	// get the formGroup
	formGroup := cellintFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(cellint_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(cellint_.Value), formDiv)
		}
	}

	cellintFormCallback.stageOfInterest.Commit()
	fillUpTable[models.CellInt](
		cellintFormCallback.stageOfInterest,
		cellintFormCallback.tableStage,
		cellintFormCallback.formStage,
		cellintFormCallback.r,
		cellintFormCallback.backRepoOfInterest,
	)
	cellintFormCallback.tableStage.Commit()
}
func NewCellStringFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	cellstring *models.CellString,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (cellstringFormCallback *CellStringFormCallback) {
	cellstringFormCallback = new(CellStringFormCallback)
	cellstringFormCallback.stageOfInterest = stageOfInterest
	cellstringFormCallback.tableStage = tableStage
	cellstringFormCallback.formStage = formStage
	cellstringFormCallback.cellstring = cellstring
	cellstringFormCallback.r = r
	cellstringFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type CellStringFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	cellstring            *models.CellString
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (cellstringFormCallback *CellStringFormCallback) OnSave() {

	log.Println("CellStringFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	cellstringFormCallback.formStage.Checkout()

	if cellstringFormCallback.cellstring == nil {
		cellstringFormCallback.cellstring = new(models.CellString).Stage(cellstringFormCallback.stageOfInterest)
	}
	cellstring_ := cellstringFormCallback.cellstring
	_ = cellstring_

	// get the formGroup
	formGroup := cellstringFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(cellstring_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(cellstring_.Value), formDiv)
		}
	}

	cellstringFormCallback.stageOfInterest.Commit()
	fillUpTable[models.CellString](
		cellstringFormCallback.stageOfInterest,
		cellstringFormCallback.tableStage,
		cellstringFormCallback.formStage,
		cellstringFormCallback.r,
		cellstringFormCallback.backRepoOfInterest,
	)
	cellstringFormCallback.tableStage.Commit()
}
func NewCheckBoxFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	checkbox *models.CheckBox,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (checkboxFormCallback *CheckBoxFormCallback) {
	checkboxFormCallback = new(CheckBoxFormCallback)
	checkboxFormCallback.stageOfInterest = stageOfInterest
	checkboxFormCallback.tableStage = tableStage
	checkboxFormCallback.formStage = formStage
	checkboxFormCallback.checkbox = checkbox
	checkboxFormCallback.r = r
	checkboxFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type CheckBoxFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	checkbox            *models.CheckBox
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (checkboxFormCallback *CheckBoxFormCallback) OnSave() {

	log.Println("CheckBoxFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	checkboxFormCallback.formStage.Checkout()

	if checkboxFormCallback.checkbox == nil {
		checkboxFormCallback.checkbox = new(models.CheckBox).Stage(checkboxFormCallback.stageOfInterest)
	}
	checkbox_ := checkboxFormCallback.checkbox
	_ = checkbox_

	// get the formGroup
	formGroup := checkboxFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(checkbox_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(checkbox_.Value), formDiv)
		}
	}

	checkboxFormCallback.stageOfInterest.Commit()
	fillUpTable[models.CheckBox](
		checkboxFormCallback.stageOfInterest,
		checkboxFormCallback.tableStage,
		checkboxFormCallback.formStage,
		checkboxFormCallback.r,
		checkboxFormCallback.backRepoOfInterest,
	)
	checkboxFormCallback.tableStage.Commit()
}
func NewDisplayedColumnFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	displayedcolumn *models.DisplayedColumn,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (displayedcolumnFormCallback *DisplayedColumnFormCallback) {
	displayedcolumnFormCallback = new(DisplayedColumnFormCallback)
	displayedcolumnFormCallback.stageOfInterest = stageOfInterest
	displayedcolumnFormCallback.tableStage = tableStage
	displayedcolumnFormCallback.formStage = formStage
	displayedcolumnFormCallback.displayedcolumn = displayedcolumn
	displayedcolumnFormCallback.r = r
	displayedcolumnFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type DisplayedColumnFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	displayedcolumn            *models.DisplayedColumn
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (displayedcolumnFormCallback *DisplayedColumnFormCallback) OnSave() {

	log.Println("DisplayedColumnFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	displayedcolumnFormCallback.formStage.Checkout()

	if displayedcolumnFormCallback.displayedcolumn == nil {
		displayedcolumnFormCallback.displayedcolumn = new(models.DisplayedColumn).Stage(displayedcolumnFormCallback.stageOfInterest)
	}
	displayedcolumn_ := displayedcolumnFormCallback.displayedcolumn
	_ = displayedcolumn_

	// get the formGroup
	formGroup := displayedcolumnFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(displayedcolumn_.Name), formDiv)
		}
	}

	displayedcolumnFormCallback.stageOfInterest.Commit()
	fillUpTable[models.DisplayedColumn](
		displayedcolumnFormCallback.stageOfInterest,
		displayedcolumnFormCallback.tableStage,
		displayedcolumnFormCallback.formStage,
		displayedcolumnFormCallback.r,
		displayedcolumnFormCallback.backRepoOfInterest,
	)
	displayedcolumnFormCallback.tableStage.Commit()
}
func NewFormDivFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	formdiv *models.FormDiv,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (formdivFormCallback *FormDivFormCallback) {
	formdivFormCallback = new(FormDivFormCallback)
	formdivFormCallback.stageOfInterest = stageOfInterest
	formdivFormCallback.tableStage = tableStage
	formdivFormCallback.formStage = formStage
	formdivFormCallback.formdiv = formdiv
	formdivFormCallback.r = r
	formdivFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type FormDivFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	formdiv            *models.FormDiv
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (formdivFormCallback *FormDivFormCallback) OnSave() {

	log.Println("FormDivFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formdivFormCallback.formStage.Checkout()

	if formdivFormCallback.formdiv == nil {
		formdivFormCallback.formdiv = new(models.FormDiv).Stage(formdivFormCallback.stageOfInterest)
	}
	formdiv_ := formdivFormCallback.formdiv
	_ = formdiv_

	// get the formGroup
	formGroup := formdivFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formdiv_.Name), formDiv)
		case "FormEditAssocButton":
			FormDivSelectFieldToField(&(formdiv_.FormEditAssocButton), formdivFormCallback.stageOfInterest, formDiv)
		case "FormSortAssocButton":
			FormDivSelectFieldToField(&(formdiv_.FormSortAssocButton), formdivFormCallback.stageOfInterest, formDiv)
		}
	}

	formdivFormCallback.stageOfInterest.Commit()
	fillUpTable[models.FormDiv](
		formdivFormCallback.stageOfInterest,
		formdivFormCallback.tableStage,
		formdivFormCallback.formStage,
		formdivFormCallback.r,
		formdivFormCallback.backRepoOfInterest,
	)
	formdivFormCallback.tableStage.Commit()
}
func NewFormEditAssocButtonFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	formeditassocbutton *models.FormEditAssocButton,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (formeditassocbuttonFormCallback *FormEditAssocButtonFormCallback) {
	formeditassocbuttonFormCallback = new(FormEditAssocButtonFormCallback)
	formeditassocbuttonFormCallback.stageOfInterest = stageOfInterest
	formeditassocbuttonFormCallback.tableStage = tableStage
	formeditassocbuttonFormCallback.formStage = formStage
	formeditassocbuttonFormCallback.formeditassocbutton = formeditassocbutton
	formeditassocbuttonFormCallback.r = r
	formeditassocbuttonFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type FormEditAssocButtonFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	formeditassocbutton            *models.FormEditAssocButton
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (formeditassocbuttonFormCallback *FormEditAssocButtonFormCallback) OnSave() {

	log.Println("FormEditAssocButtonFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formeditassocbuttonFormCallback.formStage.Checkout()

	if formeditassocbuttonFormCallback.formeditassocbutton == nil {
		formeditassocbuttonFormCallback.formeditassocbutton = new(models.FormEditAssocButton).Stage(formeditassocbuttonFormCallback.stageOfInterest)
	}
	formeditassocbutton_ := formeditassocbuttonFormCallback.formeditassocbutton
	_ = formeditassocbutton_

	// get the formGroup
	formGroup := formeditassocbuttonFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formeditassocbutton_.Name), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(formeditassocbutton_.Label), formDiv)
		}
	}

	formeditassocbuttonFormCallback.stageOfInterest.Commit()
	fillUpTable[models.FormEditAssocButton](
		formeditassocbuttonFormCallback.stageOfInterest,
		formeditassocbuttonFormCallback.tableStage,
		formeditassocbuttonFormCallback.formStage,
		formeditassocbuttonFormCallback.r,
		formeditassocbuttonFormCallback.backRepoOfInterest,
	)
	formeditassocbuttonFormCallback.tableStage.Commit()
}
func NewFormFieldFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	formfield *models.FormField,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (formfieldFormCallback *FormFieldFormCallback) {
	formfieldFormCallback = new(FormFieldFormCallback)
	formfieldFormCallback.stageOfInterest = stageOfInterest
	formfieldFormCallback.tableStage = tableStage
	formfieldFormCallback.formStage = formStage
	formfieldFormCallback.formfield = formfield
	formfieldFormCallback.r = r
	formfieldFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type FormFieldFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	formfield            *models.FormField
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (formfieldFormCallback *FormFieldFormCallback) OnSave() {

	log.Println("FormFieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldFormCallback.formStage.Checkout()

	if formfieldFormCallback.formfield == nil {
		formfieldFormCallback.formfield = new(models.FormField).Stage(formfieldFormCallback.stageOfInterest)
	}
	formfield_ := formfieldFormCallback.formfield
	_ = formfield_

	// get the formGroup
	formGroup := formfieldFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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
			FormDivSelectFieldToField(&(formfield_.FormFieldString), formfieldFormCallback.stageOfInterest, formDiv)
		case "FormFieldFloat64":
			FormDivSelectFieldToField(&(formfield_.FormFieldFloat64), formfieldFormCallback.stageOfInterest, formDiv)
		case "FormFieldInt":
			FormDivSelectFieldToField(&(formfield_.FormFieldInt), formfieldFormCallback.stageOfInterest, formDiv)
		case "FormFieldDate":
			FormDivSelectFieldToField(&(formfield_.FormFieldDate), formfieldFormCallback.stageOfInterest, formDiv)
		case "FormFieldTime":
			FormDivSelectFieldToField(&(formfield_.FormFieldTime), formfieldFormCallback.stageOfInterest, formDiv)
		case "FormFieldDateTime":
			FormDivSelectFieldToField(&(formfield_.FormFieldDateTime), formfieldFormCallback.stageOfInterest, formDiv)
		case "FormFieldSelect":
			FormDivSelectFieldToField(&(formfield_.FormFieldSelect), formfieldFormCallback.stageOfInterest, formDiv)
		case "HasBespokeWidth":
			FormDivBasicFieldToField(&(formfield_.HasBespokeWidth), formDiv)
		case "BespokeWidthPx":
			FormDivBasicFieldToField(&(formfield_.BespokeWidthPx), formDiv)
		}
	}

	formfieldFormCallback.stageOfInterest.Commit()
	fillUpTable[models.FormField](
		formfieldFormCallback.stageOfInterest,
		formfieldFormCallback.tableStage,
		formfieldFormCallback.formStage,
		formfieldFormCallback.r,
		formfieldFormCallback.backRepoOfInterest,
	)
	formfieldFormCallback.tableStage.Commit()
}
func NewFormFieldDateFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	formfielddate *models.FormFieldDate,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (formfielddateFormCallback *FormFieldDateFormCallback) {
	formfielddateFormCallback = new(FormFieldDateFormCallback)
	formfielddateFormCallback.stageOfInterest = stageOfInterest
	formfielddateFormCallback.tableStage = tableStage
	formfielddateFormCallback.formStage = formStage
	formfielddateFormCallback.formfielddate = formfielddate
	formfielddateFormCallback.r = r
	formfielddateFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type FormFieldDateFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	formfielddate            *models.FormFieldDate
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (formfielddateFormCallback *FormFieldDateFormCallback) OnSave() {

	log.Println("FormFieldDateFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfielddateFormCallback.formStage.Checkout()

	if formfielddateFormCallback.formfielddate == nil {
		formfielddateFormCallback.formfielddate = new(models.FormFieldDate).Stage(formfielddateFormCallback.stageOfInterest)
	}
	formfielddate_ := formfielddateFormCallback.formfielddate
	_ = formfielddate_

	// get the formGroup
	formGroup := formfielddateFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfielddate_.Name), formDiv)
		}
	}

	formfielddateFormCallback.stageOfInterest.Commit()
	fillUpTable[models.FormFieldDate](
		formfielddateFormCallback.stageOfInterest,
		formfielddateFormCallback.tableStage,
		formfielddateFormCallback.formStage,
		formfielddateFormCallback.r,
		formfielddateFormCallback.backRepoOfInterest,
	)
	formfielddateFormCallback.tableStage.Commit()
}
func NewFormFieldDateTimeFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	formfielddatetime *models.FormFieldDateTime,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (formfielddatetimeFormCallback *FormFieldDateTimeFormCallback) {
	formfielddatetimeFormCallback = new(FormFieldDateTimeFormCallback)
	formfielddatetimeFormCallback.stageOfInterest = stageOfInterest
	formfielddatetimeFormCallback.tableStage = tableStage
	formfielddatetimeFormCallback.formStage = formStage
	formfielddatetimeFormCallback.formfielddatetime = formfielddatetime
	formfielddatetimeFormCallback.r = r
	formfielddatetimeFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type FormFieldDateTimeFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	formfielddatetime            *models.FormFieldDateTime
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (formfielddatetimeFormCallback *FormFieldDateTimeFormCallback) OnSave() {

	log.Println("FormFieldDateTimeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfielddatetimeFormCallback.formStage.Checkout()

	if formfielddatetimeFormCallback.formfielddatetime == nil {
		formfielddatetimeFormCallback.formfielddatetime = new(models.FormFieldDateTime).Stage(formfielddatetimeFormCallback.stageOfInterest)
	}
	formfielddatetime_ := formfielddatetimeFormCallback.formfielddatetime
	_ = formfielddatetime_

	// get the formGroup
	formGroup := formfielddatetimeFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfielddatetime_.Name), formDiv)
		}
	}

	formfielddatetimeFormCallback.stageOfInterest.Commit()
	fillUpTable[models.FormFieldDateTime](
		formfielddatetimeFormCallback.stageOfInterest,
		formfielddatetimeFormCallback.tableStage,
		formfielddatetimeFormCallback.formStage,
		formfielddatetimeFormCallback.r,
		formfielddatetimeFormCallback.backRepoOfInterest,
	)
	formfielddatetimeFormCallback.tableStage.Commit()
}
func NewFormFieldFloat64FormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	formfieldfloat64 *models.FormFieldFloat64,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (formfieldfloat64FormCallback *FormFieldFloat64FormCallback) {
	formfieldfloat64FormCallback = new(FormFieldFloat64FormCallback)
	formfieldfloat64FormCallback.stageOfInterest = stageOfInterest
	formfieldfloat64FormCallback.tableStage = tableStage
	formfieldfloat64FormCallback.formStage = formStage
	formfieldfloat64FormCallback.formfieldfloat64 = formfieldfloat64
	formfieldfloat64FormCallback.r = r
	formfieldfloat64FormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type FormFieldFloat64FormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	formfieldfloat64            *models.FormFieldFloat64
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (formfieldfloat64FormCallback *FormFieldFloat64FormCallback) OnSave() {

	log.Println("FormFieldFloat64FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldfloat64FormCallback.formStage.Checkout()

	if formfieldfloat64FormCallback.formfieldfloat64 == nil {
		formfieldfloat64FormCallback.formfieldfloat64 = new(models.FormFieldFloat64).Stage(formfieldfloat64FormCallback.stageOfInterest)
	}
	formfieldfloat64_ := formfieldfloat64FormCallback.formfieldfloat64
	_ = formfieldfloat64_

	// get the formGroup
	formGroup := formfieldfloat64FormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	formfieldfloat64FormCallback.stageOfInterest.Commit()
	fillUpTable[models.FormFieldFloat64](
		formfieldfloat64FormCallback.stageOfInterest,
		formfieldfloat64FormCallback.tableStage,
		formfieldfloat64FormCallback.formStage,
		formfieldfloat64FormCallback.r,
		formfieldfloat64FormCallback.backRepoOfInterest,
	)
	formfieldfloat64FormCallback.tableStage.Commit()
}
func NewFormFieldIntFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	formfieldint *models.FormFieldInt,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (formfieldintFormCallback *FormFieldIntFormCallback) {
	formfieldintFormCallback = new(FormFieldIntFormCallback)
	formfieldintFormCallback.stageOfInterest = stageOfInterest
	formfieldintFormCallback.tableStage = tableStage
	formfieldintFormCallback.formStage = formStage
	formfieldintFormCallback.formfieldint = formfieldint
	formfieldintFormCallback.r = r
	formfieldintFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type FormFieldIntFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	formfieldint            *models.FormFieldInt
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (formfieldintFormCallback *FormFieldIntFormCallback) OnSave() {

	log.Println("FormFieldIntFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldintFormCallback.formStage.Checkout()

	if formfieldintFormCallback.formfieldint == nil {
		formfieldintFormCallback.formfieldint = new(models.FormFieldInt).Stage(formfieldintFormCallback.stageOfInterest)
	}
	formfieldint_ := formfieldintFormCallback.formfieldint
	_ = formfieldint_

	// get the formGroup
	formGroup := formfieldintFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	formfieldintFormCallback.stageOfInterest.Commit()
	fillUpTable[models.FormFieldInt](
		formfieldintFormCallback.stageOfInterest,
		formfieldintFormCallback.tableStage,
		formfieldintFormCallback.formStage,
		formfieldintFormCallback.r,
		formfieldintFormCallback.backRepoOfInterest,
	)
	formfieldintFormCallback.tableStage.Commit()
}
func NewFormFieldSelectFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	formfieldselect *models.FormFieldSelect,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (formfieldselectFormCallback *FormFieldSelectFormCallback) {
	formfieldselectFormCallback = new(FormFieldSelectFormCallback)
	formfieldselectFormCallback.stageOfInterest = stageOfInterest
	formfieldselectFormCallback.tableStage = tableStage
	formfieldselectFormCallback.formStage = formStage
	formfieldselectFormCallback.formfieldselect = formfieldselect
	formfieldselectFormCallback.r = r
	formfieldselectFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type FormFieldSelectFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	formfieldselect            *models.FormFieldSelect
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (formfieldselectFormCallback *FormFieldSelectFormCallback) OnSave() {

	log.Println("FormFieldSelectFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldselectFormCallback.formStage.Checkout()

	if formfieldselectFormCallback.formfieldselect == nil {
		formfieldselectFormCallback.formfieldselect = new(models.FormFieldSelect).Stage(formfieldselectFormCallback.stageOfInterest)
	}
	formfieldselect_ := formfieldselectFormCallback.formfieldselect
	_ = formfieldselect_

	// get the formGroup
	formGroup := formfieldselectFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfieldselect_.Name), formDiv)
		case "Value":
			FormDivSelectFieldToField(&(formfieldselect_.Value), formfieldselectFormCallback.stageOfInterest, formDiv)
		case "CanBeEmpty":
			FormDivBasicFieldToField(&(formfieldselect_.CanBeEmpty), formDiv)
		}
	}

	formfieldselectFormCallback.stageOfInterest.Commit()
	fillUpTable[models.FormFieldSelect](
		formfieldselectFormCallback.stageOfInterest,
		formfieldselectFormCallback.tableStage,
		formfieldselectFormCallback.formStage,
		formfieldselectFormCallback.r,
		formfieldselectFormCallback.backRepoOfInterest,
	)
	formfieldselectFormCallback.tableStage.Commit()
}
func NewFormFieldStringFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	formfieldstring *models.FormFieldString,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (formfieldstringFormCallback *FormFieldStringFormCallback) {
	formfieldstringFormCallback = new(FormFieldStringFormCallback)
	formfieldstringFormCallback.stageOfInterest = stageOfInterest
	formfieldstringFormCallback.tableStage = tableStage
	formfieldstringFormCallback.formStage = formStage
	formfieldstringFormCallback.formfieldstring = formfieldstring
	formfieldstringFormCallback.r = r
	formfieldstringFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type FormFieldStringFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	formfieldstring            *models.FormFieldString
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (formfieldstringFormCallback *FormFieldStringFormCallback) OnSave() {

	log.Println("FormFieldStringFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldstringFormCallback.formStage.Checkout()

	if formfieldstringFormCallback.formfieldstring == nil {
		formfieldstringFormCallback.formfieldstring = new(models.FormFieldString).Stage(formfieldstringFormCallback.stageOfInterest)
	}
	formfieldstring_ := formfieldstringFormCallback.formfieldstring
	_ = formfieldstring_

	// get the formGroup
	formGroup := formfieldstringFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfieldstring_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(formfieldstring_.Value), formDiv)
		}
	}

	formfieldstringFormCallback.stageOfInterest.Commit()
	fillUpTable[models.FormFieldString](
		formfieldstringFormCallback.stageOfInterest,
		formfieldstringFormCallback.tableStage,
		formfieldstringFormCallback.formStage,
		formfieldstringFormCallback.r,
		formfieldstringFormCallback.backRepoOfInterest,
	)
	formfieldstringFormCallback.tableStage.Commit()
}
func NewFormFieldTimeFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	formfieldtime *models.FormFieldTime,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (formfieldtimeFormCallback *FormFieldTimeFormCallback) {
	formfieldtimeFormCallback = new(FormFieldTimeFormCallback)
	formfieldtimeFormCallback.stageOfInterest = stageOfInterest
	formfieldtimeFormCallback.tableStage = tableStage
	formfieldtimeFormCallback.formStage = formStage
	formfieldtimeFormCallback.formfieldtime = formfieldtime
	formfieldtimeFormCallback.r = r
	formfieldtimeFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type FormFieldTimeFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	formfieldtime            *models.FormFieldTime
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (formfieldtimeFormCallback *FormFieldTimeFormCallback) OnSave() {

	log.Println("FormFieldTimeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldtimeFormCallback.formStage.Checkout()

	if formfieldtimeFormCallback.formfieldtime == nil {
		formfieldtimeFormCallback.formfieldtime = new(models.FormFieldTime).Stage(formfieldtimeFormCallback.stageOfInterest)
	}
	formfieldtime_ := formfieldtimeFormCallback.formfieldtime
	_ = formfieldtime_

	// get the formGroup
	formGroup := formfieldtimeFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfieldtime_.Name), formDiv)
		case "Step":
			FormDivBasicFieldToField(&(formfieldtime_.Step), formDiv)
		}
	}

	formfieldtimeFormCallback.stageOfInterest.Commit()
	fillUpTable[models.FormFieldTime](
		formfieldtimeFormCallback.stageOfInterest,
		formfieldtimeFormCallback.tableStage,
		formfieldtimeFormCallback.formStage,
		formfieldtimeFormCallback.r,
		formfieldtimeFormCallback.backRepoOfInterest,
	)
	formfieldtimeFormCallback.tableStage.Commit()
}
func NewFormGroupFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	formgroup *models.FormGroup,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (formgroupFormCallback *FormGroupFormCallback) {
	formgroupFormCallback = new(FormGroupFormCallback)
	formgroupFormCallback.stageOfInterest = stageOfInterest
	formgroupFormCallback.tableStage = tableStage
	formgroupFormCallback.formStage = formStage
	formgroupFormCallback.formgroup = formgroup
	formgroupFormCallback.r = r
	formgroupFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type FormGroupFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	formgroup            *models.FormGroup
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (formgroupFormCallback *FormGroupFormCallback) OnSave() {

	log.Println("FormGroupFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formgroupFormCallback.formStage.Checkout()

	if formgroupFormCallback.formgroup == nil {
		formgroupFormCallback.formgroup = new(models.FormGroup).Stage(formgroupFormCallback.stageOfInterest)
	}
	formgroup_ := formgroupFormCallback.formgroup
	_ = formgroup_

	// get the formGroup
	formGroup := formgroupFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formgroup_.Name), formDiv)
		}
	}

	formgroupFormCallback.stageOfInterest.Commit()
	fillUpTable[models.FormGroup](
		formgroupFormCallback.stageOfInterest,
		formgroupFormCallback.tableStage,
		formgroupFormCallback.formStage,
		formgroupFormCallback.r,
		formgroupFormCallback.backRepoOfInterest,
	)
	formgroupFormCallback.tableStage.Commit()
}
func NewFormSortAssocButtonFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	formsortassocbutton *models.FormSortAssocButton,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (formsortassocbuttonFormCallback *FormSortAssocButtonFormCallback) {
	formsortassocbuttonFormCallback = new(FormSortAssocButtonFormCallback)
	formsortassocbuttonFormCallback.stageOfInterest = stageOfInterest
	formsortassocbuttonFormCallback.tableStage = tableStage
	formsortassocbuttonFormCallback.formStage = formStage
	formsortassocbuttonFormCallback.formsortassocbutton = formsortassocbutton
	formsortassocbuttonFormCallback.r = r
	formsortassocbuttonFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type FormSortAssocButtonFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	formsortassocbutton            *models.FormSortAssocButton
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (formsortassocbuttonFormCallback *FormSortAssocButtonFormCallback) OnSave() {

	log.Println("FormSortAssocButtonFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formsortassocbuttonFormCallback.formStage.Checkout()

	if formsortassocbuttonFormCallback.formsortassocbutton == nil {
		formsortassocbuttonFormCallback.formsortassocbutton = new(models.FormSortAssocButton).Stage(formsortassocbuttonFormCallback.stageOfInterest)
	}
	formsortassocbutton_ := formsortassocbuttonFormCallback.formsortassocbutton
	_ = formsortassocbutton_

	// get the formGroup
	formGroup := formsortassocbuttonFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formsortassocbutton_.Name), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(formsortassocbutton_.Label), formDiv)
		}
	}

	formsortassocbuttonFormCallback.stageOfInterest.Commit()
	fillUpTable[models.FormSortAssocButton](
		formsortassocbuttonFormCallback.stageOfInterest,
		formsortassocbuttonFormCallback.tableStage,
		formsortassocbuttonFormCallback.formStage,
		formsortassocbuttonFormCallback.r,
		formsortassocbuttonFormCallback.backRepoOfInterest,
	)
	formsortassocbuttonFormCallback.tableStage.Commit()
}
func NewOptionFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	option *models.Option,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (optionFormCallback *OptionFormCallback) {
	optionFormCallback = new(OptionFormCallback)
	optionFormCallback.stageOfInterest = stageOfInterest
	optionFormCallback.tableStage = tableStage
	optionFormCallback.formStage = formStage
	optionFormCallback.option = option
	optionFormCallback.r = r
	optionFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type OptionFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	option            *models.Option
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (optionFormCallback *OptionFormCallback) OnSave() {

	log.Println("OptionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	optionFormCallback.formStage.Checkout()

	if optionFormCallback.option == nil {
		optionFormCallback.option = new(models.Option).Stage(optionFormCallback.stageOfInterest)
	}
	option_ := optionFormCallback.option
	_ = option_

	// get the formGroup
	formGroup := optionFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(option_.Name), formDiv)
		}
	}

	optionFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Option](
		optionFormCallback.stageOfInterest,
		optionFormCallback.tableStage,
		optionFormCallback.formStage,
		optionFormCallback.r,
		optionFormCallback.backRepoOfInterest,
	)
	optionFormCallback.tableStage.Commit()
}
func NewRowFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	row *models.Row,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (rowFormCallback *RowFormCallback) {
	rowFormCallback = new(RowFormCallback)
	rowFormCallback.stageOfInterest = stageOfInterest
	rowFormCallback.tableStage = tableStage
	rowFormCallback.formStage = formStage
	rowFormCallback.row = row
	rowFormCallback.r = r
	rowFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type RowFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	row            *models.Row
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (rowFormCallback *RowFormCallback) OnSave() {

	log.Println("RowFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rowFormCallback.formStage.Checkout()

	if rowFormCallback.row == nil {
		rowFormCallback.row = new(models.Row).Stage(rowFormCallback.stageOfInterest)
	}
	row_ := rowFormCallback.row
	_ = row_

	// get the formGroup
	formGroup := rowFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(row_.Name), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(row_.IsChecked), formDiv)
		}
	}

	rowFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Row](
		rowFormCallback.stageOfInterest,
		rowFormCallback.tableStage,
		rowFormCallback.formStage,
		rowFormCallback.r,
		rowFormCallback.backRepoOfInterest,
	)
	rowFormCallback.tableStage.Commit()
}
func NewTableFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	table *models.Table,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (tableFormCallback *TableFormCallback) {
	tableFormCallback = new(TableFormCallback)
	tableFormCallback.stageOfInterest = stageOfInterest
	tableFormCallback.tableStage = tableStage
	tableFormCallback.formStage = formStage
	tableFormCallback.table = table
	tableFormCallback.r = r
	tableFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type TableFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	table            *models.Table
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (tableFormCallback *TableFormCallback) OnSave() {

	log.Println("TableFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tableFormCallback.formStage.Checkout()

	if tableFormCallback.table == nil {
		tableFormCallback.table = new(models.Table).Stage(tableFormCallback.stageOfInterest)
	}
	table_ := tableFormCallback.table
	_ = table_

	// get the formGroup
	formGroup := tableFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	tableFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Table](
		tableFormCallback.stageOfInterest,
		tableFormCallback.tableStage,
		tableFormCallback.formStage,
		tableFormCallback.r,
		tableFormCallback.backRepoOfInterest,
	)
	tableFormCallback.tableStage.Commit()
}
