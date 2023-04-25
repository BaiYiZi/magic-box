import { MongoClient } from "mongodb";

const uri = "mongodb://root:single@127.0.0.1:27018";
const client = new MongoClient(uri);

export {
  client,
}