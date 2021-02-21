import {Component, EventEmitter, HostBinding, Input, OnInit, Output} from '@angular/core';
import { animate, state, style, transition, trigger } from '@angular/animations';

@Component({
  selector: 'app-alert-window',
  templateUrl: './alert-window.component.html',
  styleUrls: ['./alert-window.component.css'],
})
export class AlertWindowComponent{
  @Input() message: string;
  @Input() color: string;

  constructor() {}

}
