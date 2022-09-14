import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

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
    private route: ActivatedRoute
    ) { }

  ngOnInit(): void {
    this.http.get("http://localhost:5000/api/printers").subscribe((printers: any) => {
      console.log(printers)
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
      printer: printer.url
    };

    this.http.post("http://localhost:5000/api/print",  data).subscribe(m => {
      console.log(m);
    });
  }

}
