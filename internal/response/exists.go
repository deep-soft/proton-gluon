package response

import (
	"fmt"

	"github.com/ProtonMail/gluon/imap"
)

type exists struct {
	count imap.SeqID
}

func isExists(r Response) bool {
	_, ok := r.(*exists)
	return ok
}

func existsHasHigherID(a, b Response) bool {
	existsA, ok := a.(*exists)
	if !ok {
		return false
	}

	existsB, ok := b.(*exists)
	if !ok {
		return false
	}

	return existsA.count > existsB.count
}

func Exists() *exists {
	return &exists{}
}

func (r *exists) WithCount(n imap.SeqID) *exists {
	r.count = n
	return r
}

func (r *exists) Send(s Session) error {
	return s.WriteResponse(r.String())
}

func (r *exists) String() string {
	return fmt.Sprintf("* %v EXISTS", r.count)
}
