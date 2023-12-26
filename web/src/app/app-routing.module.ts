import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {AppComponent} from "./app.component";
import {DefaultLayoutComponent} from "./layouts/default-layout/default-layout.component";
import {DashboardComponent} from "./pages/dashboard/dashboard.component";
import {DevicesModule} from "./pages/devices/devices.module";

const routes: Routes = [

  {
    path: "",
    component: DefaultLayoutComponent,
    children: [
      {
        path: '',
        pathMatch: 'full',
        redirectTo: 'dashboard',
      },
      {
        path: 'devices',
        loadChildren:()=> import('./pages/devices/devices.module').then(m=> m.DevicesModule)
      },
      {
        path: 'dashboard',
        component: DashboardComponent
      }
    ]
  },

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
