import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';

import { VivianModule } from './vivian/vivian.module';
import { VivianComponent } from './vivian/vivian.component'; // Add this import
import { StarfiresComponent } from './starfires/starfires.component';
import { UserService } from './user.service';

@NgModule({
  declarations: [
    StarfiresComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    VivianModule // Import VivianModule here
  ],
  providers: [UserService],
  bootstrap: [VivianComponent] // Use VivianComponent here
})
export class AppModule { }
