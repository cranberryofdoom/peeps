package controllers

import (
	"github.com/albrow/zoom"
	"github.com/revel/revel"
	"peeps/app/models"
)

type Persons struct {
	*revel.Controller
	JsonController
}

func (c Persons) Index() revel.Result {
	persons := make([]*models.Person, 0)
	if err := zoom.NewQuery("Person").Scan(&persons); err != nil {
		return c.RenderJsonError(500, err)
	}
	return c.RenderJson(persons)
}

func (c Persons) Create(name string, age int) revel.Result {
	p := &models.Person{Name: name, Age: age}
	if err := zoom.Save(p); err != nil {
		return c.RenderJsonError(500, err)
	}
	return c.RenderJson(p)
}

func (c Persons) Show(id string) revel.Result {
	p, err := zoom.FindById("Person", id)
	if err != nil {
		return c.RenderJsonError(500, err)
	}
	return c.RenderJson(p)
}

func (c Persons) Update(name string, age int, id string) revel.Result {
	p := &models.Person{}
	if err := zoom.ScanById(id, p); err != nil {
		return c.RenderJsonError(500, err)
	}
	if _, ok := c.Params.Values["name"]; ok {
		p.Name = name
	}
	if _, ok := c.Params.Values["age"]; ok {
		p.Age = age
	}
	if err := zoom.Save(p); err != nil {
		return c.RenderJsonError(500, err)
	}
	return c.RenderJson(p)
}

func (c Persons) Delete(id string) revel.Result {
	if err := zoom.DeleteById("Person", id); err != nil {
		return c.RenderJsonError(500, err)
	}
	return c.RenderJsonOk()
}
