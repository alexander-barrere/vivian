import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { LocationService } from './location.service';
import { VivianModule } from './vivian/vivian.module';
import { VivianComponent } from './vivian/vivian.component'; // Add this import
import { StarfiresComponent } from './starfires/starfires.component';
import { UserService } from './user.service';
import { MatSnackBarModule } from '@angular/material/snack-bar';

@NgModule({
  declarations: [
    StarfiresComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    MatSnackBarModule,
    VivianModule // Import VivianModule here
  ],
  providers: [
    UserService,
    LocationService
    ],
  bootstrap: [VivianComponent] // Use VivianComponent here
})
export class AppModule { }
