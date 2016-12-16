import 'core-js/fn/object/assign';
import React from 'react';
import ReactDOM from 'react-dom';
import {Router, Route, IndexRoute, browserHistory, Redirect} from 'react-router'
import {App} from './app';
import {Admin} from './admin';
import {Index} from './views/Index';
import {NotFound} from './views/NotFound';
import {About} from './views/About';
import {ArticleList} from './views/ArticleList';
import {AdminArticleList} from './views/admin/AdminArticleList';
import {AdminArticleCreate} from './views/admin/AdminArticleCreate';
import {ArticleView} from './views/ArticleView';
import {Login} from './views/Login';
import {DaoVoiceService} from './services/DaoVoiceService';


require('normalize.css/normalize.css');
DaoVoiceService.load();

ReactDOM.render((<Router history={browserHistory}>
    <Route path="/" component={App}>
      <IndexRoute component={Index}/>
      <Route component={About} path="about"/>
      <Route component={ArticleList} path="/articleList"/>
      <Route component={ArticleView} path="/article/:articleId"/>
    </Route>
    <Route path="/smartPuppy" components={Admin}>
      <IndexRoute component={AdminArticleList}/>
      <Route component={Login} path="login"/>
      <Route component={AdminArticleList} path="aList"/>
      <Route component={AdminArticleCreate} path="a/:articleId"/>
    </Route>
    <Route path="/%E2%91%A8" component={NotFound}/>
    <Redirect path="*" to="/%E2%91%A8"/>
  </Router>
), document.getElementById('app-content'));
