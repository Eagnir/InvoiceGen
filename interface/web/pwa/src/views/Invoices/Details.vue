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
          <li><button class="dropdown-item"><i class="bi bi-card-checklist me-2"></i> Change Status</button></li>
          <li><button class="dropdown-item"><i class="bi bi-download me-2"></i> Download</button></li>
          <li><hr class="dropdown-divider"></li>
          <li><button class="dropdown-item text-danger"><i class="bi bi-trash me-2"></i> Delete</button></li>
        </ul>
      </div>
    </nav>
    <div class="row align-items-center">
      <div class="col-12 col-md text-md-start text-center">
        <h1 class="m-0">Invoice Details <span><span class="d-none d-md-inline">-</span><br class="d-inline d-md-none"></span> #{{ invoiceNumber }}</h1>
      </div>
      <div class="col-12 col-md-5 text-md-end d-none d-md-block">
        <nav class="navbar text-center justify-content-center align-items-center flex-nowrap">
          <button type="button" class="btn flex-fill btn-primary"><i class="bi bi-arrow-repeat"></i> Reset</button>
          <button type="button" class="btn flex-fill btn-success mx-3"><i class="bi bi-check2-square"></i> Save</button>
          <div class="btn-group dropdown flex-fill">
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
          <router-link :to="{ name: 'Invoices' }">Invoices</router-link>
        </li>
        <li class="breadcrumb-item active">Detail</li>
        <li class="breadcrumb-item active" aria-current="page">
          #{{ invoiceNumber }}
        </li>
      </ol>
    </nav>

    <div class="row">
      <div class="col-md mb-3 mb-md-0">
        <div class="form-floating">
          <input type="date" class="form-control" id="date">
          <label for="date">Invoice Date</label>
        </div>
        <div class="d-flex flex-row">
          <small class="flex-grow-1">Invoice #: 23052102 <a href="#" class="text-primary"><i class="bi bi-pencil-square"></i></a></small>
        </div>
      </div>
      <div class="col-md mb-3 mb-md-0">
        <div class="form-floating">
          <select class="form-select" id="client">
            <option selected>Select Client</option>
            <option value="1">Abc Pvt. Ltd.</option>
            <option value="2">John &amp; Sons Co.</option>
          </select>
          <label for="client">Client</label>
        </div>
        <div class="d-flex flex-row">
          <small class="flex-grow-1">GST #: (select a client) </small>
          <small><i class="bi bi-cash"></i> INR</small>
        </div>
      </div>
      <div class="col-md mb-3 mb-md-0">
        <div class="form-floating">
          <select class="form-select" id="taxGroup">
            <option selected>Select Applicable Tax</option>
            <option value="0">Not Applicable</option>
            <option value="1">Integrated GST</option>
            <option value="2">Goods &amp; Service Tax</option>
          </select>
          <label for="taxGroup">Tax Group</label>
        </div>
      </div>
      <div class="col-md">
        <div class="form-floating">
          <select class="form-select" id="invoiceCurrency">
            <option selected>Select Currency</option>
            <option value="1">INR</option>
            <option value="2">USD</option>
            <option value="3">CAD</option>
          </select>
          <label for="invoiceCurrency">Invoice Currency</label>
        </div>
      </div>
    </div>




    <div class="row mt-4">
      <div class="col">
        <h5>Invoice Paticulars</h5>
        <hr>
      </div>
    </div>


<div class="card">
  <div class="card-body py-0">
    <div class="row mt-3">
      <div class="col-12 col-md-9">
        <div class="input-group mb-3 form-floating">
          <button class="button btn-sm btn-danger"><i class="bi bi-trash"></i></button>
          <input type="text" class="form-control" id="particular" placeholder="Development services">
          <label for="particular" class="ps-5">Particular</label>
        </div>
      </div>
      <div class="col-12 col-md-3 mt-md-0">
          <div class="form-floating input-group mb-3">
            <span class="input-group-text">USD</span>
            <input type="text" class="form-control" v-maska="'#*.##'" v-model.trim="grossAmount" aria-label="Dollar amount (with dot and two decimal places)">
            <label class="ps-5 ms-4" for="amount">Gross Amount</label>
            <button class="btn btn-sm btn-success" type="button">Add Item</button>
          </div>
      </div>
    </div>
  </div>
</div>

    <div class="row">
      <div class="col">
        <table class="table table-sm table-striped table-hover">
          <thead>
            <tr>
              <th scope="col">#</th>
              <th scope="col">Particulars</th>
              <th class="text-end" style="min-width:100px;" scope="col">Total (INR)</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td class="text-center" colspan="3">Please add invoice particulars</td>
            </tr>
            <tr>
              <th scope="row">1</th>
              <td>Software Development for WPC</td>
              <td class="text-end align-middle" v-html="this.$options.filters.igamount(32000.86)"></td>
            </tr>
            <tr>
              <th scope="row">2</th>
              <td>IT Consultancy for 3 months</td>
              <td class="text-end align-middle" v-html="this.$options.filters.igamount(95200.08)"></td>
            </tr>
            <tr>
              <th scope="row">3</th>
              <td>1 Year software maintenance and development contract (1 Jan 2021 to 31 Dec 2021)</td>
              <td class="text-end align-middle" v-html="this.$options.filters.igamount(16520.12)"></td>
            </tr>
          </tbody>
          <tfoot>
            <tr>
              <td colspan="2" class="border-0 border-end pe-4 text-end fw-bold text-success border-end border-success">Gross Amount</td>
              <td class="text-end align-middle fw-bold border border-success text-success" v-html="this.$options.filters.igamount(120520)"></td>
            </tr>
            <tr>
              <td colspan="2" class="border-0 border-end pe-4 text-end">+ Tax</td>
              <td class="text-end align-middle" v-html="this.$options.filters.igamount(12000.56)"></td>
            </tr>
            <tr>
              <td colspan="2" class="border-0 border-end pe-4 text-end">+ Tax</td>
              <td class="text-end align-middle border-bottom border-success" v-html="this.$options.filters.igamount(12000)"></td>
            </tr>
            <tr>
              <td colspan="2" class="border-0 border-end pe-4 text-end fw-bold text-success border-end border-success">Invoice Amount</td>
              <td class="text-end align-middle fw-bold border border-success text-success" v-html="this.$options.filters.igamount(123000,'INR')"></td>
            </tr>
          </tfoot>
        </table>
      </div>
    </div>

  </div>
</template>

<script lang="ts">
import { Prop, Vue } from "vue-property-decorator";
import Component from 'vue-class-component'

Component.registerHooks([
  'beforeRouteEnter',
  'beforeRouteUpdate',
  'beforeRouteLeave'
])

@Component({
  components: {},
})
export default class InvoiceDetails extends Vue {
  @Prop({ default: "0" })
  invoiceId: number;

  @Prop({ default: "Loading..." })
  invoiceNumber: string;

  grossAmount:string = "";

  
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