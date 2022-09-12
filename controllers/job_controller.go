/*
Copyright 2022 Adam B Kaplan.

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

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	ArchivedAnnotation = "kubarchive.io/archived"
)

// DynamicReconciler reconciles an object with the given GroupVersionKind
type DynamicReconciler struct {
	client.Client
	Scheme           *runtime.Scheme
	GroupVersionKind schema.GroupVersionKind
}

func NewDynamicReconciler(client client.Client, scheme *runtime.Scheme, gvk schema.GroupVersionKind) *DynamicReconciler {
	return &DynamicReconciler{
		Client:           client,
		Scheme:           scheme,
		GroupVersionKind: gvk,
	}
}

//+kubebuilder:rbac:groups=batch,resources=jobs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=batch,resources=jobs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=batch,resources=jobs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.2/pkg/reconcile
func (r *DynamicReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	obj := r.getUnstructuredObject()
	log.V(5).Info("reconciling object")
	if err := r.Client.Get(ctx, req.NamespacedName, obj); err != nil {
		log.Error(err, "unable to fetch object")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	annotations := obj.GetAnnotations()
	if annotations == nil {
		annotations = map[string]string{}
	}
	if _, ok := annotations[ArchivedAnnotation]; ok {
		log.V(5).Info("object has already been archived")
		return ctrl.Result{}, nil
	}
	annotations[ArchivedAnnotation] = "true"
	obj.SetAnnotations(annotations)
	if err := r.Client.Update(ctx, obj); err != nil {
		log.Error(err, "failed to update object")
		return ctrl.Result{}, err
	}
	// Always requeue if there is an update
	log.Info("updated object with archive annotation")
	return ctrl.Result{Requeue: true}, nil
}

func (r *DynamicReconciler) getUnstructuredObject() *unstructured.Unstructured {
	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(r.GroupVersionKind)
	return obj
}

// SetupWithManager sets up the controller with the Manager.
func (r *DynamicReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(r.getUnstructuredObject()).
		Complete(r)
}
