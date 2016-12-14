import React from 'react';
import Moment from 'react-moment';
import ReactMarkdown from 'react-markdown';
import styles from '../../styles/Views/admin/ArticleView.scss';
import {BackButton} from '../../components/BackButton';
import {PageContent} from '../../components/PageContent';
import {ArticleService} from '../../services/ArticleService';

require('github-markdown-css');


export class AdminArticleCreate extends React.Component {
  state = {
    updateAt: {},
    title: '',
    author: '',
    content: ''
  };
  static propTypes = {
    params: React.PropTypes.object
  };

  async load(articleId) {
    const res = await ArticleService.get(articleId);
    document.title = `${res.title}-西道の狗窝`;
    this.article = res;
    this.setState({title: res.title, author: res.author, content: res.content, updateAt: res.updateAt});

  }

  async componentWillMount() {
    const articleId = this.props.params.articleId;
    if (articleId !== 'new') {
      return await this.load(articleId)
    }
    document.title = '新建文章-西道の狗窝';
    this.article = {};
  }

  changeContent(e) {
    this.setState({content: e.target.value});
  }

  changeTitle(e) {
    this.setState({title: e.target.value});
  }

  changeAuthor(e) {
    this.setState({author: e.target.value});
  }

  async save() {
    try {
      await ArticleService.save(Object.assign(this.article, {
        title: this.state.title,
        author: this.state.author,
        content: this.state.content
      }));
      this.props.router.goBack();
    } catch (err) {
      alert(`${err.errorCode}/${err.message}`)
    }


  }

  render() {
    const {title, author, updateAt, content} = this.state;
    return (
      <PageContent>
        <div className={styles.articleWrapper}>
          <input type="text" value={title} onChange={this.changeTitle.bind(this)} placeholder="标题"
                 className={styles.articleTitle}/>
          <div className={styles.articleMeta}>
            <input type="text" className={styles.author} value={author} onChange={this.changeAuthor.bind(this)}
                   placeholder="作者"/>
            <p className={styles.time}><Moment format="YYYY年MM月DD日 HH:mm" date={updateAt}/></p>
            <button onClick={this.save.bind(this)}>保存</button>
          </div>
          <div className={styles.articleContent}>
            <div className={styles.rawContent}>
              <textarea value={content} onChange={this.changeContent.bind(this)} placeholder="在此输入内容"/>
            </div>
            <div className={[styles.articleContentRendered, 'markdown-body'].join(' ')}>
              <ReactMarkdown source={content}/>
            </div>
          </div>

          <div className={styles.back}>
            <BackButton onClick={() => this.props.router.goBack()}/>
          </div>
        </div>
      </PageContent>
    );
  }
}
