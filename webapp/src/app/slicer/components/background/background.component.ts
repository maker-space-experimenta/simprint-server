import { AfterViewInit, Component, ElementRef, OnInit, ViewChild } from '@angular/core';

import * as paper from 'paper'
import { PaperOffset } from 'paperjs-offset'

@Component({
  selector: 'app-background',
  templateUrl: './background.component.html',
  styleUrls: ['./background.component.scss']
})
export class BackgroundComponent implements OnInit, AfterViewInit {

  @ViewChild('myCanvas') canvas: ElementRef | undefined;

  background = "#182e3e";
  foreground = "rgba(255, 167, 0, 0.2)";
  strokeWith = 4;
  view: paper.View | null = null
  viewWitdth = 0;
  viewHeight = 0;
  maxAmplitude = 400;
  minAmplitude = 100;
  waveLength = 100;
  animatedLines: any[] = [];


  constructor() { }

  ngOnInit(): void {
  }


  ngAfterViewInit(): void {

    paper.setup(this.canvas?.nativeElement);

    this.view = paper.project.view;
    this.viewWitdth = paper.project.view.size.width;
    this.viewHeight = paper.project.view.size.height;
    this.maxAmplitude = (this.viewHeight * 0.3);

    this.createWaves();

    setInterval(() => this.createWaves(), 10000);

    paper.project.view.onFrame = (count: number, time: number, delta: number) => {
      this.animatedLines.forEach(m => m.dashOffset += 10)
    }
  }

  createWaves() {
    this.animatedLines = [];
    paper.project.clear();

    let bg = this.setBackground();


    let p = this.createWave();
    p.dashArray = [5000];
    this.animatedLines.push(p);

    for (let i = 1; i < 2; i++) {
      this.createOffset(p, 15*i, 1500*i);

    }
  }

  createOffset(path: paper.Path, offset: number, dashOffset: number) {
    let po = PaperOffset.offsetStroke(path, offset, { cap: 'round' })
    po.fillColor = new paper.Color("rgba(0,0,0,0)");
    po.strokeColor = new paper.Color(this.foreground);
    po.strokeWidth = this.strokeWith;
    po.dashArray = [dashOffset];
    po.dashOffset = dashOffset;

    this.animatedLines.push(po);
  }


  setBackground() {
    let rect = new paper.Path.Rectangle({
      point: [0, 0],
      size: [this.viewWitdth, this.viewHeight],
      strokeColor: 'white',
      selected: true
    });

    rect.sendToBack();
    rect.fillColor = new paper.Color(this.background);

    return rect;
  }

  createCircle(p: paper.Point) {
    // let myCircle = new paper.Path.Circle(p, 5);
    // myCircle.fillColor = new paper.Color('red');
  }

  createWave() {
    let startY = (this.viewHeight * 0.3) + Math.floor(Math.random() * (this.viewHeight * 0.4));
    let endY = (this.viewHeight * 0.3) + Math.floor(Math.random() * (this.viewHeight * 0.4));
    let start: paper.Point = new paper.Point(0, startY);
    let end: paper.Point = new paper.Point(this.viewWitdth, endY);
    let length = Math.sqrt(Math.pow(Math.abs(startY - endY), 2) + Math.pow(this.viewWitdth, 2))
    console.log(length)

    let amplitude = Math.floor(Math.random() * this.maxAmplitude);
    amplitude = amplitude < this.minAmplitude ? this.minAmplitude : amplitude;
    let position = Math.floor((length * 0.15) + Math.floor(Math.random() * (length * 0.7)));

    let line = new paper.Path();
    line.add(start);
    line.add(end);

    let wave = new paper.Path();
    wave.add(start);

    let p = line.getPointAt(position);
    let ps = line.getPointAt(position - this.waveLength);
    let pps = (new paper.Path.Line(ps, p)).getPointAt((new paper.Path.Line(ps, p)).length / 2);
    let pe = line.getPointAt(position + this.waveLength);
    let ppe = (new paper.Path.Line(pe, p)).getPointAt((new paper.Path.Line(pe, p)).length / 2);


    let pt = new paper.Point(pps.x, pps.y + amplitude);
    let pb = new paper.Point(ppe.x, ppe.y - amplitude);

    wave.add(ps);
    wave.add(pt);
    wave.add(pb);

    wave.add(pe);

    wave.add(end);

    this.createCircle(p);
    this.createCircle(ps)
    this.createCircle(pps)
    this.createCircle(pe)
    this.createCircle(ppe)
    this.createCircle(pt)
    this.createCircle(pb)

    wave.strokeColor = new paper.Color(this.foreground);
    wave.strokeWidth = this.strokeWith;

    return wave;
  }

}
