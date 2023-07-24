import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { throwError } from 'rxjs';
import { catchError, tap } from 'rxjs/operators';
import { environment } from '../environments/environment';
import { BehaviorSubject } from 'rxjs';

interface LoginResponse {
  token: string;
  // Add other properties as needed
}

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private http: HttpClient) { }

  register(user: any) {
    const apiUrl = environment.apiUrl;
    return this.http.post(`${apiUrl}/register`, user)
      .pipe(
        catchError(this.handleError)
      );
  }

  private loggedIn = new BehaviorSubject<boolean>(false);

  get isLoggedIn() {
    return this.loggedIn.asObservable();
  }

  login(user: any) {
    const apiUrl = environment.apiUrl;
    return this.http.post<LoginResponse>(`${apiUrl}/login`, user)
      .pipe(
        catchError(this.handleError),
        tap(res => {
          localStorage.setItem('auth_token', res.token);
          this.loggedIn.next(true);
        })
      );
  }

  logout() {
    localStorage.removeItem('auth_token');
    this.loggedIn.next(false);
  }

  private handleError(error: HttpErrorResponse) {
    let errorMessage = '';

    if (error.status === 400) {
      if (error.error.message === 'Geocoding failed') {
        errorMessage = 'Geocoding failed. Please check your city, state, and country.';
      } else {
        errorMessage = 'Please enter a valid email address.';
      }
    } else if (error.status === 409) {
      errorMessage = 'An account with this email address already exists.';
    } else if (error.status === 401) {
      errorMessage = 'Invalid email address or password.';
    } else {
      errorMessage = 'An unknown error occurred. Please try again later.';
    }

    return throwError(errorMessage);
  }
}
