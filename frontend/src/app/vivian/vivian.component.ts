import { Component } from "@angular/core";
import { UserService } from "../user.service";
import { MatSnackBar, MatSnackBarRef } from "@angular/material/snack-bar";
import { LocationService } from '../location.service';

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
    private locationService: LocationService,
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
  
  searchLocation(): void {
    const query = `${this.user.city}, ${this.user.state}, ${this.user.country}`;
    this.locationService.searchLocation(query).subscribe(
      (response: any) => {
        if (response.results && response.results.length > 0) {
          const result = response.results[0];
          console.log(result.formatted); // This will log the formatted address
          // Proceed with the registration process
        } else {
          console.error('Invalid location');
          // Show an error message
        }
      },
      (error) => {
        console.error(error);
        // Show an error message
      }
    );
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
