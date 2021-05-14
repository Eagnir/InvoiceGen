
import { Server } from "@/IG/server";
import APIResponse, { APIResponseStatus } from "@/entity/response";
import { App as igApp, UserCredential } from "./app";
import Vue from 'vue'

class APIUrl {
  public static Auth_Credential: string = Server.Instance.APIBaseURL + "auth/credential";
  public static Auth_Heartbeat: string = Server.Instance.APIBaseURL + "auth/heartbeat";
  public static Auth_Invalidate: string = Server.Instance.APIBaseURL + "auth/invalidate";

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

  private static defaultOption: APIOptions = {
    AuthRequired: true,
    ShowToast: true,
    ShowSuccessToast: true,
    ShowWarningToast: true,
    ShowFailureToast: true,
    RejectOnFailure: true,
    ClearPreviousToasts: true
  }

  private static getHeaders(): string[][] {
    const headers: any = {};
    headers["content-type"] = "application/json";
    return headers;
  }

  private static getAuthHeaders(): string[][] {
    const authHeaders: any = Object.assign({}, this.getHeaders());
    authHeaders["token"] = igApp.Instance.User.token;
    authHeaders["email"] = igApp.Instance.User.email;
    return authHeaders;
  }

  private static postCall<T>(url: string, data?: any, option?: APIOptions): Promise<APIResponse<T>> {
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
        let toastType = "";
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
          if (option.ClearPreviousToasts)
            Vue.$toast.clear();
          const instance = Vue.$toast.open({
            type: toastType,
            message: resp.Message,
            duration: toastDuration
          });
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
      const instance = Vue.$toast.open({
        type: "error",
        message: "Fatal Error occured; '" + error + "'",
        duration: 0
      });
      console.log(error);
    }) as Promise<APIResponse<T>>;
  }

  public static authCredential(email: string, passw: string, option?:APIOptions): Promise<APIResponse<UserCredential>> {
    option = Object.assign({}, {AuthRequired:false}, option);
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

  public static heartbeat(option?:APIOptions): Promise<APIResponse<UserCredential>> {
    option = Object.assign({}, {ShowToast:false}, option);
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

  public static signout(option?:APIOptions): Promise<APIResponse<UserCredential>> {
    option = Object.assign({}, option);
    return this.postCall<UserCredential>(APIUrl.Auth_Invalidate)
      .then(resp => {
        if (resp.Status == APIResponseStatus.StatusSuccess) {
          igApp.Instance.clearUser();
        }
        return resp;
      })
  }

}