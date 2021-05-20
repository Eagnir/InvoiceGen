import _Vue, { PluginFunction } from "vue";
import Swal, { SweetAlertIcon, SweetAlertOptions, SweetAlertResult } from 'sweetalert2'

type Awaited<T> = T extends Promise<infer U> ? U : T;

type ToastType = {
    fire: (text: string, options?: SweetAlertOptions) => Promise<SweetAlertResult>;
    info: (text: string, options?: SweetAlertOptions) => Promise<SweetAlertResult>;
    error: (text: string, options?: SweetAlertOptions) => Promise<SweetAlertResult>;
    warning: (text: string, options?: SweetAlertOptions) => Promise<SweetAlertResult>;
    success: (text: string, options?: SweetAlertOptions) => Promise<SweetAlertResult>;
}

type SwalType = {
    fire: (text: string, options?: SweetAlertOptions) => Promise<SweetAlertResult>;
    alert: (text: string, options?: SweetAlertOptions) => Promise<SweetAlertResult>;
    confirm: (text: string, options?: SweetAlertOptions) => Promise<SweetAlertResult>;
    toast: ToastType;
}

export default {
    install(Vue: typeof _Vue, opt?: any) {
        const ToastDefOptions: SweetAlertOptions = {
            position: "top-end",
            timerProgressBar: true,
            timer: 3000,
            showConfirmButton: false,
            iconColor: "white",
            customClass: {
                popup: 'colored-toast'
            }
        };
        const PopupDefOptions: SweetAlertOptions = {
            buttonsStyling: false,
            customClass: {
                confirmButton: 'btn btn-primary me-3',
                cancelButton: 'btn btn-default',
              }
        };

        Vue.prototype.$swal = {
            fire: (text: string, options?: SweetAlertOptions): Promise<SweetAlertResult> => {
                return Swal.fire(Object.assign({}, PopupDefOptions, options, { toast: false, text: text }))
            },
            alert: (text: string, options?: SweetAlertOptions): Promise<SweetAlertResult> => {
                const def: SweetAlertOptions = {
                    text: text,
                    showConfirmButton: true,
                    showCloseButton: true,
                    showCancelButton: false,
                }
                return Swal.fire(Object.assign({}, PopupDefOptions, {
                    confirmButtonText: "Ok",
                }, options, def))
            },
            confirm: (text: string, options?: SweetAlertOptions): Promise<SweetAlertResult> => {
                const def: SweetAlertOptions = {
                    text: text,
                    showConfirmButton: true,
                    showCloseButton: true,
                    showCancelButton: true,
                }
                return Swal.fire(Object.assign({}, PopupDefOptions, {
                    confirmButtonText: "Confirm",
                    cancelButtonText: "Cancel"
                }, options, def));
            },
            toast: {
                fire: (text: string, options?: SweetAlertOptions): Promise<SweetAlertResult> => {
                    console.log(options);
                    return Swal.fire(Object.assign({}, ToastDefOptions, options, { toast: true, text: text }))
                },
                info: (text: string, options?: SweetAlertOptions): Promise<SweetAlertResult> => {
                    return Swal.fire(Object.assign({}, ToastDefOptions, options, { toast: true, icon: 'info', text: text }))
                },
                error: (text: string, options?: SweetAlertOptions): Promise<SweetAlertResult> => {
                    return Swal.fire(Object.assign({}, ToastDefOptions, options, { toast: true, icon: 'error', text: text }))
                },
                warning: (text: string, options?: SweetAlertOptions): Promise<SweetAlertResult> => {
                    return Swal.fire(Object.assign({}, ToastDefOptions, options, { toast: true, icon: 'warning', text: text }))
                },
                success: (text: string, options?: SweetAlertOptions): Promise<SweetAlertResult> => {
                    return Swal.fire(Object.assign({}, ToastDefOptions, options, { toast: true, icon: 'success', text: text }))
                }
            }
        };
    }
}

declare module 'vue/types/vue' {
    interface VueConstructor {
        $swal: SwalType
    }

    interface Vue {
        $swal: SwalType
    }
}