import Vue from "vue";
import App from "./App.vue";
import "./registerServiceWorker";
import router from "./router";

// Filters
import "./filters/IGDate";
import "./filters/IGTime";
import "./filters/IGAmount";

// App SCSS Files
import "@/scss/app.scss";
import "bootstrap";

// Directives


// Plugins
import IGPlugin from "@/IG/igPlugin";
import moment from "moment-timezone";
import VueMoment from "vue-moment";
import SwalPlugin from "@/libs/SwalPlugin";
import VueHotkey from "v-hotkey";
import Maska from 'maska';
import loading from 'vuejs-loading-screen';

Vue.use(loading, {
  bg: '#000000ad',
  slot:'<div class="d-flex align-items-center text-primary">\
  <div class="spinner-border ms-auto" role="status" aria-hidden="true"></div>\
  <span class="ps-2">Loading...</span>\
</div>',
  /* icon: 'bi bi-arrow-repeat',
  size: 3,
  icon_color: 'white', */
});
Vue.use(Maska);
Vue.use(IGPlugin);
Vue.use(VueMoment, {
  moment,
});
Vue.use(SwalPlugin);
Vue.use(VueHotkey);

Vue.config.productionTip = false;

Vue.config.ignoredElements = ["keymap"];

const mainVueApp = new Vue({
  router,
  render: (h) => h(App),
}).$mount("#app");

export default mainVueApp;
