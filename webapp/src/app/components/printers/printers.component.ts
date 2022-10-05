import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-printers',
  templateUrl: './printers.component.html',
  styleUrls: ['./printers.component.scss']
})
export class PrintersComponent implements OnInit {

  printers: any[] = [];
  filename: any;

  constructor(
    private http: HttpClient,
    private route: ActivatedRoute,
    private router: Router
    ) { }

  ngOnInit(): void {
    this.http.get("http://172.18.5.196:5000/api/printers").subscribe((printers: any) => {
      console.log("printers", printers)
      this.printers = printers;
    });

    this.route.queryParams
      .subscribe(params => {
        console.log(params); // { orderby: "price" }
        this.filename = params['file'];
      }
    );
  }

  print(printer: any) {
    let data = {
      file: this.filename,
      printer: printer.hostname
    };

    this.http.post("http://172.18.5.196:5000/api/jobs",  data).subscribe(m => {
      console.log(m);
      this.router.navigate(["/"]);
    });
  }

}
