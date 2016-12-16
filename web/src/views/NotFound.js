import React from 'react';
import styles from '../styles/Views/Index.scss';

export class NotFound extends React.Component {
  componentWillMount() {
    document.title = '页面不存在';

  }

  render() {
    return (
      <div className={styles.index}>
        <img src={require('../images/Cirno.jpg')} alt="⑨" style={{width: '256px'}}/>
        <p style={{fontSize: '28px', fontWeight: '500', color: '#8ab6f7', letterSpacing: '8px'}}>页面不存在</p>
      </div>
    );
  }
}
