package client

import (
	"fmt"
	"log"
	"strings"
)

type BgpInstance struct {
	Id                       string `mikrotik:".id"`
	Name                     string
	As                       int    `mikroktik:"as"`
	ClientToClientReflection bool   `mikroktik:"client-to-client-reflection"`
	ClusterId                string `mikroktik:"cluster-id"`
	Comment                  string `mikrotik:"comment"`
	Confederation            int    `mikrotik:"confederation"`
	// TODO:  maybe not include
	ConfederationPeers string `mikrotik:"confederation-peers"`
	Disabled           bool   `mikrotik:"disabled"`
	IgnoreAsPathLen    bool   `mikrotik:"ignore-as-path-len"`
	// TODO:  docs says this field is not recommended, "instead use out-filter on peer"
	OutFilter             string `mikrotik:"out-filter"`
	RedistributeConnected bool   `mikrotik:"redistribute-connected"`
	RedistributeOspf      bool   `mikrotik:"redistribute-ospf"`
	RedistributeOtherBgp  bool   `mikrotik:"redistribute-other-bgp"`
	RedistributeRip       bool   `mikrotik:"redistribute-rip"`
	RedistributeStatic    bool   `mikrotik:"redistribute-static"`
	RouterId              string `mikroktik:"router-id"`
	RoutingTable          string `mikroktik:"routing-table"`
}

func (client Mikrotik) AddBgpInstance(name string, as int, clientToClientReflection bool, clusterId string, comment string, confederation int, confederationPeers string, disabled bool, ignoreAsPathLen bool, outFilter string, redistributeConnected bool, redistributeOspf bool, redistributeOtherBgp bool, redistributeRip bool, redistributeStatic bool, routerId string, routingTable string) (*BgpInstance, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := strings.Split(fmt.Sprintf("/routing/bgp/instance/add =as=%d =name=%s =client-to-client-reflection=%s =cluster-id=%s =comment=%s =confederation=%d =confederation-peers=%s =disabled=%s =ignore-as-path-len=%s =out-filter=%s =redistribute-connected=%s =redistribute-ospf=%s =redistribute-other-bgp=%s =redistribute-rip=%s =redistribute-static=%s =router-id=%s =routing-table=%s", as, name, boolToMikrotikBool(clientToClientReflection), clusterId, comment, confederation, confederationPeers, boolToMikrotikBool(disabled), boolToMikrotikBool(ignoreAsPathLen), outFilter, boolToMikrotikBool(redistributeConnected), boolToMikrotikBool(redistributeOspf), boolToMikrotikBool(redistributeOtherBgp), boolToMikrotikBool(redistributeRip), boolToMikrotikBool(redistributeStatic), routerId, routingTable), " ")

	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)
	log.Printf("[DEBUG] /routing/bpg/instance/add returned %v", r)

	if err != nil {
		return nil, err
	}

	return client.FindBgpInstance(name)
}

func (client Mikrotik) FindBgpInstance(name string) (*BgpInstance, error) {
	c, err := client.getMikrotikClient()
	if err != nil {
		return nil, err
	}

	cmd := strings.Split(fmt.Sprintf("/routing/bgp/instance print where name=%s", name), " ")
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] Find bgp instance: `%s`", cmd)

	bgpInstance := BgpInstance{}

	err = Unmarshal(*r, &bgpInstance)

	if err != nil {
		return nil, err
	}

	if bgpInstance.Name == "" {
		return nil, NewNotFound(fmt.Sprintf("bgp instance `%s` not found", name))
	}

	return &bgpInstance, nil
}
