import { Router } from "express"
import { authProxy } from "../proxies/auth.proxy"

export const authRoutes = Router()
authRoutes.use("/", authProxy)
