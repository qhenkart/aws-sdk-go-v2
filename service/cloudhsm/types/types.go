// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// A key-value pair that identifies or specifies metadata about an AWS CloudHSM
// resource.
type Tag struct {

	// The key of the tag.
	//
	// This member is required.
	Key *string

	// The value of the tag.
	//
	// This member is required.
	Value *string
}
