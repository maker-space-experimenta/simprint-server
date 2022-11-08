import { HttpClient } from '@angular/common/http';
import { AfterViewInit, Component, ElementRef, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
// import { GCodeRenderer, Color, SpeedColorizer } from 'gcode-viewer';
import {QRCode } from 'qrcode'
import { environment } from 'src/environments/environment';

@Component({
  selector: 'app-files',
  templateUrl: './files.component.html',
  styleUrls: ['./files.component.scss']
})
export class FilesComponent implements OnInit, AfterViewInit, OnDestroy {

  @ViewChild('qrCodeCanvas')qrCodeCanvas: ElementRef | undefined;

  files: any[] = [];
  timeoutHandle: any;

  constructor(
    private http: HttpClient,
    private router: Router
  ) { }
  

  updateFiles() {
    this.http.get(environment.api + "/api/files/local").subscribe((files: any) => {
      if (files) {
        console.log(files)
        this.files = files.data;
      }
    });
  }

  ngOnInit(): void {
    this.updateFiles();
    setInterval(() => this.updateFiles(), 10000)
  }

  ngAfterViewInit(): void {
    this.timeoutHandle = setTimeout(() => {
      this.router.navigate(["/kiosk/slideshow/"]);
    }, 60000);
  }

  ngOnDestroy(): void {
    clearTimeout(this.timeoutHandle);
  }

}
