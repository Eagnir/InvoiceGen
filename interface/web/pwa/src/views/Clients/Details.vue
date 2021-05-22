<template>
  <div>
    <div class="row align-items-center">
      <div class="col-12 col-md text-md-start text-center">
        <h1 class="m-0">{{ clientName }}</h1>
      </div>
      <div class="col-12 col-md-5 text-md-end d-none d-md-block">
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
          {{ clientName }} #{{ clientId }}
        </li>
      </ol>
    </nav>

    <div class="row">
      <div class="col-md mb-3 mb-md-0">
        <div class="form-floating">
          <input type="text" class="form-control" id="companyName" placeholder="Abc Pvt. Ltd.">
          <label for="companyName">Company Name</label>
        </div>
      </div>
      <div class="col-md mb-3 mb-md-0">
        <div class="form-floating">
          <input type="text" class="form-control" id="gstNumber" placeholder="12AAAAA0000A1Z5">
          <label for="gstNumber">GST Number</label>
        </div>
      </div>
      <div class="col-md">
        <div class="form-floating">
          <select class="form-select" id="defaultCurrency">
            <option selected>Select Currency</option>
            <option value="1">INR</option>
            <option value="2">USD</option>
            <option value="3">CAD</option>
          </select>
          <label for="defaultCurrency">Default Currency</label>
        </div>
      </div>
    </div>

    <div class="row mt-3">
      <div class="col">
        <div class="form-floating">
          <input type="text" class="form-control" id="address" placeholder="Office #, Building name, street, city, zip, state, country.">
          <label for="address">Postal Address</label>
        </div>
      </div>
    </div>

    <div class="row mt-3">
      <div class="col-md mb-3 mb-md-0">
        <div class="form-floating">
          <input type="text" class="form-control" id="email" placeholder="name@domain.com">
          <label for="email">Email</label>
        </div>
      </div>
      <div class="col-md">
        <div class="form-floating">
          <input type="text" class="form-control" id="contactNumber" placeholder="1234567890">
          <label for="contactNumber">Contact Number</label>
        </div>
      </div>
    </div>

    <div class="row mt-4">
      <div class="col-12 text-center d-md-none">
        <button type="button" class="btn btn-default mb-3 me-4">Reset</button>
        <button type="button" class="btn btn-success mb-3">Save Client</button>
      </div>
    </div>

    <div class="row mt-4">
      <div class="col">
        <h5 class="d-flex">
          <span class="flex-grow-1">Client Invoices</span>
          <span>
            <span title="Pending invoices" class="badge bg-warning">3</span>&nbsp;
            <span title="Paid invoices" class="badge bg-success">4</span>&nbsp;
            <span title="Total invoices" class="badge bg-secondary">7</span>
          </span>
          </h5>
        <hr>
      </div>
    </div>

    <div class="row">
      <div class="col">
        <table class="table table-striped table-hover">
          <thead>
            <tr>
              <th scope="col">#</th>
              <th scope="col">Date</th>
              <th class="d-none d-md-table-cell" scope="col">Particulars</th>
              <th class="d-none d-md-table-cell" scope="col">Gross Amount</th>
              <th class="text-end" scope="col">Tax</th>
              <th class="text-end" scope="col">Total (INR)</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <th scope="row">1</th>
              <td>21 May 2021</td>
              <td class="d-none d-md-table-cell">2 items</td>
              <td class="d-none d-md-table-cell">20,000</td>
              <td class="text-end">12,000</td>
              <td class="text-end">32,000</td>
            </tr>
          </tbody>
        </table>
      </div>
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
  @Prop({ default: "Loading..." })
  clientName: string;

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