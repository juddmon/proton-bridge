// Copyright (c) 2023 Proton AG
//
// This file is part of Proton Mail Bridge.
//
// Proton Mail Bridge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Proton Mail Bridge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Proton Mail Bridge.  If not, see <https://www.gnu.org/licenses/>.

package imapservice

import (
	"context"

	"github.com/ProtonMail/gluon/connector"
)

type IMAPServerManager interface {
	AddIMAPUser(
		ctx context.Context,
		connector connector.Connector,
		addrID string,
		idProvider GluonIDProvider,
		syncStateProvider SyncStateProvider,
	) error

	RemoveIMAPUser(ctx context.Context, deleteData bool, provider GluonIDProvider, addrID ...string) error
}

type NullIMAPServerManager struct{}

func (n NullIMAPServerManager) AddIMAPUser(
	_ context.Context,
	_ connector.Connector,
	_ string,
	_ GluonIDProvider,
	_ SyncStateProvider,
) error {
	return nil
}

func (n NullIMAPServerManager) RemoveIMAPUser(
	_ context.Context,
	_ bool,
	_ GluonIDProvider,
	_ ...string,
) error {
	return nil
}

func NewNullIMAPServerManager() *NullIMAPServerManager {
	return &NullIMAPServerManager{}
}
