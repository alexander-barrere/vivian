import { NgModule } from "@angular/core";
import { RouterModule, Routes } from "@angular/router";
import { LoginComponent } from "./login/login.component";
import { VivianComponent } from "./vivian/vivian.component";
import { AuthGuard } from './auth.guard';

const routes: Routes = [
  { path: "register", component: VivianComponent },
  { path: "", component: LoginComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
