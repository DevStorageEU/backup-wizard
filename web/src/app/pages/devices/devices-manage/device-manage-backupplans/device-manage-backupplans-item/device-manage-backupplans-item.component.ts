import {Component, Input} from '@angular/core';
import {NgIf} from "@angular/common";
import {BackupJobStatus} from "../../../../../services/backupjob.service";
import {NgbProgressbar} from "@ng-bootstrap/ng-bootstrap";
import {RouterLink} from "@angular/router";


@Component({
  selector: 'app-device-manage-backupplans-item',
  standalone: true,
  templateUrl: './device-manage-backupplans-item.component.html',
  imports: [
    NgIf,
    NgbProgressbar,
    RouterLink
  ],
  styleUrl: './device-manage-backupplans-item.component.scss'
})
export class DeviceManageBackupplansItemComponent {

  @Input()
  public name: string | undefined;

  @Input()
  public expanded = false;

  @Input()
  public jobStatus: BackupJobStatus = BackupJobStatus.IDLE;

  BackupJobStatus : typeof BackupJobStatus = BackupJobStatus ;

}
