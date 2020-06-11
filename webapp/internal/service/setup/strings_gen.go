// Code generated by go generate; DO NOT EDIT.
// This file was generated by github.com/golangee/i18n

package setup

import (
	"fmt"
	i18n "github.com/golangee/i18n"
)

func init() {
	var tag string

	// from strings.xml
	tag = "und"

	i18n.ImportValue(i18n.NewText(tag, "btn_accept", "accept"))
	i18n.ImportValue(i18n.NewText(tag, "btn_apply", "apply"))
	i18n.ImportValue(i18n.NewText(tag, "btn_back", "back"))
	i18n.ImportValue(i18n.NewText(tag, "btn_next", "next"))
	i18n.ImportValue(i18n.NewText(tag, "btn_start", "start"))
	i18n.ImportValue(i18n.NewText(tag, "setup_db_host", "database host"))
	i18n.ImportValue(i18n.NewText(tag, "setup_db_name", "database name"))
	i18n.ImportValue(i18n.NewText(tag, "setup_driver", "driver"))
	i18n.ImportValue(i18n.NewText(tag, "setup_http_address", "http bind address"))
	i18n.ImportValue(i18n.NewText(tag, "setup_license", "worldiety Enterprise Edition (EE) License (the \"EE License\")\n        Copyright (c) 2020 worldiety GmbH\n\n        This software and the associated documentation may only be used if you or the natural\n        or legal entity you represent, have a corresponding individual agreement with the worldiety GmbH\n        about the use of this software. Unless otherwise agreed, you have, after the completion of all outstanding\n        payments,\n        a non-exclusive, worldwide, perpetual and irrevocable right to use the software. In addition, you are granted\n        permission, in accordance with the agreements made, to use the software and publish patches for the software.\n        In doing so, you agree that worldiety receives an exclusive, worldwide, unlimited and irrevocable right of use\n        of\n        the modifications. This also applies in particular to commissioned modifications, which are carried out by third\n        parties or by worldiety within the scope of service or contracts for work and services arise.\n\n        You are not permitted to copy, merge or combine this software as an independent product,\n        to publish, distribute, sublicense or sell Deviating from this. Unaffected is your permission to copy,\n        distribute,\n        sublicense or sell this software as auxiliary or accessory of your software product.\n\n        Used components of third parties, which are integrated into the worldiety software, remain under the respective\n        Original license of the respective licensor.\n    "))
	i18n.ImportValue(i18n.NewText(tag, "setup_password", "password"))
	i18n.ImportValue(i18n.NewText(tag, "setup_path", "path"))
	i18n.ImportValue(i18n.NewText(tag, "setup_port", "port"))
	i18n.ImportValue(i18n.NewText(tag, "setup_ssl_mode", "SSL mode"))
	i18n.ImportValue(i18n.NewText(tag, "setup_stepper_database", "database"))
	i18n.ImportValue(i18n.NewText(tag, "setup_stepper_http", "http"))
	i18n.ImportValue(i18n.NewText(tag, "setup_stepper_license", "license"))
	i18n.ImportValue(i18n.NewText(tag, "setup_stepper_storage", "storage"))
	i18n.ImportValue(i18n.NewText(tag, "setup_stepper_welcome", "welcome"))
	i18n.ImportValue(i18n.NewText(tag, "setup_title_applying", "setting up..."))
	i18n.ImportValue(i18n.NewText(tag, "setup_title_blobs", "blob store"))
	i18n.ImportValue(i18n.NewText(tag, "setup_title_db", "database"))
	i18n.ImportValue(i18n.NewText(tag, "setup_title_failed", "configuration failed"))
	i18n.ImportValue(i18n.NewText(tag, "setup_title_finished", "configuration finished"))
	i18n.ImportValue(i18n.NewText(tag, "setup_title_http", "http configuration"))
	i18n.ImportValue(i18n.NewText(tag, "setup_title_license", "license agreement"))
	i18n.ImportValue(i18n.NewText(tag, "setup_title_welcome", "welcome to mercurius"))
	i18n.ImportValue(i18n.NewText(tag, "setup_user", "user"))
	i18n.ImportValue(i18n.NewText(tag, "setup_welcome", "Welcome to the setup wizard for the login service.\n        In The following steps the basic configurations for the service are carried out.\n\n        After you have accepted the license agreements and made the necessary\n        configurations, the wizard initializes and starts the login service for your\n        environment.\n\n        After completion, it is possible to create users with different clients and\n        initialize sessions for them.\n\n        You can retun to the previous steps at any time to make adjustments to your settings.\n    "))
	_ = tag

}

// Resources wraps the package strings to get invoked safely.
type Resources struct {
	res *i18n.Resources
}

// NewResources creates a new localized resource instance.
func NewResources(locale string) Resources {
	return Resources{i18n.From(locale)}
}

// BtnAccept returns a translated text for "accept"
func (r Resources) BtnAccept() string {
	str, err := r.res.Text("btn_accept")
	if err != nil {
		return fmt.Errorf("MISS!btn_accept: %w", err).Error()
	}
	return str
}

// BtnApply returns a translated text for "apply"
func (r Resources) BtnApply() string {
	str, err := r.res.Text("btn_apply")
	if err != nil {
		return fmt.Errorf("MISS!btn_apply: %w", err).Error()
	}
	return str
}

// BtnBack returns a translated text for "back"
func (r Resources) BtnBack() string {
	str, err := r.res.Text("btn_back")
	if err != nil {
		return fmt.Errorf("MISS!btn_back: %w", err).Error()
	}
	return str
}

// BtnNext returns a translated text for "next"
func (r Resources) BtnNext() string {
	str, err := r.res.Text("btn_next")
	if err != nil {
		return fmt.Errorf("MISS!btn_next: %w", err).Error()
	}
	return str
}

// BtnStart returns a translated text for "start"
func (r Resources) BtnStart() string {
	str, err := r.res.Text("btn_start")
	if err != nil {
		return fmt.Errorf("MISS!btn_start: %w", err).Error()
	}
	return str
}

// SetupDbHost returns a translated text for "database host"
func (r Resources) SetupDbHost() string {
	str, err := r.res.Text("setup_db_host")
	if err != nil {
		return fmt.Errorf("MISS!setup_db_host: %w", err).Error()
	}
	return str
}

// SetupDbName returns a translated text for "database name"
func (r Resources) SetupDbName() string {
	str, err := r.res.Text("setup_db_name")
	if err != nil {
		return fmt.Errorf("MISS!setup_db_name: %w", err).Error()
	}
	return str
}

// SetupDriver returns a translated text for "driver"
func (r Resources) SetupDriver() string {
	str, err := r.res.Text("setup_driver")
	if err != nil {
		return fmt.Errorf("MISS!setup_driver: %w", err).Error()
	}
	return str
}

// SetupHttpAddress returns a translated text for "http bind address"
func (r Resources) SetupHttpAddress() string {
	str, err := r.res.Text("setup_http_address")
	if err != nil {
		return fmt.Errorf("MISS!setup_http_address: %w", err).Error()
	}
	return str
}

/*
SetupLicense returns a translated text for "worldiety Enterprise Edition (EE) License (the "EE License")
        Copyright (c) 2020 worldiety GmbH

        This software and the associated documentation may only be used if you or the natural
        or legal entity you represent, have a corresponding individual agreement with the worldiety GmbH
        about the use of this software. Unless otherwise agreed, you have, after the completion of all outstanding
        payments,
        a non-exclusive, worldwide, perpetual and irrevocable right to use the software. In addition, you are granted
        permission, in accordance with the agreements made, to use the software and publish patches for the software.
        In doing so, you agree that worldiety receives an exclusive, worldwide, unlimited and irrevocable right of use
        of
        the modifications. This also applies in particular to commissioned modifications, which are carried out by third
        parties or by worldiety within the scope of service or contracts for work and services arise.

        You are not permitted to copy, merge or combine this software as an independent product,
        to publish, distribute, sublicense or sell Deviating from this. Unaffected is your permission to copy,
        distribute,
        sublicense or sell this software as auxiliary or accessory of your software product.

        Used components of third parties, which are integrated into the worldiety software, remain under the respective
        Original license of the respective licensor.
    "
*/
func (r Resources) SetupLicense() string {
	str, err := r.res.Text("setup_license")
	if err != nil {
		return fmt.Errorf("MISS!setup_license: %w", err).Error()
	}
	return str
}

// SetupPassword returns a translated text for "password"
func (r Resources) SetupPassword() string {
	str, err := r.res.Text("setup_password")
	if err != nil {
		return fmt.Errorf("MISS!setup_password: %w", err).Error()
	}
	return str
}

// SetupPath returns a translated text for "path"
func (r Resources) SetupPath() string {
	str, err := r.res.Text("setup_path")
	if err != nil {
		return fmt.Errorf("MISS!setup_path: %w", err).Error()
	}
	return str
}

// SetupPort returns a translated text for "port"
func (r Resources) SetupPort() string {
	str, err := r.res.Text("setup_port")
	if err != nil {
		return fmt.Errorf("MISS!setup_port: %w", err).Error()
	}
	return str
}

// SetupSslMode returns a translated text for "SSL mode"
func (r Resources) SetupSslMode() string {
	str, err := r.res.Text("setup_ssl_mode")
	if err != nil {
		return fmt.Errorf("MISS!setup_ssl_mode: %w", err).Error()
	}
	return str
}

// SetupStepperDatabase returns a translated text for "database"
func (r Resources) SetupStepperDatabase() string {
	str, err := r.res.Text("setup_stepper_database")
	if err != nil {
		return fmt.Errorf("MISS!setup_stepper_database: %w", err).Error()
	}
	return str
}

// SetupStepperHttp returns a translated text for "http"
func (r Resources) SetupStepperHttp() string {
	str, err := r.res.Text("setup_stepper_http")
	if err != nil {
		return fmt.Errorf("MISS!setup_stepper_http: %w", err).Error()
	}
	return str
}

// SetupStepperLicense returns a translated text for "license"
func (r Resources) SetupStepperLicense() string {
	str, err := r.res.Text("setup_stepper_license")
	if err != nil {
		return fmt.Errorf("MISS!setup_stepper_license: %w", err).Error()
	}
	return str
}

// SetupStepperStorage returns a translated text for "storage"
func (r Resources) SetupStepperStorage() string {
	str, err := r.res.Text("setup_stepper_storage")
	if err != nil {
		return fmt.Errorf("MISS!setup_stepper_storage: %w", err).Error()
	}
	return str
}

// SetupStepperWelcome returns a translated text for "welcome"
func (r Resources) SetupStepperWelcome() string {
	str, err := r.res.Text("setup_stepper_welcome")
	if err != nil {
		return fmt.Errorf("MISS!setup_stepper_welcome: %w", err).Error()
	}
	return str
}

// SetupTitleApplying returns a translated text for "setting up..."
func (r Resources) SetupTitleApplying() string {
	str, err := r.res.Text("setup_title_applying")
	if err != nil {
		return fmt.Errorf("MISS!setup_title_applying: %w", err).Error()
	}
	return str
}

// SetupTitleBlobs returns a translated text for "blob store"
func (r Resources) SetupTitleBlobs() string {
	str, err := r.res.Text("setup_title_blobs")
	if err != nil {
		return fmt.Errorf("MISS!setup_title_blobs: %w", err).Error()
	}
	return str
}

// SetupTitleDb returns a translated text for "database"
func (r Resources) SetupTitleDb() string {
	str, err := r.res.Text("setup_title_db")
	if err != nil {
		return fmt.Errorf("MISS!setup_title_db: %w", err).Error()
	}
	return str
}

// SetupTitleFailed returns a translated text for "configuration failed"
func (r Resources) SetupTitleFailed() string {
	str, err := r.res.Text("setup_title_failed")
	if err != nil {
		return fmt.Errorf("MISS!setup_title_failed: %w", err).Error()
	}
	return str
}

// SetupTitleFinished returns a translated text for "configuration finished"
func (r Resources) SetupTitleFinished() string {
	str, err := r.res.Text("setup_title_finished")
	if err != nil {
		return fmt.Errorf("MISS!setup_title_finished: %w", err).Error()
	}
	return str
}

// SetupTitleHttp returns a translated text for "http configuration"
func (r Resources) SetupTitleHttp() string {
	str, err := r.res.Text("setup_title_http")
	if err != nil {
		return fmt.Errorf("MISS!setup_title_http: %w", err).Error()
	}
	return str
}

// SetupTitleLicense returns a translated text for "license agreement"
func (r Resources) SetupTitleLicense() string {
	str, err := r.res.Text("setup_title_license")
	if err != nil {
		return fmt.Errorf("MISS!setup_title_license: %w", err).Error()
	}
	return str
}

// SetupTitleWelcome returns a translated text for "welcome to mercurius"
func (r Resources) SetupTitleWelcome() string {
	str, err := r.res.Text("setup_title_welcome")
	if err != nil {
		return fmt.Errorf("MISS!setup_title_welcome: %w", err).Error()
	}
	return str
}

// SetupUser returns a translated text for "user"
func (r Resources) SetupUser() string {
	str, err := r.res.Text("setup_user")
	if err != nil {
		return fmt.Errorf("MISS!setup_user: %w", err).Error()
	}
	return str
}

/*
SetupWelcome returns a translated text for "Welcome to the setup wizard for the login service.
        In The following steps the basic configurations for the service are carried out.

        After you have accepted the license agreements and made the necessary
        configurations, the wizard initializes and starts the login service for your
        environment.

        After completion, it is possible to create users with different clients and
        initialize sessions for them.

        You can retun to the previous steps at any time to make adjustments to your settings.
    "
*/
func (r Resources) SetupWelcome() string {
	str, err := r.res.Text("setup_welcome")
	if err != nil {
		return fmt.Errorf("MISS!setup_welcome: %w", err).Error()
	}
	return str
}

// FuncMap returns the named functions to be used with a template
func (r Resources) FuncMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["BtnAccept"] = r.BtnAccept
	m["BtnApply"] = r.BtnApply
	m["BtnBack"] = r.BtnBack
	m["BtnNext"] = r.BtnNext
	m["BtnStart"] = r.BtnStart
	m["SetupDbHost"] = r.SetupDbHost
	m["SetupDbName"] = r.SetupDbName
	m["SetupDriver"] = r.SetupDriver
	m["SetupHttpAddress"] = r.SetupHttpAddress
	m["SetupLicense"] = r.SetupLicense
	m["SetupPassword"] = r.SetupPassword
	m["SetupPath"] = r.SetupPath
	m["SetupPort"] = r.SetupPort
	m["SetupSslMode"] = r.SetupSslMode
	m["SetupStepperDatabase"] = r.SetupStepperDatabase
	m["SetupStepperHttp"] = r.SetupStepperHttp
	m["SetupStepperLicense"] = r.SetupStepperLicense
	m["SetupStepperStorage"] = r.SetupStepperStorage
	m["SetupStepperWelcome"] = r.SetupStepperWelcome
	m["SetupTitleApplying"] = r.SetupTitleApplying
	m["SetupTitleBlobs"] = r.SetupTitleBlobs
	m["SetupTitleDb"] = r.SetupTitleDb
	m["SetupTitleFailed"] = r.SetupTitleFailed
	m["SetupTitleFinished"] = r.SetupTitleFinished
	m["SetupTitleHttp"] = r.SetupTitleHttp
	m["SetupTitleLicense"] = r.SetupTitleLicense
	m["SetupTitleWelcome"] = r.SetupTitleWelcome
	m["SetupUser"] = r.SetupUser
	m["SetupWelcome"] = r.SetupWelcome
	return m
}
