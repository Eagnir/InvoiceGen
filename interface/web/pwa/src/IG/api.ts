
import { Server } from "@/IG/server";
import APIResponse, { APIResponseStatus } from "@/entity/response";
import { App as igApp, UserCredential } from "./app";
import Vue from 'vue'
import { SweetAlertIcon } from "sweetalert2";
import mainVueApp from "@/main";

class APIUrl {
  public static Auth_Credential: string = Server.Instance.APIBaseURL + "auth/credential";
  public static Auth_Heartbeat: string = Server.Instance.APIBaseURL + "auth/heartbeat";
  public static Auth_Invalidate: string = Server.Instance.APIBaseURL + "auth/invalidate";

  public static Client_List: string = Server.Instance.APIBaseURL + "client/list";

}

export interface APIOptions {
  AuthRequired?: boolean,
  ShowToast?: boolean,
  ShowSuccessToast?: boolean,
  ShowWarningToast?: boolean,
  ShowFailureToast?: boolean,
  RejectOnFailure?: boolean,
  ClearPreviousToasts?: boolean
}

export class API {

  private static _instance: API;

  public static get Instance() {
    return this._instance || (this._instance = new this());
  }


  private defaultOption: APIOptions = {
    AuthRequired: true,
    ShowToast: true,
    ShowSuccessToast: true,
    ShowWarningToast: true,
    ShowFailureToast: true,
    RejectOnFailure: true,
    ClearPreviousToasts: true
  }

  private getHeaders(): string[][] {
    const headers: any = {};
    headers["content-type"] = "application/json";
    return headers;
  }

  private getAuthHeaders(): string[][] {
    const authHeaders: any = Object.assign({}, this.getHeaders());
    if(igApp.Instance.User != null){
      authHeaders["token"] = igApp.Instance.User.token;
      authHeaders["email"] = igApp.Instance.User.email;
    }
    return authHeaders;
  }

  private postCall<T>(url: string, data?: any, option?: APIOptions): Promise<APIResponse<T>> {
    option = Object.assign({}, this.defaultOption, option);
    const h: any = option.AuthRequired ? this.getAuthHeaders() : this.getHeaders();
    return fetch(url, {
      method: "POST",
      headers: h,
      body: data != null ? JSON.stringify(data) : null
    }).then(res => res.json()).then(res => {
      const resp = Object.assign(new APIResponse<T>(), res) as APIResponse<T>;

      if (option.ShowToast) {
        let showToast = true;
        let toastType:SweetAlertIcon = "info";
        let toastDuration = 3000;
        switch (resp.Status) {
          case APIResponseStatus.StatusFailure:
            showToast = option.ShowFailureToast;
            toastType = "error";
            toastDuration = 0;
            break;
          case APIResponseStatus.StatusFatalError:
            showToast = option.ShowFailureToast;
            toastType = "error";
            toastDuration = 0;
            break;
          case APIResponseStatus.StatusWarning:
            showToast = option.ShowWarningToast;
            toastType = "warning";
            break;
          case APIResponseStatus.StatusSuccess:
            showToast = option.ShowSuccessToast;
            toastType = "success";
            break;
          default:
            toastType = "info";
        }
        if (showToast) {

          
          //if (option.ClearPreviousToasts)
            //Vue.prototype.$swal.toast.fire()
            //Vue.$toast.clear();
          if (resp.Message != "") {
              const instance = mainVueApp.$swal.toast.fire(resp.Message, {
                icon: toastType,
              });
            }
        }
      }

      // Handle session expiry errors here

      if (resp.Status == APIResponseStatus.StatusSuccess)
        return resp;
      else if (option.RejectOnFailure)
        Promise.reject();
      else
        return resp;
    }).catch(error => {
      console.log(error);
      const instance = mainVueApp.$swal.toast.error("Fatal Error occured; '" + error + "'", {
        timer:8000
      });
      console.log(error);
    }) as Promise<APIResponse<T>>;
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

}