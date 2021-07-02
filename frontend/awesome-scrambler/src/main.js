import Router from 'vue-router'
import Vue from 'vue'
import App from './App.vue'

import {routes} from './routes'

Vue.config.productionTip = false

const router = new Router({
  routes: routes,
  mode: 'history',
})

Vue.use(Router)

new Vue({
  router,
  render: h => h(App)
}).$mount('#app');