// 运行时配置

// 全局初始化数据配置，用于 Layout 用户信息和权限初始化
// 更多信息见文档：https://umijs.org/docs/api/runtime-config#getinitialstate
import {Footer} from "@/components/footer";
import {RunTimeLayoutConfig} from "@@/plugin-layout/types";
import {AntdConfig, RuntimeAntdConfig} from "@@/plugin-antd/types";
import {Actions} from "@/components/layout";
import React from "react";
import Root from "@/root";
import {RequestConfig} from "@@/plugin-request/request";
import {notification} from "antd";
import {FormattedMessage} from "@@/plugin-locale";

export async function getInitialState(): Promise<{ name: string }> {
  return { name: 'Hello World' };
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
        const { status, response } = err?.request;
        const responseBody: ErrorResponse = JSON.parse(response);
        switch (status) {
          // 鉴权失败
          case 401:
            break;
          case 400:
            // success(responseBody.error)
            notification.error({
              description: responseBody.error,
              // message: <FormattedMessage id='request.400'/>,
              message: 'Bad Request',
            });
            break;
          case 500:
            notification.error({
              description: responseBody.error,
              // message: <FormattedMessage id='request.500'/>,
              message: '服务器内部错误',
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
            // "x-admin-token": localStorage.getItem(TOKEN_HEADER)
          }
        }
      };
    }
  ],
};
