import React from 'react';
import styles from '../styles/Views/Index.scss';
import {Link} from 'react-router';
import _ from 'lodash';
import {observer} from 'mobx-react';

import {AuthService} from '../services/AuthService';
import {ArticleService} from '../services/ArticleService';
import {DaoVoiceService} from '../services/DaoVoiceService';
import {AuthStore} from '../stores/AuthStore';
import pkg from '../../package.json';

class Action extends React.Component {
  static propTypes = {
    text: React.PropTypes.string,
    href: React.PropTypes.string,
    link: React.PropTypes.string
  };

  render() {
    if (this.props.link) return <Link to={this.props.link} className={styles.action}>{this.props.text}</Link>;
    return <a href={this.props.href} className={styles.action}>{this.props.text}</a>;
  }
}

@observer
export class Index extends React.Component {
  state = {
    articles: {}
  };

  async componentWillMount() {
    document.title = '首页-西道の狗窝';
    try {
      await AuthService.valid();
    } catch (err) {
      //Silence is gold
    }
    const res = await ArticleService.listForIndex();
    this.setState({articles: res});
    DaoVoiceService.pageVisitEvent('index');

  }

  render() {
    const store = AuthStore.getInstance();
    return (
      <div className={styles.index}>
        <div className={styles.title}>
          <h1>西道の狗窝</h1>
          <span>CDog's Kennel</span>
        </div>
        <div className={styles.actions}>
          <Action text="文章" link="/articleList"/>
          <Action text="链接" link="/lian"/>
          {_.map(this.state.articles, (article) => <Action text={article.title} link={`/article/${article.id}`}
                                                           key={article.id}/>)}
          {store.auth && <Action text="管理" link="/smartPuppy"/>}
        </div>
        <p className={styles.footer}>
          Power By <a href="https://github.com/CDog34/GBY" target="_blank">GBY v{pkg.version}</a>.
        </p>
      </div>
    );
  }
}
