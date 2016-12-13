import React from 'react';
import styles from './styles/App.scss';
import {LocalStorageService}from './services/LocalStorageService';

export class Admin extends React.Component {
  static propTypes = {
    children: React.PropTypes.node
  };

  componentWillMount() {
    LocalStorageService.set('Time', new Date())
  }

  render() {
    return <div id={styles.app}>
      {this.props.children}
    </div>
  }
}
