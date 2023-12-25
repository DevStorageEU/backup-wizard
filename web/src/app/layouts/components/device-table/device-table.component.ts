import {AfterViewInit, Component, EventEmitter, OnInit, Output, ViewChild} from '@angular/core';
import {NgbPaginator, NgbSort, NgbTableDataSource} from 'dfx-bootstrap-table';
import {FormControl} from "@angular/forms";
import {Device, DeviceIcon, ProtectionStatus} from 'src/app/services/device.service';


@Component({
  selector: 'app-device-table',
  templateUrl: './device-table.component.html',
  styleUrls: ['./device-table.component.scss']
})
export class DeviceTableComponent implements AfterViewInit {
  DeviceIcon = DeviceIcon;
  ProtectionStatus = ProtectionStatus;

  public devices: Device[] = [
    {
      id: 1,
      name: "Test Server 1",
      connection: true,
      icon: DeviceIcon.SERVER,
      selected: false,
      lastBackup: "25.12.2023 13:55",
      protection: ProtectionStatus.OK,
      agent: "OpenSSH_6.6.1p1",
      ip_addresses: ["127.0.0.2"],
      os: {
        osName: "Ubuntu 22.04",
        cpu: "AMD Epyc 7402P",
        ram: "64 GB RAM",
        totalHarddrive: "1 TB"
      },
      machineIdentifier: "33970533-952c-43e9-a999-27f5ba1c655c",
      backupPlans: []
    },
    {
      id: 2,
      name: "Test Server 2",
      connection: true,
      icon: DeviceIcon.SERVER,
      selected: false,
      lastBackup: "22.12.2023 09:33",
      protection: ProtectionStatus.WARNING,
      agent: "OpenSSH_6.6.1p1",
      ip_addresses: ["127.0.0.2", "182.322.33.134"],
      os: {
        osName: "Ubuntu 22.04",
        cpu: "AMD Epyc 7402P",
        ram: "64 GB RAM",
        totalHarddrive: "1 TB"
      },
      comment: "Hello World, this is a description of this device. How are you? How long could this text can be?",
      machineIdentifier: "33970533-952c-43e9-a999-27f5ba1c655c",
      backupPlans: []
    },
    {
      id: 3,
      name: "Test Computer",
      connection: false,
      icon: DeviceIcon.COMPUTER,
      selected: false,
      lastBackup: "25.12.2023 11:15",
      protection: ProtectionStatus.ERROR,
      agent: "OpenSSH_6.6.1p1",
      ip_addresses: ["127.0.0.2"],
      os: {
        osName: "Ubuntu 22.04",
        cpu: "AMD Epyc 7402P",
        ram: "64 GB RAM",
        totalHarddrive: "1 TB"
      },
      machineIdentifier: "33970533-952c-43e9-a999-27f5ba1c655c",
      backupPlans: []
    },
  ]

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

    this.select(this.devices[0]) //TODO: remove me (just for debugging)
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
