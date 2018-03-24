import { FormControl } from '@angular/forms';

export class PasswordValidator {

  static checkPassword(control: FormControl): any {

    return new Promise(resolve => {

      //Fake a slow response from server

      setTimeout(() => {
        if(control.value.length < 10){

          resolve({
            "password insecure": true
          });

        } else {
          resolve(null);
        }
      }, 1000);

    });
  }

}
