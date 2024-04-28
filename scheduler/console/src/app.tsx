import {Footer} from "@/components/footer";
import {RunTimeLayoutConfig} from "@@/plugin-layout/types";
import {AntdConfig, RuntimeAntdConfig} from "@@/plugin-antd/types";
import {Actions} from "@/components/layout";
import React from "react";
import Root from "@/root";
import {RequestConfig} from "@@/plugin-request/request";
import {getIntl} from "@@/plugin-locale";
import {notification} from "antd";
import {current, refreshTokens} from "@/services/account";
import {TOKEN} from "@/constants";
import {history} from "@umijs/max";

const signinPath = '/signin';

export async function getInitialState(): Promise<{
  fetchCurrent: () => Promise<ResponsePayload.Current>,
  current?: ResponsePayload.Current,
  loading: boolean,
}> {
  const getCurrent = async (): Promise<ResponsePayload.Current> => {
    return await current();
  }

  if (history.location.pathname !== signinPath) {
    const currentAccount = await getCurrent();
    return {
      fetchCurrent: getCurrent,
      current: currentAccount,
      loading: false,
    }
  } else {
    history.push('/');
  }
  return {
    fetchCurrent: current,
    loading: false,
  };
}

export const layout: RunTimeLayoutConfig = () => {
  return {
    logo: 'https://img.alicdn.com/tfs/TB1YHEpwUT1gK0jSZFhXXaAtVXa-28-27.svg',
    menu: {
      locale: true,
    },
    footerRender: () => <Footer />,
    actionsRender: () => <Actions />,
    links: [
      <a href="https://github.com">大家好</a>,
      <a href="https://github.com">大家好</a>,
    ],
    layout: 'mix',
  };
};

export const antd: RuntimeAntdConfig = (memo: AntdConfig) => {
  memo.theme ??= {
    token: {
      fontFamily: 'Arial',
      colorPrimary: '#000000',
      colorBgBase: '#f5f5f5',
      colorPrimaryHover: '#333333',
      colorBgContainer: '#f5f5f5',
    },
  };
  // memo.theme.algorithm = theme.compactAlgorithm; // 配置 antd5 的预设 dark 算法
  memo.appConfig = {
    message: {
      // 配置 message 最大显示数，超过限制时，最早的消息会被自动关闭
      maxCount: 3,
    }
  }
  return memo;
};

export function rootContainer(container: any, args: any) {
  return React.createElement(Root, null, container);
}

interface ErrorResponse {
  error: string;
}

export const request: RequestConfig = {
  timeout: 1000,
  // other axios options you want
  errorConfig: {
    errorHandler(err: any) {
      if (err?.request) {
        const intl = getIntl()
        const message400 = intl.formatMessage({ id: 'request.400' })
        const message401 = intl.formatMessage({ id: 'request.401' })
        const message500 = intl.formatMessage({ id: 'request.500' })
        const { status, response } = err?.request;
        const responseBody: ErrorResponse = JSON.parse(response);
        switch (status) {
          // 鉴权失败
          case 401:
            localStorage.removeItem(TOKEN.ACCESS_TOKEN);
            // 判断有没有refresh token 如果有去通过refresh token重新获取access token
            const refreshToken = localStorage.getItem(TOKEN.REFRESH_TOKEN);
            if (refreshToken) {
              refreshTokens({refresh_token: refreshToken})
                .then((res) => {
                  localStorage.setItem(TOKEN.ACCESS_TOKEN, res.access_token);
                  window.location.reload();
                }).catch((reason) => {
                  // reason.response.data.error
                  localStorage.removeItem(TOKEN.REFRESH_TOKEN);
                  notification.error({
                    description: reason.response.data.error,
                    message: message401,
                  });
                  history.push('/signin');
                });
            } else {
              notification.error({
                description: responseBody.error,
                message: message401,
              });
              history.push('/signin');
            }
            break;
          case 400:
            notification.error({
              description: responseBody.error,
              message: message400,
            });
            break;
          case 500:
            notification.error({
              description: responseBody.error,
              message: message500,
            });
            break;
        }
      }
    },
    errorThrower() {
    }
  },
  requestInterceptors: [
    (url: string, options: any) => {
      return {
        url: "/api" + url, // 此处可以添加域名前缀
        options: {
          ...options,
          timeout: 10000,
          headers: {
            'Authorization': 'Bearer ' + localStorage.getItem(TOKEN.ACCESS_TOKEN),
          }
        }
      };
    }
  ],
};
