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
            placeholder="name, GST # or email"
          />
          <div class="input-group-text"><i class="bi-search"></i></div>
        </div>
      </div>
    </div>
    <hr />
    <div class="row">
      <router-link
        :to="{ name: 'ClientDetails', params: { id: 1 } }"
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