import {Button, ConfigProvider} from "antd";

export default () => {
  return (
    <ConfigProvider>
      <div>
        <h1>Page login</h1>
        <Button type='primary'>你好</Button>
      </div>
    </ConfigProvider>
  )
}
