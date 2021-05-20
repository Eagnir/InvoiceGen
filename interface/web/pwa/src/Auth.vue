<template>
  <div class="root-panel">    
    <NavBar></NavBar>
    <div class="flex-fill p-lg-5 p-3 content-panel">
      <router-view />
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import NavBar from "@/views/NavBar.vue";
import { APIResponseStatus } from "./entity/response";
import { Popover, Tooltip } from "bootstrap";

@Component({
  components: {
    NavBar,
  },
})
export default class Auth extends Vue {
  beforeMount() {
    this.$ig.api
      .heartbeat({
        RejectOnFailure: false,
      })
      .then((resp) => {
        if (resp.Status != APIResponseStatus.StatusSuccess) {
          this.$swal.toast.error("Your session has expired");
          this.$router.push({ name: "Login" });
        }
      });


    this.$router.beforeEach((to, from, next) => {
      if (to.name == "Login") {
        next();
        return;
      }

        console.log(to);
      this.$ig.api
        .heartbeat({
          RejectOnFailure: false,
        })
        .then((resp) => {
          if (resp.Status != APIResponseStatus.StatusSuccess) {
            this.$swal.toast.error("Your session has expired");
            next({ name: "Login" });
          }
        });
      next();
    });


    this.$router.afterEach((to, from) => {
      
        console.log("Reached");
      this.$ig.app.initBS();
    });
  }

  mounted() {
    this.$ig.app.initBS();

  }
}
</script>