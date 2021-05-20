<template>
  <h1>Invoice Details - {{ invoiceId }}</h1>
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