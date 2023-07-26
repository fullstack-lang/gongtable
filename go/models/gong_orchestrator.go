package models

// insertion point
// CellIconOrchestrator
type CellIconOrchestrator struct {
}

func (orchestrator *CellIconOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedCellIcon, backRepoCellIcon *CellIcon) {

	stagedCellIcon.OnAfterUpdate(gongsvgStage, stagedCellIcon, backRepoCellIcon)
}

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *StageStruct) {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case CellIcon:
		stage.OnAfterCellIconUpdateCallback = new(CellIconOrchestrator)

	}

}
