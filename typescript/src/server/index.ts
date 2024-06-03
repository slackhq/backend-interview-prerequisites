import { initializeHttp } from "./http";
import { initializeTcp } from "./tcp";

const app = initializeHttp();
const httpPort = process.env.HTTP_PORT ? parseInt(process.env.HTTP_PORT) : 3033;
app.listen(httpPort, () => {
  console.log(
    `[http-server]: HTTP server is running at http://localhost:${httpPort}`,
  );
});

const server = initializeTcp();
const tcpPort = process.env.TCP_PORT ? parseInt(process.env.TCP_PORT) : 8000;
server.listen(tcpPort, "localhost", () => {
  console.log(`[tcp-server]: TCP server is running at localhost:${tcpPort}`);
});
