package client

type FirefallFilter struct {
	Id     string `mikrotik:".id"`
	Action string `mikrotik:action`
	//AddressListTimeout string `mikroktik:"address-list-timeout"` // not going to include at first
	Chain   string `mikrotik:chain`
	Comment string `mikrotik:"comment"`
	//ConnectionBytes    string `mikroktik:"connection-bytes"`     // not going to include at first
	//ConnectionLimit    string `mikroktik:"connection-limit"`     // not going to include at first
	//ConnectionMark     string `mikroktik:"connection-mark"`      // not going to include at first
	//ConnectionNatState string `mikroktik:"connection-nat-state"` // not going to include at first
	//ConnectionRate     int    `mikroktik:"connection-rate"`      // not going to include at first
	ConnectionState string `mikroktik:"connection-state"`
	//ConnectionType         string `mikroktik:"connection-type"` // not going to include at first
	Content                string `mikroktik:"content"` // not going to include at first
	Discp                  int    `mikrotik:"discp"`
	DestinationAddress     string `mikrotik:"dst-address"`
	DestinationAddressList string `mikrotik:"dst-address-list"`
	DestinationAddressType string `mikrotik:"dst-address-type"`
	DestinationLimit       string `mikrotik:"dst-limit"`
	DestinationPort        string `mikrotik:"dst-port"`
	//IcmpOptions string `mikrotik:"icmp-options"` // not going to include at first
	InBridgePort      string `mikrotik:"in-bridge-port"`
	InBridgePortList  string `mikrotik:"in-bridge-port-list"`
	InInterface       string `mikrotik:"in-interface"`
	InInterfaceList   string `mikrotik:"in-interface-list"`
	IpsecPolicy       string `mikrotik:"ipsec-policy"`
	Ip4Options        string `mikrotik:"ip4-options"`
	Limit             string `mikrotik:"limit"`
	Nth               string `mikrotik:"nth"`
	OutBridgePort     string `mikrotik:"in-bridge-port"`
	OutBridgePortList string `mikrotik:"in-bridge-port-list"`
	OutInterface      string `mikrotik:"out-interface"`
	OutInterfaceList  string `mikrotik:"out-interface-list"`
	Port              string `mikrotik:"port"`
	Priority          string `mikrotik:"priority"`
	Protocol          string `mikrotik:"protocol"`
	//RoutingTable string `mikrotik:"routing-table"` // not going to include at first
	//RoutingMark string `mikrotik:"routing-mark"` // not going to include at first
	SourceAddress     string `mikrotik:"src-address"`
	SourceAddressList string `mikrotik:"src-address-list"`
	SourceAddressType string `mikrotik:"src-address-type"`
	SourcePort        string `mikrotik:"src-port"`
	//SourceMacAddress string `mikrotik:"src-mac-address"` // not going to include at first
	Ttl int `mikrotik:"ttl"`
}
