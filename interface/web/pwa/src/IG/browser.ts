
export interface BrowserInterface {

    getOS():string;
    getCtrlKeyForOS(): {short:string, value:string};

}

export class Browser {

    static getOS():string {
        const userAgent = window.navigator.userAgent,
            platform = window.navigator.platform,
            macosPlatforms = ['Macintosh', 'MacIntel', 'MacPPC', 'Mac68K'],
            windowsPlatforms = ['Win32', 'Win64', 'Windows', 'WinCE'],
            iosPlatforms = ['iPhone', 'iPad', 'iPod'];
        let os = null;

        if (macosPlatforms.indexOf(platform) !== -1) {
            os = 'Mac OS';
        } else if (iosPlatforms.indexOf(platform) !== -1) {
            os = 'iOS';
        } else if (windowsPlatforms.indexOf(platform) !== -1) {
            os = 'Windows';
        } else if (/Android/.test(userAgent)) {
            os = 'Android';
        } else if (!os && /Linux/.test(platform)) {
            os = 'Linux';
        }
        return os;
    }
    
    static  getCtrlKeyForOS(): {short:string, value:string} {
        const keyMapping = {
            short: "",
            value:""
        };
        switch (this.getOS()) {
            case "Mac OS":
                keyMapping.short = "opt";
                keyMapping.value = "alt";
                break;
            case "Windows":
            case "Linux":
                keyMapping.short = "alt";
                keyMapping.value = "alt";
                break;
            default:
                keyMapping.short = "alt";
                keyMapping.value = "alt";
        }
        return keyMapping;
    }
}