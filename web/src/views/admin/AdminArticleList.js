import React from 'react';
import {Link} from 'react-router';
import Moment from 'react-moment';
import styles from '../../styles/Views/admin/ArticleList.scss';
import {PageContent} from '../../components/PageContent';
import {ArticleService} from '../../services/ArticleService';
import {BackButton} from '../../components/BackButton';


class ArticleItem extends React.Component {
  static propTypes = {
    article: React.PropTypes.object,
    onDelete: React.PropTypes.func,
    onRecover: React.PropTypes.func
  };


  render() {
    const article = this.props.article;
    return <Link to={`/smartPuppy/a/${article.id}`}>
      <article className={styles.listItem}>
        <h1 className={article.deleted ? styles.red : ''}>{article.title}</h1>
        <div className={styles.articleMeta}>
          <p className={styles.author}>{article.author}</p>
          {!!article.updateAt &&
          <p className={styles.time}><Moment format="YYYY年MM月DD日 HH:mm" date={article.updateAt}/></p>}
        </div>
        {!article.deleted && <button onClick={this.props.onDelete} className={styles.del}>删除</button>}
        {article.deleted && <button onClick={this.props.onRecover} className={styles.rec}>恢复</button>}
      </article>
    </Link>;
  }
}

export class AdminArticleList extends React.Component {
  state = {
    articles: []
  };

  async load() {
    const res = await ArticleService.adminList();
    this.setState({articles: res});
  }

  async componentWillMount() {
    document.title = '文章列表-西道の狗窝';
    await this.load();
  }

  deleteArticle(articleId) {
    return async(e) => {
      e.preventDefault();
      await ArticleService.adminDelete(articleId);
      await this.load()
    }
  }

  recoverArticle(articleId) {
    return async(e) => {
      e.preventDefault();
      await ArticleService.adminRecover(articleId);
      await this.load()
    }
  }

  render() {
    return (
      <PageContent>
        <div className={styles.listWrapper}>

          <div className={styles.back}>
            <BackButton to="/"/>
          </div>
          <Link to="/smartPuppy/a/new" className={styles.createNew}>+ 新建</Link>
          {this.state.articles.map((article) => <ArticleItem onDelete={this.deleteArticle(article.id)}
                                                             onRecover={this.recoverArticle(article.id)}
                                                             article={article} key={article.id}/>)}

        </div>
      </PageContent>
    );
  }
}
