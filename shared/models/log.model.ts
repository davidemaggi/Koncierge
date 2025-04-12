export class Pod {
    constructor(public label: string, public value: string) {}
}


export enum LogEntryType {
    info = 'info',
    warning = 'warning',
    error = 'error'

}