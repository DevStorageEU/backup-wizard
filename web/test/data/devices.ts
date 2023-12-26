import {Device, DeviceIcon, ProtectionStatus} from "../../src/app/services/device.service";

export const dummyDevices: Device[] = [
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
