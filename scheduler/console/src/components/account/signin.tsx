import {Button, Flex, Space} from "antd";
import {ProForm} from "@ant-design/pro-form";
import {ProFormText} from "@ant-design/pro-components";
import {FormattedMessage, useIntl} from "@@/plugin-locale";
import {Styles} from "@/components/account/styles";
import {signIn} from "@/services/account";
import CryptoJS from 'crypto-js';
import {TOKEN} from "@/constants";

export default () => {

  const intl = useIntl();
  const usernamePlaceholder = intl.formatMessage(
    {
      id: 'page.signin.input.username.placeholder',
    },
  );

  const passwordPlaceholder = intl.formatMessage(
    {
      id: 'page.signin.input.password.placeholder',
    },
  );

  const handleSignin = async (params: RequestPayload.SignIn) => {
    params.password = CryptoJS.MD5(params.password).toString();
    const { refresh_token, access_token} = await signIn(params);
    localStorage.setItem(TOKEN.ACCESS_TOKEN, access_token);
    localStorage.setItem(TOKEN.REFRESH_TOKEN, refresh_token);
    window.location.reload();
  }

  return (
    <Styles>
      <div className='content-container'>
        <div className="signin-container signin-wrapper">
          <div className='title-container'>
            <h2 className='title'><FormattedMessage id='page.signin.title'/></h2>
          </div>
          <div className="signin-form-container">
            <ProForm<RequestPayload.SignIn>
              className='signin-form'
              submitter={false}
              onFinish={handleSignin}
            >
              <Flex vertical justify='space-between' gap='large'>
                <div>
                  <ProFormText
                    style={{borderRadius: 15}}
                    required
                    name='username'
                    rules={[
                      {
                        required: true,
                        message: <FormattedMessage id='page.signin.input.username.required'/>,
                      },
                    ]}
                    label={<span><FormattedMessage id='page.signin.input.username.label'/></span>}
                    placeholder={usernamePlaceholder}
                  />
                </div>
                <div>
                  <ProFormText.Password
                    required
                    rules={[
                      {
                        required: true,
                        message: <FormattedMessage id='page.signin.input.password.required'/>,
                      },
                    ]}
                    name='password'
                    label={<span><FormattedMessage id='page.signin.input.password.label'/></span>}
                    placeholder={passwordPlaceholder}
                  />
                </div>
                <ProForm.Item className='signin-buttons'>
                  <Button style={{borderRadius: 15}} type='primary' htmlType='submit' className='signin-button'>
                    <FormattedMessage id='button.signIn'/>
                  </Button>
                </ProForm.Item>
              </Flex>
            </ProForm>
          </div>
          <div className="to-signup">
            <Space>
              <FormattedMessage id='page.signin.toSignupText'/>
              <a href='/signup'>
                <FormattedMessage id='page.signin.toSignupLink'/>
              </a>
            </Space>
          </div>
        </div>
        <div className="signin-container right-wrapper">
          <div className="signin-right">
          </div>
        </div>
      </div>
      <div className="footer-container">
      </div>
    </Styles>
  )
}

