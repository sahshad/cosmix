import { createProxyMiddleware } from "http-proxy-middleware";
import { services } from "../config/service";

export const userProxy = createProxyMiddleware({
    target: services.user,
    changeOrigin: true,
    pathRewrite: {
        "^/user": ""
    }
})