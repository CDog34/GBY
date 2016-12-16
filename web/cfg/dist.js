'use strict';

let path = require('path');
let webpack = require('webpack');

let baseConfig = require('./base');
let defaultSettings = require('./defaults');
let HtmlWebpackPlugin = require('html-webpack-plugin');

// Add needed plugins here
let BowerWebpackPlugin = require('bower-webpack-plugin');

let config = Object.assign({}, baseConfig, {
  entry: {
    bundle: [
      'babel-polyfill',
      path.join(__dirname, '../src/index')
    ],
    vendor: ['whatwg-fetch', 'es6-promise', 'babel-polyfill', 'react', 'react-moment', 'react-router', 'lodash', './package.json', 'react-markdown', 'mobx', 'mobx-react'],
  },
  cache: false,
  devtool: null,
  plugins: [
    new webpack.optimize.DedupePlugin(),
    new webpack.ProvidePlugin({
      'Promise': 'es6-promise',
      'fetch': 'imports?this=>global!exports?global.fetch!whatwg-fetch'
    }),
    new webpack.DefinePlugin({
      'process.env.NODE_ENV': '"production"'
    }),
    new BowerWebpackPlugin({
      searchResolveModulesDirectories: false
    }),
    new webpack.optimize.UglifyJsPlugin(),
    new webpack.optimize.OccurenceOrderPlugin(),
    new webpack.optimize.AggressiveMergingPlugin(),
    new webpack.NoErrorsPlugin(),
    new webpack.optimize.CommonsChunkPlugin('vendor', 'vendor.js'),
    new HtmlWebpackPlugin({
      template: path.join(__dirname, '/../src/index.html'),
      filename: path.join(__dirname, '/../dist/index.html'),
      inject: false,
      minify: {
        removeComments: true,
        collapseWhitespace: true,
        minifyCSS: true
      }
    })
  ],
  module: defaultSettings.getDefaultModules()
});

// Add needed loaders to the defaults here
config.module.loaders.push({
  test: /\.(js|jsx)$/,
  loader: 'babel',
  include: [].concat(
    config.additionalPaths,
    [path.join(__dirname, '/../src')]
  )
});

module.exports = config;
