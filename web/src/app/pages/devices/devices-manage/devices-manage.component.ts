import { Component } from '@angular/core';
import {Device} from "../../../services/device.service";
import {dummyDevices} from "../../../../../test/data/devices";

@Component({
  selector: 'app-devices-manage',
  standalone: true,
  imports: [],
  templateUrl: './devices-manage.component.html',
  styleUrl: './devices-manage.component.scss'
})
export class DevicesManageComponent {
  public device: Device | undefined;

  constructor() {
    this.device = dummyDevices[0];

  }

}
