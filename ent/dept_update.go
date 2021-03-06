// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Peanuttown/dd_contacts/ent/dept"
	"github.com/Peanuttown/dd_contacts/ent/predicate"
	"github.com/Peanuttown/dd_contacts/ent/user"
	"github.com/Peanuttown/dd_contacts/ent/userpropertyindept"
)

// DeptUpdate is the builder for updating Dept entities.
type DeptUpdate struct {
	config
	hooks    []Hook
	mutation *DeptMutation
}

// Where adds a new predicate for the DeptUpdate builder.
func (du *DeptUpdate) Where(ps ...predicate.Dept) *DeptUpdate {
	du.mutation.predicates = append(du.mutation.predicates, ps...)
	return du
}

// SetName sets the "name" field.
func (du *DeptUpdate) SetName(s string) *DeptUpdate {
	du.mutation.SetName(s)
	return du
}

// SetGeneration sets the "generation" field.
func (du *DeptUpdate) SetGeneration(u uint) *DeptUpdate {
	du.mutation.ResetGeneration()
	du.mutation.SetGeneration(u)
	return du
}

// SetNillableGeneration sets the "generation" field if the given value is not nil.
func (du *DeptUpdate) SetNillableGeneration(u *uint) *DeptUpdate {
	if u != nil {
		du.SetGeneration(*u)
	}
	return du
}

// AddGeneration adds u to the "generation" field.
func (du *DeptUpdate) AddGeneration(u uint) *DeptUpdate {
	du.mutation.AddGeneration(u)
	return du
}

// ClearGeneration clears the value of the "generation" field.
func (du *DeptUpdate) ClearGeneration() *DeptUpdate {
	du.mutation.ClearGeneration()
	return du
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (du *DeptUpdate) AddUserIDs(ids ...string) *DeptUpdate {
	du.mutation.AddUserIDs(ids...)
	return du
}

// AddUsers adds the "users" edges to the User entity.
func (du *DeptUpdate) AddUsers(u ...*User) *DeptUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return du.AddUserIDs(ids...)
}

// AddUserPropertiesInDeptIDs adds the "user_properties_in_dept" edge to the UserPropertyInDept entity by IDs.
func (du *DeptUpdate) AddUserPropertiesInDeptIDs(ids ...int) *DeptUpdate {
	du.mutation.AddUserPropertiesInDeptIDs(ids...)
	return du
}

// AddUserPropertiesInDept adds the "user_properties_in_dept" edges to the UserPropertyInDept entity.
func (du *DeptUpdate) AddUserPropertiesInDept(u ...*UserPropertyInDept) *DeptUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return du.AddUserPropertiesInDeptIDs(ids...)
}

// SetParentID sets the "parent" edge to the Dept entity by ID.
func (du *DeptUpdate) SetParentID(id uint) *DeptUpdate {
	du.mutation.SetParentID(id)
	return du
}

// SetNillableParentID sets the "parent" edge to the Dept entity by ID if the given value is not nil.
func (du *DeptUpdate) SetNillableParentID(id *uint) *DeptUpdate {
	if id != nil {
		du = du.SetParentID(*id)
	}
	return du
}

// SetParent sets the "parent" edge to the Dept entity.
func (du *DeptUpdate) SetParent(d *Dept) *DeptUpdate {
	return du.SetParentID(d.ID)
}

// AddSubDeptIDs adds the "sub_depts" edge to the Dept entity by IDs.
func (du *DeptUpdate) AddSubDeptIDs(ids ...uint) *DeptUpdate {
	du.mutation.AddSubDeptIDs(ids...)
	return du
}

// AddSubDepts adds the "sub_depts" edges to the Dept entity.
func (du *DeptUpdate) AddSubDepts(d ...*Dept) *DeptUpdate {
	ids := make([]uint, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.AddSubDeptIDs(ids...)
}

// Mutation returns the DeptMutation object of the builder.
func (du *DeptUpdate) Mutation() *DeptMutation {
	return du.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (du *DeptUpdate) ClearUsers() *DeptUpdate {
	du.mutation.ClearUsers()
	return du
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (du *DeptUpdate) RemoveUserIDs(ids ...string) *DeptUpdate {
	du.mutation.RemoveUserIDs(ids...)
	return du
}

// RemoveUsers removes "users" edges to User entities.
func (du *DeptUpdate) RemoveUsers(u ...*User) *DeptUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return du.RemoveUserIDs(ids...)
}

// ClearUserPropertiesInDept clears all "user_properties_in_dept" edges to the UserPropertyInDept entity.
func (du *DeptUpdate) ClearUserPropertiesInDept() *DeptUpdate {
	du.mutation.ClearUserPropertiesInDept()
	return du
}

// RemoveUserPropertiesInDeptIDs removes the "user_properties_in_dept" edge to UserPropertyInDept entities by IDs.
func (du *DeptUpdate) RemoveUserPropertiesInDeptIDs(ids ...int) *DeptUpdate {
	du.mutation.RemoveUserPropertiesInDeptIDs(ids...)
	return du
}

// RemoveUserPropertiesInDept removes "user_properties_in_dept" edges to UserPropertyInDept entities.
func (du *DeptUpdate) RemoveUserPropertiesInDept(u ...*UserPropertyInDept) *DeptUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return du.RemoveUserPropertiesInDeptIDs(ids...)
}

// ClearParent clears the "parent" edge to the Dept entity.
func (du *DeptUpdate) ClearParent() *DeptUpdate {
	du.mutation.ClearParent()
	return du
}

// ClearSubDepts clears all "sub_depts" edges to the Dept entity.
func (du *DeptUpdate) ClearSubDepts() *DeptUpdate {
	du.mutation.ClearSubDepts()
	return du
}

// RemoveSubDeptIDs removes the "sub_depts" edge to Dept entities by IDs.
func (du *DeptUpdate) RemoveSubDeptIDs(ids ...uint) *DeptUpdate {
	du.mutation.RemoveSubDeptIDs(ids...)
	return du
}

// RemoveSubDepts removes "sub_depts" edges to Dept entities.
func (du *DeptUpdate) RemoveSubDepts(d ...*Dept) *DeptUpdate {
	ids := make([]uint, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.RemoveSubDeptIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DeptUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeptMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DeptUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DeptUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DeptUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DeptUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dept.Table,
			Columns: dept.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: dept.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dept.FieldName,
		})
	}
	if value, ok := du.mutation.Generation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: dept.FieldGeneration,
		})
	}
	if value, ok := du.mutation.AddedGeneration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: dept.FieldGeneration,
		})
	}
	if du.mutation.GenerationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Column: dept.FieldGeneration,
		})
	}
	if du.mutation.UsersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedUsersIDs(); len(nodes) > 0 && !du.mutation.UsersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.UsersIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.UserPropertiesInDeptCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedUserPropertiesInDeptIDs(); len(nodes) > 0 && !du.mutation.UserPropertiesInDeptCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.UserPropertiesInDeptIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dept.ParentTable,
			Columns: []string{dept.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: dept.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dept.ParentTable,
			Columns: []string{dept.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: dept.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.SubDeptsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dept.SubDeptsTable,
			Columns: []string{dept.SubDeptsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: dept.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedSubDeptsIDs(); len(nodes) > 0 && !du.mutation.SubDeptsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dept.SubDeptsTable,
			Columns: []string{dept.SubDeptsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: dept.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.SubDeptsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dept.SubDeptsTable,
			Columns: []string{dept.SubDeptsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: dept.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dept.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DeptUpdateOne is the builder for updating a single Dept entity.
type DeptUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeptMutation
}

// SetName sets the "name" field.
func (duo *DeptUpdateOne) SetName(s string) *DeptUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// SetGeneration sets the "generation" field.
func (duo *DeptUpdateOne) SetGeneration(u uint) *DeptUpdateOne {
	duo.mutation.ResetGeneration()
	duo.mutation.SetGeneration(u)
	return duo
}

// SetNillableGeneration sets the "generation" field if the given value is not nil.
func (duo *DeptUpdateOne) SetNillableGeneration(u *uint) *DeptUpdateOne {
	if u != nil {
		duo.SetGeneration(*u)
	}
	return duo
}

// AddGeneration adds u to the "generation" field.
func (duo *DeptUpdateOne) AddGeneration(u uint) *DeptUpdateOne {
	duo.mutation.AddGeneration(u)
	return duo
}

// ClearGeneration clears the value of the "generation" field.
func (duo *DeptUpdateOne) ClearGeneration() *DeptUpdateOne {
	duo.mutation.ClearGeneration()
	return duo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (duo *DeptUpdateOne) AddUserIDs(ids ...string) *DeptUpdateOne {
	duo.mutation.AddUserIDs(ids...)
	return duo
}

// AddUsers adds the "users" edges to the User entity.
func (duo *DeptUpdateOne) AddUsers(u ...*User) *DeptUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return duo.AddUserIDs(ids...)
}

// AddUserPropertiesInDeptIDs adds the "user_properties_in_dept" edge to the UserPropertyInDept entity by IDs.
func (duo *DeptUpdateOne) AddUserPropertiesInDeptIDs(ids ...int) *DeptUpdateOne {
	duo.mutation.AddUserPropertiesInDeptIDs(ids...)
	return duo
}

// AddUserPropertiesInDept adds the "user_properties_in_dept" edges to the UserPropertyInDept entity.
func (duo *DeptUpdateOne) AddUserPropertiesInDept(u ...*UserPropertyInDept) *DeptUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return duo.AddUserPropertiesInDeptIDs(ids...)
}

// SetParentID sets the "parent" edge to the Dept entity by ID.
func (duo *DeptUpdateOne) SetParentID(id uint) *DeptUpdateOne {
	duo.mutation.SetParentID(id)
	return duo
}

// SetNillableParentID sets the "parent" edge to the Dept entity by ID if the given value is not nil.
func (duo *DeptUpdateOne) SetNillableParentID(id *uint) *DeptUpdateOne {
	if id != nil {
		duo = duo.SetParentID(*id)
	}
	return duo
}

// SetParent sets the "parent" edge to the Dept entity.
func (duo *DeptUpdateOne) SetParent(d *Dept) *DeptUpdateOne {
	return duo.SetParentID(d.ID)
}

// AddSubDeptIDs adds the "sub_depts" edge to the Dept entity by IDs.
func (duo *DeptUpdateOne) AddSubDeptIDs(ids ...uint) *DeptUpdateOne {
	duo.mutation.AddSubDeptIDs(ids...)
	return duo
}

// AddSubDepts adds the "sub_depts" edges to the Dept entity.
func (duo *DeptUpdateOne) AddSubDepts(d ...*Dept) *DeptUpdateOne {
	ids := make([]uint, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.AddSubDeptIDs(ids...)
}

// Mutation returns the DeptMutation object of the builder.
func (duo *DeptUpdateOne) Mutation() *DeptMutation {
	return duo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (duo *DeptUpdateOne) ClearUsers() *DeptUpdateOne {
	duo.mutation.ClearUsers()
	return duo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (duo *DeptUpdateOne) RemoveUserIDs(ids ...string) *DeptUpdateOne {
	duo.mutation.RemoveUserIDs(ids...)
	return duo
}

// RemoveUsers removes "users" edges to User entities.
func (duo *DeptUpdateOne) RemoveUsers(u ...*User) *DeptUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return duo.RemoveUserIDs(ids...)
}

// ClearUserPropertiesInDept clears all "user_properties_in_dept" edges to the UserPropertyInDept entity.
func (duo *DeptUpdateOne) ClearUserPropertiesInDept() *DeptUpdateOne {
	duo.mutation.ClearUserPropertiesInDept()
	return duo
}

// RemoveUserPropertiesInDeptIDs removes the "user_properties_in_dept" edge to UserPropertyInDept entities by IDs.
func (duo *DeptUpdateOne) RemoveUserPropertiesInDeptIDs(ids ...int) *DeptUpdateOne {
	duo.mutation.RemoveUserPropertiesInDeptIDs(ids...)
	return duo
}

// RemoveUserPropertiesInDept removes "user_properties_in_dept" edges to UserPropertyInDept entities.
func (duo *DeptUpdateOne) RemoveUserPropertiesInDept(u ...*UserPropertyInDept) *DeptUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return duo.RemoveUserPropertiesInDeptIDs(ids...)
}

// ClearParent clears the "parent" edge to the Dept entity.
func (duo *DeptUpdateOne) ClearParent() *DeptUpdateOne {
	duo.mutation.ClearParent()
	return duo
}

// ClearSubDepts clears all "sub_depts" edges to the Dept entity.
func (duo *DeptUpdateOne) ClearSubDepts() *DeptUpdateOne {
	duo.mutation.ClearSubDepts()
	return duo
}

// RemoveSubDeptIDs removes the "sub_depts" edge to Dept entities by IDs.
func (duo *DeptUpdateOne) RemoveSubDeptIDs(ids ...uint) *DeptUpdateOne {
	duo.mutation.RemoveSubDeptIDs(ids...)
	return duo
}

// RemoveSubDepts removes "sub_depts" edges to Dept entities.
func (duo *DeptUpdateOne) RemoveSubDepts(d ...*Dept) *DeptUpdateOne {
	ids := make([]uint, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.RemoveSubDeptIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DeptUpdateOne) Select(field string, fields ...string) *DeptUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Dept entity.
func (duo *DeptUpdateOne) Save(ctx context.Context) (*Dept, error) {
	var (
		err  error
		node *Dept
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeptMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DeptUpdateOne) SaveX(ctx context.Context) *Dept {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DeptUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DeptUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DeptUpdateOne) sqlSave(ctx context.Context) (_node *Dept, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dept.Table,
			Columns: dept.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: dept.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Dept.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dept.FieldID)
		for _, f := range fields {
			if !dept.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != dept.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dept.FieldName,
		})
	}
	if value, ok := duo.mutation.Generation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: dept.FieldGeneration,
		})
	}
	if value, ok := duo.mutation.AddedGeneration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: dept.FieldGeneration,
		})
	}
	if duo.mutation.GenerationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Column: dept.FieldGeneration,
		})
	}
	if duo.mutation.UsersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !duo.mutation.UsersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.UsersIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.UserPropertiesInDeptCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedUserPropertiesInDeptIDs(); len(nodes) > 0 && !duo.mutation.UserPropertiesInDeptCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.UserPropertiesInDeptIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dept.ParentTable,
			Columns: []string{dept.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: dept.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dept.ParentTable,
			Columns: []string{dept.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: dept.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.SubDeptsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dept.SubDeptsTable,
			Columns: []string{dept.SubDeptsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: dept.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedSubDeptsIDs(); len(nodes) > 0 && !duo.mutation.SubDeptsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dept.SubDeptsTable,
			Columns: []string{dept.SubDeptsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: dept.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.SubDeptsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dept.SubDeptsTable,
			Columns: []string{dept.SubDeptsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: dept.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Dept{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dept.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
