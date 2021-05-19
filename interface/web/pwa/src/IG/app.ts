import { Popover, Tooltip } from "bootstrap";

export class App {
    private static _instance: App;

    public static get Instance() {
        return this._instance || (this._instance = new this());
    }


    public static storageKey_User: string = "usr";

    private _user: UserCredential = null;
    get User(): UserCredential {
        if (this._user == null) {
            const usr = new UserCredential()
            if (usr.load()) {
                this._user = usr;
            }
        }
        return this._user;
    }
    set User(value: UserCredential) {
        if (value == null) {
            this._user = value;
            return;
        }
        if (value?.save()) {
            this._user = value;
        }
        else {
            console.log("Error setting UserCredential")
        }
    }


    clearUser(): boolean {
        try {
            localStorage.removeItem(App.storageKey_User);
            this.User = null;
            return true;
        }
        catch (ex) {
            console.log(ex);
        }
        return false;
    }

    initBS() {
        this.initBSTooltips();
        this.initBSPopovers();
    }

    initBSTooltips() {
        //Tooltips
        const tooltipTriggerList = [].slice.call(
            document.querySelectorAll('[data-bs-toggle="tooltip"]')
        );
        const tooltipList = tooltipTriggerList.map(function (tooltipTriggerEl) {
            let inst = Tooltip.getInstance(tooltipTriggerEl);
            if (inst == null) {
                inst = new Tooltip(tooltipTriggerEl, {
                    delay: { show: 100, hide: 10 }, trigger: "hover"
                });
            }
            return inst;
        });
    }

    initBSPopovers() {
        //Popovers
        const popoverTriggerList = [].slice.call(
            document.querySelectorAll('[data-bs-toggle="popover"]')
        );
        const popoverList = popoverTriggerList.map(function (popoverTriggerEl) {
            let inst = Popover.getInstance(popoverTriggerEl);
            if (inst == null) {
                inst = new Popover(popoverTriggerEl);
            }
            return inst;
        });

    }
}

export class UserCredential implements IUserCredential {
    ucid: number = -1;
    name: string = "";
    email: string = "";
    token: string = "";
    company: Company = null;
    load(): boolean {
        try {
            const usrJson = localStorage.getItem(App.storageKey_User);
            if (usrJson != null) {
                const x = JSON.parse(usrJson) as UserCredential;
                Object.assign(this, x);
                return true;
            }
            else {
                return false;
            }
        }
        catch (ex) {
            console.log(ex);
        }
        return false;
    }

    save(): boolean {
        try {
            localStorage.removeItem(App.storageKey_User);
            if (this != null) {
                localStorage.setItem(App.storageKey_User, JSON.stringify(this));
                return true;
            }
            return false;
        }
        catch (ex) {
            console.log(ex);
        }
        return false;
    }

}

class Company implements ICompany {
    cid: number = -1;
    name: string = "";
    tel: string = "";

}

interface IUserCredential {
    ucid: number;
    name: string;
    email: string;
    token: string;

    company?: Company;

    load(): boolean;
    save(): boolean;
}


interface ICompany {
    cid: number;
    name: string;
    tel: string;
}


