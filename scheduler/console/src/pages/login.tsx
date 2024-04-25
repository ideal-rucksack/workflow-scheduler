import {Login} from "@/components/account";
import {ConfigProvider} from "antd";

export default () => {
  return (
    <ConfigProvider>
      <div>
        <Login />
      </div>
    </ConfigProvider>
  )
}
