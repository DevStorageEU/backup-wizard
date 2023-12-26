import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { DashboardComponent } from './pages/dashboard/dashboard.component';
import { DefaultLayoutComponent } from './layouts/default-layout/default-layout.component';
import { LeftSidebarComponent } from './layouts/components/left-sidebar/left-sidebar.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { TopHeaderbarComponent } from './layouts/components/top-headerbar/top-headerbar.component';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { DeviceTableComponent } from './layouts/components/device-table/device-table.component';
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {DfxPaginationModule, DfxSortModule, DfxTableModule} from "dfx-bootstrap-table";
import { DeviceTableRightSidebarComponent } from './layouts/components/device-table-right-sidebar/device-table-right-sidebar.component';
import {BrowserAnimationsModule} from "@angular/platform-browser/animations";
import {ToastrModule} from "ngx-toastr";

@NgModule({
  declarations: [
    AppComponent,
    DashboardComponent,
    DefaultLayoutComponent,
    LeftSidebarComponent,
    TopHeaderbarComponent,
    DeviceTableComponent,
    DeviceTableRightSidebarComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    NgbModule,
    FontAwesomeModule,
    ReactiveFormsModule,
    DfxTableModule,
    DfxPaginationModule,
    DfxSortModule,
    FormsModule,
    BrowserAnimationsModule,
    ToastrModule.forRoot()
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
