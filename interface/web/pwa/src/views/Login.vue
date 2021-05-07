<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="text-center">InvoiceGen</h1>
        <p class="text-center">{{"please login to continue" | capitalize}}</p>
{{ new Date() | moment("dddd, MMMM Do YYYY") }}
        {{ msg }}
      </div>
    </div>
  </div>
</template>


<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component({
  components: {
    
  },
})
export default class Home extends Vue {

    msg = "";


}
</script>
