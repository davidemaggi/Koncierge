export class k8sObject {
    constructor(public nameSpace: string, public name: string, public kind:k8sObjectKind) {}
}


export enum k8sObjectKind {
    pod = 'pod'

}