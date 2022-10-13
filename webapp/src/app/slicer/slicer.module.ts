import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { UploadComponent } from './components/upload/upload.component';
import { SlicerRoutingModule } from './slicer.router.module';



@NgModule({
  declarations: [
    UploadComponent
  ],
  imports: [
    CommonModule,
    SlicerRoutingModule,
  ]
})
export class SlicerModule { }
