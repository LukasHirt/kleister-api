package teams

import (
	"context"
	"time"

	"github.com/kleister/kleister-api/pkg/model"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// LoggingRequestID returns the request ID as string for logging
type LoggingRequestID func(context.Context) string

type loggingService struct {
	service   Service
	requestID LoggingRequestID
	logger    zerolog.Logger
}

// NewLoggingService wraps the Service and provides logging for its methods.
func NewLoggingService(s Service, requestID LoggingRequestID) Service {
	return &loggingService{
		service:   s,
		requestID: requestID,
		logger:    log.With().Str("service", "teams").Logger(),
	}
}

func (s *loggingService) List(ctx context.Context) ([]*model.Team, error) {
	start := time.Now()
	records, err := s.service.List(ctx)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "list").
		Dur("duration", time.Since(start)).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to find all teams")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

func (s *loggingService) Show(ctx context.Context, name string) (*model.Team, error) {
	start := time.Now()
	record, err := s.service.Show(ctx, name)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "show").
		Dur("duration", time.Since(start)).
		Str("name", name).
		Logger()

	if err != nil && err != ErrNotFound {
		logger.Warn().
			Err(err).
			Msg("failed to find team by name")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

func (s *loggingService) Create(ctx context.Context, team *model.Team) (*model.Team, error) {
	start := time.Now()
	record, err := s.service.Create(ctx, team)

	name := ""

	if record != nil {
		name = record.Name
	}

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "create").
		Dur("duration", time.Since(start)).
		Str("name", name).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to create team")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

func (s *loggingService) Update(ctx context.Context, team *model.Team) (*model.Team, error) {
	start := time.Now()
	record, err := s.service.Update(ctx, team)

	id := ""
	name := ""

	if record != nil {
		id = record.ID
		name = record.Name
	}

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "update").
		Dur("duration", time.Since(start)).
		Str("id", id).
		Str("name", name).
		Logger()

	if err != nil && err != ErrNotFound {
		logger.Warn().
			Err(err).
			Msg("failed to update team")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

func (s *loggingService) Delete(ctx context.Context, name string) error {
	start := time.Now()
	err := s.service.Delete(ctx, name)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "delete").
		Dur("duration", time.Since(start)).
		Str("name", name).
		Logger()

	if err != nil && err != ErrNotFound {
		logger.Warn().
			Err(err).
			Msg("failed to delete team")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

func (s *loggingService) ListUsers(ctx context.Context, name string) ([]*model.TeamUser, error) {
	start := time.Now()
	records, err := s.service.ListUsers(ctx, name)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "listUsers").
		Dur("duration", time.Since(start)).
		Str("name", name).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to find all team users")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

func (s *loggingService) AppendUser(ctx context.Context, teamID, userID, perm string) error {
	start := time.Now()
	err := s.service.AppendUser(ctx, teamID, userID, perm)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "appendUser").
		Dur("duration", time.Since(start)).
		Str("team", teamID).
		Str("user", userID).
		Str("perm", perm).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to append team to user")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

func (s *loggingService) PermitUser(ctx context.Context, teamID, userID, perm string) error {
	start := time.Now()
	err := s.service.PermitUser(ctx, teamID, userID, perm)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "permitUser").
		Dur("duration", time.Since(start)).
		Str("team", teamID).
		Str("user", userID).
		Str("perm", perm).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to update user perms")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

func (s *loggingService) DropUser(ctx context.Context, teamID, userID string) error {
	start := time.Now()
	err := s.service.DropUser(ctx, teamID, userID)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "dropUser").
		Dur("duration", time.Since(start)).
		Str("team", teamID).
		Str("user", userID).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to drop team from user")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

func (s *loggingService) ListMods(ctx context.Context, name string) ([]*model.TeamMod, error) {
	start := time.Now()
	records, err := s.service.ListMods(ctx, name)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "listMods").
		Dur("duration", time.Since(start)).
		Str("name", name).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to find all team mods")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

func (s *loggingService) AppendMod(ctx context.Context, teamID, modID, perm string) error {
	start := time.Now()
	err := s.service.AppendMod(ctx, teamID, modID, perm)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "appendMod").
		Dur("duration", time.Since(start)).
		Str("team", teamID).
		Str("mod", modID).
		Str("perm", perm).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to append team to mod")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

func (s *loggingService) PermitMod(ctx context.Context, teamID, modID, perm string) error {
	start := time.Now()
	err := s.service.PermitMod(ctx, teamID, modID, perm)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "permitMod").
		Dur("duration", time.Since(start)).
		Str("team", teamID).
		Str("mod", modID).
		Str("perm", perm).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to update mod perms")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

func (s *loggingService) DropMod(ctx context.Context, teamID, modID string) error {
	start := time.Now()
	err := s.service.DropMod(ctx, teamID, modID)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "dropMod").
		Dur("duration", time.Since(start)).
		Str("team", teamID).
		Str("mod", modID).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to drop team from mod")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

func (s *loggingService) ListPacks(ctx context.Context, name string) ([]*model.TeamPack, error) {
	start := time.Now()
	records, err := s.service.ListPacks(ctx, name)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "listPacks").
		Dur("duration", time.Since(start)).
		Str("name", name).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to find all team packs")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

func (s *loggingService) AppendPack(ctx context.Context, teamID, packID, perm string) error {
	start := time.Now()
	err := s.service.AppendPack(ctx, teamID, packID, perm)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "appendPack").
		Dur("duration", time.Since(start)).
		Str("team", teamID).
		Str("pack", packID).
		Str("perm", perm).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to append team to pack")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

func (s *loggingService) PermitPack(ctx context.Context, teamID, packID, perm string) error {
	start := time.Now()
	err := s.service.PermitPack(ctx, teamID, packID, perm)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "permitPack").
		Dur("duration", time.Since(start)).
		Str("team", teamID).
		Str("pack", packID).
		Str("perm", perm).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to update pack perms")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

func (s *loggingService) DropPack(ctx context.Context, teamID, packID string) error {
	start := time.Now()
	err := s.service.DropPack(ctx, teamID, packID)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "dropPack").
		Dur("duration", time.Since(start)).
		Str("team", teamID).
		Str("pack", packID).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to drop team from pack")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}
