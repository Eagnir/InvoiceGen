<template>
  <div class="d-flex flex-row h-100">
    <div>
      <NavBar></NavBar>
    </div>
    <div class="flex-fill">
      <router-view />
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import NavBar from "@/views/NavBar.vue";
import { API as igApi } from "@/IG/api";
import { APIResponseStatus } from "./entity/response";
import { Popover, Tooltip } from "bootstrap";

@Component({
  components: {
    NavBar,
  },
})
export default class Auth extends Vue {
  beforeMount() {
    igApi
      .heartbeat({
        RejectOnFailure: false,
      })
      .then((resp) => {
        if (resp.Status != APIResponseStatus.StatusSuccess) {
          this.$toast.error("Your session has expired");
          this.$router.push({ name: "Login" });
        }
      });
  }

  initBS() {
    //Tooltips
    var tooltipTriggerList = [].slice.call(
      document.querySelectorAll('[data-bs-toggle="tooltip"]')
    );
    var tooltipList = tooltipTriggerList.map(function (tooltipTriggerEl) {
      return new Tooltip(tooltipTriggerEl, {
        delay: { show: 1000, hide: 0 },
      });
    });

    //Popovers
    var popoverTriggerList = [].slice.call(
      document.querySelectorAll('[data-bs-toggle="popover"]')
    );
    var popoverList = popoverTriggerList.map(function (popoverTriggerEl) {
      return new Popover(popoverTriggerEl);
    });
  }

  mounted() {
    this.initBS();
    this.$router.afterEach((to, from) => {
      this.initBS();
    });

    this.$router.beforeEach((to, from, next) => {
      if (to.name == "Login") {
        next();
        return;
      }

      igApi
        .heartbeat({
          RejectOnFailure: false,
        })
        .then((resp) => {
          if (resp.Status != APIResponseStatus.StatusSuccess) {
            this.$toast.error("Your session has expired");
            next({ name: "Login" });
          }
        });
      next();
    });
  }
}
</script>