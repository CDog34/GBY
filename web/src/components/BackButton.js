import React from 'react';
import {Link} from 'react-router';

import styles from '../styles/components/BackButton.scss';

export class BackButton extends React.Component {
  static propTypes = {
    to: React.PropTypes.string,
    onClick: React.PropTypes.func
  };

  render() {
    if (this.props.onClick) return (
      <a href="#" className={styles.backButton} onClick={this.props.onClick.bind(this)}>←返回</a>
    );
    return <Link to={this.props.to} className={styles.backButton}>← 返回</Link>;
  }
}
