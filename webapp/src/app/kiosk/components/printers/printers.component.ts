import { HttpClient } from '@angular/common/http';
import { AfterViewInit, Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { environment } from 'src/environments/environment';

@Component({
  selector: 'app-printers',
  templateUrl: './printers.component.html',
  styleUrls: ['./printers.component.scss']
})
export class PrintersComponent implements OnInit, AfterViewInit, OnDestroy {

  printers: any[] = [];
  filename: any;
  timeoutHandle: any;

  constructor(
    private http: HttpClient,
    private route: ActivatedRoute,
    private router: Router
    ) { }

  ngOnInit(): void {
    this.http.get(environment.api + "/api/printers").subscribe((printers: any) => {
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

  ngAfterViewInit(): void {
    this.timeoutHandle = setTimeout(() => {
      this.router.navigate(["/kiosk/slideshow/"]);
    }, 60000);
  }

  ngOnDestroy(): void {
    clearTimeout(this.timeoutHandle);
  }



  print(printer: any) {
    let data = {
      file: this.filename,
      printer: printer.hostname
    };

    this.http.post(environment.api + "/api/jobs/",  data).subscribe(m => {
      console.log(m);
      this.router.navigate(["/kiosk/success"]);
    });
  }

}
