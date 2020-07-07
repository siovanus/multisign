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

package multisign

import (
	"fmt"
	"time"

	"github.com/siovanus/multisign/common"
	"github.com/siovanus/multisign/config"
	sdk "github.com/ontio/ontology-go-sdk"
)

func MakeAuthorizeTxAndSign(ontSdk *sdk.OntologySdk) bool {
	time.Sleep(1 * time.Second)
	user, ok := common.GetAccountByPassword(ontSdk, config.DefConfig.WalletPath)
	if !ok {
		return false
	}
	if len(config.DefConfig.PublicKeyList) == 0 || len(config.DefConfig.PosList) == 0 {
		fmt.Println("PublicKeyList or PosList is not provided")
		return false
	}
	ok = makeAuthorizeTxAndSign(ontSdk, user, config.DefConfig.PublicKeyList, config.DefConfig.PosList)
	if !ok {
		return false
	}
	return true
}

func Sign(ontSdk *sdk.OntologySdk) bool {
	time.Sleep(1 * time.Second)
	user, ok := common.GetAccountByPassword(ontSdk, config.DefConfig.WalletPath)
	if !ok {
		return false
	}
	ok = sign(ontSdk, user)
	if !ok {
		return false
	}
	return true
}

func SignAndSend(ontSdk *sdk.OntologySdk) bool {
	time.Sleep(1 * time.Second)
	user, ok := common.GetAccountByPassword(ontSdk, config.DefConfig.WalletPath)
	if !ok {
		return false
	}
	ok = signAndSend(ontSdk, user)
	if !ok {
		return false
	}
	common.WaitForBlock(ontSdk)
	return true
}
