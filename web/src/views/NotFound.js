import React from 'react';
import styles from '../styles/Views/Index.scss';
import {BackButton} from '../components/BackButton';

export class NotFound extends React.Component {
  componentWillMount() {
    document.title = '页面不存在';

  }

  render() {
    return (
      <div className={styles.index}>
        <img src={require('../images/Cirno.jpg')} alt="⑨" style={{width: '256px'}}/>
        <p style={{fontSize: '28px', fontWeight: '500', color: '#8ab6f7', letterSpacing: '8px'}}>页面不存在</p>
        <p style={{textAlign: 'center', color: 'rgba(0,0,0,0.54'}}>#能遇见就是缘#</p>
        <BackButton to="/">首页</BackButton>
      </div>
    );
  }
}
