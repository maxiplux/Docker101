import { NgModule } from '@angular/core';
import { IonicPageModule } from 'ionic-angular';
import { Showcase } from './showcase';

@NgModule({
  declarations: [
    Showcase,
  ],
  imports: [
    IonicPageModule.forChild(Showcase),
  ],
  exports: [
    Showcase
  ]
})
export class ShowcaseModule {}
