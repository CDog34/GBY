import React from 'react';
import styles from './styles/App.scss';
import {NetworkIndicator, NetworkError} from './components/NetworkIndicator'

export class App extends React.Component {
  static propTypes = {
    children: React.PropTypes.node
  };

  render() {
    return <div id={styles.app}>

      <NetworkIndicator/>
      <NetworkError/>
      {this.props.children}
    </div>
  }
}
