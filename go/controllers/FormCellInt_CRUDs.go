// generated by stacks/gong/go/models/controller_file.go
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongtable/go/models"
	"github.com/fullstack-lang/gongtable/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __FormCellInt__dummysDeclaration__ models.FormCellInt
var __FormCellInt_time__dummyDeclaration time.Duration

var mutexFormCellInt sync.Mutex

// An FormCellIntID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFormCellInt updateFormCellInt deleteFormCellInt
type FormCellIntID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormCellIntInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFormCellInt updateFormCellInt
type FormCellIntInput struct {
	// The FormCellInt to submit or modify
	// in: body
	FormCellInt *orm.FormCellIntAPI
}

// GetFormCellInts
//
// swagger:route GET /formcellints formcellints getFormCellInts
//
// # Get all formcellints
//
// Responses:
// default: genericError
//
//	200: formcellintDBResponse
func (controller *Controller) GetFormCellInts(c *gin.Context) {

	// source slice
	var formcellintDBs []orm.FormCellIntDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormCellInts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoFormCellInt.GetDB()

	query := db.Find(&formcellintDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	formcellintAPIs := make([]orm.FormCellIntAPI, 0)

	// for each formcellint, update fields from the database nullable fields
	for idx := range formcellintDBs {
		formcellintDB := &formcellintDBs[idx]
		_ = formcellintDB
		var formcellintAPI orm.FormCellIntAPI

		// insertion point for updating fields
		formcellintAPI.ID = formcellintDB.ID
		formcellintDB.CopyBasicFieldsToFormCellInt(&formcellintAPI.FormCellInt)
		formcellintAPI.FormCellIntPointersEnconding = formcellintDB.FormCellIntPointersEnconding
		formcellintAPIs = append(formcellintAPIs, formcellintAPI)
	}

	c.JSON(http.StatusOK, formcellintAPIs)
}

// PostFormCellInt
//
// swagger:route POST /formcellints formcellints postFormCellInt
//
// Creates a formcellint
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFormCellInt(c *gin.Context) {

	mutexFormCellInt.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFormCellInts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoFormCellInt.GetDB()

	// Validate input
	var input orm.FormCellIntAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create formcellint
	formcellintDB := orm.FormCellIntDB{}
	formcellintDB.FormCellIntPointersEnconding = input.FormCellIntPointersEnconding
	formcellintDB.CopyBasicFieldsFromFormCellInt(&input.FormCellInt)

	query := db.Create(&formcellintDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFormCellInt.CheckoutPhaseOneInstance(&formcellintDB)
	formcellint := backRepo.BackRepoFormCellInt.Map_FormCellIntDBID_FormCellIntPtr[formcellintDB.ID]

	if formcellint != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), formcellint)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, formcellintDB)

	mutexFormCellInt.Unlock()
}

// GetFormCellInt
//
// swagger:route GET /formcellints/{ID} formcellints getFormCellInt
//
// Gets the details for a formcellint.
//
// Responses:
// default: genericError
//
//	200: formcellintDBResponse
func (controller *Controller) GetFormCellInt(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormCellInt", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoFormCellInt.GetDB()

	// Get formcellintDB in DB
	var formcellintDB orm.FormCellIntDB
	if err := db.First(&formcellintDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var formcellintAPI orm.FormCellIntAPI
	formcellintAPI.ID = formcellintDB.ID
	formcellintAPI.FormCellIntPointersEnconding = formcellintDB.FormCellIntPointersEnconding
	formcellintDB.CopyBasicFieldsToFormCellInt(&formcellintAPI.FormCellInt)

	c.JSON(http.StatusOK, formcellintAPI)
}

// UpdateFormCellInt
//
// swagger:route PATCH /formcellints/{ID} formcellints updateFormCellInt
//
// # Update a formcellint
//
// Responses:
// default: genericError
//
//	200: formcellintDBResponse
func (controller *Controller) UpdateFormCellInt(c *gin.Context) {

	mutexFormCellInt.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFormCellInt", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoFormCellInt.GetDB()

	// Validate input
	var input orm.FormCellIntAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formcellintDB orm.FormCellIntDB

	// fetch the formcellint
	query := db.First(&formcellintDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formcellintDB.CopyBasicFieldsFromFormCellInt(&input.FormCellInt)
	formcellintDB.FormCellIntPointersEnconding = input.FormCellIntPointersEnconding

	query = db.Model(&formcellintDB).Updates(formcellintDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formcellintNew := new(models.FormCellInt)
	formcellintDB.CopyBasicFieldsToFormCellInt(formcellintNew)

	// get stage instance from DB instance, and call callback function
	formcellintOld := backRepo.BackRepoFormCellInt.Map_FormCellIntDBID_FormCellIntPtr[formcellintDB.ID]
	if formcellintOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), formcellintOld, formcellintNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formcellintDB
	c.JSON(http.StatusOK, formcellintDB)

	mutexFormCellInt.Unlock()
}

// DeleteFormCellInt
//
// swagger:route DELETE /formcellints/{ID} formcellints deleteFormCellInt
//
// # Delete a formcellint
//
// default: genericError
//
//	200: formcellintDBResponse
func (controller *Controller) DeleteFormCellInt(c *gin.Context) {

	mutexFormCellInt.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFormCellInt", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoFormCellInt.GetDB()

	// Get model if exist
	var formcellintDB orm.FormCellIntDB
	if err := db.First(&formcellintDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&formcellintDB)

	// get an instance (not staged) from DB instance, and call callback function
	formcellintDeleted := new(models.FormCellInt)
	formcellintDB.CopyBasicFieldsToFormCellInt(formcellintDeleted)

	// get stage instance from DB instance, and call callback function
	formcellintStaged := backRepo.BackRepoFormCellInt.Map_FormCellIntDBID_FormCellIntPtr[formcellintDB.ID]
	if formcellintStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), formcellintStaged, formcellintDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexFormCellInt.Unlock()
}