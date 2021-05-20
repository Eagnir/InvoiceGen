import Vue from "vue";
import VueRouter, { RouteConfig } from "vue-router";
import Dashboard from "../views/Dashboard.vue";
import Error from "../views/Error.vue";
import Login from "../views/Login.vue";
import About from "../views/About.vue";

import InvoiceDefault from "../views/Invoices/Default.vue";
import InvoiceList from "../views/Invoices/List.vue";
import InvoiceDetails from "../views/Invoices/Details.vue";

import ReportDefault from "../views/Reports/Default.vue";
import ReportList from "../views/Reports/List.vue";

import ClientDefault from "../views/Clients/Default.vue";
import ClientList from "../views/Clients/List.vue";
import ClientDetails from "../views/Clients/Details.vue";

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
        path: "",
        redirect: { name: 'Dashboard' },
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
        path: "clients",
        component: ClientDefault,
        children: [
          {
            path: "",
            name:"Clients",
            redirect: { name: 'ClientList' },
          },
          {
            path: "list",
            name: "ClientList",
            component: ClientList,
          },
          {
            path: ":clientId/details/",
            name: "ClientDetails",
            props: true,
            component: ClientDetails,
          }
        ]
      },

      {
        path: "invoices",
        component: InvoiceDefault,
        children: [
          {
            path: "",
            name:"Invoices",
            redirect: { name: 'InvoiceList' },
          },
          {
            path: "list",
            name: "InvoiceList",
            component: InvoiceList,
          },
          {
            path: ":invoiceId/details",
            name: "InvoiceDetails",
            props: true,
            component: InvoiceDetails,
          }
        ]
      },


      {
        path: "reports",
        component: ReportDefault,
        children: [
          {
            path: "",
            name:"Reports",
            redirect: { name: 'ReportList' },
          },
          {
            path: "list",
            name: "ReportList",
            component: ReportList,
          },
        ]
      },

      {
        path: "settings",
        name: "Settings",
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
