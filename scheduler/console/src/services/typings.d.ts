declare namespace RequestPayload {
  interface SignIn {
    username: string;
    password: string;
  }

  interface RefreshToken {
    refresh_token: string;
  }
}

declare namespace ResponsePayload {
  interface SignIn {
    access_token: string;
    refresh_token: string;
  }

  interface Account {
    nickname: string;
    username: string;
  }

  interface Current {
    nickname: string;
    username: string;
    email: string;
    code?: string;
    Status?: 'active' | 'inactive' | 'suspended';
  }

  interface RefreshToken {
    access_token: string;
  }
}
