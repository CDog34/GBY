import React from 'react';
import styles from './styles/App.scss';
import {NetworkIndicator, NetworkError} from './components/NetworkIndicator'
import {DaoVoiceService} from'./services/DaoVoiceService';

export class App extends React.Component {
  static propTypes = {
    children: React.PropTypes.node
  };

  componentWillMount() {
    DaoVoiceService.setUpForAnonymous();
  }

  render() {

    return <div id={styles.app}>

      <NetworkIndicator/>
      <NetworkError/>
      {this.props.children}
    </div>
  }
}
