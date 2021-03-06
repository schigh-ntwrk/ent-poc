// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
	"github.com/schigh-ntwrk/ent-poc/internal/ent/clinic"
	"github.com/schigh-ntwrk/ent-poc/internal/ent/predicate"
	"github.com/schigh-ntwrk/ent-poc/internal/ent/veterinarian"
)

// ClinicUpdate is the builder for updating Clinic entities.
type ClinicUpdate struct {
	config
	hooks    []Hook
	mutation *ClinicMutation
}

// Where adds a new predicate for the builder.
func (cu *ClinicUpdate) Where(ps ...predicate.Clinic) *ClinicUpdate {
	cu.mutation.predicates = append(cu.mutation.predicates, ps...)
	return cu
}

// SetCreatedAt sets the created_at field.
func (cu *ClinicUpdate) SetCreatedAt(t time.Time) *ClinicUpdate {
	cu.mutation.SetCreatedAt(t)
	return cu
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (cu *ClinicUpdate) SetNillableCreatedAt(t *time.Time) *ClinicUpdate {
	if t != nil {
		cu.SetCreatedAt(*t)
	}
	return cu
}

// SetUpdatedAt sets the updated_at field.
func (cu *ClinicUpdate) SetUpdatedAt(t time.Time) *ClinicUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetRemovedAt sets the removed_at field.
func (cu *ClinicUpdate) SetRemovedAt(t time.Time) *ClinicUpdate {
	cu.mutation.SetRemovedAt(t)
	return cu
}

// SetNillableRemovedAt sets the removed_at field if the given value is not nil.
func (cu *ClinicUpdate) SetNillableRemovedAt(t *time.Time) *ClinicUpdate {
	if t != nil {
		cu.SetRemovedAt(*t)
	}
	return cu
}

// ClearRemovedAt clears the value of removed_at.
func (cu *ClinicUpdate) ClearRemovedAt() *ClinicUpdate {
	cu.mutation.ClearRemovedAt()
	return cu
}

// SetName sets the name field.
func (cu *ClinicUpdate) SetName(s string) *ClinicUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetAddress sets the address field.
func (cu *ClinicUpdate) SetAddress(s string) *ClinicUpdate {
	cu.mutation.SetAddress(s)
	return cu
}

// SetPhone sets the phone field.
func (cu *ClinicUpdate) SetPhone(s string) *ClinicUpdate {
	cu.mutation.SetPhone(s)
	return cu
}

// SetWebURL sets the web_url field.
func (cu *ClinicUpdate) SetWebURL(s string) *ClinicUpdate {
	cu.mutation.SetWebURL(s)
	return cu
}

// AddVeterinarianIDs adds the veterinarians edge to Veterinarian by ids.
func (cu *ClinicUpdate) AddVeterinarianIDs(ids ...uuid.UUID) *ClinicUpdate {
	cu.mutation.AddVeterinarianIDs(ids...)
	return cu
}

// AddVeterinarians adds the veterinarians edges to Veterinarian.
func (cu *ClinicUpdate) AddVeterinarians(v ...*Veterinarian) *ClinicUpdate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cu.AddVeterinarianIDs(ids...)
}

// Mutation returns the ClinicMutation object of the builder.
func (cu *ClinicUpdate) Mutation() *ClinicMutation {
	return cu.mutation
}

// ClearVeterinarians clears all "veterinarians" edges to type Veterinarian.
func (cu *ClinicUpdate) ClearVeterinarians() *ClinicUpdate {
	cu.mutation.ClearVeterinarians()
	return cu
}

// RemoveVeterinarianIDs removes the veterinarians edge to Veterinarian by ids.
func (cu *ClinicUpdate) RemoveVeterinarianIDs(ids ...uuid.UUID) *ClinicUpdate {
	cu.mutation.RemoveVeterinarianIDs(ids...)
	return cu
}

// RemoveVeterinarians removes veterinarians edges to Veterinarian.
func (cu *ClinicUpdate) RemoveVeterinarians(v ...*Veterinarian) *ClinicUpdate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cu.RemoveVeterinarianIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *ClinicUpdate) Save(ctx context.Context) (int, error) {
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
			mutation, ok := m.(*ClinicMutation)
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
func (cu *ClinicUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ClinicUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ClinicUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ClinicUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := clinic.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ClinicUpdate) check() error {
	if v, ok := cu.mutation.Name(); ok {
		if err := clinic.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := cu.mutation.Address(); ok {
		if err := clinic.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf("ent: validator failed for field \"address\": %w", err)}
		}
	}
	if v, ok := cu.mutation.Phone(); ok {
		if err := clinic.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf("ent: validator failed for field \"phone\": %w", err)}
		}
	}
	if v, ok := cu.mutation.WebURL(); ok {
		if err := clinic.WebURLValidator(v); err != nil {
			return &ValidationError{Name: "web_url", err: fmt.Errorf("ent: validator failed for field \"web_url\": %w", err)}
		}
	}
	return nil
}

func (cu *ClinicUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   clinic.Table,
			Columns: clinic.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: clinic.FieldID,
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
			Column: clinic.FieldCreatedAt,
		})
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: clinic.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.RemovedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: clinic.FieldRemovedAt,
		})
	}
	if cu.mutation.RemovedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: clinic.FieldRemovedAt,
		})
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinic.FieldName,
		})
	}
	if value, ok := cu.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinic.FieldAddress,
		})
	}
	if value, ok := cu.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinic.FieldPhone,
		})
	}
	if value, ok := cu.mutation.WebURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinic.FieldWebURL,
		})
	}
	if cu.mutation.VeterinariansCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   clinic.VeterinariansTable,
			Columns: []string{clinic.VeterinariansColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: veterinarian.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedVeterinariansIDs(); len(nodes) > 0 && !cu.mutation.VeterinariansCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   clinic.VeterinariansTable,
			Columns: []string{clinic.VeterinariansColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: veterinarian.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.VeterinariansIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   clinic.VeterinariansTable,
			Columns: []string{clinic.VeterinariansColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: veterinarian.FieldID,
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
			err = &NotFoundError{clinic.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ClinicUpdateOne is the builder for updating a single Clinic entity.
type ClinicUpdateOne struct {
	config
	hooks    []Hook
	mutation *ClinicMutation
}

// SetCreatedAt sets the created_at field.
func (cuo *ClinicUpdateOne) SetCreatedAt(t time.Time) *ClinicUpdateOne {
	cuo.mutation.SetCreatedAt(t)
	return cuo
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (cuo *ClinicUpdateOne) SetNillableCreatedAt(t *time.Time) *ClinicUpdateOne {
	if t != nil {
		cuo.SetCreatedAt(*t)
	}
	return cuo
}

// SetUpdatedAt sets the updated_at field.
func (cuo *ClinicUpdateOne) SetUpdatedAt(t time.Time) *ClinicUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetRemovedAt sets the removed_at field.
func (cuo *ClinicUpdateOne) SetRemovedAt(t time.Time) *ClinicUpdateOne {
	cuo.mutation.SetRemovedAt(t)
	return cuo
}

// SetNillableRemovedAt sets the removed_at field if the given value is not nil.
func (cuo *ClinicUpdateOne) SetNillableRemovedAt(t *time.Time) *ClinicUpdateOne {
	if t != nil {
		cuo.SetRemovedAt(*t)
	}
	return cuo
}

// ClearRemovedAt clears the value of removed_at.
func (cuo *ClinicUpdateOne) ClearRemovedAt() *ClinicUpdateOne {
	cuo.mutation.ClearRemovedAt()
	return cuo
}

// SetName sets the name field.
func (cuo *ClinicUpdateOne) SetName(s string) *ClinicUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetAddress sets the address field.
func (cuo *ClinicUpdateOne) SetAddress(s string) *ClinicUpdateOne {
	cuo.mutation.SetAddress(s)
	return cuo
}

// SetPhone sets the phone field.
func (cuo *ClinicUpdateOne) SetPhone(s string) *ClinicUpdateOne {
	cuo.mutation.SetPhone(s)
	return cuo
}

// SetWebURL sets the web_url field.
func (cuo *ClinicUpdateOne) SetWebURL(s string) *ClinicUpdateOne {
	cuo.mutation.SetWebURL(s)
	return cuo
}

// AddVeterinarianIDs adds the veterinarians edge to Veterinarian by ids.
func (cuo *ClinicUpdateOne) AddVeterinarianIDs(ids ...uuid.UUID) *ClinicUpdateOne {
	cuo.mutation.AddVeterinarianIDs(ids...)
	return cuo
}

// AddVeterinarians adds the veterinarians edges to Veterinarian.
func (cuo *ClinicUpdateOne) AddVeterinarians(v ...*Veterinarian) *ClinicUpdateOne {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cuo.AddVeterinarianIDs(ids...)
}

// Mutation returns the ClinicMutation object of the builder.
func (cuo *ClinicUpdateOne) Mutation() *ClinicMutation {
	return cuo.mutation
}

// ClearVeterinarians clears all "veterinarians" edges to type Veterinarian.
func (cuo *ClinicUpdateOne) ClearVeterinarians() *ClinicUpdateOne {
	cuo.mutation.ClearVeterinarians()
	return cuo
}

// RemoveVeterinarianIDs removes the veterinarians edge to Veterinarian by ids.
func (cuo *ClinicUpdateOne) RemoveVeterinarianIDs(ids ...uuid.UUID) *ClinicUpdateOne {
	cuo.mutation.RemoveVeterinarianIDs(ids...)
	return cuo
}

// RemoveVeterinarians removes veterinarians edges to Veterinarian.
func (cuo *ClinicUpdateOne) RemoveVeterinarians(v ...*Veterinarian) *ClinicUpdateOne {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cuo.RemoveVeterinarianIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (cuo *ClinicUpdateOne) Save(ctx context.Context) (*Clinic, error) {
	var (
		err  error
		node *Clinic
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClinicMutation)
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
func (cuo *ClinicUpdateOne) SaveX(ctx context.Context) *Clinic {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ClinicUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ClinicUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ClinicUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := clinic.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ClinicUpdateOne) check() error {
	if v, ok := cuo.mutation.Name(); ok {
		if err := clinic.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.Address(); ok {
		if err := clinic.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf("ent: validator failed for field \"address\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.Phone(); ok {
		if err := clinic.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf("ent: validator failed for field \"phone\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.WebURL(); ok {
		if err := clinic.WebURLValidator(v); err != nil {
			return &ValidationError{Name: "web_url", err: fmt.Errorf("ent: validator failed for field \"web_url\": %w", err)}
		}
	}
	return nil
}

func (cuo *ClinicUpdateOne) sqlSave(ctx context.Context) (_node *Clinic, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   clinic.Table,
			Columns: clinic.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: clinic.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Clinic.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: clinic.FieldCreatedAt,
		})
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: clinic.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.RemovedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: clinic.FieldRemovedAt,
		})
	}
	if cuo.mutation.RemovedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: clinic.FieldRemovedAt,
		})
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinic.FieldName,
		})
	}
	if value, ok := cuo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinic.FieldAddress,
		})
	}
	if value, ok := cuo.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinic.FieldPhone,
		})
	}
	if value, ok := cuo.mutation.WebURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinic.FieldWebURL,
		})
	}
	if cuo.mutation.VeterinariansCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   clinic.VeterinariansTable,
			Columns: []string{clinic.VeterinariansColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: veterinarian.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedVeterinariansIDs(); len(nodes) > 0 && !cuo.mutation.VeterinariansCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   clinic.VeterinariansTable,
			Columns: []string{clinic.VeterinariansColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: veterinarian.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.VeterinariansIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   clinic.VeterinariansTable,
			Columns: []string{clinic.VeterinariansColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: veterinarian.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Clinic{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{clinic.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
