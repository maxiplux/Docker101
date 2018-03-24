import { NgModule } from '@angular/core';
import { IonicPageModule } from 'ionic-angular';
import { AccountSettings } from './account-settings';

@NgModule({
  declarations: [
    AccountSettings,
  ],
  imports: [
    IonicPageModule.forChild(AccountSettings),
  ],
  exports: [
    AccountSettings
  ]
})
export class AccountSettingsModule {}
