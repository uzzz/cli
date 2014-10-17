// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/cli/cf/command"
	"github.com/cloudfoundry/cli/cf/command_metadata"
	"github.com/cloudfoundry/cli/cf/requirements"
	"github.com/codegangsta/cli"
)

type FakeCommand struct {
	MetadataStub        func() command_metadata.CommandMetadata
	metadataMutex       sync.RWMutex
	metadataArgsForCall []struct{}
	metadataReturns     struct {
		result1 command_metadata.CommandMetadata
	}
	GetRequirementsStub        func(requirementsFactory requirements.Factory, context *cli.Context) (reqs []requirements.Requirement, err error)
	getRequirementsMutex       sync.RWMutex
	getRequirementsArgsForCall []struct {
		requirementsFactory requirements.Factory
		context             *cli.Context
	}
	getRequirementsReturns struct {
		result1 []requirements.Requirement
		result2 error
	}
	RunStub        func(context *cli.Context)
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		context *cli.Context
	}
}

func (fake *FakeCommand) Metadata() command_metadata.CommandMetadata {
	fake.metadataMutex.Lock()
	defer fake.metadataMutex.Unlock()
	fake.metadataArgsForCall = append(fake.metadataArgsForCall, struct{}{})
	if fake.MetadataStub != nil {
		return fake.MetadataStub()
	} else {
		return fake.metadataReturns.result1
	}
}

func (fake *FakeCommand) MetadataCallCount() int {
	fake.metadataMutex.RLock()
	defer fake.metadataMutex.RUnlock()
	return len(fake.metadataArgsForCall)
}

func (fake *FakeCommand) MetadataReturns(result1 command_metadata.CommandMetadata) {
	fake.MetadataStub = nil
	fake.metadataReturns = struct {
		result1 command_metadata.CommandMetadata
	}{result1}
}

func (fake *FakeCommand) GetRequirements(requirementsFactory requirements.Factory, context *cli.Context) (reqs []requirements.Requirement, err error) {
	fake.getRequirementsMutex.Lock()
	defer fake.getRequirementsMutex.Unlock()
	fake.getRequirementsArgsForCall = append(fake.getRequirementsArgsForCall, struct {
		requirementsFactory requirements.Factory
		context             *cli.Context
	}{requirementsFactory, context})
	if fake.GetRequirementsStub != nil {
		return fake.GetRequirementsStub(requirementsFactory, context)
	} else {
		return fake.getRequirementsReturns.result1, fake.getRequirementsReturns.result2
	}
}

func (fake *FakeCommand) GetRequirementsCallCount() int {
	fake.getRequirementsMutex.RLock()
	defer fake.getRequirementsMutex.RUnlock()
	return len(fake.getRequirementsArgsForCall)
}

func (fake *FakeCommand) GetRequirementsArgsForCall(i int) (requirements.Factory, *cli.Context) {
	fake.getRequirementsMutex.RLock()
	defer fake.getRequirementsMutex.RUnlock()
	return fake.getRequirementsArgsForCall[i].requirementsFactory, fake.getRequirementsArgsForCall[i].context
}

func (fake *FakeCommand) GetRequirementsReturns(result1 []requirements.Requirement, result2 error) {
	fake.GetRequirementsStub = nil
	fake.getRequirementsReturns = struct {
		result1 []requirements.Requirement
		result2 error
	}{result1, result2}
}

func (fake *FakeCommand) Run(context *cli.Context) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		context *cli.Context
	}{context})
	if fake.RunStub != nil {
		fake.RunStub(context)
	}
}

func (fake *FakeCommand) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeCommand) RunArgsForCall(i int) *cli.Context {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].context
}

var _ command.Command = new(FakeCommand)