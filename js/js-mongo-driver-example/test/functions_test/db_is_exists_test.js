import { dbIsExists } from "../../pkg/index.js";

async function dbIsExistsTest() {
  let insertData = ["students", "test"];

  for (const item of insertData) {
    let dbName = item;
    let output = false;

    if (await dbIsExists(dbName)) {
      output = dbName + " is exists";
    } else {
      output = dbName + " not exists";
    }

    console.log(output);
  }
}

export {
  dbIsExistsTest,
}