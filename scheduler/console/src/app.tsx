// 运行时配置

// 全局初始化数据配置，用于 Layout 用户信息和权限初始化
// 更多信息见文档：https://umijs.org/docs/api/runtime-config#getinitialstate
import {Footer} from "@/components/footer";
import {RunTimeLayoutConfig} from "@@/plugin-layout/types";
import {AntdConfig, RuntimeAntdConfig} from "@@/plugin-antd/types";
import {Actions} from "@/components/layout";

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
