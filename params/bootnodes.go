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
	"enode://7bed18c87054f807bc9096501bc78f737363f357af831791bab07c4fa6c5a1a67cdcf0a097dc2cc918262ef04fb1c05c26026df5c11a6a56666f9b1fb4072210@18.178.30.66:32668",
	"enode://d67251dd3b050e555679a8abdc427a4c78a9bae174f2fd3b9163c364d27b6a69688ee067cd3214e8ceb71e6e602fd812797b085ae37ed3bf93b78e2b77ae3306@18.181.40.7:32668",
	"enode://f88bb1f5d0e42cf75ec879212b7c8477d605315d5296fba02bc4600eccf73c64427de46567a320d00985d5bc612168817ba6dff169bd6a4774e112e6db0ff6a2@18.176.66.118:32668",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
var TestnetBootnodes = []string{
	"enode://f70c9040dfb58507a78dcd87e9fd183fa82fa56b394b2969b2a3b19093088319b487c31e5154cd89a4af4b95c80c42ae432d737f5ae15650b63a7425c539fdf2@13.230.35.234:33668",
	"enode://4a15647b36f2c5e65dbe74976223903d0b966787aa9310fb792b48ca2eef64750f826b84a095db392b3bb650052c2386b357e654370e7d11572d109da09bd10b@35.73.127.28:33668",
	"enode://367fa1287885b4832a49bddfb941b6ef7defcdd805d67aa91c3261cbd83aa726fe70333632a5611b407dc76f46247039a93981aab9f204f3c572c1fe0f89721a@18.181.232.158:33668",
}

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
