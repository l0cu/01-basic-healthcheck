package mooc

import "context"

// CourseRepository defines the expected behavior from a course storage.
type CourseRepository interface {
	Save(ctx context.Context, course Course) error
}

// Course is the data structure that represents a course
type Course struct {
	id       string
	name     string
	duration string
}

// NewCourse creates a new course
func NewCourse(id, name, duration string) Course {
	return Course{
		id:       id,
		name:     name,
		duration: duration,
	}
}

// ID returns the course unique identifier.
func (c Course) ID() string {
	return c.id
}

// Name returns the course name.
func (c Course) Name() string {
	return c.name
}

// Duration returns the duration name.
func (c Course) Duration() string {
	return c.duration
}
