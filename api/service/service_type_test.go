package service_test

import (
	. "github.com/timeredbull/tsuru/api/service"
	. "launchpad.net/gocheck"
)

func (s *ServiceSuite) createServiceType() {
	s.serviceType = &ServiceType{Name: "Mysql", Charm: "mysql"}
	s.serviceType.Create()
}

func (s *ServiceSuite) TestCreateServiceType(c *C) {
	s.createServiceType()
	rows, err := s.db.Query("SELECT name, charm FROM service_type WHERE name = 'Mysql' AND charm='mysql'")
	c.Check(err, IsNil)

	var name string
	var charm string
	for rows.Next() {
		rows.Scan(&name, &charm)
	}

	c.Assert(name, Equals, "Mysql")
	c.Assert(charm, Equals, "mysql")
}

func (s *ServiceSuite) TestDeleteServiceType(c *C) {
	s.createServiceType()
	s.serviceType.Delete()

	rows, err := s.db.Query("SELECT count(*) FROM service_type WHERE name = 'Mysql' AND charm = 'mysql'")
	c.Assert(err, IsNil)

	var qtd int
	for rows.Next() {
		rows.Scan(&qtd)
	}

	c.Assert(qtd, Equals, 0)
}
