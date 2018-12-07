//  Copyright (c) 2018 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package orchestrator

import (
	"github.com/gogo/status"
	"github.com/ligato/cn-infra/datasync/kvdbsync/local"
	"github.com/ligato/cn-infra/db/keyval"
	"github.com/ligato/cn-infra/logging"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"

	"github.com/ligato/vpp-agent/api"
	"github.com/ligato/vpp-agent/api/models"
)

type grpcService struct {
	log logging.Logger
}

// ListModules implements SyncServiceServer.
func (s *grpcService) ListModules(ctx context.Context, req *api.ListModulesRequest) (*api.ListModulesResponse, error) {
	resp := &api.ListModulesResponse{
		Modules: models.GetRegisteredModules(),
	}
	return resp, nil
}

// Sync implements SyncServiceServer.
func (s *grpcService) Sync(ctx context.Context, req *api.SyncRequest) (*api.SyncResponse, error) {
	s.log.Debug("------------------------------")
	s.log.Debugf("=> GRPC SYNC: %d items", len(req.Items))
	s.log.Debug("------------------------------")
	for _, item := range req.Items {
		s.log.Debugf(" - %v", item)
	}
	s.log.Debug("------------------------------")

	// prepare a transaction
	var txn keyval.ProtoTxn
	if req.GetOptions().GetResync() {
		txn = local.NewProtoTxn(Registry.PropagateResync)
	} else {
		txn = local.NewProtoTxn(Registry.PropagateChanges)
	}

	for _, change := range req.Items {
		item := change.GetModel()
		if item == nil {
			return nil, status.Error(codes.InvalidArgument, "change item is nil")
		}
		pb, err := models.Unmarshal(item)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		key, err := models.GetKey(pb)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		if change.Delete {
			txn.Delete(key)
		} else {
			txn.Put(key, pb)
		}
	}

	// commit the transaction
	if err := txn.Commit(); err != nil {
		st := status.New(codes.FailedPrecondition, err.Error())
		return nil, st.Err()
		// TODO: use the WithDetails to return extra info to clients.
		//ds, err := st.WithDetails(&rpc.DebugInfo{Detail: "Local transaction failed!"})
		//if err != nil {
		//	return nil, st.Err()
		//}
		//return nil, ds.Err()
	}

	return &api.SyncResponse{}, nil
}

// Obtain implements SyncServiceServer.
func (*grpcService) Obtain(context.Context, *api.ObtainRequest) (*api.ObtainResponse, error) {
	st := status.New(codes.Unimplemented, "obtain not implemented")
	return nil, st.Err()
}
