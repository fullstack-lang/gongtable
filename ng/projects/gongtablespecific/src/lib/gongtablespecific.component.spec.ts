import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongtablespecificComponent } from './gongtablespecific.component';

describe('GongtablespecificComponent', () => {
  let component: GongtablespecificComponent;
  let fixture: ComponentFixture<GongtablespecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongtablespecificComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongtablespecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
