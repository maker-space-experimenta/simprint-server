import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { FilesComponent } from './components/files/files.component';
import { PrintersComponent } from './components/printers/printers.component';
import { SliceComponent } from './components/slice/slice.component';
import { SlideshowComponent } from './components/slideshow/slideshow.component';

const routes: Routes = [
  {
    component: SlideshowComponent,
    path: "slideshow"
  },
  {
    component: FilesComponent,
    path: "files"
  },
  {
    component: PrintersComponent,
    path: "printers"
  },
  {
    component: SliceComponent,
    path: "slice"
  },

  {
    redirectTo: "/slideshow",
    path: "**"
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
