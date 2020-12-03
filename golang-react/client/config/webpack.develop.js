const paths = require('./paths')
const base = require('./webpack.base.js')
const { merge } = require('webpack-merge')

module.exports = merge(base, {
  mode: 'development',

  devtool: 'inline-source-map',

  devServer: {
    contentBase: paths.build,
    host: 'localhost',
    port: 8000,
    publicPath: '/',
    historyApiFallback: true,
    compress: true,
    inline: true,
    hot: true,
  },
})
