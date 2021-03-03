import {Component, Inject, Injector, OnInit} from '@angular/core';
import {APIService} from '../../services/api.service';
import {AlertWindowService} from '../../includes/alert-window/alert-window.service';
import {createCustomElement} from '@angular/elements';
import {AlertWindowComponent} from '../../includes/alert-window/alert-window.component';
import {CurrencyModel} from '../../models/currency-model';
import {TradeModel} from '../../models/trade-model';
import {CurrencyArrayEntry} from '../../models/currency-array-entry';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
  currencys: CurrencyModel[];
  trades: TradeModel[];
  balance: number;
  wallets: CurrencyArrayEntry[];

  constructor(
    @Inject('APIService') private api: APIService,
    injector: Injector,
    public popup: AlertWindowService
  ) {
    const PopupElement = createCustomElement(AlertWindowComponent, {injector});
    customElements.define('popup-element', PopupElement);
  }

  ngOnInit(): void {

    // counter for successful API requests
    let actionCounter = 0;

    // checks token
    this.api.checkTokenStatus().subscribe(data => {
        actionCounter += 1
        if (!data.valid) {
          location.href = '/login';
        }
      this.sendLoadedMessage(actionCounter);
    });

    // queries all currencies
    this.api.getAllCurrencys().subscribe(data => {
      if (data.status) {
        actionCounter += 1
          this.currencys = data.data;
      } else {
        this.popup.showAsComponent(data.message, '#d41717');
        setTimeout(() => {
          this.popup.closePopup();
        }, 1000);
      }
      this.sendLoadedMessage(actionCounter);
    });

    // queries user specific balance
    this.api.getBalance().subscribe(data => {
      if (data.status) {
        actionCounter += 1
        this.balance = data.balance;
      } else {
        this.popup.showAsComponent(data.message, '#d41717');
        setTimeout(() => {
          this.popup.closePopup();
        }, 1000);
      }
      this.sendLoadedMessage(actionCounter);
    });

    // query all trades of user
    this.api.getAllTrades().subscribe(data => {
      if (data.status) {
        actionCounter += 1
        this.trades = data.data.reverse();
      } else {
        this.popup.showAsComponent(data.message, '#d41717');
        setTimeout(() => {
          this.popup.closePopup();
        }, 1000);
      }
      this.sendLoadedMessage(actionCounter);
    });

    // query wallet data of user
    this.api.getWalletForUser().subscribe(data => {
      if (data.status) {
        actionCounter += 1
        this.wallets = data.data;
      } else {
        this.popup.showAsComponent(data.message, '#d41717');
        setTimeout(() => {
          this.popup.closePopup();
        }, 1000);
      }
      this.sendLoadedMessage(actionCounter);
    })
  }

  // rounds value to decimals
  round(price: number, decimals: number): string {
    return price.toFixed(decimals);
  }

  // adds '+' if value is positive
  parsePositive(num: string): string {
    if (parseFloat(num) > 0) {
      return '+' + num;
    } else {
      return num
    }
  }

  // calculates color by number
  colorCalculator(value: string): string {
    if (value.indexOf('+') > -1) {
      return 'color: #00CA0C;';
    } else {
      return 'color: #E51F07;';
    }
  }

  // send successful message if all ngOnInit queries were successful
  sendLoadedMessage(actionCounter: number): void {
    if (actionCounter == 5) {
      this.popup.showAsComponent('successfully loaded data', '#1db004');
      setTimeout(() => {
        this.popup.closePopup();
      }, 1000);
    }
  }

  // redirect to specific currency
  viewCurrency(coinID: string): void {
    location.href = '/currency-view/' + coinID
  }

  // parses unix to date
  parseTime(unix: number): string {
    return new Date(unix * 1000).toLocaleString('en-US');
  }
}
