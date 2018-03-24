import { Component } from '@angular/core';
import {  ModalController ,NavController, NavParams,IonicPage} from 'ionic-angular';
import { Geolocation } from '@ionic-native/geolocation';
import { Http } from '@angular/http';
import 'rxjs/add/operator/map';

@IonicPage()
@Component({
  selector: 'page-new-item',
  templateUrl: 'new-item.html',
})
export class NewItemPage {
rating: any;
selectedRating: any;
image: string;
places: any;
itemPlace:  string;
selPlace: any;
location: any;
public  lat:  number;
public  lng:  number;
url:  string  = 'https://maps.googleapis.com/maps/api/place/nearbysearch/json?';
rank: string  = 'distance';
placeType:  string  = 'store';
apiKey: string  = 'AIzaSyBtHKsUF2kwRS2QDI6V3rG5syffGWHNqvo';

  constructor(
    public http: Http,
    private geolocation:  Geolocation,
    private navCtrl: NavController,
    private navParams: NavParams,
    public modalCtrl: ModalController
  ){
  this.image =  this.navParams.get('selImg');
  }

ionViewWillEnter(){
    //Get Location First
    this.geolocation.getCurrentPosition().then((resp) => {
      this.location = resp;
      console.log('location retrieved successfully');

       this.lat = resp.coords.latitude;
       this.lng = resp.coords.longitude;

       let apiCall: string  = `${this.url}location=${this.lat},${this.lng}&rankby=${this.rank}&type=${this.placeType}&key=${this.apiKey}`;

       this.places = null;
       this.itemPlace = null;
    // Get near places
       this.http.get(apiCall)
       .map(res => res.json()).subscribe(
           data => {
              this.places = data.results;
              this.itemPlace = data.results[0].name;
               if (this.selPlace == null) {
                 this.selPlace  = this.itemPlace;
               }
           },
           err => {
               console.log("Error getting near places!");
           }
       );

    }).catch((error) => {
      console.log('Error getting location', error);
    });
}

onLoadPlaceSelect() {
  let profileModal = this.modalCtrl.create('PlaceSelect',{latitude: this.lat,longitude: this.lng});
  profileModal.onDidDismiss(place => {
    this.selPlace = place  ==  null  ? this.itemPlace  : place.name;
  });
  profileModal.present();
}

onSetRating(rating){
this.selectedRating = (rating/100)*5;
console.log(this.selectedRating);
}

}

//import { File } from '@ionic-native/file';
//File.listDir(cordova.file.applicationDirectory, 'carpeta/subcarpeta').then(
//  (files) => { // (files : FileEntry[]) =>
//    // Aquí los usas
//    console.log(files[0].name)
//  }
//).catch(
//  (err) => {
//
//  }
//);
