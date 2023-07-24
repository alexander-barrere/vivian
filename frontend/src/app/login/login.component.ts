import { Component } from '@angular/core';
import { UserService } from '../user.service';
import { Router } from '@angular/router';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  model: any = {};

  constructor(private userService: UserService, private router: Router) { }

  onSubmit() {
    this.userService.login(this.model).subscribe(
      data => {
        this.router.navigate(['/profile']);
      },
      error => {
        // Handle error, e.g. by redirecting to registration page
      }
    );
  }  
}
