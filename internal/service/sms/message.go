// See the file LICENSE for redistribution and license information.
//
// Copyright (c) 2020 worldiety. All rights reserved.
// DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.
//
// Please contact worldiety, Marie-Curie-Straße 1, 26129 Oldenburg, Germany
// or visit www.worldiety.com if you need more information or have any questions.
//
// Authors: Torben Schinke

// +build !wasm

package sms

import (
	"context"
	"github.com/golangee/uuid"
	"time"
)

type SMS struct {
	ID        uuid.UUID `ee.sql.Name:"id"`
	CreatedAt time.Time `ee.sql.Name:"created_at"`
	Text      string    `ee.sql.Name:"text"`
}

/**
@ee.sql.Schema("""
	{
		"dialect":"mysql"
	}
	CREATE TABLE IF NOT EXISTS `sms` (
	  `id` BINARY(16) NOT NULL COMMENT 'an uuid',
	  `provider` BINARY(16) NOT NULL COMMENT 'uuid of the provider to sent the message with',
	  `recipient` VARCHAR(255) NOT NULL COMMENT 'the phone number to send to',
	  `text` TEXT NOT NULL COMMENT 'SMS text is actually limited to 160 chars per message but larger messages can be joined from an arbitrary amount.',
	  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'the time when this sms has been inserted',
	  `send_at` TIMESTAMP NOT NULL DEFAULT 0 COMMENT '0 if never sent',
	  `status` ENUM("unknown", "success", "failed") NOT NULL DEFAULT 'unknown' COMMENT 'default status is unknown',
	  `details` JSON NOT NULL DEFAULT '{}' COMMENT 'contains arbitrary status details',
	  PRIMARY KEY (`id`),
	  INDEX `fk_sms_sms_provider_idx` (`provider` ASC) VISIBLE,
	  CONSTRAINT `fk_sms_sms_provider`
		FOREIGN KEY (`provider`)
		REFERENCES `mercurius`.`sms_provider` (`id`)
		ON DELETE CASCADE
		ON UPDATE NO ACTION)
	ENGINE = InnoDB
	COMMENT = 'contains all sms which needs to be sent and which have been sent. Probably already sent sms may be removed over time.'
""")

@ee.stereotype.Repository("sms")
 */
type MessageRepository interface {
	// @ee.sql.Query("SELECT id,created_at,text FROM sms LIMIT :limit")
	FindAll(ctx context.Context, limit int) ([]SMS, error)

	// @ee.sql.Query("SELECT id,created_at,text FROM sms WHERE id = :id")
	FindById(ctx context.Context, id uuid.UUID) (SMS, error)

	// @ee.sql.Query("INSERT INTO sms (id,recipient,text) VALUES (:uuid, :recipient, :text)")
	Create(ctx context.Context, uuid uuid.UUID, recipient string, text string) error
}
