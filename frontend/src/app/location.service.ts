import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { Observable } from "rxjs";

@Injectable({
  providedIn: "root",
})
export class LocationService {
  private apiKey: string = "89d5a7e1287b40b9b8418e3e7775e054";
  private apiUrl: string = "https://api.opencagedata.com/geocode/v1/json";

  constructor(private http: HttpClient) {}

  searchLocation(query: string): Observable<any> {
    const url = `${this.apiUrl}?q=${encodeURIComponent(query)}&key=${
      this.apiKey
    }`;
    return this.http.get(url);
  }
}
