import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { Profile } from './Profile';
import { Observable, throwError } from 'rxjs';
import { retry, catchError } from 'rxjs/operators';
import { Quote } from './Quote';



@Injectable({
  providedIn: 'root'
})
export class ProfileService {

  constructor(private http: HttpClient) {}

  private _profileurl: string = "https://financialmodelingprep.com/api/v3/company/profile/";

  getProfiles(symbol): Observable<Profile> {
    return this.http.get<Profile>(this._profileurl+symbol)
      .pipe(retry(1),
        catchError(this.handleError)
      );
  }
  private _quoteurl: string = "https://financialmodelingprep.com/api/v3/quote/";
  
  getQuote(symbol): Observable<Quote[]> {
    return this.http.get<Quote[]>(this._quoteurl+symbol)
      .pipe(retry(1),
        catchError(this.handleError)
      );
  }

  handleError(error) {
    let errorMessage = '';
    if (error.error instanceof ErrorEvent) {
        // client-side error
        errorMessage = `Error: ${error.error.message}`;
    } else {
        // server-side error
        errorMessage = `Error Code: ${error.status}\nMessage: ${error.message}`;
    }
    console.log(errorMessage);
    return throwError(errorMessage);
 }
}
