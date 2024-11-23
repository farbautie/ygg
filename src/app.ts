
import express, { NextFunction, Request, Response } from 'express';
import * as morgan from './config/morgan'
import { ApiError } from './utils/apiError';
import { errorHandler } from './middlewares/error';
import { error } from 'console';

const app = express();

app.use(morgan.successHandler)
app.use(morgan.errorHandler)

app.use(express.json())
app.use(express.urlencoded({ extended: true }))

app.get('/', (req, res) => {
  res.send('Hello World!');
});

app.use((_req: Request, _res: Response, next: NextFunction) => {
  next(new ApiError('Not found', 404));
})

app.use(errorHandler)

export default app;
