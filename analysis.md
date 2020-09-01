# Steps to re-create issue

1. Create a new Kubernetes cluster with any version above 1.16
2. Install Istio version [1.6.8](https://istio.io/latest/news/releases/1.6.x/announcing-1.6.8/)
3. Enable sidecar injection in the ```default``` namespace.
```
kubectl label namespace default istio-injection=enabled
```
3. Install [Argo](https://github.com/argoproj/argo/blob/master/manifests/install.yaml) on it's own (argo) namespace.
4. Build and push the image:
```
docker build -t <registry>/k8s-api-golang .
docker push <registry>/k8s-api-golang
  ```
5. Update the job.yaml file accordingly: 
```
apiVersion: batch/v1
kind: Job
metadata:
  name: k8s-api-golang-2
spec:
  backoffLimit: 4
  template:
    spec:
      serviceAccountName: k8s-api-job-sa
      imagePullSecrets:
        - name: regcred
      containers:
      - name: k8s-api-golang
        image: <registry>/k8s-api-golang
        command: ["/k8s-job/main", "-sleepTime=5"]
      restartPolicy: Never
```
6. Run this job in the ```default``` namespace. 

The result on any Azure AKS cluster will be the one below, on AWS EKS or GCP GKE it will pass. 

```
2020/09/01 19:25:08 Created new workflow in namespace default, execution took 36.042332ms 
2020/09/01 19:25:08 There are 1 pods in the namespace default, execution took 9.230532ms 
2020/09/01 19:25:08 Sleeping for 5 minutes
panic: Post https://10.0.0.1:443/apis/argoproj.io/v1alpha1/namespaces/default/workflows: unexpected EOF

goroutine 1 [running]:
main.main()
	/k8s-job/app/main.go:98 +0x6c3
```
