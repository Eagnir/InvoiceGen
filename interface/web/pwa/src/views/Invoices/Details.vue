<template>
  <div>
    <div class="row align-items-center">
      <div class="col-12 col-md text-md-start text-center">
        <h1 class="m-0">Invoice Details <span><span class="d-none d-md-inline">-</span><br class="d-inline d-md-none"></span> #{{ invoiceNumber }}</h1>
      </div>
      <div class="col-12 col-md-5 text-md-end d-none d-md-block">
        <button type="button" class="btn btn-default mb-3 me-4">Reset</button>
        <button type="button" class="btn btn-success mb-3">Save Invoice</button>
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
          Invoice #{{ invoiceNumber }}
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
            <input type="number" class="form-control" aria-label="Dollar amount (with dot and two decimal places)">
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
              <th scope="row">1</th>
              <td>Software Development for WPC</td>
              <td class="text-end align-middle">32,000<small>.00</small></td>
            </tr>
            <tr>
              <th scope="row">2</th>
              <td>IT Consultancy for 3 months</td>
              <td class="text-end align-middle">32,000<small>.00</small></td>
            </tr>
            <tr>
              <th scope="row">3</th>
              <td>1 Year software maintenance and development contract (1 Jan 2021 to 31 Dec 2021)</td>
              <td class="text-end align-middle">32,000<small>.32</small></td>
            </tr>
          </tbody>
          <tfoot>
            <tr>
              <td colspan="2" class="border-0 border-end pe-4 text-end fw-bold text-success border-end border-success">Gross Amount</td>
              <td class="text-end align-middle fw-bold border border-success text-success">12,000<small>.56</small></td>
            </tr>
            <tr>
              <td colspan="2" class="border-0 border-end pe-4 text-end">+ Tax</td>
              <td class="text-end align-middle">12,000<small>.56</small></td>
            </tr>
            <tr>
              <td colspan="2" class="border-0 border-end pe-4 text-end">+ Tax</td>
              <td class="text-end align-middle border-bottom border-success">12,000<small>.02</small></td>
            </tr>
            <tr>
              <td colspan="2" class="border-0 border-end pe-4 text-end fw-bold text-success border-end border-success">Invoice Amount</td>
              <td class="text-end align-middle fw-bold border border-success text-success">1,23,000<small>.00</small></td>
            </tr>
          </tfoot>
        </table>
      </div>
    </div>


    <div class="row mt-4">
      <div class="col-12 text-center d-md-none">
        <button type="button" class="btn btn-default mb-3 me-4">Reset</button>
        <button type="button" class="btn btn-success mb-3">Save Invoice</button>
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