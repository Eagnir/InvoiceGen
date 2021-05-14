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
            <router-link to="/about" class="nav-link float-start"
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
        class="col-12 footer-notice align-self-end p-4 text-center text-lg-right"
      >
        <a href="https://github.com/eagnir" target="_blank">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
          >
            <path
              d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"
            />
          </svg>
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
export default class Home extends Vue {
  email: string = "";
  passw: string = "";

  public mounted() {
    this.AutoLogin();
  }

  private AutoLogin() {
    var usr = igApp.Instance.User;
    if (usr != null) {
      igApi.heartbeat().then((resp) => {
        this.$toast.info("Your session has been continued")
        this.$router.push({ name: "Home" });
      });
    }
  }

  private onLogin(): void {
    igApi.authCredential(this.email, this.passw).then((resp) => {
      igApp.Instance.User = resp.Data?.pop();
      this.$router.push({ name: "Home" });
    });
  }
}
</script>
