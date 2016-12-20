import React from 'react';
import styles from '../styles/Views/LinkView.scss';
import {BackButton} from '../components/BackButton';
import {PageContent} from '../components/PageContent';
import {LinkService} from '../services/LinkService';
import {DaoVoiceService} from '../services/DaoVoiceService';


class LinkItem extends React.Component {
  static propTypes = {
    link: React.PropTypes.object
  };

  clickLink() {
    DaoVoiceService.outLinkVisitEvent(this.props.link);
  }

  render() {
    const link = this.props.link;
    let img = link.image;
    if (!img) img = require('../images/akalin.jpg');
    return <a className={styles.listItem} onClick={this.clickLink.bind(this)} href={link.url} target="_blank">
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

  componentDidMount() {
    DaoVoiceService.pageVisitEvent('linkView');
  }

  render() {
    return (
      <PageContent>
        <div className={styles.LinkView}>
          <div className={styles.listWrapper}>
            {this.state.links.map((link) => <LinkItem link={link} key={link.id}/>)}
          </div>
          <p className={styles.apply}>
            申请请戳 <a href='#' onClick={() => DaoVoiceService.openNewMsg(
            '你好！我想和你交换网站链接。\n' +
            '名称是：「网站名称」\n' +
            '地址是：「网站链接」\n' +
            '描述是：「一句话描述」\n' +
            '图片是：「缩略图url。(如果是个人站点，建议使用cn.gravatar.com)」')}>这里</a> ～
          </p>
          <div className={styles.back}>
            <BackButton to="/"/>
          </div>
        </div>
      </PageContent>
    );
  }
}
