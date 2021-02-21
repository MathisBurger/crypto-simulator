import {ApplicationRef, ComponentFactoryResolver, Injectable, Injector} from '@angular/core';
import {AlertWindowComponent} from './alert-window.component';

@Injectable({
  providedIn: 'root'
})
export class AlertWindowService {
  popup: any;
  constructor(private injector: Injector,
              private applicationRef: ApplicationRef,
              private componentFactoryResolver: ComponentFactoryResolver) { }

  showAsComponent(message: string, color: string) {
    const popup = document.createElement('app-alert-window');
    const factory = this.componentFactoryResolver.resolveComponentFactory(AlertWindowComponent);
    const popupComponentRef = factory.create(this.injector, [], popup);
    this.applicationRef.attachView(popupComponentRef.hostView);

    // Set the message
    popupComponentRef.instance.message = message;
    popupComponentRef.instance.color = color;

    // Add to the DOM
    document.body.appendChild(popup);
    this.popup = popup;
  }

  closePopup(): void {
    document.body.removeChild(this.popup);
  }


}
