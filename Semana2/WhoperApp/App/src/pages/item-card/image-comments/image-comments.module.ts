import { NgModule } from '@angular/core';
import { IonicPageModule } from 'ionic-angular';
import { ImageComments } from './image-comments';

@NgModule({
  declarations: [
    ImageComments,
  ],
  imports: [
    IonicPageModule.forChild(ImageComments),
  ],
  exports: [
    ImageComments
  ]
})
export class ImageCommentsModule {}
