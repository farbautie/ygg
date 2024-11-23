import { z } from 'zod';

const envSchema = z.object({
  PORT: z.coerce.number({
    description: '.env files convert numbers to strings, therefoore we have to enforce them to be numbers',
  })
    .positive()
    .max(65535, `options.PORT must be less than 65535`)
    .default(3000),
});

export default envSchema.parse(process.env);
