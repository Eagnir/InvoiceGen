import APIResponse, { APIResponseStatus } from "@/entity/response";
import { API, APIOptions } from "./api";
import { App as igApp, UserCredential } from "./app";
import { Server } from "./server";


class APIUrl {
    public static Auth_Credential: string = Server.Instance.APIBaseURL + "auth/credential";
    public static Auth_Heartbeat: string = Server.Instance.APIBaseURL + "auth/heartbeat";
    public static Auth_Invalidate: string = Server.Instance.APIBaseURL + "auth/invalidate";

    public static Client_List: string = Server.Instance.APIBaseURL + "client/list";
    public static Client_Details: string = Server.Instance.APIBaseURL + "client/detail";

}


export class IGAPI extends API {

    private static _apiInstance: IGAPI;

    public static get Instance() {
        return this._apiInstance || (this._apiInstance = new this());
    }

    public authCredential(email: string, passw: string, option?: APIOptions): Promise<APIResponse<UserCredential>> {
        option = Object.assign({}, { AuthRequired: false }, option);
        return this.postCall<UserCredential>(
            APIUrl.Auth_Credential,
            {
                Email: email,
                Password: passw
            }, option)
            .then(resp => {
                if (resp.Data != null) {
                    const data: UserCredential[] = [];
                    resp.Data?.forEach(item => {
                        data.push(Object.assign(new UserCredential(), item));
                    });
                    resp.Data = data;
                }
                return resp;
            })
    }

    public heartbeat(option?: APIOptions): Promise<APIResponse<UserCredential>> {
        if (igApp.Instance.User == null) return Promise.reject();
        option = Object.assign({}, { ShowToast: false }, option);
        return this.postCall<UserCredential>(APIUrl.Auth_Heartbeat, null, option)
            .then(resp => {
                if (resp.Status != APIResponseStatus.StatusSuccess) {
                    igApp.Instance.clearUser();
                    return resp;
                }
                if (resp.Data != null) {
                    const data: UserCredential[] = [];
                    resp.Data?.forEach(item => {
                        data.push(Object.assign(new UserCredential(), item));
                    });
                    resp.Data = data;
                }
                return resp;
            })
    }

    public signout(option?: APIOptions): Promise<APIResponse<UserCredential>> {
        option = Object.assign({}, option);
        return this.postCall<UserCredential>(APIUrl.Auth_Invalidate, null, option)
            .then(resp => {
                if (resp.Status == APIResponseStatus.StatusSuccess) {
                    igApp.Instance.clearUser();
                }
                return resp;
            })
    }

    public listOfCommpanyClients(option?: APIOptions): Promise<APIResponse<any>> {
        option = Object.assign({}, option);
        return this.postCall<any>(APIUrl.Client_List, null, option)
    }

    public getClientDetails(clientId:number, option?: APIOptions): Promise<APIResponse<any>> {
        option = Object.assign({}, option);
        return this.postCall<any>(APIUrl.Client_Details, {clientId: clientId}, option)
    }

}
