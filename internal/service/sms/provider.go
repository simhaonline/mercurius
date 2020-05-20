// See the file LICENSE for redistribution and license information.
//
// Copyright (c) 2020 worldiety. All rights reserved.
// DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.
//
// Please contact worldiety, Marie-Curie-Straße 1, 26129 Oldenburg, Germany
// or visit www.worldiety.com if you need more information or have any questions.
//
// Authors: Torben Schinke

package sms

import (
	"context"
	"github.com/golangee/uuid"
)

// ProviderType tells which of the sms provider implementation should be used.
type ProviderType int

const (
	Sipgate ProviderType = 1
	API     ProviderType = 2
)

// A Provider is a concrete configuration to deliver sms messages.
type Provider struct {
	ID           uuid.UUID    `ee.sql.Name:"id"`
	Name         string       `ee.sql.Name:"name"`
	Text         string       `ee.sql.Name:"text"`
	Host         string       `ee.sql.Name:"host"`
	Type         ProviderType `ee.sql.Name:"type"`
	AccessToken  string       `ee.sql.Name:"access_token"`
	RefreshToken string       `ee.sql.Name:"refresh_token"`
}

/**
A ProviderRepository defines the repository contract to sms providers.

	@ee.sql.Schema("""
		{
			"dialect":"mysql"
		}
		CREATE TABLE IF NOT EXISTS `sms_provider` (
		  `id` BINARY(16) NOT NULL COMMENT 'uuid of the provider',
		  `type` ENUM('sipgate', 'api') NOT NULL COMMENT 'there are only 2 build-in providers. Sipgate or a custom web api',
		  `name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'human readable name of the provider configuration',
		  `host` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'the url to use as host, may be empty e.g. for sipgate',
		  `access_token` VARCHAR(2048) NOT NULL DEFAULT '' COMMENT 'the latest oauth2 access token or another kind of api token if not oauth2\n',
		  `refresh_token` VARCHAR(45) NOT NULL DEFAULT '' COMMENT 'the latest refresh token',
		  PRIMARY KEY (`id`))
		ENGINE = InnoDB
		COMMENT = 'sms_provider contains all configured senders. A sms must be sent using a specific provider.'
	""")

	@ee.stereotype.Repository("sms")
*/
type ProviderRepository interface {
	// @ee.sql.Query(SELECT "id", "name", "type", "host", "access_token", "refresh_token" FROM "sms_provider")
	FindAll(ctx context.Context) ([]Provider, error)

	// @ee.sql.Query(SELECT "id", "name", "type", "host", "access_token", "refresh_token" FROM "sms_provider" WHERE "id" = :id)
	FindById(ctx context.Context, id uuid.UUID) (SMS, error)

	// @ee.sql.Query(INSERT INTO "sms_provider" ("id", "type") VALUES (:id, :kind) )
	Create(ctx context.Context, id uuid.UUID, kind ProviderType) error

	// @ee.sql.Query("UPDATE INTO "sms_provider" SET "name" = :name, "host" = :host, "access_token" = :accessToken, "refresh_token" = :refreshToken WHERE "id" = :id)
	Update(ctx context.Context, id uuid.UUID, name, host, accessToken, refreshToken string) error

	// @ee.sql.Query(DELETE FROM "sms_provider" WHERE "id" = :id )
	Delete(ctx context.Context, id uuid.UUID)error
}
