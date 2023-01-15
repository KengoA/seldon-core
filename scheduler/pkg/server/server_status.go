/*
Copyright 2022 Seldon Technologies Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"time"

	"github.com/seldonio/seldon-core/scheduler/v2/pkg/coordinator"

	pb "github.com/seldonio/seldon-core/apis/go/v2/mlops/scheduler"
)

func (s *SchedulerServer) SubscribeModelStatus(req *pb.ModelSubscriptionRequest, stream pb.Scheduler_SubscribeModelStatusServer) error {
	logger := s.logger.WithField("func", "SubscribeModelStatus")
	logger.Infof("Received subscribe request from %s", req.GetSubscriberName())

	fin := make(chan bool)

	s.modelEventStream.mu.Lock()
	s.modelEventStream.streams[stream] = &ModelSubscription{
		name:   req.GetSubscriberName(),
		stream: stream,
		fin:    fin,
	}
	s.modelEventStream.mu.Unlock()

	err := s.sendCurrentModelStatuses(stream)
	if err != nil {
		return err
	}

	ctx := stream.Context()
	// Keep this scope alive because once this scope exits - the stream is closed
	for {
		select {
		case <-fin:
			logger.Infof("Closing stream for %s", req.GetSubscriberName())
			return nil
		case <-ctx.Done():
			logger.Infof("Stream disconnected %s", req.GetSubscriberName())
			s.modelEventStream.mu.Lock()
			delete(s.modelEventStream.streams, stream)
			s.modelEventStream.mu.Unlock()
			return nil
		}
	}
}

// TODO as this could be 1000s of models may need to look at ways to optimize?
func (s *SchedulerServer) sendCurrentModelStatuses(stream pb.Scheduler_SubscribeModelStatusServer) error {
	modelNames := s.modelStore.GetAllModels()
	for _, modelName := range modelNames {
		model, err := s.modelStore.GetModel(modelName)
		if err != nil {
			return err
		}
		ms, err := s.modelStatusImpl(model, false)
		if err != nil {
			return err
		}
		err = stream.Send(ms)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *SchedulerServer) handleModelEvent(event coordinator.ModelEventMsg) {
	logger := s.logger.WithField("func", "handleModelEvent")
	logger.Debugf("Got model event msg for %s", event.String())

	// TODO - Should this spawn a goroutine?
	// Surely if we do we're risking reordering of events, e.g. load/unload -> unload/load?
	err := s.sendModelStatusEvent(event)
	if err != nil {
		logger.WithError(err).Errorf("Failed to update model status for model %s", event.String())
	}
}

func (s *SchedulerServer) StopSendModelEvents() {
	s.modelEventStream.mu.Lock()
	defer s.modelEventStream.mu.Unlock()
	for _, subscription := range s.modelEventStream.streams {
		close(subscription.fin)
	}
}

func (s *SchedulerServer) sendModelStatusEvent(evt coordinator.ModelEventMsg) error {
	logger := s.logger.WithField("func", "sendModelStatusEvent")
	model, err := s.modelStore.GetModel(evt.ModelName)
	if err != nil {
		return err
	}
	if model.GetLatest() != nil && model.GetLatest().GetVersion() == evt.ModelVersion {
		ms, err := s.modelStatusImpl(model, false)
		if err != nil {
			logger.WithError(err).Errorf("Failed to create model status for model %s", evt.String())
			return err
		}
		s.modelEventStream.mu.Lock()
		defer s.modelEventStream.mu.Unlock()
		for stream, subscription := range s.modelEventStream.streams {
			err := stream.Send(ms)
			if err != nil {
				logger.WithError(err).Errorf("Failed to send model status event to %s for %s", subscription.name, evt.String())
			}
		}
	}
	return nil
}

func (s *SchedulerServer) SubscribeServerStatus(req *pb.ServerSubscriptionRequest, stream pb.Scheduler_SubscribeServerStatusServer) error {
	logger := s.logger.WithField("func", "SubscribeServerStatus")
	logger.Infof("Received server subscribe request from %s", req.GetSubscriberName())

	fin := make(chan bool)

	s.serverEventStream.mu.Lock()
	s.serverEventStream.streams[stream] = &ServerSubscription{
		name:   req.GetSubscriberName(),
		stream: stream,
		fin:    fin,
	}
	s.serverEventStream.mu.Unlock()

	ctx := stream.Context()
	// Keep this scope alive because once this scope exits - the stream is closed
	for {
		select {
		case <-fin:
			logger.Infof("Closing stream for %s", req.GetSubscriberName())
			return nil
		case <-ctx.Done():
			logger.Infof("Stream disconnected %s", req.GetSubscriberName())
			s.serverEventStream.mu.Lock()
			delete(s.serverEventStream.streams, stream)
			s.serverEventStream.mu.Unlock()
			return nil
		}
	}
}

// TODO - Create a ServerStatusMsg type to disambiguate?
func (s *SchedulerServer) handleServerEvent(event coordinator.ModelEventMsg) {
	logger := s.logger.WithField("func", "handleServerEvent")
	logger.Debugf("Got server state change for %s", event.String())

	// TODO - Should this spawn a goroutine?
	// Surely if we do we're risking reordering of events, e.g. load/unload -> unload/load?
	err := s.updateServerStatus(event)
	if err != nil {
		logger.WithError(err).Errorf("Failed to update server status for model event %s", event.String())
	}
}

func (s *SchedulerServer) StopSendServerEvents() {
	s.serverEventStream.mu.Lock()
	defer s.serverEventStream.mu.Unlock()
	for _, subscription := range s.serverEventStream.streams {
		close(subscription.fin)
	}
}

func (s *SchedulerServer) updateServerStatus(evt coordinator.ModelEventMsg) error {
	logger := s.logger.WithField("func", "sendServerStatusEvent")

	model, err := s.modelStore.GetModel(evt.ModelName)
	if err != nil {
		return err
	}
	modelVersion := model.GetVersion(evt.ModelVersion)
	if modelVersion == nil {
		logger.Warnf("Failed to find model version %s so ignoring event", evt.String())
		return nil
	}
	if modelVersion.Server() == "" {
		logger.Warnf("Empty server for %s so ignoring event", evt.String())
		return nil
	}

	s.serverEventStream.pendingLock.Lock()
	s.serverEventStream.pendingEvents[modelVersion.Server()] = struct{}{}
	if s.serverEventStream.trigger == nil {
		s.serverEventStream.trigger = time.AfterFunc(defaultBatchWaitMillis, s.sendServerStatus)
	}
	s.serverEventStream.pendingLock.Unlock()

	return nil
}

func (s *SchedulerServer) sendServerStatus() {
	logger := s.logger.WithField("func", "sendServerStatus")

	// Sending events may be slow, so allow a new batch to start building as we send.
	s.serverEventStream.pendingLock.Lock()
	s.serverEventStream.trigger = nil
	pendingServers := s.serverEventStream.pendingEvents
	s.serverEventStream.pendingEvents = map[string]struct{}{}
	s.serverEventStream.pendingLock.Unlock()

	// Inform subscriber
	s.serverEventStream.mu.Lock()
	for serverName := range pendingServers {
		server, err := s.modelStore.GetServer(serverName, true, true)
		if err != nil {
			logger.Errorf("Failed to get server %s", serverName)
			continue
		}
		if server == nil {
			logger.Warnf("Server %s does not exist", serverName)
			continue
		}
		ssr := createServerStatusResponse(server)

		for stream, subscription := range s.serverEventStream.streams {
			err := stream.Send(ssr)
			if err != nil {
				logger.WithError(err).Errorf("Failed to send server status event to %s", subscription.name)
			}
		}
	}
	s.serverEventStream.mu.Unlock()
}