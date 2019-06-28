package service

import (
	"context"
	"errors"

	"github.com/Danr17/microservices_project/section_2/frontend/pb"
	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
)

//SiteService describe the Stats service
type SiteService interface {
	GetTable(ctx context.Context, league string) ([]*Table, error)
	GetTeamBestPlayers(ctx context.Context, teamName string) ([]*Player, error)
	GetBestDefenders(ctx context.Context, position string) ([]*Player, error)
	GetBestAttackers(ctx context.Context, position string) ([]*Player, error)
	GetGreatPassers(ctx context.Context, position string) ([]*Player, error)
}

// NewSiteService returns a basic StatsService with all of the expected middlewares wired in.
func NewSiteService(logger log.Logger, conn *grpc.ClientConn) SiteService {
	var svc SiteService
	svc = NewBasicService(conn)
	svc = LoggingMiddleware(logger)(svc)
	return svc
}

// NewBasicService returns a naive, stateless implementation of StatsService.
func NewBasicService(conn *grpc.ClientConn) SiteService {
	return &basicService{
		gcStats: pb.NewStatsServiceClient(conn),
	}
}

type basicService struct {
	gcStats pb.StatsServiceClient
}

var (
	//ErrTeamNotFound unable to find the requested team
	ErrTeamNotFound = errors.New("team not found")
	//ErrPLayerNotFound unable to find requested player
	ErrPLayerNotFound = errors.New("player not found")
	//ErrDisplayTable unable to disply table
	ErrDisplayTable = errors.New("unable to display table")
	//ErrDisplayPlayers unable to disply table
	ErrDisplayPlayers = errors.New("unable to display players")
)

//GetTable display final league table
func (s *basicService) GetTable(ctx context.Context, league string) ([]*Table, error) {
	resp, err := s.gcStats.ListTable(context.Background(), &pb.TableRequest{
		TableName: league,
	})
	if err != nil {
		return nil, ErrDisplayTable  
	}

	teams := make([]*Table, len(resp.Teams))
	for i := range resp.Teams {
		teams[i] = &Table{
			TeamName:    resp.Teams[i].TeamName,
			TeamPlayed:  resp.Teams[i].TeamPlayed,
			TeamWon:     resp.Teams[i].TeamWon,
			TeamDrawn:   resp.Teams[i].TeamDrawn,
			TeamLost:    resp.Teams[i].TeamLost,
			TeamGF:      resp.Teams[i].TeamGF,
			TeamGA:      resp.Teams[i].TeamGA,
			TeamGD:      resp.Teams[i].TeamGD,
			TeamPoints:  resp.Teams[i].TeamPoints,
			TeamCapital: resp.Teams[i].TeamCapital,
		}
	}

	return teams, str2err(resp.Err)
}

//GetTeamBestPLayers diplay top 3 players of a team (one forward, one mid and one defender)
func (s *basicService) GetTeamBestPlayers(ctx context.Context, teamName string) ([]*Player, error) {

	resp, err := s.gcStats.ListTeamPlayers(context.Background(), &pb.TeamRequest{
		TeamName: teamName,
	})
	if err != nil {
		return nil, ErrDisplayPlayers 
	}

	players := make([]*Player, len(resp.Players))
	for i := range resp.Players {
		players[i] = &Player{
			Name:          resp.Players[i].Name,
			Team:          resp.Players[i].Team,
			Nationality:   resp.Players[i].Nationality,
			Position:      resp.Players[i].Position,
			Appearences:   resp.Players[i].Appearences,
			Goals:         resp.Players[i].Goals,
			Assists:       resp.Players[i].Assists,
			Passes:        resp.Players[i].Passes,
			Interceptions: resp.Players[i].Interceptions,
			Tackles:       resp.Players[i].Tackles,
			Fouls:         resp.Players[i].Fouls,
		}
	}

	return players, str2err(resp.Err)
}

//GetBestDefenders display top 3 league defenders
func (s *basicService) GetBestDefenders(ctx context.Context, position string) ([]*Player, error) {

	resp, err := s.gcStats.ListPositionPlayers(context.Background(), &pb.PositionRequest{
		Position: position,
	})
	if err != nil {
		return nil, ErrDisplayPlayers 
	}

	players := make([]*Player, len(resp.Players))
	for i := range resp.Players {
		players[i] = &Player{
			Name:          resp.Players[i].Name,
			Team:          resp.Players[i].Team,
			Nationality:   resp.Players[i].Nationality,
			Position:      resp.Players[i].Position,
			Appearences:   resp.Players[i].Appearences,
			Goals:         resp.Players[i].Goals,
			Assists:       resp.Players[i].Assists,
			Passes:        resp.Players[i].Passes,
			Interceptions: resp.Players[i].Interceptions,
			Tackles:       resp.Players[i].Tackles,
			Fouls:         resp.Players[i].Fouls,
		}
	}

	return players, str2err(resp.Err)
}

//GetBestAttackers display top 3 league attackers
func (s *basicService) GetBestAttackers(ctx context.Context, position string) ([]*Player, error) {

	resp, err := s.gcStats.ListPositionPlayers(context.Background(), &pb.PositionRequest{
		Position: position,
	})
	if err != nil {
		return nil, ErrDisplayPlayers 
	}

	players := make([]*Player, len(resp.Players))
	for i := range resp.Players {
		players[i] = &Player{
			Name:          resp.Players[i].Name,
			Team:          resp.Players[i].Team,
			Nationality:   resp.Players[i].Nationality,
			Position:      resp.Players[i].Position,
			Appearences:   resp.Players[i].Appearences,
			Goals:         resp.Players[i].Goals,
			Assists:       resp.Players[i].Assists,
			Passes:        resp.Players[i].Passes,
			Interceptions: resp.Players[i].Interceptions,
			Tackles:       resp.Players[i].Tackles,
			Fouls:         resp.Players[i].Fouls,
		}
	}

	return players, str2err(resp.Err)
}

//GetGreatPassers display top 3 league passers
func (s *basicService) GetGreatPassers(ctx context.Context, position string) ([]*Player, error) {

	resp, err := s.gcStats.ListPositionPlayers(context.Background(), &pb.PositionRequest{
		Position: position,
	})
	if err != nil {
		return nil, ErrDisplayPlayers 
	}

	players := make([]*Player, len(resp.Players))
	for i := range resp.Players {
		players[i] = &Player{
			Name:          resp.Players[i].Name,
			Team:          resp.Players[i].Team,
			Nationality:   resp.Players[i].Nationality,
			Position:      resp.Players[i].Position,
			Appearences:   resp.Players[i].Appearences,
			Goals:         resp.Players[i].Goals,
			Assists:       resp.Players[i].Assists,
			Passes:        resp.Players[i].Passes,
			Interceptions: resp.Players[i].Interceptions,
			Tackles:       resp.Players[i].Tackles,
			Fouls:         resp.Players[i].Fouls,
		}
	}

	return players, str2err(resp.Err)
}

// Helper function is required to translate Go error types from strings,
// which is the type we use in our IDLs to represent errors.

func str2err(s string) error {
	if s == "" {
		return nil
	}
	return errors.New(s)
}