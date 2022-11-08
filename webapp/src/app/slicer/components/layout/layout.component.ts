import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-layout',
  templateUrl: './layout.component.html',
  styleUrls: ['./layout.component.scss']
})
export class LayoutComponent implements OnInit {

  state = "upload";

  data: any = {
    file: null,
    preset: "pla_01mm_fine",
    support: false
  };

  constructor(
    private http: HttpClient
  ) { }

  ngOnInit(): void {
    
  }

  sendFormData() {
    let formData: FormData = new FormData();
    formData.append("file", this.data.file);
    formData.append("preset", this.data.preset);
    formData.append("support", this.data.support);

    const upload$ = this.http.post("/api/thumbnail-upload", formData);
    upload$.subscribe();
  }

}
