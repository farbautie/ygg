import { createLogger, format, transports } from 'winston';
import config from './config';

export const logger = createLogger({
  level: config.NODE_ENV === 'development' ? 'debug' : 'info',
  format: format.combine(
    config.NODE_ENV === 'development' ? format.colorize() : format.uncolorize(),
    format.splat(),
    format.printf(
      (info) => `${info.timestamp} ${info.level} [${info.label}]: ${info.message}`
    )
  ),
  transports: [
    new transports.Console({
      stderrLevels: ['error'],
    }),
  ],
})
