import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { catchError, map, switchMap } from 'rxjs/operators';
import * as opencage from 'opencage-api-client';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private apiUrl = 'http://localhost:8080/api';
  private openCageApiKey = '89d5a7e1287b40b9b8418e3e7775e054';

  constructor(private http: HttpClient) {}

  register(user: any): Observable<any> {
    return this.http.post(`${this.apiUrl}/register`, user).pipe(
      switchMap((response: any) => {
        console.log('Making geocode call...');
        return this.geocode(`${user.city}, ${user.state}, ${user.country}`).pipe(
          map((geometry: any) => {
            response.latitude = geometry.lat;
            response.longitude = geometry.lng;
            return response;
          }),
          catchError((error: any) => {
            console.log('Geocode failed:', error.message);
            return [response];
          })
        );
      })
    );
  }

  geocode(query: string): Observable<any> {
    return new Observable((observer) => {
      opencage
        .geocode({ key: this.openCageApiKey, q: query })
        .then((data: any) => {
          if (data.status.code === 200 && data.results.length > 0) {
            observer.next(data.results[0].geometry);
          } else {
            observer.error(new Error('No results found'));
          }
          observer.complete();
        })
        .catch((error: any) => {
          observer.error(error);
          observer.complete();
        });
    });
  }
}
