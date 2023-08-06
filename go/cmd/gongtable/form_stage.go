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
	__FormDiv__000001_Form_Div_First_Name_Name := (&models.FormDiv{Name: `Form Div First Name  - Name`}).Stage(stage)

	// Declarations of staged instances of FormField
	__FormField__000000_Age := (&models.FormField{Name: `Age`}).Stage(stage)
	__FormField__000001_FirstName := (&models.FormField{Name: `FirstName`}).Stage(stage)
	__FormField__000002_LastName := (&models.FormField{Name: `LastName`}).Stage(stage)

	// Declarations of staged instances of FormFieldBoolean

	// Declarations of staged instances of FormFieldFloat64

	// Declarations of staged instances of FormFieldInt
	__FormFieldInt__000000_Age := (&models.FormFieldInt{Name: `Age`}).Stage(stage)

	// Declarations of staged instances of FormFieldString
	__FormFieldString__000000_FirstName := (&models.FormFieldString{Name: `FirstName`}).Stage(stage)
	__FormFieldString__000001_LastName := (&models.FormFieldString{Name: `LastName`}).Stage(stage)

	// Declarations of staged instances of FormGroup
	__FormGroup__000000_Form_1 := (&models.FormGroup{Name: `Form 1`}).Stage(stage)

	// Declarations of staged instances of Row

	// Declarations of staged instances of Table

	// Setup of values

	// FormDiv values setup
	__FormDiv__000000_Age.Name = `Age`

	// FormDiv values setup
	__FormDiv__000001_Form_Div_First_Name_Name.Name = `Form Div First Name  - Name`

	// FormField values setup
	__FormField__000000_Age.Name = `Age`
	__FormField__000000_Age.InputTypeEnum = models.Number
	__FormField__000000_Age.Label = `Age`
	__FormField__000000_Age.Placeholder = `18`

	// FormField values setup
	__FormField__000001_FirstName.Name = `FirstName`
	__FormField__000001_FirstName.InputTypeEnum = models.Text
	__FormField__000001_FirstName.Label = `First Name (label)`
	__FormField__000001_FirstName.Placeholder = ``

	// FormField values setup
	__FormField__000002_LastName.Name = `LastName`
	__FormField__000002_LastName.InputTypeEnum = models.Text
	__FormField__000002_LastName.Label = `Last Name (label)`
	__FormField__000002_LastName.Placeholder = ``

	// FormFieldInt values setup
	__FormFieldInt__000000_Age.Name = `Age`
	__FormFieldInt__000000_Age.Value = 1227

	// FormFieldString values setup
	__FormFieldString__000000_FirstName.Name = `FirstName`
	__FormFieldString__000000_FirstName.Value = `charles`

	// FormFieldString values setup
	__FormFieldString__000001_LastName.Name = `LastName`
	__FormFieldString__000001_LastName.Value = `Bau`

	// FormGroup values setup
	__FormGroup__000000_Form_1.Name = `Form 1`

	// Setup of pointers
	__FormDiv__000000_Age.FormFields = append(__FormDiv__000000_Age.FormFields, __FormField__000000_Age)
	__FormDiv__000001_Form_Div_First_Name_Name.FormFields = append(__FormDiv__000001_Form_Div_First_Name_Name.FormFields, __FormField__000001_FirstName)
	__FormDiv__000001_Form_Div_First_Name_Name.FormFields = append(__FormDiv__000001_Form_Div_First_Name_Name.FormFields, __FormField__000002_LastName)
	__FormField__000000_Age.FormFieldInt = __FormFieldInt__000000_Age
	__FormField__000001_FirstName.FormFieldString = __FormFieldString__000000_FirstName
	__FormField__000002_LastName.FormFieldString = __FormFieldString__000001_LastName
	__FormGroup__000000_Form_1.FormDivs = append(__FormGroup__000000_Form_1.FormDivs, __FormDiv__000001_Form_Div_First_Name_Name)
	__FormGroup__000000_Form_1.FormDivs = append(__FormGroup__000000_Form_1.FormDivs, __FormDiv__000000_Age)
}


