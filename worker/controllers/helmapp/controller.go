package helmapp

import (
	"time"

	"github.com/goodrain/rainbond/pkg/generated/clientset/versioned"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/util/workqueue"
)

// Controller -
type Controller struct {
	storer      Storer
	stopCh      chan struct{}
	controlLoop *ControlLoop
}

func NewController(stopCh chan struct{}, clientset versioned.Interface, resyncPeriod time.Duration,
	repoFile, repoCache string) *Controller {
	queue := workqueue.New()
	storer := NewStorer(clientset, resyncPeriod, queue)

	controlLoop := NewControlLoop(clientset, storer, queue, repoFile, repoCache)

	return &Controller{
		storer:      storer,
		stopCh:      stopCh,
		controlLoop: controlLoop,
	}
}

func (c *Controller) Start() {
	logrus.Info("start helm app controller")
	go c.storer.Run(c.stopCh)

	c.controlLoop.Run()
}