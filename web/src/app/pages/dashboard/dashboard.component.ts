import { Component, OnInit } from '@angular/core';
import {Device} from "../../services/device.service";

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {

  public selectedDevice: Device | undefined;

  constructor() { }

  ngOnInit(): void {
  }

}
