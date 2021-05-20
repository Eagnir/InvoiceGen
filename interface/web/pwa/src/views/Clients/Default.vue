<template>
  <div class="client-panel">    
    <div class="flex-fill">
      <router-view />
    </div>

  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";

@Component({
  components: {},
})
export default class Clients extends Vue {
  
}
</script>