// Code generated by ent, DO NOT EDIT.

package internal

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ProtonMail/gluon/internal/db_impl/ent_db/internal/mailboxpermflag"
	"github.com/ProtonMail/gluon/internal/db_impl/ent_db/internal/predicate"
)

// MailboxPermFlagUpdate is the builder for updating MailboxPermFlag entities.
type MailboxPermFlagUpdate struct {
	config
	hooks    []Hook
	mutation *MailboxPermFlagMutation
}

// Where appends a list predicates to the MailboxPermFlagUpdate builder.
func (mpfu *MailboxPermFlagUpdate) Where(ps ...predicate.MailboxPermFlag) *MailboxPermFlagUpdate {
	mpfu.mutation.Where(ps...)
	return mpfu
}

// SetValue sets the "Value" field.
func (mpfu *MailboxPermFlagUpdate) SetValue(s string) *MailboxPermFlagUpdate {
	mpfu.mutation.SetValue(s)
	return mpfu
}

// Mutation returns the MailboxPermFlagMutation object of the builder.
func (mpfu *MailboxPermFlagUpdate) Mutation() *MailboxPermFlagMutation {
	return mpfu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mpfu *MailboxPermFlagUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, MailboxPermFlagMutation](ctx, mpfu.sqlSave, mpfu.mutation, mpfu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mpfu *MailboxPermFlagUpdate) SaveX(ctx context.Context) int {
	affected, err := mpfu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mpfu *MailboxPermFlagUpdate) Exec(ctx context.Context) error {
	_, err := mpfu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mpfu *MailboxPermFlagUpdate) ExecX(ctx context.Context) {
	if err := mpfu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mpfu *MailboxPermFlagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(mailboxpermflag.Table, mailboxpermflag.Columns, sqlgraph.NewFieldSpec(mailboxpermflag.FieldID, field.TypeInt))
	if ps := mpfu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mpfu.mutation.Value(); ok {
		_spec.SetField(mailboxpermflag.FieldValue, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mpfu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mailboxpermflag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mpfu.mutation.done = true
	return n, nil
}

// MailboxPermFlagUpdateOne is the builder for updating a single MailboxPermFlag entity.
type MailboxPermFlagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MailboxPermFlagMutation
}

// SetValue sets the "Value" field.
func (mpfuo *MailboxPermFlagUpdateOne) SetValue(s string) *MailboxPermFlagUpdateOne {
	mpfuo.mutation.SetValue(s)
	return mpfuo
}

// Mutation returns the MailboxPermFlagMutation object of the builder.
func (mpfuo *MailboxPermFlagUpdateOne) Mutation() *MailboxPermFlagMutation {
	return mpfuo.mutation
}

// Where appends a list predicates to the MailboxPermFlagUpdate builder.
func (mpfuo *MailboxPermFlagUpdateOne) Where(ps ...predicate.MailboxPermFlag) *MailboxPermFlagUpdateOne {
	mpfuo.mutation.Where(ps...)
	return mpfuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mpfuo *MailboxPermFlagUpdateOne) Select(field string, fields ...string) *MailboxPermFlagUpdateOne {
	mpfuo.fields = append([]string{field}, fields...)
	return mpfuo
}

// Save executes the query and returns the updated MailboxPermFlag entity.
func (mpfuo *MailboxPermFlagUpdateOne) Save(ctx context.Context) (*MailboxPermFlag, error) {
	return withHooks[*MailboxPermFlag, MailboxPermFlagMutation](ctx, mpfuo.sqlSave, mpfuo.mutation, mpfuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mpfuo *MailboxPermFlagUpdateOne) SaveX(ctx context.Context) *MailboxPermFlag {
	node, err := mpfuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mpfuo *MailboxPermFlagUpdateOne) Exec(ctx context.Context) error {
	_, err := mpfuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mpfuo *MailboxPermFlagUpdateOne) ExecX(ctx context.Context) {
	if err := mpfuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mpfuo *MailboxPermFlagUpdateOne) sqlSave(ctx context.Context) (_node *MailboxPermFlag, err error) {
	_spec := sqlgraph.NewUpdateSpec(mailboxpermflag.Table, mailboxpermflag.Columns, sqlgraph.NewFieldSpec(mailboxpermflag.FieldID, field.TypeInt))
	id, ok := mpfuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`internal: missing "MailboxPermFlag.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mpfuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mailboxpermflag.FieldID)
		for _, f := range fields {
			if !mailboxpermflag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("internal: invalid field %q for query", f)}
			}
			if f != mailboxpermflag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mpfuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mpfuo.mutation.Value(); ok {
		_spec.SetField(mailboxpermflag.FieldValue, field.TypeString, value)
	}
	_node = &MailboxPermFlag{config: mpfuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mpfuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mailboxpermflag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mpfuo.mutation.done = true
	return _node, nil
}