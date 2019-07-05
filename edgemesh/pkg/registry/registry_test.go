package registry

import (
utiltags "github.com/go-chassis/go-chassis/pkg/util/tags"
"github.com/kubeedge/beehive/pkg/core/context"
"github.com/kubeedge/kubeedge/edge/pkg/metamanager/client"
"testing"
)

func TestFindMicroServiceInstances(t *testing.T) {
	c:= context.GetContext(context.MsgCtxTypeChannel)
	sd := &ServiceDiscovery{
		metaClient:client.New(c),
	}

	_, _ = sd.FindMicroServiceInstances("123", "name.namespace", utiltags.Tags{})
}

