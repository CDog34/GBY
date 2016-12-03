import React from 'react';
import styles from '../styles/Views/Index.scss';


export class ArticleList extends React.Component {
  componentWillMount() {
    document.title = '文章列表-西道の狗窝';
  }

  render() {
    return (
      <div className={styles.index}>
        <p>～ 啊...好想吃火锅啊 ～</p>
      </div>
    );
  }
}
