import React from 'react';
import styles from '../styles/Views/Index.scss';

export class AppComponent extends React.Component {
  render() {
    return (
      <div className={styles.index}>
        <div className={styles.title}>
          <h1>西道の狗窝</h1>
          <span>CDog's Kennel</span>
        </div>
      </div>
    );
  }
}
