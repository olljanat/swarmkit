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
	{StubManagerInit("macvlan", datastore.LocalScope, datastore.GlobalScope), "macvlan"},
	{StubManagerInit("bridge", datastore.LocalScope, datastore.LocalScope), "bridge"},
	{StubManagerInit("ipvlan", datastore.LocalScope, datastore.GlobalScope), "ipvlan"},
	{StubManagerInit("host", datastore.LocalScope, datastore.LocalScope), "host"},
}

// PredefinedNetworks returns the list of predefined network structures
func PredefinedNetworks() []networkallocator.PredefinedNetworkData {
	return []networkallocator.PredefinedNetworkData{
		{Name: "bridge", Driver: "bridge"},
		{Name: "host", Driver: "host"},
	}
}
