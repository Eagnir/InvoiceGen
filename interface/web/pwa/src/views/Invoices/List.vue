<template>
  <div>
    <div class="row align-items-center">
      <div class="col">
        <h1 class="m-0">Invoices</h1>
      </div>
      <div class="col-12 col-md-5 text-end">
        <div class="input-group mt-2 mt-sm-0">
          <input
            type="text"
            class="form-control"
            id="autoSizingInputGroup"
            placeholder="invoice #, client name, GST # or date"
          />
          <div class="input-group-text"><i class="bi-search"></i></div>
        </div>
      </div>
    </div>
    <hr />

    <div class="row mb-3">
      <div class="col-auto pe-0">Filter:</div>
      <div class="col">
        <span class="badge bg-primary me-2">All</span>
        <span class="badge bg-warning me-2">Pending</span>
        <span class="badge bg-info me-2">Parital Paid</span>
        <span class="badge bg-success me-2">Paid</span>
      </div>
    </div>

    <div class="accordion" id="accordionPanelsStayOpenExample">
      <div class="accordion-item">
        <h2 class="accordion-header" id="panelsStayOpen-headingOne">
          <button
            class="accordion-button"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#panelsStayOpen-collapseOne"
            aria-expanded="true"
            aria-controls="panelsStayOpen-collapseOne"
          >
            Pending Invoices
          </button>
        </h2>
        <div
          id="panelsStayOpen-collapseOne"
          data-bs-parent="#accordionPanelsStayOpenExample"
          class="accordion-collapse collapse show"
          aria-labelledby="panelsStayOpen-headingOne"
        >
          <div class="accordion-body">
            

            <div class="row">
      <router-link
        :to="{
          name: 'InvoiceDetails',
          params: { invoiceNumber: '22052101' },
        }"
        custom
        v-slot="{ navigate }"
      >
        <div class="col col-md-6 col-xl-3 c-pointer" @click="navigate">
          <div class="card mb-3 d-flex bg-warning text-dark">
            <h5 class="card-header d-flex pe-0">
              <span class="flex-grow-1">#22052101</span>
              <span class="text-end fs-6 d-md-none" @click.stop="showInvoiceOptions(1)">22 May 2021
                <i class="bi-three-dots-vertical pe-2"></i>
              </span>
              <div class="btn-group d-none d-md-inline-flex">
                <span class="text-end fs-6" data-bs-toggle="dropdown" data-bs-auto-close="true">22 May 2021
                  <i class="bi-three-dots-vertical pe-2"></i>
                </span>
                <ul class="dropdown-menu" aria-labelledby="defaultDropdown">
                   <li><h6 class="dropdown-header">Actions</h6></li>
                  <li><a class="dropdown-item" href="#"><i class="bi-pencil-square"></i> Edit</a></li>
                  <li><hr class="dropdown-divider"></li>
                  <li><a class="dropdown-item" href="#"><i class="bi bi-file-earmark-arrow-down"></i> Download</a></li>
                  <li><hr class="dropdown-divider"></li>
                  <li><a class="dropdown-item text-danger" href="#"><i class="bi bi-trash"></i> Delete</a></li>
                </ul>
              </div>
            </h5>
            <div class="d-flex">
              <div class="card-body">
                <p class="card-text mb-1">
                  <i class="bi-people" title="Default Currency"></i> Client Name
                  Here
                </p>
                <hr class="my-1 mx-0" />

                <p class="card-text mb-1 d-flex" title="USD 4,526/-">
                  <small
                    class="bi-list-check me-1 text-black-50"
                    title="Current Financial Year"
                  ></small>
                  <small class="flex-grow-1 text-black-50"
                    >Product/Service</small
                  >
                  <small class="text-black-50">10,20,000</small>
                </p>
                <p class="card-text mb-1 d-flex" title="USD 4,526/-">
                  <small
                    class="bi-plus-square me-1 text-black-50"
                    title="Current Financial Year"
                  ></small>
                  <small class="flex-grow-1 text-black-50">Tax (18%)</small>
                  <small class="text-black-50">10,2000</small>
                </p>
              </div>
            </div>
            <div class="card-footer">
              <p class="card-text mb-1 d-flex fw-bold" title="USD 4,526/-">
                <small class="flex-grow-1">Invoice Total (USD)</small>
                <small class="">20,30,2000</small>
              </p>
            </div>
          </div>
        </div>
      </router-link>

      
    </div>



          </div>
        </div>
      </div>
      <div class="accordion-item">
        <h2 class="accordion-header" id="panelsStayOpen-headingTwo">
          <button
            class="accordion-button collapsed"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#panelsStayOpen-collapseTwo"
            aria-expanded="false"
            aria-controls="panelsStayOpen-collapseTwo"
          >
            May 2021
          </button>
        </h2>
        <div
          id="panelsStayOpen-collapseTwo"
          data-bs-parent="#accordionPanelsStayOpenExample"
          class="accordion-collapse collapse"
          aria-labelledby="panelsStayOpen-headingTwo"
        >
          <div class="accordion-body">Current Month</div>
        </div>
      </div>
      <div class="accordion-item">
        <h2 class="accordion-header" id="panelsStayOpen-headingThree">
          <button
            class="accordion-button collapsed"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#panelsStayOpen-collapseThree"
            aria-expanded="false"
            aria-controls="panelsStayOpen-collapseThree"
          >
            April 2021
          </button>
        </h2>
        <div
          id="panelsStayOpen-collapseThree"
          data-bs-parent="#accordionPanelsStayOpenExample"
          class="accordion-collapse collapse"
          aria-labelledby="panelsStayOpen-headingThree"
        >
          <div class="accordion-body">Previous Months</div>
        </div>
      </div>
    </div>

    
    <!-- Modal -->
    <div class="modal fade" id="invoiceOptions" tabindex="-1">
      <div class="modal-dialog modal-xs modal-dialog-centered">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="exampleModalLabel">
              Invoice #22052101
            </h5>
          </div>
          <div class="modal-body">
            <div class="d-flex flex-column">
              <button type="button" class="btn btn-info mb-2"><i class="bi bi-pencil-square"></i> Edit</button>
              <hr>
              <button type="button" class="btn btn-secondary mb-2"><i class="bi bi-file-earmark-arrow-down"></i> Download</button>
              <hr>
              <button type="button" class="btn btn-danger mb-2"><i class="bi bi-trash"></i> Delete</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Modal } from 'bootstrap';
import Vue from 'vue'
import Component from 'vue-class-component';

@Component
export default class InvoiceList extends Vue {

  invoiceOptionsPopup: Modal;

  public mounted() {
    const el = document.getElementById("invoiceOptions");
    this.invoiceOptionsPopup = new Modal(el);
  }

  showInvoiceOptions(invoiceId:number) {
    this.invoiceOptionsPopup.show();
  }

  showInvoice() 
  {
    console.log("hello");
  }

}
</script>