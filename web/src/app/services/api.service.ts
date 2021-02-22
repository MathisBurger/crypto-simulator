import { Injectable } from '@angular/core';
import {Observable, throwError} from 'rxjs';
import {DefaultResponse} from '../models/default-response';
import {HttpClient, HttpErrorResponse, HttpParams} from '@angular/common/http';
import {catchError} from 'rxjs/operators';
import {LoginResponse} from '../models/login-response';
import {TokenStatusResponse} from '../models/token-status-response';
import {CookieService} from './cookie.service';
import {GetAllCurrencysResponse} from '../models/get-all-currencys-response';
import {GetBalanceResponse} from '../models/get-balance-response';

const BASE_URL = 'http://localhost:8080/api';

@Injectable({
  providedIn: 'root'
})

export class APIService {
  constructor(private client: HttpClient) { }

  private handleError(error: HttpErrorResponse) {
    if (error.error instanceof ErrorEvent) {
      console.error('An error occurred:', error.error.message);
    } else {
      console.error(
        `Backend returned code ${error.status}, ` +
        `body was: ${error.error}`);
    }
    return throwError(
      'Something bad happened; please try again later.');
  }


  register(username: string, password: string): Observable<DefaultResponse> {
    return this.client.post<DefaultResponse>(
      BASE_URL + '/register',
      {
      username: username,
      password: password
      }
    ).pipe(catchError(this.handleError));
  }

  login(username: string, password: string): Observable<LoginResponse> {
    return this.client.post<LoginResponse>(
      BASE_URL + '/login',
      {
        username: username,
        password: password
      }
    ).pipe(catchError(this.handleError));
  }

  checkTokenStatus(): Observable<TokenStatusResponse> {
    let creds = new CookieService().getLoginCredentials();
    let params = new HttpParams();
    params = params.append('username', creds[0]);
    params = params.append('token', creds[1]);
    return this.client.get<TokenStatusResponse>(BASE_URL + '/checkTokenStatus', {params: params})
      .pipe(catchError(this.handleError));
  }

  getAllCurrencys(): Observable<GetAllCurrencysResponse> {
    let creds = new CookieService().getLoginCredentials();
    let params = new HttpParams();
    params = params.append('username', creds[0]);
    params = params.append('token', creds[1]);
    return this.client.get<GetAllCurrencysResponse>(BASE_URL + '/getAllCurrencys', {params: params})
      .pipe(catchError(this.handleError));
  }

  getBalance(): Observable<GetBalanceResponse> {
    let creds = new CookieService().getLoginCredentials();
    let params = new HttpParams();
    params = params.append('username', creds[0]);
    params = params.append('auth_token', creds[1]);
    return this.client.get<GetBalanceResponse>(BASE_URL + '/checkBalance', {params: params})
      .pipe(catchError(this.handleError))
  }


}
