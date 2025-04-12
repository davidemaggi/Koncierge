import {V1PodList} from "@kubernetes/client-node";
import {k8sPod} from "../shared/models/k8sPods.model";

const k8s = require('@kubernetes/client-node');

async function fetchPods(): Promise<k8sPod[]> {

    const kc = new k8s.KubeConfig();
    kc.loadFromDefault();

    const k8sApi = kc.makeApiClient(k8s.CoreV1Api);

    const pods: V1PodList = await k8sApi.listNamespacedPod({namespace: 'default'});

    return pods.items.map(p =>k8sPod.fromV1(p)) ?? [];
}

export {
    fetchPods,
}






