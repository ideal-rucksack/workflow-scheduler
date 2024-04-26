import {App, ConfigProvider} from "antd";

export default (props: any) => {
  const locale = getLocale();
  return (
    <ConfigProvider locale={locale}>
      <App>
          {props?.children}
      </App>
    </ConfigProvider>
)
}

import {getLocale} from "@@/plugin-locale";
