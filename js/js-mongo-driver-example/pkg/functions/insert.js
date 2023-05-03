import { client } from "../config.js";

export async function insertOne(doc, dbName, collectionName) {
  try {
    await client.connect();

    const database = client.db(dbName);
    const collection = database.collection(collectionName);

    await collection.insertOne(doc);
    return true;

  } catch (err) {
    console.error(err);
    return false;

  } finally {
    await client.close();
  }
}