package orm

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/fullstack-lang/gongtable/go/models"

	"github.com/tealeg/xlsx/v3"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoCell BackRepoCellStruct

	BackRepoCellFloat64 BackRepoCellFloat64Struct

	BackRepoCellIcon BackRepoCellIconStruct

	BackRepoCellInt BackRepoCellIntStruct

	BackRepoCellString BackRepoCellStringStruct

	BackRepoDisplayedColumn BackRepoDisplayedColumnStruct

	BackRepoRow BackRepoRowStruct

	BackRepoTable BackRepoTableStruct

	CommitFromBackNb uint // records commit increments when performed by the back

	PushFromFrontNb uint // records commit increments when performed by the front

	stage *models.StageStruct
}

func NewBackRepo(stage *models.StageStruct, filename string) (backRepo *BackRepoStruct) {

	// adjust naming strategy to the stack
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "github_com_fullstack_lang_gong_test_go_", // table name prefix
		},
	}
	db, err := gorm.Open(sqlite.Open(filename), gormConfig)

	// since testsim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	dbDB_inMemory, err := db.DB()
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	// it is mandatory to allow parallel access, otherwise, bizarre errors occurs
	dbDB_inMemory.SetMaxOpenConns(1)

	if err != nil {
		panic("Failed to connect to database!")
	}

	// adjust naming strategy to the stack
	db.Config.NamingStrategy = &schema.NamingStrategy{
		TablePrefix: "github_com_fullstack_lang_gong_test_go_", // table name prefix
	}

	err = db.AutoMigrate( // insertion point for reference to structs
		&CellDB{},
		&CellFloat64DB{},
		&CellIconDB{},
		&CellIntDB{},
		&CellStringDB{},
		&DisplayedColumnDB{},
		&RowDB{},
		&TableDB{},
	)

	if err != nil {
		msg := err.Error()
		panic("problem with migration " + msg + " on package github.com/fullstack-lang/gong/test/go")
	}

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoCell = BackRepoCellStruct{
		Map_CellDBID_CellPtr: make(map[uint]*models.Cell, 0),
		Map_CellDBID_CellDB:  make(map[uint]*CellDB, 0),
		Map_CellPtr_CellDBID: make(map[*models.Cell]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoCellFloat64 = BackRepoCellFloat64Struct{
		Map_CellFloat64DBID_CellFloat64Ptr: make(map[uint]*models.CellFloat64, 0),
		Map_CellFloat64DBID_CellFloat64DB:  make(map[uint]*CellFloat64DB, 0),
		Map_CellFloat64Ptr_CellFloat64DBID: make(map[*models.CellFloat64]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoCellIcon = BackRepoCellIconStruct{
		Map_CellIconDBID_CellIconPtr: make(map[uint]*models.CellIcon, 0),
		Map_CellIconDBID_CellIconDB:  make(map[uint]*CellIconDB, 0),
		Map_CellIconPtr_CellIconDBID: make(map[*models.CellIcon]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoCellInt = BackRepoCellIntStruct{
		Map_CellIntDBID_CellIntPtr: make(map[uint]*models.CellInt, 0),
		Map_CellIntDBID_CellIntDB:  make(map[uint]*CellIntDB, 0),
		Map_CellIntPtr_CellIntDBID: make(map[*models.CellInt]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoCellString = BackRepoCellStringStruct{
		Map_CellStringDBID_CellStringPtr: make(map[uint]*models.CellString, 0),
		Map_CellStringDBID_CellStringDB:  make(map[uint]*CellStringDB, 0),
		Map_CellStringPtr_CellStringDBID: make(map[*models.CellString]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDisplayedColumn = BackRepoDisplayedColumnStruct{
		Map_DisplayedColumnDBID_DisplayedColumnPtr: make(map[uint]*models.DisplayedColumn, 0),
		Map_DisplayedColumnDBID_DisplayedColumnDB:  make(map[uint]*DisplayedColumnDB, 0),
		Map_DisplayedColumnPtr_DisplayedColumnDBID: make(map[*models.DisplayedColumn]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoRow = BackRepoRowStruct{
		Map_RowDBID_RowPtr: make(map[uint]*models.Row, 0),
		Map_RowDBID_RowDB:  make(map[uint]*RowDB, 0),
		Map_RowPtr_RowDBID: make(map[*models.Row]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTable = BackRepoTableStruct{
		Map_TableDBID_TablePtr: make(map[uint]*models.Table, 0),
		Map_TableDBID_TableDB:  make(map[uint]*TableDB, 0),
		Map_TablePtr_TableDBID: make(map[*models.Table]uint, 0),

		db:    db,
		stage: stage,
	}

	stage.BackRepo = backRepo
	backRepo.stage = stage

	return
}

func (backRepo *BackRepoStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepo.stage
	return
}

func (backRepo *BackRepoStruct) GetLastCommitFromBackNb() uint {
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) GetLastPushFromFrontNb() uint {
	return backRepo.PushFromFrontNb
}

func (backRepo *BackRepoStruct) IncrementCommitFromBackNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromBackCallback != nil {
		backRepo.stage.OnInitCommitFromBackCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) IncrementPushFromFrontNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromFrontCallback != nil {
		backRepo.stage.OnInitCommitFromFrontCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.PushFromFrontNb = backRepo.PushFromFrontNb + 1
	return backRepo.CommitFromBackNb
}

// Commit the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) Commit(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoCell.CommitPhaseOne(stage)
	backRepo.BackRepoCellFloat64.CommitPhaseOne(stage)
	backRepo.BackRepoCellIcon.CommitPhaseOne(stage)
	backRepo.BackRepoCellInt.CommitPhaseOne(stage)
	backRepo.BackRepoCellString.CommitPhaseOne(stage)
	backRepo.BackRepoDisplayedColumn.CommitPhaseOne(stage)
	backRepo.BackRepoRow.CommitPhaseOne(stage)
	backRepo.BackRepoTable.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoCell.CommitPhaseTwo(backRepo)
	backRepo.BackRepoCellFloat64.CommitPhaseTwo(backRepo)
	backRepo.BackRepoCellIcon.CommitPhaseTwo(backRepo)
	backRepo.BackRepoCellInt.CommitPhaseTwo(backRepo)
	backRepo.BackRepoCellString.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDisplayedColumn.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRow.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTable.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoCell.CheckoutPhaseOne()
	backRepo.BackRepoCellFloat64.CheckoutPhaseOne()
	backRepo.BackRepoCellIcon.CheckoutPhaseOne()
	backRepo.BackRepoCellInt.CheckoutPhaseOne()
	backRepo.BackRepoCellString.CheckoutPhaseOne()
	backRepo.BackRepoDisplayedColumn.CheckoutPhaseOne()
	backRepo.BackRepoRow.CheckoutPhaseOne()
	backRepo.BackRepoTable.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoCell.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoCellFloat64.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoCellIcon.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoCellInt.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoCellString.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDisplayedColumn.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRow.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTable.CheckoutPhaseTwo(backRepo)
}

var _backRepo *BackRepoStruct

var once sync.Once

func GetDefaultBackRepo() *BackRepoStruct {
	once.Do(func() {
		_backRepo = NewBackRepo(models.GetDefaultStage(), "")
	})
	return _backRepo
}

func GetLastCommitFromBackNb() uint {
	return GetDefaultBackRepo().GetLastCommitFromBackNb()
}

func GetLastPushFromFrontNb() uint {
	return GetDefaultBackRepo().GetLastPushFromFrontNb()
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoCell.Backup(dirPath)
	backRepo.BackRepoCellFloat64.Backup(dirPath)
	backRepo.BackRepoCellIcon.Backup(dirPath)
	backRepo.BackRepoCellInt.Backup(dirPath)
	backRepo.BackRepoCellString.Backup(dirPath)
	backRepo.BackRepoDisplayedColumn.Backup(dirPath)
	backRepo.BackRepoRow.Backup(dirPath)
	backRepo.BackRepoTable.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoCell.BackupXL(file)
	backRepo.BackRepoCellFloat64.BackupXL(file)
	backRepo.BackRepoCellIcon.BackupXL(file)
	backRepo.BackRepoCellInt.BackupXL(file)
	backRepo.BackRepoCellString.BackupXL(file)
	backRepo.BackRepoDisplayedColumn.BackupXL(file)
	backRepo.BackRepoRow.BackupXL(file)
	backRepo.BackRepoTable.BackupXL(file)

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	file.Write(writer)
	theBytes := b.Bytes()

	filename := filepath.Join(dirPath, "bckp.xlsx")
	err := ioutil.WriteFile(filename, theBytes, 0644)
	if err != nil {
		log.Panic("Cannot write the XL file", err.Error())
	}
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) Restore(stage *models.StageStruct, dirPath string) {
	backRepo.stage.Commit()
	backRepo.stage.Reset()
	backRepo.stage.Checkout()

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoCell.RestorePhaseOne(dirPath)
	backRepo.BackRepoCellFloat64.RestorePhaseOne(dirPath)
	backRepo.BackRepoCellIcon.RestorePhaseOne(dirPath)
	backRepo.BackRepoCellInt.RestorePhaseOne(dirPath)
	backRepo.BackRepoCellString.RestorePhaseOne(dirPath)
	backRepo.BackRepoDisplayedColumn.RestorePhaseOne(dirPath)
	backRepo.BackRepoRow.RestorePhaseOne(dirPath)
	backRepo.BackRepoTable.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoCell.RestorePhaseTwo()
	backRepo.BackRepoCellFloat64.RestorePhaseTwo()
	backRepo.BackRepoCellIcon.RestorePhaseTwo()
	backRepo.BackRepoCellInt.RestorePhaseTwo()
	backRepo.BackRepoCellString.RestorePhaseTwo()
	backRepo.BackRepoDisplayedColumn.RestorePhaseTwo()
	backRepo.BackRepoRow.RestorePhaseTwo()
	backRepo.BackRepoTable.RestorePhaseTwo()

	backRepo.stage.Checkout()
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) RestoreXL(stage *models.StageStruct, dirPath string) {

	// clean the stage
	backRepo.stage.Reset()

	// commit the cleaned stage
	backRepo.stage.Commit()

	// open an existing file
	filename := filepath.Join(dirPath, "bckp.xlsx")
	file, err := xlsx.OpenFile(filename)
	_ = file

	if err != nil {
		log.Panic("Cannot read the XL file", err.Error())
	}

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoCell.RestoreXLPhaseOne(file)
	backRepo.BackRepoCellFloat64.RestoreXLPhaseOne(file)
	backRepo.BackRepoCellIcon.RestoreXLPhaseOne(file)
	backRepo.BackRepoCellInt.RestoreXLPhaseOne(file)
	backRepo.BackRepoCellString.RestoreXLPhaseOne(file)
	backRepo.BackRepoDisplayedColumn.RestoreXLPhaseOne(file)
	backRepo.BackRepoRow.RestoreXLPhaseOne(file)
	backRepo.BackRepoTable.RestoreXLPhaseOne(file)

	// commit the restored stage
	backRepo.stage.Commit()
}
