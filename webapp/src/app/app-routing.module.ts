import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { Path } from 'three';

const routes: Routes = [
  {
    path: "slicer",
    loadChildren: () => {
      return new Promise((resolve, reject) => {
        import('./slicer/slicer.module').then(m => {
          resolve(m.SlicerModule);
        });
      });
    }
  },
  {
    path: "kiosk",
    loadChildren: () => {
      return new Promise((resolve, reject) => {
        import('./kiosk/kiosk.module').then(m => {
          resolve(m.KioskModule);
        });
      });
    }
  },
  {
    redirectTo: "/kiosk/slideshow",
    path: "**"
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
