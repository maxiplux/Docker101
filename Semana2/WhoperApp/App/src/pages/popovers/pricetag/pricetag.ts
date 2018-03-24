import { Component } from '@angular/core';
import { ViewController } from 'ionic-angular';
@Component({
  templateUrl: 'pricetag.html'
})

export class PopoverPricetag {
constructor(public viewCtrl: ViewController) {}
close() {
    this.viewCtrl.dismiss();
  }
}
