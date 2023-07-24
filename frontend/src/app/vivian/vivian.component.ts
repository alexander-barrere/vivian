import { Component } from "@angular/core";
import { UserService } from "../user.service";
import { MatSnackBar, MatSnackBarRef } from "@angular/material/snack-bar";

@Component({
  selector: "vivian",
  templateUrl: "./vivian.component.html",
  styleUrls: ["./vivian.component.css"],
})
export class VivianComponent {
  errorMessage = "";
  returnedLatitude: string = "";
  returnedLongitude: string = "";
  snackBarRef: MatSnackBarRef<any> | null = null;
  user = {
    firstName: "",
    lastName: "",
    email: "",
    password: "",
    birthDate: "",
    birthTime: "",
    city: "",
    state: "",
    country: "",
  };

  constructor(
    private userService: UserService,
    private snackBar: MatSnackBar
  ) {}

  register(): void {
    if (this.snackBarRef) {
      this.snackBarRef.dismiss();
    }
  
    const requestBody = {
      first_name: this.user.firstName,
      last_name: this.user.lastName,
      email: this.user.email,
      password: this.user.password,
      birth_date: this.user.birthDate,
      birth_time: this.user.birthTime,
      city: this.user.city,
      state: this.user.state,
      country: this.user.country,
    };
  
    this.userService.register(requestBody).subscribe(
      (response: any) => {
        this.returnedLatitude = response.latitude;
        this.returnedLongitude = response.longitude;
        this.errorMessage = ''; // Clear the error message
        this.showSuccessMessage("Registration successful");
      },
      (error) => {
        console.error(error);
        this.showErrorMessage(error);
      }
    );
  }  

  showErrorMessage(errorMessage: string) {
    if (this.snackBarRef) {
      this.snackBarRef.dismiss();
    }
  
    this.snackBarRef = this.snackBar.open(errorMessage, "Close", {
      duration: 3000,
      verticalPosition: "top",
      horizontalPosition: "center",
      panelClass: "mat-snackbar-error",
    });
  
    this.snackBarRef.afterOpened().subscribe(() => {
      if (this.snackBarRef) {
        this.snackBarRef.dismiss();
      }
    });
  }
  
  showSuccessMessage(message: string) {
    if (this.snackBarRef) {
      this.snackBarRef.dismiss();
    }
  
    this.snackBarRef = this.snackBar.open(message, "Close", {
      duration: 3000,
      verticalPosition: "top",
      horizontalPosition: "center",
      panelClass: "mat-snackbar-success",
    });
  
    this.snackBarRef.afterOpened().subscribe(() => {
      if (this.snackBarRef) {
        this.snackBarRef.dismiss();
      }
    });
  }  
}
