// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Peanuttown/dd_contacts/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Depts holds the value of the depts edge.
	Depts []*Dept `json:"depts,omitempty"`
	// PropertiesInDept holds the value of the properties_in_dept edge.
	PropertiesInDept []*UserPropertyInDept `json:"properties_in_dept,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// DeptsOrErr returns the Depts value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) DeptsOrErr() ([]*Dept, error) {
	if e.loadedTypes[0] {
		return e.Depts, nil
	}
	return nil, &NotLoadedError{edge: "depts"}
}

// PropertiesInDeptOrErr returns the PropertiesInDept value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PropertiesInDeptOrErr() ([]*UserPropertyInDept, error) {
	if e.loadedTypes[1] {
		return e.PropertiesInDept, nil
	}
	return nil, &NotLoadedError{edge: "properties_in_dept"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID, user.FieldName, user.FieldPhone:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				u.ID = value.String
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				u.Phone = value.String
			}
		}
	}
	return nil
}

// QueryDepts queries the "depts" edge of the User entity.
func (u *User) QueryDepts() *DeptQuery {
	return (&UserClient{config: u.config}).QueryDepts(u)
}

// QueryPropertiesInDept queries the "properties_in_dept" edge of the User entity.
func (u *User) QueryPropertiesInDept() *UserPropertyInDeptQuery {
	return (&UserClient{config: u.config}).QueryPropertiesInDept(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteString(", phone=")
	builder.WriteString(u.Phone)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
