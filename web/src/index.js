import 'core-js/fn/object/assign';
import React from 'react';
import ReactDOM from 'react-dom';
import {Router, Route, IndexRoute, browserHistory} from 'react-router'
import {App} from './app';
import {Admin} from './admin';
import {Index} from './views/Index';
import {About} from './views/About';
import {ArticleList} from './views/ArticleList';
import {AdminArticleList} from './views/AdminArticleList';
import {ArticleView} from './views/ArticleView';
import {Login} from './views/Login';


require('normalize.css/normalize.css');

ReactDOM.render((<Router history={browserHistory}>
    <Route path="/" component={App}>
      <IndexRoute component={Index}/>
      <Route component={About} path="about"/>
      <Route component={ArticleList} path="/articleList"/>
      <Route component={ArticleView} path="/article/:articleId"/>
    </Route>
    <Route path="/smartPuppy" components={Admin}>
      <Route component={Login} path="login"/>
      <Route component={AdminArticleList} path="aList"/>
    </Route>
  </Router>
), document.getElementById('app-content'));
