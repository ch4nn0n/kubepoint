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
	"fmt"
	"github.com/go-logr/logr"
	"gopkg.in/yaml.v2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	kp "kubepoint.io/kubepoint/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// DashboardReconciler reconciles a Dashboard object
type DashboardReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=kubepoint.io,resources=dashboards,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=kubepoint.io,resources=dashboards/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=kubepoint.io,resources=dashboards/finalizers,verbs=update
//+kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Dashboard object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *DashboardReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("dashboard", req.NamespacedName)

	var dashboard kp.Dashboard
	if err := r.Get(ctx, req.NamespacedName, &dashboard); err != nil {
		log.Error(err, "unable to fetch Dashboard")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	constructDashboardConfig := func(dashboard *kp.Dashboard) string {
		yaml, err := yaml.Marshal(dashboard.Spec)
		if err != nil {
			log.Error(err, "unable to unmarshal config to yaml")
		}
		return string(yaml)
	}

	constructDashboardConfigMap := func(dashboard *kp.Dashboard) (*corev1.ConfigMap, error) {
		name := fmt.Sprintf("%s-%s", "kubepoint", dashboard.Name)

		configMap := &corev1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: dashboard.Namespace,
			},
			Data: map[string]string{
				"config.yml": constructDashboardConfig(dashboard),
			},
		}
		if err := ctrl.SetControllerReference(dashboard, configMap, r.Scheme); err != nil {
			return nil, err
		}

		return configMap, nil
	}
	configMap, err := constructDashboardConfigMap(&dashboard)
	if err != nil {
		log.Error(err, "unable to construct configmap from template")
	}

	constructDashboardReplicaSet := func(dashboard *kp.Dashboard) (*appsv1.ReplicaSet, error) {
		name := fmt.Sprintf("%s-%s", "kubepoint", dashboard.Name)

		replicaSet := &appsv1.ReplicaSet{
			ObjectMeta: metav1.ObjectMeta{},
			Spec:       appsv1.ReplicaSetSpec{},
		}
	}

	if err := r.Create(ctx, configMap); err != nil {
		log.Error(err, "unable to create config for Dashboard")
		return ctrl.Result{}, err
	}

	log.V(1).Info("created configMap for Dashboard")

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DashboardReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&kp.Dashboard{}).
		Complete(r)
}
