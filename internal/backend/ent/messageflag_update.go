// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ProtonMail/gluon/internal/backend/ent/messageflag"
	"github.com/ProtonMail/gluon/internal/backend/ent/predicate"
)

// MessageFlagUpdate is the builder for updating MessageFlag entities.
type MessageFlagUpdate struct {
	config
	hooks    []Hook
	mutation *MessageFlagMutation
}

// Where appends a list predicates to the MessageFlagUpdate builder.
func (mfu *MessageFlagUpdate) Where(ps ...predicate.MessageFlag) *MessageFlagUpdate {
	mfu.mutation.Where(ps...)
	return mfu
}

// SetValue sets the "Value" field.
func (mfu *MessageFlagUpdate) SetValue(s string) *MessageFlagUpdate {
	mfu.mutation.SetValue(s)
	return mfu
}

// Mutation returns the MessageFlagMutation object of the builder.
func (mfu *MessageFlagUpdate) Mutation() *MessageFlagMutation {
	return mfu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mfu *MessageFlagUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mfu.hooks) == 0 {
		affected, err = mfu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageFlagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mfu.mutation = mutation
			affected, err = mfu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mfu.hooks) - 1; i >= 0; i-- {
			if mfu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mfu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mfu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mfu *MessageFlagUpdate) SaveX(ctx context.Context) int {
	affected, err := mfu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mfu *MessageFlagUpdate) Exec(ctx context.Context) error {
	_, err := mfu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mfu *MessageFlagUpdate) ExecX(ctx context.Context) {
	if err := mfu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mfu *MessageFlagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   messageflag.Table,
			Columns: messageflag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: messageflag.FieldID,
			},
		},
	}
	if ps := mfu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mfu.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: messageflag.FieldValue,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mfu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{messageflag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// MessageFlagUpdateOne is the builder for updating a single MessageFlag entity.
type MessageFlagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MessageFlagMutation
}

// SetValue sets the "Value" field.
func (mfuo *MessageFlagUpdateOne) SetValue(s string) *MessageFlagUpdateOne {
	mfuo.mutation.SetValue(s)
	return mfuo
}

// Mutation returns the MessageFlagMutation object of the builder.
func (mfuo *MessageFlagUpdateOne) Mutation() *MessageFlagMutation {
	return mfuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mfuo *MessageFlagUpdateOne) Select(field string, fields ...string) *MessageFlagUpdateOne {
	mfuo.fields = append([]string{field}, fields...)
	return mfuo
}

// Save executes the query and returns the updated MessageFlag entity.
func (mfuo *MessageFlagUpdateOne) Save(ctx context.Context) (*MessageFlag, error) {
	var (
		err  error
		node *MessageFlag
	)
	if len(mfuo.hooks) == 0 {
		node, err = mfuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageFlagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mfuo.mutation = mutation
			node, err = mfuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mfuo.hooks) - 1; i >= 0; i-- {
			if mfuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mfuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, mfuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*MessageFlag)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MessageFlagMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (mfuo *MessageFlagUpdateOne) SaveX(ctx context.Context) *MessageFlag {
	node, err := mfuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mfuo *MessageFlagUpdateOne) Exec(ctx context.Context) error {
	_, err := mfuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mfuo *MessageFlagUpdateOne) ExecX(ctx context.Context) {
	if err := mfuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mfuo *MessageFlagUpdateOne) sqlSave(ctx context.Context) (_node *MessageFlag, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   messageflag.Table,
			Columns: messageflag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: messageflag.FieldID,
			},
		},
	}
	id, ok := mfuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MessageFlag.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mfuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, messageflag.FieldID)
		for _, f := range fields {
			if !messageflag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != messageflag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mfuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mfuo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: messageflag.FieldValue,
		})
	}
	_node = &MessageFlag{config: mfuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mfuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{messageflag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
