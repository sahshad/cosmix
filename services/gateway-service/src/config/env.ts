import dotenv from "dotenv"

dotenv.config()

export const env = {
  port: Number(process.env.PORT) || 3000,
  authServiceUrl: process.env.AUTH_SERVICE_URL!,
  jwtPublicKey: process.env.JWT_PUBLIC_KEY!
}
