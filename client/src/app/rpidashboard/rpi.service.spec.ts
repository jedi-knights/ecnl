import { TestBed } from '@angular/core/testing';

import { RpiService } from './rpi.service';

describe('RpiService', () => {
  let service: RpiService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(RpiService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
