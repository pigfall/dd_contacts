// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Peanuttown/dd_contacts/ent/dept"
	"github.com/Peanuttown/dd_contacts/ent/user"
	"github.com/Peanuttown/dd_contacts/ent/userpropertyindept"
)

// DeptCreate is the builder for creating a Dept entity.
type DeptCreate struct {
	config
	mutation *DeptMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (dc *DeptCreate) SetName(s string) *DeptCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetID sets the "id" field.
func (dc *DeptCreate) SetID(u uint) *DeptCreate {
	dc.mutation.SetID(u)
	return dc
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (dc *DeptCreate) AddUserIDs(ids ...string) *DeptCreate {
	dc.mutation.AddUserIDs(ids...)
	return dc
}

// AddUsers adds the "users" edges to the User entity.
func (dc *DeptCreate) AddUsers(u ...*User) *DeptCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return dc.AddUserIDs(ids...)
}

// AddUserPropertiesInDeptIDs adds the "user_properties_in_dept" edge to the UserPropertyInDept entity by IDs.
func (dc *DeptCreate) AddUserPropertiesInDeptIDs(ids ...int) *DeptCreate {
	dc.mutation.AddUserPropertiesInDeptIDs(ids...)
	return dc
}

// AddUserPropertiesInDept adds the "user_properties_in_dept" edges to the UserPropertyInDept entity.
func (dc *DeptCreate) AddUserPropertiesInDept(u ...*UserPropertyInDept) *DeptCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return dc.AddUserPropertiesInDeptIDs(ids...)
}

// Mutation returns the DeptMutation object of the builder.
func (dc *DeptCreate) Mutation() *DeptMutation {
	return dc.mutation
}

// Save creates the Dept in the database.
func (dc *DeptCreate) Save(ctx context.Context) (*Dept, error) {
	var (
		err  error
		node *Dept
	)
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeptMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dc.check(); err != nil {
				return nil, err
			}
			dc.mutation = mutation
			node, err = dc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DeptCreate) SaveX(ctx context.Context) *Dept {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (dc *DeptCreate) check() error {
	if _, ok := dc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	return nil
}

func (dc *DeptCreate) sqlSave(ctx context.Context) (*Dept, error) {
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if _node.ID == 0 {
		id := _spec.ID.Value.(int64)
		_node.ID = uint(id)
	}
	return _node, nil
}

func (dc *DeptCreate) createSpec() (*Dept, *sqlgraph.CreateSpec) {
	var (
		_node = &Dept{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: dept.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: dept.FieldID,
			},
		}
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dept.FieldName,
		})
		_node.Name = value
	}
	if nodes := dc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   dept.UsersTable,
			Columns: dept.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.UserPropertiesInDeptIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dept.UserPropertiesInDeptTable,
			Columns: []string{dept.UserPropertiesInDeptColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userpropertyindept.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DeptCreateBulk is the builder for creating many Dept entities in bulk.
type DeptCreateBulk struct {
	config
	builders []*DeptCreate
}

// Save creates the Dept entities in the database.
func (dcb *DeptCreateBulk) Save(ctx context.Context) ([]*Dept, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Dept, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeptMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DeptCreateBulk) SaveX(ctx context.Context) []*Dept {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}