import React from 'react';
import styles from '../styles/Views/Index.scss';

class Action extends React.Component {
  static propTypes = {
    text: React.PropTypes.string,
    href: React.PropTypes.string
  };

  render() {
    return <a className={styles.action} href={this.props.href}>
      {this.props.text}

    </a>
  }
}

export class AppComponent extends React.Component {
  componentWillMount() {
    document.title = '首页-西道の狗窝';
  }

  render() {
    return (
      <div className={styles.index}>
        <div className={styles.title}>
          <h1>西道の狗窝</h1>
          <span>CDog's Kennel</span>
        </div>
        <div className={styles.actions}>
          <Action text="首页" href="http://cdog.me"/>
          <Action text="文章"/>
          <Action text="关于"/>
        </div>
        <p className={styles.footer}>
          Power By <a href="https://github.com/CDog34/GBY" target="_blank">GBY</a>
        </p>
      </div>
    );
  }
}
