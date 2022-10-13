import { HttpClient } from '@angular/common/http';
import { AfterViewInit, Component, ElementRef, OnInit, ViewChild } from '@angular/core';
// import { GCodeRenderer, Color, SpeedColorizer } from 'gcode-viewer';
import {QRCode } from 'qrcode'

@Component({
  selector: 'app-files',
  templateUrl: './files.component.html',
  styleUrls: ['./files.component.scss']
})
export class FilesComponent implements OnInit, AfterViewInit {

  @ViewChild('qrCodeCanvas')qrCodeCanvas: ElementRef | undefined;

  files: any[] = [];

  constructor(
    private http: HttpClient
  ) { }

  updateFiles() {
    this.http.get("http://172.18.5.196:5000/api/files/local").subscribe((files: any) => {
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

    // QRCode

    // QRCode.toCanvas(this.qrCodeCanvas, 'sample text', (err: any) => {
    //   console.log(err);
    // })
  }

}
