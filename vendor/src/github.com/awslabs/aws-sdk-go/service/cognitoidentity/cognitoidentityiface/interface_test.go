// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cognitoidentityiface_test

import (
	"testing"

	"github.com/awslabs/aws-sdk-go/service/cognitoidentity"
	"github.com/awslabs/aws-sdk-go/service/cognitoidentity/cognitoidentityiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*cognitoidentityiface.CognitoIdentityAPI)(nil), cognitoidentity.New(nil))
}