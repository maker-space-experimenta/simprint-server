import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';

@Component({
  selector: 'app-upload',
  templateUrl: './upload.component.html',
  styleUrls: ['./upload.component.scss']
})
export class UploadComponent implements OnInit {

  fileName: string = "";

  @Input() data: any;
  @Output() dataChange = new EventEmitter();

  constructor() { }

  ngOnInit(): void {
  }

  onFileSelected(e: any) {
    const file:File = e.target.files[0];

    if (file) {

        this.fileName = file.name;
this.data.file = file;
    }
  }

}
