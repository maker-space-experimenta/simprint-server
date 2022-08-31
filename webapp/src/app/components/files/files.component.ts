import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
// import { GCodeRenderer, Color, SpeedColorizer } from 'gcode-viewer';

@Component({
  selector: 'app-files',
  templateUrl: './files.component.html',
  styleUrls: ['./files.component.scss']
})
export class FilesComponent implements OnInit {

  files: any[] = [];

  constructor(
    private http: HttpClient
  ) { }

  updateFiles() {
    this.http.get("http://localhost:5000/api/files/local").subscribe((files: any) => {
      console.log(files)
      this.files = files;

    });
  }

  ngOnInit(): void {
    
    this.updateFiles();
    setInterval(() => this.updateFiles(), 10000)
  }

}
