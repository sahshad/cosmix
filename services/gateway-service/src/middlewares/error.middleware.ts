import { Request, Response, NextFunction } from "express"

export interface ApiError extends Error {
  statusCode?: number
  details?: unknown
}

/**
 * Centralized error handler
 * - Ensures consistent response shape
 * - Prevents leaking internal errors
 * - Mirrors Go-style explicit error handling
 */
export function errorHandler(
  err: ApiError,
  _req: Request,
  res: Response,
  _next: NextFunction
) {
  const statusCode = err.statusCode ?? 500

  // Log full error internally (can be replaced with Winston later)
  console.error("API Gateway Error:", {
    message: err.message,
    statusCode,
    details: err.details,
    stack: err.stack
  })

  res.status(statusCode).json({
    success: false,
    message:
      statusCode >= 500
        ? "Internal server error"
        : err.message,
    ...(err.details, { details: err.details })
  })
}
