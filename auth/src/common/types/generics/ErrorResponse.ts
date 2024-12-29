export default interface ErrorResponse<Code = string> {
  message: string | object;
  code: Code;
}
