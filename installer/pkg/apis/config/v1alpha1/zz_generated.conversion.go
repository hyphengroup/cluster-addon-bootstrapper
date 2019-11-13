// +build !ignore_autogenerated

/*

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
// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	config "sigs.k8s.io/addon-operators/installer/pkg/apis/config"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Addon)(nil), (*config.Addon)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Addon_To_config_Addon(a.(*Addon), b.(*config.Addon), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Addon)(nil), (*Addon)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Addon_To_v1alpha1_Addon(a.(*config.Addon), b.(*Addon), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AddonInstallerConfiguration)(nil), (*config.AddonInstallerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AddonInstallerConfiguration_To_config_AddonInstallerConfiguration(a.(*AddonInstallerConfiguration), b.(*config.AddonInstallerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.AddonInstallerConfiguration)(nil), (*AddonInstallerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_AddonInstallerConfiguration_To_v1alpha1_AddonInstallerConfiguration(a.(*config.AddonInstallerConfiguration), b.(*AddonInstallerConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_Addon_To_config_Addon(in *Addon, out *config.Addon, s conversion.Scope) error {
	out.Name = in.Name
	out.KustomizeRef = in.KustomizeRef
	out.ManifestRef = in.ManifestRef
	return nil
}

// Convert_v1alpha1_Addon_To_config_Addon is an autogenerated conversion function.
func Convert_v1alpha1_Addon_To_config_Addon(in *Addon, out *config.Addon, s conversion.Scope) error {
	return autoConvert_v1alpha1_Addon_To_config_Addon(in, out, s)
}

func autoConvert_config_Addon_To_v1alpha1_Addon(in *config.Addon, out *Addon, s conversion.Scope) error {
	out.Name = in.Name
	out.KustomizeRef = in.KustomizeRef
	out.ManifestRef = in.ManifestRef
	return nil
}

// Convert_config_Addon_To_v1alpha1_Addon is an autogenerated conversion function.
func Convert_config_Addon_To_v1alpha1_Addon(in *config.Addon, out *Addon, s conversion.Scope) error {
	return autoConvert_config_Addon_To_v1alpha1_Addon(in, out, s)
}

func autoConvert_v1alpha1_AddonInstallerConfiguration_To_config_AddonInstallerConfiguration(in *AddonInstallerConfiguration, out *config.AddonInstallerConfiguration, s conversion.Scope) error {
	out.DryRun = in.DryRun
	out.Addons = *(*[]config.Addon)(unsafe.Pointer(&in.Addons))
	return nil
}

// Convert_v1alpha1_AddonInstallerConfiguration_To_config_AddonInstallerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_AddonInstallerConfiguration_To_config_AddonInstallerConfiguration(in *AddonInstallerConfiguration, out *config.AddonInstallerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_AddonInstallerConfiguration_To_config_AddonInstallerConfiguration(in, out, s)
}

func autoConvert_config_AddonInstallerConfiguration_To_v1alpha1_AddonInstallerConfiguration(in *config.AddonInstallerConfiguration, out *AddonInstallerConfiguration, s conversion.Scope) error {
	out.DryRun = in.DryRun
	out.Addons = *(*[]Addon)(unsafe.Pointer(&in.Addons))
	return nil
}

// Convert_config_AddonInstallerConfiguration_To_v1alpha1_AddonInstallerConfiguration is an autogenerated conversion function.
func Convert_config_AddonInstallerConfiguration_To_v1alpha1_AddonInstallerConfiguration(in *config.AddonInstallerConfiguration, out *AddonInstallerConfiguration, s conversion.Scope) error {
	return autoConvert_config_AddonInstallerConfiguration_To_v1alpha1_AddonInstallerConfiguration(in, out, s)
}
