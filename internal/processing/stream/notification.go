/*
   GoToSocial
   Copyright (C) 2021-2023 GoToSocial Authors admin@gotosocial.org

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package stream

import (
	"encoding/json"
	"fmt"

	apimodel "github.com/superseriousbusiness/gotosocial/internal/api/model"
	"github.com/superseriousbusiness/gotosocial/internal/gtsmodel"
	"github.com/superseriousbusiness/gotosocial/internal/stream"
)

// Notify streams the given notification to any open, appropriate streams belonging to the given account.
func (p *Processor) Notify(n *apimodel.Notification, account *gtsmodel.Account) error {
	bytes, err := json.Marshal(n)
	if err != nil {
		return fmt.Errorf("error marshalling notification to json: %s", err)
	}

	return p.toAccount(string(bytes), stream.EventTypeNotification, []string{stream.TimelineNotifications, stream.TimelineHome}, account.ID)
}
