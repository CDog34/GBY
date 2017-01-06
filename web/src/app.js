import React from 'react';
import styles from './styles/App.scss';
import {NetworkIndicator, NetworkError} from './components/NetworkIndicator'
import {DaoVoiceService} from'./services/DaoVoiceService';
import {VisitorService} from'./services/VisitorService';
import config from 'config';

export class App extends React.Component {
  static propTypes = {
    children: React.PropTypes.node
  };

  componentWillMount() {
    if (config.anonymousTracking) {
      const visitor = VisitorService.getVisitor();
      DaoVoiceService.setUpForVisitor(visitor.uuid);
    } else {
      DaoVoiceService.setUpForAnonymous()
    }

  }

  render() {

    return <div id={styles.app}>

      <NetworkIndicator/>
      <NetworkError/>
      {this.props.children}
    </div>
  }
}
