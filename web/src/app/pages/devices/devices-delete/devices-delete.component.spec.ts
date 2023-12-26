import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DevicesDeleteComponent } from './devices-delete.component';

describe('DevicesDeleteComponent', () => {
  let component: DevicesDeleteComponent;
  let fixture: ComponentFixture<DevicesDeleteComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DevicesDeleteComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DevicesDeleteComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
