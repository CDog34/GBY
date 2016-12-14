import React from 'react';
import styles from '../styles/Views/Index.scss';
import {AuthService} from '../services/AuthService';


export class Login extends React.Component {
  state = {
    name: '',
    pwd: ''
  };

  async componentWillMount() {
    document.title = '登录-西道の狗窝';
    try {
      await AuthService.valid();
      this.props.router.push('/smartPuppy');
    } catch (err) {
      //Silence is gold
    }

  }

  async login() {
    try {
      const userName = this.state.name;
      const password = this.state.pwd;
      await AuthService.login(userName, password);
      this.props.router.push('/smartPuppy/aList');
    } catch (err) {
      alert(`${err.errorCode}/${err.message}`);
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
