package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"fmt"

	"github.com/RouteHub-Link/routehub-service-graphql/graph/model"
	"github.com/google/uuid"
)

// CreateTicket is the resolver for the createTicket field.
func (r *mutationResolver) CreateTicket(ctx context.Context, input model.TicketInput) (*model.Ticket, error) {
	panic(fmt.Errorf("not implemented: CreateTicket - createTicket"))
}

// UpdateTicket is the resolver for the updateTicket field.
func (r *mutationResolver) UpdateTicket(ctx context.Context, id uuid.UUID, input model.TicketInput) (*model.Ticket, error) {
	panic(fmt.Errorf("not implemented: UpdateTicket - updateTicket"))
}

// DeleteTicket is the resolver for the deleteTicket field.
func (r *mutationResolver) DeleteTicket(ctx context.Context, id uuid.UUID) (*model.Ticket, error) {
	panic(fmt.Errorf("not implemented: DeleteTicket - deleteTicket"))
}

// CreateTicketMessage is the resolver for the createTicketMessage field.
func (r *mutationResolver) CreateTicketMessage(ctx context.Context, ticketID uuid.UUID, input model.TicketMessageInput) (*model.TicketMessage, error) {
	panic(fmt.Errorf("not implemented: CreateTicketMessage - createTicketMessage"))
}

// UpdateTicketMessage is the resolver for the updateTicketMessage field.
func (r *mutationResolver) UpdateTicketMessage(ctx context.Context, id uuid.UUID, input model.TicketMessageInput) (*model.TicketMessage, error) {
	panic(fmt.Errorf("not implemented: UpdateTicketMessage - updateTicketMessage"))
}

// DeleteTicketMessage is the resolver for the deleteTicketMessage field.
func (r *mutationResolver) DeleteTicketMessage(ctx context.Context, id uuid.UUID) (*model.TicketMessage, error) {
	panic(fmt.Errorf("not implemented: DeleteTicketMessage - deleteTicketMessage"))
}

// Ticket is the resolver for the ticket field.
func (r *queryResolver) Ticket(ctx context.Context, id uuid.UUID) (*model.Ticket, error) {
	panic(fmt.Errorf("not implemented: Ticket - ticket"))
}

// Tickets is the resolver for the tickets field.
func (r *queryResolver) Tickets(ctx context.Context) ([]*model.Ticket, error) {
	panic(fmt.Errorf("not implemented: Tickets - tickets"))
}
