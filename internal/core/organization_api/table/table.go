// Package table manages all of the Dynamo calls (query, scan, get, write, etc).
package table

/**
 * Panther is a scalable, powerful, cloud-native SIEM written in Golang/React.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

	"github.com/panther-labs/panther/api/lambda/organization/models"
)

// General settings are stored in a single Dynamo row with this key.
var settingsKey = DynamoItem{"id": {S: aws.String("generalSettings")}}

// API defines the interface for the table which can be used for mocking.
type API interface {
	Get() (*models.GeneralSettings, error)
	Update(settings *models.GeneralSettings) (*models.GeneralSettings, error)
}

// OrganizationsTable encapsulates a connection to the Dynamo table.
type OrganizationsTable struct {
	Name   *string
	client dynamodbiface.DynamoDBAPI
}

// The OrganizationsTable must satisfy the API interface.
var _ API = (*OrganizationsTable)(nil)

// New creates a new Dynamo client which talks to the given table name.
func New(tableName string, sess *session.Session) *OrganizationsTable {
	return &OrganizationsTable{Name: aws.String(tableName), client: dynamodb.New(sess)}
}

// DynamoItem is a type alias for the item format expected by the Dynamo SDK.
type DynamoItem = map[string]*dynamodb.AttributeValue
