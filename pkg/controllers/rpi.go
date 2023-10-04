package controllers

import (
	"context"
	"github.com/jedi-knights/ecnl/pkg/dal"
	"github.com/jedi-knights/ecnl/pkg/models"
	"github.com/jedi-knights/rpi/pkg/match"
	"github.com/jedi-knights/rpi/pkg/schedule"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"slices"
	"sort"
	"time"
)

type RIPer interface {
	GetRanking()
}

type RPI struct{}

func NewRPI() *RPI {
	return &RPI{}
}

func (r *RPI) GenerateRankings(ageGroup string) ([]models.RPIRankingData, error) {
	var (
		err         error
		client      *mongo.Client
		ctx         context.Context
		rpi         float64
		teamNames   []string
		matches     []models.MatchEvent
		rpiSchedule *schedule.Schedule
		data        []models.RPIRankingData
	)

	log.Printf("processing age group %s\n", ageGroup)

	ctx, _ = context.WithTimeout(context.Background(), 1*time.Minute)

	// get the client
	client = dal.MustGetClient(ctx)

	// get the database
	database := client.Database("ecnl")

	// create collection
	matchesCollection := database.Collection("matches")

	// create data access objects
	matchDAO := dal.NewMatchEventDAO(ctx, matchesCollection)

	// This should return with the latest matches for the ECNL
	if matches, err = matchDAO.GetECNLByAgeGroup(ageGroup); err != nil {
		return nil, err
	}

	// convert the matches to the scheule for RPI computation
	rpiSchedule = schedule.NewSchedule()

	for _, m := range matches {
		rpiMatch := match.NewMatch()

		if rpiMatch.Date, err = time.Parse(m.GameDate, "2023-09-09T12:00:00"); err != nil {
			rpiMatch.Date = time.Now() // date doesn't matter to the computation
		}

		if len(m.HomeTeamName) == 0 || len(m.AwayTeamName) == 0 {
			continue
		}

		if !slices.Contains(teamNames, m.HomeTeamName) {
			teamNames = append(teamNames, m.HomeTeamName)
		}

		if !slices.Contains(teamNames, m.AwayTeamName) {
			teamNames = append(teamNames, m.AwayTeamName)
		}

		rpiMatch.Home.Name = m.HomeTeamName
		rpiMatch.Away.Name = m.AwayTeamName
		rpiMatch.Home.Score = m.HomeTeamScore
		rpiMatch.Away.Score = m.AwayTeamScore

		rpiSchedule.AddMatch(rpiMatch)
	}

	// sort the team names
	sort.Slice(teamNames, func(i, j int) bool {
		return teamNames[i] < teamNames[j]
	})

	// at this point the match data is loaded and RPI values can be computed

	// Get the list of teams then calculate the RPI for each team.
	for _, teamName := range teamNames {
		if teamName == "" {
			continue
		}

		if rpi, err = rpiSchedule.CalculateRPI(teamName); err != nil {
			panic(err)
		}

		data = append(data, models.RPIRankingData{
			TeamName: teamName,
			RPI:      rpi,
			Ranking:  -1, // will update this later after sorting
		})
	}

	// Sort the data by RPI
	sort.Slice(data, func(i, j int) bool {
		return data[i].RPI > data[j].RPI
	})

	// Update the ranking
	for i := range data {
		data[i].Ranking = i + 1
	}

	return data, nil
}
