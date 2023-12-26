import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DevicesAddComponent } from './devices-add/devices-add.component';
import { DevicesDeleteComponent } from './devices-delete/devices-delete.component';
import {RouterModule, Routes} from "@angular/router";
import {DefaultLayoutComponent} from "../../layouts/default-layout/default-layout.component";
import {FormsModule} from "@angular/forms";
import {NgbTooltipModule} from "@ng-bootstrap/ng-bootstrap";



const routes: Routes = [
      {
        path: 'add',
        component: DevicesAddComponent
      },
      {
        path: 'delete',
        component: DevicesDeleteComponent
      }

];

@NgModule({
  declarations: [
    DevicesAddComponent,
    DevicesDeleteComponent
  ],
  imports: [
    CommonModule,
    RouterModule.forChild(routes),
    FormsModule,
    NgbTooltipModule
  ]
})
export class DevicesModule { }
