import { FormControl } from '@angular/forms';
import { Injectable } from '@angular/core';

import { AuthServiceProvider } from '../providers/auth-service/auth-service';

@Injectable()
export class UsernameValidator {
        debouncer: any;

        constructor(public authProvider: AuthServiceProvider){

        }

        checkUsername(control: FormControl): any {

          clearTimeout(this.debouncer);

          return new Promise(resolve => {

            this.debouncer = setTimeout(() => {

              this.authProvider.validateUsername(control.value).subscribe((res) => {
                if(res == 1){
                  resolve({'username In Use': true});
                } else {
                  resolve(null);
                }
              });

            }, 1000);

          });
        }

}
