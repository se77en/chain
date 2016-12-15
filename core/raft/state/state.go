package state

import (
	"context"

	"github.com/golang/protobuf/proto"

	"chain/core/raft/internal/statepb"
	"chain/errors"
	"chain/log"
)

var ErrAlreadyApplied = errors.New("entry already applied")

// TODO(kr): what data type should we really use?

// State is a general-purpose data store designed to accumulate
// and apply replicated updates from a raft log.
// The zero value is an empty State ready to use.
type State struct {
	state        map[string]string
	peers        map[uint64]string // id -> addr
	appliedIndex uint64
	nextNodeID   uint64
}

// New returns a new State
func New() *State {
	return &State{
		state:      make(map[string]string),
		peers:      make(map[uint64]string),
		nextNodeID: 2,
	}
}

// SetPeerAddr sets the address for the given peer.
func (s *State) SetPeerAddr(id uint64, addr string) {
	s.peers[id] = addr
}

// GetPeerAddr gets the current address for the given peer, if set.
func (s *State) GetPeerAddr(id uint64) (addr string) {
	return s.peers[id]
}

// RemovePeerAddr deletes the current address for the given peer if it exists.
func (s *State) RemovePeerAddr(id uint64) {
	delete(s.peers, id)
}

// RestoreSnapshot decodes data and overwrites the contents of s.
// It should be called with the retrieved snapshot
// when bootstrapping a new node from an existing cluster
// or when recovering from a file on disk.
func (s *State) RestoreSnapshot(data []byte, index uint64) error {
	s.appliedIndex = index
	//TODO (ameets): think about having statepb in state for restore
	snapshot := &statepb.Snapshot{}
	err := proto.Unmarshal(data, snapshot)
	s.peers = snapshot.Peers
	s.state = snapshot.State
	s.nextNodeID = snapshot.NextNodeId
	log.Messagef(context.Background(), "decoded snapshot %#v (err %v)", s, err)
	return errors.Wrap(err)
}

// Snapshot returns an encoded copy of s
// suitable for RestoreSnapshot.
func (s *State) Snapshot() ([]byte, uint64, error) {
	log.Messagef(context.Background(), "encoding snapshot %#v", s)
	data, err := proto.Marshal(&statepb.Snapshot{
		NextNodeId: s.nextNodeID,
		State:      s.state,
		Peers:      s.peers,
	})
	return data, s.appliedIndex, errors.Wrap(err)
}

// Apply applies a raft log entry payload to s.
// For conditional operations returns whether codition was satisfied
// in addition to any errors.
func (s *State) Apply(data []byte, index uint64) (satisfied bool, err error) {
	if index < s.appliedIndex {
		return false, ErrAlreadyApplied
	}
	// TODO(kr): figure out a better entry encoding
	op := &statepb.Op{}
	err = proto.Unmarshal(data, op)
	if err != nil {
		// An error here indicates a malformed update
		// was written to the raft log. We do version
		// negotiation in the transport layer, so this
		// should be impossible; by this point, we are
		// all speaking the same version.
		return false, errors.Wrap(err)
	}
	switch op.Type {
	case statepb.Op_INCREMENT_NEXT_NODE_ID:
		if op.NextNodeId == s.nextNodeID {
			s.nextNodeID++
			satisfied = true
		}
	case statepb.Op_SET:
		s.state[op.Key] = op.Value
		satisfied = true
	default:
		return false, errors.New("unknown operation type")
	}

	s.appliedIndex = index
	return satisfied, nil
}

// Provisional read operation.
func (s *State) Get(key string) (value string) {
	return s.state[key]
}

// Set encodes a set operation setting key to value.
// The encoded op should be committed to the raft log,
// then it can be applied with Apply.
func Set(key, value string) []byte {
	// TODO(kr): make a way to delete things
	// TODO(kr): we prob need other operations too, like conditional writes
	b, _ := proto.Marshal(&statepb.Op{
		Type:  statepb.Op_SET,
		Key:   key,
		Value: value,
	})

	return b
}

// AppliedIndex returns the raft log index (applied index) of current state
func (s *State) AppliedIndex() uint64 {
	return s.appliedIndex
}

// IDCounter
func (s *State) NextNodeID() uint64 {
	return s.nextNodeID
}

func IncrementNextNodeID(oldID uint64) []byte {
	b, _ := proto.Marshal(&statepb.Op{
		Type:       statepb.Op_INCREMENT_NEXT_NODE_ID,
		NextNodeId: oldID,
	})

	return b
}
