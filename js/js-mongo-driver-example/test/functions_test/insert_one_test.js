import { insertOne } from "../../pkg/index.js";

async function insertOneTest() {
  let insertDocs = [
    { msg: "test msg" },
    { info: "test info" },
  ];

  let dbName = "test";
  let collectionName = "test";

  console.log(insertDocs);
  for (const doc of insertDocs) {
    let ok = await insertOne(doc, dbName, collectionName);
    if (ok) {
      console.log("insert", doc);
    }
  }
  console.log(insertDocs);
}

export {
  insertOneTest,
}