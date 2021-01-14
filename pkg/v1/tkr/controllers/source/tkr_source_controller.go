// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package source

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	runv1 "github.com/vmware-tanzu-private/core/apis/run/v1alpha1"
	mgrcontext "github.com/vmware-tanzu-private/core/pkg/v1/tkr/pkg/context"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const (
	controllerName = "tkr-source-controller"
	tkrFinalizer   = "tanzukubernetesrelease.run.tanzu.vmware.com"
)

// TanzuKubernetesReleaseReconciler reconciles a TanzuKubernetesRelease object
type reconciler struct {
	ctx context.Context
	client.Client
	log    logr.Logger
	scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=run.tanzu.vmware.com,resources=tanzukubernetesreleases,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=run.tanzu.vmware.com,resources=tanzukubernetesreleases/status,verbs=get;update;patch

func (r *reconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.log.WithValues("tanzukubernetesrelease", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

// AddToManager adds this package's controller to the provided manager.
func AddToManager(ctx *mgrcontext.ControllerManagerContext, mgr ctrl.Manager) error {
	c, err := controller.New(controllerName, mgr, controller.Options{
		Reconciler: newReconciler(ctx),
	})
	if err != nil {
		return errors.Wrapf(err, "error constructing controller '%s'", controllerName)
	}
	return c.Watch(&source.Kind{Type: &runv1.TanzuKubernetesRelease{}}, &handler.EnqueueRequestForObject{})
}

func newReconciler(ctx *mgrcontext.ControllerManagerContext) *reconciler {
	return &reconciler{
		ctx:    ctx.Context,
		Client: ctx.Client,
		log:    ctx.Logger,
		scheme: ctx.Scheme,
	}
}