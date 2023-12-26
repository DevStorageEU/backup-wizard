import {Component, Input} from '@angular/core';
import {Device, ProtectionStatus} from "../../../../services/device.service";
import {NgForOf, NgIf} from "@angular/common";

@Component({
  selector: 'app-device-manage-devicestatus',
  standalone: true,
  imports: [
    NgForOf,
    NgIf
  ],
  templateUrl: './device-manage-devicestatus.component.html',
  styleUrl: './device-manage-devicestatus.component.scss'
})
export class DeviceManageDevicestatusComponent {

  @Input()
  public device: Device | undefined;

  constructor() {
  }

  protected readonly ProtectionStatus = ProtectionStatus;
}
