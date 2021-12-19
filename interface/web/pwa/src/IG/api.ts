
import { Server } from "@/IG/server";
import APIResponse, { APIResponseStatus } from "@/entity/response";
import { App as igApp, UserCredential } from "./app";
import Vue from 'vue'
import { SweetAlertIcon } from "sweetalert2";
import mainVueApp from "@/main";

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

  protected postCall<T>(url: string, data?: any, option?: APIOptions): Promise<APIResponse<T>> {
    try{mainVueApp.$isLoading(true);}catch(err){console.log(err);}
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
      //console.log(error);
      const instance = mainVueApp.$swal.toast.error("Fatal Error occured; '" + error + "'", {
        timer:8000
      });
      throw new Error(error);
    }).finally(()=>{
      try{setTimeout(()=>{mainVueApp.$ig.app.initBS();},0); mainVueApp.$isLoading(false);
      }catch(err){console.log(err);}
    }) as Promise<APIResponse<T>>;
  }
}