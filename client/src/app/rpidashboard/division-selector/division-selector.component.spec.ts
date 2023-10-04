import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DivisionSelectorComponent } from './division-selector.component';

describe('DivisionSelectorComponent', () => {
  let component: DivisionSelectorComponent;
  let fixture: ComponentFixture<DivisionSelectorComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [DivisionSelectorComponent]
    });
    fixture = TestBed.createComponent(DivisionSelectorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
