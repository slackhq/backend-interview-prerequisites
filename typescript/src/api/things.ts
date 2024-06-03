import { sqlConnection } from "../database";

interface Thing {
  id: number;
  name: string;
}

export async function getThings(): Promise<Thing[]> {
  const db = await sqlConnection();
  return await db.all<Thing>("SELECT * FROM `things`");
}
