import { NgModule } from '@angular/core';
import { IonicPageModule } from 'ionic-angular';
import { ImgSelect } from './img-select';

@NgModule({
  declarations: [
    ImgSelect,
  ],
  imports: [
    IonicPageModule.forChild(ImgSelect),
  ],
  exports: [
    ImgSelect
  ]
})
export class ImgSelectPageModule {}
