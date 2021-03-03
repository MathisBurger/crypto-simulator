import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class CookieService {

  constructor() { }

  // set login credentials
  setLoginCredentials(username: string, password: string, auth_token: string): void {
    CookieService.setCookie('username', username, 7);
    CookieService.setCookie('password', password, 7);
    CookieService.setCookie('auth_token', auth_token, 7);
  }

  // return login credentials
  getLoginCredentials(): string[] {
    return [CookieService.getCookie('username'), CookieService.getCookie('auth_token')]
  }

  // get specific cookie
  private static getCookie(name: string) {
    let ca: Array<string> = document.cookie.split(';');
    let caLen: number = ca.length;
    let cookieName = `${name}=`;
    let c: string;

    for (let i: number = 0; i < caLen; i += 1) {
      c = ca[i].replace(/^\s+/g, '');
      if (c.indexOf(cookieName) == 0) {
        return c.substring(cookieName.length, c.length);
      }
    }
    return '';
  }

  // set specific cookie
  private static setCookie(name: string, value: string, expireDays: number, path: string = '') {
    let d:Date = new Date();
    d.setTime(d.getTime() + expireDays * 24 * 60 * 60 * 1000);
    let expires:string = `expires=${d.toUTCString()}`;
    let cpath:string = path ? `; path=${path}` : '';
    document.cookie = `${name}=${value}; ${expires}${cpath}`;
  }
}
