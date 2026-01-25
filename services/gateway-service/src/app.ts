import express from "express"
import helmet from "helmet"
import cors from "cors"

import { rateLimiter } from "./middlewares/rateLimit.middleware"
import { errorHandler } from "./middlewares/error.middleware"

import { authRoutes } from "./routes/auth.routes"
import { healthRoutes } from "./routes/health.routes"
import { userRoutes } from "./routes/user.routes"

export const app = express()

app.use(helmet())
app.use(cors())
app.use(rateLimiter)

app.use("/api/health", healthRoutes)
app.use("/api/auth", authRoutes)
app.use("/api/user", userRoutes)

app.use(errorHandler)
