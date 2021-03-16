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

  // data for currency-view.component.html
  // This data is being handled to display
  // data of the Rest service
  currencys: CurrencyModel[];
  trades: TradeModel[];
  balance: number;
  wallets: CurrencyArrayEntry[];

  constructor(
    // implementation of the API-service and the
    // popup service. It is made for fetching data
    // and showing status popups
    @Inject('APIService') private api: APIService,
    injector: Injector,
    public popup: AlertWindowService
  ) {
    // defines the popup element for showing
    // data in the popup
    const PopupElement = createCustomElement(AlertWindowComponent, {injector});
    customElements.define('popup-element', PopupElement);
  }

  ngOnInit(): void {

    // get JWT access token for communication
    // with the API. It is using an refresh token
    // auth system.
    this.api.getAccessToken().subscribe(data => {

      if (data == 'unauthorized') {

        location.href = '/login';
      } else {

        // set the current JWT session token
        this.api.sessionToken = data.token;

        // executes data queries
        // which are containing data
        // needed in the frontend
        this.executeQueries();
      }
    })


  }

  executeQueries() {
    // queries all currencies
    this.api.getAllCurrencys().subscribe(data => {
      if (data.status) {
        this.currencys = data.data;
      } else {
        this.ngOnInit();
      }
    });

    // queries user specific balance
    this.api.getBalance().subscribe(data => {
      if (data.status) {
        this.balance = data.balance;
      } else {
        this.ngOnInit();
      }
    });

    // query all trades of user
    this.api.getAllTrades().subscribe(data => {
      if (data.status) {
        this.trades = data.data.reverse();
      } else {
        this.ngOnInit();
      }
    });

    // query wallet data of user
    this.api.getWalletForUser().subscribe(data => {
      if (data.status) {
        this.wallets = data.data;
      } else {
        this.ngOnInit();
      }
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

  // redirect to specific currency
  viewCurrency(coinID: string): void {
    location.href = '/currency-view/' + coinID
  }

  // parses unix to date
  parseTime(unix: number): string {
    return new Date(unix * 1000).toLocaleString('en-US');
  }
}
