// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/examples/traversal/ent/group"
	"entgo.io/ent/examples/traversal/ent/predicate"
	"entgo.io/ent/examples/traversal/ent/user"
	"entgo.io/ent/schema/field"
)

// GroupUpdate is the builder for updating Group entities.
type GroupUpdate struct {
	config
	hooks    []Hook
	mutation *GroupMutation
}

// Where appends a list predicates to the GroupUpdate builder.
func (gu *GroupUpdate) Where(ps ...predicate.Group) *GroupUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetName sets the "name" field.
func (gu *GroupUpdate) SetName(s string) *GroupUpdate {
	gu.mutation.SetName(s)
	return gu
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (gu *GroupUpdate) AddUserIDs(ids ...int) *GroupUpdate {
	gu.mutation.AddUserIDs(ids...)
	return gu
}

// AddUsers adds the "users" edges to the User entity.
func (gu *GroupUpdate) AddUsers(u ...*User) *GroupUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gu.AddUserIDs(ids...)
}

// SetAdminID sets the "admin" edge to the User entity by ID.
func (gu *GroupUpdate) SetAdminID(id int) *GroupUpdate {
	gu.mutation.SetAdminID(id)
	return gu
}

// SetNillableAdminID sets the "admin" edge to the User entity by ID if the given value is not nil.
func (gu *GroupUpdate) SetNillableAdminID(id *int) *GroupUpdate {
	if id != nil {
		gu = gu.SetAdminID(*id)
	}
	return gu
}

// SetAdmin sets the "admin" edge to the User entity.
func (gu *GroupUpdate) SetAdmin(u *User) *GroupUpdate {
	return gu.SetAdminID(u.ID)
}

// Mutation returns the GroupMutation object of the builder.
func (gu *GroupUpdate) Mutation() *GroupMutation {
	return gu.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (gu *GroupUpdate) ClearUsers() *GroupUpdate {
	gu.mutation.ClearUsers()
	return gu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (gu *GroupUpdate) RemoveUserIDs(ids ...int) *GroupUpdate {
	gu.mutation.RemoveUserIDs(ids...)
	return gu
}

// RemoveUsers removes "users" edges to User entities.
func (gu *GroupUpdate) RemoveUsers(u ...*User) *GroupUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gu.RemoveUserIDs(ids...)
}

// ClearAdmin clears the "admin" edge to the User entity.
func (gu *GroupUpdate) ClearAdmin() *GroupUpdate {
	gu.mutation.ClearAdmin()
	return gu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GroupUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, GroupMutation](ctx, gu.sqlSave, gu.mutation, gu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GroupUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GroupUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GroupUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gu *GroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(group.Table, group.Columns, sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt))
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.Name(); ok {
		_spec.SetField(group.FieldName, field.TypeString, value)
	}
	if gu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !gu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.AdminCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   group.AdminTable,
			Columns: []string{group.AdminColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.AdminIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   group.AdminTable,
			Columns: []string{group.AdminColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gu.mutation.done = true
	return n, nil
}

// GroupUpdateOne is the builder for updating a single Group entity.
type GroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GroupMutation
}

// SetName sets the "name" field.
func (guo *GroupUpdateOne) SetName(s string) *GroupUpdateOne {
	guo.mutation.SetName(s)
	return guo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (guo *GroupUpdateOne) AddUserIDs(ids ...int) *GroupUpdateOne {
	guo.mutation.AddUserIDs(ids...)
	return guo
}

// AddUsers adds the "users" edges to the User entity.
func (guo *GroupUpdateOne) AddUsers(u ...*User) *GroupUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return guo.AddUserIDs(ids...)
}

// SetAdminID sets the "admin" edge to the User entity by ID.
func (guo *GroupUpdateOne) SetAdminID(id int) *GroupUpdateOne {
	guo.mutation.SetAdminID(id)
	return guo
}

// SetNillableAdminID sets the "admin" edge to the User entity by ID if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableAdminID(id *int) *GroupUpdateOne {
	if id != nil {
		guo = guo.SetAdminID(*id)
	}
	return guo
}

// SetAdmin sets the "admin" edge to the User entity.
func (guo *GroupUpdateOne) SetAdmin(u *User) *GroupUpdateOne {
	return guo.SetAdminID(u.ID)
}

// Mutation returns the GroupMutation object of the builder.
func (guo *GroupUpdateOne) Mutation() *GroupMutation {
	return guo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (guo *GroupUpdateOne) ClearUsers() *GroupUpdateOne {
	guo.mutation.ClearUsers()
	return guo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (guo *GroupUpdateOne) RemoveUserIDs(ids ...int) *GroupUpdateOne {
	guo.mutation.RemoveUserIDs(ids...)
	return guo
}

// RemoveUsers removes "users" edges to User entities.
func (guo *GroupUpdateOne) RemoveUsers(u ...*User) *GroupUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return guo.RemoveUserIDs(ids...)
}

// ClearAdmin clears the "admin" edge to the User entity.
func (guo *GroupUpdateOne) ClearAdmin() *GroupUpdateOne {
	guo.mutation.ClearAdmin()
	return guo
}

// Where appends a list predicates to the GroupUpdate builder.
func (guo *GroupUpdateOne) Where(ps ...predicate.Group) *GroupUpdateOne {
	guo.mutation.Where(ps...)
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GroupUpdateOne) Select(field string, fields ...string) *GroupUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Group entity.
func (guo *GroupUpdateOne) Save(ctx context.Context) (*Group, error) {
	return withHooks[*Group, GroupMutation](ctx, guo.sqlSave, guo.mutation, guo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GroupUpdateOne) SaveX(ctx context.Context) *Group {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GroupUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GroupUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (guo *GroupUpdateOne) sqlSave(ctx context.Context) (_node *Group, err error) {
	_spec := sqlgraph.NewUpdateSpec(group.Table, group.Columns, sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt))
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Group.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, group.FieldID)
		for _, f := range fields {
			if !group.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != group.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.Name(); ok {
		_spec.SetField(group.FieldName, field.TypeString, value)
	}
	if guo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !guo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.AdminCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   group.AdminTable,
			Columns: []string{group.AdminColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.AdminIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   group.AdminTable,
			Columns: []string{group.AdminColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Group{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	guo.mutation.done = true
	return _node, nil
}
