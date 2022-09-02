package backend

import (
	"context"

	"github.com/ProtonMail/gluon/internal/db"
	"github.com/ProtonMail/gluon/internal/db/ent"
	"github.com/ProtonMail/gluon/internal/state"
	"github.com/ProtonMail/gluon/store"
)

// StateUserInterfaceImpl should be used to interface with the user type from a State type. This is meant to control
// the API boundary layer.
type StateUserInterfaceImpl struct {
	u *user
	c state.Connector
}

func newStateUserInterfaceImpl(u *user, connector state.Connector) *StateUserInterfaceImpl {
	return &StateUserInterfaceImpl{u: u, c: connector}
}

func (s *StateUserInterfaceImpl) GetUserID() string {
	return s.u.userID
}

func (s *StateUserInterfaceImpl) GetDelimiter() string {
	return s.u.delimiter
}

func (s *StateUserInterfaceImpl) GetDB() *db.DB {
	return s.u.db
}

func (s *StateUserInterfaceImpl) GetRemote() state.Connector {
	return s.c
}

func (s *StateUserInterfaceImpl) GetStore() store.Store {
	return s.u.store
}

func (s *StateUserInterfaceImpl) QueueOrApplyStateUpdate(ctx context.Context, tx *ent.Tx, update state.Update) error {
	// If we detect a state id in the context, it means this function call is a result of a User interaction.
	// When that happens the update needs to be applied to the state matching the state ID immediately. If no such
	// stateID exists or the context information is not present, all updates are queued for later execution.
	stateID, ok := state.GetStateIDFromContext(ctx)
	if !ok {
		return s.u.forState(func(state *state.State) error {
			state.QueueUpdates(update)
			return nil
		})
	} else {
		return s.u.forState(func(state *state.State) error {
			if state.StateID != stateID {
				state.QueueUpdates(update)

				return nil
			} else {
				if !update.Filter(state) {
					return nil
				}

				return update.Apply(ctx, tx, state)
			}
		})
	}
}

func (s *StateUserInterfaceImpl) ReleaseState(ctx context.Context, st *state.State) error {
	return s.u.removeState(ctx, st)
}