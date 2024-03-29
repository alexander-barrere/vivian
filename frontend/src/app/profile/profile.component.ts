import { Component, OnInit } from '@angular/core';
import { UserService } from '../user.service';
import { latLng, tileLayer, marker, icon, Layer } from 'leaflet';
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
  options = {
    layers: [
      tileLayer('http://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', { maxZoom: 18, attribution: '...' })
    ],
    zoom: 5,
    center: latLng(46.879966, -121.726909)
  };
  layers: Layer[] = [];

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
        iconUrl: '../../assets/leaflet/images/marker-icon.png',
        shadowUrl: '../../assets/leaflet/images/marker-shadow.png'
      })
    }).addTo(this.map);
  }
}
