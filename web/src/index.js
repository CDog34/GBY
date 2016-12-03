import 'core-js/fn/object/assign';
import React from 'react';
import ReactDOM from 'react-dom';
import {App} from './app';
import {Index} from './views/Index';
require('normalize.css/normalize.css');
import {Router, Route, IndexRoute, hashHistory} from 'react-router'

// Render the main component into the dom
ReactDOM.render((<Router history={hashHistory}>
    <Route path="/" component={App}>
      <IndexRoute component={Index}/>
    </Route>
  </Router>
), document.body);
