const ExtractTextPlugin = require('extract-text-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const path = require('path');
const webpack = require('webpack');

module.exports = {
  entry: {
    app: './app/app.tsx',
    vendor: ['react', 'react-dom'],
  },
  output: {
    path: path.resolve(__dirname, '../dist'),
    filename: './static/[name].[hash].js',
  },
  resolve: {
    extensions: ['.ts', '.tsx', '.js'],
  },
  module: {
    rules: [
      {
        test: /\.(js|jsx)$/,
        loaders: ['babel-loader'],
      },
      {
        test: /\.ts(x)?$/,
        loaders: ['babel-loader', 'ts-loader'],
      },
      {
        test: /\.scss$/,
        loader: ExtractTextPlugin.extract({
          fallbackLoader: 'style-loader',
          loader: 'css-loader!postcss-loader!sass-loader',
        }),
      },
      {
        test: /\.css$/,
        loader: ExtractTextPlugin.extract({
          fallbackLoader: 'style-loader',
          loader: 'css-loader',
        }),
      },
      {
        test: /\.svg$/,
        loader:
          'url-loader?limit=65000&mimetype=image/svg+xml&name=static/[name].[ext]&publicPath=../',
      },
      {
        test: /\.woff$/,
        loader:
          'url-loader?limit=65000&mimetype=application/font-woff&name=static/[name].[ext]&publicPath=../',
      },
      {
        test: /\.woff2$/,
        loader:
          'url-loader?limit=65000&mimetype=application/font-woff2&name=static/[name].[ext]&publicPath=../',
      },
      {
        test: /\.[ot]tf$/,
        loader:
          'url-loader?limit=65000&mimetype=application/octet-stream&name=static/[name].[ext]&publicPath=../',
      },
      {
        test: /\.eot$/,
        loader:
          'url-loader?limit=65000&mimetype=application/vnd.ms-fontobject&name=static/[name].[ext]&publicPath=../',
      },
    ],
  },
  plugins: [
    new ExtractTextPlugin({
      filename: '/static/[name].[hash].css',
      disable: false,
      allChunks: true,
    }),
    new HtmlWebpackPlugin({
      filename: 'index.html',
      template: './index.html',
    }),
    new webpack.optimize.CommonsChunkPlugin({
      name: ['vendor', 'manifest'],
      minChunks: 'Infinity',
    }),
    new webpack.DefinePlugin({
      'process.env': {
        NODE_ENV: JSON.stringify(process.env.NODE_ENV),
      },
    }),
  ],
};
