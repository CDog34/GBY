import React from 'react';
import {observer} from 'mobx-react'
import styles from './styles/App.scss';
import {AuthService}from './services/AuthService';
import {AuthStore} from './stores/AuthStore';
import {NetworkIndicator, NetworkError} from './components/NetworkIndicator';
import {Link} from 'react-router';

@observer
export class Admin extends React.Component {
  store;
  static propTypes = {
    children: React.PropTypes.node
  };

  async componentWillMount() {
    this.store = AuthStore.getInstance();
    try {
      await AuthService.valid()
    } catch (err) {
      this.props.router.push('/smartPuppy/login')
    }

  }

  async logout() {
    await AuthService.logout();
    this.props.router.push('/');
  }

  renderMenu() {
    return <div className={styles.adminMenu}>
      <div className={styles.mainBtn}>
        <Link to="/">首页</Link>
        <Link to="/smartPuppy/aList">文章列表</Link>
        <Link to="/smartPuppy/lList">链接列表</Link>
      </div>
      <button onClick={this.logout.bind(this)} className={styles.logoutBtn}>登出</button>
    </div>

  }

  render() {
    const isAuth = this.store.auth;
    return <div id={styles.app}>
      <NetworkIndicator/>
      <NetworkError/>
      {isAuth && this.renderMenu()}
      {this.props.children}
    </div>
  }
}
