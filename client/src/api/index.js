/* eslint-disable */
import axios from 'axios';
import store from '../store';

const request = (method, uri, data = null) => {
  if (!method) {
    console.error('API function call requires method argument');
    return;
  }

  if (!uri) {
    console.error('API function call requires uri argument');
    return;
  }

  const url = store.state.serverURI + uri;
  // console.log('api', url);
  return axios({ method, url, data });
};

/*
const getData = (method, uri, data = null, successMutation = null) => {
  // console.log('getData', store);
  store.commit('LOADING');
  return request('get', uri)
    .then((response) => {
      // console.log('getData response', response);
      if (response.data.status !== 'ok') {
        store.commit('LOADING_ERROR', { message: 'Response status is not ok', code: 406 });
        return;
      }

      store.commit('LOADED');
      return response.data.data;
    })
    .catch((error) => {
      store.commit('LOADING_ERROR', error);
    });
};
*/
export default {
  request,
  // getData,
};
