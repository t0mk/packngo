package metadata

import (
	"fmt"
	"net"
	"time"
)

type CurrentDevice struct {
	ID             string          `json:"id"`
	Hostname       string          `json:"hostname"`
	IQN            string          `json:"iqn"`
	OS             OperatingSystem `json:"operating_system"`
	Plan           string          `json:"plan"`
	Reserved       bool            `json:"reserved"`
	Class          string          `json:"class"`
	Facility       string          `json:"facility"`
	PrivateSubnets []string        `json:"private_subnets"`
	Tags           []string        `json:"tags"`
	SSHKeys        []string        `json:"ssh_keys"`
	Storage        Storage         `json:"storage"`
	Network        Network         `json:"network"`
	CustomData     interface{}     `json:"customdata"`
	Specs          Specs           `json:"specs"`
	SwitchShortID  string          `json:"switch_short_id"`
	Volumes        []Volume        `json:"volumes"`
	APIURL         string          `json:"api_url"`
	PhoneHomeURL   string          `json:"phone_home_url"`
	UserStateURL   string          `json:"user_state_url"`
}

type AddressFamily int

const (
	IPv4 = AddressFamily(4)
	IPv6 = AddressFamily(6)
)

type BondingMode int

const (
	BondingBalanceRR    = BondingMode(0)
	BondingActiveBackup = BondingMode(1)
	BondingBalanceXOR   = BondingMode(2)
	BondingBroadcast    = BondingMode(3)
	BondingLACP         = BondingMode(4)
	BondingBalanceTLB   = BondingMode(5)
	BondingBalanceALB   = BondingMode(6)
)

var bondingModeStrings = map[BondingMode]string{
	BondingBalanceRR:    "balance-rr",
	BondingActiveBackup: "active-backup",
	BondingBalanceXOR:   "balance-xor",
	BondingBroadcast:    "broadcast",
	BondingLACP:         "802.3ad",
	BondingBalanceTLB:   "balance-tlb",
	BondingBalanceALB:   "balance-alb",
}

func (m BondingMode) String() string {
	if str, ok := bondingModeStrings[m]; ok {
		return str
	}
	return fmt.Sprintf("%d", m)
}

type Capacity struct {
	Size int    `json:"size,string"`
	Unit string `json:"unit"`
}
type Volume struct {
	IPs      []net.IP `json:"ips"`
	Name     string   `json:"name"`
	Capacity Capacity `json:"capacity"`
	IQN      string   `json:"iqn"`
}

type LicenseActivation struct {
	State string `json:"state"`
}
type OperatingSystem struct {
	Slug              string            `json:"slug"`
	Distro            string            `json:"distro"`
	Version           string            `json:"version"`
	LicenseActivation LicenseActivation `json:"license_activation"`
	ImageTag          string            `json:"image_tag"`
}
type Partitions struct {
	Label  string `json:"label"`
	Number int    `json:"number"`
	// Partition size is sometimes str and sometimes int
	/*
	 "storage": {
	    "disks": [
	      {
	        "device": "/dev/sda",
	        "wipeTable": true,
	        "partitions": [
	          {
	            "label": "BIOS",
	            "number": 1,
	            "size": 4096
	          },
	          {
	            "label": "SWAP",
	            "number": 2,
	            "size": "3993600"
	          },
	          {
	            "label": "ROOT",
	            "number": 3,
	            "size": 0
	          }
	        ]
	      }
	    ],

	*/
	//Size   int    `json:"size"`
}
type Disks struct {
	Device     string       `json:"device"`
	WipeTable  bool         `json:"wipeTable"`
	Partitions []Partitions `json:"partitions"`
}
type Create struct {
	Options []string `json:"options"`
}
type Mount struct {
	Device string `json:"device"`
	Format string `json:"format"`
	Point  string `json:"point"`
	Create Create `json:"create"`
}
type Filesystems struct {
	Mount Mount `json:"mount"`
}
type Storage struct {
	Disks       []Disks       `json:"disks"`
	Filesystems []Filesystems `json:"filesystems"`
}
type Bonding struct {
	Mode            BondingMode `json:"mode"`
	LinkAggregation string      `json:"link_aggregation"`
	MAC             string      `json:"mac"`
}

type InterfaceInfo struct {
	Name string `json:"name"`
	MAC  string `json:"mac"`
	Bond string `json:"bond"`
}

func (i *InterfaceInfo) ParseMAC() (net.HardwareAddr, error) {
	return net.ParseMAC(i.MAC)
}

type ParentBlock struct {
	Network string `json:"network"`
	Netmask string `json:"netmask"`
	Cidr    int    `json:"cidr"`
	Href    string `json:"href"`
}
type Addresses struct {
	ID            string        `json:"id"`
	AddressFamily AddressFamily `json:"address_family"`
	Netmask       net.IP        `json:"netmask"`
	CreatedAt     time.Time     `json:"created_at"`
	Public        bool          `json:"public"`
	Cidr          int           `json:"cidr"`
	Management    bool          `json:"management"`
	Enabled       bool          `json:"enabled"`
	Network       net.IP        `json:"network"`
	Address       net.IP        `json:"address"`
	Gateway       net.IP        `json:"gateway"`
	ParentBlock   ParentBlock   `json:"parent_block"`
}
type Network struct {
	Bonding    Bonding         `json:"bonding"`
	Interfaces []InterfaceInfo `json:"interfaces"`
	Addresses  []Addresses     `json:"addresses"`
}

func (n *Network) BondingMode() BondingMode {
	return n.Bonding.Mode
}

type Cpus struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

type Memory struct {
	Total string `json:"total"`
}

type Drives struct {
	Count int    `json:"count"`
	Size  string `json:"size"`
	Type  string `json:"type"`
}

type Nics struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

type Features struct {
	Raid bool `json:"raid"`
	Txt  bool `json:"txt"`
	Uefi bool `json:"uefi"`
}

type Specs struct {
	Cpus     []Cpus   `json:"cpus"`
	Memory   Memory   `json:"memory"`
	Drives   []Drives `json:"drives"`
	Nics     []Nics   `json:"nics"`
	Features Features `json:"features"`
}
