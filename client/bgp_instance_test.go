package client

import (
	"fmt"
	"testing"
)

var bgpName string = "test-bgp"
var as int = 65534
var clientToClientReflection bool = true
var clusterID string = "172.21.16.1"
var noClusterID string = ""
var bgpComment string = ""
var confederation int = 7
var noConfederation int = -1
var confederationPeers string = ""
var disabled bool = false
var ignoreAsPathLen bool = false
var outFilter string = ""
var redistributeConnected bool = false
var redistributeOspf bool = false
var redistributeOtherBgp bool = false
var redistributeRip bool = false
var redistributeStatic bool = false
var routerID string = "172.21.16.2"
var routingTable string = ""

func TestAddBgpInstanceIncludingOptionalFieldsAndDeleteBgpInstance(t *testing.T) {
	c := NewClient(GetConfigFromEnv())

	expectedBgpInstance := &BgpInstance{
		Name:                     bgpName,
		As:                       as,
		ClientToClientReflection: clientToClientReflection,
		Comment:                  bgpComment,
		ConfederationPeers:       confederationPeers,
		Disabled:                 disabled,
		IgnoreAsPathLen:          ignoreAsPathLen,
		OutFilter:                outFilter,
		RedistributeConnected:    redistributeConnected,
		RedistributeOspf:         redistributeOspf,
		RedistributeOtherBgp:     redistributeOtherBgp,
		RedistributeRip:          redistributeRip,
		RedistributeStatic:       redistributeStatic,
		RouterID:                 routerID,
		RoutingTable:             routingTable,
		ClusterID:                clusterID,
		Confederation:            confederation,
	}
	bgpInstance, err := c.AddBgpInstance(
		bgpName,
		as,
		clientToClientReflection,
		bgpComment,
		confederationPeers,
		disabled,
		ignoreAsPathLen,
		outFilter,
		redistributeConnected,
		redistributeOspf,
		redistributeOtherBgp,
		redistributeRip,
		redistributeStatic,
		routerID,
		routingTable,
		clusterID,
		confederation,
	)
	if err != nil {
		t.Fatalf("Error creating a bpg instance with: %v", err)
	}
	fmt.Println(bgpInstance)
	fmt.Println(expectedBgpInstance)
	updatedBgpInstance, err := c.UpdateBgpInstance(
		bgpName,
		as,
		clientToClientReflection,
		bgpComment,
		confederationPeers,
		disabled,
		ignoreAsPathLen,
		outFilter,
		redistributeConnected,
		redistributeOspf,
		redistributeOtherBgp,
		redistributeRip,
		redistributeStatic,
		routerID,
		routingTable,
		noClusterID,
		9,
	)

	fmt.Println(updatedBgpInstance)

	err = c.DeleteBgpInstance(bgpInstance.Name)

	if err != nil {
		t.Errorf("Error deleting bgp instance with: %v", err)
	}
}
