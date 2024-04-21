package main

import (
	"fmt"
	"time"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/operator-framework/operator-sdk/pkg/util/k8sutil"
	sdkVersion "github.com/operator-framework/operator-sdk/version"

	"github.com/example/app-operator/pkg/apis"
	"github.com/example/app-operator/pkg/controller"
)

func main() {
	sdk.ExposeMetricsPort()

	resource := "app.example.com/v1/sampledbs"
	kind := "SampleDB"
	namespace, err := k8sutil.GetWatchNamespace()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	resyncPeriod := time.Duration(5) * time.Second
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	ctx := sdk.New(sdkVersion.Version, sdk.WithLogger(logger), sdk.WithNamespace(namespace))

	sdk.Watch(resource, kind, namespace, resyncPeriod)
	sdk.Handle(controller.NewHandler() )
}

