const paths = require('./paths')
const HtmlWebpackPlugin = require('html-webpack-plugin')

module.exports = {
  entry: {
    bundle: paths.scripts + '/index.js'
  },

  output: {
    path: paths.build,
    filename: '[name].js',
    publicPath: '/'
  },

  plugins: [
    new HtmlWebpackPlugin({
      template: paths.html + '/index.html',
      filename: 'index.html',
    }),
  ],

  module: {
    rules: [{
      test: /\.(js)$/,
      exclude: /node_modules/,
      use: {
        loader: 'babel-loader',
        options: {
          presets: [
            '@babel/preset-env',
            '@babel/preset-react'
          ]
        }
      }
    }]
  },

  externals: {
    'Config': JSON.stringify({
      serverURL: "http://localhost:9000"
    })
  },
};
