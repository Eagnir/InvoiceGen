import { Browser as igBrowser } from "@/IG/browser";

export class Keymap {

    static CreateInvoice(func: any) {
        const key = igBrowser.getCtrlKeyForOS("o");
        const obj: any = {};
        obj[key] = func;
        return obj;
    }

}