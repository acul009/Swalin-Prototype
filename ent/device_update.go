// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"github.com/rahn-it/svalin/ent/device"
	"github.com/rahn-it/svalin/ent/hostconfig"
	"github.com/rahn-it/svalin/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DeviceUpdate is the builder for updating Device entities.
type DeviceUpdate struct {
	config
	hooks    []Hook
	mutation *DeviceMutation
}

// Where appends a list predicates to the DeviceUpdate builder.
func (du *DeviceUpdate) Where(ps ...predicate.Device) *DeviceUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetCertificate sets the "certificate" field.
func (du *DeviceUpdate) SetCertificate(s string) *DeviceUpdate {
	du.mutation.SetCertificate(s)
	return du
}

// SetNillableCertificate sets the "certificate" field if the given value is not nil.
func (du *DeviceUpdate) SetNillableCertificate(s *string) *DeviceUpdate {
	if s != nil {
		du.SetCertificate(*s)
	}
	return du
}

// AddConfigIDs adds the "configs" edge to the HostConfig entity by IDs.
func (du *DeviceUpdate) AddConfigIDs(ids ...int) *DeviceUpdate {
	du.mutation.AddConfigIDs(ids...)
	return du
}

// AddConfigs adds the "configs" edges to the HostConfig entity.
func (du *DeviceUpdate) AddConfigs(h ...*HostConfig) *DeviceUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return du.AddConfigIDs(ids...)
}

// Mutation returns the DeviceMutation object of the builder.
func (du *DeviceUpdate) Mutation() *DeviceMutation {
	return du.mutation
}

// ClearConfigs clears all "configs" edges to the HostConfig entity.
func (du *DeviceUpdate) ClearConfigs() *DeviceUpdate {
	du.mutation.ClearConfigs()
	return du
}

// RemoveConfigIDs removes the "configs" edge to HostConfig entities by IDs.
func (du *DeviceUpdate) RemoveConfigIDs(ids ...int) *DeviceUpdate {
	du.mutation.RemoveConfigIDs(ids...)
	return du
}

// RemoveConfigs removes "configs" edges to HostConfig entities.
func (du *DeviceUpdate) RemoveConfigs(h ...*HostConfig) *DeviceUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return du.RemoveConfigIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DeviceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DeviceUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DeviceUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DeviceUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DeviceUpdate) check() error {
	if v, ok := du.mutation.Certificate(); ok {
		if err := device.CertificateValidator(v); err != nil {
			return &ValidationError{Name: "certificate", err: fmt.Errorf(`ent: validator failed for field "Device.certificate": %w`, err)}
		}
	}
	return nil
}

func (du *DeviceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := du.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(device.Table, device.Columns, sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt))
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Certificate(); ok {
		_spec.SetField(device.FieldCertificate, field.TypeString, value)
	}
	if du.mutation.ConfigsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.ConfigsTable,
			Columns: []string{device.ConfigsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostconfig.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedConfigsIDs(); len(nodes) > 0 && !du.mutation.ConfigsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.ConfigsTable,
			Columns: []string{device.ConfigsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostconfig.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.ConfigsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.ConfigsTable,
			Columns: []string{device.ConfigsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostconfig.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{device.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DeviceUpdateOne is the builder for updating a single Device entity.
type DeviceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeviceMutation
}

// SetCertificate sets the "certificate" field.
func (duo *DeviceUpdateOne) SetCertificate(s string) *DeviceUpdateOne {
	duo.mutation.SetCertificate(s)
	return duo
}

// SetNillableCertificate sets the "certificate" field if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableCertificate(s *string) *DeviceUpdateOne {
	if s != nil {
		duo.SetCertificate(*s)
	}
	return duo
}

// AddConfigIDs adds the "configs" edge to the HostConfig entity by IDs.
func (duo *DeviceUpdateOne) AddConfigIDs(ids ...int) *DeviceUpdateOne {
	duo.mutation.AddConfigIDs(ids...)
	return duo
}

// AddConfigs adds the "configs" edges to the HostConfig entity.
func (duo *DeviceUpdateOne) AddConfigs(h ...*HostConfig) *DeviceUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return duo.AddConfigIDs(ids...)
}

// Mutation returns the DeviceMutation object of the builder.
func (duo *DeviceUpdateOne) Mutation() *DeviceMutation {
	return duo.mutation
}

// ClearConfigs clears all "configs" edges to the HostConfig entity.
func (duo *DeviceUpdateOne) ClearConfigs() *DeviceUpdateOne {
	duo.mutation.ClearConfigs()
	return duo
}

// RemoveConfigIDs removes the "configs" edge to HostConfig entities by IDs.
func (duo *DeviceUpdateOne) RemoveConfigIDs(ids ...int) *DeviceUpdateOne {
	duo.mutation.RemoveConfigIDs(ids...)
	return duo
}

// RemoveConfigs removes "configs" edges to HostConfig entities.
func (duo *DeviceUpdateOne) RemoveConfigs(h ...*HostConfig) *DeviceUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return duo.RemoveConfigIDs(ids...)
}

// Where appends a list predicates to the DeviceUpdate builder.
func (duo *DeviceUpdateOne) Where(ps ...predicate.Device) *DeviceUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DeviceUpdateOne) Select(field string, fields ...string) *DeviceUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Device entity.
func (duo *DeviceUpdateOne) Save(ctx context.Context) (*Device, error) {
	return withHooks(ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DeviceUpdateOne) SaveX(ctx context.Context) *Device {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DeviceUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DeviceUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DeviceUpdateOne) check() error {
	if v, ok := duo.mutation.Certificate(); ok {
		if err := device.CertificateValidator(v); err != nil {
			return &ValidationError{Name: "certificate", err: fmt.Errorf(`ent: validator failed for field "Device.certificate": %w`, err)}
		}
	}
	return nil
}

func (duo *DeviceUpdateOne) sqlSave(ctx context.Context) (_node *Device, err error) {
	if err := duo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(device.Table, device.Columns, sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt))
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Device.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, device.FieldID)
		for _, f := range fields {
			if !device.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != device.FieldID {
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
	if value, ok := duo.mutation.Certificate(); ok {
		_spec.SetField(device.FieldCertificate, field.TypeString, value)
	}
	if duo.mutation.ConfigsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.ConfigsTable,
			Columns: []string{device.ConfigsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostconfig.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedConfigsIDs(); len(nodes) > 0 && !duo.mutation.ConfigsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.ConfigsTable,
			Columns: []string{device.ConfigsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostconfig.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.ConfigsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.ConfigsTable,
			Columns: []string{device.ConfigsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostconfig.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Device{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{device.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}
