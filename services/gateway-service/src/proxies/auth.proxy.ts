import { createProxyMiddleware } from "http-proxy-middleware"
import { services } from "../config/service"

export const authProxy = createProxyMiddleware({
  target: services.auth,
  changeOrigin: true,
  pathRewrite: {
    "^/auth": ""
  }
})
