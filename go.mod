module k8sapi-test

go 1.13.1

// need to create a test with these settings https://github.com/kubeflow/pipelines/blob/master/go.mod
require (
	github.com/argoproj/argo v2.3.0+incompatible
	github.com/cenkalti/backoff v2.0.0+incompatible
	github.com/emicklei/go-restful v2.14.0+incompatible // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-openapi/spec v0.19.9 // indirect
	github.com/gogo/protobuf v1.1.1
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.2
	github.com/google/btree v1.0.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/googleapis/gnostic v0.4.0
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/pkg/errors v0.8.0
	github.com/sirupsen/logrus v1.6.0
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	google.golang.org/genproto v0.0.0-20190801165951-fa694d86fc64
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/api v0.0.0-20180712090710-2d6f90ab1293
	k8s.io/apimachinery v0.0.0-20180621070125-103fd098999d
	k8s.io/client-go v0.0.0-20180718001006-59698c7d9724
	k8s.io/kube-openapi v0.0.0-20180719232738-d8ea2fe547a4 // indirect
	k8s.io/kubernetes v1.11.1
	sigs.k8s.io/controller-runtime v0.0.0-20181121180216-5558165425ef
	sigs.k8s.io/testing_frameworks v0.1.1 // indirect
	sigs.k8s.io/yaml v1.2.0
)
