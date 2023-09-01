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
	__CellString__000000_A := (&models.CellString{Name: `A`}).Stage(stage)

	// Declarations of staged instances of CheckBox
	__CheckBox__000000_Boolean := (&models.CheckBox{Name: `Boolean`}).Stage(stage)

	// Declarations of staged instances of DisplayedColumn

	// Declarations of staged instances of FormDiv
	__FormDiv__000000_CheckBoxs := (&models.FormDiv{Name: `CheckBoxs`}).Stage(stage)
	__FormDiv__000001_Date_Time := (&models.FormDiv{Name: `Date - Time`}).Stage(stage)
	__FormDiv__000002_DateTime := (&models.FormDiv{Name: `DateTime`}).Stage(stage)
	__FormDiv__000003_Edit_Assoc := (&models.FormDiv{Name: `Edit Assoc`}).Stage(stage)
	__FormDiv__000004_Float64_level_of_x := (&models.FormDiv{Name: `Float64 - level of x`}).Stage(stage)
	__FormDiv__000005_Form_Div_First_Name_Name := (&models.FormDiv{Name: `Form Div First Name  - Name`}).Stage(stage)
	__FormDiv__000006_Int_Age := (&models.FormDiv{Name: `Int - Age`}).Stage(stage)
	__FormDiv__000007_Select := (&models.FormDiv{Name: `Select`}).Stage(stage)

	// Declarations of staged instances of FormEditAssocButton
	__FormEditAssocButton__000000_Edit_Assoc := (&models.FormEditAssocButton{Name: `Edit Assoc`}).Stage(stage)

	// Declarations of staged instances of FormField
	__FormField__000000_Age := (&models.FormField{Name: `Age`}).Stage(stage)
	__FormField__000001_Date := (&models.FormField{Name: `Date`}).Stage(stage)
	__FormField__000002_DateTime := (&models.FormField{Name: `DateTime`}).Stage(stage)
	__FormField__000003_FirstName := (&models.FormField{Name: `FirstName`}).Stage(stage)
	__FormField__000004_Float64_level_of_x := (&models.FormField{Name: `Float64 - level of x`}).Stage(stage)
	__FormField__000005_LastName := (&models.FormField{Name: `LastName`}).Stage(stage)
	__FormField__000006_Select := (&models.FormField{Name: `Select`}).Stage(stage)
	__FormField__000007_Time := (&models.FormField{Name: `Time`}).Stage(stage)

	// Declarations of staged instances of FormFieldDate
	__FormFieldDate__000000_Time := (&models.FormFieldDate{Name: `Time`}).Stage(stage)

	// Declarations of staged instances of FormFieldDateTime
	__FormFieldDateTime__000000_DateTime := (&models.FormFieldDateTime{Name: `DateTime`}).Stage(stage)

	// Declarations of staged instances of FormFieldFloat64
	__FormFieldFloat64__000000_Float64_level_of_x := (&models.FormFieldFloat64{Name: `Float64 - level of x`}).Stage(stage)

	// Declarations of staged instances of FormFieldInt
	__FormFieldInt__000000_Age := (&models.FormFieldInt{Name: `Age`}).Stage(stage)

	// Declarations of staged instances of FormFieldSelect
	__FormFieldSelect__000000_Select := (&models.FormFieldSelect{Name: `Select`}).Stage(stage)

	// Declarations of staged instances of FormFieldString
	__FormFieldString__000000_FirstName := (&models.FormFieldString{Name: `FirstName`}).Stage(stage)
	__FormFieldString__000001_LastName := (&models.FormFieldString{Name: `LastName`}).Stage(stage)

	// Declarations of staged instances of FormFieldTime
	__FormFieldTime__000000_Time := (&models.FormFieldTime{Name: `Time`}).Stage(stage)

	// Declarations of staged instances of FormGroup
	__FormGroup__000000_Form_1 := (&models.FormGroup{Name: `Form 1`}).Stage(stage)

	// Declarations of staged instances of FormSortAssocButton
	__FormSortAssocButton__000000_Sort_Button := (&models.FormSortAssocButton{Name: `Sort Button`}).Stage(stage)

	// Declarations of staged instances of Option
	__Option__000000_Option_1 := (&models.Option{Name: `Option 1`}).Stage(stage)
	__Option__000001_Option_2 := (&models.Option{Name: `Option 2`}).Stage(stage)
	__Option__000002_Option_3 := (&models.Option{Name: `Option 3`}).Stage(stage)
	__Option__000003_Option_4 := (&models.Option{Name: `Option 4`}).Stage(stage)

	// Declarations of staged instances of Row

	// Declarations of staged instances of Table

	// Setup of values

	// CellString values setup
	__CellString__000000_A.Name = `A`
	__CellString__000000_A.Value = `A`

	// CheckBox values setup
	__CheckBox__000000_Boolean.Name = `Boolean`
	__CheckBox__000000_Boolean.Value = true

	// FormDiv values setup
	__FormDiv__000000_CheckBoxs.Name = `CheckBoxs`

	// FormDiv values setup
	__FormDiv__000001_Date_Time.Name = `Date - Time`

	// FormDiv values setup
	__FormDiv__000002_DateTime.Name = `DateTime`

	// FormDiv values setup
	__FormDiv__000003_Edit_Assoc.Name = `Edit Assoc`

	// FormDiv values setup
	__FormDiv__000004_Float64_level_of_x.Name = `Float64 - level of x`

	// FormDiv values setup
	__FormDiv__000005_Form_Div_First_Name_Name.Name = `Form Div First Name  - Name`

	// FormDiv values setup
	__FormDiv__000006_Int_Age.Name = `Int - Age`

	// FormDiv values setup
	__FormDiv__000007_Select.Name = `Select`

	// FormEditAssocButton values setup
	__FormEditAssocButton__000000_Edit_Assoc.Name = `Edit Assoc`
	__FormEditAssocButton__000000_Edit_Assoc.Label = `Association to B`

	// FormField values setup
	__FormField__000000_Age.Name = `Age`
	__FormField__000000_Age.InputTypeEnum = models.Number
	__FormField__000000_Age.Label = `Age`
	__FormField__000000_Age.Placeholder = `18`
	__FormField__000000_Age.HasBespokeWidth = true
	__FormField__000000_Age.BespokeWidthPx = 90

	// FormField values setup
	__FormField__000001_Date.Name = `Date`
	__FormField__000001_Date.InputTypeEnum = models.Date
	__FormField__000001_Date.Label = `Date`
	__FormField__000001_Date.Placeholder = ``
	__FormField__000001_Date.HasBespokeWidth = false
	__FormField__000001_Date.BespokeWidthPx = 0

	// FormField values setup
	__FormField__000002_DateTime.Name = `DateTime`
	__FormField__000002_DateTime.Label = ``
	__FormField__000002_DateTime.Placeholder = ``
	__FormField__000002_DateTime.HasBespokeWidth = false
	__FormField__000002_DateTime.BespokeWidthPx = 0

	// FormField values setup
	__FormField__000003_FirstName.Name = `FirstName`
	__FormField__000003_FirstName.InputTypeEnum = models.Text
	__FormField__000003_FirstName.Label = `First Name (label)`
	__FormField__000003_FirstName.Placeholder = ``
	__FormField__000003_FirstName.HasBespokeWidth = false
	__FormField__000003_FirstName.BespokeWidthPx = 0

	// FormField values setup
	__FormField__000004_Float64_level_of_x.Name = `Float64 - level of x`
	__FormField__000004_Float64_level_of_x.Label = `Float`
	__FormField__000004_Float64_level_of_x.Placeholder = ``
	__FormField__000004_Float64_level_of_x.HasBespokeWidth = false
	__FormField__000004_Float64_level_of_x.BespokeWidthPx = 0

	// FormField values setup
	__FormField__000005_LastName.Name = `LastName`
	__FormField__000005_LastName.InputTypeEnum = models.Text
	__FormField__000005_LastName.Label = `Last Name (label)`
	__FormField__000005_LastName.Placeholder = ``
	__FormField__000005_LastName.HasBespokeWidth = false
	__FormField__000005_LastName.BespokeWidthPx = 0

	// FormField values setup
	__FormField__000006_Select.Name = `Select`
	__FormField__000006_Select.Label = ``
	__FormField__000006_Select.Placeholder = ``
	__FormField__000006_Select.HasBespokeWidth = false
	__FormField__000006_Select.BespokeWidthPx = 0

	// FormField values setup
	__FormField__000007_Time.Name = `Time`
	__FormField__000007_Time.InputTypeEnum = models.Time
	__FormField__000007_Time.Label = `Time`
	__FormField__000007_Time.Placeholder = ``
	__FormField__000007_Time.HasBespokeWidth = false
	__FormField__000007_Time.BespokeWidthPx = 0

	// FormFieldDate values setup
	__FormFieldDate__000000_Time.Name = `Time`
	__FormFieldDate__000000_Time.Value, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-03-02 00:00:00 +0000 UTC")

	// FormFieldDateTime values setup
	__FormFieldDateTime__000000_DateTime.Name = `DateTime`
	__FormFieldDateTime__000000_DateTime.Value, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2001-01-01 01:02:03 +0000 UTC")

	// FormFieldFloat64 values setup
	__FormFieldFloat64__000000_Float64_level_of_x.Name = `Float64 - level of x`
	__FormFieldFloat64__000000_Float64_level_of_x.Value = 64.255500
	__FormFieldFloat64__000000_Float64_level_of_x.HasMinValidator = false
	__FormFieldFloat64__000000_Float64_level_of_x.MinValue = 0.000000
	__FormFieldFloat64__000000_Float64_level_of_x.HasMaxValidator = false
	__FormFieldFloat64__000000_Float64_level_of_x.MaxValue = 0.000000

	// FormFieldInt values setup
	__FormFieldInt__000000_Age.Name = `Age`
	__FormFieldInt__000000_Age.Value = 0
	__FormFieldInt__000000_Age.HasMinValidator = true
	__FormFieldInt__000000_Age.MinValue = 0
	__FormFieldInt__000000_Age.HasMaxValidator = true
	__FormFieldInt__000000_Age.MaxValue = 59

	// FormFieldSelect values setup
	__FormFieldSelect__000000_Select.Name = `Select`

	// FormFieldString values setup
	__FormFieldString__000000_FirstName.Name = `FirstName`
	__FormFieldString__000000_FirstName.Value = `charles`

	// FormFieldString values setup
	__FormFieldString__000001_LastName.Name = `LastName`
	__FormFieldString__000001_LastName.Value = `Bau`

	// FormFieldTime values setup
	__FormFieldTime__000000_Time.Name = `Time`
	__FormFieldTime__000000_Time.Value, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1970-01-01 09:12:10 +0000 UTC")
	__FormFieldTime__000000_Time.Step = 1.000000

	// FormGroup values setup
	__FormGroup__000000_Form_1.Name = `Form 1`

	// FormSortAssocButton values setup
	__FormSortAssocButton__000000_Sort_Button.Name = `Sort Button`
	__FormSortAssocButton__000000_Sort_Button.Label = `Sort Button`

	// Option values setup
	__Option__000000_Option_1.Name = `Option 1`

	// Option values setup
	__Option__000001_Option_2.Name = `Option 2`

	// Option values setup
	__Option__000002_Option_3.Name = `Option 3`

	// Option values setup
	__Option__000003_Option_4.Name = `Option 4`

	// Setup of pointers
	__FormDiv__000000_CheckBoxs.CheckBoxs = append(__FormDiv__000000_CheckBoxs.CheckBoxs, __CheckBox__000000_Boolean)
	__FormDiv__000001_Date_Time.FormFields = append(__FormDiv__000001_Date_Time.FormFields, __FormField__000001_Date)
	__FormDiv__000001_Date_Time.FormFields = append(__FormDiv__000001_Date_Time.FormFields, __FormField__000007_Time)
	__FormDiv__000002_DateTime.FormFields = append(__FormDiv__000002_DateTime.FormFields, __FormField__000002_DateTime)
	__FormDiv__000003_Edit_Assoc.FormEditAssocButton = __FormEditAssocButton__000000_Edit_Assoc
	__FormDiv__000003_Edit_Assoc.FormSortAssocButton = __FormSortAssocButton__000000_Sort_Button
	__FormDiv__000004_Float64_level_of_x.FormFields = append(__FormDiv__000004_Float64_level_of_x.FormFields, __FormField__000004_Float64_level_of_x)
	__FormDiv__000005_Form_Div_First_Name_Name.FormFields = append(__FormDiv__000005_Form_Div_First_Name_Name.FormFields, __FormField__000003_FirstName)
	__FormDiv__000005_Form_Div_First_Name_Name.FormFields = append(__FormDiv__000005_Form_Div_First_Name_Name.FormFields, __FormField__000005_LastName)
	__FormDiv__000006_Int_Age.FormFields = append(__FormDiv__000006_Int_Age.FormFields, __FormField__000000_Age)
	__FormDiv__000007_Select.FormFields = append(__FormDiv__000007_Select.FormFields, __FormField__000006_Select)
	__FormField__000000_Age.FormFieldInt = __FormFieldInt__000000_Age
	__FormField__000001_Date.FormFieldDate = __FormFieldDate__000000_Time
	__FormField__000002_DateTime.FormFieldDateTime = __FormFieldDateTime__000000_DateTime
	__FormField__000003_FirstName.FormFieldString = __FormFieldString__000000_FirstName
	__FormField__000004_Float64_level_of_x.FormFieldFloat64 = __FormFieldFloat64__000000_Float64_level_of_x
	__FormField__000005_LastName.FormFieldString = __FormFieldString__000001_LastName
	__FormField__000006_Select.FormFieldSelect = __FormFieldSelect__000000_Select
	__FormField__000007_Time.FormFieldTime = __FormFieldTime__000000_Time
	__FormFieldSelect__000000_Select.Value = __Option__000000_Option_1
	__FormFieldSelect__000000_Select.Options = append(__FormFieldSelect__000000_Select.Options, __Option__000002_Option_3)
	__FormFieldSelect__000000_Select.Options = append(__FormFieldSelect__000000_Select.Options, __Option__000003_Option_4)
	__FormFieldSelect__000000_Select.Options = append(__FormFieldSelect__000000_Select.Options, __Option__000000_Option_1)
	__FormFieldSelect__000000_Select.Options = append(__FormFieldSelect__000000_Select.Options, __Option__000001_Option_2)
	__FormGroup__000000_Form_1.FormDivs = append(__FormGroup__000000_Form_1.FormDivs, __FormDiv__000000_CheckBoxs)
	__FormGroup__000000_Form_1.FormDivs = append(__FormGroup__000000_Form_1.FormDivs, __FormDiv__000003_Edit_Assoc)
	__FormGroup__000000_Form_1.FormDivs = append(__FormGroup__000000_Form_1.FormDivs, __FormDiv__000006_Int_Age)
	__FormGroup__000000_Form_1.FormDivs = append(__FormGroup__000000_Form_1.FormDivs, __FormDiv__000004_Float64_level_of_x)
	__FormGroup__000000_Form_1.FormDivs = append(__FormGroup__000000_Form_1.FormDivs, __FormDiv__000007_Select)
	__FormGroup__000000_Form_1.FormDivs = append(__FormGroup__000000_Form_1.FormDivs, __FormDiv__000005_Form_Div_First_Name_Name)
	__FormGroup__000000_Form_1.FormDivs = append(__FormGroup__000000_Form_1.FormDivs, __FormDiv__000001_Date_Time)
	__FormGroup__000000_Form_1.FormDivs = append(__FormGroup__000000_Form_1.FormDivs, __FormDiv__000002_DateTime)
}


