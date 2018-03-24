import { Component } from '@angular/core';
import { NavController,IonicPage } from 'ionic-angular';
import { Camera } from '@ionic-native/camera';

@IonicPage()
@Component({
  selector: 'page-img-select',
  templateUrl: 'img-select.html'
})

export class ImgSelect {
  private base64Image: string;

images = ['Aeris_Gainsborough.jpg', 'Barret_Wallace.jpg', 'Cloud_Strife.jpg', 'VincentValentine.jpg'];
preImg = this.images[0];
  constructor(public navCtrl: NavController,private camera: Camera) {
    //
  }

onLoadNewItem(preImg) {
  this.navCtrl.push('NewItemPage',{
    selImg: preImg
  });
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

onSelectImg(img) {
this.preImg = img;
}

}
