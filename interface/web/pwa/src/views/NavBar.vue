<template>
  <div
    class="d-flex flex-column p-3 text-white bg-dark h-100"
    style="width: 280px"
  >
    <router-link
      to="/auth/"
      exact
      class="d-flex align-items-center mb-3 mb-md-0 me-md-auto text-white text-decoration-none"
    >
      <i class="bi-building" style="margin-right: 10px"></i
      ><span class="fs-5">{{ companyName }}</span>
    </router-link>
    <hr />
    <ul class="nav nav-pills flex-column mb-auto">
      <li>
        <button
          class="btn btn-warning w-100"
          @click="createInvoice"
          data-bs-toggle="tooltip"
          data-bs-placement="right"
          :title="this.$ig.keymap.CreateInvoice()"
          v-hotkey.stop="this.$ig.keymap.CreateInvoice(this.createInvoice)"
        >
          <i class="bi-plus-square"></i> Create Invoice
        </button>
      </li>
      <li>
        <hr />
      </li>
      <li>
        <router-link
          to="/auth/dashboard"
          exact
          class="nav-link text-white"
          data-bs-toggle="tooltip"
          data-bs-placement="right"
          :title="this.$ig.keymap.NavDashboard()"
          v-hotkey.stop="this.$ig.keymap.NavDashboard(this.NavTo('Dashboard'))"
          ><i class="bi-clipboard-data me-3"></i>
          <keymapped>D</keymapped>ashboard</router-link
        >
      </li>
      <li>
        <router-link
          to="/auth/invoices"
          class="nav-link text-white"
          data-bs-toggle="tooltip"
          data-bs-placement="right"
          :title="this.$ig.keymap.NavInvoices()"
          v-hotkey.stop="this.$ig.keymap.NavInvoices(this.NavTo('Invoices'))"
          ><i class="bi-receipt-cutoff me-3"></i>
          <keymapped>I</keymapped>nvoices</router-link
        >
      </li>
      <li>
        <router-link
          to="/auth/clients"
          class="nav-link text-white"
          data-bs-toggle="tooltip"
          data-bs-placement="right"
          :title="this.$ig.keymap.NavClients()"
          v-hotkey.stop="this.$ig.keymap.NavClients(this.NavTo('Clients'))"
          ><i class="bi-people me-3"></i>
          <keymapped>C</keymapped>lients</router-link
        >
      </li>
      <li>
        <router-link
          to="/auth/reports"
          class="nav-link text-white"
          data-bs-toggle="tooltip"
          data-bs-placement="right"
          :title="this.$ig.keymap.NavReports()"
          v-hotkey.stop="this.$ig.keymap.NavReports(this.NavTo('Reports'))"
          ><i class="bi-file-earmark-text me-3"></i>
          <keymapped>R</keymapped>eports</router-link
        >
      </li>
    </ul>
    <hr />
    <div class="dropdown">
      <a
        href="#"
        class="d-flex align-items-center text-white text-decoration-none dropdown-toggle"
        id="dropdownUser1"
        data-bs-toggle="dropdown"
        aria-expanded="false"
      >
        <img
          src="https://github.com/mdo.png"
          alt=""
          width="32"
          height="32"
          class="rounded-circle me-2 d-none"
        />
        <span class="me-auto d-flex"
          ><i class="bi-person-badge me-3"></i>
          <strong class="wrap-text">{{ userName }}</strong></span
        >
      </a>
      <ul
        class="dropdown-menu dropdown-menu-dark text-small shadow"
        aria-labelledby="dropdownUser1"
      >
        <li>
          <router-link
            to="/auth/about"
            class="nav-link text-white"
            data-bs-toggle="tooltip"
            data-bs-placement="right"
            title="View details about this site"
            ><i class="bi-info-square me-1"></i> About</router-link
          >
        </li>
        <li>
          <router-link
            to="/auth/settings"
            class="nav-link text-white"
            data-bs-toggle="tooltip"
            data-bs-placement="right"
            :title="this.$ig.keymap.NavSettings()"
            v-hotkey.stop="this.$ig.keymap.NavSettings(this.NavTo('Settings'))"
            ><i class="bi-gear me-1"></i>
            <keymapped>S</keymapped>ettings</router-link
          >
        </li>
        <li><hr class="dropdown-divider" /></li>
        <li>
          <a
            :title="this.$ig.keymap.SignOut()"
            v-hotkey.stop="this.$ig.keymap.SignOut(this.confirmSignOut)"
            class="dropdown-item"
            href="#"
            @click.prevent="confirmSignOut"
            data-bs-toggle="tooltip"
            data-bs-placement="right"
            ><i class="bi-door-open me-1"></i> Sign
            <keymapped>o</keymapped>ut</a
          >
        </li>
      </ul>
    </div>

    <!-- Modal -->
    <div class="modal fade" id="confirmSignOut" tabindex="-1">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="exampleModalLabel">
              Confirm Signing Out
            </h5>
          </div>
          <div class="modal-body">Do you really want to sign out?</div>
          <div class="modal-footer">
            <button
              type="button"
              class="btn btn-secondary"
              data-bs-dismiss="modal"
            >
              Cancel
            </button>
            <button
              type="button"
              class="btn btn-primary confirm"
              autofocus
              @click="signout"
            >
              Yes, Sign Out
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
  <!-- 
<nav class="navbar navbar-expand-lg navbar-light bg-light">
  <div class="container-fluid">
    <a class="navbar-brand" href="#">InvoiceGen</a>
    <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNavDropdown" aria-controls="navbarNavDropdown" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarNavDropdown">
      <ul class="navbar-nav">
        <li class="nav-item">
          <a class="nav-link active" aria-current="page" href="#">Dashboard</a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="#">|CompanyName|</a>
        </li>        
        <li class="nav-item dropdown">
          <a class="nav-link dropdown-toggle" href="#" id="navbarDropdownMenuLink" role="button" data-bs-toggle="dropdown" aria-expanded="false">
            Invoices
          </a>
          <ul class="dropdown-menu" aria-labelledby="navbarDropdownMenuLink">
            <li><a class="dropdown-item" href="#">All Invoices</a></li>
            <li><a class="dropdown-item" href="#">Unpaid Invoices</a></li>
          </ul>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="#">Clients</a>
        </li>
      </ul>
      <button class="btn btn-primary"><i class="bi-plus-square"></i>
 Create Invoice</button>
    </div>
  </div>
</nav> -->

  <!-- 
  <div id="nav">
    <router-link to="/">Dashboard</router-link> |
    <router-link to="/about">Company</router-link> |
    <router-link to="/login">Clients</router-link>
    <router-link to="/login">Invoices</router-link>
    <router-link to="/login">Purchases</router-link>
  </div> -->
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { API as igApi } from "@/IG/api";
import { App as igApp, UserCredential } from "@/IG/app";
import { APIResponseStatus } from "@/entity/response";
import { Modal } from "bootstrap";

@Component({
  components: {},
})
export default class Home extends Vue {
  userName: string = "";
  companyName: string = "";

  signOutConfirmationPopup: Modal;

  public mounted() {
    this.userName = igApp.Instance.User.name;
    this.companyName = igApp.Instance.User.company.name;

    const el = document.getElementById("confirmSignOut");
    this.signOutConfirmationPopup = new Modal(el);
    el.addEventListener("shown.bs.modal", function (event) {
      var x:HTMLElement = el.querySelector(".confirm");
      x.focus();
    });
  }

  createInvoice() {
    alert("Hello World");
  }

  NavTo(routeName: string) {
    return () => {
      if (this.$router.currentRoute.name != routeName) {
        this.$router.push({ name: routeName });
      } else this.$toast.info("You are already in " + routeName);
    };
  }

  confirmSignOut() {
    this.signOutConfirmationPopup.show();
  }

  signout() {
    this.signOutConfirmationPopup.hide();
    this.$ig.api.signout().then((resp) => {
      this.$router.push({ name: "Login" });
    });
  }
}
</script>