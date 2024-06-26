package main

import (
	"time"

	"github.com/fullstack-lang/gongtable/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_form_stage_issue30 models.StageStruct
var ___dummy__Time_form_stage_issue30 time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_form_stage_issue30 map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["form_stage_issue30"] = form_stage_issue30Injection
// }

// form_stage_issue30Injection will stage objects of database "form_stage_issue30"
func form_stage_issue30Injection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Cell

	// Declarations of staged instances of CellBoolean

	// Declarations of staged instances of CellFloat64

	// Declarations of staged instances of CellIcon

	// Declarations of staged instances of CellInt

	// Declarations of staged instances of CellString

	// Declarations of staged instances of CheckBox

	// Declarations of staged instances of DisplayedColumn

	// Declarations of staged instances of FormDiv
	__FormDiv__000000_Date := (&models.FormDiv{Name: `Date`}).Stage(stage)
	__FormDiv__000001_DateTime := (&models.FormDiv{Name: `DateTime`}).Stage(stage)

	// Declarations of staged instances of FormEditAssocButton

	// Declarations of staged instances of FormField
	__FormField__000000_Date := (&models.FormField{Name: `Date`}).Stage(stage)
	__FormField__000001_DateTime := (&models.FormField{Name: `DateTime`}).Stage(stage)

	// Declarations of staged instances of FormFieldDate
	__FormFieldDate__000000_Date := (&models.FormFieldDate{Name: `Date`}).Stage(stage)

	// Declarations of staged instances of FormFieldDateTime
	__FormFieldDateTime__000000_FormFieldDateTime := (&models.FormFieldDateTime{Name: `FormFieldDateTime`}).Stage(stage)

	// Declarations of staged instances of FormFieldFloat64

	// Declarations of staged instances of FormFieldInt

	// Declarations of staged instances of FormFieldSelect

	// Declarations of staged instances of FormFieldString

	// Declarations of staged instances of FormFieldTime

	// Declarations of staged instances of FormGroup
	__FormGroup__000000_Form_1 := (&models.FormGroup{Name: `Form 1`}).Stage(stage)

	// Declarations of staged instances of FormSortAssocButton

	// Declarations of staged instances of Option

	// Declarations of staged instances of Row

	// Declarations of staged instances of Table

	// Setup of values

	// FormDiv values setup
	__FormDiv__000000_Date.Name = `Date`

	// FormDiv values setup
	__FormDiv__000001_DateTime.Name = `DateTime`

	// FormField values setup
	__FormField__000000_Date.Name = `Date`
	__FormField__000000_Date.Label = ``
	__FormField__000000_Date.Placeholder = ``
	__FormField__000000_Date.HasBespokeWidth = false
	__FormField__000000_Date.BespokeWidthPx = 0
	__FormField__000000_Date.HasBespokeHeight = false
	__FormField__000000_Date.BespokeHeightPx = 0

	// FormField values setup
	__FormField__000001_DateTime.Name = `DateTime`
	__FormField__000001_DateTime.Label = `DateTime`
	__FormField__000001_DateTime.Placeholder = `DateTime`
	__FormField__000001_DateTime.HasBespokeWidth = false
	__FormField__000001_DateTime.BespokeWidthPx = 0
	__FormField__000001_DateTime.HasBespokeHeight = false
	__FormField__000001_DateTime.BespokeHeightPx = 0

	// FormFieldDate values setup
	__FormFieldDate__000000_Date.Name = `Date`
	__FormFieldDate__000000_Date.Value, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2020-10-10 00:00:00 +0000 UTC")

	// FormFieldDateTime values setup
	__FormFieldDateTime__000000_FormFieldDateTime.Name = `FormFieldDateTime`
	__FormFieldDateTime__000000_FormFieldDateTime.Value, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2024-03-29 03:52:43.214 +0000 UTC")

	// FormGroup values setup
	__FormGroup__000000_Form_1.Name = `Form 1`
	__FormGroup__000000_Form_1.Label = `Form 1`
	__FormGroup__000000_Form_1.HasSuppressButton = false
	__FormGroup__000000_Form_1.HasSuppressButtonBeenPressed = false

	// Setup of pointers
	__FormDiv__000000_Date.FormFields = append(__FormDiv__000000_Date.FormFields, __FormField__000000_Date)
	__FormDiv__000001_DateTime.FormFields = append(__FormDiv__000001_DateTime.FormFields, __FormField__000001_DateTime)
	__FormField__000000_Date.FormFieldDate = __FormFieldDate__000000_Date
	__FormField__000001_DateTime.FormFieldDateTime = __FormFieldDateTime__000000_FormFieldDateTime
	__FormGroup__000000_Form_1.FormDivs = append(__FormGroup__000000_Form_1.FormDivs, __FormDiv__000001_DateTime)
	__FormGroup__000000_Form_1.FormDivs = append(__FormGroup__000000_Form_1.FormDivs, __FormDiv__000000_Date)
}


