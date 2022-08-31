import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Component, ElementRef, Input, OnInit, ViewChild } from '@angular/core';
import { Color, GCodeRenderer, SpeedColorizer } from 'gcode-viewer';

@Component({
  selector: 'app-file-card',
  templateUrl: './file-card.component.html',
  styleUrls: ['./file-card.component.scss']
})
export class FileCardComponent implements OnInit {

  @Input('data') file: any;

  content: any;
  image: any;

  constructor(
    private http: HttpClient
  ) { }

  ngOnInit(): void {
    console.log(this.file);

    this.image = "data:image/png;base64," + this.file.image;    

  }

}
