/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */
package main

import (
	"flag"
	"fmt"
	"github.com/siovanus/multisign/config"
	"github.com/siovanus/multisign/core"
	_ "github.com/siovanus/multisign/methods"
)

var (
	Method string //Method
)

func init() {
	flag.StringVar(&Method, "t", "", "MakeAuthorizeTxAndSign: read input data from config.json to make authorizeForPeer tx and sign it, write signed tx to tx.txt file \n"+
		"MakeUnAuthorizeTxAndSign: read input data from config.json to make unAuthorizeForPeer tx and sign it, write signed tx to tx.txt file \n"+
		"MakeWithdrawTxAndSign: read input data from config.json to make withdraw tx and sign it, write signed tx to tx.txt file \n"+
		"Sign: read raw tx from tx.txt and sign it \n"+
		"SignAndSend: read raw tx from tx.txt, sign it and send it to ontology network configured in config.json")
	flag.Parse()
}

func main() {
	err := config.DefConfig.Init("./config.json")
	if err != nil {
		fmt.Println("DefConfig.Init error:", err)
		return
	}

	err = config.DefPubKeysGroup.Init("./pubKeysGroup.json")
	if err != nil {
		fmt.Println("DefPubKeysGroup.Init error:", err)
		return
	}

	core.OntTool.Start(Method)
}
