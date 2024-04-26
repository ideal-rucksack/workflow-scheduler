declare namespace RequestPayload {
  interface SignInParams {
    username: string;
    password: string;
  }
}

declare namespace ResponsePayload {
  interface SignInResponse {
    access_token: string;
    refresh_token: string;
  }

}
