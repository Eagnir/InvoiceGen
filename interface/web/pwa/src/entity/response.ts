export default class APIResponse<Type> {

    Command: string = "";
    Pagination?: APIPagination;
    Data?:Type[];
    Status?: APIResponseStatus;
    StatusText: string = "";
    Message: string = "";
    Error: any;
    DateTime?: Date;

}

class APIPagination {
    PageSize: number = 10
    Page: number = 0
    TotalPages: number = 0

}
export enum APIResponseStatus {
    StatusSuccess = 0,
	StatusWarning,
	StatusFailure,
	StatusFatalError
}