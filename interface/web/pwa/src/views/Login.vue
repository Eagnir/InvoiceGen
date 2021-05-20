<template>
  <div class="container-fluid" style="height: 100%">
    <div class="row h-100 justify-content-center">
      <div class="col-12 align-self-end"></div>
      <div class="col-12 col-md-6 col-lg-3 align-self-center">
        <h1 class="text-center mb-3">Welcome</h1>
        <div class="row">
          <div class="col">
            <div class="mb-2 form-floating">
              <input
                v-model.trim="email"
                type="email"
                class="form-control"
                id="email"
                placeholder="name@domain.com"
                value=""
                autofocus
              />
              <label for="email">Email Address</label>
            </div>

            <div class="mb-2 form-floating">
              <input
                v-model.trim="passw"
                v-on:keyup.enter="onLogin"
                type="password"
                class="form-control"
                id="passw"
                placeholder="try to remember :)"
                value=""
              />
              <label for="passw">Password</label>
            </div>
            <router-link to="/about" class="float-start"
              ><i class="bi-info-square me-1"></i> About</router-link
            >
            <button
              type="submit"
              @click="onLogin"
              class="btn btn-primary float-end"
            >
              Sign In
            </button>
          </div>
        </div>
        <!--
        <p class="text-center">{{ "please login to continue" | capitalize }}</p>
        {{ new Date() | moment("dddd, MMMM Do YYYY") }}
        {{ msg }}
        -->
      </div>
      <div
        class="col-12 footer-notice align-self-end p-4 text-center text-lg-right d-flex justify-content-center"
      >
        <a href="https://github.com/eagnir" target="_blank" class="d-flex align-items-center">
          <i class="bi-github fs-2 me-2"></i>
          @Eagnir (Nirav Shah)
        </a>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { API as igApi } from "@/IG/api";
import { App as igApp, UserCredential } from "@/IG/app";
import { APIResponseStatus } from "@/entity/response";

@Component({
  components: {},
})
export default class Login extends Vue {
  email: string = "";
  passw: string = "";

  public mounted() {
    this.AutoLogin();

    //this.$swal.toast.fire("Hello Default");
    //this.$swal.toast.info("Hello Info");
    //this.$swal.toast.warning("Hello Warning");
    //this.$swal.toast.error("Hello Error");
    //this.$swal.toast.success("Hello Success");

  }

  private AutoLogin() {
    var usr = igApp.Instance.User;
    if (usr != null) {
      this.$ig.api.heartbeat().then((resp) => {
        this.$swal.toast.info("Your session has been continued")
        this.$router.push({ name: "Dashboard" });
      });
    }
  }

  private onLogin(): void {
    this.$ig.api.authCredential(this.email, this.passw).then((resp) => {
      igApp.Instance.User = resp.Data?.pop();
      this.$router.push({ name: "Dashboard" });
    });
  }
}
</script>
