<template>
  <div>
    <div class="row align-items-center">
      <div class="col">
        <h1 class="m-0">Clients</h1>
      </div>
      <div class="col-12 col-md-5 text-end">
        <div class="input-group mt-2 mt-sm-0">
          <input
            type="text"
            class="form-control"
            id="autoSizingInputGroup"
            placeholder="name, GST # or email"
            v-model.trim="searchTerm"
            @keyup="search"
          />
          <div class="input-group-text c-pointer" @click="clearSearch"><i :class="{'bi-search':this.searchTerm=='','bi-x-circle':this.searchTerm!='',}"></i></div>
        </div>
      </div>
    </div>
    <hr />
    <div class="row">
      <router-link
        :to="{ name: 'ClientDetails', params: { clientId: 1, clientName:'Client Name Here' } }"
        custom
        v-slot="{ navigate }"
        v-for="client in clientResults" :key="client.ClientId"
      >
        <div class="col col-md-6 col-xl-4 c-pointer" @click="navigate">
          <div class="card mb-3 d-flex">
            <div class="d-flex">
              <div class="card-body w-100">
                <h5 class="card-title d-flex">
                  <span class="flex-grow-1 d-inline-block text-truncate" title="Abc Pvt. Ltd.">{{ client.Name }}</span>
                  <span class="card-text ms-3 text-end d-md-none text-nowrap" @click.stop="showClientOptions(client)"><i class="bi-globe2" title="Country"></i> {{client.Country}}
                    <i class="bi-three-dots-vertical"></i>
                  </span>
                  <div class="btn-group d-none d-md-inline-flex">
                    <span class="text-end fs-6" data-bs-toggle="dropdown" data-bs-auto-close="true"><i class="bi-globe2" title="Country"></i> {{client.Country}}
                      <i class="bi-three-dots-vertical"></i>
                    </span>
                    <ul class="dropdown-menu" aria-labelledby="defaultDropdown">
                      <li><h6 class="dropdown-header">Actions</h6></li>
                      <li><a class="dropdown-item" href="#"><i class="bi-pencil-square"></i> Edit</a></li>
                      <li><hr class="dropdown-divider"></li>
                      <li><a class="dropdown-item" href="#"><i class="bi bi-receipt-cutoff"></i> Raise Invoice</a></li>
                      <li><hr class="dropdown-divider"></li>
                      <li><a class="dropdown-item text-danger" href="#"><i class="bi bi-trash"></i> Delete</a></li>
                    </ul>
                  </div>
                </h5>
                <div class="d-flex">
                  <p class="card-text mb-1 flex-grow-1"><i class="bi-cash" title="Default Currency"></i> {{client.DefaultCurrency.ShortName}}</p>
                  <small class="card-text mb-1" @click.stop="copyGST(client.GSTNumber)" v-if="client.GSTNumber!=''"><i class="bi-hash" title="GST Number"></i>{{client.GSTNumber}}</small>
                </div>
                <p class="card-text mb-1" :title="client.InvoiceStats.TotalAmount"><i class="bi-calendar3" title="Current Financial Year"></i> <span v-html="$options.filters.igamount(client.InvoiceStats.TotalAmount / client.DefaultCurrency.Conversion, 'INR')"></span> <small v-if="client.DefaultCurrency.ShortName != $ig.app.User.company.defaultCurrency.shortName">(<span v-html="$options.filters.igamount(client.InvoiceStats.TotalAmount,  client.DefaultCurrency.ShortName)"></span>)</small></p>
                <p class="card-text m-0 d-flex">
                  <small class="text-muted flex-grow-1" v-if="client.InvoiceStats.LastInvoiceDate !=null" :title="client.InvoiceStats.LastInvoiceDate | moment('Do MMM yyyy, h:mm:ss a')">last invoice, <span>{{ client.InvoiceStats.LastInvoiceDate | moment("from") }}</span></small>
                  <span class="badge bg-warning ms-auto" :class="{'invisible':client.InvoiceStats.PendingInvoiceCount<=0}">{{ client.InvoiceStats.PendingInvoiceCount }}</span>
                </p>
              </div>
            </div>
          </div>
        </div>
      </router-link>

      
    </div>

    <!-- Modal -->
    <div class="modal fade" id="clientOptions" tabindex="-1">
      <div class="modal-dialog modal-xs modal-dialog-centered">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="exampleModalLabel">
              {{ selectedClient.Name }}
            </h5>
          </div>
          <div class="modal-body">
            <div class="d-flex flex-column">
              <button type="button" class="btn btn-info mb-2"><i class="bi bi-pencil-square"></i> Edit</button>
              <hr>
              <button type="button" class="btn btn-secondary mb-2"><i class="bi bi-receipt-cutoff"></i> Raise Invoice</button>
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
import { Vue } from "vue-property-decorator";
import Component from 'vue-class-component'

@Component({
  components: {},
})
export default class ClientList extends Vue {

  clientOptionsPopup: Modal;

  searchTerm:string = "";

  selectedClient: any = {};
  
  clientResults: Array<any> = [];

  clients:Array<any> = [];

  public mounted() {
    const el = document.getElementById("clientOptions");
    this.clientOptionsPopup = new Modal(el);
    this.refreshClients();
  }

  public refreshClients() {
    this.$ig.api.listOfCommpanyClients().then(resp => {
      this.clients = resp.Data;
      this.search();
    })
  }

  public search() {
    if(this.searchTerm=="" || this.searchTerm==null) {
      this.clientResults = this.clients;
      return;
    }

    var words = this.searchTerm.toLowerCase().split(" ");

    this.clientResults = [];
    this.clients.forEach(client => {
      words.forEach(word => {
        if(client.Name.toLowerCase().includes(word) ||
          client.GSTNumber.toLowerCase().includes(word) ||
          client.Email.toLowerCase().includes(word) ||
          client.Country.toLowerCase().includes(word) ||
          client.DefaultCurrency.ShortName.toLowerCase().includes(word) ||
          client.InvoiceStats.PendingInvoiceCount == parseInt(word))
          this.clientResults.push(client);
      });
    });
  }

  public clearSearch() {
    this.searchTerm = "";
    this.search();
  }

  public copyGST(gstNumber:string) {
    this.$ig.browser.copyToClipboard(gstNumber);
    this.$swal.toast.success("GST Number, copied to clipboard");
  }

  showClientOptions(client:any) {
    this.selectedClient = client;
    this.clientOptionsPopup.show();
  }

  //showInvoice() 
  //{
  //  console.log("hello");
 // }

}
</script>