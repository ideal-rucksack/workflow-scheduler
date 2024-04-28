import {history, Outlet, useModel} from "@umijs/max";
import React from "react";
import {Navigate} from "react-router";

export default () => {
  const {initialState} = useModel('@@initialState');
  console.log(history.location.pathname + ': initialState', initialState);
  if (initialState?.current) {
    return <Outlet />;
  } else {
    return <Navigate to='/signin' />;
  }
}
