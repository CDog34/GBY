import React from 'react';
import {Link} from 'react-router';
import styles from '../styles/Views/ArticleList.scss';
import {BackButton} from '../components/BackButton';
import {PageContent} from '../components/PageContent';


class ArticleItem extends React.Component {
  render() {
    return <Link to={`/article/${'testArticle'}`}>
      <article className={styles.listItem}>
        <h1>震撼世界：中一颗卫星</h1>
        <div className={styles.articleMeta}>
          <p className={styles.author}>西道</p>
          <p className={styles.time}>1995年11月23日 13:20</p>
        </div>
      </article>
    </Link>;
  }
}

export class ArticleList extends React.Component {
  componentWillMount() {
    document.title = '文章列表-西道の狗窝';
  }

  render() {
    return (
      <PageContent>
        <div className={styles.listWrapper}>
          <ArticleItem/>
          <ArticleItem/>
          <ArticleItem/>
          <ArticleItem/>
          <ArticleItem/>
          <ArticleItem/>
          <ArticleItem/>
          <ArticleItem/>
          <ArticleItem/>
          <div className={styles.back}>
            <BackButton to="/"/>
          </div>
        </div>
      </PageContent>
    );
  }
}
