package client

import (
	"fmt"
	"testing"
)

var bgp_name string = "test-bgp"
var as int = 65534
var clientToClientReflection bool = true
var clusterId string = "172.21.16.1"
var bgp_comment string = ""
var confederation int = 7
var confederationPeers string = ""
var disabled bool = false
var ignoreAsPathLen bool = false
var outFilter string = ""
var redistributeConnected bool = false
var redistributeOspf bool = false
var redistributeOtherBgp bool = false
var redistributeRip bool = false
var redistributeStatic bool = false
var routerId string = "172.21.16.2"
var routingTable string = ""

func TestAddBgpInstanceAndDeleteBgpInstance(t *testing.T) {
	c := NewClient(GetConfigFromEnv())

	expectedBgpInstance := &BgpInstance{
		Name:                     bgp_name,
		As:                       as,
		ClientToClientReflection: clientToClientReflection,
		ClusterId:                clusterId,
		Comment:                  bgp_comment,
		Confederation:            confederation,
		ConfederationPeers:       confederationPeers,
		Disabled:                 disabled,
		IgnoreAsPathLen:          ignoreAsPathLen,
		OutFilter:                outFilter,
		RedistributeConnected:    redistributeConnected,
		RedistributeOspf:         redistributeOspf,
		RedistributeOtherBgp:     redistributeOtherBgp,
		RedistributeRip:          redistributeRip,
		RedistributeStatic:       redistributeStatic,
		RouterId:                 routerId,
		RoutingTable:             routingTable,
	}
	bgpInstance, err := c.AddBgpInstance(
		bgp_name,
		as,
		clientToClientReflection,
		clusterId,
		bgp_comment,
		confederation,
		confederationPeers,
		disabled,
		ignoreAsPathLen,
		outFilter,
		redistributeConnected,
		redistributeOspf,
		redistributeOtherBgp,
		redistributeRip,
		redistributeStatic,
		routerId,
		routingTable,
	)

	if err != nil {
		t.Fatalf("Error creating a bpg instance with: %v", err)
	}
	fmt.Println(bgpInstance)
	fmt.Println(expectedBgpInstance)

	//if len(pool.Id) < 1 {
	//	t.Errorf("The created pool does not have an Id: %v", pool)
	//}

	//if pool.Name != expectedPool.Name {
	//	t.Errorf("The pool Name fields do not match. actual: %v expected: %v", pool.Name, expectedPool.Name)
	//}

	//if pool.Ranges != expectedPool.Ranges {
	//	t.Errorf("The pool Ranges fields do not match. actual: %v expected: %v", pool.Ranges, expectedPool.Ranges)
	//}

	//if pool.Comment != expectedPool.Comment {
	//	t.Errorf("The pool Comment fields do not match. actual: %v expected: %v", pool.Comment, expectedPool.Comment)
	//}

	//foundPool, err := c.FindPool(pool.Id)

	//if err != nil {
	//	t.Errorf("Error getting pool with: %v", err)
	//}

	//if !reflect.DeepEqual(pool, foundPool) {
	//	t.Errorf("Created pool and found pool do not match. actual: %v expected: %v", foundPool, pool)
	//}

	//err = c.DeletePool(pool.Id)

	//if err != nil {
	//	t.Errorf("Error deleting pool with: %v", err)
	//}
}
