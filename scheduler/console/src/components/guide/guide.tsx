import { Layout, Row, Typography } from 'antd';
import React from 'react';
import styles from './guide.less';
import {FormattedMessage} from "@@/plugin-locale";

interface Props {
  name?: string;
}

// 脚手架示例组件
const Guide: React.FC<Props> = (props) => {
  const { name } = props;
  return (
    <Layout>
      <Row>
        <Typography.Title level={3} className={styles.title}>
          您好 {name}! <br />
          感谢您使用 Workflow Scheduler
          <div>
            <FormattedMessage id='welcome' />
          </div>
        </Typography.Title>
      </Row>
    </Layout>
  );
};

export default Guide;
