import { find } from "../../pkg/index.js";

async function findTest() {
  let query = {};
  let dbName = "test";
  let collectionName = "test";

  let docs = await find(query, dbName, collectionName);
  console.log(docs);
}

export {
  findTest,
}