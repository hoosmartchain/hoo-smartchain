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
	"enode://eb565f4c270e45ee13aaf4e27e452ee11591c2dc4ab795a8987f705a50fd3604f751f7c0d8b669287b2ba1b3713effea13bb0c25d245066b1a703bd0c477f86a@35.73.253.58:32668",
	"enode://98eda17a83decd4a9b606b926f12b4c965b7b6b851a277e895861338eb7eebae8d4b48523eee33fafee5ad9a5f8d3fb14e56ad75ea70e80de815dc98962d9904@52.192.142.14:32668",
	"enode://d7ee8747aa8a0e4ab61ef27e18045a36014f1458a018394f587411d4959c9cf41742bd1402f693e9eb8a76b9fdc4e4d12b5a49c0000e7b08092e7f357b72a4c1@3.115.0.208:32668",
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
