import {Button} from "antd";
import {ProForm} from "@ant-design/pro-form";
import {ProFormText} from "@ant-design/pro-components";
import {FormattedMessage, getLocale, setLocale, useIntl} from "@@/plugin-locale";
import {Styles} from "@/components/account/styles";
import {signIn} from "@/services/account";

export default () => {

  const locale = getLocale();
  const intl = useIntl();
  const usernamePlaceholder = intl.formatMessage(
    {
      id: 'page.login.username.placeholder',
    },
  );

  const passwordPlaceholder = intl.formatMessage(
    {
      id: 'page.login.password.placeholder',
    },
  );

  return (
    <Styles>
      <div>
        <div className="sign-in-container">
          <div className='title-container'>
            <h2 className='title'><FormattedMessage id='page.login.title'/></h2>
          </div>
          <div className="sign-in-form">
            <ProForm<RequestPayload.SignInParams>
              submitter={false}
              onFinish={async (params) => {
                const {refresh_token, access_token} = await signIn(params);
                console.log('refresh_token', refresh_token, 'access_token', access_token);
              }}
            >
              <ProFormText
                name='username'
                label={<span><FormattedMessage id='page.login.username.label'/>:</span>}
                placeholder={usernamePlaceholder}
              />
              <ProFormText.Password
                name='password'
                label={<span><FormattedMessage id='page.login.password.label'/>:</span>}
                placeholder={passwordPlaceholder}
              />
              <ProForm.Item className='login-buttons'>
                <Button type='primary' htmlType='submit' className='login-button'>
                  <FormattedMessage id='button.signIn'/>
                </Button>
              </ProForm.Item>
            </ProForm>
          </div>
        </div>
        <div className="sign-in-container">
          <div className="sign-in-bg">
            登陆背景
          </div>
        </div>
      </div>
      <Button type='primary' onClick={() => {
        if (locale === 'zh-CN') {
          setLocale('en-US');
        } else {
          setLocale('zh-CN');
        }
      }}>切换语言</Button>
    </Styles>
  )
}

