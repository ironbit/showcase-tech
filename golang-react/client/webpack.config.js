const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin')

module.exports = {
  entry: {
    bundle: path.join(__dirname, 'scripts', 'index.js')
  },

  output: {
    path: path.resolve(__dirname, 'build'),
    filename: '[name].js',
    publicPath: '/'
  },

  plugins: [
    new HtmlWebpackPlugin({
      template: path.join(__dirname, 'html', 'index.html'),
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


  mode: 'development',

  devtool: 'inline-source-map',

  devServer: {
    contentBase: path.resolve(__dirname, 'build'),
    publicPath: '/',
    historyApiFallback: true,
    inline: true,
    hot: true,
  },
};
