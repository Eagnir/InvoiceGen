<template>
  <div>
    <div class="row align-items-center">
      <div class="col">
        <h1 class="m-0">{{ clientName }}</h1>
      </div>
      <div class="col-12 col-md-5 text-end">
        <button type="button" class="btn btn-default mb-3 me-4">Reset</button>
        <button type="button" class="btn btn-success mb-3">Save Client</button>
      </div>
    </div>
    <hr />
    <nav aria-label="breadcrumb">
      <ol class="breadcrumb">
        <li class="breadcrumb-item">
          <router-link :to="{ name: 'Clients' }">Clients</router-link>
        </li>
        <li class="breadcrumb-item active">Detail</li>
        <li class="breadcrumb-item active" aria-current="page">
          #{{ clientId }}
        </li>
      </ol>
    </nav>
    <div class="row">
      <router-link
        :to="{ name: 'ClientDetails', params: { clientId: 1 } }"
        custom
        v-slot="{ navigate }"
      >
        <div class="col col-md-6 col-xl-4 c-pointer" @click="navigate">
          <div class="card mb-3 d-flex">
            <div class="d-flex">
              <img src="https://via.placeholder.com/120" />
              <div class="card-body">
                <h5 class="card-title">Client Name</h5>
                <p class="card-text mb-1">FY: INR 10,20,000/-</p>
                <p class="card-text m-0">
                  <small class="text-muted">3 months ago</small>
                </p>
              </div>
            </div>
          </div>
        </div>
      </router-link>
    </div>
  </div>
</template>
<script lang="ts">
import { Prop, Vue } from "vue-property-decorator";
import Component from "vue-class-component";

Component.registerHooks([
  "beforeRouteEnter",
  "beforeRouteUpdate",
  "beforeRouteLeave",
]);

@Component({
  components: {},
})
export default class ClientDetails extends Vue {
  @Prop({ default: "0" })
  clientId: number;
  clientName: string = "Loading...";

  beforeRouteLeave(to: any, from: any, next: any) {
    this.$swal.confirm("Do you want to discard changes?")
      .then((res) => {
        if(!res.isConfirmed)
          return;
          
        this.$swal.toast.info("Changes are discarded");
        next();
      })
      .catch((x) => {
        next(false);
      });
  }
}
</script>