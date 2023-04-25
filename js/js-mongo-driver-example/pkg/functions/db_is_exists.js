import { client } from "../config.js";

export async function dbIsExists(lookForDb) {
  try {
    await client.connect();

    let dbs = (await client.db().admin().listDatabases());

    for (const item of dbs.databases) {
      if (item.name === lookForDb) {
        return true;
      }
    }

    return false;

  } catch (err) {
    console.error(err);

  } finally {
    await client.close();
  }
}