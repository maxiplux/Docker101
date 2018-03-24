import { Injectable } from '@angular/core';
import {Observable} from 'rxjs/Observable';
import { Http } from '@angular/http';
import {Storage} from "@ionic/storage";
import 'rxjs/add/operator/map';

@Injectable()
export class AuthServiceProvider {
  users: any;
  URL: string ='http://localhost:8081/auth/validate-'
  SIGNUP_URL: string;
  private authjwt: { id: number,  user_login: string, salt: string, token: string }

  constructor(public http: Http, private storage: Storage){
  }

  validateUsername(username){
    return  this.http.get('http://localhost:8081/auth/validate-username/' + username).map(res => res.json());
  }

  validateEmail(domain){
    return  this.http.get('http://api.mailtest.in/v1/' + domain).map(res => res.json());
  }

  public loginE(credentials) {
    if (credentials.nameoremail === null || credentials.password === null) {
      return Observable.throw("Please insert credentials");
    } else {
      return Observable.create(observer => {
          //
          this.SIGNUP_URL = this.URL+'email?'+'email='+credentials.nameoremail+'&pwd='+credentials.password
          this.http.post(this.SIGNUP_URL, JSON.stringify('SIGNUP_URL'))
          .map(res => res.json()).subscribe(data =>
              {this.authjwt = data;
                if (this.authjwt == null) {
                      observer.next(false);
                      observer.complete();
                } else {
                    this.authSuccess(this.authjwt)
                      observer.next(true);
                      observer.complete();
                }
              }
            );
          //
      });
    }
  }

  public loginU(credentials) {
      if (credentials.nameoremail === null || credentials.password === null) {
        return Observable.throw("Please insert credentials");
      } else {
        return Observable.create(observer => {
            //
            this.SIGNUP_URL = this.URL+'username?'+'user_login='+credentials.nameoremail+'&pwd='+credentials.password
            this.http.post(this.SIGNUP_URL, JSON.stringify('SIGNUP_URL'))
            .map(res => res.json()).subscribe(data =>
                {this.authjwt = data;
                  if (this.authjwt == null) {
                        observer.next(false);
                        observer.complete();
                  } else {
                      this.authSuccess(this.authjwt)
                        observer.next(true);
                        observer.complete();
                  }
                }
              );
            //
        });
      }
    }


  public register(credentials) {
    if (credentials.email === null || credentials.password === null || credentials.username === null) {
      return Observable.throw("Please insert credentials");
    } else {
      return Observable.create(observer => {
        observer.next(true);
        observer.complete();
      });
    }
  }

  authSuccess(jwt) {
     this.storage.set('jwt', jwt);
   }

   getjwt() {
   return this.storage.get('jwt')
     .then(
       (jwt) => {
         this.authjwt = jwt;
         return this.authjwt;
       }
     );
   }

  public logout() {
    return Observable.create(observer => {
      this.storage.remove('jwt')
      observer.next(true);
      observer.complete();
    });
  }
}
