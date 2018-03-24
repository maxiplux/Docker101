import { Component } from '@angular/core';
import { NavController, AlertController, LoadingController, Loading, IonicPage } from 'ionic-angular';

import { AuthServiceProvider } from '../../providers/auth-service/auth-service';

@IonicPage()
@Component({
  selector: 'page-login',
  templateUrl: 'login.html',
})
export class LoginPage {
  loading: Loading;
  registerCredentials = { nameoremail: '', password: '' };
  private authjwt: { id: number,  user_login: string, salt: string, token: string }

  constructor(private nav: NavController, private auth: AuthServiceProvider,
    private alertCtrl: AlertController, private loadingCtrl: LoadingController) {

    }

  public createAccount() {
    this.nav.push('RegisterPage');
  }

  public login() {
    if (this.registerCredentials.nameoremail.indexOf('@') > -1)
    {
      this.loginWithEmail()
    } else {
      this.loginWithUserName()
    }
  }

  public loginWithEmail() {
    this.showLoading()
    this.auth.loginE(this.registerCredentials).subscribe(allowed => {
      if (allowed) {
        this.nav.setRoot('TabsPage');
      } else {
        this.showError("Access Denied");
      }
    },
      error => {
        this.showError(error);
      });
  }

  public loginWithUserName() {
    this.showLoading()
    this.auth.loginU(this.registerCredentials).subscribe(allowed => {
      if (allowed) {
        this.nav.setRoot('TabsPage');
      } else {
        this.showError("Access Denied");
      }
    },
      error => {
        this.showError(error);
      });
  }

  showLoading() {
    this.loading = this.loadingCtrl.create({
      content: 'Please wait...',
      dismissOnPageChange: true
    });
    this.loading.present();
  }

  showError(text) {
    this.loading.dismiss();

    let alert = this.alertCtrl.create({
      title: 'Fail',
      subTitle: text,
      buttons: ['OK']
    });
    alert.present(prompt);
  }
}
