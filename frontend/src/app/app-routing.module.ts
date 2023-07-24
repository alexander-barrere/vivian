import { NgModule } from "@angular/core";
import { RouterModule, Routes } from "@angular/router";
import { LoginComponent } from "./login/login.component";
import { VivianComponent } from "./vivian/vivian.component";
import { AuthGuard } from './auth.guard';

const routes: Routes = [
  { path: "register", component: VivianComponent },
  { path: "", component: LoginComponent },
  { path: 'protected-route', component: ProtectedComponent, canActivate: [AuthGuard] },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
