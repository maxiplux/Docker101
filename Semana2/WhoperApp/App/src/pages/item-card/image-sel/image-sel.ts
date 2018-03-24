import { Component } from '@angular/core';
import { ViewController,NavController,NavParams,PopoverController } from 'ionic-angular';


@Component({
  selector: 'page-image-sel',
  templateUrl: 'image-sel.html'
})
export class ImageSel {

  constructor(
public navCtrl: NavController,
private navParams: NavParams,
private popoverCtrl: PopoverController
  ) {
    //
  }

  onReturn() {
    this.navCtrl.pop(
    );
  }

  onLoadImageSettings(ev) {
    let popover = this.popoverCtrl.create('PopoverImageSettings');
    popover.present({
      ev: ev
    });
  }

}
