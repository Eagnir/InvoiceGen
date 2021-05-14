declare global {
  interface Window {
    igconfig: IGConfig
    srv:Server
  }
}

type IGConfig = {
  apiSubdomain: string;
  protocol: string;
  host: string;
  port: string;
}

export class Server {
  private static _instance: Server;

    public static get Instance() {
      if (this._instance == null)
        this._instance = new this();
      this._instance.config = window.igconfig;
      return this._instance
    }

  public config: IGConfig = window.igconfig;

  public get APIBaseURL():string {
    return this.config?.protocol + "://" + this.config?.apiSubdomain + "." + this.config?.host + ":" + this.config?.port + "/";
  }

}