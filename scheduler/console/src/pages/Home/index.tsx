import Guide from '@/components/guide';
import { PageContainer } from '@ant-design/pro-components';
import { useModel } from '@umijs/max';
import styles from './index.less';
import React from "react";

const HomePage: React.FC = () => {
  const { name } = useModel('global');
  const {initialState} = useModel('@@initialState');

  return (
    <PageContainer ghost>
      <div className={styles.container}>
        <Guide name={initialState?.current?.nickname} />
      </div>
      <div style={{height: 1000}}></div>
    </PageContainer>
  );
}

export default HomePage;
