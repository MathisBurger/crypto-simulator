import { Injectable } from '@angular/core';
import {Observable, throwError} from 'rxjs';
import {DefaultResponse} from '../models/default-response';
import {HttpClient, HttpErrorResponse, HttpHeaders, HttpParams} from '@angular/common/http';
import {catchError} from 'rxjs/operators';
import {GetAllCurrencysResponse} from '../models/get-all-currencys-response';
import {GetBalanceResponse} from '../models/get-balance-response';
import {CurrencyHistoryResponse} from '../models/currency-history-response';
import {GetCurrencyResponse} from '../models/get-currency-response';
import {GetAllTradesResponse} from '../models/get-all-trades-response';
import {BuyCryptoResponse} from '../models/buy-crypto-response';
import {SellCryptoResponse} from '../models/sell-crypto-response';
import {GetWalletsForUserResponse} from '../models/get-wallets-for-user-response';
import {AccessToken} from '../models/access-token';

const BASE_URL = 'https://crypto.mathis-burger.de/api';

//const BASE_URL = 'http://127.0.0.1:8080/api';

@Injectable({
  providedIn: 'root'
})

export class APIService {

  constructor(private client: HttpClient) { }

  public sessionToken: string;

  // handles bad API response
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

  // handles the error of login functions
  // returns "FAILED" as observable
  private handleAuthError(error: HttpErrorResponse): Observable<string> {
    return new Observable<string>(subscriber => {
      subscriber.next('FAILED');
      subscriber.complete();
    });
  }

  // handles the error if a refresh token is not valid
  // only allowed for requesting new accessToken
  private handleAuthTokenError(error: HttpErrorResponse): Observable<string> {
    return new Observable<string>(subscriber => {
      subscriber.next('unauthorized');
      subscriber.complete();
    })
  }

  // This functions is used to register a new user
  // At this point no refresh tokens are needed
  register(username: string, password: string): Observable<DefaultResponse> {
    return this.client.post<DefaultResponse>(
      BASE_URL + '/register',
      {
      username: username,
      password: password
      }
    ).pipe(catchError(this.handleError));
  }

  // logs the user into the game.
  // sets automatically the send refresh token as cookie
  login(username: string, password: string): Observable<string>{
    return this.client.post<any>(
      BASE_URL + '/auth/login',
      {
        username: username,
        password: password
      },
      {withCredentials: true, responseType: 'text' as 'json'}
    ).pipe(catchError(this.handleAuthError));
  }

  // queries the API to generate a short life accessToken
  // bases on the long life refreshToken
  getAccessToken(): Observable<any> {
    return this.client.get<AccessToken>(BASE_URL + '/auth/accessToken', {
      withCredentials: true
    }).pipe(catchError(this.handleAuthTokenError));
  }

  // This auth endpoint revokes the current session
  // and deletes all data about this session in
  // the database, except trading information
  revokeSession(): Observable<any> {
    return this.client.post<any>(BASE_URL + '/auth/revokeSession', {}, {
      withCredentials: true,
      headers: new HttpHeaders({'Authorization': 'accessToken ' + this.sessionToken})
    }).pipe(catchError(this.handleAuthError));
  }


  //////////////////////////////////////////////////
  // Functions to get values and execute          //
  // actions on backend                           //
  //////////////////////////////////////////////////


  getAllCurrencys(): Observable<GetAllCurrencysResponse> {
    return this.client.get<GetAllCurrencysResponse>(BASE_URL + '/getAllCurrencys', {
      headers: new HttpHeaders({'Authorization': 'accessToken ' + this.sessionToken})
    }).pipe(catchError(this.handleError));
  }

  getBalance(): Observable<GetBalanceResponse> {
    return this.client.get<GetBalanceResponse>(BASE_URL + '/checkBalance', {
      headers: new HttpHeaders({'Authorization': 'accessToken ' + this.sessionToken})
    }).pipe(catchError(this.handleError))
  }

  fetchCurrencyHistory(name: string, time: number): Observable<CurrencyHistoryResponse> {
    let interval = '';
    let end = Math.floor(Date.now());
    let start = end - time;
    if (time <= 3600000) {
      interval = 'm1';
    } else if (time <= 14400000) {
      interval = 'm5';
    } else if (time <= 43200000) {
      interval = 'm15';
    } else if (time <= 86400000) {
      interval = 'm30';
    } else if (time <= 604800000) {
      interval = 'h6';
    } else if (time <= 1209600000) {
      interval = 'h12';
    } else {
      interval = 'd1';
    }
    return this.client.get<CurrencyHistoryResponse>('https://api.coincap.io/v2/assets/' + name + '/history?interval=' + interval + '&start=' + start + '&end=' + end)
      .pipe(catchError(this.handleError));
  }

  getCurrency(name: string): Observable<GetCurrencyResponse> {
    let params = new HttpParams();
    params = params.append('currency', name);
    return this.client.get<GetCurrencyResponse>(BASE_URL + '/getCurrency',
      {params: params, headers: new HttpHeaders({'Authorization': 'accessToken ' + this.sessionToken})})
      .pipe(catchError(this.handleError));
  }

  getAllTrades(): Observable<GetAllTradesResponse> {
    return this.client.get<GetAllTradesResponse>(BASE_URL + '/getAllTrades', {headers: new HttpHeaders({'Authorization': 'accessToken ' + this.sessionToken})})
      .pipe(catchError(this.handleError));
  }

  buyCrypto(currencyID: string, amount: number): Observable<BuyCryptoResponse> {
    return this.client.post<BuyCryptoResponse>(BASE_URL + '/buyCrypto',
      {
        currency_id: currencyID,
        amount: amount
      }, {headers: new HttpHeaders({'Authorization': 'accessToken ' + this.sessionToken})}).pipe(catchError(this.handleError));
  }

  sellCrypto(currencyID: string, amount: number): Observable<SellCryptoResponse> {
    return this.client.post<SellCryptoResponse>(BASE_URL + '/sellCrypto',
      {
        currency_id: currencyID,
        amount: amount
      }, {headers: new HttpHeaders({'Authorization': 'accessToken ' + this.sessionToken})}).pipe(catchError(this.handleError));
  }

  getWalletForUser(): Observable<GetWalletsForUserResponse> {
    return this.client.get<GetWalletsForUserResponse>(BASE_URL + '/getWalletsForUser',
      {headers: new HttpHeaders({'Authorization': 'accessToken ' + this.sessionToken})})
      .pipe(catchError(this.handleError));
  }
}
