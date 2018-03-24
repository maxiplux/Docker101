import { Component } from '@angular/core';
import { ViewController,NavController,NavParams } from 'ionic-angular';
import { Geolocation } from '@ionic-native/geolocation';
import { Http } from '@angular/http';
import 'rxjs/add/operator/map';

@Component({
  selector: 'page-place-select',
  templateUrl: 'place-select.html'
})
export class PlaceSelect {

noGPS:  string  = null;
places: any;
selPlace: any;
location: any;
public  lat:  number;
public  lng:  number;
url:  string  = 'https://maps.googleapis.com/maps/api/place/nearbysearch/json?';
rank: string  = 'distance';
placeType:  string  = 'store';
apiKey: string  = 'AIzaSyBtHKsUF2kwRS2QDI6V3rG5syffGWHNqvo';

  constructor(
public navCtrl: NavController,
public viewCtrl: ViewController,
public navParams: NavParams,
public http: Http,
private geolocation:  Geolocation,
  ) {
this.lat =  this.navParams.get('latitude');
this.lng =  this.navParams.get('longitude');
  }

  onSelectPlace(place) {
    this.viewCtrl.dismiss(place);
  }

  onReturn() {
    this.navCtrl.pop(
    );
  }

  ionViewWillEnter(){
         let apiCall: string  = `${this.url}location=${this.lat},${this.lng}&rankby=${this.rank}&type=${this.placeType}&key=${this.apiKey}`;
         this.places = null;
      // Get near places
         this.http.get(apiCall)
         .map(res => res.json()).subscribe(
             data => {
                 this.places = data.results;
             },
             err => {
               this.noGPS = "Can´t retrieve location";
                 console.log("Error getting near places!");
             }
         );
  }

}
