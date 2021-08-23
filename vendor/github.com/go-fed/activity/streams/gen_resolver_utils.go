// Code generated by astool. DO NOT EDIT.

package streams

import (
	"context"
	"errors"
	vocab "github.com/go-fed/activity/streams/vocab"
)

// ErrNoCallbackMatch indicates a Resolver could not match the ActivityStreams value to a callback function.
var ErrNoCallbackMatch error = errors.New("activity stream did not match the callback function")

// ErrUnhandledType indicates that an ActivityStreams value has a type that is not handled by the code that has been generated.
var ErrUnhandledType error = errors.New("activity stream did not match any known types")

// ErrPredicateUnmatched indicates that a predicate is accepting a type or interface that does not match an ActivityStreams value's type or interface.
var ErrPredicateUnmatched error = errors.New("activity stream did not match type demanded by predicate")

// errCannotTypeAssertType indicates that the 'type' property returned by the ActivityStreams value cannot be type-asserted to its interface form.
var errCannotTypeAssertType error = errors.New("activity stream type cannot be asserted to its interface")

// ActivityStreamsInterface represents any ActivityStream value code-generated by
// go-fed or compatible with the generated interfaces.
type ActivityStreamsInterface interface {
	// GetTypeName returns the ActiivtyStreams value's type.
	GetTypeName() string
	// VocabularyURI returns the vocabulary's URI as a string.
	VocabularyURI() string
}

// Resolver represents any TypeResolver.
type Resolver interface {
	// Resolve will attempt to resolve an untyped ActivityStreams value into a
	// Go concrete type.
	Resolve(ctx context.Context, o ActivityStreamsInterface) error
}

// IsUnmatchedErr is true when the error indicates that a Resolver was
// unsuccessful due to the ActivityStreams value not matching its callbacks or
// predicates.
func IsUnmatchedErr(err error) bool {
	return err == ErrPredicateUnmatched || err == ErrUnhandledType || err == ErrNoCallbackMatch
}

// ToType attempts to resolve the generic JSON map into a Type.
func ToType(c context.Context, m map[string]interface{}) (t vocab.Type, err error) {
	var r *JSONResolver
	r, err = NewJSONResolver(func(ctx context.Context, i vocab.ActivityStreamsAccept) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsActivity) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsAdd) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsAnnounce) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsApplication) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsArrive) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsArticle) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsAudio) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsBlock) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ForgeFedBranch) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsCollection) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsCollectionPage) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ForgeFedCommit) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsCreate) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsDelete) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsDislike) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsDocument) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.TootEmoji) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsEvent) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsFlag) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsFollow) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsGroup) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.TootIdentityProof) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsIgnore) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsImage) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsIntransitiveActivity) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsInvite) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsJoin) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsLeave) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsLike) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsLink) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsListen) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsMention) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsMove) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsNote) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsObject) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsOffer) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsOrderedCollection) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsOrderedCollectionPage) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsOrganization) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsPage) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsPerson) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsPlace) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsProfile) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.W3IDSecurityV1PublicKey) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ForgeFedPush) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsQuestion) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsRead) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsReject) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsRelationship) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsRemove) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ForgeFedRepository) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsService) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsTentativeAccept) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsTentativeReject) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ForgeFedTicket) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ForgeFedTicketDependency) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsTombstone) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsTravel) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsUndo) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsUpdate) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsVideo) error {
		t = i
		return nil
	}, func(ctx context.Context, i vocab.ActivityStreamsView) error {
		t = i
		return nil
	})
	if err != nil {
		return
	}
	err = r.Resolve(c, m)
	return
}