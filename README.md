# Cash Service OPerator

The Cash Service OPerator (CSOP) is a Kubernetes operator that codifies the operational knowledge required to run a service (based on Misk or otherwise) in a Kubernetes cluster.

CSOP uses CRDs to allow service owners to easily deploy a service that conforms to the container contract specified by CSOP. The resources created by CSOP will be configured with sane defaults that emphasize security. CSOP provides the following features:
- Autoscaling
- Scheduled Scaling
- Autodeployment
- Canary Deployments
- Jobs Deployments
- Debug Containers
- Resource Customization

## Deployment

Service owners deploy a CSOP for each service they want to run in a namespace. They will then deploy a CashService Custom Resource (CR) with the small amount of desired configuration for the service. CSOP will observe the creation of the CashService and create kubernetes resources including a Deployment, HorizontalPodAutoscaler, NetworkPolicy, Service, Secrets, and ConfigMaps. These resources will be owned and managed by CSOP to ensure the service always operates in the way CSOP describes.

## Autoscaling

CSOP will create a HorizontalPodAutoscaler for the service using the configuration provided in the CashService CR.

## Scheduled Scaling (Coming Soon)

CSOP is able to scale up a service at a specific date and time using the CashServiceScheduledScale CR. Service owners are able to add replicas to their deployment in anticipation of an event that increases traffic.

## Canary Deployments (Coming Soon)

CSOP will allow service owners to create a small, secondary deployment that can be used to test changes while taking a fraction of the traffic.

## Jobs Deployment (Coming Soon)

CSOP enables service owners to create a separate deployment that can be used to handle jobs created by the service.

## Debug Containers

CSOP will enable service owners to easily add a container to the service pods that contains a number of tools that can be used to debug the service. NOTE: This will require a restart of the service.

## Resource Customization (Coming Soon)

CSOP provides the ability to modify the managed resources outside of the CashService CRD using the CashServiceExtension CR. The CashServiceExtension contains raw yaml that will be applied to the specified resource. Service owners will be customize the CSOP managed resources in case the CashService CR is not flexible enough and the default values are not suitable for their cluster.
