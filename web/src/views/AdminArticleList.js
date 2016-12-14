import React from 'react';
import {Link} from 'react-router';
import Moment from 'react-moment';
import styles from '../styles/Views/ArticleList.scss';
import {BackButton} from '../components/BackButton';
import {PageContent} from '../components/PageContent';
import {ArticleService} from '../services/ArticleService';


class ArticleItem extends React.Component {
  static propTypes = {
    article: React.PropTypes.object
  };

  render() {
    const article = this.props.article;
    return <Link to={`/article/${article.id}`}>
      <article className={styles.listItem}>
        <h1>{article.title}</h1>
        <div className={styles.articleMeta}>
          <p className={styles.author}>{article.author}</p>
          {!!article.updateAt &&
          <p className={styles.time}><Moment format="YYYY年MM月DD日 HH:mm" date={article.updateAt}/></p>}
        </div>
      </article>
    </Link>;
  }
}

export class AdminArticleList extends React.Component {
  state = {
    articles: []
  };

  async componentWillMount() {
    document.title = '文章列表-西道の狗窝';
    const res = await ArticleService.list();
    this.setState({articles: res});
  }

  render() {
    return (
      <PageContent>
        <div className={styles.listWrapper}>
          {this.state.articles.map((article) => <ArticleItem article={article} key={article.id}/>)}
          <div className={styles.back}>
            <BackButton to="/"/>
          </div>
        </div>
      </PageContent>
    );
  }
}
