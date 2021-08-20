package sqlmw

import (
	"context"
	"database/sql/driver"
)

type fakeDriver struct {
	conn driver.Conn
}

func (d *fakeDriver) Open(_ string) (driver.Conn, error) {
	return d.conn, nil
}

type fakeStmt struct {
	called bool // nolint:structcheck // ignore unused warning, it is accessed via reflection
}

type fakeStmtWithCheckNamedValue struct {
	fakeStmt
}

type fakeStmtWithoutCheckNamedValue struct {
	fakeStmt
}

type fakeStmtWithValStore struct {
	fakeStmt
	val []driver.Value
}

func (s *fakeStmtWithValStore) Query(v []driver.Value) (driver.Rows, error) {
	s.val = v
	return nil, nil
}

func (s fakeStmt) Close() error {
	return nil
}

func (s fakeStmt) NumInput() int {
	return 1
}

func (s fakeStmt) Exec(_ []driver.Value) (driver.Result, error) {
	return nil, nil
}

func (s fakeStmt) Query(_ []driver.Value) (driver.Rows, error) {
	return nil, nil
}

func (s *fakeStmtWithCheckNamedValue) CheckNamedValue(_ *driver.NamedValue) (err error) {
	s.called = true
	return
}

type fakeRows struct {
	con         *fakeConn
	closeCalled bool // nolint:structcheck,unused // ignore unused warning, it is accessed via reflection
}

func (r *fakeRows) Close() error {
	r.con.rowsCloseCalled = true
	return nil
}

func (r *fakeRows) Columns() []string {
	return nil
}

func (r *fakeRows) Next(_ []driver.Value) error {
	return nil
}

type fakeConn struct {
	called          bool // nolint:structcheck // ignore unused warning, it is accessed via reflection
	rowsCloseCalled bool
	stmt            driver.Stmt
}

type fakeConnWithCheckNamedValue struct {
	fakeConn
}

type fakeConnWithoutCheckNamedValue struct {
	fakeConn
}

func (c *fakeConn) Prepare(_ string) (driver.Stmt, error) {
	return nil, nil
}

func (c *fakeConn) PrepareContext(_ context.Context, _ string) (driver.Stmt, error) {
	return c.stmt, nil
}

func (c *fakeConn) Close() error { return nil }

func (c *fakeConn) Begin() (driver.Tx, error) { return nil, nil }

func (c *fakeConn) QueryContext(_ context.Context, _ string, _ []driver.NamedValue) (driver.Rows, error) {
	return &fakeRows{con: c}, nil
}

func (c *fakeConnWithCheckNamedValue) CheckNamedValue(_ *driver.NamedValue) (err error) {
	c.called = true
	return
}

type fakeInterceptor struct {
	NullInterceptor
}
