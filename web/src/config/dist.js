'use strict';

import baseConfig from './base';

let config = {
  appEnv: 'dist',  // feel free to remove the appEnv property here
  apiBaseUrl: '//blog-api.cdog.me/'
};

export default Object.freeze(Object.assign({}, baseConfig, config));
