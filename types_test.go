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
			event:   StoredEvent{EventTimeStamp: "timestamp", EventType: "eventtype", EventSource: "eventsource", EventValues: StoredEvent{"value1": "1", "Value2": "34"}},
			wantOut: "timestamp,eventtype,eventsource,value1,1,Value2,34",
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
