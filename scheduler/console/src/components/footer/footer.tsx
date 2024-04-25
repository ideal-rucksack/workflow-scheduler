import {DefaultFooter} from "@ant-design/pro-layout";
import React from "react";
import {GithubOutlined} from "@ant-design/icons";
import {Props} from "react-intl/lib/components/message";
import {FormattedMessage} from "@@/plugin-locale";

const Footer: React.FC<Props> = () => {
  return (
    <DefaultFooter
      copyright="@2024 Jixiangup"
      links={[
        {
          key: 'Ant Design Pro',
          title: <FormattedMessage id='footer.link.antdPro'/>,
          href: 'https://pro.ant.design',
          blankTarget: true,
        },
        {
          key: 'github',
          title: <GithubOutlined />,
          href: 'https://github.com/ideal-rucksack/workflow-scheduler',
          blankTarget: true,
        },
        {
          key: 'Ant Design',
          title: <FormattedMessage id='footer.link.antd'/>,
          href: 'https://ant.design',
          blankTarget: true,
        },
        {
          key: 'Umi js',
          title: <FormattedMessage id='footer.link.umi'/>,
          href: 'https://umijs.org',
          blankTarget: true,
        },
      ]}
    />
  );
}

export default Footer;
