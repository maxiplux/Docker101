import { FormControl } from '@angular/forms';
import { Injectable } from '@angular/core';

import { AuthServiceProvider } from '../providers/auth-service/auth-service';

@Injectable()
export class EmailValidator {
        debouncer: any;

        constructor(public authProvider: AuthServiceProvider){

        }

        checkEmail(control: FormControl): any {

          clearTimeout(this.debouncer);

          return new Promise(resolve => {

            this.debouncer = setTimeout(() => {
              var email = control.value;
              var domain = email.substring(email.lastIndexOf("@") +1);
              console.log(domain)
              this.authProvider.validateEmail(domain).subscribe((res) => {
                console.log(res)
                if(res.code != "01"){
                  resolve({'Email not valid': true});
                } else {
                  resolve(null);
                }
              });

            }, 1000);

          });
        }

}
