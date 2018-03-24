import { Component } from '@angular/core';
import { NavParams,NavController,ModalController,PopoverController,IonicPage } from 'ionic-angular';
import { Camera } from '@ionic-native/camera';


@IonicPage()
@Component({
  selector: 'page-home',
  templateUrl: 'home.html'
})

export class HomePage {
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

/*  itemsSorted = this.items.sort((item1, item2) => {
    if (item1.rating > item2.rating) {
        return -1;
    }

    if (item1.rating < item2.rating) {
        return 1;
    }

    if (item1.rating == item2.rating) {
      if (item1.price < item2.price) {
          return 1;
      }
      if (item1.price > item2.price) {
          return -1;
      }
      return 0;
    }

});*/

  presentSettingsPopover(ev) {
    let popover = this.popoverCtrl.create('PopoverImageSettings');
    popover.present({
      ev: ev
    });
  }

  presentPricetagPopover(ev) {
    let popover = this.popoverCtrl.create('PopoverPricetag');
    popover.present({
      ev: ev
    });
  }

  onClickImage() {
    let imageModal = this.modalCtrl.create('ImageSel');
    imageModal.onDidDismiss(place => {

    });
    imageModal.present();
  }

  onClickComment() {
    let imageModal = this.modalCtrl.create('ImageComments');
    imageModal.onDidDismiss(place => {

    });
    imageModal.present();
  }

  onClickCard() {
    let cardModal = this.modalCtrl.create('CardSel');
    cardModal.onDidDismiss(place => {

    });
    cardModal.present();
  }

onLoadNewItem() {
  this.navCtrl.push('ImgSelect');
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
