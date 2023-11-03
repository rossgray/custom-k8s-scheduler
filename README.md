# Creating a Custom Kubernetes Scheduler

This is a sample repo demonstrating how to create a custom scheduler in Kubernetes.

For this proof-of-concept we just add a custom scheduler plugin that will schedule pods only on nodes which match a given label.
Obviously, it would not make sense to use this plugin in reality, since the same behaviour can easily be achieved using built-in functionality, such as [using a `nodeSelector` in your deployment][assign pods to nodes].

If we wanted more control over how our pods get scheduled however, it would be reasonably straightforward to modify this plugin to add the desired behaviour.

## Demo

See [demo/README.md](demo/README.md) for instructions on how to try out this custom scheduler in a local minikube deployment.

## Useful resources

- [This repo][crl-scheduler] for a custom scheduler written by Chris Seto at Cockroach Labs was very useful as a reference
- The [Kubernetes Scheduling Framework][scheduling framework]
- Source code for other [scheduler plugins][scheduler plugins]

[assign pods to nodes]: https://kubernetes.io/docs/tasks/configure-pod-container/assign-pods-nodes/
[scheduling framework]: https://kubernetes.io/docs/concepts/scheduling-eviction/scheduling-framework/
[configure multiple schedulers]: https://kubernetes.io/docs/tasks/extend-kubernetes/configure-multiple-schedulers/
[crl-scheduler]: https://github.com/cockroachlabs/crl-scheduler/tree/master
[scheduler plugins]: https://github.com/kubernetes-sigs/scheduler-plugins