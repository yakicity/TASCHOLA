export interface AuthResponse {
  "code": number,
  "expires": string,
  "message": string,
  "token": string, // JWT: JSON Web Token
}

export interface AuthRequest {
  "name": string,
  "password": string,
}
