package cnmallocator

import (
	"github.com/docker/docker/libnetwork/datastore"
	"github.com/docker/docker/libnetwork/drivers/overlay/ovmanager"
	"github.com/docker/docker/libnetwork/drivers/remote"
	"github.com/moby/swarmkit/v2/manager/allocator/networkallocator"
)

var initializers = []initializer{
	{remote.Init, "remote"},
	{ovmanager.Init, "overlay"},
	{StubManagerInit("internal", datastore.LocalScope, datastore.LocalScope), "internal"},
	{StubManagerInit("l2bridge", datastore.LocalScope, datastore.LocalScope), "l2bridge"},
	{StubManagerInit("nat", datastore.LocalScope, datastore.LocalScope), "nat"},
}

// PredefinedNetworks returns the list of predefined network structures
func PredefinedNetworks() []networkallocator.PredefinedNetworkData {
	return []networkallocator.PredefinedNetworkData{
		{Name: "nat", Driver: "nat"},
	}
}
