import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { environment } from '../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private http: HttpClient) { }

  register(user: any) {
    const apiUrl = environment.apiUrl;
    return this.http.post(`${apiUrl}/register`, user)
      .pipe(
        catchError((error: HttpErrorResponse) => {
          let errorMessage = '';

          if (error.status === 400) {
            if (error.error.message === 'Geocoding failed') {
              errorMessage = 'Geocoding failed. Please check your city, state, and country.';
            } else {
              errorMessage = 'Please enter a valid email address.';
            }
          } else if (error.status === 409) {
            errorMessage = 'An account with this email address already exists.';
          } else {
            errorMessage = 'An unknown error occurred. Please try again later.';
          }

          return throwError(errorMessage);
        })
      );
  }
}
