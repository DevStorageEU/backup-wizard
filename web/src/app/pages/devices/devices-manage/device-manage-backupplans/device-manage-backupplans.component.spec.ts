import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DeviceManageBackupplansComponent } from './device-manage-backupplans.component';

describe('DeviceManageBackupplansComponent', () => {
  let component: DeviceManageBackupplansComponent;
  let fixture: ComponentFixture<DeviceManageBackupplansComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DeviceManageBackupplansComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(DeviceManageBackupplansComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
