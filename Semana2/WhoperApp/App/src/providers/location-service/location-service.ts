import { Geolocation } from '@ionic-native/geolocation';
import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import 'rxjs/add/operator/map';

@Injectable()
export class LocationServiceProvider {
private location: {lat: number, lng:  number};
private resp: any;

  constructor (
      private geolocation: Geolocation,
      public http: Http
  ) {

  }

  getLocation() {
return  this.geolocation.getCurrentPosition().then((resp) => {
    console.log('location retrieved successfully');
  this.location.lat = this.resp.coords.latitude;
  this.location.lng = this.resp.coords.longitude;
  return this.location;
  }).catch((error) => {
    console.log('Error getting location', error);
  });
  }
}
