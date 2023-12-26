import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DevicesAddComponent } from './devices-add.component';

describe('DevicesAddComponent', () => {
  let component: DevicesAddComponent;
  let fixture: ComponentFixture<DevicesAddComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DevicesAddComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DevicesAddComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
