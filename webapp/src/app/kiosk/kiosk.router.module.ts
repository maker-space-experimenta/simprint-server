import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { FilesComponent } from './components/files/files.component';
import { PrintersComponent } from './components/printers/printers.component';
import { SlideshowComponent } from './components/slideshow/slideshow.component';
import { SuccessComponent } from './components/success/success.component';


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
        component: SuccessComponent,
        path: "success"
      },
    
      {
        redirectTo: "/kiosk/slideshow/",
        path: "**"
      }
];

@NgModule({
    imports: [
        RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class KioskRoutingModule { }
