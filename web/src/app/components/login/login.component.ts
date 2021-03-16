import {Component, Inject, Injector, OnInit} from '@angular/core';
import {APIService} from '../../services/api.service';
import {AlertWindowService} from '../../includes/alert-window/alert-window.service';
import {createCustomElement} from '@angular/elements';
import {AlertWindowComponent} from '../../includes/alert-window/alert-window.component';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  // defines the variables of the login
  // and register input
  // They are updated on type
  username: string;
  password: string;


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
  }

  // execute on login button click
  login(): void {
    this.api.login(this.username, this.password).subscribe(data => {
      if (data == 'OK') {
        location.href = '/dashboard';
      } else {
        this.popup.showAsComponent('401 Unauthorized', '#d41717');
        setTimeout(() => {
          this.popup.closePopup();
        }, 1500);
      }
    });
  }

}
