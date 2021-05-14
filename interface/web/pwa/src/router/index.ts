import Vue from "vue";
import VueRouter, { RouteConfig } from "vue-router";
import Dashboard from "../views/Dashboard.vue";
import Error from "../views/Error.vue";
import Login from "../views/Login.vue";
import About from "../views/About.vue";
import Invoices from "../views/Invoices.vue";
import Clients from "../views/Clients.vue";
import Reports from "../views/Reports.vue";
import Settings from "../views/Settings.vue";

import Auth from "../Auth.vue";


Vue.use(VueRouter);

const routes: Array<RouteConfig> = [

  {
    path: "/",
    name: "Login",
    component: Login,
  },
  {
    path: "*",
    name: "LoginDefault",
    component: Login,
  },
  {
    path: "/about",
    name: "About",
    component: About,
  },
  {
    path: '/auth',
    component: Auth,
    children: [
      {
        path: "/",
        name: "Home",
        component: Dashboard,
      },
      {
        path: "dashboard",
        name: "Dashboard",
        component: Dashboard,
      },
      {
        path: "about",
        name: "AuthAbout",
        component: About,
      },

      {
        path: "invoices",
        name: "Invoices",
        component: Invoices,
      },

      {
        path: "clients",
        name: "Clients",
        component: Clients,
      },

      {
        path: "reports",
        name: "Reports",
        component: Reports,
      },

      {
        path: "settings",
        name: "settings",
        component: Settings,
      },
    ]
  },
  { path: "/error/msg?", component: Error },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  linkActiveClass: "active",
  linkExactActiveClass: "exact-active",
  routes,
});

export default router;
