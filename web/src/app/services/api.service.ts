import { Injectable } from '@angular/core';
import {Observable, throwError} from 'rxjs';
import {DefaultResponse} from '../models/default-response';
import {HttpClient, HttpErrorResponse} from '@angular/common/http';
import {catchError} from 'rxjs/operators';
import {LoginResponse} from '../models/login-response';

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
}
