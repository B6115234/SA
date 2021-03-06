// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/ACER/app/ent/course"
	"github.com/ACER/app/ent/courseitem"
	"github.com/ACER/app/ent/subject"
	"github.com/ACER/app/ent/subjecttype"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// CourseItemCreate is the builder for creating a CourseItem entity.
type CourseItemCreate struct {
	config
	mutation *CourseItemMutation
	hooks    []Hook
}

// SetCoursesID sets the courses edge to Course by id.
func (cic *CourseItemCreate) SetCoursesID(id int) *CourseItemCreate {
	cic.mutation.SetCoursesID(id)
	return cic
}

// SetNillableCoursesID sets the courses edge to Course by id if the given value is not nil.
func (cic *CourseItemCreate) SetNillableCoursesID(id *int) *CourseItemCreate {
	if id != nil {
		cic = cic.SetCoursesID(*id)
	}
	return cic
}

// SetCourses sets the courses edge to Course.
func (cic *CourseItemCreate) SetCourses(c *Course) *CourseItemCreate {
	return cic.SetCoursesID(c.ID)
}

// SetSubjectsID sets the subjects edge to Subject by id.
func (cic *CourseItemCreate) SetSubjectsID(id int) *CourseItemCreate {
	cic.mutation.SetSubjectsID(id)
	return cic
}

// SetNillableSubjectsID sets the subjects edge to Subject by id if the given value is not nil.
func (cic *CourseItemCreate) SetNillableSubjectsID(id *int) *CourseItemCreate {
	if id != nil {
		cic = cic.SetSubjectsID(*id)
	}
	return cic
}

// SetSubjects sets the subjects edge to Subject.
func (cic *CourseItemCreate) SetSubjects(s *Subject) *CourseItemCreate {
	return cic.SetSubjectsID(s.ID)
}

// SetTypesID sets the types edge to SubjectType by id.
func (cic *CourseItemCreate) SetTypesID(id int) *CourseItemCreate {
	cic.mutation.SetTypesID(id)
	return cic
}

// SetNillableTypesID sets the types edge to SubjectType by id if the given value is not nil.
func (cic *CourseItemCreate) SetNillableTypesID(id *int) *CourseItemCreate {
	if id != nil {
		cic = cic.SetTypesID(*id)
	}
	return cic
}

// SetTypes sets the types edge to SubjectType.
func (cic *CourseItemCreate) SetTypes(s *SubjectType) *CourseItemCreate {
	return cic.SetTypesID(s.ID)
}

// Mutation returns the CourseItemMutation object of the builder.
func (cic *CourseItemCreate) Mutation() *CourseItemMutation {
	return cic.mutation
}

// Save creates the CourseItem in the database.
func (cic *CourseItemCreate) Save(ctx context.Context) (*CourseItem, error) {
	var (
		err  error
		node *CourseItem
	)
	if len(cic.hooks) == 0 {
		node, err = cic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CourseItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cic.mutation = mutation
			node, err = cic.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cic.hooks) - 1; i >= 0; i-- {
			mut = cic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cic *CourseItemCreate) SaveX(ctx context.Context) *CourseItem {
	v, err := cic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cic *CourseItemCreate) sqlSave(ctx context.Context) (*CourseItem, error) {
	ci, _spec := cic.createSpec()
	if err := sqlgraph.CreateNode(ctx, cic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	ci.ID = int(id)
	return ci, nil
}

func (cic *CourseItemCreate) createSpec() (*CourseItem, *sqlgraph.CreateSpec) {
	var (
		ci    = &CourseItem{config: cic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: courseitem.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: courseitem.FieldID,
			},
		}
	)
	if nodes := cic.mutation.CoursesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   courseitem.CoursesTable,
			Columns: []string{courseitem.CoursesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: course.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cic.mutation.SubjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   courseitem.SubjectsTable,
			Columns: []string{courseitem.SubjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subject.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cic.mutation.TypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   courseitem.TypesTable,
			Columns: []string{courseitem.TypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subjecttype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return ci, _spec
}
