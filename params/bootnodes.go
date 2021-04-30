// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes//TODO
	"enode://f0f6efd7ef4fa12816c41629bb992f456f406630f7e1f486cd1091b11d307423d1f5920402d5faec13c0a57246dd62352704845b2acf3e4746972823daa015d7@35.73.253.58:32668",
	"enode://83081cb37c3fc39c8c7cf774cc55dbde11b1c92512551bd5dcc11fce7cb13e54a17798fd236642ccd76f90a32a418d86c5a1ad0ea1a759392c7d8f746a501f95@52.192.142.14:32668",
	"enode://4cbd81edde76f560f734fe6471c7620eb4dc693a2db1e72401b9f09efd89b5711f7f545d833136f6c6cbfae0050b9726adac384605580bac3b799e574dd3c95e@3.115.0.208:32668",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
var TestnetBootnodes = []string{
	"enode://160ae2ec30b079ace074cfa28c73757f2493fd7afd807e0da901e68eff31ead944ff6eec7b854dfffffb1b4d4cc886829455e62f8422edd937694498544f9c40@13.230.35.234:33668",
	"enode://d4718eb176c63c4752adeaf050df31b93a85de7e37e69b82fbe7a6f28f4e556c976a227250367a9e9ffd2e08b6d513448f7dff4969e4afaabb170b71eac2907b@35.73.127.28:33668",
	"enode://bfe0bdaaeed5767f8a79204f18a45afc5a802eed6428b2c0aebf757f3a98276268090d5b6f1f839330f6b0cd397e22ad4f99affefd7a132421d43446c0a3f8a8@18.181.232.158:33668",
}

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
