import { Component } from '@angular/core';
import {Device} from "../../../services/device.service";
import {dummyDevices} from "../../../../../test/data/devices";
import {ActivatedRoute, Router} from "@angular/router";
import {
  NgbAccordionBody,
  NgbAccordionButton, NgbAccordionCollapse,
  NgbAccordionDirective,
  NgbAccordionHeader,
  NgbAccordionItem
} from "@ng-bootstrap/ng-bootstrap";
import {DeviceManageBackupplansComponent} from "./device-manage-backupplans/device-manage-backupplans.component";

@Component({
  selector: 'app-devices-manage',
  standalone: true,
  imports: [
    NgbAccordionDirective,
    NgbAccordionHeader,
    NgbAccordionItem,
    NgbAccordionButton,
    NgbAccordionCollapse,
    NgbAccordionBody,
    DeviceManageBackupplansComponent,
  ],
  templateUrl: './devices-manage.component.html',
  styleUrl: './devices-manage.component.scss'
})
export class DevicesManageComponent {
  public device: Device;

  constructor(private router: Router, private activatedRoute: ActivatedRoute) {
    const deviceId = activatedRoute.snapshot.params["id"];
    this.device = dummyDevices.find(value => value.id == deviceId)!;

  }

}
