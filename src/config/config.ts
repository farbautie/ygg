import { z } from 'zod';
import dotenv from 'dotenv';

dotenv.config();

const envSchema = z.object({
  NODE_ENV: z.enum(['development', 'production'], {
    description: 'The environment the app is running in',
  }).default('development'),
  STORAGE_PATH: z.string({
    description: 'The path to the storage directory',
  }),
  PORT: z.coerce.number({
    description: '.env files convert numbers to strings, therefoore we have to enforce them to be numbers',
  })
    .positive()
    .max(65535, `options.PORT must be less than 65535`)
    .default(3000),
});

export default envSchema.parse(process.env);
