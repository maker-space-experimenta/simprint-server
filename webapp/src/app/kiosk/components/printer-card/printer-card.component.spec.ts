import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PrinterCardComponent } from './printer-card.component';

describe('PrinterCardComponent', () => {
  let component: PrinterCardComponent;
  let fixture: ComponentFixture<PrinterCardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PrinterCardComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PrinterCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
