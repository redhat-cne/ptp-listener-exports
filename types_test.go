package exports

import "testing"

func TestStoredEvent_ToCsv(t *testing.T) {
	tests := []struct {
		name    string
		event   StoredEvent
		wantOut string
	}{
		{
			name:    "test1",
			event:   StoredEvent{EventTimeStamp: "timestamp", EventName: "event1", EventType: "eventtype", EventSource: "eventsource", EventValuesFull: "eventfull", EventValuesShort: "eventshort"},
			wantOut: "timestamp,event1,eventtype,eventsource,eventfull,eventshort",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := tt.event.ToCsv(); gotOut != tt.wantOut {
				t.Errorf("StoredEvent.ToCsv() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
