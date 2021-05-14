import { Browser as igBrowser } from "@/IG/browser";

export interface KeymapInterface {
    CreateInvoice(func: any): any;
    SignOut(func: any): any;
    
    NavDashboard(func: any): any;
    NavInvoices(func: any): any;
    NavClients(func: any): any;
    NavReports(func: any): any;
    NavSettings(func: any): any;
}

class KeymapObject {
    [Key: string]: any
}

export class Keymap {

    static ControlKey: string = igBrowser.getCtrlKeyForOS().value;
    static ControlKeyText: string = igBrowser.getCtrlKeyForOS().short;

    // Keymapping
    static CreateInvoice_Key: string = "n";
    static SignOut_Key: string = "o";

    static Nav_Dashboard_Key: string = "d";
    static Nav_Invoices_Key: string = "i";
    static Nav_Clients_Key: string = "c";
    static Nav_Reports_Key: string = "r";
    static Nav_Setting_Key: string = "s";


    static CreateInvoice(func: any): any {
        const key:string = this.CreateInvoice_Key;
        if (func == null)
            return this.ControlKeyText + " + " + key;
        const map = new KeymapObject()
        map[this.ControlKey + " + " + key] = func;
        return map;
    }
    static SignOut(func: any): any {
        const key:string = this.SignOut_Key;
        if (func == null)
            return this.ControlKeyText + " + " + key;
        const map = new KeymapObject()
        map[this.ControlKey + " + " + key] = func;
        return map;
    }
    
    static NavDashboard(func: any): any {
        const key:string = this.Nav_Dashboard_Key;
        if (func == null)
            return this.ControlKeyText + " + " + key;
        const map = new KeymapObject()
        map[this.ControlKey + " + " + key] = func;
        return map;
    }
    static NavInvoices(func: any): any {
        const key:string = this.Nav_Invoices_Key;
        if (func == null)
            return this.ControlKeyText + " + " + key;
        const map = new KeymapObject()
        map[this.ControlKey + " + " + key] = func;
        return map;
    }
    static NavClients(func: any): any {
        const key:string = this.Nav_Clients_Key;
        if (func == null)
            return this.ControlKeyText + " + " + key;
        const map = new KeymapObject()
        map[this.ControlKey + " + " + key] = func;
        return map;
    }
    static NavReports(func: any): any {
        const key:string = this.Nav_Reports_Key;
        if (func == null)
            return this.ControlKeyText + " + " + key;
        const map = new KeymapObject()
        map[this.ControlKey + " + " + key] = func;
        return map;
    }
    static NavSettings(func: any): any {
        const key:string = this.Nav_Setting_Key;
        if (func == null)
            return this.ControlKeyText + " + " + key;
        const map = new KeymapObject()
        map[this.ControlKey + " + " + key] = func;
        return map;
    }

}