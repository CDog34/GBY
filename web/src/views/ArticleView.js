import React from 'react';
import {browserHistory} from 'react-router'
import styles from '../styles/Views/ArticleView.scss';
import {BackButton} from '../components/BackButton';
import {PageContent} from '../components/PageContent';


export class ArticleView extends React.Component {
  static propTypes = {
    params: React.PropTypes.object
  };

  componentWillMount() {
    document.title = `文章${this.props.params.articleId}-西道の狗窝`;
  }

  render() {
    return (
      <PageContent>
        <div className={styles.articleWrapper}>
          <h1 className={styles.articleTitle}>美商务部终裁中国产大型洗衣机存在倾销行为</h1>
          <div className={styles.articleMeta}>
            <p className={styles.author}>凤凰社</p>
            <p className={styles.time}>2016-12-10 13:24:00</p>
          </div>
          <div className={styles.articleContent}>
            <p>新华社华盛顿12月9日电(记者郑启航 高攀)美国商务部9日宣布终裁结果，认定从中国进口的大型洗衣机存在倾销行为。
              根据当天发表的声明，中国厂商的倾销幅度为32.12%至52.51%。美国商务部将通知美国海关按最终确定的倾销幅度对上述产品的中国生产商和出口商征收保证金。
              为回应美国惠而浦公司的申诉，美国商务部于今年1月对从中国进口的上述产品发起反倾销调查。
              根据美国贸易救济政策程序，最终是否征收反倾销税还需另一家联邦机构美国国际贸易委员会作出裁决。按照最新日程，国际贸易委员会将于2017年1月23日左右作出终裁。
              根据美国商务部的数据，2015年美国从中国进口的这类产品金额约为11亿美元。
              对于中美之间的贸易摩擦，中国商务部多次表示，希望美国政府恪守反对贸易保护主义承诺，共同维护自由、开放、公正的国际贸易环境，以更加理性的方法妥善处理贸易摩擦。
            </p>
          </div>
          <div className={styles.back}>
            <BackButton onClick={() => browserHistory.goBack()}/>
          </div>
        </div>
      </PageContent>
    );
  }
}
