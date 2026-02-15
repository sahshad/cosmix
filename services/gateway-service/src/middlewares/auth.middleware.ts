import { Request, Response, NextFunction } from "express"
import jwt from "jsonwebtoken"
import { env } from "../config/env"

export interface AuthenticatedRequest extends Request {
  user?: jwt.JwtPayload
}

export function authenticate(
  req: AuthenticatedRequest,
  res: Response,
  next: NextFunction
) {
  const header = req.headers.authorization

  if (!header) {
    return res.status(401).json({ message: "Authorization header missing" })
  }

  const token = header.split(" ")[1]

  try {
    req.user = jwt.verify(token, env.jwtPublicKey) as jwt.JwtPayload
    req.headers["X-User-Id"] = req.user.userId.toString()
    next()
  } catch {
    return res.status(401).json({ message: "Invalid token" })
  }
}
