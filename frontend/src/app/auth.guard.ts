import { Injectable } from "@angular/core";
import { Router } from "@angular/router"; // Add CanActivate here
import { UserService } from "./user.service";

@Injectable({
  providedIn: "root",
})
export class AuthGuard  { // Use CanActivate here
  constructor(private userService: UserService, private router: Router) {}

  canActivate(): boolean { // Use canActivate here
    if (!this.userService.isLoggedIn) {
      this.router.navigate(["/login"]);
      return false;
    }
    return true;
  }
}
