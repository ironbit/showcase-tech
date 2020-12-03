import Router from './Router.js';

export default {
  fetchProducts(url, process) {
    url += "/products"
    Router.get(url, process)
  },

  storeProduct(url, data, process) {
    url = url + "/product";
    data = JSON.stringify(data);
    Router.post(url, data, process);
  }
}
