package main

import (
	"errors"
	"log"
	"net/http"
	"reflect"
	"time"

	appointment "github.com/harsh-zen/basic-jaal/schema"
	"go.appointy.com/jaal"
	"go.appointy.com/jaal/introspection"
	"go.appointy.com/jaal/schemabuilder"
)

func main() {
	sb := schemabuilder.NewSchema()
	appointment.RegisterPayload(sb)
	appointment.RegisterInput(sb)
	typ := reflect.TypeOf(time.Time{})
	schemabuilder.RegisterScalar(typ, "DateTime", func(value interface{}, dest reflect.Value) error {
		v, ok := value.(string)
		if !ok {
			return errors.New("invalid type expected string")
		}

		t, err := time.Parse(time.RFC3339, v)
		if err != nil {
			return err
		}

		dest.Set(reflect.ValueOf(t))

		return nil
	})

	s:= &appointment.Server{
		Appointments: []*appointment.Appointment{
			{
				Id: "1",
				Title: "Meeting 1",
				Description: nil,
				StartTime: time.Now(),
				EndTime: time.Now().Add(time.Hour),
			},
			{
				Id: "2",
				Title: "Meeting 2",
				Description: nil,
				StartTime: time.Now().Add(2 * time.Hour),
				EndTime: time.Now().Add(3 * time.Hour),
			},},
	}

	s.RegisterOperations(sb)
	schema, err := sb.Build()
	if err != nil {
		log.Fatal(err)
	}

	introspection.AddIntrospectionToSchema(schema)

	http.Handle("/graphql", jaal.HTTPHandler(schema))
	log.Println("Server ready at 9000")
    if err := http.ListenAndServe(":9000", nil); err!= nil {
        panic(err)
    }
}