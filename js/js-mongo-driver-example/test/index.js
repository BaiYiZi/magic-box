import { insertTest } from "./functions_test/insert_test.js";
import { dbIsExistsTest } from "./functions_test/db_is_exists_test.js";
import { findTest } from "./functions_test/find_test.js";

const testFileName = process.argv.slice()[2]

switch (testFileName) {
  case "db_is_exists_test":
    dbIsExistsTest();
    break;

  case "find_test":
    findTest();
    break;

  case "insert_test":
    insertTest();
    break;

  default:
    console.log("input arg is not support");
    break;
}