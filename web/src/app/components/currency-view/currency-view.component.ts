import {Component, Inject, Injector, OnInit} from '@angular/core';
import {APIService} from '../../services/api.service';
import {AlertWindowService} from '../../includes/alert-window/alert-window.service';
import {ActivatedRoute} from '@angular/router';
import {createCustomElement} from '@angular/elements';
import {AlertWindowComponent} from '../../includes/alert-window/alert-window.component';

@Component({
  selector: 'app-currency-view',
  templateUrl: './currency-view.component.html',
  styleUrls: ['./currency-view.component.css']
})
export class CurrencyViewComponent implements OnInit {
  currency: string;

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
  }

  sendLoadedMessage(actionCounter: number): void {
    if (actionCounter == 1) {
      this.popup.showAsComponent('successfully loaded data', '#1db004');
      setTimeout(() => {
        this.popup.closePopup();
      }, 1000);
    }
  }

}
