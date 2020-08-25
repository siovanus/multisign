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

package core

import (
	"fmt"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/siovanus/multisign/config"
)

var OntTool = NewOntologyTool()

type Method func(sdk *sdk.OntologySdk) bool

type OntologyTool struct {
	//Map name to method
	methodsMap map[string]Method
}

func NewOntologyTool() *OntologyTool {
	return &OntologyTool{
		methodsMap: make(map[string]Method, 0),
	}
}

func (this *OntologyTool) RegMethod(name string, method Method) {
	this.methodsMap[name] = method
}

//Start run
func (this *OntologyTool) Start(method string) {
	if method != "" {
		this.runMethod(method)
		return
	}
	fmt.Println("No method to run")
	return
}

func (this *OntologyTool) runMethod(method string) {
	this.onStart()
	ontSdk := sdk.NewOntologySdk()
	ontSdk.NewRpcClient().SetAddress(config.DefConfig.JsonRpcAddress)
	function := this.getMethodByName(method)
	ok := function(ontSdk)
	if ok {
		fmt.Println("---------------------------------------------------------------")
		fmt.Println("Success:", method)
	}
	if !ok {
		fmt.Println("---------------------------------------------------------------")
		fmt.Println("Failed:", method)
	}
}

func (this *OntologyTool) onStart() {
	fmt.Println("===============================================================")
	fmt.Println("-------Ontology Tool Start-------")
	fmt.Println("===============================================================")
	fmt.Println("")
}

func (this *OntologyTool) getMethodByName(name string) Method {
	return this.methodsMap[name]
}
