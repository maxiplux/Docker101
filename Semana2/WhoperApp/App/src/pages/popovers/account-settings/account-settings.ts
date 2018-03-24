import { Component } from '@angular/core';
import { ViewController,NavController,IonicPage,LoadingController } from 'ionic-angular';
import {Storage} from "@ionic/storage";

@IonicPage()
@Component({
  templateUrl: 'account-settings.html'
})

export class AccountSettings {
constructor(
  public viewCtrl: ViewController,
  public navCtrl: NavController,
  public storage: Storage,
  public loadingCtrl: LoadingController) {}
close() {
    this.viewCtrl.dismiss();
  }

  onReturn() {
    this.navCtrl.pop(
    );
  }

  logout() {
  let loading = this.loadingCtrl.create({
  content: 'Login Off...'
  });

  loading.present();
  this.storage.remove('jwt');

  setTimeout(() => {
    this.navCtrl.push('LoginPage');
    loading.dismiss();
}, 3500);

}

}
