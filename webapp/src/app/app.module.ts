import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { SlideshowComponent } from './components/slideshow/slideshow.component';
import { FilesComponent } from './components/files/files.component';
import { FileCardComponent } from './components/file-card/file-card.component';

@NgModule({
  declarations: [
    AppComponent,
    SlideshowComponent,
    FilesComponent,
    FileCardComponent
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
