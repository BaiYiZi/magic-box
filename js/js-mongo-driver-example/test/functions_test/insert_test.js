import { insert } from "../../pkg/index.js";

async function insertTest() {
  let insertDocs = [
    { msg: "test msg" },
    { info: "test info" },
  ];

  let dbName = "test";
  let collectionName = "test";

  for (const doc of insertDocs) {
    let ok = await insert(doc, dbName, collectionName);
    if (ok) {
      console.log("insert", doc);
    }
  }
}

export {
  insertTest,
}