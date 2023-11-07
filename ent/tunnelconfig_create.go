// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"rahnit-rmm/ent/device"
	"rahnit-rmm/ent/tunnelconfig"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TunnelConfigCreate is the builder for creating a TunnelConfig entity.
type TunnelConfigCreate struct {
	config
	mutation *TunnelConfigMutation
	hooks    []Hook
}

// SetDeviceID sets the "device" edge to the Device entity by ID.
func (tcc *TunnelConfigCreate) SetDeviceID(id int) *TunnelConfigCreate {
	tcc.mutation.SetDeviceID(id)
	return tcc
}

// SetDevice sets the "device" edge to the Device entity.
func (tcc *TunnelConfigCreate) SetDevice(d *Device) *TunnelConfigCreate {
	return tcc.SetDeviceID(d.ID)
}

// Mutation returns the TunnelConfigMutation object of the builder.
func (tcc *TunnelConfigCreate) Mutation() *TunnelConfigMutation {
	return tcc.mutation
}

// Save creates the TunnelConfig in the database.
func (tcc *TunnelConfigCreate) Save(ctx context.Context) (*TunnelConfig, error) {
	return withHooks(ctx, tcc.sqlSave, tcc.mutation, tcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tcc *TunnelConfigCreate) SaveX(ctx context.Context) *TunnelConfig {
	v, err := tcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcc *TunnelConfigCreate) Exec(ctx context.Context) error {
	_, err := tcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcc *TunnelConfigCreate) ExecX(ctx context.Context) {
	if err := tcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tcc *TunnelConfigCreate) check() error {
	if _, ok := tcc.mutation.DeviceID(); !ok {
		return &ValidationError{Name: "device", err: errors.New(`ent: missing required edge "TunnelConfig.device"`)}
	}
	return nil
}

func (tcc *TunnelConfigCreate) sqlSave(ctx context.Context) (*TunnelConfig, error) {
	if err := tcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tcc.mutation.id = &_node.ID
	tcc.mutation.done = true
	return _node, nil
}

func (tcc *TunnelConfigCreate) createSpec() (*TunnelConfig, *sqlgraph.CreateSpec) {
	var (
		_node = &TunnelConfig{config: tcc.config}
		_spec = sqlgraph.NewCreateSpec(tunnelconfig.Table, sqlgraph.NewFieldSpec(tunnelconfig.FieldID, field.TypeInt))
	)
	if nodes := tcc.mutation.DeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   tunnelconfig.DeviceTable,
			Columns: []string{tunnelconfig.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TunnelConfigCreateBulk is the builder for creating many TunnelConfig entities in bulk.
type TunnelConfigCreateBulk struct {
	config
	err      error
	builders []*TunnelConfigCreate
}

// Save creates the TunnelConfig entities in the database.
func (tccb *TunnelConfigCreateBulk) Save(ctx context.Context) ([]*TunnelConfig, error) {
	if tccb.err != nil {
		return nil, tccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tccb.builders))
	nodes := make([]*TunnelConfig, len(tccb.builders))
	mutators := make([]Mutator, len(tccb.builders))
	for i := range tccb.builders {
		func(i int, root context.Context) {
			builder := tccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TunnelConfigMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tccb *TunnelConfigCreateBulk) SaveX(ctx context.Context) []*TunnelConfig {
	v, err := tccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tccb *TunnelConfigCreateBulk) Exec(ctx context.Context) error {
	_, err := tccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tccb *TunnelConfigCreateBulk) ExecX(ctx context.Context) {
	if err := tccb.Exec(ctx); err != nil {
		panic(err)
	}
}
