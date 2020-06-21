package models

import (
	"time"

	"github.com/lissteron/loghole/dashboard/internal/app/domain"
)

type Entry struct {
	Time        time.Time `db:"time"`
	NSec        int64     `db:"nsec"`
	Namespace   string    `db:"namespace"`
	Source      string    `db:"source"`
	Host        string    `db:"host"`
	TraceID     string    `db:"trace_id"`
	Message     string    `db:"message"`
	Params      string    `db:"params"`
	BuildCommit string    `db:"build_commit"`
	ConfigHash  string    `db:"config_hash"`
}

func (e *Entry) ToDomain() *domain.Entry {
	entry := &domain.Entry{
		Time:        e.Time,
		NSec:        e.NSec,
		Namespace:   e.Namespace,
		Source:      e.Source,
		Host:        e.Host,
		TraceID:     e.TraceID,
		Message:     e.Message,
		BuildCommit: e.BuildCommit,
		ConfigHash:  e.ConfigHash,
	}

	if e.Params != "" {
		entry.Params = []byte(e.Params)
	}

	return entry
}
