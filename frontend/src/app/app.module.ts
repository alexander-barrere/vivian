import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { AppRoutingModule } from './app-routing.module';
import { LocationService } from './location.service';
import { VivianModule } from './vivian/vivian.module';
import { LoginModule } from './login/login.module';
import { StarfiresComponent } from './starfires/starfires.component';
import { UserService } from './user.service';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { LoginComponent } from './login/login.component';
import { RouterModule } from '@angular/router';
import { AppComponent } from './app.component';

@NgModule({
  declarations: [
    AppComponent,
    StarfiresComponent,
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    MatSnackBarModule,
    AppRoutingModule,
    VivianModule,
    LoginModule,
    RouterModule
  ],
  providers: [
    UserService,
    LocationService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
