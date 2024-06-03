import { resetDb, sqlConnection } from "../database";
import { getThings } from "./things";

describe("things api", () => {
  beforeEach(() => {
    resetDb();
  });

  async function insertThing(id: number, name: string) {
    const db = await sqlConnection();
    await db.run("INSERT INTO `things` (`id`, `name`) VALUES ($id, $name)", {
      $id: id,
      $name: name,
    });
  }

  it("returns all of the things when they exist", async () => {
    await insertThing(1, "thing 1");
    await insertThing(2, "thing 2");

    const response = await getThings();

    expect(response).toEqual([
      {
        id: 1,
        name: "thing 1",
      },
      {
        id: 2,
        name: "thing 2",
      },
    ]);
  });

  it("returns nothing when they dont exist", async () => {
    const response = await getThings();
    expect(response).toEqual([]);
  });
});
