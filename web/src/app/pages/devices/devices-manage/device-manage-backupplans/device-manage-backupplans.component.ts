import { Component } from '@angular/core';
import {
    NgbAccordionBody,
    NgbAccordionButton,
    NgbAccordionCollapse,
    NgbAccordionDirective, NgbAccordionHeader, NgbAccordionItem
} from "@ng-bootstrap/ng-bootstrap";
import {
  DeviceManageBackupplansItemComponent
} from "./device-manage-backupplans-item/device-manage-backupplans-item.component";
import {BackupJobStatus} from "../../../../services/backupjob.service";

@Component({
  selector: 'app-device-manage-backupplans',
  standalone: true,
  imports: [
    NgbAccordionBody,
    NgbAccordionButton,
    NgbAccordionCollapse,
    NgbAccordionDirective,
    NgbAccordionHeader,
    NgbAccordionItem,
    DeviceManageBackupplansItemComponent
  ],
  templateUrl: './device-manage-backupplans.component.html',
  styleUrl: './device-manage-backupplans.component.scss'
})
export class DeviceManageBackupplansComponent {

  BackupJobStatus = BackupJobStatus;

}
