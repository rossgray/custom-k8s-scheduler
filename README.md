# Creating a Custom Kubernetes Scheduler

This is a sample repo demonstrating how to create a custom scheduler in Kubernetes.

For this proof-of-concept we just add a custom scheduler plugin that will schedule pods only on nodes which match a given label.
Obviously, it would not make sense to use this plugin in reality, since the same behaviour can easily be achieved using built-in functionality, such as [using a `nodeSelector` in your deployment][assign pods to nodes].

If we wanted more control over how our pods get scheduled however, it would be reasonably straightforward to modify this plugin to add the desired behaviour.

## Demo

See [demo/README.md](demo/README.md) for instructions on how to try out this custom scheduler in a local minikube deployment.

## Motivation

While there are several resources in the Kubernetes documentation regarding how the scheduler works and how it can be extended, it was not clear to me how to go about actually implementing and deploying a custom scheduler.

At first, it appeared as though you can just write your own [custom scheduler][configure multiple schedulers], but then I discovered there was a better way, using the [Scheduling Framework][scheduling framework]. As mentioned in the docs, _The scheduling framework is a pluggable architecture for the Kubernetes scheduler. The APIs allow most scheduling features to be implemented as plugins, while keeping the scheduling "core" lightweight and maintainable._ This means we can just write a custom plugin for the specific point of the scheduling cycle we're interested in.

Although it appears fairly straightforward to write your own scheduler plugin, being neither an expert in Kubernetes or Go, I ran into a few issues along the way (such as [fixing Go dependencies](fix-deps.sh), and what [configuration](deploy/custom-scheduler.yaml) to use when deploying my plugin). Therefore, perhaps this repo will be useful to others who want to implement their own custom scheduler logic.

## Useful resources

- [This repo][crl-scheduler] for a custom scheduler written by Chris Seto at Cockroach Labs was very useful as a reference
- The [Kubernetes Scheduling Framework][scheduling framework]
- Source code for other [scheduler plugins][scheduler plugins]

[assign pods to nodes]: https://kubernetes.io/docs/tasks/configure-pod-container/assign-pods-nodes/
[scheduling framework]: https://kubernetes.io/docs/concepts/scheduling-eviction/scheduling-framework/
[configure multiple schedulers]: https://kubernetes.io/docs/tasks/extend-kubernetes/configure-multiple-schedulers/
[crl-scheduler]: https://github.com/cockroachlabs/crl-scheduler/tree/master
[scheduler plugins]: https://github.com/kubernetes-sigs/scheduler-plugins