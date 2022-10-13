import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

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
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
