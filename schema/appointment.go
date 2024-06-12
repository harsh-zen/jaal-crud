package appointment

import (
	"context"
	"time"

	"go.appointy.com/jaal/schemabuilder"
)


type Appointment struct {
	Id          string
	Title       string
	Description *string
	StartTime   time.Time
	EndTime     time.Time
}

func RegisterPayload(schema *schemabuilder.Schema){
	payloadObj:= schema.Object("Appointment", Appointment{})
	payloadObj.FieldFunc("id", func(a *Appointment) *schemabuilder.ID {
		return &schemabuilder.ID{Value: a.Id}
	})
	payloadObj.FieldFunc("title", func(a *Appointment) string {
		return a.Title
	})
	payloadObj.FieldFunc("description", func(a *Appointment) *string {
		return a.Description
	})
	payloadObj.FieldFunc("startTime", func(a *Appointment) schemabuilder.Timestamp {
		return a.StartTime
	})
	payloadObj.FieldFunc("endTime", func(a *Appointment) time.Time {
		return a.EndTime
	})
}




type Server struct {
    Appointments []*Appointment
}

func (s *Server) RegisterOperations(schema *schemabuilder.Schema) {
	query := schema.Query()
	query.FieldFunc("appointments", func(ctx context.Context, args struct{}) []*Appointment {
		return s.Appointments
	})

	query.FieldFunc("appointment", func(ctx context.Context, args struct{ ID string }) *Appointment {
		for _, a := range s.Appointments {
			if a.Id == args.ID {
				return a
			}
		}
		return nil
	})
}
