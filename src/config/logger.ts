import { createLogger, format, transports } from 'winston';
import config from './config';

export const logger = createLogger({
  level: config.NODE_ENV === 'development' ? 'debug' : 'info',
  format: format.combine(
    config.NODE_ENV === 'development' ? format.colorize() : format.uncolorize(),
    format.splat(),
    format.printf(({ level, message }) => `${level}: ${message}`)
  ),
  transports: [
    new transports.Console({
      stderrLevels: ['error'],
    }),
  ],
})
