package control

import (
	"database/sql"
	"time"
)

type Profile struct {
	Name        string
	Gender      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DateOfBirth time.Time
	State       string
}

type User struct {
	ID          int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	VerifiedAt  time.Time
	DestroyAt   time.Time
	Filter      string
	Role        string
	Banned      bool
	Vip         bool
	InFilter    string
	Name        string
	Gender      int
	DateOfBirth time.Time
	BannedTo    time.Time
	RealPerson  bool

	CurrentProfile   Profile
	CurrentProfileId sql.NullInt64
	Profiles         []Profile
}

type Call struct {
	ID                  int
	SourceAnswer        bool
	DestinationAnswer   bool
	Incognito           bool
	SourceReveal        bool
	DestinationReveal   bool
	SourceAccept        bool
	DestinationAccept   bool
	Status              int
	CreatedAt           time.Time
	UpdatedAt           time.Time
	FinishedAt          time.Time
	SourceAnswerAt      time.Time
	DestinationAnswerAt time.Time
	SourceRevealAt      time.Time
	DestinationRevealAt time.Time
	AcceptedAt          time.Time
	RejectedAt          time.Time
	Type                string
	Destination         User `sql:"-"`
	DestinationID       int
	SourceID            int
	Source              User `sql:"-"`
	callTimer           time.Timer
}

type Message struct {
	Type          string `mapstructure:"type" json:"type"`
	Source        User   `json:"-"`
	SourceId      int    `json:"source"`
	DestinationId int    `mapstructure:"destination" json:"destination"`
	Text          string `mapstructure:"text" json:"text"`
	Action        string `mapstructure:"action" json:"action"`
	CallId        int    `json:"call_id,omitempty"`
	CreatedAt     string `mapstructure:"created_at" json:"created_at"`
	ReadAt        string `json:"read_at"`
	Incognito     bool
}
