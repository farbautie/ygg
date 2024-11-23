import { Server } from "node:http";
import app from "./app";
import config from "./config/config";
import { logger } from "./config/logger";

const server: Server = app.listen(config.PORT, () => {
  logger.info(`Server listening on port ${config.PORT}`);
});


function exitHandler() {
  logger.info("Cleaning up");
  server.close(() => {
    logger.info("Server closed");
    process.exit(1);
  });
}

// Handle SIGHUP
process.on("SIGINT", () => {
  logger.info("Received SIGINT, shutting down gracefully");
  server.close(() => {
    logger.info("Server closed");
  });
});

process.on("uncaughtException", (err) => {
  logger.error("Uncaught exception", err);
  exitHandler()
});

process.on("unhandledRejection", (err) => {
  logger.error("Unhandled rejection", err);
  exitHandler()
});
