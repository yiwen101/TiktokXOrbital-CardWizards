package service

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/thriftgo/pkg/test"
	"github.com/yiwen101/CardWizards/common"
	"github.com/yiwen101/CardWizards/common/descriptor"
	"github.com/yiwen101/CardWizards/service/clients"
)

/* tested: cache's get, save, HandlerForAnnotatedRoutes, HandlerForRoute, GetHandlerManager */
func TestHandlerCache(t *testing.T) {
	h := handlerCache{}
	f, ok := h.get("test", "test")
	test.Assert(t, ok == false)
	test.Assert(t, f == nil)
	testEffect := 0
	theFunc := func(ctx context.Context, c *app.RequestContext) { testEffect++ }
	h.save("test", "test", theFunc)
	f, ok = h.get("test", "test")
	test.Assert(t, ok == true)
	test.Assert(t, f != nil)
	f(context.Background(), nil)
	test.Assert(t, testEffect == 1)
}

func TestHandlerManager(t *testing.T) {
	descriptor.BuildDescriptorManager(common.RelativePathToIDLFromTest2)
	clients.BuildGenericClients(common.RelativePathToIDLFromTest2)
	hm, err := GetHandlerManager()
	test.Assert(t, err == nil)
	f, err := hm.HandlerForAnnotatedRoutes("GET")
	test.Assert(t, err == nil)
	test.Assert(t, f != nil)
	f, err = hm.HandlerForRoute("arithmetic", "Add")
	test.Assert(t, err == nil)
	test.Assert(t, f != nil)
	f, err = hm.HandlerForRoute("fake", "fake")
	test.Assert(t, err != nil)
	test.Assert(t, f == nil)
}