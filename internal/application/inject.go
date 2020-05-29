package application

import (
	"github.com/golangee/sql"
	"github.com/worldiety/mercurius/internal/service/setup"
	"github.com/worldiety/mercurius/internal/service/sms"
)

// InjectionContainer is something what happens in java spring boot completely opaque. This is something
// which is neither hard nor a bad thing to maintain, quite the contrary.
type InjectionContainer struct {
	server               *Server
	setupController      *setup.RestController
	smsController        *sms.RestController
	smsMessageRepository sms.MessageRepository
}

func NewInjectionContext(server *Server) *InjectionContainer {
	return &InjectionContainer{server: server}
}

func (i *InjectionContainer) SetupReloader() setup.Reloader {
	return i.server
}

func (i *InjectionContainer) SetupController() *setup.RestController {
	if i.setupController == nil {
		i.setupController = setup.NewRestController(i.SetupReloader())
	}
	return i.setupController
}

func (i *InjectionContainer) SMSMessageRepository() sms.MessageRepository {
	if i.smsMessageRepository == nil {
		dialect := i.server.settings.Database.Dialect()
		if err := sql.NewRepository(dialect, &i.smsMessageRepository); err != nil {
			panic(err) // programing error
		}
	}

	return i.smsMessageRepository
}

func (i *InjectionContainer) SMSController() *sms.RestController {
	if i.smsController == nil {
		i.smsController = sms.NewRestController(i.SMSMessageRepository())
	}
	return i.smsController
}
