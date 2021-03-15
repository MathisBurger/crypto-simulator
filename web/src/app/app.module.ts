import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import { AppComponent } from './app.component';
import { LoginComponent } from './components/login/login.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { RegisterComponent } from './components/register/register.component';
import {FormsModule} from '@angular/forms';
import {HttpClientModule} from '@angular/common/http';
import {APIService} from './services/api.service';
import { AlertWindowComponent } from './includes/alert-window/alert-window.component';
import { NavbarComponent } from './includes/navbar/navbar.component';
import { CurrencyViewComponent } from './components/currency-view/currency-view.component';

const routes: Routes = [
  {path: '', redirectTo: 'dashboard', pathMatch: 'full'},
  {path: 'dashboard', component: DashboardComponent},
  {path: 'login', component: LoginComponent},
  {path: 'register', component: RegisterComponent},
  {path: 'currency-view/:currency', component: CurrencyViewComponent}
];

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    DashboardComponent,
    RegisterComponent,
    AlertWindowComponent,
    NavbarComponent,
    CurrencyViewComponent
  ],
  imports: [
    BrowserModule,
    RouterModule.forRoot(routes, { relativeLinkResolution: 'legacy' }),
    FormsModule,
    HttpClientModule
  ],
  providers: [
    {
      provide: 'APIService',
      useClass: APIService
    },
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
