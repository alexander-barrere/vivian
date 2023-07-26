import { Component, OnInit } from '@angular/core';
import { UserService } from '../user.service';
import { latLng, tileLayer, marker, icon } from 'leaflet';
import { Map } from 'leaflet';
import { Chart } from 'astrochart';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
  user: any;
  options = {
    layers: [
      tileLayer('http://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', { maxZoom: 18, attribution: '...' })
    ],
    zoom: 5,
    center: latLng(46.879966, -121.726909)
  };
  layers = [];

  constructor(private userService: UserService) { }

  map: Map;
  onMapReady(map: Map) {
    this.map = map;
  }

  ngOnInit(): void {
    this.user = this.userService.getCurrentUser();
    if (this.user && this.user.latitude && this.user.longitude) {
      this.options.center = latLng(this.user.latitude, this.user.longitude);
      this.layers.push(marker([ this.user.latitude, this.user.longitude ], {
        icon: icon({
          iconSize: [ 25, 41 ],
          iconAnchor: [ 13, 41 ],
          iconUrl: 'assets/leaflet/images/marker-icon.png',
          shadowUrl: 'assets/leaflet/images/marker-shadow.png'
        })
      }));
  
      // Generate the natal chart
      let birthDate = new Date(this.user.birthDate);
      let chart = new Chart({
        element: document.getElementById('chart'),
        width: 600,
        height: 600,
        birth: {
          year: birthDate.getFullYear(),
          month: birthDate.getMonth() + 1,
          day: birthDate.getDate(),
          hour: birthDate.getHours(),
          minute: birthDate.getMinutes(),
          lat: this.user.latitude,
          lng: this.user.longitude
        }
      });
    }
  }   
}
