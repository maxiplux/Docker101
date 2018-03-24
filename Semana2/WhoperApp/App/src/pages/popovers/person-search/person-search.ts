import { Component } from '@angular/core';
import { ViewController,NavController } from 'ionic-angular';
@Component({
  templateUrl: 'person-search.html'
})

export class PersonSearch {
  
  images = ['Aeris_Gainsborough.jpg', 'Barret_Wallace.jpg', 'Cloud_Strife.jpg', 'VincentValentine.jpg'
  , 'Barret_Wallace.jpg', 'Cloud_Strife.jpg', 'VincentValentine.jpg'
  , 'Barret_Wallace.jpg', 'Cloud_Strife.jpg', 'VincentValentine.jpg'
  , 'Barret_Wallace.jpg', 'Cloud_Strife.jpg', 'VincentValentine.jpg'];

constructor(
  public viewCtrl: ViewController,
  public navCtrl: NavController) {}
close() {
    this.viewCtrl.dismiss();
  }

  onReturn() {
    this.navCtrl.pop(
    );
  }

}
