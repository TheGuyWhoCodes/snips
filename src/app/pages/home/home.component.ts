import { Component, OnInit } from '@angular/core';
import { faCut } from '@fortawesome/free-solid-svg-icons';
import { HttpClient } from '@angular/common/http';
import { PostResponse } from 'src/app/interfaces/post.interfaces';
import { Router } from '@angular/router';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {
    faCut = faCut;
    body: string;

    constructor(private router: Router, private http: HttpClient) { }
    ngOnInit(): void {

    }

  submitPost() : void {
    this.http.post<PostResponse>("http://localhost:9000/api/v0/writeBody/", {      
        "title": "Test Title",
        "body": this.body,
        "markdown": false
      }).subscribe(data => {
        this.router.navigate([data.Id], {queryParams: {success:true}})
      });
  }

}
