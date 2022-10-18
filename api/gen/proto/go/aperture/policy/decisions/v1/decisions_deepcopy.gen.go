// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package decisionsv1

import (
	proto "github.com/golang/protobuf/proto"
)

// DeepCopyInto supports using LoadDecision within kubernetes types, where deepcopy-gen is used.
func (in *LoadDecision) DeepCopyInto(out *LoadDecision) {
	p := proto.Clone(in).(*LoadDecision)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadDecision. Required by controller-gen.
func (in *LoadDecision) DeepCopy() *LoadDecision {
	if in == nil {
		return nil
	}
	out := new(LoadDecision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LoadDecision. Required by controller-gen.
func (in *LoadDecision) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TokensDecision within kubernetes types, where deepcopy-gen is used.
func (in *TokensDecision) DeepCopyInto(out *TokensDecision) {
	p := proto.Clone(in).(*TokensDecision)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokensDecision. Required by controller-gen.
func (in *TokensDecision) DeepCopy() *TokensDecision {
	if in == nil {
		return nil
	}
	out := new(TokensDecision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TokensDecision. Required by controller-gen.
func (in *TokensDecision) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RateLimiterDecision within kubernetes types, where deepcopy-gen is used.
func (in *RateLimiterDecision) DeepCopyInto(out *RateLimiterDecision) {
	p := proto.Clone(in).(*RateLimiterDecision)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimiterDecision. Required by controller-gen.
func (in *RateLimiterDecision) DeepCopy() *RateLimiterDecision {
	if in == nil {
		return nil
	}
	out := new(RateLimiterDecision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RateLimiterDecision. Required by controller-gen.
func (in *RateLimiterDecision) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
