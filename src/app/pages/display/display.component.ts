import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { PostResponse } from 'src/app/interfaces/post.interfaces';
import { ActivatedRoute } from '@angular/router';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'app-display',
  templateUrl: './display.component.html',
  styleUrls: ['./display.component.scss']
})
export class DisplayComponent implements OnInit {
    public key: string = ""
    public body: string = ""

    constructor(private http: HttpClient, private activatedRoute: ActivatedRoute,private titleService: Title) { }

    ngOnInit(): void {
        this.titleService.setTitle("snips :: "+this.activatedRoute.snapshot.params.id)
        this.key = this.activatedRoute.snapshot.params.id
        this.getPost(this.key).subscribe(x => {
            this.body = x["body"]
        });
    }

  getPost(key: string) {
    return this.http.get("http://localhost:9000/api/v0/getPost/", {
    params: {
        id: key
    }
    })
}
}
