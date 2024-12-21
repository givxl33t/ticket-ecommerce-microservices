export interface JWTPayload {
  readonly id: string;
  readonly email: string;
  readonly exp: number;
}
