package main

import (
	"time"

	"github.com/fullstack-lang/gongtable/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_form_stage models.StageStruct
var ___dummy__Time_form_stage time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_form_stage map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["form_stage"] = form_stageInjection
// }

// form_stageInjection will stage objects of database "form_stage"
func form_stageInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Cell

	// Declarations of staged instances of CellBoolean

	// Declarations of staged instances of CellFloat64

	// Declarations of staged instances of CellIcon

	// Declarations of staged instances of CellInt

	// Declarations of staged instances of CellString

	// Declarations of staged instances of DisplayedColumn

	// Declarations of staged instances of FormDiv
	__FormDiv__000000_Age := (&models.FormDiv{Name: `Age`}).Stage(stage)
	__FormDiv__000001_Date_Time := (&models.FormDiv{Name: `Date - Time`}).Stage(stage)
	__FormDiv__000002_DateTime := (&models.FormDiv{Name: `DateTime`}).Stage(stage)
	__FormDiv__000003_Form_Div_First_Name_Name := (&models.FormDiv{Name: `Form Div First Name  - Name`}).Stage(stage)

	// Declarations of staged instances of FormField
	__FormField__000000_Age := (&models.FormField{Name: `Age`}).Stage(stage)
	__FormField__000001_Date := (&models.FormField{Name: `Date`}).Stage(stage)
	__FormField__000002_DateTime := (&models.FormField{Name: `DateTime`}).Stage(stage)
	__FormField__000003_FirstName := (&models.FormField{Name: `FirstName`}).Stage(stage)
	__FormField__000004_LastName := (&models.FormField{Name: `LastName`}).Stage(stage)
	__FormField__000005_Time := (&models.FormField{Name: `Time`}).Stage(stage)

	// Declarations of staged instances of FormFieldBoolean

	// Declarations of staged instances of FormFieldDate
	__FormFieldDate__000000_Time := (&models.FormFieldDate{Name: `Time`}).Stage(stage)

	// Declarations of staged instances of FormFieldDateTime
	__FormFieldDateTime__000000_DateTime := (&models.FormFieldDateTime{Name: `DateTime`}).Stage(stage)

	// Declarations of staged instances of FormFieldFloat64

	// Declarations of staged instances of FormFieldInt
	__FormFieldInt__000000_Age := (&models.FormFieldInt{Name: `Age`}).Stage(stage)

	// Declarations of staged instances of FormFieldString
	__FormFieldString__000000_FirstName := (&models.FormFieldString{Name: `FirstName`}).Stage(stage)
	__FormFieldString__000001_LastName := (&models.FormFieldString{Name: `LastName`}).Stage(stage)

	// Declarations of staged instances of FormFieldTime
	__FormFieldTime__000000_Time := (&models.FormFieldTime{Name: `Time`}).Stage(stage)

	// Declarations of staged instances of FormGroup
	__FormGroup__000000_Form_1 := (&models.FormGroup{Name: `Form 1`}).Stage(stage)

	// Declarations of staged instances of Row

	// Declarations of staged instances of Table

	// Setup of values

	// FormDiv values setup
	__FormDiv__000000_Age.Name = `Age`

	// FormDiv values setup
	__FormDiv__000001_Date_Time.Name = `Date - Time`

	// FormDiv values setup
	__FormDiv__000002_DateTime.Name = `DateTime`

	// FormDiv values setup
	__FormDiv__000003_Form_Div_First_Name_Name.Name = `Form Div First Name  - Name`

	// FormField values setup
	__FormField__000000_Age.Name = `Age`
	__FormField__000000_Age.InputTypeEnum = models.Number
	__FormField__000000_Age.Label = `Age`
	__FormField__000000_Age.Placeholder = `18`

	// FormField values setup
	__FormField__000001_Date.Name = `Date`
	__FormField__000001_Date.InputTypeEnum = models.Date
	__FormField__000001_Date.Label = `Date`
	__FormField__000001_Date.Placeholder = ``

	// FormField values setup
	__FormField__000002_DateTime.Name = `DateTime`
	__FormField__000002_DateTime.Label = ``
	__FormField__000002_DateTime.Placeholder = ``

	// FormField values setup
	__FormField__000003_FirstName.Name = `FirstName`
	__FormField__000003_FirstName.InputTypeEnum = models.Text
	__FormField__000003_FirstName.Label = `First Name (label)`
	__FormField__000003_FirstName.Placeholder = ``

	// FormField values setup
	__FormField__000004_LastName.Name = `LastName`
	__FormField__000004_LastName.InputTypeEnum = models.Text
	__FormField__000004_LastName.Label = `Last Name (label)`
	__FormField__000004_LastName.Placeholder = ``

	// FormField values setup
	__FormField__000005_Time.Name = `Time`
	__FormField__000005_Time.InputTypeEnum = models.Time
	__FormField__000005_Time.Label = `Time`
	__FormField__000005_Time.Placeholder = ``

	// FormFieldDate values setup
	__FormFieldDate__000000_Time.Name = `Time`
	__FormFieldDate__000000_Time.Value, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-03-02 00:00:00 +0000 UTC")

	// FormFieldDateTime values setup
	__FormFieldDateTime__000000_DateTime.Name = `DateTime`
	__FormFieldDateTime__000000_DateTime.Value, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2001-01-01 01:02:03 +0000 UTC")

	// FormFieldInt values setup
	__FormFieldInt__000000_Age.Name = `Age`
	__FormFieldInt__000000_Age.Value = 1227

	// FormFieldString values setup
	__FormFieldString__000000_FirstName.Name = `FirstName`
	__FormFieldString__000000_FirstName.Value = `charles`

	// FormFieldString values setup
	__FormFieldString__000001_LastName.Name = `LastName`
	__FormFieldString__000001_LastName.Value = `Bau`

	// FormFieldTime values setup
	__FormFieldTime__000000_Time.Name = `Time`
	__FormFieldTime__000000_Time.Value, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1970-01-01 12:03:06 +0000 UTC")
	__FormFieldTime__000000_Time.Step = 1.000000

	// FormGroup values setup
	__FormGroup__000000_Form_1.Name = `Form 1`

	// Setup of pointers
	__FormDiv__000000_Age.FormFields = append(__FormDiv__000000_Age.FormFields, __FormField__000000_Age)
	__FormDiv__000001_Date_Time.FormFields = append(__FormDiv__000001_Date_Time.FormFields, __FormField__000001_Date)
	__FormDiv__000001_Date_Time.FormFields = append(__FormDiv__000001_Date_Time.FormFields, __FormField__000005_Time)
	__FormDiv__000002_DateTime.FormFields = append(__FormDiv__000002_DateTime.FormFields, __FormField__000002_DateTime)
	__FormDiv__000003_Form_Div_First_Name_Name.FormFields = append(__FormDiv__000003_Form_Div_First_Name_Name.FormFields, __FormField__000003_FirstName)
	__FormDiv__000003_Form_Div_First_Name_Name.FormFields = append(__FormDiv__000003_Form_Div_First_Name_Name.FormFields, __FormField__000004_LastName)
	__FormField__000000_Age.FormFieldInt = __FormFieldInt__000000_Age
	__FormField__000001_Date.FormFieldDate = __FormFieldDate__000000_Time
	__FormField__000002_DateTime.FormFieldDateTime = __FormFieldDateTime__000000_DateTime
	__FormField__000003_FirstName.FormFieldString = __FormFieldString__000000_FirstName
	__FormField__000004_LastName.FormFieldString = __FormFieldString__000001_LastName
	__FormField__000005_Time.FormFieldTime = __FormFieldTime__000000_Time
	__FormGroup__000000_Form_1.FormDivs = append(__FormGroup__000000_Form_1.FormDivs, __FormDiv__000003_Form_Div_First_Name_Name)
	__FormGroup__000000_Form_1.FormDivs = append(__FormGroup__000000_Form_1.FormDivs, __FormDiv__000000_Age)
	__FormGroup__000000_Form_1.FormDivs = append(__FormGroup__000000_Form_1.FormDivs, __FormDiv__000001_Date_Time)
	__FormGroup__000000_Form_1.FormDivs = append(__FormGroup__000000_Form_1.FormDivs, __FormDiv__000002_DateTime)
}


