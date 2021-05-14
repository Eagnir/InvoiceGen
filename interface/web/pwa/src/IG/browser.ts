
export class Browser {

    static getOS() {
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
    
    static  getCtrlKeyForOS(secondaryKey: string, showInTip: boolean = false) {
        let keyMapping = "";
        switch (this.getOS()) {
            case "Mac OS":
                keyMapping = "cmd + " + secondaryKey;
                break;
            case "Windows":
            case "Linux":
                keyMapping = "ctrl + " + secondaryKey;
                break;
            default:
                keyMapping = "ctrl + " + secondaryKey;
        }
        if (showInTip)
            return keyMapping;
        else
            return keyMapping.replaceAll("cmd", "command");
    }
}