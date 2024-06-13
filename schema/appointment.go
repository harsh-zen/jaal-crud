package appointment

import (
	"context"
	"strconv"
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

type CreateAppointmentRequest struct {
	Title       string
	StartTime   time.Time
	EndTime     time.Time
	Description *string
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
	payloadObj.FieldFunc("startTime", func(a *Appointment) time.Time{
		return a.StartTime
	})
	payloadObj.FieldFunc("endTime", func(a *Appointment) time.Time {
		return a.EndTime
	})
}

func RegisterInput(schema *schemabuilder.Schema){
	inputObj := schema.InputObject(("CreateAppointmentRequest"), CreateAppointmentRequest{})
	inputObj.FieldFunc("title", func(a *CreateAppointmentRequest, source string) {
		a.Title = source
	})
	inputObj.FieldFunc("startTime", func(a *CreateAppointmentRequest, source time.Time) {
		a.StartTime = source
	})
	inputObj.FieldFunc("endTime", func(a *CreateAppointmentRequest, source time.Time) {
		a.EndTime = source
	})
	inputObj.FieldFunc("description", func(a *CreateAppointmentRequest, source *string) {
		a.Description = source
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
	mutation := schema.Mutation()
	mutation.FieldFunc("createAppointment", func(ctx context.Context, args struct{ Input CreateAppointmentRequest }) *Appointment {
		a := &Appointment{
			Id:          strconv.Itoa(len(s.Appointments) + 1),
			Title:       args.Input.Title,
			Description: args.Input.Description,
			StartTime:   args.Input.StartTime,
			EndTime:     args.Input.EndTime,
		}
		s.Appointments = append(s.Appointments, a)
		return a
	})

	mutation.FieldFunc("updateAppointment", func(ctx context.Context, args struct{ ID string; Input CreateAppointmentRequest }) *Appointment {
		for _, a := range s.Appointments {
			if a.Id == args.ID {
				a.Title = args.Input.Title
				a.Description = args.Input.Description
				a.StartTime = args.Input.StartTime
				a.EndTime = args.Input.EndTime
				return a
			}
		}
		return nil
	})

	mutation.FieldFunc("deleteAppointment", func(ctx context.Context, args struct{ ID string }) *Appointment {
		for i, a := range s.Appointments {
			if a.Id == args.ID {
				s.Appointments = append(s.Appointments[:i], s.Appointments[i+1:]...)
				return a
			}
		}
		return nil
	})
}
