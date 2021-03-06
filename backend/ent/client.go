// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/ACER/app/ent/migrate"

	"github.com/ACER/app/ent/course"
	"github.com/ACER/app/ent/courseitem"
	"github.com/ACER/app/ent/subject"
	"github.com/ACER/app/ent/subjecttype"
	"github.com/ACER/app/ent/teacher"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Course is the client for interacting with the Course builders.
	Course *CourseClient
	// CourseItem is the client for interacting with the CourseItem builders.
	CourseItem *CourseItemClient
	// Subject is the client for interacting with the Subject builders.
	Subject *SubjectClient
	// SubjectType is the client for interacting with the SubjectType builders.
	SubjectType *SubjectTypeClient
	// Teacher is the client for interacting with the Teacher builders.
	Teacher *TeacherClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Course = NewCourseClient(c.config)
	c.CourseItem = NewCourseItemClient(c.config)
	c.Subject = NewSubjectClient(c.config)
	c.SubjectType = NewSubjectTypeClient(c.config)
	c.Teacher = NewTeacherClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Course:      NewCourseClient(cfg),
		CourseItem:  NewCourseItemClient(cfg),
		Subject:     NewSubjectClient(cfg),
		SubjectType: NewSubjectTypeClient(cfg),
		Teacher:     NewTeacherClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:      cfg,
		Course:      NewCourseClient(cfg),
		CourseItem:  NewCourseItemClient(cfg),
		Subject:     NewSubjectClient(cfg),
		SubjectType: NewSubjectTypeClient(cfg),
		Teacher:     NewTeacherClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Course.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Course.Use(hooks...)
	c.CourseItem.Use(hooks...)
	c.Subject.Use(hooks...)
	c.SubjectType.Use(hooks...)
	c.Teacher.Use(hooks...)
}

// CourseClient is a client for the Course schema.
type CourseClient struct {
	config
}

// NewCourseClient returns a client for the Course from the given config.
func NewCourseClient(c config) *CourseClient {
	return &CourseClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `course.Hooks(f(g(h())))`.
func (c *CourseClient) Use(hooks ...Hook) {
	c.hooks.Course = append(c.hooks.Course, hooks...)
}

// Create returns a create builder for Course.
func (c *CourseClient) Create() *CourseCreate {
	mutation := newCourseMutation(c.config, OpCreate)
	return &CourseCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Course.
func (c *CourseClient) Update() *CourseUpdate {
	mutation := newCourseMutation(c.config, OpUpdate)
	return &CourseUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CourseClient) UpdateOne(co *Course) *CourseUpdateOne {
	mutation := newCourseMutation(c.config, OpUpdateOne, withCourse(co))
	return &CourseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CourseClient) UpdateOneID(id int) *CourseUpdateOne {
	mutation := newCourseMutation(c.config, OpUpdateOne, withCourseID(id))
	return &CourseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Course.
func (c *CourseClient) Delete() *CourseDelete {
	mutation := newCourseMutation(c.config, OpDelete)
	return &CourseDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CourseClient) DeleteOne(co *Course) *CourseDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CourseClient) DeleteOneID(id int) *CourseDeleteOne {
	builder := c.Delete().Where(course.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CourseDeleteOne{builder}
}

// Create returns a query builder for Course.
func (c *CourseClient) Query() *CourseQuery {
	return &CourseQuery{config: c.config}
}

// Get returns a Course entity by its id.
func (c *CourseClient) Get(ctx context.Context, id int) (*Course, error) {
	return c.Query().Where(course.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CourseClient) GetX(ctx context.Context, id int) *Course {
	co, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return co
}

// QueryOwner queries the owner edge of a Course.
func (c *CourseClient) QueryOwner(co *Course) *TeacherQuery {
	query := &TeacherQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(course.Table, course.FieldID, id),
			sqlgraph.To(teacher.Table, teacher.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, course.OwnerTable, course.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCourseItems queries the course_items edge of a Course.
func (c *CourseClient) QueryCourseItems(co *Course) *CourseItemQuery {
	query := &CourseItemQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(course.Table, course.FieldID, id),
			sqlgraph.To(courseitem.Table, courseitem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, course.CourseItemsTable, course.CourseItemsColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CourseClient) Hooks() []Hook {
	return c.hooks.Course
}

// CourseItemClient is a client for the CourseItem schema.
type CourseItemClient struct {
	config
}

// NewCourseItemClient returns a client for the CourseItem from the given config.
func NewCourseItemClient(c config) *CourseItemClient {
	return &CourseItemClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `courseitem.Hooks(f(g(h())))`.
func (c *CourseItemClient) Use(hooks ...Hook) {
	c.hooks.CourseItem = append(c.hooks.CourseItem, hooks...)
}

// Create returns a create builder for CourseItem.
func (c *CourseItemClient) Create() *CourseItemCreate {
	mutation := newCourseItemMutation(c.config, OpCreate)
	return &CourseItemCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for CourseItem.
func (c *CourseItemClient) Update() *CourseItemUpdate {
	mutation := newCourseItemMutation(c.config, OpUpdate)
	return &CourseItemUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CourseItemClient) UpdateOne(ci *CourseItem) *CourseItemUpdateOne {
	mutation := newCourseItemMutation(c.config, OpUpdateOne, withCourseItem(ci))
	return &CourseItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CourseItemClient) UpdateOneID(id int) *CourseItemUpdateOne {
	mutation := newCourseItemMutation(c.config, OpUpdateOne, withCourseItemID(id))
	return &CourseItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CourseItem.
func (c *CourseItemClient) Delete() *CourseItemDelete {
	mutation := newCourseItemMutation(c.config, OpDelete)
	return &CourseItemDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CourseItemClient) DeleteOne(ci *CourseItem) *CourseItemDeleteOne {
	return c.DeleteOneID(ci.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CourseItemClient) DeleteOneID(id int) *CourseItemDeleteOne {
	builder := c.Delete().Where(courseitem.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CourseItemDeleteOne{builder}
}

// Create returns a query builder for CourseItem.
func (c *CourseItemClient) Query() *CourseItemQuery {
	return &CourseItemQuery{config: c.config}
}

// Get returns a CourseItem entity by its id.
func (c *CourseItemClient) Get(ctx context.Context, id int) (*CourseItem, error) {
	return c.Query().Where(courseitem.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CourseItemClient) GetX(ctx context.Context, id int) *CourseItem {
	ci, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return ci
}

// QueryCourses queries the courses edge of a CourseItem.
func (c *CourseItemClient) QueryCourses(ci *CourseItem) *CourseQuery {
	query := &CourseQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ci.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(courseitem.Table, courseitem.FieldID, id),
			sqlgraph.To(course.Table, course.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, courseitem.CoursesTable, courseitem.CoursesColumn),
		)
		fromV = sqlgraph.Neighbors(ci.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySubjects queries the subjects edge of a CourseItem.
func (c *CourseItemClient) QuerySubjects(ci *CourseItem) *SubjectQuery {
	query := &SubjectQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ci.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(courseitem.Table, courseitem.FieldID, id),
			sqlgraph.To(subject.Table, subject.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, courseitem.SubjectsTable, courseitem.SubjectsColumn),
		)
		fromV = sqlgraph.Neighbors(ci.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTypes queries the types edge of a CourseItem.
func (c *CourseItemClient) QueryTypes(ci *CourseItem) *SubjectTypeQuery {
	query := &SubjectTypeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ci.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(courseitem.Table, courseitem.FieldID, id),
			sqlgraph.To(subjecttype.Table, subjecttype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, courseitem.TypesTable, courseitem.TypesColumn),
		)
		fromV = sqlgraph.Neighbors(ci.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CourseItemClient) Hooks() []Hook {
	return c.hooks.CourseItem
}

// SubjectClient is a client for the Subject schema.
type SubjectClient struct {
	config
}

// NewSubjectClient returns a client for the Subject from the given config.
func NewSubjectClient(c config) *SubjectClient {
	return &SubjectClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `subject.Hooks(f(g(h())))`.
func (c *SubjectClient) Use(hooks ...Hook) {
	c.hooks.Subject = append(c.hooks.Subject, hooks...)
}

// Create returns a create builder for Subject.
func (c *SubjectClient) Create() *SubjectCreate {
	mutation := newSubjectMutation(c.config, OpCreate)
	return &SubjectCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Subject.
func (c *SubjectClient) Update() *SubjectUpdate {
	mutation := newSubjectMutation(c.config, OpUpdate)
	return &SubjectUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SubjectClient) UpdateOne(s *Subject) *SubjectUpdateOne {
	mutation := newSubjectMutation(c.config, OpUpdateOne, withSubject(s))
	return &SubjectUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SubjectClient) UpdateOneID(id int) *SubjectUpdateOne {
	mutation := newSubjectMutation(c.config, OpUpdateOne, withSubjectID(id))
	return &SubjectUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Subject.
func (c *SubjectClient) Delete() *SubjectDelete {
	mutation := newSubjectMutation(c.config, OpDelete)
	return &SubjectDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SubjectClient) DeleteOne(s *Subject) *SubjectDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SubjectClient) DeleteOneID(id int) *SubjectDeleteOne {
	builder := c.Delete().Where(subject.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SubjectDeleteOne{builder}
}

// Create returns a query builder for Subject.
func (c *SubjectClient) Query() *SubjectQuery {
	return &SubjectQuery{config: c.config}
}

// Get returns a Subject entity by its id.
func (c *SubjectClient) Get(ctx context.Context, id int) (*Subject, error) {
	return c.Query().Where(subject.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SubjectClient) GetX(ctx context.Context, id int) *Subject {
	s, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return s
}

// QueryOwner queries the owner edge of a Subject.
func (c *SubjectClient) QueryOwner(s *Subject) *TeacherQuery {
	query := &TeacherQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(subject.Table, subject.FieldID, id),
			sqlgraph.To(teacher.Table, teacher.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, subject.OwnerTable, subject.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCourseItems queries the course_items edge of a Subject.
func (c *SubjectClient) QueryCourseItems(s *Subject) *CourseItemQuery {
	query := &CourseItemQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(subject.Table, subject.FieldID, id),
			sqlgraph.To(courseitem.Table, courseitem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, subject.CourseItemsTable, subject.CourseItemsColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SubjectClient) Hooks() []Hook {
	return c.hooks.Subject
}

// SubjectTypeClient is a client for the SubjectType schema.
type SubjectTypeClient struct {
	config
}

// NewSubjectTypeClient returns a client for the SubjectType from the given config.
func NewSubjectTypeClient(c config) *SubjectTypeClient {
	return &SubjectTypeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `subjecttype.Hooks(f(g(h())))`.
func (c *SubjectTypeClient) Use(hooks ...Hook) {
	c.hooks.SubjectType = append(c.hooks.SubjectType, hooks...)
}

// Create returns a create builder for SubjectType.
func (c *SubjectTypeClient) Create() *SubjectTypeCreate {
	mutation := newSubjectTypeMutation(c.config, OpCreate)
	return &SubjectTypeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for SubjectType.
func (c *SubjectTypeClient) Update() *SubjectTypeUpdate {
	mutation := newSubjectTypeMutation(c.config, OpUpdate)
	return &SubjectTypeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SubjectTypeClient) UpdateOne(st *SubjectType) *SubjectTypeUpdateOne {
	mutation := newSubjectTypeMutation(c.config, OpUpdateOne, withSubjectType(st))
	return &SubjectTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SubjectTypeClient) UpdateOneID(id int) *SubjectTypeUpdateOne {
	mutation := newSubjectTypeMutation(c.config, OpUpdateOne, withSubjectTypeID(id))
	return &SubjectTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SubjectType.
func (c *SubjectTypeClient) Delete() *SubjectTypeDelete {
	mutation := newSubjectTypeMutation(c.config, OpDelete)
	return &SubjectTypeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SubjectTypeClient) DeleteOne(st *SubjectType) *SubjectTypeDeleteOne {
	return c.DeleteOneID(st.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SubjectTypeClient) DeleteOneID(id int) *SubjectTypeDeleteOne {
	builder := c.Delete().Where(subjecttype.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SubjectTypeDeleteOne{builder}
}

// Create returns a query builder for SubjectType.
func (c *SubjectTypeClient) Query() *SubjectTypeQuery {
	return &SubjectTypeQuery{config: c.config}
}

// Get returns a SubjectType entity by its id.
func (c *SubjectTypeClient) Get(ctx context.Context, id int) (*SubjectType, error) {
	return c.Query().Where(subjecttype.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SubjectTypeClient) GetX(ctx context.Context, id int) *SubjectType {
	st, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return st
}

// QueryCourseItems queries the course_items edge of a SubjectType.
func (c *SubjectTypeClient) QueryCourseItems(st *SubjectType) *CourseItemQuery {
	query := &CourseItemQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := st.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(subjecttype.Table, subjecttype.FieldID, id),
			sqlgraph.To(courseitem.Table, courseitem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, subjecttype.CourseItemsTable, subjecttype.CourseItemsColumn),
		)
		fromV = sqlgraph.Neighbors(st.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SubjectTypeClient) Hooks() []Hook {
	return c.hooks.SubjectType
}

// TeacherClient is a client for the Teacher schema.
type TeacherClient struct {
	config
}

// NewTeacherClient returns a client for the Teacher from the given config.
func NewTeacherClient(c config) *TeacherClient {
	return &TeacherClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `teacher.Hooks(f(g(h())))`.
func (c *TeacherClient) Use(hooks ...Hook) {
	c.hooks.Teacher = append(c.hooks.Teacher, hooks...)
}

// Create returns a create builder for Teacher.
func (c *TeacherClient) Create() *TeacherCreate {
	mutation := newTeacherMutation(c.config, OpCreate)
	return &TeacherCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Teacher.
func (c *TeacherClient) Update() *TeacherUpdate {
	mutation := newTeacherMutation(c.config, OpUpdate)
	return &TeacherUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TeacherClient) UpdateOne(t *Teacher) *TeacherUpdateOne {
	mutation := newTeacherMutation(c.config, OpUpdateOne, withTeacher(t))
	return &TeacherUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TeacherClient) UpdateOneID(id int) *TeacherUpdateOne {
	mutation := newTeacherMutation(c.config, OpUpdateOne, withTeacherID(id))
	return &TeacherUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Teacher.
func (c *TeacherClient) Delete() *TeacherDelete {
	mutation := newTeacherMutation(c.config, OpDelete)
	return &TeacherDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TeacherClient) DeleteOne(t *Teacher) *TeacherDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TeacherClient) DeleteOneID(id int) *TeacherDeleteOne {
	builder := c.Delete().Where(teacher.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TeacherDeleteOne{builder}
}

// Create returns a query builder for Teacher.
func (c *TeacherClient) Query() *TeacherQuery {
	return &TeacherQuery{config: c.config}
}

// Get returns a Teacher entity by its id.
func (c *TeacherClient) Get(ctx context.Context, id int) (*Teacher, error) {
	return c.Query().Where(teacher.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TeacherClient) GetX(ctx context.Context, id int) *Teacher {
	t, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return t
}

// QuerySubjects queries the subjects edge of a Teacher.
func (c *TeacherClient) QuerySubjects(t *Teacher) *SubjectQuery {
	query := &SubjectQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(teacher.Table, teacher.FieldID, id),
			sqlgraph.To(subject.Table, subject.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, teacher.SubjectsTable, teacher.SubjectsColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCourses queries the courses edge of a Teacher.
func (c *TeacherClient) QueryCourses(t *Teacher) *CourseQuery {
	query := &CourseQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(teacher.Table, teacher.FieldID, id),
			sqlgraph.To(course.Table, course.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, teacher.CoursesTable, teacher.CoursesColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TeacherClient) Hooks() []Hook {
	return c.hooks.Teacher
}
