import { Component,ViewChild } from '@angular/core';
import { NavController, AlertController, IonicPage } from 'ionic-angular';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Geolocation } from '@ionic-native/geolocation';

import { AuthServiceProvider } from '../../providers/auth-service/auth-service';
import { AgeValidator } from  '../../validators/age';
import { UsernameValidator } from  '../../validators/username';
import { EmailValidator } from  '../../validators/email';
import { PasswordValidator } from  '../../validators/password';
import { Storage } from '@ionic/storage';
import {Headers, Http} from "@angular/http";

@IonicPage()
@Component({
  selector: 'page-register',
  templateUrl: 'register.html',
})
export class RegisterPage {
  URL: string ='http://localhost:8081/users?'
  SIGNUP_URL: string;
  location: any;
  public  lat:  number;
  public  lng:  number;
  createSuccess = false;
  submitAttempt: boolean = false;

@ViewChild('signupSlider') signupSlider: any;

slideOneForm: FormGroup;
slideTwoForm: FormGroup;

  constructor(private nav: NavController, private auth: AuthServiceProvider,
     private alertCtrl: AlertController, public formBuilder: FormBuilder,
      public usernameValidator: UsernameValidator, private geolocation:  Geolocation,
    private storage: Storage, public emailValidator :EmailValidator,
  public http: Http) {
       this.slideOneForm = formBuilder.group({
           firstName: ['', Validators.compose([Validators.maxLength(30), Validators.pattern('[a-zA-Z ]*'), Validators.required])],
           lastName: ['', Validators.compose([Validators.maxLength(30), Validators.pattern('[a-zA-Z ]*'), Validators.required])],
           age: ['', AgeValidator.isValid]
       });

       this.slideTwoForm = formBuilder.group({
           username: ['', Validators.compose([Validators.required, Validators.pattern('[a-zA-Z]*')]), usernameValidator.checkUsername.bind(usernameValidator)],
           email: ['', Validators.compose([Validators.required]), emailValidator.checkEmail.bind(emailValidator)],
           password: ['', Validators.compose([Validators.required]), PasswordValidator.checkPassword],
           gender: ['', Validators.required],
       });
     }

      next(){
          this.signupSlider.slideNext();
      }

  save(){
      this.submitAttempt = true;

      if(!this.slideOneForm.valid){
          this.signupSlider.slideTo(0);
      }
      else if(!this.slideTwoForm.valid){
          this.signupSlider.slideTo(1);
      }
      else {
          this.SIGNUP_URL=
            this.URL+
            'user_name='+
            encodeURIComponent(this.slideOneForm.get('firstName').value)+
            '%20'+
            encodeURIComponent(this.slideOneForm.get('lastName').value)+
            '&user_login='+
            this.slideTwoForm.get('username').value+
            '&pwd='+
            this.slideTwoForm.get('password').value+
            '&email='+
            this.slideTwoForm.get('email').value+
            '&gender='+
            this.slideTwoForm.get('gender').value+
            '&location='+
            '('+this.lat+','+this.lng+')'

            this.http.post(this.SIGNUP_URL, JSON.stringify('SIGNUP_URL'))
            .map(res => res.json()).subscribe(data =>
                console.log(data)
              );
              this.showPopup("Success", "Account created.")
              this.nav.popToRoot()
      }

  }

  showPopup(title, text) {
    let alert = this.alertCtrl.create({
      title: title,
      subTitle: text,
      buttons: [
        {
          text: 'Ok',
          handler: data => {
            if (this.createSuccess) {
              this.nav.popToRoot();
            }
          }
        }
      ]
    });
    alert.present();
  }

  ionViewWillEnter(){
      //Get Location First
      this.geolocation.getCurrentPosition().then((resp) => {
        this.location = resp;
        console.log('location retrieved successfully');

         this.lat = resp.coords.latitude;
         this.lng = resp.coords.longitude;

      }).catch((error) => {
        console.log('Error getting location', error);
      });
  }
}
