import {AfterViewInit, Component, EventEmitter, OnInit, Output, ViewChild} from '@angular/core';
import {NgbPaginator, NgbSort, NgbTableDataSource} from 'dfx-bootstrap-table';
import {FormControl} from "@angular/forms";
import {Device, DeviceIcon, ProtectionStatus} from 'src/app/services/device.service';
import {dummyDevices} from "../../../../../test/data/devices";


@Component({
  selector: 'app-device-table',
  templateUrl: './device-table.component.html',
  styleUrls: ['./device-table.component.scss']
})
export class DeviceTableComponent implements AfterViewInit {
  DeviceIcon = DeviceIcon;
  ProtectionStatus = ProtectionStatus;

  public devices = dummyDevices;

  // Filtering
  public filter = new FormControl();

  // Sorting
  @ViewChild(NgbSort) sort!: NgbSort;
  @ViewChild(NgbPaginator) paginator!: NgbPaginator;

  @Output() public deviceSelect = new EventEmitter<Device>();

  public columnsToDisplay = ['selected', 'name', 'protection', 'connection', 'lastBackup', 'action'];
  public dataSource: NgbTableDataSource<Device> = new NgbTableDataSource(this.devices);

  ngAfterViewInit(): void {
    // Sort has to be set after template initializing
    this.dataSource.sort = this.sort;
    this.dataSource.paginator = this.paginator;

  }

  toggleSelectAll(enabled: boolean) {
    this.devices.forEach(value => {
      value.selected = enabled;
    })
  }

  select(device: Device) {
    this.deviceSelect.emit(device);
  }
}
