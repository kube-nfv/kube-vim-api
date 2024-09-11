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

// SwitchLBRuleInformer provides access to a shared informer and lister for
// SwitchLBRules.
type SwitchLBRuleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() kubeovnv1.SwitchLBRuleLister
}

type switchLBRuleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewSwitchLBRuleInformer constructs a new informer for SwitchLBRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSwitchLBRuleInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSwitchLBRuleInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredSwitchLBRuleInformer constructs a new informer for SwitchLBRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSwitchLBRuleInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeovnV1().SwitchLBRules().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeovnV1().SwitchLBRules().Watch(context.TODO(), options)
			},
		},
		&apiskubeovnv1.SwitchLBRule{},
		resyncPeriod,
		indexers,
	)
}

func (f *switchLBRuleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSwitchLBRuleInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *switchLBRuleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiskubeovnv1.SwitchLBRule{}, f.defaultInformer)
}

func (f *switchLBRuleInformer) Lister() kubeovnv1.SwitchLBRuleLister {
	return kubeovnv1.NewSwitchLBRuleLister(f.Informer().GetIndexer())
}