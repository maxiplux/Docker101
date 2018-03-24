import { NgModule } from '@angular/core';
import { IonicPageModule } from 'ionic-angular';
import { CardSel } from './card-sel';

@NgModule({
  declarations: [
    CardSel,
  ],
  imports: [
    IonicPageModule.forChild(CardSel),
  ],
  exports: [
    CardSel
  ]
})
export class CardSelModule {}
