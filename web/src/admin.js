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

  render() {
    return <div id={styles.app}>
      {this.props.children}
    </div>
  }
}
