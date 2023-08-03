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

	// Declarations of staged instances of Form
	__Form__000000_Form_1 := (&models.Form{Name: `Form 1`}).Stage(stage)

	// Declarations of staged instances of FormField
	__FormField__000000_Form_Cell_Int := (&models.FormField{Name: `Form Cell Int`}).Stage(stage)

	// Declarations of staged instances of FormFieldBoolean

	// Declarations of staged instances of FormFieldFloat64

	// Declarations of staged instances of FormFieldInt
	__FormFieldInt__000000_Form_Cell_Int := (&models.FormFieldInt{Name: `Form Cell Int`}).Stage(stage)

	// Declarations of staged instances of FormFieldString

	// Declarations of staged instances of Row

	// Declarations of staged instances of Table

	// Setup of values

	// Form values setup
	__Form__000000_Form_1.Name = `Form 1`

	// FormField values setup
	__FormField__000000_Form_Cell_Int.Name = `Form Cell Int`

	// FormFieldInt values setup
	__FormFieldInt__000000_Form_Cell_Int.Name = `Form Cell Int`
	__FormFieldInt__000000_Form_Cell_Int.Value = 12

	// Setup of pointers
	__Form__000000_Form_1.FormCells = append(__Form__000000_Form_1.FormCells, __FormField__000000_Form_Cell_Int)
	__FormField__000000_Form_Cell_Int.FormFieldInt = __FormFieldInt__000000_Form_Cell_Int
}


