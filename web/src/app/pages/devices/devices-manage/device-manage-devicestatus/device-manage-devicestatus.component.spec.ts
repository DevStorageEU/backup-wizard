import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DeviceManageDevicestatusComponent } from './device-manage-devicestatus.component';

describe('DeviceManageDevicestatusComponent', () => {
  let component: DeviceManageDevicestatusComponent;
  let fixture: ComponentFixture<DeviceManageDevicestatusComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DeviceManageDevicestatusComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(DeviceManageDevicestatusComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
