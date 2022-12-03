import { HttpClient } from '@angular/common/http';
import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { environment } from 'src/environments/environment.prod';

@Component({
  selector: 'app-upload',
  templateUrl: './upload.component.html',
  styleUrls: ['./upload.component.scss']
})
export class UploadComponent implements OnInit {

  fileName: string = "";

  @Input() data: any;
  @Output() dataChange = new EventEmitter();

  constructor(
    private http: HttpClient
  ) { }

  ngOnInit(): void {
  }

  onFileSelected(e: any) {
    const file: File = e.target.files[0];

    if (file) {
      // this.fileName = file.name;
      // this.data.file = file;

      let formData = new FormData();
      formData.append("file", file);

      let upload_url = "http://172.18.5.196:5000/api/slicer";
      this.http.post(upload_url, formData).subscribe(result => {
        console.log(result);
      })

    }
  }

}
