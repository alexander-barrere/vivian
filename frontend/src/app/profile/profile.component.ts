import { Component, OnInit } from '@angular/core';
import { UserService } from '../user.service';
import { latLng, tileLayer, marker, icon } from 'leaflet';

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

  ngOnInit(): void {
    this.user = this.userService.getCurrentUser();
    if (this.user && this.user.latitude && this.user.longitude) {
      this.options.center = latLng(this.user.latitude, this.user.longitude);
      this.layers.push(marker([ this.user.latitude, this.user.longitude ], {
        icon: icon({
          iconSize: [ 25, 41 ],
          iconAnchor: [ 13, 41 ],
          iconUrl: 'leaflet/marker-icon.png',
          shadowUrl: 'leaflet/marker-shadow.png'
        })
      }));
    }
  }  
}
