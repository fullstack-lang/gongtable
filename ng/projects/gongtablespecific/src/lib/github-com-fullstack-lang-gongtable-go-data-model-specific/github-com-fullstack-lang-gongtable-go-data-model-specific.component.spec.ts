import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GithubComFullstackLangGongtableGoDataModelSpecificComponent } from './github-com-fullstack-lang-gongtable-go-data-model-specific.component';

describe('GithubComFullstackLangGongtableGoDataModelSpecificComponent', () => {
  let component: GithubComFullstackLangGongtableGoDataModelSpecificComponent;
  let fixture: ComponentFixture<GithubComFullstackLangGongtableGoDataModelSpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GithubComFullstackLangGongtableGoDataModelSpecificComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GithubComFullstackLangGongtableGoDataModelSpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
