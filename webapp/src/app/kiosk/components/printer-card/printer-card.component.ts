import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-printer-card',
  templateUrl: './printer-card.component.html',
  styleUrls: ['./printer-card.component.scss']
})
export class PrinterCardComponent implements OnInit {

  @Input('data') printer: any;

  constructor() { }

  ngOnInit(): void {
  }

}
