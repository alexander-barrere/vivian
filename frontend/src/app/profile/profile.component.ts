import { Component, OnInit } from '@angular/core';
import { UserService } from '../user.service';
import { latLng, tileLayer, marker, icon } from 'leaflet';
import { Map } from 'leaflet';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
  user: any;
  natalChartPath: string;
  compositeChartPath: string;
  transitChartPath: string;
  map: Map;
  marker: any;

  constructor(private userService: UserService, private http: HttpClient) { }

  ngOnInit(): void {
    this.user = this.userService.getCurrentUser();
    this.http.get(`http://localhost:8080/profile/${this.user.id}`).subscribe((response: any) => {
      this.natalChartPath = response.natalChartPath;
      this.compositeChartPath = response.compositeChartPath;
      this.transitChartPath = response.transitChartPath;
    });
  }

  onMapReady(map: Map) {
    this.map = map;
    this.marker = marker([this.user.latitude, this.user.longitude], {
      icon: icon({
        iconSize: [ 25, 41 ],
        iconAnchor: [ 13, 41 ],
        iconUrl: 'leaflet/marker-icon.png',
        shadowUrl: 'leaflet/marker-shadow.png'
      })
    }).addTo(this.map);
  }
}
