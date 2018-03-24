import { Component,ViewChild, ElementRef } from '@angular/core';
import { NavParams,NavController,ModalController,PopoverController,IonicPage } from 'ionic-angular';
import { Camera } from '@ionic-native/camera';
// My Services
import {ImgSelect} from "../new-item/img-select/img-select";
// My pages
import {ImageSel} from "../item-card/image-sel/image-sel";
import {CardSel} from "../item-card/card-sel/card-sel";
import {ImageComments} from "../item-card/image-comments/image-comments";
import {PopoverImageSettings} from "../popovers/image-settings/image-settings";
import {PopoverPricetag} from "../popovers/pricetag/pricetag";

@IonicPage()
@Component({
  selector: 'page-showcase',
  templateUrl: 'showcase.html'
})

export class Showcase {
  // Vars
  private base64Image: string;

  constructor(
    public navCtrl: NavController,
    private camera: Camera,
    public modalCtrl: ModalController,
    private navParams: NavParams,
    private popoverCtrl: PopoverController
) {
//
  }
  presentSettingsPopover(ev) {
    let popover = this.popoverCtrl.create(PopoverImageSettings);
    popover.present({
      ev: ev
    });
  }

  presentPricetagPopover(ev) {
    let popover = this.popoverCtrl.create(PopoverPricetag);
    popover.present({
      ev: ev
    });
  }

  onClickImage() {
    let imageModal = this.modalCtrl.create(ImageSel);
    imageModal.onDidDismiss(place => {

    });
    imageModal.present();
  }

  onClickComment() {
    let imageModal = this.modalCtrl.create(ImageComments);
    imageModal.onDidDismiss(place => {

    });
    imageModal.present();
  }

  onClickCard() {
    let cardModal = this.modalCtrl.create(CardSel);
    cardModal.onDidDismiss(place => {

    });
    cardModal.present();
  }

onLoadNewItem() {
  this.navCtrl.push(ImgSelect);
}

onTakePicture() {
  this.camera.getPicture({
    quality: 50,
    destinationType: this.camera.DestinationType.DATA_URL,
    encodingType: this.camera.EncodingType.JPEG,
    mediaType: this.camera.MediaType.PICTURE
  }).then((imageData) => {
   // imageData is either a base64 encoded string or a file URI
   // If it's base64:
   this.base64Image = 'data:image/jpeg;base64,' + imageData;
  }, (err) => {
   // Handle error
   console.log(err);
  });
}

}
