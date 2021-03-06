// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/ACER/app/ent/course"
	"github.com/ACER/app/ent/courseitem"
	"github.com/ACER/app/ent/predicate"
	"github.com/ACER/app/ent/subject"
	"github.com/ACER/app/ent/subjecttype"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// CourseItemUpdate is the builder for updating CourseItem entities.
type CourseItemUpdate struct {
	config
	hooks      []Hook
	mutation   *CourseItemMutation
	predicates []predicate.CourseItem
}

// Where adds a new predicate for the builder.
func (ciu *CourseItemUpdate) Where(ps ...predicate.CourseItem) *CourseItemUpdate {
	ciu.predicates = append(ciu.predicates, ps...)
	return ciu
}

// SetCoursesID sets the courses edge to Course by id.
func (ciu *CourseItemUpdate) SetCoursesID(id int) *CourseItemUpdate {
	ciu.mutation.SetCoursesID(id)
	return ciu
}

// SetNillableCoursesID sets the courses edge to Course by id if the given value is not nil.
func (ciu *CourseItemUpdate) SetNillableCoursesID(id *int) *CourseItemUpdate {
	if id != nil {
		ciu = ciu.SetCoursesID(*id)
	}
	return ciu
}

// SetCourses sets the courses edge to Course.
func (ciu *CourseItemUpdate) SetCourses(c *Course) *CourseItemUpdate {
	return ciu.SetCoursesID(c.ID)
}

// SetSubjectsID sets the subjects edge to Subject by id.
func (ciu *CourseItemUpdate) SetSubjectsID(id int) *CourseItemUpdate {
	ciu.mutation.SetSubjectsID(id)
	return ciu
}

// SetNillableSubjectsID sets the subjects edge to Subject by id if the given value is not nil.
func (ciu *CourseItemUpdate) SetNillableSubjectsID(id *int) *CourseItemUpdate {
	if id != nil {
		ciu = ciu.SetSubjectsID(*id)
	}
	return ciu
}

// SetSubjects sets the subjects edge to Subject.
func (ciu *CourseItemUpdate) SetSubjects(s *Subject) *CourseItemUpdate {
	return ciu.SetSubjectsID(s.ID)
}

// SetTypesID sets the types edge to SubjectType by id.
func (ciu *CourseItemUpdate) SetTypesID(id int) *CourseItemUpdate {
	ciu.mutation.SetTypesID(id)
	return ciu
}

// SetNillableTypesID sets the types edge to SubjectType by id if the given value is not nil.
func (ciu *CourseItemUpdate) SetNillableTypesID(id *int) *CourseItemUpdate {
	if id != nil {
		ciu = ciu.SetTypesID(*id)
	}
	return ciu
}

// SetTypes sets the types edge to SubjectType.
func (ciu *CourseItemUpdate) SetTypes(s *SubjectType) *CourseItemUpdate {
	return ciu.SetTypesID(s.ID)
}

// Mutation returns the CourseItemMutation object of the builder.
func (ciu *CourseItemUpdate) Mutation() *CourseItemMutation {
	return ciu.mutation
}

// ClearCourses clears the courses edge to Course.
func (ciu *CourseItemUpdate) ClearCourses() *CourseItemUpdate {
	ciu.mutation.ClearCourses()
	return ciu
}

// ClearSubjects clears the subjects edge to Subject.
func (ciu *CourseItemUpdate) ClearSubjects() *CourseItemUpdate {
	ciu.mutation.ClearSubjects()
	return ciu
}

// ClearTypes clears the types edge to SubjectType.
func (ciu *CourseItemUpdate) ClearTypes() *CourseItemUpdate {
	ciu.mutation.ClearTypes()
	return ciu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ciu *CourseItemUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(ciu.hooks) == 0 {
		affected, err = ciu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CourseItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ciu.mutation = mutation
			affected, err = ciu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ciu.hooks) - 1; i >= 0; i-- {
			mut = ciu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ciu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ciu *CourseItemUpdate) SaveX(ctx context.Context) int {
	affected, err := ciu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ciu *CourseItemUpdate) Exec(ctx context.Context) error {
	_, err := ciu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciu *CourseItemUpdate) ExecX(ctx context.Context) {
	if err := ciu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ciu *CourseItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   courseitem.Table,
			Columns: courseitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: courseitem.FieldID,
			},
		},
	}
	if ps := ciu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ciu.mutation.CoursesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciu.mutation.CoursesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ciu.mutation.SubjectsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciu.mutation.SubjectsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ciu.mutation.TypesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciu.mutation.TypesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ciu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{courseitem.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CourseItemUpdateOne is the builder for updating a single CourseItem entity.
type CourseItemUpdateOne struct {
	config
	hooks    []Hook
	mutation *CourseItemMutation
}

// SetCoursesID sets the courses edge to Course by id.
func (ciuo *CourseItemUpdateOne) SetCoursesID(id int) *CourseItemUpdateOne {
	ciuo.mutation.SetCoursesID(id)
	return ciuo
}

// SetNillableCoursesID sets the courses edge to Course by id if the given value is not nil.
func (ciuo *CourseItemUpdateOne) SetNillableCoursesID(id *int) *CourseItemUpdateOne {
	if id != nil {
		ciuo = ciuo.SetCoursesID(*id)
	}
	return ciuo
}

// SetCourses sets the courses edge to Course.
func (ciuo *CourseItemUpdateOne) SetCourses(c *Course) *CourseItemUpdateOne {
	return ciuo.SetCoursesID(c.ID)
}

// SetSubjectsID sets the subjects edge to Subject by id.
func (ciuo *CourseItemUpdateOne) SetSubjectsID(id int) *CourseItemUpdateOne {
	ciuo.mutation.SetSubjectsID(id)
	return ciuo
}

// SetNillableSubjectsID sets the subjects edge to Subject by id if the given value is not nil.
func (ciuo *CourseItemUpdateOne) SetNillableSubjectsID(id *int) *CourseItemUpdateOne {
	if id != nil {
		ciuo = ciuo.SetSubjectsID(*id)
	}
	return ciuo
}

// SetSubjects sets the subjects edge to Subject.
func (ciuo *CourseItemUpdateOne) SetSubjects(s *Subject) *CourseItemUpdateOne {
	return ciuo.SetSubjectsID(s.ID)
}

// SetTypesID sets the types edge to SubjectType by id.
func (ciuo *CourseItemUpdateOne) SetTypesID(id int) *CourseItemUpdateOne {
	ciuo.mutation.SetTypesID(id)
	return ciuo
}

// SetNillableTypesID sets the types edge to SubjectType by id if the given value is not nil.
func (ciuo *CourseItemUpdateOne) SetNillableTypesID(id *int) *CourseItemUpdateOne {
	if id != nil {
		ciuo = ciuo.SetTypesID(*id)
	}
	return ciuo
}

// SetTypes sets the types edge to SubjectType.
func (ciuo *CourseItemUpdateOne) SetTypes(s *SubjectType) *CourseItemUpdateOne {
	return ciuo.SetTypesID(s.ID)
}

// Mutation returns the CourseItemMutation object of the builder.
func (ciuo *CourseItemUpdateOne) Mutation() *CourseItemMutation {
	return ciuo.mutation
}

// ClearCourses clears the courses edge to Course.
func (ciuo *CourseItemUpdateOne) ClearCourses() *CourseItemUpdateOne {
	ciuo.mutation.ClearCourses()
	return ciuo
}

// ClearSubjects clears the subjects edge to Subject.
func (ciuo *CourseItemUpdateOne) ClearSubjects() *CourseItemUpdateOne {
	ciuo.mutation.ClearSubjects()
	return ciuo
}

// ClearTypes clears the types edge to SubjectType.
func (ciuo *CourseItemUpdateOne) ClearTypes() *CourseItemUpdateOne {
	ciuo.mutation.ClearTypes()
	return ciuo
}

// Save executes the query and returns the updated entity.
func (ciuo *CourseItemUpdateOne) Save(ctx context.Context) (*CourseItem, error) {

	var (
		err  error
		node *CourseItem
	)
	if len(ciuo.hooks) == 0 {
		node, err = ciuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CourseItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ciuo.mutation = mutation
			node, err = ciuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ciuo.hooks) - 1; i >= 0; i-- {
			mut = ciuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ciuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ciuo *CourseItemUpdateOne) SaveX(ctx context.Context) *CourseItem {
	ci, err := ciuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return ci
}

// Exec executes the query on the entity.
func (ciuo *CourseItemUpdateOne) Exec(ctx context.Context) error {
	_, err := ciuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciuo *CourseItemUpdateOne) ExecX(ctx context.Context) {
	if err := ciuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ciuo *CourseItemUpdateOne) sqlSave(ctx context.Context) (ci *CourseItem, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   courseitem.Table,
			Columns: courseitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: courseitem.FieldID,
			},
		},
	}
	id, ok := ciuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing CourseItem.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ciuo.mutation.CoursesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciuo.mutation.CoursesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ciuo.mutation.SubjectsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciuo.mutation.SubjectsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ciuo.mutation.TypesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciuo.mutation.TypesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	ci = &CourseItem{config: ciuo.config}
	_spec.Assign = ci.assignValues
	_spec.ScanValues = ci.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ciuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{courseitem.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return ci, nil
}
