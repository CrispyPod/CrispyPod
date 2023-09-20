package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.37

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"crispypod.com/crispypod/db"
	"crispypod.com/crispypod/graph/model"
	"crispypod.com/crispypod/helpers"
	"crispypod.com/crispypod/models"
	"crispypod.com/crispypod/rssfeed"
	"github.com/google/uuid"
)

// CreateEpisode is the resolver for the createEpisode field.
func (r *mutationResolver) CreateEpisode(ctx context.Context, input *model.NewEpisode) (*model.Episode, error) {
	userName := helpers.JWTFromContext(ctx)
	if len(userName) == 0 {
		return nil, errors.New("authorization failed")
	}

	var jwtDbUser models.DbUser
	if err := db.DB.Model(models.DbUser{UserName: userName}).Find(&jwtDbUser).Error; err != nil {
		return nil, errors.New("user not found")
	}

	newEpisode := models.Episode{
		ID:            uuid.New(),
		Title:         input.Title,
		Description:   input.Description,
		EpisodeStatus: models.EpisodeStatus_Draft,
		UserID:        jwtDbUser.ID,
		CreateTime:    time.Now(),
	}

	if input.AudioFileName != nil {
		newEpisode.AudioFileName.String = *input.AudioFileName
	}

	if input.AudioFileUploadName != nil {
		newEpisode.AudioFileUploadName.String = *input.AudioFileUploadName
	}

	if input.AudioFileDuration != nil {
		newEpisode.AudioFileDuration.Int64 = int64(*input.AudioFileDuration)
	}

	if input.ThumbnailFileName != nil {
		newEpisode.ThumbnailFileName.String = *input.ThumbnailFileName
	}

	if input.ThumbnailFileUploadName != nil {
		newEpisode.ThumbnailUploadName.String = *input.ThumbnailFileUploadName
	}

	db.DB.Create(newEpisode)

	return newEpisode.ToGQLEpisode(), nil
}

// ModifyEpisode is the resolver for the modifyEpisode field.
func (r *mutationResolver) ModifyEpisode(ctx context.Context, id string, input *model.NewEpisode) (*model.Episode, error) {
	userName := helpers.JWTFromContext(ctx)
	if len(userName) == 0 {
		return nil, errors.New("authorization failed")
	}

	var jwtDbUser models.DbUser
	if err := db.DB.Model(models.DbUser{UserName: userName}).Find(&jwtDbUser).Error; err != nil {
		return nil, errors.New("user not found")
	}

	var dbEpisode models.Episode
	if err := db.DB.Model(models.Episode{ID: uuid.Must(uuid.Parse(id))}).Find(&dbEpisode).Error; err != nil {
		return nil, errors.New("episode not found")
	}

	dbEpisode.Title = input.Title
	dbEpisode.Description = input.Description
	dbEpisode.EpisodeStatus = models.EpisodeStatusType(*input.EpisodeStatus)

	if input.AudioFileName != nil {
		dbEpisode.AudioFileName = sql.NullString{String: *input.AudioFileName, Valid: true}
	}

	if input.AudioFileUploadName != nil {
		dbEpisode.AudioFileUploadName = sql.NullString{String: *input.AudioFileUploadName, Valid: true}
	}

	if input.AudioFileDuration != nil {
		dbEpisode.AudioFileDuration = sql.NullInt64{Int64: int64(*input.AudioFileDuration), Valid: true}
	}

	if input.ThumbnailFileName != nil {
		dbEpisode.ThumbnailFileName = sql.NullString{String: *input.ThumbnailFileName, Valid: true}
	}

	if input.ThumbnailFileUploadName != nil {
		dbEpisode.ThumbnailUploadName = sql.NullString{String: *input.ThumbnailFileUploadName, Valid: true}
	}

	db.DB.Save(dbEpisode)
	rssfeed.GenerateRSSFeed()

	return dbEpisode.ToGQLEpisode(), nil
}

// Episodes is the resolver for the episodes field.
func (r *queryResolver) Episodes(ctx context.Context, pagination model.Pagination) (*model.EpisodesResult, error) {
	var episodes []models.Episode
	var rtEpisodes []*model.Episode
	var count int64
	if userName := helpers.JWTFromContext(ctx); len(userName) == 0 {
		err := db.DB.
			Scopes(helpers.Paginate(pagination.PageIndex, pagination.PerPage)).
			Find(&episodes, models.Episode{EpisodeStatus: models.EpisodeStatus_Published}).
			Count(&count).Error
		if err != nil {
			return nil, errors.New("episodes not found")
		}
	} else {
		if err := db.DB.Scopes(helpers.Paginate(pagination.PageIndex, pagination.PerPage)).
			Find(&episodes).
			Count(&count).Error; err != nil {
			return nil, errors.New("episodes not found")
		}
	}
	// Convert
	for _, e := range episodes {
		rtEpisodes = append(rtEpisodes, e.ToGQLEpisode())
	}

	pageInfo := helpers.GetPageInfo(pagination.PageIndex, pagination.PerPage, int(count))
	// return rtEpisodes, nil
	return &model.EpisodesResult{
		TotalCount: int(count),
		Items:      rtEpisodes,
		PageInfo:   &pageInfo,
	}, nil
}

// Episode is the resolver for the episode field.
func (r *queryResolver) Episode(ctx context.Context, id string) (*model.Episode, error) {
	userName := helpers.JWTFromContext(ctx)
	var dbEpisode models.Episode
	if err := db.DB.Find(&dbEpisode, models.Episode{ID: uuid.Must(uuid.Parse(id))}).Error; err != nil {
		return nil, errors.New("episode not found")
	}
	if len(userName) == 0 && dbEpisode.EpisodeStatus == models.EpisodeStatus_Draft {
		// check if episode status was draft, if so, we will not return it.
		return nil, errors.New("episode not accessable")
	}
	return dbEpisode.ToGQLEpisode(), nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, pagination model.Pagination) (*model.UsersResult, error) {
	var count int64
	userName := helpers.JWTFromContext(ctx)
	if len(userName) == 0 {
		return nil, errors.New("authorization failed")
	}

	var dbJWTUser models.DbUser
	if err := db.DB.Find(&dbJWTUser, models.DbUser{UserName: userName}).Count(&count).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	if !dbJWTUser.IsAdmin {
		return nil, errors.New("user not to access this data")
	}

	var users []models.DbUser
	if err := db.DB.Model(&models.DbUser{}).Count(&count).Find(&users).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	var rtUsers []*model.User
	for _, u := range users {
		rtUsers = append(rtUsers, u.ToGQLUser())
	}
	// return rtUsers, nil
	// return nil, nil
	pageInfo := helpers.GetPageInfo(pagination.PageIndex, pagination.PerPage, int(count))
	return &model.UsersResult{
		TotalCount: int(count),
		Items:      rtUsers,
		PageInfo:   &pageInfo,
	}, nil
}

// Login is the resolver for the login field.
func (r *queryResolver) Login(ctx context.Context, credential model.Credential) (*model.LoginData, error) {
	var user models.DbUser
	if err := db.DB.Model(models.DbUser{UserName: credential.UserName}).First(&user).Error; err != nil {
		return nil, errors.New("user with provided credentials not found")
	}
	if helpers.CheckPasswordHash(credential.Password, user.Password) {
		token, _ := helpers.GenerateJwt(user.UserName)
		return &model.LoginData{
			Token: token,
		}, nil
	}
	return nil, errors.New("user with provided credentials not found")
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	userName := helpers.JWTFromContext(ctx)
	if len(userName) == 0 {
		return nil, errors.New("authorization failed")
	}

	var dbJWTUser models.DbUser
	if err := db.DB.Model(&models.DbUser{UserName: userName}).Find(&dbJWTUser).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	return dbJWTUser.ToGQLUser(), nil
}

// SiteConfig is the resolver for the siteConfig field.
func (r *queryResolver) SiteConfig(ctx context.Context) (*model.SiteConfig, error) {
	var siteConfig models.SiteConfig
	if err := db.DB.First(&siteConfig).Error; err != nil {
		siteConfig = models.SiteConfig{
			ID:              uuid.New(),
			SiteName:        "CrispyPod",
			SiteDescription: "Super awesome podcast!",
			SiteUrl:         "https://crispypod.com",
		}
		db.DB.Create(&siteConfig)
	}

	return siteConfig.ToGQLSiteConfig(), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
