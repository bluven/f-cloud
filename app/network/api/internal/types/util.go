package types

import (
	"github.com/bluven/f-cloud/app/network/model"
)

func FromNetwork(network model.Network) *Network {
	return &Network{
		ID:         network.ID,
		Name:       network.Name,
		IPv4Addr:   network.IPv4Addr,
		Bandwidth:  network.Bandwidth,
		Traffic:    network.Traffic,
		InstanceID: network.InstanceID,
		CreatedAt:  network.CreatedAt.Unix(),
		UpdatedAt:  network.UpdatedAt.Unix(),
	}
}

func FromLoadBalancer(lb model.LoadBalancer) *LoadBalancer {
	var network *Network
	if lb.Network != nil {
		network = FromNetwork(*lb.Network)
	}
	return &LoadBalancer{
		ID:        lb.ID,
		Name:      lb.Name,
		NetworkID: lb.NetworkID,
		Network:   network,
		Algorithm: lb.Algorithm,
		CreatedAt: lb.CreatedAt.Unix(),
		UpdatedAt: lb.UpdatedAt.Unix(),
	}
}
