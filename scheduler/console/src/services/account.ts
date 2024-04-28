import {request} from '@umijs/max';

export const signIn = async (params: RequestPayload.SignIn, options?: {
  [key: string]: any
}): Promise<ResponsePayload.SignIn> => {
  return await request<ResponsePayload.SignIn>('/account/signin', {
    method: 'POST',
    data: params,
    ...(options || {}),
  });
}

export const current = async (options?: {
  [key: string]: any
}): Promise<ResponsePayload.Current> => {
  return await request<ResponsePayload.Current>('/account/current', {
    method: 'GET',
    ...(options || {}),
  });
}

export const refreshTokens = async (params: RequestPayload.RefreshToken, options?: {
  [key: string]: any
}): Promise<ResponsePayload.RefreshToken> => {
  return await request<ResponsePayload.RefreshToken>('/account/refresh_token', {
    method: 'POST',
    data: params,
    ...(options || {}),
  });
}
