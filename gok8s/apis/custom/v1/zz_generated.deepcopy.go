// +build !ignore_autogenerated

/*
Copyright 2019 The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlueGreenSpec) DeepCopyInto(out *BlueGreenSpec) {
	*out = *in
	if in.LabelNames != nil {
		in, out := &in.LabelNames, &out.LabelNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlueGreenSpec.
func (in *BlueGreenSpec) DeepCopy() *BlueGreenSpec {
	if in == nil {
		return nil
	}
	out := new(BlueGreenSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigSpec) DeepCopyInto(out *ConfigSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigSpec.
func (in *ConfigSpec) DeepCopy() *ConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerSpec) DeepCopyInto(out *ContainerSpec) {
	*out = *in
	if in.Verify != nil {
		in, out := &in.Verify, &out.Verify
		*out = make([]VerifySpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerSpec.
func (in *ContainerSpec) DeepCopy() *ContainerSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HistorySpec) DeepCopyInto(out *HistorySpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HistorySpec.
func (in *HistorySpec) DeepCopy() *HistorySpec {
	if in == nil {
		return nil
	}
	out := new(HistorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KCD) DeepCopyInto(out *KCD) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KCD.
func (in *KCD) DeepCopy() *KCD {
	if in == nil {
		return nil
	}
	out := new(KCD)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KCD) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KCDList) DeepCopyInto(out *KCDList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KCD, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KCDList.
func (in *KCDList) DeepCopy() *KCDList {
	if in == nil {
		return nil
	}
	out := new(KCDList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KCDList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KCDSpec) DeepCopyInto(out *KCDSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Container.DeepCopyInto(&out.Container)
	in.Strategy.DeepCopyInto(&out.Strategy)
	out.History = in.History
	out.Rollback = in.Rollback
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		if *in == nil {
			*out = nil
		} else {
			*out = new(ConfigSpec)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KCDSpec.
func (in *KCDSpec) DeepCopy() *KCDSpec {
	if in == nil {
		return nil
	}
	out := new(KCDSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KCDStatus) DeepCopyInto(out *KCDStatus) {
	*out = *in
	in.CurrStatusTime.DeepCopyInto(&out.CurrStatusTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KCDStatus.
func (in *KCDStatus) DeepCopy() *KCDStatus {
	if in == nil {
		return nil
	}
	out := new(KCDStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollbackSpec) DeepCopyInto(out *RollbackSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollbackSpec.
func (in *RollbackSpec) DeepCopy() *RollbackSpec {
	if in == nil {
		return nil
	}
	out := new(RollbackSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StrategySpec) DeepCopyInto(out *StrategySpec) {
	*out = *in
	if in.BlueGreen != nil {
		in, out := &in.BlueGreen, &out.BlueGreen
		if *in == nil {
			*out = nil
		} else {
			*out = new(BlueGreenSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Verify != nil {
		in, out := &in.Verify, &out.Verify
		*out = make([]VerifySpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StrategySpec.
func (in *StrategySpec) DeepCopy() *StrategySpec {
	if in == nil {
		return nil
	}
	out := new(StrategySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerifySpec) DeepCopyInto(out *VerifySpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerifySpec.
func (in *VerifySpec) DeepCopy() *VerifySpec {
	if in == nil {
		return nil
	}
	out := new(VerifySpec)
	in.DeepCopyInto(out)
	return out
}
