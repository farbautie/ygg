import config from "@/config/config";
import { ApiError } from "@/utils/apiError";
import { NextFunction, Request, Response } from "express";

export function errorHandler(err: ApiError, req: Request, res: Response, next: NextFunction) {
  let { statusCode, message } = err
  if (config.NODE_ENV === 'production') {
    statusCode = 500
    message = 'Internal Server Error'
  }

  const response = {
    code: statusCode,
    message,
    ...(config.NODE_ENV === 'development' ? { stack: err.stack } : {}),
  }

  res.status(statusCode).json(response)
}

