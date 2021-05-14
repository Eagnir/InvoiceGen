<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { API as igApi } from "@/IG/api";
import { APIResponseStatus } from "./entity/response";

@Component({
})
export default class App extends Vue {
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

  mounted() {
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