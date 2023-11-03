# Custom scheduler demo

Below are some instructions for testing out the custom scheduler using minikube.

1. Ensure you have minikube installed (see [docs](https://minikube.sigs.k8s.io/docs/start/))
2. Start a new multi-node cluster
    ```shell
    minikube start --nodes 3 -p custom-scheduler-demo
    ```
3. We need to enable the minikube [registry](https://minikube.sigs.k8s.io/docs/handbook/pushing/#4-pushing-to-an-in-cluster-using-registry-addon) plugin to be able to deploy our custom scheduler container:
    ```shell
    minikube addons enable registry -p custom-scheduler-demo
    ```
4. We also need to [enable insecure registries](https://minikube.sigs.k8s.io/docs/handbook/registry/#enabling-insecure-registries) so that we can push images without needing to provision TLS certificates:
    ```shell
    docker run --rm -it --network=host alpine ash -c "apk add socat && socat TCP-LISTEN:5000,reuseaddr,fork TCP:$(minikube ip -p custom-scheduler-demo):5000"
    ```
5. Next, we can build and push our custom scheduler container to the registry in minikube:
    ```shell
    docker build --tag localhost:5000/custom-scheduler .
    docker push localhost:5000/custom-scheduler
    ```
6. Deploy our custom scheduler:
    ```shell
    kubectl apply -f deploy/custom-scheduler.yaml
    ```
7. Since the custom scheduler expects there to be a `nodeGroup` label on the nodes, we need to add these using:
    ```shell
    kubectl label nodes custom-scheduler-demo-m02 nodeGroup=1
    kubectl label nodes custom-scheduler-demo-m03 nodeGroup=2
    ```
8. Deploy two demo apps (these are just taken from the minikube [tutorial](https://minikube.sigs.k8s.io/docs/start/)):
    ```shell
    kubectl apply -f demo/demo-app.yaml
    ```
9. Verify it all worked. Since each of the app's deployments has a different `nodeGroup` label, you should see all pods from `hello-minikube-1` deployed to node `custom-scheduler-demo-m02` and all pods from `hello-minikube-2` deployed to node `custom-scheduler-demo-m03`:
    ```
    kubectl get pods -o wide

    NAME                                READY   STATUS    RESTARTS   AGE   IP            NODE                        NOMINATED NODE   READINESS GATES
    hello-minikube-1-7bc54fdc74-2vrlg   1/1     Running   0          66s   10.244.1.16   custom-scheduler-demo-m02   <none>           <none>
    hello-minikube-1-7bc54fdc74-zms9g   1/1     Running   0          67s   10.244.1.15   custom-scheduler-demo-m02   <none>           <none>
    hello-minikube-2-569d97b8bf-cqtjm   1/1     Running   0          64s   10.244.2.15   custom-scheduler-demo-m03   <none>           <none>
    hello-minikube-2-569d97b8bf-pp9dx   1/1     Running   0          62s   10.244.2.16   custom-scheduler-demo-m03   <none>           <none>
    ```