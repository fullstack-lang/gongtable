import { Component, Input, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';

import * as gongtable from 'gongtable'

import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import * as moment from 'moment';

@Component({
  selector: 'lib-material-form',
  templateUrl: './material-form.component.html',
  styleUrls: ['./material-form.component.css']
})
export class MaterialFormComponent implements OnInit {

  @Input() DataStack: string = ""

  // within the same stack, there can be multiple form. This one is the form to display
  @Input() FormName: string = ""

  // the component is refreshed when modification are performed in the back repo 
  // 
  // the checkCommitNbFromBackTimer polls the commit number of the back repo
  // if the commit number has increased, it pulls the front repo and redraw the diagram
  private commutNbFromBackSubscription: Subscription = new Subscription
  lastCommitNbFromBack = -1
  lastPushFromFrontNb = -1
  currTime: number = 0
  dateOfLastTimerEmission: Date = new Date

  public gongtableFrontRepo?: gongtable.FrontRepo

  // for selection
  selectedFormGroup: gongtable.FormGroupDB | undefined = undefined;

  // angular stuff
  myFormSample: FormGroup | undefined

  // generated form by getting info from the back
  generatedForm: FormGroup | undefined

  constructor(
    private gongtableFrontRepoService: gongtable.FrontRepoService,
    private gongtableCommitNbFromBackService: gongtable.CommitNbFromBackService,
    private formBuilder: FormBuilder,
    private formFieldStringService: gongtable.FormFieldStringService,
    private formFieldIntService: gongtable.FormFieldIntService,
    private formFieldDateService: gongtable.FormFieldDateService,
    private formFieldTimeService: gongtable.FormFieldTimeService,
    private formFieldDateTimeService: gongtable.FormFieldDateTimeService,
  ) {

  }

  ngOnInit(): void {
    this.startAutoRefresh(500); // Refresh every 500 ms (half second)

  }

  ngOnDestroy(): void {
    this.stopAutoRefresh();
  }


  stopAutoRefresh(): void {
    if (this.commutNbFromBackSubscription) {
      this.commutNbFromBackSubscription.unsubscribe();
    }
  }

  startAutoRefresh(intervalMs: number): void {
    this.commutNbFromBackSubscription = this.gongtableCommitNbFromBackService
      .getCommitNbFromBack(intervalMs, this.DataStack)
      .subscribe((commitNbFromBack: number) => {
        // console.log("OutletComponent, last commit nb " + this.lastCommitNbFromBack + " new: " + commitNbFromBack)

        if (this.lastCommitNbFromBack < commitNbFromBack) {
          const d = new Date()
          console.log("OutletComponent:", this.DataStack, d.toLocaleTimeString() + `.${d.getMilliseconds()}` +
            ", last commit increased nb " + this.lastCommitNbFromBack + " new: " + commitNbFromBack)
          this.lastCommitNbFromBack = commitNbFromBack
          this.refresh()
        }
      }
      )
  }

  refresh(): void {

    this.gongtableFrontRepoService.pull(this.DataStack).subscribe(
      gongtablesFrontRepo => {
        this.gongtableFrontRepo = gongtablesFrontRepo

        this.selectedFormGroup = undefined

        for (let form of this.gongtableFrontRepo.FormGroups_array) {
          if (form.Name == this.FormName) {
            this.selectedFormGroup = form
          }
        }

        this.generatedForm = undefined
        if (this.selectedFormGroup == undefined) {
          return
        }

        let generatedFormGroupConfig: { [key: string]: [string, any] } = {};

        if (this.selectedFormGroup.FormDivs == undefined) {
          return
        }

        this.selectedFormGroup.FormDivs.sort((t1, t2) => {
          if (t1.FormGroup_FormDivsDBID_Index && t2.FormGroup_FormDivsDBID_Index) {
            if (t1.FormGroup_FormDivsDBID_Index.Int64 > t2.FormGroup_FormDivsDBID_Index.Int64) {
              return 1;
            }
            if (t1.FormGroup_FormDivsDBID_Index.Int64 < t2.FormGroup_FormDivsDBID_Index.Int64) {
              return -1;
            }
          }
          return 0;
        })

        for (let formDiv of this.selectedFormGroup.FormDivs) {
          if (formDiv.FormFields == undefined) {
            continue
          }


          formDiv.FormFields.sort((t1, t2) => {
            if (t1.FormDiv_FormFieldsDBID_Index && t2.FormDiv_FormFieldsDBID_Index) {
              if (t1.FormDiv_FormFieldsDBID_Index.Int64 > t2.FormDiv_FormFieldsDBID_Index.Int64) {
                return 1;
              }
              if (t1.FormDiv_FormFieldsDBID_Index.Int64 < t2.FormDiv_FormFieldsDBID_Index.Int64) {
                return -1;
              }
            }
            return 0;
          })

          for (let formField of formDiv.FormFields) {
            if (formField.FormFieldString) {
              generatedFormGroupConfig[formField.Name] = [formField.FormFieldString.Value, Validators.required]
            }
            if (formField.FormFieldInt) {
              generatedFormGroupConfig[formField.Name] = [formField.FormFieldInt.Value.toString(), Validators.required]
            }
            if (formField.FormFieldDate) {
              let displayedString = formField.FormFieldDate.Value.toString().substring(0, 10)
              generatedFormGroupConfig[formField.Name] = [displayedString, Validators.required]
            }
            if (formField.FormFieldDateTime) {
              let displayedString = formField.FormFieldDateTime.Value.toString()
              generatedFormGroupConfig[formField.Name] = [displayedString, Validators.required]
            }
            if (formField.FormFieldTime) {
              let timeValue = new Date(formField.FormFieldTime.Value)
              let hours = timeValue.getUTCHours();
              let minutes = timeValue.getUTCMinutes();
              let seconds = timeValue.getUTCSeconds();

              let hoursStr = hours < 10 ? '0' + hours.toString() : hours.toString();
              let minutesStr = minutes < 10 ? '0' + minutes.toString() : minutes.toString();
              let secondsStr = seconds < 10 ? '0' + seconds.toString() : seconds.toString();

              const timeStr = `${hoursStr}:${minutesStr}:${secondsStr}`

              generatedFormGroupConfig[formField.Name] = [timeStr, Validators.required]
            }
          }
        }

        this.generatedForm = this.formBuilder.group(generatedFormGroupConfig)


        const formGroupConfig = {
          firstName: ['', Validators.required],
          lastName: ['', Validators.required],
          email: ['', [Validators.required, Validators.email]],
          age: ['', [Validators.required, Validators.min(18), Validators.max(20)]], // integer field with validation for minimum age
          password: ['', [Validators.required, Validators.minLength(6)]],
          choice: ['', Validators.required], // dropdown field
          date: ['', Validators.required],
          time: ['', Validators.required],
          acceptTerms: [false, Validators.requiredTrue] // boolean field
        };

        this.myFormSample = this.formBuilder.group(formGroupConfig);

      }
    )
  }

  options = [
    { label: 'Option 1', value: 'option1' },
    { label: 'Option 2', value: 'option2' },
    { label: 'Option 3', value: 'option3' },
    { label: 'Option 4', value: 'option4' },
    { label: 'Option 5', value: 'option5' },
  ];

  submitForm() {
    if (this.generatedForm?.valid) {
      console.log(this.generatedForm.valueChanges)

      if (this.selectedFormGroup == undefined) {
        return
      }

      if (this.selectedFormGroup.FormDivs == undefined) {
        return
      }

      for (let formDiv of this.selectedFormGroup.FormDivs) {
        if (formDiv.FormFields == undefined) {
          continue
        }
        for (let formField of formDiv.FormFields) {
          if (formField.FormFieldString) {
            let formFieldString = formField.FormFieldString
            let newValue = this.generatedForm.value[formField.Name]

            if (newValue != formFieldString.Value) {
              formFieldString.Value = newValue
              this.formFieldStringService.updateFormFieldString(formFieldString, this.DataStack).subscribe(
                () => {
                  console.log("String Form Field updated")
                }
              )
            }
          }
          if (formField.FormFieldInt) {
            let formFieldInt = formField.FormFieldInt
            let newValue: number = +this.generatedForm.value[formField.Name]

            if (newValue != formFieldInt.Value) {
              formFieldInt.Value = newValue
              this.formFieldIntService.updateFormFieldInt(formFieldInt, this.DataStack).subscribe(
                () => {
                  console.log("Int Form Field updated")
                }
              )
            }
          }
          if (formField.FormFieldDate) {
            let formFieldDate = formField.FormFieldDate

            let date = moment(this.generatedForm.value[formField.Name])
            let formattedDate = date.utc().format('YYYY-MM-DDTHH:mm:ss[Z]')
            let dateObject = moment(formattedDate).toDate()

            console.log(dateObject)
            console.log(formFieldDate.Value)

            let moment1 = moment(this.generatedForm.value[formField.Name]).startOf('day');
            let moment2 = moment(formFieldDate.Value).utc().startOf('day');

            console.log(moment1.isSame(moment2, 'day')); // Outputs: true if they are the same day

            if (!moment1.isSame(moment2, 'day')) {
              formFieldDate.Value = dateObject
              this.formFieldDateService.updateFormFieldDate(formFieldDate, this.DataStack).subscribe(
                () => {
                  console.log("Date Form Field updated")
                }
              )
            }
          }
          if (formField.FormFieldTime) {
            let formFieldTime = formField.FormFieldTime

            const [hours, minutes, seconds] = this.generatedForm.value[formField.Name].split(':').map(Number);
            const date = new Date("1970-01-01")
            date.setUTCHours(hours, minutes, seconds);
            // console.log("date for time", date.toUTCString())
            // console.log("date for backend time", new Date(formFieldTime.Value).toUTCString())

            if (date.getTime() != new Date(formFieldTime.Value).getTime()) {
              formFieldTime.Value = date
              this.formFieldTimeService.updateFormFieldTime(formFieldTime, this.DataStack).subscribe(
                () => {
                  console.log("Time Form Field updated")
                }
              )
            }
          }
          if (formField.FormFieldDateTime) {
            let formFieldDateTime = formField.FormFieldDateTime

            let newValue = this.generatedForm.value[formField.Name]

            if (newValue != formFieldDateTime.Value) {
              formFieldDateTime.Value = newValue
              this.formFieldDateTimeService.updateFormFieldDateTime(formFieldDateTime, this.DataStack).subscribe(
                () => {
                  console.log("Date Time Form Field updated")
                }
              )
            }

          }
        }
      }
    }
    if (this.myFormSample?.valid) {
      console.log(this.myFormSample.value);
      // handle your form submission
    }
  }
}
