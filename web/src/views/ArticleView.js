import React from 'react';
import Moment from 'react-moment';
import ReactMarkdown from 'react-markdown';
import styles from '../styles/Views/ArticleView.scss';
import {BackButton} from '../components/BackButton';
import {PageContent} from '../components/PageContent';
import {ArticleService} from '../services/ArticleService';
import {DaoVoiceService} from '../services/DaoVoiceService';

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
    DaoVoiceService.pageVisitEvent('articleView', {
      articleId: res.id,
      articleTitle: res.title
    });
  }

  render() {
    const article = this.state.article;
    return (
      <PageContent>
        <div className={styles.articleWrapper}>
          <h1 className={styles.articleTitle}>{article.title}</h1>
          <div className={styles.articleMeta}>
            <p>
              {article.author}
              {!!article.updateAt && <Moment format="作于YYYY年MM月DD日" date={article.updateAt}/>}
            </p>
          </div>
          <div className={[styles.articleContent, 'markdown-body'].join(' ')}>
            {article.content && <ReactMarkdown softBreak='br' source={article.content} className={styles.container}/>}
          </div>
          <div className={styles.back}>
            <BackButton onClick={() => this.props.router.goBack()}/>
          </div>
        </div>
      </PageContent>
    );
  }
}
