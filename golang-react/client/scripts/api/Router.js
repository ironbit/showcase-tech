import axios from 'axios';

export default {
  get(url, process) {
    axios({
      method: 'get',
      url
    }).then(function(response) {
      process(response);
    })
  },

  post(url, data, process) {
    axios({
      method: 'post',
      url,
      data,
      headers: {
        'Content-Type': 'application/json'
      }
    }).then(function(response) {
      process(response);
    })
  }
};
