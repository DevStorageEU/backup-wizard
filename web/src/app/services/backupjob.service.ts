import { Injectable } from '@angular/core';
export enum BackupJobStatus {
  DISABLED, IDLE, ACTIVE, WARNING, ERROR
}
@Injectable({
  providedIn: 'root'
})
export class BackupjobService {

  constructor() { }
}
