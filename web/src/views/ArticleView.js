import React from 'react';
import {browserHistory} from 'react-router'
import Moment from 'react-moment';
import ReactMarkdown from 'react-markdown';
import styles from '../styles/Views/ArticleView.scss';
import {BackButton} from '../components/BackButton';
import {PageContent} from '../components/PageContent';
import {ArticleService} from '../services/ArticleService';

require('github-markdown-css');


export class ArticleView extends React.Component {
  state = {
    article: {}
  };
  static propTypes = {
    params: React.PropTypes.object
  };

  async componentWillMount() {
    const articleId = this.props.params.articleId;
    const res = await ArticleService.get(articleId);
    document.title = `${res.title}-西道の狗窝`;
    this.setState({article: res});
  }

  render() {
    const article = this.state.article;
    return (
      <PageContent>
        <div className={styles.articleWrapper}>
          <h1 className={styles.articleTitle}>{article.title}</h1>
          <div className={styles.articleMeta}>
            <p className={styles.author}>{article.author}</p>
            {!!article.updateAt &&
            <p className={styles.time}><Moment format="YYYY年MM月DD日 HH:mm" date={article.updateAt}/></p>}
          </div>
          <div className={[styles.articleContent, 'markdown-body'].join(' ')}>
            {article.content && <ReactMarkdown source={article.content}/>}
          </div>
          <div className={styles.back}>
            <BackButton onClick={() => browserHistory.goBack()}/>
          </div>
        </div>
      </PageContent>
    );
  }
}
