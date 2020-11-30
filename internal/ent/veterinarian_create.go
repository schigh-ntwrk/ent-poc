// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
	"github.com/schigh-ntwrk/ent-poc/internal/ent/appointment"
	"github.com/schigh-ntwrk/ent-poc/internal/ent/clinic"
	"github.com/schigh-ntwrk/ent-poc/internal/ent/user"
	"github.com/schigh-ntwrk/ent-poc/internal/ent/veterinarian"
)

// VeterinarianCreate is the builder for creating a Veterinarian entity.
type VeterinarianCreate struct {
	config
	mutation *VeterinarianMutation
	hooks    []Hook
}

// SetCreatedAt sets the created_at field.
func (vc *VeterinarianCreate) SetCreatedAt(t time.Time) *VeterinarianCreate {
	vc.mutation.SetCreatedAt(t)
	return vc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (vc *VeterinarianCreate) SetNillableCreatedAt(t *time.Time) *VeterinarianCreate {
	if t != nil {
		vc.SetCreatedAt(*t)
	}
	return vc
}

// SetUpdatedAt sets the updated_at field.
func (vc *VeterinarianCreate) SetUpdatedAt(t time.Time) *VeterinarianCreate {
	vc.mutation.SetUpdatedAt(t)
	return vc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (vc *VeterinarianCreate) SetNillableUpdatedAt(t *time.Time) *VeterinarianCreate {
	if t != nil {
		vc.SetUpdatedAt(*t)
	}
	return vc
}

// SetRemovedAt sets the removed_at field.
func (vc *VeterinarianCreate) SetRemovedAt(t time.Time) *VeterinarianCreate {
	vc.mutation.SetRemovedAt(t)
	return vc
}

// SetNillableRemovedAt sets the removed_at field if the given value is not nil.
func (vc *VeterinarianCreate) SetNillableRemovedAt(t *time.Time) *VeterinarianCreate {
	if t != nil {
		vc.SetRemovedAt(*t)
	}
	return vc
}

// SetPhone sets the phone field.
func (vc *VeterinarianCreate) SetPhone(s string) *VeterinarianCreate {
	vc.mutation.SetPhone(s)
	return vc
}

// SetID sets the id field.
func (vc *VeterinarianCreate) SetID(u uuid.UUID) *VeterinarianCreate {
	vc.mutation.SetID(u)
	return vc
}

// SetUserID sets the user edge to User by id.
func (vc *VeterinarianCreate) SetUserID(id uuid.UUID) *VeterinarianCreate {
	vc.mutation.SetUserID(id)
	return vc
}

// SetUser sets the user edge to User.
func (vc *VeterinarianCreate) SetUser(u *User) *VeterinarianCreate {
	return vc.SetUserID(u.ID)
}

// SetClinicID sets the clinic edge to Clinic by id.
func (vc *VeterinarianCreate) SetClinicID(id uuid.UUID) *VeterinarianCreate {
	vc.mutation.SetClinicID(id)
	return vc
}

// SetClinic sets the clinic edge to Clinic.
func (vc *VeterinarianCreate) SetClinic(c *Clinic) *VeterinarianCreate {
	return vc.SetClinicID(c.ID)
}

// AddAppointmentIDs adds the appointments edge to Appointment by ids.
func (vc *VeterinarianCreate) AddAppointmentIDs(ids ...uuid.UUID) *VeterinarianCreate {
	vc.mutation.AddAppointmentIDs(ids...)
	return vc
}

// AddAppointments adds the appointments edges to Appointment.
func (vc *VeterinarianCreate) AddAppointments(a ...*Appointment) *VeterinarianCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return vc.AddAppointmentIDs(ids...)
}

// Mutation returns the VeterinarianMutation object of the builder.
func (vc *VeterinarianCreate) Mutation() *VeterinarianMutation {
	return vc.mutation
}

// Save creates the Veterinarian in the database.
func (vc *VeterinarianCreate) Save(ctx context.Context) (*Veterinarian, error) {
	var (
		err  error
		node *Veterinarian
	)
	vc.defaults()
	if len(vc.hooks) == 0 {
		if err = vc.check(); err != nil {
			return nil, err
		}
		node, err = vc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VeterinarianMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vc.check(); err != nil {
				return nil, err
			}
			vc.mutation = mutation
			node, err = vc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(vc.hooks) - 1; i >= 0; i-- {
			mut = vc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VeterinarianCreate) SaveX(ctx context.Context) *Veterinarian {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (vc *VeterinarianCreate) defaults() {
	if _, ok := vc.mutation.CreatedAt(); !ok {
		v := veterinarian.DefaultCreatedAt()
		vc.mutation.SetCreatedAt(v)
	}
	if _, ok := vc.mutation.UpdatedAt(); !ok {
		v := veterinarian.DefaultUpdatedAt()
		vc.mutation.SetUpdatedAt(v)
	}
	if _, ok := vc.mutation.ID(); !ok {
		v := veterinarian.DefaultID()
		vc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VeterinarianCreate) check() error {
	if _, ok := vc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := vc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	if _, ok := vc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New("ent: missing required field \"phone\"")}
	}
	if v, ok := vc.mutation.Phone(); ok {
		if err := veterinarian.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf("ent: validator failed for field \"phone\": %w", err)}
		}
	}
	if _, ok := vc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New("ent: missing required edge \"user\"")}
	}
	if _, ok := vc.mutation.ClinicID(); !ok {
		return &ValidationError{Name: "clinic", err: errors.New("ent: missing required edge \"clinic\"")}
	}
	return nil
}

func (vc *VeterinarianCreate) sqlSave(ctx context.Context) (*Veterinarian, error) {
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (vc *VeterinarianCreate) createSpec() (*Veterinarian, *sqlgraph.CreateSpec) {
	var (
		_node = &Veterinarian{config: vc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: veterinarian.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: veterinarian.FieldID,
			},
		}
	)
	if id, ok := vc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := vc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: veterinarian.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := vc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: veterinarian.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := vc.mutation.RemovedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: veterinarian.FieldRemovedAt,
		})
		_node.RemovedAt = &value
	}
	if value, ok := vc.mutation.Phone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: veterinarian.FieldPhone,
		})
		_node.Phone = value
	}
	if nodes := vc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   veterinarian.UserTable,
			Columns: []string{veterinarian.UserColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.ClinicIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   veterinarian.ClinicTable,
			Columns: []string{veterinarian.ClinicColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: clinic.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.AppointmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   veterinarian.AppointmentsTable,
			Columns: []string{veterinarian.AppointmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: appointment.FieldID,
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

// VeterinarianCreateBulk is the builder for creating a bulk of Veterinarian entities.
type VeterinarianCreateBulk struct {
	config
	builders []*VeterinarianCreate
}

// Save creates the Veterinarian entities in the database.
func (vcb *VeterinarianCreateBulk) Save(ctx context.Context) ([]*Veterinarian, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Veterinarian, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VeterinarianMutation)
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
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
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
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (vcb *VeterinarianCreateBulk) SaveX(ctx context.Context) []*Veterinarian {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
