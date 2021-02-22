import {Component, ElementRef, Inject, Injector, OnInit, ViewChild} from '@angular/core';
import {APIService} from '../../services/api.service';
import {AlertWindowService} from '../../includes/alert-window/alert-window.service';
import {AlertWindowComponent} from '../../includes/alert-window/alert-window.component';
import { createCustomElement } from '@angular/elements';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})


export class RegisterComponent implements OnInit {
  username: string;
  password: string;

  constructor(
    @Inject('APIService') private api: APIService,
    injector: Injector,
    public popup: AlertWindowService
  ) {
    const PopupElement = createCustomElement(AlertWindowComponent, {injector});
    customElements.define('popup-element', PopupElement);
  }

  ngOnInit(): void {}

  register(): void {
    this.api.register(this.username, this.password).subscribe(data => {
      if (data.status) {
        location.href = '/login';
      } else {
        this.popup.showAsComponent(data.message, '#d41717');
        setTimeout(() => {
          this.popup.closePopup();
        }, 2000);
      }
    });
  }

}
