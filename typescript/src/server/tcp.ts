import { createServer, Server } from "net";

export function initializeTcp(): Server {
  const server = createServer();

  server.on("connection", (socket) => {
    if (!socket.remotePort) {
      return;
    }

    function writeResponse(ok: boolean, res: object) {
      socket.write(
        JSON.stringify({
          ok,
          ...res,
        }) + "\n\n"
      );
    }

    console.log("CONNECTED: " + socket.remoteAddress + ":" + socket.remotePort);

    socket.on("data", (data) => {
      if (!socket.remotePort) {
        return;
      }
      console.log("DATA " + socket.remoteAddress + ": " + data);

      const requests = data.toString("utf8").trim().split("\n");
      for (const rawRequest of requests) {
        try {
          const req = JSON.parse(rawRequest);
          switch (req.type) {
            case "hello.get":
            case "hello.post":
              writeResponse(true, { msg: "hello", req });
              break;
            default:
              writeResponse(false, { error: "unknown_type" });
          }
        } catch (e) {
          console.error("unknown_error", { error: e });
          writeResponse(false, { error: "unknown_error", error_detail: e });
        }
      }
    });

    // Add a 'close' event handler to this instance of socket
    socket.on("close", () => {
      if (!socket.remotePort) {
        return;
      }
      console.log("GOODBYE");
    });
  });
  return server;
}
