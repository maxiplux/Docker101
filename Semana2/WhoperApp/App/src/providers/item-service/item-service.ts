import { Storage } from  '@ionic/storage';
import { Injectable } from '@angular/core';

@Injectable()
export class ItemServiceProvider  {
  private items: { rating: number,  price: number }[] = [];
  constructor (private storage: Storage) {}

  addItem(item: { rating: number,price: number }) {
    this.items.push(item);
    this.storage.set('items', this.items);
  }

  getItems() {
  return this.storage.get('items')
    .then(
      (items) => {
        this.items = items == null ? [] : items;
        return this.items.slice();
      }
    );
  }

  rmItem() {
    this.storage.remove('items');
  }
}
