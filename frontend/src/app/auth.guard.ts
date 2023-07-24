import { Injectable } from "@angular/core";
import { CanActivateChild, Router } from "@angular/router"; // changed this line
import { UserService } from "./user.service";

@Injectable({
  providedIn: "root",
})
export class AuthGuard implements CanActivateChild { // changed this line
  constructor(private userService: UserService, private router: Router) {}

  canActivateChild(): boolean { // changed this line
    if (!this.userService.isLoggedIn) {
      this.router.navigate(["/login"]);
      return false;
    }
    return true;
  }
}
