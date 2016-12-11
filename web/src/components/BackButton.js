import React from 'react';
import {Link} from 'react-router';

import styles from '../styles/components/BackButton.scss';

export class BackButton extends React.Component {
  static propTypes = {
    to: React.PropTypes.string
  };

  render() {
    return <Link to={this.props.to} className={styles.backButton}>← 返回</Link>;
  }
}
