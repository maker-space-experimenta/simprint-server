import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { SlideshowComponent } from './components/slideshow/slideshow.component';
import { FilesComponent } from './components/files/files.component';
import { FileCardComponent } from './components/file-card/file-card.component';
import { PrintersComponent } from './components/printers/printers.component';
import { PrinterCardComponent } from './components/printer-card/printer-card.component';
import { SliceComponent } from './components/slice/slice.component';

@NgModule({
  declarations: [
    AppComponent,
    SlideshowComponent,
    FilesComponent,
    FileCardComponent,
    PrintersComponent,
    PrinterCardComponent,
    SliceComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
