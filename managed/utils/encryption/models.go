// Copyright (C) 2024 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package encryption

import "github.com/google/tink/go/tink"

type Encryption struct {
	Path      string
	Key       string
	Primitive tink.AEAD
}

type DatabaseConnection struct {
	Host, User, Password string
	Port                 int16
	DBName               string
	SSLMode              string
	SSLCAPath            string
	SSLKeyPath           string
	SSLCertPath          string
	EncryptedItems       []EncryptedItem
}

type EncryptedItem struct {
	Database, Table string
	Identificators  []string
	Columns         []string
}

type QueryValues struct {
	Query       string
	SetValues   [][]any
	WhereValues [][]any
}
