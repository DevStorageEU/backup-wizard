import {Component, Input, OnInit} from '@angular/core';
import {Device, DeviceIcon} from "../../../services/device.service";

@Component({
  selector: 'app-device-table-right-sidebar',
  templateUrl: './device-table-right-sidebar.component.html',
  styleUrls: ['./device-table-right-sidebar.component.scss']
})
export class DeviceTableRightSidebarComponent implements OnInit {

  @Input()
  public device: Device | undefined;

  DeviceIcon = DeviceIcon;

  constructor() { }

  ngOnInit(): void {
  }

}
