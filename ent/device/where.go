// Code generated by ent, DO NOT EDIT.

package device

import (
	"github.com/rahn-it/svalin/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldID, id))
}

// PublicKey applies equality check predicate on the "public_key" field. It's identical to PublicKeyEQ.
func PublicKey(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldPublicKey, v))
}

// Certificate applies equality check predicate on the "certificate" field. It's identical to CertificateEQ.
func Certificate(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldCertificate, v))
}

// PublicKeyEQ applies the EQ predicate on the "public_key" field.
func PublicKeyEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldPublicKey, v))
}

// PublicKeyNEQ applies the NEQ predicate on the "public_key" field.
func PublicKeyNEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldPublicKey, v))
}

// PublicKeyIn applies the In predicate on the "public_key" field.
func PublicKeyIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldPublicKey, vs...))
}

// PublicKeyNotIn applies the NotIn predicate on the "public_key" field.
func PublicKeyNotIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldPublicKey, vs...))
}

// PublicKeyGT applies the GT predicate on the "public_key" field.
func PublicKeyGT(v string) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldPublicKey, v))
}

// PublicKeyGTE applies the GTE predicate on the "public_key" field.
func PublicKeyGTE(v string) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldPublicKey, v))
}

// PublicKeyLT applies the LT predicate on the "public_key" field.
func PublicKeyLT(v string) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldPublicKey, v))
}

// PublicKeyLTE applies the LTE predicate on the "public_key" field.
func PublicKeyLTE(v string) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldPublicKey, v))
}

// PublicKeyContains applies the Contains predicate on the "public_key" field.
func PublicKeyContains(v string) predicate.Device {
	return predicate.Device(sql.FieldContains(FieldPublicKey, v))
}

// PublicKeyHasPrefix applies the HasPrefix predicate on the "public_key" field.
func PublicKeyHasPrefix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasPrefix(FieldPublicKey, v))
}

// PublicKeyHasSuffix applies the HasSuffix predicate on the "public_key" field.
func PublicKeyHasSuffix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasSuffix(FieldPublicKey, v))
}

// PublicKeyEqualFold applies the EqualFold predicate on the "public_key" field.
func PublicKeyEqualFold(v string) predicate.Device {
	return predicate.Device(sql.FieldEqualFold(FieldPublicKey, v))
}

// PublicKeyContainsFold applies the ContainsFold predicate on the "public_key" field.
func PublicKeyContainsFold(v string) predicate.Device {
	return predicate.Device(sql.FieldContainsFold(FieldPublicKey, v))
}

// CertificateEQ applies the EQ predicate on the "certificate" field.
func CertificateEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldCertificate, v))
}

// CertificateNEQ applies the NEQ predicate on the "certificate" field.
func CertificateNEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldCertificate, v))
}

// CertificateIn applies the In predicate on the "certificate" field.
func CertificateIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldCertificate, vs...))
}

// CertificateNotIn applies the NotIn predicate on the "certificate" field.
func CertificateNotIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldCertificate, vs...))
}

// CertificateGT applies the GT predicate on the "certificate" field.
func CertificateGT(v string) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldCertificate, v))
}

// CertificateGTE applies the GTE predicate on the "certificate" field.
func CertificateGTE(v string) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldCertificate, v))
}

// CertificateLT applies the LT predicate on the "certificate" field.
func CertificateLT(v string) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldCertificate, v))
}

// CertificateLTE applies the LTE predicate on the "certificate" field.
func CertificateLTE(v string) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldCertificate, v))
}

// CertificateContains applies the Contains predicate on the "certificate" field.
func CertificateContains(v string) predicate.Device {
	return predicate.Device(sql.FieldContains(FieldCertificate, v))
}

// CertificateHasPrefix applies the HasPrefix predicate on the "certificate" field.
func CertificateHasPrefix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasPrefix(FieldCertificate, v))
}

// CertificateHasSuffix applies the HasSuffix predicate on the "certificate" field.
func CertificateHasSuffix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasSuffix(FieldCertificate, v))
}

// CertificateEqualFold applies the EqualFold predicate on the "certificate" field.
func CertificateEqualFold(v string) predicate.Device {
	return predicate.Device(sql.FieldEqualFold(FieldCertificate, v))
}

// CertificateContainsFold applies the ContainsFold predicate on the "certificate" field.
func CertificateContainsFold(v string) predicate.Device {
	return predicate.Device(sql.FieldContainsFold(FieldCertificate, v))
}

// HasConfigs applies the HasEdge predicate on the "configs" edge.
func HasConfigs() predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ConfigsTable, ConfigsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasConfigsWith applies the HasEdge predicate on the "configs" edge with a given conditions (other predicates).
func HasConfigsWith(preds ...predicate.HostConfig) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := newConfigsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Device) predicate.Device {
	return predicate.Device(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Device) predicate.Device {
	return predicate.Device(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Device) predicate.Device {
	return predicate.Device(sql.NotPredicates(p))
}
