import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LayoutComponent } from './components/layout/layout.component';
import { UploadComponent } from './components/upload/upload.component';

const routes: Routes = [

    {
        path: "",
        component: LayoutComponent,
        children:[
            {
                path: "upload",
                component: UploadComponent
            }, 
        ]
    },
    {
        path: "**",
        redirectTo: '/slicer/upload',
        pathMatch: 'full'
    }

];

@NgModule({
    imports: [
        RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class SlicerRoutingModule { }
