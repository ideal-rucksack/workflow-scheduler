import React from 'react';
import { Link } from 'react-router-dom';
import styled, { keyframes } from "styled-components";
import {FormattedMessage} from "@@/plugin-locale";

const float = keyframes`
  0% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-20px);
  }
  100% {
    transform: translateY(0px);
  }
`;

const NotFoundPage = () => {
  return (
    <Styles>
      <div className="not-found-page">
        <h1 className="float">404</h1>
        <p><FormattedMessage id='page.404.content'/></p>
        <Link to="/"><FormattedMessage id='page.404.back'/></Link>
      </div>
    </Styles>
  );
};

const Styles = styled.div`
.not-found-page {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100vh;
  text-align: center;
  background-color: #282c34; // 修改背景颜色为深色
  color: #fff; // 修改文字颜色为白色
}

.not-found-page h1 {
  font-size: 6em;
  margin-bottom: 0.5em;
}

.not-found-page .float {
  animation: ${float} 2s ease-in-out infinite;
}

.not-found-page p {
  font-size: 1.5em;
  margin-bottom: 1em;
}

.not-found-page a {
  color: #fff; // 修改链接颜色为白色
  text-decoration: none;
  border: 1px solid #fff; // 修改边框颜色为白色
  padding: 10px 20px;
  border-radius: 5px;
  transition: background-color 0.3s ease;

  &:hover {
    background-color: #fff; // 修改悬停背景颜色为白色
    color: #282c34; // 修改悬停文字颜色为深色
  }
}
`

export default NotFoundPage;
