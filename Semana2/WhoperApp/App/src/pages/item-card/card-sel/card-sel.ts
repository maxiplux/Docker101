import { Component } from '@angular/core';
import { NavController } from 'ionic-angular';

// My Services

@Component({
  selector: 'page-card-sel',
  templateUrl: 'card-sel.html'
})

export class CardSel {
  // Vars
  private base64Image: string;


  constructor(
    public navCtrl: NavController
) {
//
  }
  onReturn() {
    this.navCtrl.pop(
    );
  }

}
