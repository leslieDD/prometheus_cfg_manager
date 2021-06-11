'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

if (process.env.NODE_ENV === 'development') {
  module.exports = merge(prodEnv, {
    NODE_ENV: '"development"',
    VUE_APP_BASE_API: '"http://127.0.0.1:8200"'
  })
} else {
  module.exports = prodEnv
}
