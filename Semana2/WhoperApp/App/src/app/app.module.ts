import { NgModule, ErrorHandler } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { IonicApp, IonicModule, IonicErrorHandler } from 'ionic-angular';

import { Whoper } from './app.component';
import { Geolocation } from '@ionic-native/geolocation';
import { Camera } from '@ionic-native/camera';

import { StatusBar } from '@ionic-native/status-bar';
import { SplashScreen } from '@ionic-native/splash-screen';

import { IonicStorageModule } from '@ionic/storage';
import { IonicImageViewerModule } from 'ionic-img-viewer'
import { Network } from '@ionic-native/network';
import { HttpModule } from '@angular/http';
import { AuthServiceProvider } from '../providers/auth-service/auth-service';
import { ItemServiceProvider } from '../providers/item-service/item-service';
import { LocationServiceProvider } from '../providers/location-service/location-service';
import { UsernameValidator } from '../validators/username';
import { EmailValidator } from '../validators/email';


@NgModule({
  declarations: [
    Whoper

  ],
  imports: [
    BrowserModule,
    IonicImageViewerModule,
    IonicStorageModule.forRoot(),
    HttpModule,
    IonicModule.forRoot(Whoper,{tabsHideOnSubPages:true})

  ],
  bootstrap: [IonicApp],
  entryComponents: [
    Whoper
  ],
  providers: [
    StatusBar,
    SplashScreen,
    Geolocation,
    Camera,
    Network,
    {provide: ErrorHandler, useClass: IonicErrorHandler},
    AuthServiceProvider,
    ItemServiceProvider,
    LocationServiceProvider,
    UsernameValidator,
    EmailValidator

  ]
})
export class AppModule {}
