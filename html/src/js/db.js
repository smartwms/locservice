import firebase from "firebase";
import "firebase/firestore";

export default class Database {
  constructor() {
    firebase.initializeApp({
      apiKey: "AIzaSyBsROSlzwaz13xPf9x0WOmuAhK--hbWx4A",
      authDomain: "smartwms.firebaseapp.com",
      databaseURL: "https://smartwms.firebaseio.com",
      projectId: "smartwms",
      storageBucket: "smartwms.appspot.com",
      messagingSenderId: "707215956702",
      appId: "1:707215956702:web:dd4afbba939e1c8e8e004f",
    });

    this._db = firebase.firestore();
  }

  async getMeasures(tag) {
    let query = await this._db
      .collection("measures")
      .where("tag", "==", tag)
      .get();

    query.forEach(doc => {
      console.log(`${doc.id} =>`, doc.data());
    });
  }

  async getSensors() {
    let query = await this._db.collection("sensors").get();

    let sensors = [];
    query.forEach(doc => {
      console.log(`${doc.id} =>`, doc.data());
      sensors.push(doc.data());
    });

    return sensors;
  }

  async getAnchors() {
    let query = await this._db.collection("anchors").get();

    let anchors = [];
    query.forEach(doc => {
      console.log(`${doc.id} =>`, doc.data());
      anchors.push(doc.data());
    });

    return anchors;
  }
}
