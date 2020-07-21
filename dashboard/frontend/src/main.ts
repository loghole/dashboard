import Vue from 'vue';
import Buefy from 'buefy';
import axios from 'axios';
import VueAxios from 'vue-axios';
import Fragment from 'vue-fragment';
import App from './App.vue';
import router from './router';
import store from './store';
import './assets/scss/app.scss';

Vue.use(VueAxios, axios);
Vue.use(Buefy);
Vue.use(Fragment.Plugin);

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
