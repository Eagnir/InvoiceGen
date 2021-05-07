import Vue from "vue";
import App from "./App.vue";
import "./registerServiceWorker";
import router from "./router";

// Filters
import "./filters/IGDate";
import "./filters/IGTime";
import VueMoment from "vue-moment";
import moment from "moment-timezone";

//import "normalize.css"
import "@/scss/app.scss";


Vue.use(VueMoment, {
  moment,
})

Vue.config.productionTip = false;

new Vue({
  router,
  render: (h) => h(App),
}).$mount("#app");
