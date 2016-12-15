import React from 'react'
import {observer} from 'mobx-react';
import {NetworkStore} from '../stores/NetworkStore';
import styles from '../styles/components/NetworkIndicator.scss';

const errors = {
  10000: '资源不存在',
  20000: '字段验证失败',
  20001: '必填字段缺失',
  20002: '字段类型不符合',
  20003: '异常的资源编号',
  30001: '需要授权',
  30002: '授权失败',
  30003: '已存在的邮箱地址',
  30004: '密码不破配',
  30005: '未登录',
  30006: '未授权该项操作',
  66666: '客户端错误'
};

@observer
export class NetworkIndicator extends React.Component {
  store;

  componentWillMount() {
    this.store = NetworkStore.getInstance()
  }

  render() {
    if (!this.store.isBusy) return null;
    return <div className={styles.indicator}>
      <div className={styles['cssload-loader']}>
        <div className={[styles['cssload-inner'], styles['cssload-one']].join(' ')}></div>
        <div className={[styles['cssload-inner'], styles['cssload-two']].join(' ')}></div>
        <div className={[styles['cssload-inner'], styles['cssload-three']].join(' ')}></div>
      </div>
    </div>
  }
}

@observer
export class NetworkError extends React.Component {
  store;

  componentWillMount() {
    this.store = NetworkStore.getInstance()
  }

  getErrorText(errorCode) {
    const res = errors[errorCode];
    if (!res) return `未知错误 错误码：${errorCode}`;
    return res;
  }

  render() {
    if (!this.store.isError) return null;
    const {err}=this.store;
    return <div className={styles.error}>
      <p className={styles.title}>请求错误</p>
      <p className={styles.type}>类型：{this.getErrorText(err.errorCode)}</p>
      <p className={styles.message}>消息：{err.message}</p>

    </div>
  }
}
