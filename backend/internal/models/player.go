package models

import (
	"time"

	"github.com/dev-xero/chess.io/internal/utils"
	"github.com/gofiber/fiber/v3/log"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Player database model.
type Player struct {
	Id             uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"player_id"`
	Username       string         `gorm:"uniqueIndex;not null" json:"username"`
	Bio            string         `gorm:"not null" json:"bio"`
	PasswordHash   string         `gorm:"not null" json:"-"`
	SecretQuestion string         `gorm:"not null" json:"secret_question"`
	EloRating      int            `gorm:"not null" json:"elo_rating"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Player) TableName() string {
	return "players"
}

// BeforeCreate is a hook that runs right before gorm creates a new player record.
//
// This hook takes in the password received from the web client, hashes that, then
// encodes it as a string using Argon2 helper methods.
func (p *Player) BeforeCreate(tx *gorm.DB) (err error) {
	params := &utils.Argon2Params{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   16,
	}
	encodedPasswordHash, err := utils.GenerateEncodedPasswordHash(p.PasswordHash, params)
	if err != nil {
		log.Fatal("Failed to generate encoded password hash from the Argon2 utility method: ", err)
	}
	p.PasswordHash = encodedPasswordHash
	return
}
