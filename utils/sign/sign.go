/*********************************************************************

rss3go_hub: An alternative version of RSS3-Hub written in go

Copyright (C) 2021 Nyawork, Candinya

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

 ********************************************************************/

package sign

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetSigner(msg string, signHex string) (string, error) {

	msgHash := crypto.Keccak256Hash([]byte(msg))

	signatureReceived, err := hex.DecodeString(signHex)
	if err != nil {
		return "", err
	}

	recoveredPubBytes, err := crypto.Ecrecover(msgHash.Bytes(), signatureReceived)
	if err != nil {
		return "", err
	}

	recoveredPub, err := crypto.UnmarshalPubkey(recoveredPubBytes)
	if err != nil {
		return "", err
	}

	signer := crypto.PubkeyToAddress(*recoveredPub).String()

	return signer, nil

}
