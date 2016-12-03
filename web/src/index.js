import 'core-js/fn/object/assign';
import React from 'react';
import ReactDOM from 'react-dom';
import App from './views/Main';
require('normalize.css/normalize.css');
require('./styles/App.css');

// Render the main component into the dom
ReactDOM.render(<App />, document.getElementById('app'));
