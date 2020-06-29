package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/JazzXP/proof-viewer-server/graph/generated"
	"github.com/JazzXP/proof-viewer-server/graph/model"
)

func (r *mutationResolver) AddToShortlist(ctx context.Context, proofID string, image *string) ([]*model.GalleryImage, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveFromShortlist(ctx context.Context, proofID string, image *string) ([]*model.GalleryImage, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddToFavourites(ctx context.Context, proofID string, image *string) ([]*model.GalleryImage, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveFromFavourites(ctx context.Context, proofID string, image *string) ([]*model.GalleryImage, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddGalleryAccess(ctx context.Context, proofID string, email *string) ([]*model.ProofGallery, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveGalleryAccess(ctx context.Context, proofID string, email *string) ([]*model.ProofGallery, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Proofs(ctx context.Context) ([]*model.ProofGallery, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
