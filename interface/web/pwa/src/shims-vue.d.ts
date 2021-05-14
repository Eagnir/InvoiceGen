import { App } from "@/IG/App";

declare module "*.vue" {
  import Vue from "vue";
  export default Vue;
}

declare module 'v-hotkey' {
  import { DirectiveOptions, PluginFunction } from 'vue';

  type Plugin = {
    install: PluginFunction<{ [alias in string]?: number }>;
    directive: DirectiveOptions;
  };

  const plugin: Plugin;

  export default plugin;
}


declare module 'vue/types/vue' {
  interface Vue {
    $igApp: App
  }
}