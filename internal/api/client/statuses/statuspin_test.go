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

package statuses_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/superseriousbusiness/gotosocial/internal/ap"
	"github.com/superseriousbusiness/gotosocial/internal/api/client/statuses"
	apimodel "github.com/superseriousbusiness/gotosocial/internal/api/model"
	"github.com/superseriousbusiness/gotosocial/internal/config"
	"github.com/superseriousbusiness/gotosocial/internal/gtserror"
	"github.com/superseriousbusiness/gotosocial/internal/gtsmodel"
	"github.com/superseriousbusiness/gotosocial/internal/id"
	"github.com/superseriousbusiness/gotosocial/internal/oauth"
	"github.com/superseriousbusiness/gotosocial/testrig"
)

type StatusPinTestSuite struct {
	StatusStandardTestSuite
}

func (suite *StatusPinTestSuite) createPin(
	expectedHTTPStatus int,
	expectedBody string,
	targetStatusID string,
) (*apimodel.Status, error) {
	// instantiate recorder + test context
	recorder := httptest.NewRecorder()
	ctx, _ := testrig.CreateGinTestContext(recorder, nil)
	ctx.Set(oauth.SessionAuthorizedAccount, suite.testAccounts["local_account_1"])
	ctx.Set(oauth.SessionAuthorizedToken, oauth.DBTokenToToken(suite.testTokens["local_account_1"]))
	ctx.Set(oauth.SessionAuthorizedApplication, suite.testApplications["application_1"])
	ctx.Set(oauth.SessionAuthorizedUser, suite.testUsers["local_account_1"])

	// create the request
	ctx.Request = httptest.NewRequest(http.MethodPost, config.GetProtocol()+"://"+config.GetHost()+"/api/"+statuses.BasePath+"/"+targetStatusID+"/pin", nil)
	ctx.Request.Header.Set("accept", "application/json")
	ctx.AddParam(statuses.IDKey, targetStatusID)

	// trigger the handler
	suite.statusModule.StatusPinPOSTHandler(ctx)

	// read the response
	result := recorder.Result()
	defer result.Body.Close()

	b, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	errs := gtserror.MultiError{}

	// check code + body
	if resultCode := recorder.Code; expectedHTTPStatus != resultCode {
		errs = append(errs, fmt.Sprintf("expected %d got %d", expectedHTTPStatus, resultCode))
	}

	// if we got an expected body, return early
	if expectedBody != "" && string(b) != expectedBody {
		errs = append(errs, fmt.Sprintf("expected %s got %s", expectedBody, string(b)))
	}

	if len(errs) > 0 {
		return nil, errs.Combine()
	}

	resp := &apimodel.Status{}
	if err := json.Unmarshal(b, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (suite *StatusPinTestSuite) TestPinStatusPublicOK() {
	// Pin an unpinned public status that this account owns.
	targetStatus := suite.testStatuses["local_account_1_status_1"]

	resp, err := suite.createPin(http.StatusOK, "", targetStatus.ID)
	if err != nil {
		suite.FailNow(err.Error())
	}

	suite.True(resp.Pinned)
}

func (suite *StatusPinTestSuite) TestPinStatusFollowersOnlyOK() {
	// Pin an unpinned followers only status that this account owns.
	targetStatus := suite.testStatuses["local_account_1_status_5"]

	resp, err := suite.createPin(http.StatusOK, "", targetStatus.ID)
	if err != nil {
		suite.FailNow(err.Error())
	}

	suite.True(resp.Pinned)
}

func (suite *StatusPinTestSuite) TestPinStatusTwiceError() {
	// Try to pin a status that's already been pinned.
	targetStatus := &gtsmodel.Status{}
	*targetStatus = *suite.testStatuses["local_account_1_status_5"]
	targetStatus.PinnedAt = time.Now()

	if err := suite.db.UpdateStatus(context.Background(), targetStatus); err != nil {
		suite.FailNow(err.Error())
	}

	if _, err := suite.createPin(
		http.StatusUnprocessableEntity,
		`{"error":"Unprocessable Entity: status already pinned"}`,
		targetStatus.ID,
	); err != nil {
		suite.FailNow(err.Error())
	}
}

func (suite *StatusPinTestSuite) TestPinStatusOtherAccountError() {
	// Try to pin a status that doesn't belong to us.
	targetStatus := suite.testStatuses["admin_account_status_1"]

	if _, err := suite.createPin(
		http.StatusUnprocessableEntity,
		`{"error":"Unprocessable Entity: status 01F8MH75CBF9JFX4ZAD54N0W0R does not belong to account 01F8MH1H7YV1Z7D2C8K2730QBF"}`,
		targetStatus.ID,
	); err != nil {
		suite.FailNow(err.Error())
	}
}

func (suite *StatusPinTestSuite) TestPinStatusTooManyPins() {
	// Test pinning too many statuses.
	testAccount := suite.testAccounts["local_account_1"]

	// Spam 10 pinned statuses into the database.
	ctx := context.Background()
	for i := range make([]interface{}, 10) {
		status := &gtsmodel.Status{
			ID:                  id.NewULID(),
			PinnedAt:            time.Now(),
			URL:                 "stub " + strconv.Itoa(i),
			URI:                 "stub " + strconv.Itoa(i),
			Local:               testrig.TrueBool(),
			AccountID:           testAccount.ID,
			AccountURI:          testAccount.URI,
			Visibility:          gtsmodel.VisibilityPublic,
			Federated:           testrig.TrueBool(),
			Boostable:           testrig.TrueBool(),
			Replyable:           testrig.TrueBool(),
			Likeable:            testrig.TrueBool(),
			ActivityStreamsType: ap.ObjectNote,
		}
		if err := suite.db.PutStatus(ctx, status); err != nil {
			suite.FailNow(err.Error())
		}
	}

	// Try to pin one more status as a treat.
	targetStatus := suite.testStatuses["local_account_1_status_1"]
	if _, err := suite.createPin(
		http.StatusUnprocessableEntity,
		`{"error":"Unprocessable Entity: status pin limit exceeded, you've already pinned 10 status(es) out of 10"}`,
		targetStatus.ID,
	); err != nil {
		suite.FailNow(err.Error())
	}
}

func TestStatusPinTestSuite(t *testing.T) {
	suite.Run(t, new(StatusPinTestSuite))
}
