import {defineConfig} from "@umijs/max";
import layout from "./layout";
import routes from "./routes";
import locale from "./locale";
import proxy from "./proxy";

export default defineConfig({
  antd: {
    configProvider: {},
    appConfig: {},
  },
  proxy: {
    ...proxy
  },
  styledComponents: {},
  access: {},
  model: {},
  initialState: {},
  request: {},
  layout: {
    ...layout
  },
  routes: [
    ...routes
  ],
  locale: {
    ...locale
  },
  npmClient: 'yarn',
})
