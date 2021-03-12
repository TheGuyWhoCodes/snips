import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ActivatedRoute } from '@angular/router';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'app-raw',
  templateUrl: './raw.component.html',
  styleUrls: ['./raw.component.scss']
})
export class RawComponent implements OnInit {
    private key: string = ""
    constructor(private http: HttpClient, private activatedRoute: ActivatedRoute,private titleService: Title) { }
    public body: string = ""
    ngOnInit(): void {
        this.titleService.setTitle("snips :: "+this.activatedRoute.snapshot.params.id)
        this.key = this.activatedRoute.snapshot.params.id
        this.getPost(this.activatedRoute.snapshot.params.id).toPromise().then((x) => {
            if(x["body"] == null) {
              this.titleService.setTitle("lync.rip :: error!")
              return
            }

            this.body = x["body"]
          })
    }

    getPost(key: string) {
        return this.http.get("http://localhost:9000/api/v0/getPost/", {
        params: {
            id: key
        }
        })
    }
}
//"ReligionismWorld-ShatteringShudder"