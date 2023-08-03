// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { FormCellDB } from '../formcell-db'
import { FormCellService } from '../formcell.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { FormDB } from '../form-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// FormCellDetailComponent is initilizaed from different routes
// FormCellDetailComponentState detail different cases 
enum FormCellDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Form_FormCells_SET,
}

@Component({
	selector: 'app-formcell-detail',
	templateUrl: './formcell-detail.component.html',
	styleUrls: ['./formcell-detail.component.css'],
})
export class FormCellDetailComponent implements OnInit {

	// insertion point for declarations

	// the FormCellDB of interest
	formcell: FormCellDB = new FormCellDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: FormCellDetailComponentState = FormCellDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private formcellService: FormCellService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private activatedRoute: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {
		this.GONG__StackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')!;

		this.activatedRoute.params.subscribe(params => {
			this.onChangedActivatedRoute()
		});
	}
	onChangedActivatedRoute(): void {

		// compute state
		this.id = +this.activatedRoute.snapshot.paramMap.get('id')!;
		this.originStruct = this.activatedRoute.snapshot.paramMap.get('originStruct')!;
		this.originStructFieldName = this.activatedRoute.snapshot.paramMap.get('originStructFieldName')!;

		this.GONG__StackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')!;

		const association = this.activatedRoute.snapshot.paramMap.get('association');
		if (this.id == 0) {
			this.state = FormCellDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = FormCellDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "FormCells":
						// console.log("FormCell" + " is instanciated with back pointer to instance " + this.id + " Form association FormCells")
						this.state = FormCellDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Form_FormCells_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getFormCell()

		// observable for changes in structs
		this.formcellService.FormCellServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getFormCell()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getFormCell(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case FormCellDetailComponentState.CREATE_INSTANCE:
						this.formcell = new (FormCellDB)
						break;
					case FormCellDetailComponentState.UPDATE_INSTANCE:
						let formcell = frontRepo.FormCells.get(this.id)
						console.assert(formcell != undefined, "missing formcell with id:" + this.id)
						this.formcell = formcell!
						break;
					// insertion point for init of association field
					case FormCellDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Form_FormCells_SET:
						this.formcell = new (FormCellDB)
						this.formcell.Form_FormCells_reverse = frontRepo.Forms.get(this.id)!
						break;
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field
		if (this.formcell.FormCellStringID == undefined) {
			this.formcell.FormCellStringID = new NullInt64
		}
		if (this.formcell.FormCellString != undefined) {
			this.formcell.FormCellStringID.Int64 = this.formcell.FormCellString.ID
			this.formcell.FormCellStringID.Valid = true
		} else {
			this.formcell.FormCellStringID.Int64 = 0
			this.formcell.FormCellStringID.Valid = true
		}
		if (this.formcell.FormCellFloat64ID == undefined) {
			this.formcell.FormCellFloat64ID = new NullInt64
		}
		if (this.formcell.FormCellFloat64 != undefined) {
			this.formcell.FormCellFloat64ID.Int64 = this.formcell.FormCellFloat64.ID
			this.formcell.FormCellFloat64ID.Valid = true
		} else {
			this.formcell.FormCellFloat64ID.Int64 = 0
			this.formcell.FormCellFloat64ID.Valid = true
		}
		if (this.formcell.FormCellIntID == undefined) {
			this.formcell.FormCellIntID = new NullInt64
		}
		if (this.formcell.FormCellInt != undefined) {
			this.formcell.FormCellIntID.Int64 = this.formcell.FormCellInt.ID
			this.formcell.FormCellIntID.Valid = true
		} else {
			this.formcell.FormCellIntID.Int64 = 0
			this.formcell.FormCellIntID.Valid = true
		}
		if (this.formcell.FormCellBoolID == undefined) {
			this.formcell.FormCellBoolID = new NullInt64
		}
		if (this.formcell.FormCellBool != undefined) {
			this.formcell.FormCellBoolID.Int64 = this.formcell.FormCellBool.ID
			this.formcell.FormCellBoolID.Valid = true
		} else {
			this.formcell.FormCellBoolID.Int64 = 0
			this.formcell.FormCellBoolID.Valid = true
		}

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.formcell.Form_FormCells_reverse != undefined) {
			if (this.formcell.Form_FormCellsDBID == undefined) {
				this.formcell.Form_FormCellsDBID = new NullInt64
			}
			this.formcell.Form_FormCellsDBID.Int64 = this.formcell.Form_FormCells_reverse.ID
			this.formcell.Form_FormCellsDBID.Valid = true
			if (this.formcell.Form_FormCellsDBID_Index == undefined) {
				this.formcell.Form_FormCellsDBID_Index = new NullInt64
			}
			this.formcell.Form_FormCellsDBID_Index.Valid = true
			this.formcell.Form_FormCells_reverse = new FormDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case FormCellDetailComponentState.UPDATE_INSTANCE:
				this.formcellService.updateFormCell(this.formcell, this.GONG__StackPath)
					.subscribe(formcell => {
						this.formcellService.FormCellServiceChanged.next("update")
					});
				break;
			default:
				this.formcellService.postFormCell(this.formcell, this.GONG__StackPath).subscribe(formcell => {
					this.formcellService.FormCellServiceChanged.next("post")
					this.formcell = new (FormCellDB) // reset fields
				});
		}
	}

	// openReverseSelection is a generic function that calls dialog for the edition of 
	// ONE-MANY association
	// It uses the MapOfComponent provided by the front repo
	openReverseSelection(AssociatedStruct: string, reverseField: string, selectionMode: string,
		sourceField: string, intermediateStructField: string, nextAssociatedStruct: string) {

		console.log("mode " + selectionMode)

		const dialogConfig = new MatDialogConfig();

		let dialogData = new DialogData();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.width = "50%"
		dialogConfig.height = "50%"
		if (selectionMode == SelectionMode.ONE_MANY_ASSOCIATION_MODE) {

			dialogData.ID = this.formcell.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			dialogConfig.data = dialogData
			const dialogRef: MatDialogRef<string, any> = this.dialog.open(
				MapOfComponents.get(AssociatedStruct).get(
					AssociatedStruct + 'sTableComponent'
				),
				dialogConfig
			);
			dialogRef.afterClosed().subscribe(result => {
			});
		}
		if (selectionMode == SelectionMode.MANY_MANY_ASSOCIATION_MODE) {
			dialogData.ID = this.formcell.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "FormCell"
			dialogData.SourceField = sourceField

			// set up the intermediate struct
			dialogData.IntermediateStruct = AssociatedStruct
			dialogData.IntermediateStructField = intermediateStructField

			// set up the end struct
			dialogData.NextAssociationStruct = nextAssociatedStruct

			dialogConfig.data = dialogData
			const dialogRef: MatDialogRef<string, any> = this.dialog.open(
				MapOfComponents.get(nextAssociatedStruct).get(
					nextAssociatedStruct + 'sTableComponent'
				),
				dialogConfig
			);
			dialogRef.afterClosed().subscribe(result => {
			});
		}

	}

	openDragAndDropOrdering(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.data = {
			ID: this.formcell.ID,
			ReversePointer: reverseField,
			OrderingMode: true,
			GONG__StackPath: this.GONG__StackPath,
		};
		const dialogRef: MatDialogRef<string, any> = this.dialog.open(
			MapOfSortingComponents.get(AssociatedStruct).get(
				AssociatedStruct + 'SortingComponent'
			),
			dialogConfig
		);

		dialogRef.afterClosed().subscribe(result => {
		});
	}

	fillUpNameIfEmpty(event: { value: { Name: string; }; }) {
		if (this.formcell.Name == "") {
			this.formcell.Name = event.value.Name
		}
	}

	toggleTextArea(fieldName: string) {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			let displayAsTextArea = this.mapFields_displayAsTextArea.get(fieldName)
			this.mapFields_displayAsTextArea.set(fieldName, !displayAsTextArea)
		} else {
			this.mapFields_displayAsTextArea.set(fieldName, true)
		}
	}

	isATextArea(fieldName: string): boolean {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			return this.mapFields_displayAsTextArea.get(fieldName)!
		} else {
			return false
		}
	}

	compareObjects(o1: any, o2: any) {
		if (o1?.ID == o2?.ID) {
			return true;
		}
		else {
			return false
		}
	}
}
