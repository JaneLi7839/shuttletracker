package database

import (
	"errors"
	"time"

	"github.com/wtg/shuttletracker/model"
)

// Database is an interface that can be implemented by a database backend.
type Database interface {
	// Routes
	CreateRoute(route *model.Route) error
	DeleteRoute(routeID int64) error
	GetRoute(routeID int64) (model.Route, error)
	GetRoutes() ([]model.Route, error)
	ModifyRoute(route *model.Route) error

	// Stops
	CreateStop(stop *model.Stop) error
	DeleteStop(stopID string) error
	GetStops() ([]model.Stop, error)
	// GetStopsForRoute(routeID string) ([]model.Stop, error)
	// ModifyStop(stop *model.Stop) error
	AddStopToRoute(stopID string, routeID string) error

	// Vehicles
	CreateVehicle(vehicle *model.Vehicle) error
	DeleteVehicle(vehicleID int64) error
	GetVehicle(vehicleID int64) (model.Vehicle, error)
	GetVehicles() ([]model.Vehicle, error)
	GetVehicleByITrakID(itrakID int) (model.Vehicle, error)
	GetEnabledVehicles() ([]model.Vehicle, error)
	ModifyVehicle(vehicle *model.Vehicle) error

	// Updates
	CreateUpdate(update *model.Update) error
	DeleteUpdatesBefore(before time.Time) (int64, error)
	// GetUpdatesSince(since time.Time) ([]model.VehicleUpdate, error)
	GetUpdatesForVehicleSince(vehicleID int64, since time.Time) ([]model.Update, error)
	GetLastUpdateForVehicle(vehicleID int64) (model.Update, error)

	// Users
	GetUsers() ([]model.User, error)
}

var (
	ErrVehicleNotFound = errors.New("Vehicle not found.")
	ErrUpdateNotFound  = errors.New("Update not found.")
)
