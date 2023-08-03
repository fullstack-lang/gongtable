package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Cell:
		if stage.OnAfterCellCreateCallback != nil {
			stage.OnAfterCellCreateCallback.OnAfterCreate(stage, target)
		}
	case *CellBoolean:
		if stage.OnAfterCellBooleanCreateCallback != nil {
			stage.OnAfterCellBooleanCreateCallback.OnAfterCreate(stage, target)
		}
	case *CellFloat64:
		if stage.OnAfterCellFloat64CreateCallback != nil {
			stage.OnAfterCellFloat64CreateCallback.OnAfterCreate(stage, target)
		}
	case *CellIcon:
		if stage.OnAfterCellIconCreateCallback != nil {
			stage.OnAfterCellIconCreateCallback.OnAfterCreate(stage, target)
		}
	case *CellInt:
		if stage.OnAfterCellIntCreateCallback != nil {
			stage.OnAfterCellIntCreateCallback.OnAfterCreate(stage, target)
		}
	case *CellString:
		if stage.OnAfterCellStringCreateCallback != nil {
			stage.OnAfterCellStringCreateCallback.OnAfterCreate(stage, target)
		}
	case *DisplayedColumn:
		if stage.OnAfterDisplayedColumnCreateCallback != nil {
			stage.OnAfterDisplayedColumnCreateCallback.OnAfterCreate(stage, target)
		}
	case *Form:
		if stage.OnAfterFormCreateCallback != nil {
			stage.OnAfterFormCreateCallback.OnAfterCreate(stage, target)
		}
	case *FormCell:
		if stage.OnAfterFormCellCreateCallback != nil {
			stage.OnAfterFormCellCreateCallback.OnAfterCreate(stage, target)
		}
	case *FormCellBoolean:
		if stage.OnAfterFormCellBooleanCreateCallback != nil {
			stage.OnAfterFormCellBooleanCreateCallback.OnAfterCreate(stage, target)
		}
	case *FormCellFloat64:
		if stage.OnAfterFormCellFloat64CreateCallback != nil {
			stage.OnAfterFormCellFloat64CreateCallback.OnAfterCreate(stage, target)
		}
	case *FormCellInt:
		if stage.OnAfterFormCellIntCreateCallback != nil {
			stage.OnAfterFormCellIntCreateCallback.OnAfterCreate(stage, target)
		}
	case *FormCellString:
		if stage.OnAfterFormCellStringCreateCallback != nil {
			stage.OnAfterFormCellStringCreateCallback.OnAfterCreate(stage, target)
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
	case *CellBoolean:
		newTarget := any(new).(*CellBoolean)
		if stage.OnAfterCellBooleanUpdateCallback != nil {
			stage.OnAfterCellBooleanUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *CellFloat64:
		newTarget := any(new).(*CellFloat64)
		if stage.OnAfterCellFloat64UpdateCallback != nil {
			stage.OnAfterCellFloat64UpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *CellIcon:
		newTarget := any(new).(*CellIcon)
		if stage.OnAfterCellIconUpdateCallback != nil {
			stage.OnAfterCellIconUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *CellInt:
		newTarget := any(new).(*CellInt)
		if stage.OnAfterCellIntUpdateCallback != nil {
			stage.OnAfterCellIntUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *Form:
		newTarget := any(new).(*Form)
		if stage.OnAfterFormUpdateCallback != nil {
			stage.OnAfterFormUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormCell:
		newTarget := any(new).(*FormCell)
		if stage.OnAfterFormCellUpdateCallback != nil {
			stage.OnAfterFormCellUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormCellBoolean:
		newTarget := any(new).(*FormCellBoolean)
		if stage.OnAfterFormCellBooleanUpdateCallback != nil {
			stage.OnAfterFormCellBooleanUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormCellFloat64:
		newTarget := any(new).(*FormCellFloat64)
		if stage.OnAfterFormCellFloat64UpdateCallback != nil {
			stage.OnAfterFormCellFloat64UpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormCellInt:
		newTarget := any(new).(*FormCellInt)
		if stage.OnAfterFormCellIntUpdateCallback != nil {
			stage.OnAfterFormCellIntUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormCellString:
		newTarget := any(new).(*FormCellString)
		if stage.OnAfterFormCellStringUpdateCallback != nil {
			stage.OnAfterFormCellStringUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *CellBoolean:
		if stage.OnAfterCellBooleanDeleteCallback != nil {
			staged := any(staged).(*CellBoolean)
			stage.OnAfterCellBooleanDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *CellFloat64:
		if stage.OnAfterCellFloat64DeleteCallback != nil {
			staged := any(staged).(*CellFloat64)
			stage.OnAfterCellFloat64DeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *CellIcon:
		if stage.OnAfterCellIconDeleteCallback != nil {
			staged := any(staged).(*CellIcon)
			stage.OnAfterCellIconDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *CellInt:
		if stage.OnAfterCellIntDeleteCallback != nil {
			staged := any(staged).(*CellInt)
			stage.OnAfterCellIntDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *Form:
		if stage.OnAfterFormDeleteCallback != nil {
			staged := any(staged).(*Form)
			stage.OnAfterFormDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormCell:
		if stage.OnAfterFormCellDeleteCallback != nil {
			staged := any(staged).(*FormCell)
			stage.OnAfterFormCellDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormCellBoolean:
		if stage.OnAfterFormCellBooleanDeleteCallback != nil {
			staged := any(staged).(*FormCellBoolean)
			stage.OnAfterFormCellBooleanDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormCellFloat64:
		if stage.OnAfterFormCellFloat64DeleteCallback != nil {
			staged := any(staged).(*FormCellFloat64)
			stage.OnAfterFormCellFloat64DeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormCellInt:
		if stage.OnAfterFormCellIntDeleteCallback != nil {
			staged := any(staged).(*FormCellInt)
			stage.OnAfterFormCellIntDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormCellString:
		if stage.OnAfterFormCellStringDeleteCallback != nil {
			staged := any(staged).(*FormCellString)
			stage.OnAfterFormCellStringDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *CellBoolean:
		if stage.OnAfterCellBooleanReadCallback != nil {
			stage.OnAfterCellBooleanReadCallback.OnAfterRead(stage, target)
		}
	case *CellFloat64:
		if stage.OnAfterCellFloat64ReadCallback != nil {
			stage.OnAfterCellFloat64ReadCallback.OnAfterRead(stage, target)
		}
	case *CellIcon:
		if stage.OnAfterCellIconReadCallback != nil {
			stage.OnAfterCellIconReadCallback.OnAfterRead(stage, target)
		}
	case *CellInt:
		if stage.OnAfterCellIntReadCallback != nil {
			stage.OnAfterCellIntReadCallback.OnAfterRead(stage, target)
		}
	case *CellString:
		if stage.OnAfterCellStringReadCallback != nil {
			stage.OnAfterCellStringReadCallback.OnAfterRead(stage, target)
		}
	case *DisplayedColumn:
		if stage.OnAfterDisplayedColumnReadCallback != nil {
			stage.OnAfterDisplayedColumnReadCallback.OnAfterRead(stage, target)
		}
	case *Form:
		if stage.OnAfterFormReadCallback != nil {
			stage.OnAfterFormReadCallback.OnAfterRead(stage, target)
		}
	case *FormCell:
		if stage.OnAfterFormCellReadCallback != nil {
			stage.OnAfterFormCellReadCallback.OnAfterRead(stage, target)
		}
	case *FormCellBoolean:
		if stage.OnAfterFormCellBooleanReadCallback != nil {
			stage.OnAfterFormCellBooleanReadCallback.OnAfterRead(stage, target)
		}
	case *FormCellFloat64:
		if stage.OnAfterFormCellFloat64ReadCallback != nil {
			stage.OnAfterFormCellFloat64ReadCallback.OnAfterRead(stage, target)
		}
	case *FormCellInt:
		if stage.OnAfterFormCellIntReadCallback != nil {
			stage.OnAfterFormCellIntReadCallback.OnAfterRead(stage, target)
		}
	case *FormCellString:
		if stage.OnAfterFormCellStringReadCallback != nil {
			stage.OnAfterFormCellStringReadCallback.OnAfterRead(stage, target)
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
	
	case *CellBoolean:
		stage.OnAfterCellBooleanUpdateCallback = any(callback).(OnAfterUpdateInterface[CellBoolean])
	
	case *CellFloat64:
		stage.OnAfterCellFloat64UpdateCallback = any(callback).(OnAfterUpdateInterface[CellFloat64])
	
	case *CellIcon:
		stage.OnAfterCellIconUpdateCallback = any(callback).(OnAfterUpdateInterface[CellIcon])
	
	case *CellInt:
		stage.OnAfterCellIntUpdateCallback = any(callback).(OnAfterUpdateInterface[CellInt])
	
	case *CellString:
		stage.OnAfterCellStringUpdateCallback = any(callback).(OnAfterUpdateInterface[CellString])
	
	case *DisplayedColumn:
		stage.OnAfterDisplayedColumnUpdateCallback = any(callback).(OnAfterUpdateInterface[DisplayedColumn])
	
	case *Form:
		stage.OnAfterFormUpdateCallback = any(callback).(OnAfterUpdateInterface[Form])
	
	case *FormCell:
		stage.OnAfterFormCellUpdateCallback = any(callback).(OnAfterUpdateInterface[FormCell])
	
	case *FormCellBoolean:
		stage.OnAfterFormCellBooleanUpdateCallback = any(callback).(OnAfterUpdateInterface[FormCellBoolean])
	
	case *FormCellFloat64:
		stage.OnAfterFormCellFloat64UpdateCallback = any(callback).(OnAfterUpdateInterface[FormCellFloat64])
	
	case *FormCellInt:
		stage.OnAfterFormCellIntUpdateCallback = any(callback).(OnAfterUpdateInterface[FormCellInt])
	
	case *FormCellString:
		stage.OnAfterFormCellStringUpdateCallback = any(callback).(OnAfterUpdateInterface[FormCellString])
	
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
	
	case *CellBoolean:
		stage.OnAfterCellBooleanCreateCallback = any(callback).(OnAfterCreateInterface[CellBoolean])
	
	case *CellFloat64:
		stage.OnAfterCellFloat64CreateCallback = any(callback).(OnAfterCreateInterface[CellFloat64])
	
	case *CellIcon:
		stage.OnAfterCellIconCreateCallback = any(callback).(OnAfterCreateInterface[CellIcon])
	
	case *CellInt:
		stage.OnAfterCellIntCreateCallback = any(callback).(OnAfterCreateInterface[CellInt])
	
	case *CellString:
		stage.OnAfterCellStringCreateCallback = any(callback).(OnAfterCreateInterface[CellString])
	
	case *DisplayedColumn:
		stage.OnAfterDisplayedColumnCreateCallback = any(callback).(OnAfterCreateInterface[DisplayedColumn])
	
	case *Form:
		stage.OnAfterFormCreateCallback = any(callback).(OnAfterCreateInterface[Form])
	
	case *FormCell:
		stage.OnAfterFormCellCreateCallback = any(callback).(OnAfterCreateInterface[FormCell])
	
	case *FormCellBoolean:
		stage.OnAfterFormCellBooleanCreateCallback = any(callback).(OnAfterCreateInterface[FormCellBoolean])
	
	case *FormCellFloat64:
		stage.OnAfterFormCellFloat64CreateCallback = any(callback).(OnAfterCreateInterface[FormCellFloat64])
	
	case *FormCellInt:
		stage.OnAfterFormCellIntCreateCallback = any(callback).(OnAfterCreateInterface[FormCellInt])
	
	case *FormCellString:
		stage.OnAfterFormCellStringCreateCallback = any(callback).(OnAfterCreateInterface[FormCellString])
	
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
	
	case *CellBoolean:
		stage.OnAfterCellBooleanDeleteCallback = any(callback).(OnAfterDeleteInterface[CellBoolean])
	
	case *CellFloat64:
		stage.OnAfterCellFloat64DeleteCallback = any(callback).(OnAfterDeleteInterface[CellFloat64])
	
	case *CellIcon:
		stage.OnAfterCellIconDeleteCallback = any(callback).(OnAfterDeleteInterface[CellIcon])
	
	case *CellInt:
		stage.OnAfterCellIntDeleteCallback = any(callback).(OnAfterDeleteInterface[CellInt])
	
	case *CellString:
		stage.OnAfterCellStringDeleteCallback = any(callback).(OnAfterDeleteInterface[CellString])
	
	case *DisplayedColumn:
		stage.OnAfterDisplayedColumnDeleteCallback = any(callback).(OnAfterDeleteInterface[DisplayedColumn])
	
	case *Form:
		stage.OnAfterFormDeleteCallback = any(callback).(OnAfterDeleteInterface[Form])
	
	case *FormCell:
		stage.OnAfterFormCellDeleteCallback = any(callback).(OnAfterDeleteInterface[FormCell])
	
	case *FormCellBoolean:
		stage.OnAfterFormCellBooleanDeleteCallback = any(callback).(OnAfterDeleteInterface[FormCellBoolean])
	
	case *FormCellFloat64:
		stage.OnAfterFormCellFloat64DeleteCallback = any(callback).(OnAfterDeleteInterface[FormCellFloat64])
	
	case *FormCellInt:
		stage.OnAfterFormCellIntDeleteCallback = any(callback).(OnAfterDeleteInterface[FormCellInt])
	
	case *FormCellString:
		stage.OnAfterFormCellStringDeleteCallback = any(callback).(OnAfterDeleteInterface[FormCellString])
	
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
	
	case *CellBoolean:
		stage.OnAfterCellBooleanReadCallback = any(callback).(OnAfterReadInterface[CellBoolean])
	
	case *CellFloat64:
		stage.OnAfterCellFloat64ReadCallback = any(callback).(OnAfterReadInterface[CellFloat64])
	
	case *CellIcon:
		stage.OnAfterCellIconReadCallback = any(callback).(OnAfterReadInterface[CellIcon])
	
	case *CellInt:
		stage.OnAfterCellIntReadCallback = any(callback).(OnAfterReadInterface[CellInt])
	
	case *CellString:
		stage.OnAfterCellStringReadCallback = any(callback).(OnAfterReadInterface[CellString])
	
	case *DisplayedColumn:
		stage.OnAfterDisplayedColumnReadCallback = any(callback).(OnAfterReadInterface[DisplayedColumn])
	
	case *Form:
		stage.OnAfterFormReadCallback = any(callback).(OnAfterReadInterface[Form])
	
	case *FormCell:
		stage.OnAfterFormCellReadCallback = any(callback).(OnAfterReadInterface[FormCell])
	
	case *FormCellBoolean:
		stage.OnAfterFormCellBooleanReadCallback = any(callback).(OnAfterReadInterface[FormCellBoolean])
	
	case *FormCellFloat64:
		stage.OnAfterFormCellFloat64ReadCallback = any(callback).(OnAfterReadInterface[FormCellFloat64])
	
	case *FormCellInt:
		stage.OnAfterFormCellIntReadCallback = any(callback).(OnAfterReadInterface[FormCellInt])
	
	case *FormCellString:
		stage.OnAfterFormCellStringReadCallback = any(callback).(OnAfterReadInterface[FormCellString])
	
	case *Row:
		stage.OnAfterRowReadCallback = any(callback).(OnAfterReadInterface[Row])
	
	case *Table:
		stage.OnAfterTableReadCallback = any(callback).(OnAfterReadInterface[Table])
	
	}
}
