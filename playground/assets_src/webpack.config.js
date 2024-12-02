const path = require('path');

module.exports = {
  entry: './base.js',
  output: {
    path: path.resolve(__dirname, '../assets/js'),
    filename: 'bundle.js',
  },
};
