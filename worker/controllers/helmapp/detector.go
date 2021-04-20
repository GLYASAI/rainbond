package helmapp

import (
	"github.com/goodrain/rainbond/pkg/apis/rainbond/v1alpha1"
	"github.com/goodrain/rainbond/worker/controllers/helmapp/helm"
	corev1 "k8s.io/api/core/v1"
)

type Detector struct {
	helmApp *v1alpha1.HelmApp
	status  *Status
	repo    *helm.Repo
	app     *helm.App
}

func NewDetector(helmApp *v1alpha1.HelmApp, status *Status, app *helm.App, repo *helm.Repo) *Detector {
	return &Detector{
		helmApp: helmApp,
		status:  status,
		repo:    repo,
		app:     app,
	}
}

func (d *Detector) Detect() error {
	if d.status.isDetected() {
		return nil
	}

	// add repo
	if !d.status.IsConditionTrue(v1alpha1.HelmAppChartReady) {
		appStore := d.helmApp.Spec.AppStore
		if err := d.repo.Add(appStore.Name, appStore.URL, "", ""); err != nil {
			d.status.SetCondition(*v1alpha1.NewHelmAppCondition(
				v1alpha1.HelmAppChartReady, corev1.ConditionFalse, "RepoFailed", err.Error()))
			return err
		}
	}

	// pull chart
	if !d.status.IsConditionTrue(v1alpha1.HelmAppChartReady) {
		err := d.app.Pull()
		if err != nil {
			d.status.UpdateCondition(v1alpha1.NewHelmAppCondition(
				v1alpha1.HelmAppChartReady, corev1.ConditionFalse, "ChartFailed", err.Error()))
			return err
		}
		d.status.UpdateConditionStatus(v1alpha1.HelmAppChartReady, corev1.ConditionTrue)
		return nil
	}

	// check if the chart is valid
	if !d.status.IsConditionTrue(v1alpha1.HelmAppPreInstalled) {
		if err := d.app.PreInstall(); err != nil {
			d.status.UpdateCondition(v1alpha1.NewHelmAppCondition(
				v1alpha1.HelmAppPreInstalled, corev1.ConditionFalse, "PreInstallFailed", err.Error()))
			return err
		}
		d.status.UpdateConditionStatus(v1alpha1.HelmAppPreInstalled, corev1.ConditionTrue)
		return nil
	}

	// parse chart
	if !d.status.IsConditionTrue(v1alpha1.HelmAppChartParsed) {
		values, readme, err := d.app.ParseChart()
		if err != nil {
			d.status.UpdateCondition(v1alpha1.NewHelmAppCondition(
				v1alpha1.HelmAppChartParsed, corev1.ConditionFalse, "ChartParsed", err.Error()))
			return err
		}
		d.status.UpdateConditionStatus(v1alpha1.HelmAppChartParsed, corev1.ConditionTrue)
		d.status.ValuesTemplate = values
		d.status.Readme = readme
	}

	return nil
}