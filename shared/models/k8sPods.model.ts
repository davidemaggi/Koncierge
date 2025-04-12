import {k8sObject, k8sObjectKind} from "./k8sObject.model";
import {V1Pod} from "@kubernetes/client-node";

export class k8sPod extends k8sObject {
    constructor(public nameSpace:string, public name: string) {
        super(nameSpace, name , k8sObjectKind.pod);

    }

    public static fromV1(pod:V1Pod) {
        return new k8sPod(pod.metadata?.namespace ?? "", pod.metadata?.name??"");
    }
}