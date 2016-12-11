import React from 'react';
import styles from '../styles/Views/Index.scss';
import {BackButton} from '../components/BackButton';


export class About extends React.Component {
  componentWillMount() {
    document.title = '关于-西道の狗窝';
  }

  render() {
    return (
      <div className={styles.index}>

        <p>～ 我是一只很萌很萌的大狗狗 ～</p>
        <BackButton to="/"/>
      </div>
    );
  }
}
