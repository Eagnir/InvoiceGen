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

declare module 'vuejs-loading-screen' {

export default {
  install(Vue: typeof _Vue, opt?: any) }
  declare module 'vue/types/vue' {
      interface Vue {
        $isLoading: (isLoading:boolean) => void
    }
  }
};
