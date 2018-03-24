import { NgModule } from '@angular/core';
import { IonicPageModule } from 'ionic-angular';
import { PlaceSelect } from './place-select';

@NgModule({
  declarations: [
    PlaceSelect,
  ],
  imports: [
    IonicPageModule.forChild(PlaceSelect),
  ],
  exports: [
    PlaceSelect
  ]
})
export class PlaceSelectModule {}
