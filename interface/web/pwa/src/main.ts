import Vue from "vue";
import App from "./App.vue";
import "./registerServiceWorker";
import router from "./router";

// Filters
import "./filters/IGDate";
import "./filters/IGTime";
import VueMoment from "vue-moment";
import moment from "moment-timezone";

import "@/scss/app.scss";


import VueToast from 'vue-toast-notification';
import 'vue-toast-notification/dist/theme-default.css';
//import 'vue-toast-notification/dist/theme-sugar.css';

import IGPlugin from "@/IG/igPlugin";
Vue.use(IGPlugin);

import VueHotkey from 'v-hotkey';

Vue.use(VueHotkey)
import "bootstrap";

Vue.use(VueMoment, {
  moment,
})

Vue.use(VueToast,{
  // One of the options
  position: 'top-right'
});

Vue.config.productionTip = false;

Vue.config.ignoredElements = [
  'keymap'
]

new Vue({
  router,
  render: (h) => h(App),
}).$mount("#app");