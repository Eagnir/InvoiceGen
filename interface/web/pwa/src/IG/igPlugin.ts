import _Vue from "vue";
import { API as igApi } from "@/IG/api";
import { Server as igServer } from "@/IG/server";
import { Browser as igBrowser, BrowserInterface } from "@/IG/browser";
import { Keymap as igKeymap, KeymapInterface } from "@/IG/keymap";
import { App as igApp } from "@/IG/app";


export default {
    install(Vue: typeof _Vue, opt?: any) {
        Vue.prototype.$ig = {
            app: igApp.Instance,
            api: igApi.Instance,
            server: igServer.Instance,
            browser: igBrowser,
            keymap: igKeymap,
        }

    }
}

declare module 'vue/types/vue' {
    interface Vue {
        $ig: {
            app: igApp;
            api: igApi;
            server: igServer;
            browser: BrowserInterface;
            keymap: KeymapInterface;
        }
    }
}