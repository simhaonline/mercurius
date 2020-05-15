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
	CREATE TABLE IF NOT EXISTS "sms" (
	  "id" BINARY(16) NOT NULL,
	  "recipient" VARCHAR(255) NOT NULL COMMENT 'the phone number to send to',
	  "text" TEXT NOT NULL COMMENT 'SMS text is limited to 160 chars (non multibyte?) per message but can be joined from an arbitrary amount of sms messages.',
	  "created_at" TIMESTAMP NOT NULL,
	  "send_at" TIMESTAMP NOT NULL DEFAULT 0,
	  "status" ENUM('unknown', 'success', 'failed') NOT NULL DEFAULT 'unknown',
	  "details" JSON NOT NULL DEFAULT '{}' COMMENT 'contains arbitrary status details',
	  PRIMARY KEY ("id"))
	ENGINE = InnoDB;
""")

@ee.stereotype.Repository("sms")
 */
type Repository interface {
	// @ee.sql.Query("SELECT id,created_at,text FROM sms LIMIT :limit")
	FindAll(ctx context.Context, limit int) ([]SMS, error)

	// @ee.sql.Query("SELECT id,created_at,text FROM sms WHERE id = :id")
	FindById(ctx context.Context, id uuid.UUID) (SMS, error)

	// @ee.sql.Query("INSERT INTO sms (id,recipient,text) VALUES (:uuid, :recipient, :text)")
	Create(ctx context.Context, uuid uuid.UUID, recipient string, text string) error
}
