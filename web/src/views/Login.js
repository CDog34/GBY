import React from 'react';
import styles from '../styles/Views/Index.scss';
import {AuthService} from '../services/AuthService';


export class Login extends React.Component {
  state = {
    name: '',
    pwd: ''
  };

  componentWillMount() {
    document.title = '登录-西道の狗窝';
  }

  async login() {
    try {
      const userName = this.state.name;
      const password = this.state.pwd;
      await AuthService.login(userName, password);
    } catch (err) {
      alert(err);
    }
  }

  render() {
    return (
      <div className={styles.index}>

        <input type="text" placeholder="登录名" value={this.state.name}
               onChange={(t) => this.setState({name: t.target.value})}/>
        <input type="password" placeholder="密码" value={this.state.pwd}
               onChange={(t) => this.setState({pwd: t.target.value})}/>
        <button onClick={this.login.bind(this)}>GO</button>
      </div>
    );
  }
}
