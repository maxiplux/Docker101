import { Component } from '@angular/core';
import { NavController,IonicPage } from 'ionic-angular';

@IonicPage()
@Component({
  selector: 'page-profile',
  templateUrl: 'profile.html'
})
export class ProfilePage {

images = ['Aeris_Gainsborough.jpg', 'Barret_Wallace.jpg', 'Cloud_Strife.jpg', 'VincentValentine.jpg'
, 'Barret_Wallace.jpg', 'Cloud_Strife.jpg', 'VincentValentine.jpg'
, 'Barret_Wallace.jpg', 'Cloud_Strife.jpg', 'VincentValentine.jpg'
, 'Barret_Wallace.jpg', 'Cloud_Strife.jpg', 'VincentValentine.jpg'];
  constructor(
    public navCtrl: NavController

  ) {

  }
  onAddPerson() {
    this.navCtrl.push('PersonSearch');
  }

  onLoadSettings() {
    this.navCtrl.push('AccountSettings');
  }

}
