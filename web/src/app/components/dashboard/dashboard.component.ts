import {Component, Inject, Injector, OnInit} from '@angular/core';
import {APIService} from '../../services/api.service';
import {AlertWindowService} from '../../includes/alert-window/alert-window.service';
import {createCustomElement} from '@angular/elements';
import {AlertWindowComponent} from '../../includes/alert-window/alert-window.component';
import {CurrencyModel} from '../../models/currency-model';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
  currencys: CurrencyModel[];

  constructor(
    @Inject('APIService') private api: APIService,
    injector: Injector,
    public popup: AlertWindowService
  ) {
    const PopupElement = createCustomElement(AlertWindowComponent, {injector});
    customElements.define('popup-element', PopupElement);
  }

  ngOnInit(): void {
    this.api.checkTokenStatus().subscribe(data => {
      if (data.status) {
        if (!data.valid) {
          location.href = '/login';
        }
      } else {
        this.popup.showAsComponent(data.message, '#d41717');
        setTimeout(() => {
          this.popup.closePopup();
        }, 1000);
      }
    });
    this.api.getAllCurrencys().subscribe(data => {
      if (data.status) {
          console.log(data.data);
          this.currencys = data.data;
      } else {
        this.popup.showAsComponent(data.message, '#d41717');
        setTimeout(() => {
          this.popup.closePopup();
        }, 1000);
      }
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
}
