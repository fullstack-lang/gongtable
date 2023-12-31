// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *GongBasicField:
		ok = stage.IsStagedGongBasicField(target)

	case *GongEnum:
		ok = stage.IsStagedGongEnum(target)

	case *GongEnumValue:
		ok = stage.IsStagedGongEnumValue(target)

	case *GongLink:
		ok = stage.IsStagedGongLink(target)

	case *GongNote:
		ok = stage.IsStagedGongNote(target)

	case *GongStruct:
		ok = stage.IsStagedGongStruct(target)

	case *GongTimeField:
		ok = stage.IsStagedGongTimeField(target)

	case *Meta:
		ok = stage.IsStagedMeta(target)

	case *MetaReference:
		ok = stage.IsStagedMetaReference(target)

	case *ModelPkg:
		ok = stage.IsStagedModelPkg(target)

	case *PointerToGongStructField:
		ok = stage.IsStagedPointerToGongStructField(target)

	case *SliceOfPointerToGongStructField:
		ok = stage.IsStagedSliceOfPointerToGongStructField(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
	func (stage *StageStruct) IsStagedGongBasicField(gongbasicfield *GongBasicField) (ok bool) {

		_, ok = stage.GongBasicFields[gongbasicfield]
	
		return
	}

	func (stage *StageStruct) IsStagedGongEnum(gongenum *GongEnum) (ok bool) {

		_, ok = stage.GongEnums[gongenum]
	
		return
	}

	func (stage *StageStruct) IsStagedGongEnumValue(gongenumvalue *GongEnumValue) (ok bool) {

		_, ok = stage.GongEnumValues[gongenumvalue]
	
		return
	}

	func (stage *StageStruct) IsStagedGongLink(gonglink *GongLink) (ok bool) {

		_, ok = stage.GongLinks[gonglink]
	
		return
	}

	func (stage *StageStruct) IsStagedGongNote(gongnote *GongNote) (ok bool) {

		_, ok = stage.GongNotes[gongnote]
	
		return
	}

	func (stage *StageStruct) IsStagedGongStruct(gongstruct *GongStruct) (ok bool) {

		_, ok = stage.GongStructs[gongstruct]
	
		return
	}

	func (stage *StageStruct) IsStagedGongTimeField(gongtimefield *GongTimeField) (ok bool) {

		_, ok = stage.GongTimeFields[gongtimefield]
	
		return
	}

	func (stage *StageStruct) IsStagedMeta(meta *Meta) (ok bool) {

		_, ok = stage.Metas[meta]
	
		return
	}

	func (stage *StageStruct) IsStagedMetaReference(metareference *MetaReference) (ok bool) {

		_, ok = stage.MetaReferences[metareference]
	
		return
	}

	func (stage *StageStruct) IsStagedModelPkg(modelpkg *ModelPkg) (ok bool) {

		_, ok = stage.ModelPkgs[modelpkg]
	
		return
	}

	func (stage *StageStruct) IsStagedPointerToGongStructField(pointertogongstructfield *PointerToGongStructField) (ok bool) {

		_, ok = stage.PointerToGongStructFields[pointertogongstructfield]
	
		return
	}

	func (stage *StageStruct) IsStagedSliceOfPointerToGongStructField(sliceofpointertogongstructfield *SliceOfPointerToGongStructField) (ok bool) {

		_, ok = stage.SliceOfPointerToGongStructFields[sliceofpointertogongstructfield]
	
		return
	}


// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *GongBasicField:
		stage.StageBranchGongBasicField(target)

	case *GongEnum:
		stage.StageBranchGongEnum(target)

	case *GongEnumValue:
		stage.StageBranchGongEnumValue(target)

	case *GongLink:
		stage.StageBranchGongLink(target)

	case *GongNote:
		stage.StageBranchGongNote(target)

	case *GongStruct:
		stage.StageBranchGongStruct(target)

	case *GongTimeField:
		stage.StageBranchGongTimeField(target)

	case *Meta:
		stage.StageBranchMeta(target)

	case *MetaReference:
		stage.StageBranchMetaReference(target)

	case *ModelPkg:
		stage.StageBranchModelPkg(target)

	case *PointerToGongStructField:
		stage.StageBranchPointerToGongStructField(target)

	case *SliceOfPointerToGongStructField:
		stage.StageBranchSliceOfPointerToGongStructField(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchGongBasicField(gongbasicfield *GongBasicField) {

	// check if instance is already staged
	if IsStaged(stage, gongbasicfield) {
		return
	}

	gongbasicfield.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if gongbasicfield.GongEnum != nil {
		StageBranch(stage, gongbasicfield.GongEnum)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchGongEnum(gongenum *GongEnum) {

	// check if instance is already staged
	if IsStaged(stage, gongenum) {
		return
	}

	gongenum.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongenumvalue := range gongenum.GongEnumValues {
		StageBranch(stage, _gongenumvalue)
	}

}

func (stage *StageStruct) StageBranchGongEnumValue(gongenumvalue *GongEnumValue) {

	// check if instance is already staged
	if IsStaged(stage, gongenumvalue) {
		return
	}

	gongenumvalue.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchGongLink(gonglink *GongLink) {

	// check if instance is already staged
	if IsStaged(stage, gonglink) {
		return
	}

	gonglink.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchGongNote(gongnote *GongNote) {

	// check if instance is already staged
	if IsStaged(stage, gongnote) {
		return
	}

	gongnote.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gonglink := range gongnote.Links {
		StageBranch(stage, _gonglink)
	}

}

func (stage *StageStruct) StageBranchGongStruct(gongstruct *GongStruct) {

	// check if instance is already staged
	if IsStaged(stage, gongstruct) {
		return
	}

	gongstruct.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongbasicfield := range gongstruct.GongBasicFields {
		StageBranch(stage, _gongbasicfield)
	}
	for _, _gongtimefield := range gongstruct.GongTimeFields {
		StageBranch(stage, _gongtimefield)
	}
	for _, _pointertogongstructfield := range gongstruct.PointerToGongStructFields {
		StageBranch(stage, _pointertogongstructfield)
	}
	for _, _sliceofpointertogongstructfield := range gongstruct.SliceOfPointerToGongStructFields {
		StageBranch(stage, _sliceofpointertogongstructfield)
	}

}

func (stage *StageStruct) StageBranchGongTimeField(gongtimefield *GongTimeField) {

	// check if instance is already staged
	if IsStaged(stage, gongtimefield) {
		return
	}

	gongtimefield.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchMeta(meta *Meta) {

	// check if instance is already staged
	if IsStaged(stage, meta) {
		return
	}

	meta.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _metareference := range meta.MetaReferences {
		StageBranch(stage, _metareference)
	}

}

func (stage *StageStruct) StageBranchMetaReference(metareference *MetaReference) {

	// check if instance is already staged
	if IsStaged(stage, metareference) {
		return
	}

	metareference.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchModelPkg(modelpkg *ModelPkg) {

	// check if instance is already staged
	if IsStaged(stage, modelpkg) {
		return
	}

	modelpkg.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchPointerToGongStructField(pointertogongstructfield *PointerToGongStructField) {

	// check if instance is already staged
	if IsStaged(stage, pointertogongstructfield) {
		return
	}

	pointertogongstructfield.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if pointertogongstructfield.GongStruct != nil {
		StageBranch(stage, pointertogongstructfield.GongStruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSliceOfPointerToGongStructField(sliceofpointertogongstructfield *SliceOfPointerToGongStructField) {

	// check if instance is already staged
	if IsStaged(stage, sliceofpointertogongstructfield) {
		return
	}

	sliceofpointertogongstructfield.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if sliceofpointertogongstructfield.GongStruct != nil {
		StageBranch(stage, sliceofpointertogongstructfield.GongStruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}


// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *GongBasicField:
		stage.UnstageBranchGongBasicField(target)

	case *GongEnum:
		stage.UnstageBranchGongEnum(target)

	case *GongEnumValue:
		stage.UnstageBranchGongEnumValue(target)

	case *GongLink:
		stage.UnstageBranchGongLink(target)

	case *GongNote:
		stage.UnstageBranchGongNote(target)

	case *GongStruct:
		stage.UnstageBranchGongStruct(target)

	case *GongTimeField:
		stage.UnstageBranchGongTimeField(target)

	case *Meta:
		stage.UnstageBranchMeta(target)

	case *MetaReference:
		stage.UnstageBranchMetaReference(target)

	case *ModelPkg:
		stage.UnstageBranchModelPkg(target)

	case *PointerToGongStructField:
		stage.UnstageBranchPointerToGongStructField(target)

	case *SliceOfPointerToGongStructField:
		stage.UnstageBranchSliceOfPointerToGongStructField(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchGongBasicField(gongbasicfield *GongBasicField) {

	// check if instance is already staged
	if ! IsStaged(stage, gongbasicfield) {
		return
	}

	gongbasicfield.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if gongbasicfield.GongEnum != nil {
		UnstageBranch(stage, gongbasicfield.GongEnum)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchGongEnum(gongenum *GongEnum) {

	// check if instance is already staged
	if ! IsStaged(stage, gongenum) {
		return
	}

	gongenum.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongenumvalue := range gongenum.GongEnumValues {
		UnstageBranch(stage, _gongenumvalue)
	}

}

func (stage *StageStruct) UnstageBranchGongEnumValue(gongenumvalue *GongEnumValue) {

	// check if instance is already staged
	if ! IsStaged(stage, gongenumvalue) {
		return
	}

	gongenumvalue.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchGongLink(gonglink *GongLink) {

	// check if instance is already staged
	if ! IsStaged(stage, gonglink) {
		return
	}

	gonglink.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchGongNote(gongnote *GongNote) {

	// check if instance is already staged
	if ! IsStaged(stage, gongnote) {
		return
	}

	gongnote.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gonglink := range gongnote.Links {
		UnstageBranch(stage, _gonglink)
	}

}

func (stage *StageStruct) UnstageBranchGongStruct(gongstruct *GongStruct) {

	// check if instance is already staged
	if ! IsStaged(stage, gongstruct) {
		return
	}

	gongstruct.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongbasicfield := range gongstruct.GongBasicFields {
		UnstageBranch(stage, _gongbasicfield)
	}
	for _, _gongtimefield := range gongstruct.GongTimeFields {
		UnstageBranch(stage, _gongtimefield)
	}
	for _, _pointertogongstructfield := range gongstruct.PointerToGongStructFields {
		UnstageBranch(stage, _pointertogongstructfield)
	}
	for _, _sliceofpointertogongstructfield := range gongstruct.SliceOfPointerToGongStructFields {
		UnstageBranch(stage, _sliceofpointertogongstructfield)
	}

}

func (stage *StageStruct) UnstageBranchGongTimeField(gongtimefield *GongTimeField) {

	// check if instance is already staged
	if ! IsStaged(stage, gongtimefield) {
		return
	}

	gongtimefield.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchMeta(meta *Meta) {

	// check if instance is already staged
	if ! IsStaged(stage, meta) {
		return
	}

	meta.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _metareference := range meta.MetaReferences {
		UnstageBranch(stage, _metareference)
	}

}

func (stage *StageStruct) UnstageBranchMetaReference(metareference *MetaReference) {

	// check if instance is already staged
	if ! IsStaged(stage, metareference) {
		return
	}

	metareference.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchModelPkg(modelpkg *ModelPkg) {

	// check if instance is already staged
	if ! IsStaged(stage, modelpkg) {
		return
	}

	modelpkg.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchPointerToGongStructField(pointertogongstructfield *PointerToGongStructField) {

	// check if instance is already staged
	if ! IsStaged(stage, pointertogongstructfield) {
		return
	}

	pointertogongstructfield.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if pointertogongstructfield.GongStruct != nil {
		UnstageBranch(stage, pointertogongstructfield.GongStruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSliceOfPointerToGongStructField(sliceofpointertogongstructfield *SliceOfPointerToGongStructField) {

	// check if instance is already staged
	if ! IsStaged(stage, sliceofpointertogongstructfield) {
		return
	}

	sliceofpointertogongstructfield.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if sliceofpointertogongstructfield.GongStruct != nil {
		UnstageBranch(stage, sliceofpointertogongstructfield.GongStruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

