// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package mock

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	lifecyclev1alpha1 "github.com/ironcore-dev/lifecycle-manager/api/lifecycle/v1alpha1"
)

type MachineMockBuilder struct {
	inner *lifecyclev1alpha1.Machine
}

func (b *UnstructuredBuilder) MachineFromUnstructured() *MachineMockBuilder {
	var m lifecyclev1alpha1.Machine
	b.inner.SetAPIVersion("lifecycle.ironcore.dev/v1alpha1")
	b.inner.SetKind("Machine")
	_ = runtime.DefaultUnstructuredConverter.FromUnstructured(b.inner.Object, &m)
	return &MachineMockBuilder{inner: &m}
}

func (b *MachineMockBuilder) WithMachineTypeRef(name string) *MachineMockBuilder {
	b.inner.Spec.MachineTypeRef.Name = name
	return b
}

func (b *MachineMockBuilder) WithOOBMachineRef(name string) *MachineMockBuilder {
	b.inner.Spec.OOBMachineRef.Name = name
	return b
}

func (b *MachineMockBuilder) WithDesiredPackages(packages ...lifecyclev1alpha1.PackageVersion) *MachineMockBuilder {
	packagesToApply := make([]lifecyclev1alpha1.PackageVersion, len(packages))
	copy(packagesToApply, packages)
	b.inner.Spec.Packages = packagesToApply
	return b
}

func (b *MachineMockBuilder) WithInstalledPackages(packages ...lifecyclev1alpha1.PackageVersion) *MachineMockBuilder {
	packagesToApply := make([]lifecyclev1alpha1.PackageVersion, len(packages))
	copy(packagesToApply, packages)
	b.inner.Status.InstalledPackages = packagesToApply
	return b
}

func (b *MachineMockBuilder) WithLastScanTime(timestamp metav1.Time) *MachineMockBuilder {
	b.inner.Status.LastScanTime = timestamp
	return b
}

func (b *MachineMockBuilder) Complete() *lifecyclev1alpha1.Machine {
	return b.inner
}
