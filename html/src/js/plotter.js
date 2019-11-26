import "chart.js";

export default class Plotter {
  constructor(canvasId) {
    let ctx = document.getElementById(canvasId).getContext("2d");

    this.labels = {
      sensors: 0,
      anchors: 1,
      actual_pos: 2,
      ch37: 3,
      ch38: 4,
      ch39: 5,
      avg: 6,
    };

    this.values = {
      37: [],
      38: [],
      39: [],
    };

    this.chart = new Chart(ctx, {
      type: "line",
      options: {
        responsive: false,
        scales: {
          xAxes: [
            {
              type: "time",
              display: false,
            },
          ],
          yAxes: [
            {
              display: true,
              ticks: {
                beginAtZero: true,
                min: -100,
                max: -30,
              },
            },
          ],
        },
      },
      data: {
        datasets: [
          {
            label: "ch39",
            data: this.values[39],
            borderColor: "rgba(255, 0, 0, 1)",
            fill: false,
          },
          {
            label: "ch38",
            data: this.values[38],
            borderColor: "rgba(0, 255, 0, 1)",
            fill: false,
          },
          {
            label: "ch37",
            data: this.values[37],
            borderColor: "rgba(0, 0, 255, 1)",
            fill: false,
          },
        ],
      },
    });
  }

  addSensor(pos) {
    console.log(`Adding sensor at (${pos.x},${pos.y}) to chart.`);
    this.chart.data.datasets[this.labels["sensors"]].data.push(pos);
    this.chart.update();
  }

  addAnchor(pos) {
    console.log(`Adding anchor at (${pos.x},${pos.y})  to chart.`);
    this.chart.data.datasets[this.labels["anchors"]].data.push(pos);
    this.chart.update();
  }

  addPoint(ch, ts, rssi) {
    console.log(`Adding point at (${ts},${rssi}) to ${ch}.`);

    this.values[ch].push({ x: ts, y: -rssi });
    if (this.values[ch].length > 15) {
      this.values[ch].shift();
    }

    this.chart.update();
  }
}
