import {Component, Inject, Injector, OnInit} from '@angular/core';
import {APIService} from '../../services/api.service';
import {AlertWindowService} from '../../includes/alert-window/alert-window.service';
import {createCustomElement} from '@angular/elements';
import {AlertWindowComponent} from '../../includes/alert-window/alert-window.component';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {

  constructor(
    @Inject('APIService') private api: APIService,
    injector: Injector,
    public popup: AlertWindowService
  ) {
    const PopupElement = createCustomElement(AlertWindowComponent, {injector});
    customElements.define('popup-element', PopupElement);
  }

  ngOnInit(): void {
    this.checkNavbarWidth()
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
  }

  checkNavbarWidth(): void {
    let width = document.querySelector('#picture-box').clientWidth;
    (document.querySelector('#picture-box') as HTMLDivElement).style.height = width + 'px';
    (document.querySelector('#picture-box') as HTMLDivElement).style.width = width + 'px';
    window.addEventListener('resize', () => {
      console.log('pop');
      let width = document.querySelector('#picture-box').clientWidth;
      (document.querySelector('#picture-box') as HTMLDivElement).style.height = width + 'px';
      (document.querySelector('#picture-box') as HTMLDivElement).style.width = width + 'px';
    });
  }

}
