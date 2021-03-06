// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/suisrc/zgo/demo/model/ent/demo"
)

// Demo is the model entity for the Demo schema.
type Demo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Memo holds the value of the "memo" field.
	Memo string `json:"memo,omitempty"`
	// Status holds the value of the "status" field.
	Status int `json:"status,omitempty"`
	// Creator holds the value of the "creator" field.
	Creator string `json:"creator,omitempty"`
	// Updator holds the value of the "updator" field.
	Updator string `json:"updator,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DemoQuery when eager-loading is set.
	Edges         DemoEdges `json:"edges"`
	demo_children *int
}

// DemoEdges holds the relations/edges for other nodes in the graph.
type DemoEdges struct {
	// Parent holds the value of the parent edge.
	Parent *Demo
	// Children holds the value of the children edge.
	Children []*Demo
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DemoEdges) ParentOrErr() (*Demo, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: demo.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e DemoEdges) ChildrenOrErr() ([]*Demo, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Demo) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // code
		&sql.NullString{}, // name
		&sql.NullString{}, // memo
		&sql.NullInt64{},  // status
		&sql.NullString{}, // creator
		&sql.NullString{}, // updator
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Demo) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // demo_children
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Demo fields.
func (d *Demo) assignValues(values ...interface{}) error {
	if m, n := len(values), len(demo.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	d.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field code", values[0])
	} else if value.Valid {
		d.Code = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[1])
	} else if value.Valid {
		d.Name = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field memo", values[2])
	} else if value.Valid {
		d.Memo = value.String
	}
	if value, ok := values[3].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field status", values[3])
	} else if value.Valid {
		d.Status = int(value.Int64)
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field creator", values[4])
	} else if value.Valid {
		d.Creator = value.String
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field updator", values[5])
	} else if value.Valid {
		d.Updator = value.String
	}
	if value, ok := values[6].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[6])
	} else if value.Valid {
		d.CreatedAt = value.Time
	}
	if value, ok := values[7].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[7])
	} else if value.Valid {
		d.UpdatedAt = value.Time
	}
	values = values[8:]
	if len(values) == len(demo.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field demo_children", value)
		} else if value.Valid {
			d.demo_children = new(int)
			*d.demo_children = int(value.Int64)
		}
	}
	return nil
}

// QueryParent queries the parent edge of the Demo.
func (d *Demo) QueryParent() *DemoQuery {
	return (&DemoClient{config: d.config}).QueryParent(d)
}

// QueryChildren queries the children edge of the Demo.
func (d *Demo) QueryChildren() *DemoQuery {
	return (&DemoClient{config: d.config}).QueryChildren(d)
}

// Update returns a builder for updating this Demo.
// Note that, you need to call Demo.Unwrap() before calling this method, if this Demo
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Demo) Update() *DemoUpdateOne {
	return (&DemoClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (d *Demo) Unwrap() *Demo {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Demo is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Demo) String() string {
	var builder strings.Builder
	builder.WriteString("Demo(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", code=")
	builder.WriteString(d.Code)
	builder.WriteString(", name=")
	builder.WriteString(d.Name)
	builder.WriteString(", memo=")
	builder.WriteString(d.Memo)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", d.Status))
	builder.WriteString(", creator=")
	builder.WriteString(d.Creator)
	builder.WriteString(", updator=")
	builder.WriteString(d.Updator)
	builder.WriteString(", created_at=")
	builder.WriteString(d.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(d.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Demos is a parsable slice of Demo.
type Demos []*Demo

func (d Demos) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
