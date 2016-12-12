'use strict';

import baseConfig from './base';

let config = {
  appEnv: 'staging',  // feel free to remove the appEnv property here
  apiBaseUrl: '//api.gby.isues.net/'
};

export default Object.freeze(Object.assign({}, baseConfig, config));
