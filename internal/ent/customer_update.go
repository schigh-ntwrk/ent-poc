// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
	"github.com/schigh-ntwrk/entc-poc/internal/ent/customer"
	"github.com/schigh-ntwrk/entc-poc/internal/ent/pet"
	"github.com/schigh-ntwrk/entc-poc/internal/ent/predicate"
	"github.com/schigh-ntwrk/entc-poc/internal/ent/user"
)

// CustomerUpdate is the builder for updating Customer entities.
type CustomerUpdate struct {
	config
	hooks    []Hook
	mutation *CustomerMutation
}

// Where adds a new predicate for the builder.
func (cu *CustomerUpdate) Where(ps ...predicate.Customer) *CustomerUpdate {
	cu.mutation.predicates = append(cu.mutation.predicates, ps...)
	return cu
}

// SetCreatedAt sets the created_at field.
func (cu *CustomerUpdate) SetCreatedAt(t time.Time) *CustomerUpdate {
	cu.mutation.SetCreatedAt(t)
	return cu
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableCreatedAt(t *time.Time) *CustomerUpdate {
	if t != nil {
		cu.SetCreatedAt(*t)
	}
	return cu
}

// SetUpdatedAt sets the updated_at field.
func (cu *CustomerUpdate) SetUpdatedAt(t time.Time) *CustomerUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetRemovedAt sets the removed_at field.
func (cu *CustomerUpdate) SetRemovedAt(t time.Time) *CustomerUpdate {
	cu.mutation.SetRemovedAt(t)
	return cu
}

// SetNillableRemovedAt sets the removed_at field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableRemovedAt(t *time.Time) *CustomerUpdate {
	if t != nil {
		cu.SetRemovedAt(*t)
	}
	return cu
}

// ClearRemovedAt clears the value of removed_at.
func (cu *CustomerUpdate) ClearRemovedAt() *CustomerUpdate {
	cu.mutation.ClearRemovedAt()
	return cu
}

// SetAddress sets the address field.
func (cu *CustomerUpdate) SetAddress(s string) *CustomerUpdate {
	cu.mutation.SetAddress(s)
	return cu
}

// SetPhone sets the phone field.
func (cu *CustomerUpdate) SetPhone(s string) *CustomerUpdate {
	cu.mutation.SetPhone(s)
	return cu
}

// SetUserID sets the user edge to User by id.
func (cu *CustomerUpdate) SetUserID(id uuid.UUID) *CustomerUpdate {
	cu.mutation.SetUserID(id)
	return cu
}

// SetUser sets the user edge to User.
func (cu *CustomerUpdate) SetUser(u *User) *CustomerUpdate {
	return cu.SetUserID(u.ID)
}

// AddPetIDs adds the pets edge to Pet by ids.
func (cu *CustomerUpdate) AddPetIDs(ids ...uuid.UUID) *CustomerUpdate {
	cu.mutation.AddPetIDs(ids...)
	return cu
}

// AddPets adds the pets edges to Pet.
func (cu *CustomerUpdate) AddPets(p ...*Pet) *CustomerUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.AddPetIDs(ids...)
}

// Mutation returns the CustomerMutation object of the builder.
func (cu *CustomerUpdate) Mutation() *CustomerMutation {
	return cu.mutation
}

// ClearUser clears the "user" edge to type User.
func (cu *CustomerUpdate) ClearUser() *CustomerUpdate {
	cu.mutation.ClearUser()
	return cu
}

// ClearPets clears all "pets" edges to type Pet.
func (cu *CustomerUpdate) ClearPets() *CustomerUpdate {
	cu.mutation.ClearPets()
	return cu
}

// RemovePetIDs removes the pets edge to Pet by ids.
func (cu *CustomerUpdate) RemovePetIDs(ids ...uuid.UUID) *CustomerUpdate {
	cu.mutation.RemovePetIDs(ids...)
	return cu
}

// RemovePets removes pets edges to Pet.
func (cu *CustomerUpdate) RemovePets(p ...*Pet) *CustomerUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.RemovePetIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *CustomerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CustomerUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CustomerUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CustomerUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CustomerUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := customer.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CustomerUpdate) check() error {
	if v, ok := cu.mutation.Address(); ok {
		if err := customer.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf("ent: validator failed for field \"address\": %w", err)}
		}
	}
	if v, ok := cu.mutation.Phone(); ok {
		if err := customer.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf("ent: validator failed for field \"phone\": %w", err)}
		}
	}
	if _, ok := cu.mutation.UserID(); cu.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (cu *CustomerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   customer.Table,
			Columns: customer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: customer.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customer.FieldCreatedAt,
		})
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customer.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.RemovedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customer.FieldRemovedAt,
		})
	}
	if cu.mutation.RemovedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: customer.FieldRemovedAt,
		})
	}
	if value, ok := cu.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldAddress,
		})
	}
	if value, ok := cu.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldPhone,
		})
	}
	if cu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   customer.UserTable,
			Columns: []string{customer.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   customer.UserTable,
			Columns: []string{customer.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.PetsTable,
			Columns: []string{customer.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: pet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedPetsIDs(); len(nodes) > 0 && !cu.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.PetsTable,
			Columns: []string{customer.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: pet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.PetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.PetsTable,
			Columns: []string{customer.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: pet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customer.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CustomerUpdateOne is the builder for updating a single Customer entity.
type CustomerUpdateOne struct {
	config
	hooks    []Hook
	mutation *CustomerMutation
}

// SetCreatedAt sets the created_at field.
func (cuo *CustomerUpdateOne) SetCreatedAt(t time.Time) *CustomerUpdateOne {
	cuo.mutation.SetCreatedAt(t)
	return cuo
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableCreatedAt(t *time.Time) *CustomerUpdateOne {
	if t != nil {
		cuo.SetCreatedAt(*t)
	}
	return cuo
}

// SetUpdatedAt sets the updated_at field.
func (cuo *CustomerUpdateOne) SetUpdatedAt(t time.Time) *CustomerUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetRemovedAt sets the removed_at field.
func (cuo *CustomerUpdateOne) SetRemovedAt(t time.Time) *CustomerUpdateOne {
	cuo.mutation.SetRemovedAt(t)
	return cuo
}

// SetNillableRemovedAt sets the removed_at field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableRemovedAt(t *time.Time) *CustomerUpdateOne {
	if t != nil {
		cuo.SetRemovedAt(*t)
	}
	return cuo
}

// ClearRemovedAt clears the value of removed_at.
func (cuo *CustomerUpdateOne) ClearRemovedAt() *CustomerUpdateOne {
	cuo.mutation.ClearRemovedAt()
	return cuo
}

// SetAddress sets the address field.
func (cuo *CustomerUpdateOne) SetAddress(s string) *CustomerUpdateOne {
	cuo.mutation.SetAddress(s)
	return cuo
}

// SetPhone sets the phone field.
func (cuo *CustomerUpdateOne) SetPhone(s string) *CustomerUpdateOne {
	cuo.mutation.SetPhone(s)
	return cuo
}

// SetUserID sets the user edge to User by id.
func (cuo *CustomerUpdateOne) SetUserID(id uuid.UUID) *CustomerUpdateOne {
	cuo.mutation.SetUserID(id)
	return cuo
}

// SetUser sets the user edge to User.
func (cuo *CustomerUpdateOne) SetUser(u *User) *CustomerUpdateOne {
	return cuo.SetUserID(u.ID)
}

// AddPetIDs adds the pets edge to Pet by ids.
func (cuo *CustomerUpdateOne) AddPetIDs(ids ...uuid.UUID) *CustomerUpdateOne {
	cuo.mutation.AddPetIDs(ids...)
	return cuo
}

// AddPets adds the pets edges to Pet.
func (cuo *CustomerUpdateOne) AddPets(p ...*Pet) *CustomerUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.AddPetIDs(ids...)
}

// Mutation returns the CustomerMutation object of the builder.
func (cuo *CustomerUpdateOne) Mutation() *CustomerMutation {
	return cuo.mutation
}

// ClearUser clears the "user" edge to type User.
func (cuo *CustomerUpdateOne) ClearUser() *CustomerUpdateOne {
	cuo.mutation.ClearUser()
	return cuo
}

// ClearPets clears all "pets" edges to type Pet.
func (cuo *CustomerUpdateOne) ClearPets() *CustomerUpdateOne {
	cuo.mutation.ClearPets()
	return cuo
}

// RemovePetIDs removes the pets edge to Pet by ids.
func (cuo *CustomerUpdateOne) RemovePetIDs(ids ...uuid.UUID) *CustomerUpdateOne {
	cuo.mutation.RemovePetIDs(ids...)
	return cuo
}

// RemovePets removes pets edges to Pet.
func (cuo *CustomerUpdateOne) RemovePets(p ...*Pet) *CustomerUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.RemovePetIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (cuo *CustomerUpdateOne) Save(ctx context.Context) (*Customer, error) {
	var (
		err  error
		node *Customer
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CustomerUpdateOne) SaveX(ctx context.Context) *Customer {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CustomerUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CustomerUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CustomerUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := customer.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CustomerUpdateOne) check() error {
	if v, ok := cuo.mutation.Address(); ok {
		if err := customer.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf("ent: validator failed for field \"address\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.Phone(); ok {
		if err := customer.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf("ent: validator failed for field \"phone\": %w", err)}
		}
	}
	if _, ok := cuo.mutation.UserID(); cuo.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (cuo *CustomerUpdateOne) sqlSave(ctx context.Context) (_node *Customer, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   customer.Table,
			Columns: customer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: customer.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Customer.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customer.FieldCreatedAt,
		})
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customer.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.RemovedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customer.FieldRemovedAt,
		})
	}
	if cuo.mutation.RemovedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: customer.FieldRemovedAt,
		})
	}
	if value, ok := cuo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldAddress,
		})
	}
	if value, ok := cuo.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldPhone,
		})
	}
	if cuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   customer.UserTable,
			Columns: []string{customer.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   customer.UserTable,
			Columns: []string{customer.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.PetsTable,
			Columns: []string{customer.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: pet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedPetsIDs(); len(nodes) > 0 && !cuo.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.PetsTable,
			Columns: []string{customer.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: pet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.PetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.PetsTable,
			Columns: []string{customer.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: pet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Customer{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customer.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
