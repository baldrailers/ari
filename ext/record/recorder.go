package record

import "github.com/CyCoreSystems/ari"

// A Recorder is anything which can "Record"
type Recorder interface {
	Record(string, *ari.RecordingOptions) (*ari.LiveRecordingHandle, error)

	StageRecord(string, *ari.RecordingOptions) (*ari.LiveRecordingHandle, error)

	Subscribe(...string) ari.Subscription
}
