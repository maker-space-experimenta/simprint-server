import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { FilesComponent } from './components/files/files.component';
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
    redirectTo: "/slideshow",
    path: "**"
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
