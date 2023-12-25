import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DeviceTableRightSidebarComponent } from './device-table-right-sidebar.component';

describe('DeviceTableRightSidebarComponent', () => {
  let component: DeviceTableRightSidebarComponent;
  let fixture: ComponentFixture<DeviceTableRightSidebarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DeviceTableRightSidebarComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DeviceTableRightSidebarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
