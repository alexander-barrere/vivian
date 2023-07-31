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
import { LeafletModule } from '@asymmetrik/ngx-leaflet';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations'; // Import LeafletModule
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatButtonModule } from '@angular/material/button';
import { MatMenuModule } from '@angular/material/menu';

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
    LeafletModule,
    BrowserAnimationsModule, // Add LeafletModule to imports array
    MatToolbarModule,
    MatButtonModule,
    MatMenuModule
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
