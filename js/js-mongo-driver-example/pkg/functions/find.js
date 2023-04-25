import { client } from "../config.js";

export async function find(query, dbName, collectionName) {
  try {
    await client.connect()

    let database = client.db(dbName);
    let collection = database.collection(collectionName);

    let docs = collection.find(query);
    let result = [];

    await docs.forEach(v => { result.push(v) });
    return result;

  } catch (err) {
    console.error(err);

  } finally {
    await client.close();
  }
}