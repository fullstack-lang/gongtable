// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongtable/go/models"
	"github.com/fullstack-lang/gongtable/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__CellFormCallback(
	cell *models.Cell,
	playground *Playground,
) (cellFormCallback *CellFormCallback) {
	cellFormCallback = new(CellFormCallback)
	cellFormCallback.playground = playground
	cellFormCallback.cell = cell

	cellFormCallback.CreationMode = (cell == nil)

	return
}

type CellFormCallback struct {
	cell *models.Cell

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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
		case "Row:Cells":
			// we need to retrieve the field owner before the change
			var pastRowOwner *models.Row
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Row"
			rf.Fieldname = "Cells"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				cellFormCallback.playground.stageOfInterest,
				cellFormCallback.playground.backRepoOfInterest,
				cell_,
				&rf)

			if reverseFieldOwner != nil {
				pastRowOwner = reverseFieldOwner.(*models.Row)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRowOwner != nil {
					idx := slices.Index(pastRowOwner.Cells, cell_)
					pastRowOwner.Cells = slices.Delete(pastRowOwner.Cells, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _row := range *models.GetGongstructInstancesSet[models.Row](cellFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _row.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRowOwner := _row // we have a match
						if pastRowOwner != nil {
							if newRowOwner != pastRowOwner {
								idx := slices.Index(pastRowOwner.Cells, cell_)
								pastRowOwner.Cells = slices.Delete(pastRowOwner.Cells, idx, idx+1)
								newRowOwner.Cells = append(newRowOwner.Cells, cell_)
							}
						} else {
							newRowOwner.Cells = append(newRowOwner.Cells, cell_)
						}
					}
				}
			}
		}
	}

	cellFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Cell](
		cellFormCallback.playground,
	)
	cellFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if cellFormCallback.CreationMode {
		cellFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__CellFormCallback(
				nil,
				cellFormCallback.playground,
			),
		}).Stage(cellFormCallback.playground.formStage)
		cell := new(models.Cell)
		FillUpForm(cell, newFormGroup, cellFormCallback.playground)
		cellFormCallback.playground.formStage.Commit()
	}

	fillUpTree(cellFormCallback.playground)
}
func __gong__New__CellBooleanFormCallback(
	cellboolean *models.CellBoolean,
	playground *Playground,
) (cellbooleanFormCallback *CellBooleanFormCallback) {
	cellbooleanFormCallback = new(CellBooleanFormCallback)
	cellbooleanFormCallback.playground = playground
	cellbooleanFormCallback.cellboolean = cellboolean

	cellbooleanFormCallback.CreationMode = (cellboolean == nil)

	return
}

type CellBooleanFormCallback struct {
	cellboolean *models.CellBoolean

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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

	// display a new form by reset the form stage
	if cellbooleanFormCallback.CreationMode {
		cellbooleanFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__CellBooleanFormCallback(
				nil,
				cellbooleanFormCallback.playground,
			),
		}).Stage(cellbooleanFormCallback.playground.formStage)
		cellboolean := new(models.CellBoolean)
		FillUpForm(cellboolean, newFormGroup, cellbooleanFormCallback.playground)
		cellbooleanFormCallback.playground.formStage.Commit()
	}

	fillUpTree(cellbooleanFormCallback.playground)
}
func __gong__New__CellFloat64FormCallback(
	cellfloat64 *models.CellFloat64,
	playground *Playground,
) (cellfloat64FormCallback *CellFloat64FormCallback) {
	cellfloat64FormCallback = new(CellFloat64FormCallback)
	cellfloat64FormCallback.playground = playground
	cellfloat64FormCallback.cellfloat64 = cellfloat64

	cellfloat64FormCallback.CreationMode = (cellfloat64 == nil)

	return
}

type CellFloat64FormCallback struct {
	cellfloat64 *models.CellFloat64

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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

	// display a new form by reset the form stage
	if cellfloat64FormCallback.CreationMode {
		cellfloat64FormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__CellFloat64FormCallback(
				nil,
				cellfloat64FormCallback.playground,
			),
		}).Stage(cellfloat64FormCallback.playground.formStage)
		cellfloat64 := new(models.CellFloat64)
		FillUpForm(cellfloat64, newFormGroup, cellfloat64FormCallback.playground)
		cellfloat64FormCallback.playground.formStage.Commit()
	}

	fillUpTree(cellfloat64FormCallback.playground)
}
func __gong__New__CellIconFormCallback(
	cellicon *models.CellIcon,
	playground *Playground,
) (celliconFormCallback *CellIconFormCallback) {
	celliconFormCallback = new(CellIconFormCallback)
	celliconFormCallback.playground = playground
	celliconFormCallback.cellicon = cellicon

	celliconFormCallback.CreationMode = (cellicon == nil)

	return
}

type CellIconFormCallback struct {
	cellicon *models.CellIcon

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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

	// display a new form by reset the form stage
	if celliconFormCallback.CreationMode {
		celliconFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__CellIconFormCallback(
				nil,
				celliconFormCallback.playground,
			),
		}).Stage(celliconFormCallback.playground.formStage)
		cellicon := new(models.CellIcon)
		FillUpForm(cellicon, newFormGroup, celliconFormCallback.playground)
		celliconFormCallback.playground.formStage.Commit()
	}

	fillUpTree(celliconFormCallback.playground)
}
func __gong__New__CellIntFormCallback(
	cellint *models.CellInt,
	playground *Playground,
) (cellintFormCallback *CellIntFormCallback) {
	cellintFormCallback = new(CellIntFormCallback)
	cellintFormCallback.playground = playground
	cellintFormCallback.cellint = cellint

	cellintFormCallback.CreationMode = (cellint == nil)

	return
}

type CellIntFormCallback struct {
	cellint *models.CellInt

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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

	// display a new form by reset the form stage
	if cellintFormCallback.CreationMode {
		cellintFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__CellIntFormCallback(
				nil,
				cellintFormCallback.playground,
			),
		}).Stage(cellintFormCallback.playground.formStage)
		cellint := new(models.CellInt)
		FillUpForm(cellint, newFormGroup, cellintFormCallback.playground)
		cellintFormCallback.playground.formStage.Commit()
	}

	fillUpTree(cellintFormCallback.playground)
}
func __gong__New__CellStringFormCallback(
	cellstring *models.CellString,
	playground *Playground,
) (cellstringFormCallback *CellStringFormCallback) {
	cellstringFormCallback = new(CellStringFormCallback)
	cellstringFormCallback.playground = playground
	cellstringFormCallback.cellstring = cellstring

	cellstringFormCallback.CreationMode = (cellstring == nil)

	return
}

type CellStringFormCallback struct {
	cellstring *models.CellString

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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

	// display a new form by reset the form stage
	if cellstringFormCallback.CreationMode {
		cellstringFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__CellStringFormCallback(
				nil,
				cellstringFormCallback.playground,
			),
		}).Stage(cellstringFormCallback.playground.formStage)
		cellstring := new(models.CellString)
		FillUpForm(cellstring, newFormGroup, cellstringFormCallback.playground)
		cellstringFormCallback.playground.formStage.Commit()
	}

	fillUpTree(cellstringFormCallback.playground)
}
func __gong__New__CheckBoxFormCallback(
	checkbox *models.CheckBox,
	playground *Playground,
) (checkboxFormCallback *CheckBoxFormCallback) {
	checkboxFormCallback = new(CheckBoxFormCallback)
	checkboxFormCallback.playground = playground
	checkboxFormCallback.checkbox = checkbox

	checkboxFormCallback.CreationMode = (checkbox == nil)

	return
}

type CheckBoxFormCallback struct {
	checkbox *models.CheckBox

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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
		case "FormDiv:CheckBoxs":
			// we need to retrieve the field owner before the change
			var pastFormDivOwner *models.FormDiv
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "FormDiv"
			rf.Fieldname = "CheckBoxs"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				checkboxFormCallback.playground.stageOfInterest,
				checkboxFormCallback.playground.backRepoOfInterest,
				checkbox_,
				&rf)

			if reverseFieldOwner != nil {
				pastFormDivOwner = reverseFieldOwner.(*models.FormDiv)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastFormDivOwner != nil {
					idx := slices.Index(pastFormDivOwner.CheckBoxs, checkbox_)
					pastFormDivOwner.CheckBoxs = slices.Delete(pastFormDivOwner.CheckBoxs, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _formdiv := range *models.GetGongstructInstancesSet[models.FormDiv](checkboxFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _formdiv.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newFormDivOwner := _formdiv // we have a match
						if pastFormDivOwner != nil {
							if newFormDivOwner != pastFormDivOwner {
								idx := slices.Index(pastFormDivOwner.CheckBoxs, checkbox_)
								pastFormDivOwner.CheckBoxs = slices.Delete(pastFormDivOwner.CheckBoxs, idx, idx+1)
								newFormDivOwner.CheckBoxs = append(newFormDivOwner.CheckBoxs, checkbox_)
							}
						} else {
							newFormDivOwner.CheckBoxs = append(newFormDivOwner.CheckBoxs, checkbox_)
						}
					}
				}
			}
		}
	}

	checkboxFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.CheckBox](
		checkboxFormCallback.playground,
	)
	checkboxFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if checkboxFormCallback.CreationMode {
		checkboxFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__CheckBoxFormCallback(
				nil,
				checkboxFormCallback.playground,
			),
		}).Stage(checkboxFormCallback.playground.formStage)
		checkbox := new(models.CheckBox)
		FillUpForm(checkbox, newFormGroup, checkboxFormCallback.playground)
		checkboxFormCallback.playground.formStage.Commit()
	}

	fillUpTree(checkboxFormCallback.playground)
}
func __gong__New__DisplayedColumnFormCallback(
	displayedcolumn *models.DisplayedColumn,
	playground *Playground,
) (displayedcolumnFormCallback *DisplayedColumnFormCallback) {
	displayedcolumnFormCallback = new(DisplayedColumnFormCallback)
	displayedcolumnFormCallback.playground = playground
	displayedcolumnFormCallback.displayedcolumn = displayedcolumn

	displayedcolumnFormCallback.CreationMode = (displayedcolumn == nil)

	return
}

type DisplayedColumnFormCallback struct {
	displayedcolumn *models.DisplayedColumn

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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
		case "Table:DisplayedColumns":
			// we need to retrieve the field owner before the change
			var pastTableOwner *models.Table
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Table"
			rf.Fieldname = "DisplayedColumns"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				displayedcolumnFormCallback.playground.stageOfInterest,
				displayedcolumnFormCallback.playground.backRepoOfInterest,
				displayedcolumn_,
				&rf)

			if reverseFieldOwner != nil {
				pastTableOwner = reverseFieldOwner.(*models.Table)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastTableOwner != nil {
					idx := slices.Index(pastTableOwner.DisplayedColumns, displayedcolumn_)
					pastTableOwner.DisplayedColumns = slices.Delete(pastTableOwner.DisplayedColumns, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _table := range *models.GetGongstructInstancesSet[models.Table](displayedcolumnFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _table.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newTableOwner := _table // we have a match
						if pastTableOwner != nil {
							if newTableOwner != pastTableOwner {
								idx := slices.Index(pastTableOwner.DisplayedColumns, displayedcolumn_)
								pastTableOwner.DisplayedColumns = slices.Delete(pastTableOwner.DisplayedColumns, idx, idx+1)
								newTableOwner.DisplayedColumns = append(newTableOwner.DisplayedColumns, displayedcolumn_)
							}
						} else {
							newTableOwner.DisplayedColumns = append(newTableOwner.DisplayedColumns, displayedcolumn_)
						}
					}
				}
			}
		}
	}

	displayedcolumnFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.DisplayedColumn](
		displayedcolumnFormCallback.playground,
	)
	displayedcolumnFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if displayedcolumnFormCallback.CreationMode {
		displayedcolumnFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__DisplayedColumnFormCallback(
				nil,
				displayedcolumnFormCallback.playground,
			),
		}).Stage(displayedcolumnFormCallback.playground.formStage)
		displayedcolumn := new(models.DisplayedColumn)
		FillUpForm(displayedcolumn, newFormGroup, displayedcolumnFormCallback.playground)
		displayedcolumnFormCallback.playground.formStage.Commit()
	}

	fillUpTree(displayedcolumnFormCallback.playground)
}
func __gong__New__FormDivFormCallback(
	formdiv *models.FormDiv,
	playground *Playground,
) (formdivFormCallback *FormDivFormCallback) {
	formdivFormCallback = new(FormDivFormCallback)
	formdivFormCallback.playground = playground
	formdivFormCallback.formdiv = formdiv

	formdivFormCallback.CreationMode = (formdiv == nil)

	return
}

type FormDivFormCallback struct {
	formdiv *models.FormDiv

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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
		case "FormGroup:FormDivs":
			// we need to retrieve the field owner before the change
			var pastFormGroupOwner *models.FormGroup
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "FormGroup"
			rf.Fieldname = "FormDivs"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				formdivFormCallback.playground.stageOfInterest,
				formdivFormCallback.playground.backRepoOfInterest,
				formdiv_,
				&rf)

			if reverseFieldOwner != nil {
				pastFormGroupOwner = reverseFieldOwner.(*models.FormGroup)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastFormGroupOwner != nil {
					idx := slices.Index(pastFormGroupOwner.FormDivs, formdiv_)
					pastFormGroupOwner.FormDivs = slices.Delete(pastFormGroupOwner.FormDivs, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _formgroup := range *models.GetGongstructInstancesSet[models.FormGroup](formdivFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _formgroup.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newFormGroupOwner := _formgroup // we have a match
						if pastFormGroupOwner != nil {
							if newFormGroupOwner != pastFormGroupOwner {
								idx := slices.Index(pastFormGroupOwner.FormDivs, formdiv_)
								pastFormGroupOwner.FormDivs = slices.Delete(pastFormGroupOwner.FormDivs, idx, idx+1)
								newFormGroupOwner.FormDivs = append(newFormGroupOwner.FormDivs, formdiv_)
							}
						} else {
							newFormGroupOwner.FormDivs = append(newFormGroupOwner.FormDivs, formdiv_)
						}
					}
				}
			}
		}
	}

	formdivFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormDiv](
		formdivFormCallback.playground,
	)
	formdivFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if formdivFormCallback.CreationMode {
		formdivFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FormDivFormCallback(
				nil,
				formdivFormCallback.playground,
			),
		}).Stage(formdivFormCallback.playground.formStage)
		formdiv := new(models.FormDiv)
		FillUpForm(formdiv, newFormGroup, formdivFormCallback.playground)
		formdivFormCallback.playground.formStage.Commit()
	}

	fillUpTree(formdivFormCallback.playground)
}
func __gong__New__FormEditAssocButtonFormCallback(
	formeditassocbutton *models.FormEditAssocButton,
	playground *Playground,
) (formeditassocbuttonFormCallback *FormEditAssocButtonFormCallback) {
	formeditassocbuttonFormCallback = new(FormEditAssocButtonFormCallback)
	formeditassocbuttonFormCallback.playground = playground
	formeditassocbuttonFormCallback.formeditassocbutton = formeditassocbutton

	formeditassocbuttonFormCallback.CreationMode = (formeditassocbutton == nil)

	return
}

type FormEditAssocButtonFormCallback struct {
	formeditassocbutton *models.FormEditAssocButton

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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

	// display a new form by reset the form stage
	if formeditassocbuttonFormCallback.CreationMode {
		formeditassocbuttonFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FormEditAssocButtonFormCallback(
				nil,
				formeditassocbuttonFormCallback.playground,
			),
		}).Stage(formeditassocbuttonFormCallback.playground.formStage)
		formeditassocbutton := new(models.FormEditAssocButton)
		FillUpForm(formeditassocbutton, newFormGroup, formeditassocbuttonFormCallback.playground)
		formeditassocbuttonFormCallback.playground.formStage.Commit()
	}

	fillUpTree(formeditassocbuttonFormCallback.playground)
}
func __gong__New__FormFieldFormCallback(
	formfield *models.FormField,
	playground *Playground,
) (formfieldFormCallback *FormFieldFormCallback) {
	formfieldFormCallback = new(FormFieldFormCallback)
	formfieldFormCallback.playground = playground
	formfieldFormCallback.formfield = formfield

	formfieldFormCallback.CreationMode = (formfield == nil)

	return
}

type FormFieldFormCallback struct {
	formfield *models.FormField

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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
		case "FormDiv:FormFields":
			// we need to retrieve the field owner before the change
			var pastFormDivOwner *models.FormDiv
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "FormDiv"
			rf.Fieldname = "FormFields"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				formfieldFormCallback.playground.stageOfInterest,
				formfieldFormCallback.playground.backRepoOfInterest,
				formfield_,
				&rf)

			if reverseFieldOwner != nil {
				pastFormDivOwner = reverseFieldOwner.(*models.FormDiv)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastFormDivOwner != nil {
					idx := slices.Index(pastFormDivOwner.FormFields, formfield_)
					pastFormDivOwner.FormFields = slices.Delete(pastFormDivOwner.FormFields, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _formdiv := range *models.GetGongstructInstancesSet[models.FormDiv](formfieldFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _formdiv.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newFormDivOwner := _formdiv // we have a match
						if pastFormDivOwner != nil {
							if newFormDivOwner != pastFormDivOwner {
								idx := slices.Index(pastFormDivOwner.FormFields, formfield_)
								pastFormDivOwner.FormFields = slices.Delete(pastFormDivOwner.FormFields, idx, idx+1)
								newFormDivOwner.FormFields = append(newFormDivOwner.FormFields, formfield_)
							}
						} else {
							newFormDivOwner.FormFields = append(newFormDivOwner.FormFields, formfield_)
						}
					}
				}
			}
		}
	}

	formfieldFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormField](
		formfieldFormCallback.playground,
	)
	formfieldFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if formfieldFormCallback.CreationMode {
		formfieldFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FormFieldFormCallback(
				nil,
				formfieldFormCallback.playground,
			),
		}).Stage(formfieldFormCallback.playground.formStage)
		formfield := new(models.FormField)
		FillUpForm(formfield, newFormGroup, formfieldFormCallback.playground)
		formfieldFormCallback.playground.formStage.Commit()
	}

	fillUpTree(formfieldFormCallback.playground)
}
func __gong__New__FormFieldDateFormCallback(
	formfielddate *models.FormFieldDate,
	playground *Playground,
) (formfielddateFormCallback *FormFieldDateFormCallback) {
	formfielddateFormCallback = new(FormFieldDateFormCallback)
	formfielddateFormCallback.playground = playground
	formfielddateFormCallback.formfielddate = formfielddate

	formfielddateFormCallback.CreationMode = (formfielddate == nil)

	return
}

type FormFieldDateFormCallback struct {
	formfielddate *models.FormFieldDate

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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

	// display a new form by reset the form stage
	if formfielddateFormCallback.CreationMode {
		formfielddateFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FormFieldDateFormCallback(
				nil,
				formfielddateFormCallback.playground,
			),
		}).Stage(formfielddateFormCallback.playground.formStage)
		formfielddate := new(models.FormFieldDate)
		FillUpForm(formfielddate, newFormGroup, formfielddateFormCallback.playground)
		formfielddateFormCallback.playground.formStage.Commit()
	}

	fillUpTree(formfielddateFormCallback.playground)
}
func __gong__New__FormFieldDateTimeFormCallback(
	formfielddatetime *models.FormFieldDateTime,
	playground *Playground,
) (formfielddatetimeFormCallback *FormFieldDateTimeFormCallback) {
	formfielddatetimeFormCallback = new(FormFieldDateTimeFormCallback)
	formfielddatetimeFormCallback.playground = playground
	formfielddatetimeFormCallback.formfielddatetime = formfielddatetime

	formfielddatetimeFormCallback.CreationMode = (formfielddatetime == nil)

	return
}

type FormFieldDateTimeFormCallback struct {
	formfielddatetime *models.FormFieldDateTime

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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

	// display a new form by reset the form stage
	if formfielddatetimeFormCallback.CreationMode {
		formfielddatetimeFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FormFieldDateTimeFormCallback(
				nil,
				formfielddatetimeFormCallback.playground,
			),
		}).Stage(formfielddatetimeFormCallback.playground.formStage)
		formfielddatetime := new(models.FormFieldDateTime)
		FillUpForm(formfielddatetime, newFormGroup, formfielddatetimeFormCallback.playground)
		formfielddatetimeFormCallback.playground.formStage.Commit()
	}

	fillUpTree(formfielddatetimeFormCallback.playground)
}
func __gong__New__FormFieldFloat64FormCallback(
	formfieldfloat64 *models.FormFieldFloat64,
	playground *Playground,
) (formfieldfloat64FormCallback *FormFieldFloat64FormCallback) {
	formfieldfloat64FormCallback = new(FormFieldFloat64FormCallback)
	formfieldfloat64FormCallback.playground = playground
	formfieldfloat64FormCallback.formfieldfloat64 = formfieldfloat64

	formfieldfloat64FormCallback.CreationMode = (formfieldfloat64 == nil)

	return
}

type FormFieldFloat64FormCallback struct {
	formfieldfloat64 *models.FormFieldFloat64

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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

	// display a new form by reset the form stage
	if formfieldfloat64FormCallback.CreationMode {
		formfieldfloat64FormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FormFieldFloat64FormCallback(
				nil,
				formfieldfloat64FormCallback.playground,
			),
		}).Stage(formfieldfloat64FormCallback.playground.formStage)
		formfieldfloat64 := new(models.FormFieldFloat64)
		FillUpForm(formfieldfloat64, newFormGroup, formfieldfloat64FormCallback.playground)
		formfieldfloat64FormCallback.playground.formStage.Commit()
	}

	fillUpTree(formfieldfloat64FormCallback.playground)
}
func __gong__New__FormFieldIntFormCallback(
	formfieldint *models.FormFieldInt,
	playground *Playground,
) (formfieldintFormCallback *FormFieldIntFormCallback) {
	formfieldintFormCallback = new(FormFieldIntFormCallback)
	formfieldintFormCallback.playground = playground
	formfieldintFormCallback.formfieldint = formfieldint

	formfieldintFormCallback.CreationMode = (formfieldint == nil)

	return
}

type FormFieldIntFormCallback struct {
	formfieldint *models.FormFieldInt

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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

	// display a new form by reset the form stage
	if formfieldintFormCallback.CreationMode {
		formfieldintFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FormFieldIntFormCallback(
				nil,
				formfieldintFormCallback.playground,
			),
		}).Stage(formfieldintFormCallback.playground.formStage)
		formfieldint := new(models.FormFieldInt)
		FillUpForm(formfieldint, newFormGroup, formfieldintFormCallback.playground)
		formfieldintFormCallback.playground.formStage.Commit()
	}

	fillUpTree(formfieldintFormCallback.playground)
}
func __gong__New__FormFieldSelectFormCallback(
	formfieldselect *models.FormFieldSelect,
	playground *Playground,
) (formfieldselectFormCallback *FormFieldSelectFormCallback) {
	formfieldselectFormCallback = new(FormFieldSelectFormCallback)
	formfieldselectFormCallback.playground = playground
	formfieldselectFormCallback.formfieldselect = formfieldselect

	formfieldselectFormCallback.CreationMode = (formfieldselect == nil)

	return
}

type FormFieldSelectFormCallback struct {
	formfieldselect *models.FormFieldSelect

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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

	// display a new form by reset the form stage
	if formfieldselectFormCallback.CreationMode {
		formfieldselectFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FormFieldSelectFormCallback(
				nil,
				formfieldselectFormCallback.playground,
			),
		}).Stage(formfieldselectFormCallback.playground.formStage)
		formfieldselect := new(models.FormFieldSelect)
		FillUpForm(formfieldselect, newFormGroup, formfieldselectFormCallback.playground)
		formfieldselectFormCallback.playground.formStage.Commit()
	}

	fillUpTree(formfieldselectFormCallback.playground)
}
func __gong__New__FormFieldStringFormCallback(
	formfieldstring *models.FormFieldString,
	playground *Playground,
) (formfieldstringFormCallback *FormFieldStringFormCallback) {
	formfieldstringFormCallback = new(FormFieldStringFormCallback)
	formfieldstringFormCallback.playground = playground
	formfieldstringFormCallback.formfieldstring = formfieldstring

	formfieldstringFormCallback.CreationMode = (formfieldstring == nil)

	return
}

type FormFieldStringFormCallback struct {
	formfieldstring *models.FormFieldString

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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
		case "IsTextArea":
			FormDivBasicFieldToField(&(formfieldstring_.IsTextArea), formDiv)
		}
	}

	formfieldstringFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormFieldString](
		formfieldstringFormCallback.playground,
	)
	formfieldstringFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if formfieldstringFormCallback.CreationMode {
		formfieldstringFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FormFieldStringFormCallback(
				nil,
				formfieldstringFormCallback.playground,
			),
		}).Stage(formfieldstringFormCallback.playground.formStage)
		formfieldstring := new(models.FormFieldString)
		FillUpForm(formfieldstring, newFormGroup, formfieldstringFormCallback.playground)
		formfieldstringFormCallback.playground.formStage.Commit()
	}

	fillUpTree(formfieldstringFormCallback.playground)
}
func __gong__New__FormFieldTimeFormCallback(
	formfieldtime *models.FormFieldTime,
	playground *Playground,
) (formfieldtimeFormCallback *FormFieldTimeFormCallback) {
	formfieldtimeFormCallback = new(FormFieldTimeFormCallback)
	formfieldtimeFormCallback.playground = playground
	formfieldtimeFormCallback.formfieldtime = formfieldtime

	formfieldtimeFormCallback.CreationMode = (formfieldtime == nil)

	return
}

type FormFieldTimeFormCallback struct {
	formfieldtime *models.FormFieldTime

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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

	// display a new form by reset the form stage
	if formfieldtimeFormCallback.CreationMode {
		formfieldtimeFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FormFieldTimeFormCallback(
				nil,
				formfieldtimeFormCallback.playground,
			),
		}).Stage(formfieldtimeFormCallback.playground.formStage)
		formfieldtime := new(models.FormFieldTime)
		FillUpForm(formfieldtime, newFormGroup, formfieldtimeFormCallback.playground)
		formfieldtimeFormCallback.playground.formStage.Commit()
	}

	fillUpTree(formfieldtimeFormCallback.playground)
}
func __gong__New__FormGroupFormCallback(
	formgroup *models.FormGroup,
	playground *Playground,
) (formgroupFormCallback *FormGroupFormCallback) {
	formgroupFormCallback = new(FormGroupFormCallback)
	formgroupFormCallback.playground = playground
	formgroupFormCallback.formgroup = formgroup

	formgroupFormCallback.CreationMode = (formgroup == nil)

	return
}

type FormGroupFormCallback struct {
	formgroup *models.FormGroup

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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
		case "Label":
			FormDivBasicFieldToField(&(formgroup_.Label), formDiv)
		}
	}

	formgroupFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormGroup](
		formgroupFormCallback.playground,
	)
	formgroupFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if formgroupFormCallback.CreationMode {
		formgroupFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FormGroupFormCallback(
				nil,
				formgroupFormCallback.playground,
			),
		}).Stage(formgroupFormCallback.playground.formStage)
		formgroup := new(models.FormGroup)
		FillUpForm(formgroup, newFormGroup, formgroupFormCallback.playground)
		formgroupFormCallback.playground.formStage.Commit()
	}

	fillUpTree(formgroupFormCallback.playground)
}
func __gong__New__FormSortAssocButtonFormCallback(
	formsortassocbutton *models.FormSortAssocButton,
	playground *Playground,
) (formsortassocbuttonFormCallback *FormSortAssocButtonFormCallback) {
	formsortassocbuttonFormCallback = new(FormSortAssocButtonFormCallback)
	formsortassocbuttonFormCallback.playground = playground
	formsortassocbuttonFormCallback.formsortassocbutton = formsortassocbutton

	formsortassocbuttonFormCallback.CreationMode = (formsortassocbutton == nil)

	return
}

type FormSortAssocButtonFormCallback struct {
	formsortassocbutton *models.FormSortAssocButton

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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

	// display a new form by reset the form stage
	if formsortassocbuttonFormCallback.CreationMode {
		formsortassocbuttonFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FormSortAssocButtonFormCallback(
				nil,
				formsortassocbuttonFormCallback.playground,
			),
		}).Stage(formsortassocbuttonFormCallback.playground.formStage)
		formsortassocbutton := new(models.FormSortAssocButton)
		FillUpForm(formsortassocbutton, newFormGroup, formsortassocbuttonFormCallback.playground)
		formsortassocbuttonFormCallback.playground.formStage.Commit()
	}

	fillUpTree(formsortassocbuttonFormCallback.playground)
}
func __gong__New__OptionFormCallback(
	option *models.Option,
	playground *Playground,
) (optionFormCallback *OptionFormCallback) {
	optionFormCallback = new(OptionFormCallback)
	optionFormCallback.playground = playground
	optionFormCallback.option = option

	optionFormCallback.CreationMode = (option == nil)

	return
}

type OptionFormCallback struct {
	option *models.Option

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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
		case "FormFieldSelect:Options":
			// we need to retrieve the field owner before the change
			var pastFormFieldSelectOwner *models.FormFieldSelect
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "FormFieldSelect"
			rf.Fieldname = "Options"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				optionFormCallback.playground.stageOfInterest,
				optionFormCallback.playground.backRepoOfInterest,
				option_,
				&rf)

			if reverseFieldOwner != nil {
				pastFormFieldSelectOwner = reverseFieldOwner.(*models.FormFieldSelect)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastFormFieldSelectOwner != nil {
					idx := slices.Index(pastFormFieldSelectOwner.Options, option_)
					pastFormFieldSelectOwner.Options = slices.Delete(pastFormFieldSelectOwner.Options, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _formfieldselect := range *models.GetGongstructInstancesSet[models.FormFieldSelect](optionFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _formfieldselect.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newFormFieldSelectOwner := _formfieldselect // we have a match
						if pastFormFieldSelectOwner != nil {
							if newFormFieldSelectOwner != pastFormFieldSelectOwner {
								idx := slices.Index(pastFormFieldSelectOwner.Options, option_)
								pastFormFieldSelectOwner.Options = slices.Delete(pastFormFieldSelectOwner.Options, idx, idx+1)
								newFormFieldSelectOwner.Options = append(newFormFieldSelectOwner.Options, option_)
							}
						} else {
							newFormFieldSelectOwner.Options = append(newFormFieldSelectOwner.Options, option_)
						}
					}
				}
			}
		}
	}

	optionFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Option](
		optionFormCallback.playground,
	)
	optionFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if optionFormCallback.CreationMode {
		optionFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__OptionFormCallback(
				nil,
				optionFormCallback.playground,
			),
		}).Stage(optionFormCallback.playground.formStage)
		option := new(models.Option)
		FillUpForm(option, newFormGroup, optionFormCallback.playground)
		optionFormCallback.playground.formStage.Commit()
	}

	fillUpTree(optionFormCallback.playground)
}
func __gong__New__RowFormCallback(
	row *models.Row,
	playground *Playground,
) (rowFormCallback *RowFormCallback) {
	rowFormCallback = new(RowFormCallback)
	rowFormCallback.playground = playground
	rowFormCallback.row = row

	rowFormCallback.CreationMode = (row == nil)

	return
}

type RowFormCallback struct {
	row *models.Row

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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
		case "Table:Rows":
			// we need to retrieve the field owner before the change
			var pastTableOwner *models.Table
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Table"
			rf.Fieldname = "Rows"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				rowFormCallback.playground.stageOfInterest,
				rowFormCallback.playground.backRepoOfInterest,
				row_,
				&rf)

			if reverseFieldOwner != nil {
				pastTableOwner = reverseFieldOwner.(*models.Table)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastTableOwner != nil {
					idx := slices.Index(pastTableOwner.Rows, row_)
					pastTableOwner.Rows = slices.Delete(pastTableOwner.Rows, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _table := range *models.GetGongstructInstancesSet[models.Table](rowFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _table.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newTableOwner := _table // we have a match
						if pastTableOwner != nil {
							if newTableOwner != pastTableOwner {
								idx := slices.Index(pastTableOwner.Rows, row_)
								pastTableOwner.Rows = slices.Delete(pastTableOwner.Rows, idx, idx+1)
								newTableOwner.Rows = append(newTableOwner.Rows, row_)
							}
						} else {
							newTableOwner.Rows = append(newTableOwner.Rows, row_)
						}
					}
				}
			}
		}
	}

	rowFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Row](
		rowFormCallback.playground,
	)
	rowFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if rowFormCallback.CreationMode {
		rowFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__RowFormCallback(
				nil,
				rowFormCallback.playground,
			),
		}).Stage(rowFormCallback.playground.formStage)
		row := new(models.Row)
		FillUpForm(row, newFormGroup, rowFormCallback.playground)
		rowFormCallback.playground.formStage.Commit()
	}

	fillUpTree(rowFormCallback.playground)
}
func __gong__New__TableFormCallback(
	table *models.Table,
	playground *Playground,
) (tableFormCallback *TableFormCallback) {
	tableFormCallback = new(TableFormCallback)
	tableFormCallback.playground = playground
	tableFormCallback.table = table

	tableFormCallback.CreationMode = (table == nil)

	return
}

type TableFormCallback struct {
	table *models.Table

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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

	// display a new form by reset the form stage
	if tableFormCallback.CreationMode {
		tableFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__TableFormCallback(
				nil,
				tableFormCallback.playground,
			),
		}).Stage(tableFormCallback.playground.formStage)
		table := new(models.Table)
		FillUpForm(table, newFormGroup, tableFormCallback.playground)
		tableFormCallback.playground.formStage.Commit()
	}

	fillUpTree(tableFormCallback.playground)
}
