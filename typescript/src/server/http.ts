import express, { Express, Request, Response, json } from "express";

export function initializeHttp(): Express {
  const app = express();
  app.use(json());

  app.get(`/api/hello.get`, async (req: Request, res: Response) => {
    res.json({ ok: true, msg: "hello", params: req.query })
  });

  app.post(`/api/hello.post`, async (req: Request, res: Response) => {
    res.json({ ok: true, msg: "hello", params: req.query, body: req.body });
  });

  return app;
}
