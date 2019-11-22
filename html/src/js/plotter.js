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

    this.chart = new Chart(ctx, {
      type: "scatter",
      data: {
        datasets: [
          {
            label: "sensors",
            data: [],
            backgroundColor: "rgba(255, 0, 0, 1)",
          },
          {
            label: "anchors",
            data: [],
            backgroundColor: "rgba(0, 0, 0, 1)",
            pointStyle: "triangle",
          },
          {
            label: "actual_pos",
            data: [],
            backgroundColor: "rgba(0, 0, 0, 1)",
          },
          {
            label: "ch37",
            data: [],
            backgroundColor: "rgba(255, 0, 0, 0.3)",
          },
          {
            label: "ch38",
            data: [],
            backgroundColor: "rgba(0, 255, 0, 0.3)",
          },
          {
            label: "ch39",
            data: [],
            backgroundColor: "rgba(0, 0, 255, 0.3)",
          },
          {
            label: "biggest",
            data: [],
            backgroundColor: "rgba(0, 0, 0, 0.3)",
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
}
