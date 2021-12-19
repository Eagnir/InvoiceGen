<template>
  <div class="mb-5 mb-md-0 pb-4 pb-md-0">
    <nav class="d-md-none navbar fixed-bottom navbar-light bg-light sub-nav border border-dark border-start-0 border-end-0 text-center justify-content-center align-items-center px-3 flex-nowrap">
      <button type="button" class="btn flex-fill btn-primary"><i class="bi bi-arrow-repeat"></i> Reset</button>
      <button type="button" class="btn flex-fill btn-success mx-3"><i class="bi bi-check2-square"></i> Save</button>
      <div class="btn-group dropup flex-fill">
        <button type="button" class="btn btn-secondary dropdown-toggle remove-toggle-icon" data-bs-toggle="dropdown" aria-expanded="false">
          <i class="bi bi-list"></i> Actions
        </button>
        <ul class="dropdown-menu dropdown-menu-end">
          <li><button class="dropdown-item text-danger"><i class="bi bi-trash me-2"></i> Delete</button></li>
        </ul>
      </div>
    </nav>

    <div class="row align-items-center">
      <div class="col-12 col-md text-md-start text-center">
        <h1 class="m-0">{{ clientName }}</h1>
      </div>
      <div class="col-12 col-md-5 text-md-end d-none d-md-block">
        <nav class="navbar text-center justify-content-center align-items-center flex-nowrap">
          <button type="button" class="btn flex-fill btn-primary"><i class="bi bi-arrow-repeat"></i> Reset</button>
          <button type="button" class="btn flex-fill btn-success ms-3"><i class="bi bi-check2-square"></i> Save</button>
          <div class="btn-group dropdown flex-fill invisible order-first">
            <button type="button" class="btn btn-secondary dropdown-toggle remove-toggle-icon" data-bs-toggle="dropdown" aria-expanded="false">
              <i class="bi bi-list"></i> Actions
            </button>
            <ul class="dropdown-menu dropdown-menu-end">
              <li><button class="dropdown-item"><i class="bi bi-card-checklist me-2"></i> Change Status</button></li>
              <li><button class="dropdown-item"><i class="bi bi-download me-2"></i> Download</button></li>
              <li><hr class="dropdown-divider"></li>
              <li><button class="dropdown-item text-danger"><i class="bi bi-trash me-2"></i> Delete</button></li>
            </ul>
          </div>
        </nav>
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
          <input type="text" class="form-control" id="companyName" placeholder="Abc Pvt. Ltd." v-model.trim="client.Name">
          <label for="companyName">Company Name</label>
        </div>
      </div>
      <div class="col-md mb-3 mb-md-0">
        <div class="form-floating">
          <input type="text" class="form-control" id="gstNumber" placeholder="12AAAAA0000A1Z5" v-model.trim="client.GSTNumber">
          <label for="gstNumber">GST Number</label>
        </div>
      </div>
      <div class="col-md">
        <div class="form-floating">
          <select class="form-select" id="defaultCurrency"  v-model.trim="client.DefaultCurrency.ShortName">
            <option selected>Select Currency</option>
            <option value="INR">INR</option>
            <option value="USD">USD</option>
            <option value="CAD">CAD</option>
          </select>
          <label for="defaultCurrency">Default Currency</label>
        </div>
      </div>
    </div>

    <div class="row mt-3">
      <div class="col">
        <div class="form-floating">
          <input type="text" class="form-control" id="address" placeholder="Office #, Building name, street, city, zip, state, country." v-model.trim="client.Address">
          <label for="address">Postal Address</label>
        </div>
      </div>
    </div>

    <div class="row mt-3">
      <div class="col-md mb-3 mb-md-0">
        <div class="form-floating">
          <input type="text" class="form-control" id="email" placeholder="name@domain.com" v-model.trim="client.Email">
          <label for="email">Email</label>
        </div>
      </div>
      <div class="col-md">
        <div class="form-floating">
          <input type="text" class="form-control" id="contactNumber" placeholder="1234567890" v-model.trim="client.ContactNumber">
          <label for="contactNumber">Contact Number</label>
        </div>
      </div>
    </div>

    <div class="row mt-4">
      <div class="col">
        <h5 class="d-flex">
          <span class="flex-grow-1">Client Invoices</span>
          <span>
            <span title="Pending invoices" class="badge bg-warning">{{getTotalInvoicesForStatus(1)}}</span>&nbsp;
            <span title="Paid invoices" class="badge bg-success">{{getTotalInvoicesForStatus(2)}}</span>&nbsp;
            <span title="Cancelled invoices" class="badge bg-danger">{{getTotalInvoicesForStatus(3)}}</span>
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
              <th scope="col">Invoice Number</th>
              <th scope="col">Date</th>
              <th class="d-none d-md-table-cell" scope="col">Particulars</th>
              <th class="d-none d-md-table-cell" scope="col">Gross Amount</th>
              <th class="text-end" scope="col">Tax</th>
              <th class="text-end" scope="col">Total</th>
            </tr>
          </thead>
          <tbody>
            <tr v-if="client.Invoices == null || client.Invoices.length<=0">
              <td class="text-center" colspan="7">No invoices found</td>
            </tr>
            <tr v-for="(invoice, index) in client.Invoices" :key="invoice.InvoiceNumber">
              <th scope="row">{{ index+1 }}</th>
              <th>{{ invoice.InvoiceNumber }} <span v-html="getInvoiceStatusText(invoice.Status)"></span>
              </th>
              <td>{{ invoice.CreatedAt | moment('D MMM yyyy') }}</td>
              <td class="d-none d-md-table-cell">
                {{ invoice.InvoiceItems!=null ? invoice.InvoiceItems.length : '--' }}
              </td>
              <td class="d-none d-md-table-cell">
                <span 
                :class="invoice.Currency.ShortName != $ig.app.User.company.defaultCurrency.shortName ? 'convertedCurrency' : ''"
                data-bs-toggle="tooltip"
                data-bs-placement="right"
                :title="invoice.Currency.ShortName != $ig.app.User.company.defaultCurrency.shortName ? $options.filters.igamountRaw(invoice.TaxableAmount,  invoice.Currency.ShortName) : ''"
                v-html="$options.filters.igamount(invoice.TaxableAmount / invoice.Currency.Conversion, 'INR')">
                </span>
              </td>
              <td class="text-end">
                <span 
                :class="invoice.Currency.ShortName != $ig.app.User.company.defaultCurrency.shortName ? 'convertedCurrency' : ''"
                data-bs-toggle="tooltip"
                data-bs-placement="right"
                :title="invoice.Currency.ShortName != $ig.app.User.company.defaultCurrency.shortName ? $options.filters.igamountRaw(invoice.TaxPayable,  invoice.Currency.ShortName) : ''"
                v-html="$options.filters.igamount(invoice.TaxPayable / invoice.Currency.Conversion, 'INR')">
                </span>
                &nbsp;<small>({{ invoice.TaxPercentage }}%)</small>
              </td>
              <td class="text-end">
                <span 
                :class="invoice.Currency.ShortName != $ig.app.User.company.defaultCurrency.shortName ? 'convertedCurrency' : ''"
                data-bs-toggle="tooltip"
                data-bs-placement="right"
                :title="invoice.Currency.ShortName != $ig.app.User.company.defaultCurrency.shortName ? $options.filters.igamountRaw(invoice.InvoiceAmount,  invoice.Currency.ShortName) : ''"
                v-html="$options.filters.igamount(invoice.InvoiceAmount / invoice.Currency.Conversion, 'INR')">
                </span>
                </td>
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
import { APIResponseStatus } from "@/entity/response";

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

  client: any = {
    DefaultCurrency: {}
  };

  didChange: boolean;

  mounted() {
    this.$ig.api.getClientDetails(this.clientId).then((resp) => {
        this.client = resp.Data[0];
        //this.$swal.toast.info("Client successfully loaded");
        this.$nextTick(function(){
          this.$ig.app.initBSTooltips();
        })
    })
  }

  beforeRouteLeave(to: any, from: any, next: any) {
    if(!this.didChange) return next();

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

  getInvoiceStatusText(status:number) {
    var statusText: string = "";
    var statusClass: string = "";
    switch(status) {
      case 0: // Created
          statusText = "Created";
          statusClass = "bg-default";
          break;
        case 1: // Pending
          statusText = "Pending";
          statusClass = "bg-warning";
          break;
        case 2: // Paid
          statusText = "Paid";
          statusClass = "bg-success";
          break;
        case 3: // Cancelled
          statusText = "Cancelled";
          statusClass = "bg-danger";
          break;
    }
    return '&nbsp;<small class="badge ' + statusClass + '">' + statusText + '</small>';
  }

  getTotalInvoicesForStatus(status:number) {
    if(this.client.Invoices==null || this.client.Invoices.length<=0) return 0;
    var cnt = 0;
    this.client.Invoices.map((x:any)=>{
      return x.Status == status ? cnt++ : 0;
    });
    return cnt;
  }
}
</script>