// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	context "context"
	time "time"

	apiskubeovnv1 "github.com/DiMalovanyy/kube-vim-api/kube-ovn-api/pkg/apis/kubeovn/v1"
	versioned "github.com/DiMalovanyy/kube-vim-api/kube-ovn-api/pkg/client/clientset/versioned"
	internalinterfaces "github.com/DiMalovanyy/kube-vim-api/kube-ovn-api/pkg/client/informers/externalversions/internalinterfaces"
	kubeovnv1 "github.com/DiMalovanyy/kube-vim-api/kube-ovn-api/pkg/client/listers/kubeovn/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VipInformer provides access to a shared informer and lister for
// Vips.
type VipInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() kubeovnv1.VipLister
}

type vipInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewVipInformer constructs a new informer for Vip type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVipInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVipInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredVipInformer constructs a new informer for Vip type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVipInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeovnV1().Vips().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeovnV1().Vips().Watch(context.TODO(), options)
			},
		},
		&apiskubeovnv1.Vip{},
		resyncPeriod,
		indexers,
	)
}

func (f *vipInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVipInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *vipInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiskubeovnv1.Vip{}, f.defaultInformer)
}

func (f *vipInformer) Lister() kubeovnv1.VipLister {
	return kubeovnv1.NewVipLister(f.Informer().GetIndexer())
}