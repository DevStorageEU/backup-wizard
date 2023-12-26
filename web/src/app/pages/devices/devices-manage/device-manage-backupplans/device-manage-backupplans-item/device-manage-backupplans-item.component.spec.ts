import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DeviceManageBackupplansItemComponent } from './device-manage-backupplans-item.component';

describe('DeviceManageBackupplansItemComponent', () => {
  let component: DeviceManageBackupplansItemComponent;
  let fixture: ComponentFixture<DeviceManageBackupplansItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DeviceManageBackupplansItemComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(DeviceManageBackupplansItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
