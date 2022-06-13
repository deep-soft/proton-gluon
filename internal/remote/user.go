package remote

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/ProtonMail/gluon/connector"
	"github.com/ProtonMail/gluon/imap"
	"github.com/ProtonMail/gluon/internal/pchan"
)

// User performs operations against a remote server using a connector.
type User struct {
	userID string

	// path is the path at which the operation queue will be saved to disk.
	path string

	// conn is what the user uses to perform API operations.
	conn connector.Connector

	// updatesCh is the channel that delivers API updates to the mailserver.
	updatesCh chan imap.Update

	// queue is channel of operations that must be performed on the API.
	queue *pchan.PChan[operation]

	// lastOp holds an operation while it has been popped off the queue but not yet executed.
	lastOp operation

	// closed holds whether the operation queue has been closed.
	closed     bool
	closedLock sync.RWMutex

	connMetadataStore connMetadataStore

	// processWG is used to ensure we wait until the process goroutine has finished executing after we close the queue.
	processWG sync.WaitGroup
}

// newUser constructs a new user with the given (IMAP) credentials.
// It serializes its operation queue to a file at the given filepath,
// and performs remote operations using the given connector.
func newUser(userID, path string, conn connector.Connector) (*User, error) {
	user := &User{
		userID:            userID,
		path:              path,
		conn:              conn,
		updatesCh:         make(chan imap.Update),
		queue:             pchan.New[operation](),
		connMetadataStore: newConnMetadataStore(),
	}

	// load any saved operations that were not processed fully before.
	if err := user.load(); err != nil {
		return nil, err
	}

	// send connector updates along to the mailserver.
	go user.forward(conn.GetUpdates())

	user.processWG.Add(1)
	// process remote operations on the operation queue.
	go user.process()

	return user, nil
}

// GetUpdates returns a channel on which updates from the server are sent.
func (user *User) GetUpdates() <-chan imap.Update {
	return user.updatesCh
}

// Close closes the remote user.
func (user *User) Close() error {
	ops, err := user.closeQueue()
	if err != nil {
		return fmt.Errorf("failed to close queue: %w", err)
	}

	// Wait until any remaining operations popped by the process go routine finish executing
	user.processWG.Wait()

	if user.lastOp != nil {
		ops = append([]operation{user.lastOp}, ops...)
	}

	// Append delete operations to make sure that when we reprocess the queue after loading from disk, the
	// stored values in connMetadataStore get erased and don't conflict with new sessions
	for _, id := range user.connMetadataStore.GetActiveStoreIDs() {
		ops = append(ops, &OpConnMetadataStoreDelete{
			OperationBase: OperationBase{MetadataID: id},
		})
	}

	serializeData := userSerializedData{
		PendingOps:        ops,
		ConnMetadataStore: user.connMetadataStore,
	}

	if err := serializeData.saveToFile(user.path); err != nil {
		return err
	}

	return nil
}

// forward pulls updates off the stream and forwards them to the outgoing update channel.
func (user *User) forward(updateCh <-chan imap.Update) {
	for update := range updateCh {
		user.send(update)
	}
}

// send sends the update on the user's updates channel, optionally blocking until it has been processed.
func (user *User) send(update imap.Update, withBlock ...bool) {
	user.updatesCh <- update

	if len(withBlock) > 0 && withBlock[0] {
		update.Wait()
	}
}

// load reads queued remote operations from disk and fills the operation queue with them.
func (user *User) load() error {
	serializedData := userSerializedData{
		PendingOps:        []operation{},
		ConnMetadataStore: newConnMetadataStore(),
	}

	if err := serializedData.loadFromFile(user.path); err != nil {
		return err
	}

	if err := os.Remove(user.path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		} else if err != nil {
			return err
		}
	}

	for _, op := range serializedData.PendingOps {
		if err := user.pushOp(op); err != nil {
			return err
		}
	}

	user.connMetadataStore = serializedData.ConnMetadataStore

	return nil
}

type userSerializedData struct {
	PendingOps        []operation
	ConnMetadataStore connMetadataStore
}

func (usd *userSerializedData) saveToFile(path string) error {
	b, err := usd.saveToBytes()
	if err != nil {
		return err
	}

	if err := os.WriteFile(path, b, 0o600); err != nil {
		return err
	}

	return nil
}

func (usd *userSerializedData) saveToBytes() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := gob.NewEncoder(buf).Encode(usd); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (usd *userSerializedData) loadFromBytes(data []byte) error {
	return gob.NewDecoder(bytes.NewReader(data)).Decode(usd)
}

func (usd *userSerializedData) loadFromFile(path string) error {
	f, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	} else if err != nil {
		return err
	}

	b, err := io.ReadAll(f)
	if err != nil {
		f.Close()
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	if err := usd.loadFromBytes(b); err != nil {
		return err
	}

	return nil
}
