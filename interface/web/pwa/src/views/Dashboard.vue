<template>
  <div>
    <h1>Dashboard</h1>
    <hr />

    <div class="accordion" id="accordionExample">
      <div class="accordion-item">
        <h2 class="accordion-header" id="headingOne">
          <button
            class="accordion-button"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#collapseOne"
            aria-expanded="true"
            aria-controls="collapseOne"
          >
            Financial Year (2020 - 2021)
          </button>
        </h2>
        <div
          id="collapseOne"
          class="accordion-collapse collapse show"
          aria-labelledby="headingOne"
        >
          <div class="accordion-body">
            <div class="row">
              <div class="col-md-3 col-12">
                <div class="card text-white bg-success mb-3">
                  <div class="card-header">Yearly Collection</div>
                  <div class="card-body">
                    <h2 class="card-title"><small>gross</small>INR 20,00,350/-</h2>
                    <div class="row card-stats">
                      <div class="col">
                        <small>+Tax</small><br />
                        <small>INR 15,250.00</small>
                      </div>
                      <div class="col text-end">
                        <small>Total Inwards</small><br />
                        <small>INR 15,250.00</small>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <div class="col-md-3 col-12">
                <div class="card text-white bg-success mb-3">
                  <div class="card-header">This Month (May 2021)</div>
                  <div class="card-body">
                    <h2 class="card-title"><small>gross</small>INR 50,264/-</h2>
                    <div class="row card-stats">
                      <div class="col">
                        <small>+Tax</small><br />
                        <small>INR 15,250.00</small>
                      </div>
                      <div class="col text-end">
                        <small>Total Inwards</small><br />
                        <small>INR 15,250.00</small>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <div class="col-md-3 col-12">
                <div class="card text-dark bg-warning mb-3">
                  <div class="card-header">Pending Collection</div>
                  <div class="card-body">
                    <h2 class="card-title"><small>gross</small>INR 50,350/-</h2>
                    <div class="row card-stats">
                      <div class="col">
                        <small>+Tax</small><br />
                        <small>INR 15,250.00</small>
                      </div>
                      <div class="col text-end">
                        <small>Total Inwards</small><br />
                        <small>INR 15,250.00</small>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <div class="col-md-3 col-12">
                <div class="card text-dark bg-info mb-3">
                  <div class="card-header">Gross Income</div>
                  <div class="card-body">
                    <h2 class="card-title"><small>gross</small>INR 1,25,019/-</h2>
                    <div class="row card-stats">
                      <div class="col">
                        <small>+Tax</small><br />
                        <small>INR 15,250.00</small>
                      </div>
                      <div class="col text-end">
                        <small>Total Inwards</small><br />
                        <small>INR 15,250.00</small>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="accordion-item">
        <h2 class="accordion-header" id="headingTwo">
          <button
            class="accordion-button collapsed"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#collapseTwo"
            aria-expanded="false"
            aria-controls="collapseTwo"
          >
            Last Financial Year (2019 - 2020)
          </button>
        </h2>
        <div
          id="collapseTwo"
          class="accordion-collapse collapse"
          aria-labelledby="headingTwo"
        >
          <div class="accordion-body">
            <div class="row">
              

              <div class="col-md-3 col-12">
                <div class="card bg-default mb-3">
                  <div class="card-header">Yearly Collection</div>
                  <div class="card-body">
                    <h2 class="card-title"><small>gross</small>INR 20,00,350/-</h2>
                    <div class="row card-stats">
                      <div class="col">
                        <small>+Tax</small><br />
                        <small>INR 15,250.00</small>
                      </div>
                      <div class="col text-end">
                        <small>Total Inwards</small><br />
                        <small>INR 15,250.00</small>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <div class="col-md-3 col-12">
                <div class="card bg-default mb-3">
                  <div class="card-header">This Month (May 2021)</div>
                  <div class="card-body">
                    <h2 class="card-title"><small>gross</small>INR 50,264/-</h2>
                    <div class="row card-stats">
                      <div class="col">
                        <small>+Tax</small><br />
                        <small>INR 15,250.00</small>
                      </div>
                      <div class="col text-end">
                        <small>Total Inwards</small><br />
                        <small>INR 15,250.00</small>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <div class="col-md-3 col-12">
                <div class="card bg-default mb-3">
                  <div class="card-header">Pending Collection</div>
                  <div class="card-body">
                    <h2 class="card-title"><small>gross</small>INR 50,350/-</h2>
                    <div class="row card-stats">
                      <div class="col">
                        <small>+Tax</small><br />
                        <small>INR 15,250.00</small>
                      </div>
                      <div class="col text-end">
                        <small>Total Inwards</small><br />
                        <small>INR 15,250.00</small>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <div class="col-md-3 col-12">
                <div class="card bg-default mb-3">
                  <div class="card-header">Gross Income</div>
                  <div class="card-body">
                    <h2 class="card-title"><small>gross</small>INR 1,25,019/-</h2>
                    <div class="row card-stats">
                      <div class="col">
                        <small>+Tax</small><br />
                        <small>INR 15,250.00</small>
                      </div>
                      <div class="col text-end">
                        <small>Total Inwards</small><br />
                        <small>INR 15,250.00</small>
                      </div>
                    </div>
                  </div>
                </div>
              </div>


            </div>
          </div>
        </div>
      </div>
      <div class="accordion-item">
        <h2 class="accordion-header" id="headingThree">
          <button
            class="accordion-button"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#collapseThree"
            aria-expanded="false"
            aria-controls="collapseThree"
          >
            Lifetime Statistics
          </button>
        </h2>
        <div
          id="collapseThree"
          class="accordion-collapse collapse show"
          aria-labelledby="headingThree"
        >
          <div class="accordion-body">
            <div class="row">
              <div class="col-12 col-md-4 mb-3">
                <div class="card">
                  <div class="card-header">Last 10 invoices</div>
                  <ul class="list-group list-group-flush">
                    <li class="list-group-item">2 days ago</li>
                    <li class="list-group-item">15 days ago</li>
                    <li class="list-group-item">18 days ago</li>
                  </ul>
                </div>
              </div>
              <div class="col-12 col-md-4 mb-3">
                <div class="card">
                  <div class="card-header">Client statistics</div>
                  <ul class="list-group list-group-flush">
                    <li class="list-group-item">
                      Abc Inc. <small>avg. 15 days</small>
                      <span class="float-end">Rs. 4,03,200</span>
                    </li>
                    <li class="list-group-item">Xyz Ltd.</li>
                    <li class="list-group-item">John &amp; Sons Pvt. Ltd.</li>
                  </ul>
                </div>
              </div>
              <div class="col-12 col-md-4 mb-3">
                <div class="card">
                  <div class="card-header">Last 5 years</div>
                  <ul class="list-group list-group-flush">
                    <li class="list-group-item">2020 - 2021</li>
                    <li class="list-group-item">2019 - 2020</li>
                    <li class="list-group-item">2018 - 2019</li>
                  </ul>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="accordion-item">
        <h2 class="accordion-header" id="chartAccor">
          <button
            class="accordion-button"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#chartAccorContent"
            aria-expanded="false"
            aria-controls="chartAccorContent"
          >
            Charts &amp; Graphs
          </button>
        </h2>
        <div
          id="chartAccorContent"
          class="accordion-collapse collapse show"
          aria-labelledby="chartAccor"
        >
          <div class="accordion-body">
            <div class="row">
              <div class="col-12 col-md-6 mb-3">
                <LineChart :ChartData="this.chartData"></LineChart>
              </div>
              <div class="col-12 col-md-6 mb-3">
                <LineChart :ChartData="this.chartData"></LineChart>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import LineChart from "../components/Charts/LineChart";

@Component({
  components: {
    LineChart,
  },
})
export default class Dashboard extends Vue {
  chartData = {
    labels: [
      "January",
      "Feburary",
      "March",
      "April",
      "May",
      "June",
      "July",
      "August",
      "September",
      "October",
      "November",
      "December",
    ],
    datasets: [
      {
        label: "Income",
        data: [65, 59, 80, 81, 56, 55, 40, 81, 56, 55, 40, 20],
        fill: false,
        borderColor: "rgb(75, 192, 192)",
        tension: 0.1,
      },
      {
        label: "Average",
        data: [],
        fill: true,
        borderColor: "rgb(86, 86, 86)",
        tension: 0.1,
      },
    ],
  };

  beforeMount() {
    //var avgs = [];
    var sum = 0;
    this.chartData.datasets[0].data.forEach((element) => {
      sum += element;
    });

    console.log(Math.round(sum / this.chartData.datasets[0].data.length));
    this.chartData.datasets[0].data.forEach((element) => {
      this.chartData.datasets[1].data.push(
        Math.round(sum / this.chartData.datasets[0].data.length)
      );
    });
  }
}
</script>