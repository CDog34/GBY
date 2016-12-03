import 'core-js/fn/object/assign';
import React from 'react';
import ReactDOM from 'react-dom';
import {AppComponent} from './views/Index';
require('normalize.css/normalize.css');
require('./styles/App.css');

// Render the main component into the dom
ReactDOM.render(<AppComponent />, document.getElementById('app'));
