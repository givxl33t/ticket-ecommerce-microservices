import { JWTPayload } from "@/src/model/Auth";

import { JWTContext } from "./JWTContext";

export interface UserContext extends JWTContext {
  readonly currentUser: JWTPayload;
}
