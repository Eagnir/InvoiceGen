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

import VueHotkey from 'v-hotkey';

//import "@/IG/app";
//import { API as igApi } from "@/IG/api";
//import { Server as igServer } from "@/IG/server";
//import { Browser as igBrowser } from "@/IG/browser";

//Vue.prototype.$igApi = igApp;
//Vue.prototype.$igApi = igApi;
//Vue.prototype.$igServer = igServer;
//Vue.prototype.$igBrowser = igBrowser;

//Vue.use(App,{})



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

new Vue({
  router,
  render: (h) => h(App),
}).$mount("#app");