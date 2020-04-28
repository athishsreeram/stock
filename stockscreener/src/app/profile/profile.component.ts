import { Component, OnInit } from '@angular/core';
import { ProfileService } from '../svc/profile.service';
import { Profile } from '../svc/Profile';
import { Quote } from '../svc/Quote';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {

  profiles: Profile;
  quotes: Quote[];
  err: String;
  symbol: String;

  constructor(private profileService: ProfileService) { }

  ngOnInit() {
      
  }

  getProfile(){

    if(this.symbol != ""){
      this.profileService.getProfiles(this.symbol).
      subscribe(data => this.profiles=data,
        error => this.err = error);

        this.profileService.getQuote(this.symbol).
        subscribe(data => this.quotes = data,
          error => this.err = error);

    }else{
      console.log("empty symbol")
    }
  }


}
