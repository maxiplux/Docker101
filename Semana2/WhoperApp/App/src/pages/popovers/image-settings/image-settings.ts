import { Component } from '@angular/core';
import { ViewController } from 'ionic-angular';
@Component({
  templateUrl: 'image-settings.html'
})

export class PopoverImageSettings {
constructor(public viewCtrl: ViewController) {}
close() {
    this.viewCtrl.dismiss();
  }
}
