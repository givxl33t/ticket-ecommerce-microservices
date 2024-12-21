import { PaginationMeta } from "./PaginationMeta";

export default interface SuccessResponse<Data> {
  message: string;
  data?: Data;
  meta?: PaginationMeta;
}
