package client

import (
	"fmt"
	"testing"
)

var bgp_name string = "test-bgp"
var as int = 65534
var clientToClientReflection bool = true
var clusterId string = "172.21.16.1"
var noClusterId string = ""
var bgp_comment string = ""
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
var routerId string = "172.21.16.2"
var routingTable string = ""

func TestAddBgpInstanceIncludingOptionalFieldsAndDeleteBgpInstance(t *testing.T) {
	c := NewClient(GetConfigFromEnv())

	expectedBgpInstance := &BgpInstance{
		Name:                     bgp_name,
		As:                       as,
		ClientToClientReflection: clientToClientReflection,
		Comment:                  bgp_comment,
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
		ClusterId:                clusterId,
		Confederation:            confederation,
	}
	bgpInstance, err := c.AddBgpInstance(
		bgp_name,
		as,
		clientToClientReflection,
		bgp_comment,
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
		clusterId,
		confederation,
	)
	if err != nil {
		t.Fatalf("Error creating a bpg instance with: %v", err)
	}
	fmt.Println(bgpInstance)
	fmt.Println(expectedBgpInstance)
	updatedBgpInstance, err := c.UpdateBgpInstance(
		bgp_name,
		as,
		clientToClientReflection,
		bgp_comment,
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
		noClusterId,
		9,
	)

	fmt.Println(updatedBgpInstance)

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

	//err = c.DeleteBgpInstance(bgpInstance.Name)

	//if err != nil {
	//	t.Errorf("Error deleting bgp instance with: %v", err)
	//}
}

//func TestAddBgpInstanceExcludingOptionalFieldsAndDeleteBgpInstance(t *testing.T) {
//	c := NewClient(GetConfigFromEnv())
//
//	expectedBgpInstance := &BgpInstance{
//		Name:                     bgp_name,
//		As:                       as,
//		ClientToClientReflection: clientToClientReflection,
//		Comment:                  bgp_comment,
//		ConfederationPeers:       confederationPeers,
//		Disabled:                 disabled,
//		IgnoreAsPathLen:          ignoreAsPathLen,
//		OutFilter:                outFilter,
//		RedistributeConnected:    redistributeConnected,
//		RedistributeOspf:         redistributeOspf,
//		RedistributeOtherBgp:     redistributeOtherBgp,
//		RedistributeRip:          redistributeRip,
//		RedistributeStatic:       redistributeStatic,
//		RouterId:                 routerId,
//		RoutingTable:             routingTable,
//		ClusterId:                noClusterId,
//		Confederation:            noConfederation,
//	}
//	bgpInstance, err := c.AddBgpInstance(
//		bgp_name,
//		as,
//		clientToClientReflection,
//		bgp_comment,
//		confederationPeers,
//		disabled,
//		ignoreAsPathLen,
//		outFilter,
//		redistributeConnected,
//		redistributeOspf,
//		redistributeOtherBgp,
//		redistributeRip,
//		redistributeStatic,
//		routerId,
//		routingTable,
//		noClusterId,
//		noConfederation,
//	)
//
//	if err != nil {
//		t.Fatalf("Error creating a bpg instance with: %v", err)
//	}
//	fmt.Println(bgpInstance)
//	fmt.Println(expectedBgpInstance)
//
//	//if len(pool.Id) < 1 {
//	//	t.Errorf("The created pool does not have an Id: %v", pool)
//	//}
//
//	//if pool.Name != expectedPool.Name {
//	//	t.Errorf("The pool Name fields do not match. actual: %v expected: %v", pool.Name, expectedPool.Name)
//	//}
//
//	//if pool.Ranges != expectedPool.Ranges {
//	//	t.Errorf("The pool Ranges fields do not match. actual: %v expected: %v", pool.Ranges, expectedPool.Ranges)
//	//}
//
//	//if pool.Comment != expectedPool.Comment {
//	//	t.Errorf("The pool Comment fields do not match. actual: %v expected: %v", pool.Comment, expectedPool.Comment)
//	//}
//
//	//foundPool, err := c.FindPool(pool.Id)
//
//	//if err != nil {
//	//	t.Errorf("Error getting pool with: %v", err)
//	//}
//
//	//if !reflect.DeepEqual(pool, foundPool) {
//	//	t.Errorf("Created pool and found pool do not match. actual: %v expected: %v", foundPool, pool)
//	//}
//
//	//err = c.DeletePool(pool.Id)
//
//	//if err != nil {
//	//	t.Errorf("Error deleting pool with: %v", err)
//	//}
//}
