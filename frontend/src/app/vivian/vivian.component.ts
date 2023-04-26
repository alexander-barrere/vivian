import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { UserService } from '../user.service';
import { environment } from '../../environments/environment';

@Component({
  selector: 'vivian',
  templateUrl: './vivian.component.html',
  styleUrls: ['./vivian.component.css']
})

export class VivianComponent {
  user = {
    firstName: '',
    lastName: '',
    email: '',
    password: '',
    birthDate: '',
    birthTime: '',
    city: '',
    state: '',
    country: ''
  };

  constructor(private http: HttpClient) {}

  register(): void {
    const requestBody = {
      first_name: this.user.firstName,
      last_name: this.user.lastName,
      email: this.user.email,
      password: this.user.password,
      birth_date: this.user.birthDate,
      birth_time: this.user.birthTime,
      city: this.user.city,
      state: this.user.state,
      country: this.user.country
    };

    this.http
      .post(`${environment.apiUrl}/register`, requestBody)
      .subscribe(
        (response) => {
          // Handle successful registration
          console.log('Registration successful:', response);
        },
        (error) => {
          // Handle error
          console.error('Registration error:', error);
        }
      );
  }
}
