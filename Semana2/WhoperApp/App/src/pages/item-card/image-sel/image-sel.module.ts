import { NgModule } from '@angular/core';
import { IonicPageModule } from 'ionic-angular';
import { ImageSel } from './image-sel';

@NgModule({
  declarations: [
    ImageSel,
  ],
  imports: [
    IonicPageModule.forChild(ImageSel),
  ],
  exports: [
    ImageSel
  ]
})
export class ImageSelModule {}
