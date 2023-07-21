import { TestBed } from '@angular/core/testing';

import { GongtablespecificService } from './gongtablespecific.service';

describe('GongtablespecificService', () => {
  let service: GongtablespecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongtablespecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
