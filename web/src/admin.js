import React from 'react';
import styles from './styles/App.scss';
import {AuthService}from './services/AuthService';

export class Admin extends React.Component {
  static propTypes = {
    children: React.PropTypes.node
  };

  async componentWillMount() {
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

  render() {
    const isAuth = AuthService.isAuth();
    return <div id={styles.app}>
      {isAuth && <button onClick={this.logout.bind(this)} className={styles.logoutBtn}>登出</button>}
      {this.props.children}
    </div>
  }
}
