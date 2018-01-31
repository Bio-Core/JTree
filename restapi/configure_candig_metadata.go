// Code generated by go-swagger;
//Removed DNE becuase VS Code keeps giving warning about editing

package restapi

import (
	"crypto/tls"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	database "github.com/CanDIG/candig_mds/database"
	repos "github.com/CanDIG/candig_mds/repos"

	"github.com/CanDIG/candig_mds/models"
	operations "github.com/CanDIG/candig_mds/restapi/operations"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name candig_metadata --spec ../swagger.yaml

//
// Use a simple in-memory int -> model map for Individual and biosample
//

var individuals = make(map[int64]*models.Individual)
var lastIndividualId int64
var individualLock = &sync.Mutex{}

var biosamples = make(map[int64]*models.Biosample)
var lastBiosampleId int64
var biosampleLock = &sync.Mutex{}

func newIndividualID() int64 {
	return atomic.AddInt64(&lastIndividualId, 1)
}

func newBiosampleID() int64 {
	return atomic.AddInt64(&lastBiosampleId, 1)
}

func addIndividual(individual *models.Individual) error {
	if individual == nil {
		return errors.New(500, "item must be present")
	}

	individualLock.Lock()
	defer individualLock.Unlock()

	var newID = newIndividualID()
	individual.ID = strconv.FormatInt(newID, 10)
	repos.InsertIndividual("individual", individual)

	return nil
}

func addBiosample(biosample *models.Biosample) error {
	if biosample == nil {
		return errors.New(500, "item must be present")
	}

	biosampleLock.Lock()
	defer biosampleLock.Unlock()

	var newID = newBiosampleID()
	biosample.ID = strconv.FormatInt(newID, 10)
	repos.InsertBiosample("biosample", biosample)

	return nil
}

func allIndividuals(query string) (result []*models.Individual) {
	individuals := repos.GetAllIndividuals("individual")
	result = make([]*models.Individual, 0)
	for _, item := range individuals {
		result = append(result, item)
	}
	return
}

func allBiosamples(query string) (result []*models.Biosample) {
	list := repos.GetAllBiosamples("biosample")
	result = make([]*models.Biosample, 0)
	for _, item := range list {
		result = append(result, item)
	}
	return
}

func configureFlags(api *operations.CandigMetadataAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CandigMetadataAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	//Configure database connection
	database.DatabaseInit("candig", "mongodb://localhost:27017/")

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.AddBiosampleHandler = operations.AddBiosampleHandlerFunc(func(params operations.AddBiosampleParams) middleware.Responder {
		if err := addBiosample(params.Biosample); err != nil {
			return operations.NewAddBiosampleBadRequest()
		}
		return operations.NewAddBiosampleCreated()
	})
	api.AddIndividualHandler = operations.AddIndividualHandlerFunc(func(params operations.AddIndividualParams) middleware.Responder {
		if err := addIndividual(params.Individual); err != nil {
			return operations.NewAddIndividualBadRequest()
		}
		return operations.NewAddIndividualCreated()
	})
	api.GetBiosampleHandler = operations.GetBiosampleHandlerFunc(func(params operations.GetBiosampleParams) middleware.Responder {
		return operations.NewGetBiosampleOK().WithPayload(allBiosamples(params.BiosampleID))
	})
	api.GetIndividualHandler = operations.GetIndividualHandlerFunc(func(params operations.GetIndividualParams) middleware.Responder {
		return operations.NewGetIndividualOK().WithPayload(allIndividuals(params.IndividualID))
	})
	api.SearchBiosampleHandler = operations.SearchBiosampleHandlerFunc(func(params operations.SearchBiosampleParams) middleware.Responder {
		return operations.NewGetBiosampleOK().WithPayload(allBiosamples(""))
	})
	api.SearchIndividualHandler = operations.SearchIndividualHandlerFunc(func(params operations.SearchIndividualParams) middleware.Responder {
		return operations.NewGetIndividualOK().WithPayload(allIndividuals(""))
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
