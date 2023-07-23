package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Cell:
		if stage.OnAfterCellCreateCallback != nil {
			stage.OnAfterCellCreateCallback.OnAfterCreate(stage, target)
		}
	case *CellString:
		if stage.OnAfterCellStringCreateCallback != nil {
			stage.OnAfterCellStringCreateCallback.OnAfterCreate(stage, target)
		}
	case *DisplayedColumn:
		if stage.OnAfterDisplayedColumnCreateCallback != nil {
			stage.OnAfterDisplayedColumnCreateCallback.OnAfterCreate(stage, target)
		}
	case *Row:
		if stage.OnAfterRowCreateCallback != nil {
			stage.OnAfterRowCreateCallback.OnAfterCreate(stage, target)
		}
	case *Table:
		if stage.OnAfterTableCreateCallback != nil {
			stage.OnAfterTableCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Cell:
		newTarget := any(new).(*Cell)
		if stage.OnAfterCellUpdateCallback != nil {
			stage.OnAfterCellUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *CellString:
		newTarget := any(new).(*CellString)
		if stage.OnAfterCellStringUpdateCallback != nil {
			stage.OnAfterCellStringUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DisplayedColumn:
		newTarget := any(new).(*DisplayedColumn)
		if stage.OnAfterDisplayedColumnUpdateCallback != nil {
			stage.OnAfterDisplayedColumnUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Row:
		newTarget := any(new).(*Row)
		if stage.OnAfterRowUpdateCallback != nil {
			stage.OnAfterRowUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Table:
		newTarget := any(new).(*Table)
		if stage.OnAfterTableUpdateCallback != nil {
			stage.OnAfterTableUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Cell:
		if stage.OnAfterCellDeleteCallback != nil {
			staged := any(staged).(*Cell)
			stage.OnAfterCellDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *CellString:
		if stage.OnAfterCellStringDeleteCallback != nil {
			staged := any(staged).(*CellString)
			stage.OnAfterCellStringDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DisplayedColumn:
		if stage.OnAfterDisplayedColumnDeleteCallback != nil {
			staged := any(staged).(*DisplayedColumn)
			stage.OnAfterDisplayedColumnDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Row:
		if stage.OnAfterRowDeleteCallback != nil {
			staged := any(staged).(*Row)
			stage.OnAfterRowDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Table:
		if stage.OnAfterTableDeleteCallback != nil {
			staged := any(staged).(*Table)
			stage.OnAfterTableDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Cell:
		if stage.OnAfterCellReadCallback != nil {
			stage.OnAfterCellReadCallback.OnAfterRead(stage, target)
		}
	case *CellString:
		if stage.OnAfterCellStringReadCallback != nil {
			stage.OnAfterCellStringReadCallback.OnAfterRead(stage, target)
		}
	case *DisplayedColumn:
		if stage.OnAfterDisplayedColumnReadCallback != nil {
			stage.OnAfterDisplayedColumnReadCallback.OnAfterRead(stage, target)
		}
	case *Row:
		if stage.OnAfterRowReadCallback != nil {
			stage.OnAfterRowReadCallback.OnAfterRead(stage, target)
		}
	case *Table:
		if stage.OnAfterTableReadCallback != nil {
			stage.OnAfterTableReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Cell:
		stage.OnAfterCellUpdateCallback = any(callback).(OnAfterUpdateInterface[Cell])
	
	case *CellString:
		stage.OnAfterCellStringUpdateCallback = any(callback).(OnAfterUpdateInterface[CellString])
	
	case *DisplayedColumn:
		stage.OnAfterDisplayedColumnUpdateCallback = any(callback).(OnAfterUpdateInterface[DisplayedColumn])
	
	case *Row:
		stage.OnAfterRowUpdateCallback = any(callback).(OnAfterUpdateInterface[Row])
	
	case *Table:
		stage.OnAfterTableUpdateCallback = any(callback).(OnAfterUpdateInterface[Table])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Cell:
		stage.OnAfterCellCreateCallback = any(callback).(OnAfterCreateInterface[Cell])
	
	case *CellString:
		stage.OnAfterCellStringCreateCallback = any(callback).(OnAfterCreateInterface[CellString])
	
	case *DisplayedColumn:
		stage.OnAfterDisplayedColumnCreateCallback = any(callback).(OnAfterCreateInterface[DisplayedColumn])
	
	case *Row:
		stage.OnAfterRowCreateCallback = any(callback).(OnAfterCreateInterface[Row])
	
	case *Table:
		stage.OnAfterTableCreateCallback = any(callback).(OnAfterCreateInterface[Table])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Cell:
		stage.OnAfterCellDeleteCallback = any(callback).(OnAfterDeleteInterface[Cell])
	
	case *CellString:
		stage.OnAfterCellStringDeleteCallback = any(callback).(OnAfterDeleteInterface[CellString])
	
	case *DisplayedColumn:
		stage.OnAfterDisplayedColumnDeleteCallback = any(callback).(OnAfterDeleteInterface[DisplayedColumn])
	
	case *Row:
		stage.OnAfterRowDeleteCallback = any(callback).(OnAfterDeleteInterface[Row])
	
	case *Table:
		stage.OnAfterTableDeleteCallback = any(callback).(OnAfterDeleteInterface[Table])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Cell:
		stage.OnAfterCellReadCallback = any(callback).(OnAfterReadInterface[Cell])
	
	case *CellString:
		stage.OnAfterCellStringReadCallback = any(callback).(OnAfterReadInterface[CellString])
	
	case *DisplayedColumn:
		stage.OnAfterDisplayedColumnReadCallback = any(callback).(OnAfterReadInterface[DisplayedColumn])
	
	case *Row:
		stage.OnAfterRowReadCallback = any(callback).(OnAfterReadInterface[Row])
	
	case *Table:
		stage.OnAfterTableReadCallback = any(callback).(OnAfterReadInterface[Table])
	
	}
}
