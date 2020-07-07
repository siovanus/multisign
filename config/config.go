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

//common use fot ontology-tool
package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Default config instance
var DefConfig = NewConfig()

var DefPubKeysGroup = NewPubKeysGroup()

type PubKeysGroup struct {
	PubKeysGroup [][]string
}

//Config object used by ontology-instance
type Config struct {
	//JsonRpcAddress of ontology
	JsonRpcAddress string
	WalletPath     string
	//Gas Price of transaction
	GasPrice uint64
	//Gas Limit of invoke transaction
	GasLimit      uint64
	PublicKeyList []string
	PosList       []uint32
}

//NewConfig return a Config instance
func NewConfig() *Config {
	return &Config{}
}

//Init Config with a config file
func (this *Config) Init(fileName string) error {
	err := this.loadConfig(fileName)
	if err != nil {
		return fmt.Errorf("loadConfig error:%s", err)
	}
	return nil
}

func (this *Config) loadConfig(fileName string) error {
	data, err := ReadFile(fileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, this)
	if err != nil {
		return fmt.Errorf("json.Unmarshal TestConfig:%s error:%s", data, err)
	}
	return nil
}

func ReadFile(fileName string) ([]byte, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("OpenFile %s error %s", fileName, err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("File %s close error %s", fileName, err)
		}
	}()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll %s error %s", fileName, err)
	}
	return data, nil
}

func NewPubKeysGroup() *PubKeysGroup {
	return &PubKeysGroup{}
}

func (this *PubKeysGroup) Init(fileName string) error {
	err := this.loadPubKeysGroup(fileName)
	if err != nil {
		return fmt.Errorf("loadPubKeysGroup error:%s", err)
	}
	return nil
}

func (this *PubKeysGroup) loadPubKeysGroup(fileName string) error {
	data, err := ReadFile(fileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, this)
	if err != nil {
		return fmt.Errorf("json.Unmarshal TestConfig:%s error:%s", data, err)
	}
	return nil
}
