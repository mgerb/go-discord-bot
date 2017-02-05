var webpack = require("./webpack.config.js")

webpack.plugins[0].filename = "/static/[name].[hash].css";
webpack.output.filename = '/static/[name].[hash].js';

module.exports = webpack;