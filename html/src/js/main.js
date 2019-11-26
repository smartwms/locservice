import "babel-polyfill";
import Database from "./db";
import Plotter from "./plotter";

import "../sass/style.scss";

const sensor = "5ccf7fdb3643";
const tag = "dbd93de08ed7";

const main = async () => {
  let db = new Database();
  //   let lastRaw = await db.getLastRaw(sensor, tag);
  //   console.log(lastRaw);

  let plot = new Plotter("myChart");
  //   for (let raw of lastRaw) {
  //     plot.addPoint(raw.ch, raw.ts, raw.rssi);
  //   }

  db.subscribeToRaw(sensor, tag, raw => {
    plot.addPoint(raw.ch, raw.ts, raw.rssi);
  });
};

main();
