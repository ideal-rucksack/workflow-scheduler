import {Avatar, Badge, Button, ConfigProvider, Dropdown, Space} from "antd";
import type {MenuProps} from 'antd';
import React from "react";
import styled from "styled-components";
import {FormattedMessage, getAllLocales, getLocale, SelectLang, setLocale} from "@@/plugin-locale";
import {HiLogout, HiTranslate} from "react-icons/hi";
import {useModel} from "@umijs/max";

const defaultLangUConfigMap: Map<string, any> = new Map<string, any>([
  ['en-US', {
      lang: 'en-US',
      label: 'English',
    }
  ],
  ['zh-CN', {
    lang: 'zh-CN',
    label: '简体中文',
  }]
]);

export default () => {

  const allLocales = getAllLocales();
  const currentLocal = getLocale();
  const {initialState} = useModel('@@initialState');

  const profileItems: MenuProps['items'] = [
    {
      key: 'sign_out',
      label: (
        <Styles>
          <span className='dropdown-item'>
            <HiLogout/>&nbsp;
            <FormattedMessage id='signOut'/>
          </span>
        </Styles>
      ),
    },
  ];
  const translateItems: MenuProps['items'] = allLocales.map(e => {
    if (e === currentLocal) {
      return null;
    }
    return {
      key: e,
      label: (
        <Styles>
          <div className={currentLocal === e ? 'dropdown-item dropdown-item-selected' : 'dropdown-item'} onClick={(event) => {
            event.defaultPrevented;
            setLocale(e);
          }}>
            {defaultLangUConfigMap.get(e)?.label}
          </div>
        </Styles>
      ),
    }
  })

  return (
    <ConfigProvider>
      <Styles>
        <div>
          <Space size='middle'>
            <div className='active-item translate'>
              <Dropdown menu={{items: translateItems}}>
                <HiTranslate size={24} />
              </Dropdown>
            </div>
            <div className='active-item profile'>
              <Dropdown menu={{items: profileItems}}>
                <div>
                  <Badge dot>
                    <Avatar shape="square" src={'https://avatars.githubusercontent.com/u/75556346?v=4'}/>
                  </Badge>
                  <div>{initialState?.current?.nickname}</div>
                </div>
              </Dropdown>
            </div>
          </Space>
        </div>
      </Styles>
    </ConfigProvider>
  );
}

const Styles = styled.div`
  * {
      display: flex;
      justify-content: center;
      align-items: center;
  }
  .active-item {
      :hover {
          cursor: pointer;
      }
      display: flex;
      justify-content: center;
      align-items: center;
  }
  .dropdown-item {
      display: flex;
      justify-content: center;
      align-items: center;
      padding: 0;
      margin: 0;
  }
`
