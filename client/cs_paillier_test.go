/*
 * Copyright 2017 XLAB d.o.o.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package client

import (
	"math/big"
	"testing"

	"path/filepath"

	"github.com/stretchr/testify/assert"
	"github.com/xlab-si/emmy/config"
	"github.com/xlab-si/emmy/crypto/common"
)

func testCSPaillier(m, l *big.Int, pubKeyPath string) error {
	c, err := NewCSPaillierClient(testGrpcClientConn, pubKeyPath, m, l)
	if err != nil {
		return err
	}
	return c.Run()
}

func TestCSPaillier(t *testing.T) {
	m := common.GetRandomInt(big.NewInt(8685849))
	l := common.GetRandomInt(big.NewInt(340002223232))
	pubKeyPath := filepath.Join(config.LoadTestdataDir(), "cspaillierpubkey.txt")

	assert.Nil(t, testCSPaillier(m, l, pubKeyPath), "should finish without error")
}
