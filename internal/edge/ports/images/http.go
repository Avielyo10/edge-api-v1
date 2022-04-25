package ports

import (
	"encoding/json"
	"errors"
	"net/http"

	httperr "github.com/Avielyo10/edge-api/internal/common/server/httperr"
	"github.com/Avielyo10/edge-api/internal/edge/app"
	"github.com/Avielyo10/edge-api/internal/edge/app/command"
	"github.com/Avielyo10/edge-api/internal/edge/domain/image"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// HttpServer is a port for the http server.
type HttpServer struct {
	app app.Application
}

// NewHttpServer returns a new HttpServer.
func NewHttpServer(application app.Application) HttpServer {
	return HttpServer{
		app: application,
	}
}

// CreateImage creates a new image. Implementing ports.ServerInterface
func (h HttpServer) CreateImage(w http.ResponseWriter, r *http.Request) {
	var req CreateImageRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		httperr.HandleImageErrors(w, r, err)
		return
	}
	if err := CheckCreateRequest(req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.Respond(w, r, httperr.NewBadRequest(err.Error()))
		return
	}
	ctx := r.Context()
	cmd := command.CreateImage{
		UUID:         uuid.NewString(),
		Name:         string(*req.Name),
		Description:  string(*req.Description),
		Distribution: string(*req.Distribution),
		Username:     string(*req.Username),
		SSHKey:       string(*req.SshKey),
		OutputType:   *req.OutputType,
		Tags:         *req.Tags,
		Packages:     *req.Packages,
		Status:       image.Building.String(),
		Version:      1,
		Repos:        reposToInterfaces(req.Repositories),
	}
	image, err := h.app.Commands.CreateImage.Handle(ctx, cmd)
	if err != nil {
		httperr.HandleImageErrors(w, r, err)
		return
	}
	imageRes := imageToResponse(image)
	render.Status(r, http.StatusCreated)
	render.Respond(w, r, imageRes)
}

// GetImage returns the image with the given uuid. Implementing ports.ServerInterface
func (h HttpServer) GetImage(w http.ResponseWriter, r *http.Request, imageId string) {
	ctx := r.Context()
	image, err := h.app.Queries.GetImage.Handle(ctx, imageId)

	if err != nil {
		httperr.HandleImageErrors(w, r, err)
		return
	}
	imageRes := imageToResponse(image)
	render.Status(r, http.StatusOK)
	render.Respond(w, r, imageRes)
}

// UpdateImage updates the image with the given uuid. Implementing ports.ServerInterface
func (h HttpServer) UpdateImage(w http.ResponseWriter, r *http.Request, imageId string) {
	var req UpdateImageRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		httperr.HandleImageErrors(w, r, err)
		return
	}

	ctx := r.Context()
	cmd := command.UpdateImage{
		UUIDToUpdate: imageId,
	}
	if req.Name != nil {
		cmd.Name = string(*req.Name)
	}
	if req.Description != nil {
		cmd.Description = string(*req.Description)
	}
	if req.Tags != nil {
		if req.Tags.Add != nil {
			cmd.TagsToAdd = *req.Tags.Add
		}
		if req.Tags.Remove != nil {
			cmd.TagsToRemove = *req.Tags.Remove
		}
	}
	err = h.app.Commands.UpdateImage.Handle(ctx, cmd)
	if err != nil {
		httperr.HandleImageErrors(w, r, err)
		return
	}
	render.Status(r, http.StatusNoContent)
	render.Respond(w, r, nil)
}

// DeleteImage deletes the image with the given uuid. Implementing ports.ServerInterface
func (h HttpServer) DeleteImage(w http.ResponseWriter, r *http.Request, imageId string) {
	ctx := r.Context()
	err := h.app.Commands.DeleteImage.Handle(ctx, imageId)
	if err != nil {
		httperr.HandleImageErrors(w, r, err)
		return
	}
	render.Status(r, http.StatusNoContent)
	render.Respond(w, r, nil)
}

// GetImages returns all images. Implementing ports.ServerInterface - TODO: work on params
func (h HttpServer) GetImages(w http.ResponseWriter, r *http.Request, params GetImagesParams) {
	ctx := r.Context()
	images, err := h.app.Queries.GetImages.Handle(ctx)
	if err != nil {
		httperr.HandleImageErrors(w, r, err)
		return
	}
	imagesRes := imagesToResponse(images)
	render.Status(r, http.StatusOK)
	res := map[string]interface{}{
		"count": len(imagesRes),
		"items": imagesRes,
	}
	render.Respond(w, r, res)
}

// DeleteImagesImageIdUpdate cancels the image update with the given uuid. Implementing ports.ServerInterface
func (h HttpServer) DeleteImagesImageIdUpdate(w http.ResponseWriter, r *http.Request, imageId string) {
	ctx := r.Context()
	err := h.app.Commands.CancelUpgradeImage.Handle(ctx, imageId)
	if err != nil {
		httperr.HandleImageErrors(w, r, err)
		return
	}
	render.Status(r, http.StatusNoContent)
	render.Respond(w, r, nil)
}

// CreateNewVersion creates a new version of the image with the given uuid (upgrade process). Implementing ports.ServerInterface
func (h HttpServer) CreateNewVersion(w http.ResponseWriter, r *http.Request, imageId string) {
	var req UpgradeImageRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		httperr.HandleImageErrors(w, r, err)
		return
	}

	ctx := r.Context()
	cmd := command.UpgradeImage{
		UUIDToUpgrade:    imageId,
		Name:             string(*req.Name),
		Description:      string(*req.Description),
		TagsToRemove:     *req.Tags.Remove,
		TagsToAdd:        *req.Tags.Add,
		PackagesToRemove: *req.Packages.Remove,
		PackagesToAdd:    *req.Packages.Add,
	}
	err = h.app.Commands.UpgradeImage.Handle(ctx, cmd)
	if err != nil {
		httperr.HandleImageErrors(w, r, err)
		return
	}
	render.Status(r, http.StatusNoContent)
	render.Respond(w, r, nil)
}

// imagesToResponse converts a slice of images to a slice of image responses.
func imagesToResponse(images []*image.Image) []ImageResponse {
	imagesRes := make([]ImageResponse, len(images))
	for i, image := range images {
		imagesRes[i] = imageToResponse(image)
	}
	return imagesRes
}

// imageToResponse converts an image to a response.
func imageToResponse(image *image.Image) ImageResponse {
	description := Description(image.Description())
	name := Name(image.Name().String())
	status := Status(image.Status().String())
	uuid := UUID(image.UUID())
	version := Version(image.Version().Uint())
	outputTypes := make(OutputTypes, len(image.OutputTypes()))
	for i, outputType := range image.OutputTypes() {
		outputTypes[i] = outputType.String()
	}
	resp := ImageResponse{
		Description: &description,
		Name:        &name,
		Status:      &status,
		Uuid:        &uuid,
		Version:     &version,
		OutputType:  &outputTypes,
	}
	createdAt := CreatedAt(image.CreatedAt())
	if !image.CreatedAt().IsZero() {
		logrus.Infof("image created at: %s", image.CreatedAt().String())
		resp.CreatedAt = &createdAt
	} else {
		resp.CreatedAt = nil
	}
	deletedAt := DeletedAt(image.DeletedAt())
	if !image.DeletedAt().IsZero() {
		resp.DeletedAt = &deletedAt
	}
	updatedAt := UpdatedAt(image.UpdatedAt())
	if !image.UpdatedAt().IsZero() {
		resp.UpdatedAt = &updatedAt
	}
	return resp
}

// reposToInterfaces converts repositories to a slice of repository interfaces.
func reposToInterfaces(repos *Repositories) []interface{} {
	// iterate over repositories
	repositories := make([]interface{}, len(*repos))
	for i, repo := range *repos {
		repositories[i] = map[string]interface{}{
			"name": repo.Name,
			"url":  repo.Url,
		}
	}
	return repositories
}

// CheckCreateRequest checks the create a new image request.
func CheckCreateRequest(req CreateImageRequest) error {
	if req.Name == nil {
		return errors.New("name is required")
	}
	if req.Description == nil {
		return errors.New("description is required")
	}
	if req.Distribution == nil {
		return errors.New("distribution is required")
	}
	if req.SshKey == nil {
		return errors.New("ssh key is required")
	}
	if req.Username == nil {
		return errors.New("username is required")
	}
	return nil
}
