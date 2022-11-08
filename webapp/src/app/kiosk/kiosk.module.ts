import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';

import { SlideshowComponent } from './components/slideshow/slideshow.component';
import { FilesComponent } from './components/files/files.component';
import { FileCardComponent } from './components/file-card/file-card.component';
import { PrintersComponent } from './components/printers/printers.component';
import { PrinterCardComponent } from './components/printer-card/printer-card.component';
import { KioskRoutingModule } from './kiosk.router.module';
import { SuccessComponent } from './components/success/success.component';


@NgModule({
  declarations: [
    SlideshowComponent,
    FilesComponent,
    FileCardComponent,
    PrintersComponent,
    PrinterCardComponent,
    SuccessComponent
  ],
  imports: [
    CommonModule,
    KioskRoutingModule,
    HttpClientModule,
  ]
})
export class KioskModule { }
