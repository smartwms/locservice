import "babel-polyfill";
import Database from "./db";
import Plotter from "./plotter";

import "../sass/style.scss";

const main = async () => {
  let db = new Database();
  let sensors = await db.getSensors();
  console.log(sensors);

  let plot = new Plotter("myChart");
  for (let sensor of sensors) {
    plot.addSensor(sensor.pos);
  }

  let anchors = await db.getAnchors();
  console.log(anchors);

  for (let anchor of anchors) {
    plot.addAnchor(anchor.pos);
  }
};

main();
