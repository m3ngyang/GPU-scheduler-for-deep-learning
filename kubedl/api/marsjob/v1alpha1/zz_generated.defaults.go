/*
Copyright 2020 The Alibaba Authors.

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
// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&MarsJob{}, func(obj interface{}) { SetObjectDefaults_MarsJob(obj.(*MarsJob)) })
	scheme.AddTypeDefaultingFunc(&MarsJobList{}, func(obj interface{}) { SetObjectDefaults_MarsJobList(obj.(*MarsJobList)) })
	return nil
}

func SetObjectDefaults_MarsJob(in *MarsJob) {
	SetDefaults_MarsJob(in)
}

func SetObjectDefaults_MarsJobList(in *MarsJobList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_MarsJob(a)
	}
}
