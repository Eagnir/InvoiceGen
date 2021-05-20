<template>
  <div class="nav-panel">
    <div
      class="flex-column p-3 bg-dark h-100 sidebar"
      :class="{ iconsOnlyBar: sidebarIconsOnly }"
    >
      <router-link
        :to="{ name: 'Dashboard' }"
        exact
        class="logo d-flex align-items-center mb-3 mb-md-0 me-md-auto text-decoration-none"
      >
        <i class="bi-building me-2"></i
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
            <i class="bi-plus-square"></i> <span>Create Invoice</span>
          </button>
        </li>
        <li>
          <hr />
        </li>
        <li>
          <router-link
            :to="{ name: 'Dashboard' }"
            exact
            class="nav-link"
            data-bs-toggle="tooltip"
            data-bs-placement="right"
            :title="this.$ig.keymap.NavDashboard()"
            v-hotkey.stop="
              this.$ig.keymap.NavDashboard(this.NavTo('Dashboard'))
            "
            ><i class="bi-clipboard-data me-3"></i>
            <span><keymap>D</keymap>ashboard</span></router-link
          >
        </li>
        <li>
          <router-link
            :to="{ name: 'Clients'}"
            class="nav-link"
            data-bs-toggle="tooltip"
            data-bs-placement="right"
            :title="this.$ig.keymap.NavClients()"
            v-hotkey.stop="this.$ig.keymap.NavClients(this.NavTo('Clients'))"
            ><i class="bi-people me-3"></i>
            <span><keymap>C</keymap>lients</span></router-link
          >
        </li>
        <li>
          <router-link
            :to="{ name: 'Invoices'}"
            class="nav-link"
            data-bs-toggle="tooltip"
            data-bs-placement="right"
            :title="this.$ig.keymap.NavInvoices()"
            v-hotkey.stop="this.$ig.keymap.NavInvoices(this.NavTo('Invoices'))"
            ><i class="bi-receipt-cutoff me-3"></i>
            <span><keymap>I</keymap>nvoices</span></router-link
          >
        </li>
        <li>
          <router-link
            :to="{ name: 'Reports'}"
            class="nav-link"
            data-bs-toggle="tooltip"
            data-bs-placement="right"
            :title="this.$ig.keymap.NavReports()"
            v-hotkey.stop="this.$ig.keymap.NavReports(this.NavTo('Reports'))"
            ><i class="bi-file-earmark-text me-3"></i>
            <span><keymap>R</keymap>eports</span></router-link
          >
        </li>
      </ul>
      <a
        class="align-self-end"
        href="#"
        @click="toggleSidebar"
        data-bs-toggle="tooltip"
        data-bs-placement="right"
        :data-bs-original-title="
          (sidebarIconsOnly ? 'Show all (' : 'Icons only (') +
          this.$ig.keymap.ToggleSidebar() +
          ')'
        "
        v-hotkey.stop="this.$ig.keymap.ToggleSidebar(this.toggleSidebar)"
      >
        <i
          :class="
            sidebarIconsOnly ? 'bi-arrow-right-square' : 'bi-arrow-left-square'
          "
        ></i>
      </a>
      <hr />
      <div class="dropdown">
        <a
          href="#"
          class="d-flex align-items-center text-decoration-none dropdown-toggle"
          id="dropdownUser1"
          data-bs-toggle="dropdown"
          aria-expanded="false"
        >
          <div class="me-auto d-flex">
            <i class="bi-person-badge me-3"></i
            ><span>
              <strong class="wrap-text">{{ userName }}</strong></span
            >
          </div>
        </a>
        <ul
          class="dropdown-menu dropdown-menu-dark text-small shadow"
          aria-labelledby="dropdownUser1"
        >
          <li>
            <router-link
              :to="{ name: 'AuthAbout' }"
              class="nav-link"
              data-bs-toggle="tooltip"
              data-bs-placement="right"
              title="View details about this site"
              ><i class="bi-info-square me-1"></i> About</router-link
            >
          </li>
          <li>
            <router-link
              :to="{ name: 'Settings' }"
              class="nav-link"
              data-bs-toggle="tooltip"
              data-bs-placement="right"
              :title="this.$ig.keymap.NavSettings()"
              v-hotkey.stop="
                this.$ig.keymap.NavSettings(this.NavTo('Settings'))
              "
              ><i class="bi-gear me-1"></i
              >&nbsp;<keymap>S</keymap>ettings</router-link
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
              ><i class="bi-door-open me-1"></i> Sign <keymap>O</keymap>ut</a
            >
          </li>
        </ul>
      </div>
    </div>

    <div class="bottom-bar-title">
      <a class="nofocus" href="#" @click.prevent="NavTo('Dashboard', true)">
        <i class="bi-building me-2"></i
        ><span class="fs-5">{{ companyName }}</span>
      </a>
    </div>
    <div class="bottom-bar">
      <router-link :to="{ name: 'Dashboard' }" class="nav-link"
        ><i class="bi-clipboard-data fs-2"></i>
      </router-link>
      <router-link :to="{ name: 'Clients' }" class="nav-link"
        ><i class="bi-people fs-2"></i>
      </router-link>
      <router-link :to="{ name: 'Invoices' }" class="nav-link"
        ><i class="bi-receipt-cutoff fs-2"></i>
      </router-link>
      <router-link :to="{ name: 'Reports' }" class="nav-link"
        ><i class="bi-file-earmark-text fs-2"></i>
      </router-link>

      <div class="dropdown">
        <a
          href="#"
          class="nav-link"
          id="dropdownUser1"
          data-bs-toggle="dropdown"
          aria-expanded="false"
        >
          <i class="bi-person-badge fs-2"></i>
        </a>
        <ul
          class="dropdown-menu dropdown-menu-dark text-small shadow"
          aria-labelledby="dropdownUser1"
        >
          <li>
            <router-link
              :to="{ name: 'AuthAbout' }"
              class="nav-link"
              data-bs-toggle="tooltip"
              data-bs-placement="right"
              title="View details about this site"
              ><i class="bi-info-square me-1"></i> About</router-link
            >
          </li>
          <li>
            <router-link
              :to="{ name: 'Settings' }"
              class="nav-link"
              data-bs-toggle="tooltip"
              data-bs-placement="right"
              ><i class="bi-gear me-1"></i>&nbsp;Settings</router-link
            >
          </li>
          <li><hr class="dropdown-divider" /></li>
          <li>
            <a
              class="dropdown-item"
              href="#"
              @click.prevent="confirmSignOut"
              data-bs-toggle="tooltip"
              data-bs-placement="right"
              ><i class="bi-door-open me-1"></i> Sign Out</a
            >
          </li>
        </ul>
      </div>
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

  sidebarIconsOnly: boolean = false;

  public mounted() {
    this.userName = igApp.Instance.User.name;
    this.companyName = igApp.Instance.User.company.name;

    const el = document.getElementById("confirmSignOut");
    this.signOutConfirmationPopup = new Modal(el);
    el.addEventListener("shown.bs.modal", function (event) {
      var x: HTMLElement = el.querySelector(".confirm");
      x.focus();
    });
  }

  toggleSidebar() {
    this.sidebarIconsOnly = this.sidebarIconsOnly ? false : true;
  }

  createInvoice() {
    alert("Hello World");
  }

  NavTo(routeName: string, exec: boolean = false) {
    var nfn = () => {
      console.log(this.$router.currentRoute.name);
      if (this.$router.currentRoute.name != routeName) {
        this.$router.push({ name: routeName });
      } else this.$swal.toast.info("You are already in " + routeName);
    };
    if (exec) nfn();
    return nfn;
  }

  confirmSignOut() {
    this.$swal.confirm("Do you want to sign out?",{
      confirmButtonText: "Yes, sign out"
    })
      .then((res) => {
        if(res.isConfirmed)
          this.signout();
      });
  }

  signout() {
    this.signOutConfirmationPopup.hide();
    this.$ig.api.signout().then((resp) => {
      this.$router.push({ name: "Login" });
    });
  }
}
</script>