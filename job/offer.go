package job

import (
	"sort"
)

type JobOffer struct {
	Job Job

	// MachineIDs represents a set of machines for which this offer is valid.
	// If nil or len == 0 then all machines. Must be sorted.
	// MachineIDs started life as MachineBootIds in the datastore.
	// It cannot be changed without a migration
	MachineIDs []string `json:"MachineBootIds"`
}

func NewOfferFromJob(j Job, machineIDs []string) *JobOffer {
	return &JobOffer{
		Job:            j,
		MachineIDs: machineIDs,
	}
}

// OfferedTo returns true if job is being offered to specified machine.
func (jo *JobOffer) OfferedTo(machineID string) bool {
	// for backward compatibility: if not populated, assume all machines are considered
	if len(jo.MachineIDs) == 0 {
		return true
	}

	k := sort.SearchStrings(jo.MachineIDs, machineID)
	return k < len(jo.MachineIDs) && jo.MachineIDs[k] == machineID
}

type JobBid struct {
	JobName string

	// MachineID started life as MachineBootId in the datastore.
	// It cannot be changed without a migration.
	MachineID string `json:"MachineBootId"`
}

func NewBid(jobName string, machID string) *JobBid {
	return &JobBid{jobName, machID}
}
