import * as sc from "./modules/serverConnector.js"
import * as dm from "./modules/domManipulator.js";

// Now, usually the app just works fine, but since I'm in development
// and occasionally I want to call reloadData() from the browser console
// I apparently have to do this. I have no idea what's going on
self.sc = sc
self.dm = dm


sc.reloadData()
// setInterval(sc.reloadData, 20000)