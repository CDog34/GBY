'use strict';

import baseConfig from './base';

let config = {
  appEnv: 'dist',  // feel free to remove the appEnv property here
  apiBaseUrl: '//blog-api.cdog.me/',
  enableDaoVoice: true
};

export default Object.freeze(Object.assign({}, baseConfig, config));
