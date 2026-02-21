package reconciler

// +kubebuilder:rbac:groups=s3-controller.homelab-helper.benfiola.com,resources=buckets,verbs=get;list;watch
// +kubebuilder:rbac:groups=s3-controller.homelab-helper.benfiola.com,resources=buckets/status,verbs=get;patch;update

import (
	"context"

	"github.com/benfiola/s3-controller/internal/logging"
	apiv1 "github.com/benfiola/s3-controller/pkg/api/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	controllerruntime "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

const (
	ConditionTypeReady = "Ready"

	Finalizer = "s3-controller.benfiola.com/finalizer"

	ReasonFinalizerFailed         = "FinalizerFailed"
	ReasonReconciliationSucceeded = "ReconciliationSucceeded"
)

type BucketReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *BucketReconciler) Register(manager controllerruntime.Manager) error {
	return controllerruntime.
		NewControllerManagedBy(manager).
		For(&apiv1.Bucket{}).
		Complete(r)
}

func (r *BucketReconciler) setCondition(b *apiv1.Bucket, reason string, message string) {
	cstatus := metav1.ConditionFalse
	if reason == ReasonReconciliationSucceeded {
		cstatus = metav1.ConditionTrue
	}

	meta.SetStatusCondition(&b.Status.Conditions, metav1.Condition{
		Type:               ConditionTypeReady,
		Status:             cstatus,
		ObservedGeneration: b.Generation,
		Reason:             reason,
		Message:            message,
	})
}

func (r *BucketReconciler) Reconcile(pctx context.Context, request controllerruntime.Request) (controllerruntime.Result, error) {
	logger := logging.FromContext(pctx).With("resource", request.NamespacedName)
	ctx := logging.WithLogger(pctx, logger)

	bucket := apiv1.Bucket{}
	err := r.Get(ctx, request.NamespacedName, &bucket)
	if err != nil {
		if client.IgnoreNotFound(err) == nil {
			return controllerruntime.Result{}, nil
		}
		logger.Error("failed to fetch bucket", "error", err)
		return controllerruntime.Result{}, err
	}

	if bucket.DeletionTimestamp != nil {
		controllerutil.RemoveFinalizer(&bucket, Finalizer)
		err = r.Update(ctx, &bucket)
		if err != nil {
			logger.Error("failed to remove finalizer during deletion", "error", err)
			r.setCondition(&bucket, ReasonFinalizerFailed, err.Error())
			r.Status().Update(ctx, &bucket)
			return controllerruntime.Result{}, err
		}

		return controllerruntime.Result{}, nil
	}

	if !controllerutil.ContainsFinalizer(&bucket, Finalizer) {
		controllerutil.AddFinalizer(&bucket, Finalizer)
		err = r.Update(ctx, &bucket)
		if err != nil {
			logger.Error("failed to add finalizer", "error", err)
			r.setCondition(&bucket, ReasonFinalizerFailed, err.Error())
			r.Status().Update(ctx, &bucket)
			return controllerruntime.Result{}, err
		}
	}

	return controllerruntime.Result{}, nil
}
