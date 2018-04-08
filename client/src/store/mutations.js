// let state = require('./state')
// let getters = require('./getters')
// let actions = require('./actions')
// let merge = require('webpack-merge')
// let baseMutations = require('../../../common/store/mutations')

module.exports = {
  LOADING(state) {
    state.status.loading = true;
  },
};
