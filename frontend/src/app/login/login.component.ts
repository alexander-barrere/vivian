import { Component } from '@angular/core';
import { UserService } from '../user.service';  // Make sure to import your UserService

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  model: any = {};

  constructor(private userService: UserService) { }  // Inject the UserService

  onSubmit() {
    this.userService.login(this.model).subscribe(
      data => {
        // Handle successful login, e.g. by redirecting to profile page
      },
      error => {
        // Handle error, e.g. by redirecting to registration page
      }
    );
  }
}
