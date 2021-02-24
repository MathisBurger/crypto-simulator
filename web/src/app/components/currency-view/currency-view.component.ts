import {Component, Inject, Injector, OnInit} from '@angular/core';
import {APIService} from '../../services/api.service';
import {AlertWindowService} from '../../includes/alert-window/alert-window.service';
import {ActivatedRoute} from '@angular/router';
import {createCustomElement} from '@angular/elements';
import {AlertWindowComponent} from '../../includes/alert-window/alert-window.component';
import * as Chart from 'chart.js';
import {CurrencyModel} from '../../models/currency-model';

@Component({
  selector: 'app-currency-view',
  templateUrl: './currency-view.component.html',
  styleUrls: ['./currency-view.component.css']
})
export class CurrencyViewComponent implements OnInit {
  currency: string;
  time = 14400000;
  currModel: CurrencyModel;
  activeBtn: string = '4h';

  constructor(
    @Inject('APIService') private api: APIService,
    injector: Injector,
    public popup: AlertWindowService,
    private route: ActivatedRoute
  ) {
    const PopupElement = createCustomElement(AlertWindowComponent, {injector});
    customElements.define('popup-element', PopupElement);
  }

  ngOnInit(): void {
    let actionCounter = 0;
    this.currency = this.route.snapshot.paramMap.get('currency');
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
    this.chartUpdater();
    this.api.getCurrency(this.currency).subscribe(data => {
      if (data.status) {
        actionCounter += 1
        this.currModel = data.data;
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
    if (actionCounter == 2) {
      this.popup.showAsComponent('successfully loaded data', '#1db004');
      setTimeout(() => {
        this.popup.closePopup();
      }, 1000);
    }
  }

  parseDate(interval: number, unix: number): string {
      const date = new Date(unix).toLocaleString('en-US').split(', ');
      if (interval <= 86400000) {
        let spl = date[1].split(':');
        return spl[0] + ':' + spl[1];
      } else {
        if (interval > 604800000) {
          return date[0];
        }
        let spl = date[1].split(':');
        return date[0] + ' ' + spl[0] + ':' + spl[1];
      }
  }

  chartUpdater(): void {
    this.api.fetchCurrencyHistory(this.currency, this.time).subscribe(data => {
      let prices: number[] = [];
      let labels: string[] = [];
      let supplys: number[] = [];
      for (let i=0; i<data.data.length; i++) {
        prices[prices.length] = +data.data[i].priceUsd;
        labels[labels.length] = this.parseDate(this.time, data.data[i].time);
        supplys[supplys.length] = +data.data[i].circulatingSupply;
      }
      new Chart('currency-chart', {
        type: 'line',
        data: {
          datasets: [
            {
              data: prices,
              borderColor: ['rgba(12, 96, 39, 1)'],
              backgroundColor: ['rgba(12, 96, 39, 0.2)'],
              yAxisID: 'A'
            },
            {
              data: supplys,
              borderColor: ['#0D76EE'],
              yAxisID: 'B'
            }
          ],
          labels: labels
        },
        options: {
          legend: {
            display: false
          },
          responsive: true,
          scales: {
            yAxes: [
              {
                id: 'A',
                type: 'linear',
                position: 'left'
              },
              {
                id: 'B',
                type: 'linear',
                position: 'right'
              }
            ]
          }
        }
      });
    });
  }

  changeTimeRange(value: number, nowActive: string): void {
    this.time = value;
    (document.getElementById(this.activeBtn + '-btn') as HTMLDivElement).classList.remove('button-row-element-active');
    (document.getElementById(nowActive + '-btn') as HTMLDivElement).classList.add('button-row-element-active');
    this.activeBtn = nowActive;
    this.chartUpdater();
  }

  openModal(): void {
    var modal = document.querySelector('#sellModal') as HTMLDivElement;
    modal.style.display = 'block';
  }

  closeModal(): void {
    var modal = document.querySelector('#sellModal') as HTMLDivElement;
    modal.style.display = 'none';
  }

  buyCrypto(amount: string): void {
    this.api.buyCrypto(this.currency, +amount).subscribe(data => {
      if (data.status) {
        this.closeModal();
        this.popup.showAsComponent(data.message, '#1db004');
        setTimeout(() => {
          this.popup.closePopup();
        }, 1000);
      } else {
        this.popup.showAsComponent(data.message, '#d41717');
        setTimeout(() => {
          this.popup.closePopup();
        }, 1000);
      }
    })
  }
}
