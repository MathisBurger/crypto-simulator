import {Component, Inject, Injector, OnInit} from '@angular/core';
import {APIService} from '../../services/api.service';
import {AlertWindowService} from '../../includes/alert-window/alert-window.service';
import {createCustomElement} from '@angular/elements';
import {AlertWindowComponent} from '../../includes/alert-window/alert-window.component';
import {CurrencyModel} from '../../models/currency-model';
import {TradeModel} from '../../models/trade-model';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
  currencys: CurrencyModel[];
  trades: TradeModel[];
  balance: number;

  constructor(
    @Inject('APIService') private api: APIService,
    injector: Injector,
    public popup: AlertWindowService
  ) {
    const PopupElement = createCustomElement(AlertWindowComponent, {injector});
    customElements.define('popup-element', PopupElement);
  }

  ngOnInit(): void {
    let actionCounter = 0;
    this.api.checkTokenStatus().subscribe(data => {
      if (data.status) {
        actionCounter += 1
        if (!data.valid) {
          location.href = '/login';
        }
      } else {
        this.popup.showAsComponent(data.message, '#d41717');
        setTimeout(() => {
          this.popup.closePopup();
        }, 1000);
      }
      this.sendLoadedMessage(actionCounter);
    });

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
  }

  round(price: number, decimals: number): string {
    return price.toFixed(decimals);
  }

  parsePositive(num: string): string {
    if (parseFloat(num) > 0) {
      return '+' + num;
    } else {
      return num
    }
  }

  colorCalculator(value: string): string {
    if (value.indexOf('+') > -1) {
      return 'color: #00CA0C;';
    } else {
      return 'color: #E51F07;';
    }
  }

  sendLoadedMessage(actionCounter: number): void {
    if (actionCounter == 4) {
      this.popup.showAsComponent('successfully loaded data', '#1db004');
      setTimeout(() => {
        this.popup.closePopup();
      }, 1000);
    }
  }

  viewCurrency(coinID: string): void {
    location.href = '/currency-view/' + coinID
  }

  parseTime(unix: number): string {
    return new Date(unix * 1000).toLocaleString('en-US');
  }
}
