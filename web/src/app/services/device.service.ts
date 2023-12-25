import { Injectable } from '@angular/core';



export enum ProtectionStatus {
  OK, WARNING, ERROR
}
export enum DeviceIcon {
  SERVER, COMPUTER, MOBILE
}
export interface OperatingSystem {
  osName: string; // windoof
  cpu: string; // eg. AMD Ryzen X 3900X
  ram: string; // eg. 32 GB
  totalHarddrive: string; //eg. 2 TB
}
export interface Device {
  id: number,
  selected: boolean,
  name: string;
  icon: DeviceIcon,
  protection: ProtectionStatus,
  connection: boolean,
  lastBackup?: string,
  ip_addresses: string[]; // all IP-Addresses (primary and some fallback/secondary addresses)
  agent: string; //eg. rsync version 11.xx
  machineIdentifier: string; // uuid or something like that
  os: OperatingSystem,
  comment?: string; //open string
  backupPlans: any[]; //all backup plans of this device
}



@Injectable({
  providedIn: 'root'
})
export class DeviceService {

  constructor() { }
}
