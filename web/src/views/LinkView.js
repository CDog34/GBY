import React from 'react';
import styles from '../styles/Views/LinkView.scss';
import {BackButton} from '../components/BackButton';
import {PageContent} from '../components/PageContent';
import {LinkService} from '../services/LinkService';


class LinkItem extends React.Component {
  static propTypes = {
    link: React.PropTypes.object
  };

  render() {
    const link = this.props.link;
    let img = link.image;
    if (!img) img = require('../images/akalin.jpg');
    return <a className={styles.listItem} href={link.url} target="_blank">
      <div className={styles.imageWrapper}>
        <img src={img} alt={link.name}/>
      </div>
      <div className={styles.intro}>
        <p className={styles.name}>#{link.name}</p>
        <p className={styles.desc}>{link.description}</p>
      </div>
    </a>;
  }
}

export class LinkView extends React.Component {
  state = {
    links: []
  };

  async componentWillMount() {
    document.title = '链接列表-西道の狗窝';
    const res = await LinkService.list();
    this.setState({links: res});
  }

  render() {
    return (
      <PageContent>
        <div className={styles.LinkView}>
          <div className={styles.listWrapper}>
            {this.state.links.map((link) => <LinkItem link={link} key={link.id}/>)}
          </div>
          <p className={styles.apply}>
            申请请戳右下角～
          </p>
          <div className={styles.back}>
            <BackButton to="/"/>
          </div>
        </div>
      </PageContent>
    );
  }
}
