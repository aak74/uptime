import api from '../api';

const loadAll = ({ state }) => {
  console.warn('api', api);
  api.request('get', 'projects', {})
    .then((response) => {
      console.warn('response', response);
      state.data.sites = response.data;
    });
  state.status.loading = false;
};
/*

const ping = ({
  state,
}) => {
  state.status.pinging = true;
};

const start = ({
  state,
}) => {
  // state.interval = setInterval(next(state), 2000)
  state.status.pinging = true;
  state.interval = setInterval(() => {
    // console.log('next', state.current)
    const current = state.current;
    const url = state.data.sites[current].url;
    state.data.sites[current].status = 'pinging';
    state.data.sites.splice(state.data.sites.length);
    axios.get(url)
      .then((response) => {
        // console.log('response', url, response.status)
        // console.log('response', url, response.code)
        state.data.sites[current].status = 'warning';
        state.data.sites[current].code = response.status;
        if (response.status === 200) {
          state.data.sites[current].status = 'success';
        }
        state.data.sites.splice(state.data.sites.length);
      })
      .catch((error) => {
        // console.error('error', url, error);
        state.data.sites[current].status = 'danger';
        state.data.sites.splice(state.data.sites.length);
      });
    if (state.current + 1 < state.data.sites.length) {
      state.current++;
    } else {
      state.current = 0;
    }
  }, 500);
};

const stop = ({
  state,
}) => {
  state.status.pinging = false;
  clearInterval(state.interval);
};
/*
const next = (state) => {
  console.log('next', state.current)
  if (state.current < state.data.sites.length) {
    state.current++
  } else {
    state.current = 0
  }
}
*/
export default {
  loadAll,
  // ping,
  // start,
  // stop,
};
