import React from 'react';
import styles from '../styles/Views/Index.scss';
import {Link} from 'react-router';

export class NotFound extends React.Component {
  componentWillMount() {
    document.title = '页面不存在';

  }

  render() {
    return (
      <div className={styles.index}>
        <img src={require('../images/Cirno.jpg')} alt="⑨" style={{width: '256px'}}/>
        <p style={{fontSize: '28px', fontWeight: '500', color: '#8ab6f7', letterSpacing: '8px'}}>页面不存在</p>
        <p style={{textAlign: 'center', color: 'rgba(0,0,0,0.54'}}><span
          style={{fontSize: '18px', color: 'rgba(0,0,0,0.87'}}>#能遇见就是缘#</span> 既然来了，不如 <Link to="/">去首页</Link>看看吧</p>
      </div>
    );
  }
}
