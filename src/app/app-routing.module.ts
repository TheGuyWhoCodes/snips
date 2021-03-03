import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { HomeComponent } from './pages/home/home.component';
import { DisplayComponent } from './pages/display/display.component';
import { RawComponent } from './pages/raw/raw.component';


const routes: Routes = [
  {
    path:'',
    component: HomeComponent
  },
  {
    path:':id',
    component: DisplayComponent
  },
  {
    path:'raw/:id',
    component: RawComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
