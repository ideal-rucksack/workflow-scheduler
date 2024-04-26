import {request} from '@umijs/max';

export const signIn = async (params: RequestPayload.SignInParams, options?: {
  [key: string]: any
}): Promise<ResponsePayload.SignInResponse> => {
  return await request<ResponsePayload.SignInResponse>('/account/signin', {
    method: 'POST',
    data: params,
    ...(options || {}),
  });
}
