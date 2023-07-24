import { BrowserModule } from "@angular/platform-browser";
import { NgModule } from "@angular/core";
import { HttpClientModule } from "@angular/common/http";
import { AppRoutingModule } from "./app-routing.module";
import { VivianModule } from "./vivian/vivian.module";
import { LoginModule } from "./login/login.module";
import { UserService } from "./user.service";
import { MatSnackBarModule } from "@angular/material/snack-bar";
import { LoginComponent } from "./login/login.component";
import { RouterModule } from "@angular/router";
import { AppComponent } from "./app.component";
import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { AuthInterceptor } from './auth.interceptor';
import { AuthGuard } from './auth.guard';
import { ProfileComponent } from './profile/profile.component';
import { NavigationComponent } from './navigation/navigation.component';


@NgModule({
  declarations: [AppComponent, ProfileComponent, NavigationComponent],
  imports: [
    BrowserModule,
    HttpClientModule,
    MatSnackBarModule,
    AppRoutingModule,
    VivianModule,
    LoginModule,
    RouterModule,
  ],
  providers: [
    UserService,
    {
      provide: HTTP_INTERCEPTORS,
      useClass: AuthInterceptor,
      multi: true
    }
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}
