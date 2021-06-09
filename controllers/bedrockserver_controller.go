/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	bedrockv1 "github.com/jasoncwik/kube-bedrock/api/v1"
)

// BedrockServerReconciler reconciles a BedrockServer object
type BedrockServerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=bedrock.cwik.org,resources=bedrockservers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=bedrock.cwik.org,resources=bedrockservers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=bedrock.cwik.org,resources=bedrockservers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the BedrockServer object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *BedrockServerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var log = log.FromContext(ctx).WithValues("bedrockserver", req.NamespacedName)

	// your logic here
	var server bedrockv1.BedrockServer
	if err := r.Get(ctx, req.NamespacedName, &server); err != nil {
		log.Error(err, "unable to fetch BedrockServer")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// See if there's a corresponding StatefulSet
	var ssName = types.NamespacedName{
		Name: "bedrock-" + req.NamespacedName.Name,
		Namespace: req.NamespacedName.Namespace,
	}
	log.Info("Looking for StatefulSet", "name", ssName)

	var serverStatefulSet v1.StatefulSet
	if err := r.Get(ctx, ssName, &serverStatefulSet); err != nil {
		if errors.IsNotFound(err) {
			// The StatefulSet doesn't exist. Create it and return.
			return CreateStatefulSet(ctx, req, ssName), nil
		}
	} else {
		// Look for changes
	}

	// Update status
	if err := r.Status().Update(ctx, &serverStatefulSet); err != nil {
		log.Error(err, "unable to update CronJob status")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func CreateStatefulSet(ctx context.Context, req ctrl.Request, ssName types.NamespacedName) ctrl.Result {
	var log = log.FromContext(ctx).WithValues("bedrockserver", req.NamespacedName)

	log.Info("Creating StatefulSet for server", "Name", req.NamespacedName, "ssName", ssName)

	return ctrl.Result{}
}

// SetupWithManager sets up the controller with the Manager.
func (r *BedrockServerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&bedrockv1.BedrockServer{}).
		Complete(r)
}
