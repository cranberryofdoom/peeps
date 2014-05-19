package controllers

import (
	"fmt"
	"github.com/albrow/zoom"
	"github.com/revel/revel"
	"peeps/app/models"
)

type Persons struct {
	*revel.Controller
}

type JsonError struct {
	err error
}

func (e JsonError) Error() string {
	return fmt.Sprintf("{'error': %s}", e.err)
}

func NewJsonError(e error) JsonError {
	return JsonError{err: e}
}

func (c Persons) Index() revel.Result {
	persons := make([]*models.Person, 0)
	if err := zoom.NewQuery("Person").Scan(&persons); err != nil {
		return c.RenderError(err)
	}
	return c.RenderJson(persons)
}

func (c Persons) Create(name string, age int) revel.Result {
	p := &models.Person{Name: name, Age: age}
	if err := zoom.Save(p); err != nil {
		return c.RenderError(err)
	}
	return c.RenderJson(p)
}

func (c Persons) Show(id string) revel.Result {
	p, err := zoom.FindById("Person", id)
	if err != nil {
		return c.RenderError(err)
	}
	return c.RenderJson(p)
}

func (c Persons) Update(name string, age int, id string) revel.Result {
	p := &models.Person{}
	if err := zoom.ScanById(id, p); err != nil {
		return c.RenderText("{'error': %s}", err)
	}
	if _, ok := c.Params.Values["name"]; ok {
		p.Name = name
	}
	if _, ok := c.Params.Values["age"]; ok {
		p.Age = age
	}
	if err := zoom.Save(p); err != nil {
		return c.RenderError(err)
	}
	return c.RenderJson(p)
}

func (c Persons) Delete(id string) revel.Result {
	if err := zoom.DeleteById("Person", id); err != nil {
		return c.RenderError(err)
	}
	return c.RenderText("{'message': 'ok'}")
}
