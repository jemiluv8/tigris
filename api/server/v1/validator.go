// Copyright 2022 Tigris Data, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Validator interface {
	Validate() error
}

func (x *InsertRequest) Validate() error {
	if err := isValidCollectionAndDatabase(x.Collection, x.Db); err != nil {
		return err
	}

	if len(x.Documents) == 0 {
		return status.Errorf(codes.InvalidArgument, "empty documents received")
	}
	return nil
}

func (x *ReplaceRequest) Validate() error {
	if err := isValidCollectionAndDatabase(x.Collection, x.Db); err != nil {
		return err
	}

	if len(x.Documents) == 0 {
		return status.Errorf(codes.InvalidArgument, "empty documents received")
	}
	return nil
}

func (x *UpdateRequest) Validate() error {
	if err := isValidCollectionAndDatabase(x.Collection, x.Db); err != nil {
		return err
	}

	if x.Filter == nil {
		return status.Errorf(codes.InvalidArgument, "filter is a required field")
	}
	return nil
}

func (x *DeleteRequest) Validate() error {
	if err := isValidCollectionAndDatabase(x.Collection, x.Db); err != nil {
		return err
	}

	if x.Filter == nil {
		return status.Errorf(codes.InvalidArgument, "filter is a required field")
	}
	return nil
}

func (x *ReadRequest) Validate() error {
	if err := isValidCollectionAndDatabase(x.Collection, x.Db); err != nil {
		return err
	}

	return nil
}

func (x *CreateCollectionRequest) Validate() error {
	if err := isValidCollectionAndDatabase(x.Collection, x.Db); err != nil {
		return err
	}

	if x.Schema == nil {
		return status.Errorf(codes.InvalidArgument, "schema is a required during collection creation")
	}

	return nil
}

func (x *DropCollectionRequest) Validate() error {
	if err := isValidCollectionAndDatabase(x.Collection, x.Db); err != nil {
		return err
	}

	return nil
}

func (x *AlterCollectionRequest) Validate() error {
	if err := isValidCollectionAndDatabase(x.Collection, x.Db); err != nil {
		return err
	}

	return nil
}

func (x *TruncateCollectionRequest) Validate() error {
	if err := isValidCollectionAndDatabase(x.Collection, x.Db); err != nil {
		return err
	}

	return nil
}

func isValidCollection(name string) error {
	if len(name) == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid collection name %s", name)
	}

	return nil
}

func isValidDatabase(name string) error {
	if len(name) == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid database name %s", name)
	}

	return nil
}

func isValidCollectionAndDatabase(c string, db string) error {
	if err := isValidCollection(c); err != nil {
		return err
	}

	if err := isValidDatabase(db); err != nil {
		return err
	}

	return nil
}
