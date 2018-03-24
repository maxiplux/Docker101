import { Component } from '@angular/core';
import { ViewController,NavController,NavParams,PopoverController } from 'ionic-angular';
import {PopoverImageSettings} from "../../popovers/image-settings/image-settings";


@Component({
  selector: 'page-image-comments',
  templateUrl: 'image-comments.html'
})
export class ImageComments {

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

}
