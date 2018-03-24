import { Component } from '@angular/core';
import { Platform } from 'ionic-angular';
import { StatusBar } from '@ionic-native/status-bar';
import { SplashScreen } from '@ionic-native/splash-screen';
import { AuthServiceProvider } from '../providers/auth-service/auth-service'
import {Storage} from "@ionic/storage";

@Component({
  templateUrl: 'app.html'
})
export class Whoper {

  rootPage:any;
  private authjwt: { id: number,  user_login: string, salt: string, token: string }

  constructor(platform: Platform, statusBar: StatusBar,
    splashScreen: SplashScreen, storage: Storage) {

      storage.get('jwt').then(
             (jwt) => {
               this.authjwt = jwt;
                     if (this.authjwt == null) {
                       this.rootPage = 'LoginPage';
                     } else {
                       this.rootPage = 'TabsPage';
                     }
             }
           );


    platform.ready().then(() => {
      statusBar.styleDefault();
      splashScreen.hide();
    });
  }
}
