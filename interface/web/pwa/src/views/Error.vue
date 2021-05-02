<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png" />
    <h1>404 Resource Not Found</h1>
    <p>{{ this.$route.query.msg }}</p>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component
export default class Error extends Vue {
  
}
</script>

