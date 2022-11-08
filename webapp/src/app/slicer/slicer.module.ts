import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { UploadComponent } from './components/upload/upload.component';
import { SlicerRoutingModule } from './slicer.router.module';
import { LayoutComponent } from './components/layout/layout.component';
import { BackgroundComponent } from './components/background/background.component';



@NgModule({
  declarations: [
    UploadComponent,
    LayoutComponent,
    BackgroundComponent
  ],
  imports: [
    CommonModule,
    SlicerRoutingModule,
  ]
})
export class SlicerModule { }
