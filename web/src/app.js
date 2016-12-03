import React from 'react';
import styles from './styles/App.scss';

export class App extends React.Component {
  static propTypes = {
    children: React.PropTypes.node
  };

  render() {
    return <div id={styles.app}>
      {this.props.children}
    </div>
  }
}
