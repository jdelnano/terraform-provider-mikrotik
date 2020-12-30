package client

import (
	"fmt"
	"log"
	"strings"
)

// BgpInstance Mikrotik resource
type BgpInstance struct {
	ID                       string `mikrotik:".id"`
	Name                     string
	As                       int    `mikroktik:"as"`
	ClientToClientReflection bool   `mikroktik:"client-to-client-reflection"`
	Comment                  string `mikrotik:"comment"`
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
	RouterID              string `mikroktik:"router-id"`
	RoutingTable          string `mikroktik:"routing-table"`
	ClusterID             string `mikroktik:"cluster-id"`
	Confederation         int    `mikrotik:"confederation"`
}

// AddBgpInstance Mikrotik resource
func (client Mikrotik) AddBgpInstance(name string, as int, clientToClientReflection bool, comment string, confederationPeers string, disabled bool, ignoreAsPathLen bool, outFilter string, redistributeConnected bool, redistributeOspf bool, redistributeOtherBgp bool, redistributeRip bool, redistributeStatic bool, routerID string, routingTable string, clusterID string, confederation int) (*BgpInstance, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := strings.Split(fmt.Sprintf("/routing/bgp/instance/add =as=%d =name=%s =client-to-client-reflection=%s =comment=%s =confederation-peers=%s =disabled=%s =ignore-as-path-len=%s =out-filter=%s =redistribute-connected=%s =redistribute-ospf=%s =redistribute-other-bgp=%s =redistribute-rip=%s =redistribute-static=%s =router-id=%s =routing-table=%s", as, name, boolToMikrotikBool(clientToClientReflection), comment, confederationPeers, boolToMikrotikBool(disabled), boolToMikrotikBool(ignoreAsPathLen), outFilter, boolToMikrotikBool(redistributeConnected), boolToMikrotikBool(redistributeOspf), boolToMikrotikBool(redistributeOtherBgp), boolToMikrotikBool(redistributeRip), boolToMikrotikBool(redistributeStatic), routerID, routingTable), " ")

	// optionally append fields if they are set
	// cannot include them empty otherwise Mikrotik has fit
	if confederation != -1 {
		cmd = append(cmd, fmt.Sprintf("=confederation=%d", confederation))
	}
	if clusterID != "" {
		cmd = append(cmd, fmt.Sprintf("=cluster-id=%s", clusterID))
	}

	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)
	log.Printf("[DEBUG] /routing/bgp/instance/add returned %v", r)

	if err != nil {
		return nil, err
	}

	return client.FindBgpInstance(name)
}

// FindBgpInstance Mikrotik resource
func (client Mikrotik) FindBgpInstance(name string) (*BgpInstance, error) {
	c, err := client.getMikrotikClient()
	if err != nil {
		return nil, err
	}

	cmd := strings.Split(fmt.Sprintf("/routing/bgp/instance/print ?name=%s", name), " ")
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] Find bgp instance: `%v`", cmd)

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

// UpdateBgpInstance Mikrotik resource
func (client Mikrotik) UpdateBgpInstance(name string, as int, clientToClientReflection bool, comment string, confederationPeers string, disabled bool, ignoreAsPathLen bool, outFilter string, redistributeConnected bool, redistributeOspf bool, redistributeOtherBgp bool, redistributeRip bool, redistributeStatic bool, routerID string, routingTable string, clusterID string, confederation int) (*BgpInstance, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	bgpInstance, err := client.FindBgpInstance(name)

	if err != nil {
		return bgpInstance, err
	}
	cmd := strings.Split(fmt.Sprintf("/routing/bgp/instance/set =numbers=%s =as=%d =name=%s =client-to-client-reflection=%s =comment=%s =confederation-peers=%s =disabled=%s =ignore-as-path-len=%s =out-filter=%s =redistribute-connected=%s =redistribute-ospf=%s =redistribute-other-bgp=%s =redistribute-rip=%s =redistribute-static=%s =router-id=%s =routing-table=%s", bgpInstance.ID, as, name, boolToMikrotikBool(clientToClientReflection), comment, confederationPeers, boolToMikrotikBool(disabled), boolToMikrotikBool(ignoreAsPathLen), outFilter, boolToMikrotikBool(redistributeConnected), boolToMikrotikBool(redistributeOspf), boolToMikrotikBool(redistributeOtherBgp), boolToMikrotikBool(redistributeRip), boolToMikrotikBool(redistributeStatic), routerID, routingTable), " ")

	// optionally append fields if they are set
	// cannot include them empty otherwise Mikrotik has fit
	// TODO:  run `unset` command
	if confederation != -1 {
		cmd = append(cmd, fmt.Sprintf("=confederation=%d", confederation))
	}
	// TODO: `unset` doesn't work :(
	if clusterID != "" {
		cmd = append(cmd, fmt.Sprintf("=cluster-id=%s", clusterID))
	}

	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	_, err = c.RunArgs(cmd)

	if err != nil {
		return nil, err
	}

	return client.FindBgpInstance(name)
}

// DeleteBgpInstance Mikrotik resource
func (client Mikrotik) DeleteBgpInstance(name string) error {
	c, err := client.getMikrotikClient()

	bgpInstance, err := client.FindBgpInstance(name)

	if err != nil {
		return err
	}

	cmd := strings.Split(fmt.Sprintf("/routing/bgp/instance/remove =numbers=%s", bgpInstance.Name), " ")
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)
	log.Printf("[DEBUG] Remove bgp instance via mikrotik api: %v", r)

	return err
}
