import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { VivianComponent } from './vivian.component';

@NgModule({
  declarations: [VivianComponent],
  imports: [
    CommonModule,
    FormsModule // Import FormsModule here
  ],
  exports: [VivianComponent]
})
export class VivianModule { }
