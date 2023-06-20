package clients

import (
	"testing"

	"github.com/cloudwego/thriftgo/pkg/test"
	"github.com/yiwen101/CardWizards/common"
	"github.com/yiwen101/CardWizards/common/descriptor"
)

/*
tested: GetGenericClientforService, BuildGenericClients, buildGenericClientFromPath
*/
func TestBuildGenericClientFromPath(t *testing.T) {
	filename := "arithmetic.thrift"
	includeDir := common.RelativePathToIDLFromTest
	_, e := buildGenericClientFromPath("arithmetic", filename, includeDir)
	test.Assert(t, e == nil)
}

func TestBuildGenericClientsAndGetGenericClientforService(t *testing.T) {
	descriptor.BuildDescriptorManager(common.RelativePathToIDLFromTest)
	err := BuildGenericClients(common.RelativePathToIDLFromTest)
	test.Assert(t, err == nil)
	g1, err := GetGenericClientforService("arithmetic")
	test.Assert(t, g1 != nil)
	test.Assert(t, err == nil)
	g2, err := GetGenericClientforService("arithmetic2")
	test.Assert(t, g2 == nil)
	test.Assert(t, err != nil)
}