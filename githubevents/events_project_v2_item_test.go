// Copyright 2022 The GithubEvents Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v68/github"
	"sync"
	"testing"
)

func TestOnProjectV2ItemEventAny(t *testing.T) {
	type args struct {
		callbacks []ProjectV2ItemEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectV2ItemEventHandleFunc",
			args: args{
				[]ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectV2ItemEventHandleFuncs",
			args: args{
				[]ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectV2ItemEventAny(tt.args.callbacks...)
			if len(g.onProjectV2ItemEvent[ProjectV2ItemEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectV2ItemEvent[ProjectV2ItemEventAnyAction]))
			}
		})
	}
}

func TestSetOnProjectV2ItemEventAny(t *testing.T) {
	type args struct {
		callbacks []ProjectV2ItemEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectV2ItemEventHandleFunc",
			args: args{
				[]ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectV2ItemEventHandleFuncs",
			args: args{
				[]ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnProjectV2ItemEventAny(func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
				return nil
			})
			g.SetOnProjectV2ItemEventAny(tt.args.callbacks...)
			if len(g.onProjectV2ItemEvent[ProjectV2ItemEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectV2ItemEvent[ProjectV2ItemEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectV2ItemEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectV2ItemEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",

				event: &github.ProjectV2ItemEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",

				event: &github.ProjectV2ItemEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectV2ItemEventAny(func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectV2ItemEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleProjectV2ItemEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectItemEventCreated(t *testing.T) {
	type args struct {
		callbacks []ProjectV2ItemEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectV2ItemEventHandleFunc",
			args: args{
				callbacks: []ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectV2ItemEventHandleFunc",
			args: args{
				callbacks: []ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectItemEventCreated(tt.args.callbacks...)
			if len(g.onProjectV2ItemEvent[ProjectItemEventCreatedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectV2ItemEvent[ProjectItemEventCreatedAction]))
			}
		})
	}
}

func TestSetOnProjectItemEventCreated(t *testing.T) {
	type args struct {
		callbacks []ProjectV2ItemEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectV2ItemEventHandleFunc",
			args: args{
				[]ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectV2ItemEventHandleFuncs",
			args: args{
				[]ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnProjectItemEventCreated(func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
				return nil
			})
			g.SetOnProjectItemEventCreated(tt.args.callbacks...)
			if len(g.onProjectV2ItemEvent[ProjectItemEventCreatedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectV2ItemEvent[ProjectItemEventCreatedAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectItemEventCreated(t *testing.T) {
	action := ProjectItemEventCreatedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectV2ItemEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectItemEventCreated(func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectItemEventCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectItemEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectItemEventEdited(t *testing.T) {
	type args struct {
		callbacks []ProjectV2ItemEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectV2ItemEventHandleFunc",
			args: args{
				callbacks: []ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectV2ItemEventHandleFunc",
			args: args{
				callbacks: []ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectItemEventEdited(tt.args.callbacks...)
			if len(g.onProjectV2ItemEvent[ProjectItemEventEditedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectV2ItemEvent[ProjectItemEventEditedAction]))
			}
		})
	}
}

func TestSetOnProjectItemEventEdited(t *testing.T) {
	type args struct {
		callbacks []ProjectV2ItemEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectV2ItemEventHandleFunc",
			args: args{
				[]ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectV2ItemEventHandleFuncs",
			args: args{
				[]ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnProjectItemEventEdited(func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
				return nil
			})
			g.SetOnProjectItemEventEdited(tt.args.callbacks...)
			if len(g.onProjectV2ItemEvent[ProjectItemEventEditedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectV2ItemEvent[ProjectItemEventEditedAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectItemEventEdited(t *testing.T) {
	action := ProjectItemEventEditedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectV2ItemEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectItemEventEdited(func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectItemEventEdited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectItemEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectItemEventClosed(t *testing.T) {
	type args struct {
		callbacks []ProjectV2ItemEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectV2ItemEventHandleFunc",
			args: args{
				callbacks: []ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectV2ItemEventHandleFunc",
			args: args{
				callbacks: []ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectItemEventClosed(tt.args.callbacks...)
			if len(g.onProjectV2ItemEvent[ProjectItemEventClosedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectV2ItemEvent[ProjectItemEventClosedAction]))
			}
		})
	}
}

func TestSetOnProjectItemEventClosed(t *testing.T) {
	type args struct {
		callbacks []ProjectV2ItemEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectV2ItemEventHandleFunc",
			args: args{
				[]ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectV2ItemEventHandleFuncs",
			args: args{
				[]ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnProjectItemEventClosed(func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
				return nil
			})
			g.SetOnProjectItemEventClosed(tt.args.callbacks...)
			if len(g.onProjectV2ItemEvent[ProjectItemEventClosedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectV2ItemEvent[ProjectItemEventClosedAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectItemEventClosed(t *testing.T) {
	action := ProjectItemEventClosedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectV2ItemEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectItemEventClosed(func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectItemEventClosed(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectItemEventClosed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectItemEventReopened(t *testing.T) {
	type args struct {
		callbacks []ProjectV2ItemEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectV2ItemEventHandleFunc",
			args: args{
				callbacks: []ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectV2ItemEventHandleFunc",
			args: args{
				callbacks: []ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectItemEventReopened(tt.args.callbacks...)
			if len(g.onProjectV2ItemEvent[ProjectItemEventReopenedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectV2ItemEvent[ProjectItemEventReopenedAction]))
			}
		})
	}
}

func TestSetOnProjectItemEventReopened(t *testing.T) {
	type args struct {
		callbacks []ProjectV2ItemEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectV2ItemEventHandleFunc",
			args: args{
				[]ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectV2ItemEventHandleFuncs",
			args: args{
				[]ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnProjectItemEventReopened(func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
				return nil
			})
			g.SetOnProjectItemEventReopened(tt.args.callbacks...)
			if len(g.onProjectV2ItemEvent[ProjectItemEventReopenedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectV2ItemEvent[ProjectItemEventReopenedAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectItemEventReopened(t *testing.T) {
	action := ProjectItemEventReopenedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectV2ItemEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectItemEventReopened(func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectItemEventReopened(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectItemEventReopened() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectItemEventDeleted(t *testing.T) {
	type args struct {
		callbacks []ProjectV2ItemEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectV2ItemEventHandleFunc",
			args: args{
				callbacks: []ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectV2ItemEventHandleFunc",
			args: args{
				callbacks: []ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectItemEventDeleted(tt.args.callbacks...)
			if len(g.onProjectV2ItemEvent[ProjectItemEventDeletedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectV2ItemEvent[ProjectItemEventDeletedAction]))
			}
		})
	}
}

func TestSetOnProjectItemEventDeleted(t *testing.T) {
	type args struct {
		callbacks []ProjectV2ItemEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectV2ItemEventHandleFunc",
			args: args{
				[]ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectV2ItemEventHandleFuncs",
			args: args{
				[]ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnProjectItemEventDeleted(func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
				return nil
			})
			g.SetOnProjectItemEventDeleted(tt.args.callbacks...)
			if len(g.onProjectV2ItemEvent[ProjectItemEventDeletedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectV2ItemEvent[ProjectItemEventDeletedAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectItemEventDeleted(t *testing.T) {
	action := ProjectItemEventDeletedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectV2ItemEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectItemEventDeleted(func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectItemEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectItemEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectItemEventConverted(t *testing.T) {
	type args struct {
		callbacks []ProjectV2ItemEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectV2ItemEventHandleFunc",
			args: args{
				callbacks: []ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectV2ItemEventHandleFunc",
			args: args{
				callbacks: []ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectItemEventConverted(tt.args.callbacks...)
			if len(g.onProjectV2ItemEvent[ProjectItemEventConvertedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectV2ItemEvent[ProjectItemEventConvertedAction]))
			}
		})
	}
}

func TestSetOnProjectItemEventConverted(t *testing.T) {
	type args struct {
		callbacks []ProjectV2ItemEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectV2ItemEventHandleFunc",
			args: args{
				[]ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectV2ItemEventHandleFuncs",
			args: args{
				[]ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnProjectItemEventConverted(func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
				return nil
			})
			g.SetOnProjectItemEventConverted(tt.args.callbacks...)
			if len(g.onProjectV2ItemEvent[ProjectItemEventConvertedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectV2ItemEvent[ProjectItemEventConvertedAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectItemEventConverted(t *testing.T) {
	action := ProjectItemEventConvertedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectV2ItemEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectItemEventConverted(func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectItemEventConverted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectItemEventConverted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectItemEventRestored(t *testing.T) {
	type args struct {
		callbacks []ProjectV2ItemEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectV2ItemEventHandleFunc",
			args: args{
				callbacks: []ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectV2ItemEventHandleFunc",
			args: args{
				callbacks: []ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectItemEventRestored(tt.args.callbacks...)
			if len(g.onProjectV2ItemEvent[ProjectItemEventRestoredAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectV2ItemEvent[ProjectItemEventRestoredAction]))
			}
		})
	}
}

func TestSetOnProjectItemEventRestored(t *testing.T) {
	type args struct {
		callbacks []ProjectV2ItemEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectV2ItemEventHandleFunc",
			args: args{
				[]ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectV2ItemEventHandleFuncs",
			args: args{
				[]ProjectV2ItemEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnProjectItemEventRestored(func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
				return nil
			})
			g.SetOnProjectItemEventRestored(tt.args.callbacks...)
			if len(g.onProjectV2ItemEvent[ProjectItemEventRestoredAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectV2ItemEvent[ProjectItemEventRestoredAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectItemEventRestored(t *testing.T) {
	action := ProjectItemEventRestoredAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectV2ItemEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectItemEventRestored(func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectItemEventRestored(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectItemEventRestored() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProjectV2ItemEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectV2ItemEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger ProjectV2ItemEventAny with unknown event action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2ItemEvent: map[string][]ProjectV2ItemEventHandleFunc{
						ProjectV2ItemEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  ProjectV2ItemEvent,

				event: &github.ProjectV2ItemEvent{Action: ptrString("unknown")},
			},
			wantErr: false,
		},

		{
			name: "must trigger ProjectItemEventCreated",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2ItemEvent: map[string][]ProjectV2ItemEventHandleFunc{
						ProjectV2ItemEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectItemEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Logf("%s action called", ProjectItemEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: ptrString(ProjectItemEventCreatedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail ProjectItemEventCreated with empty action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2ItemEvent: map[string][]ProjectV2ItemEventHandleFunc{
						ProjectV2ItemEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectItemEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Logf("%s action called", ProjectItemEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail ProjectItemEventCreated with nil action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2ItemEvent: map[string][]ProjectV2ItemEventHandleFunc{
						ProjectV2ItemEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectItemEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Logf("%s action called", ProjectItemEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger ProjectItemEventEdited",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2ItemEvent: map[string][]ProjectV2ItemEventHandleFunc{
						ProjectV2ItemEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectItemEventEditedAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Logf("%s action called", ProjectItemEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: ptrString(ProjectItemEventEditedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail ProjectItemEventEdited with empty action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2ItemEvent: map[string][]ProjectV2ItemEventHandleFunc{
						ProjectV2ItemEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectItemEventEditedAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Logf("%s action called", ProjectItemEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail ProjectItemEventEdited with nil action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2ItemEvent: map[string][]ProjectV2ItemEventHandleFunc{
						ProjectV2ItemEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectItemEventEditedAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Logf("%s action called", ProjectItemEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger ProjectItemEventClosed",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2ItemEvent: map[string][]ProjectV2ItemEventHandleFunc{
						ProjectV2ItemEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectItemEventClosedAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Logf("%s action called", ProjectItemEventClosedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: ptrString(ProjectItemEventClosedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail ProjectItemEventClosed with empty action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2ItemEvent: map[string][]ProjectV2ItemEventHandleFunc{
						ProjectV2ItemEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectItemEventClosedAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Logf("%s action called", ProjectItemEventClosedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail ProjectItemEventClosed with nil action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2ItemEvent: map[string][]ProjectV2ItemEventHandleFunc{
						ProjectV2ItemEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectItemEventClosedAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Logf("%s action called", ProjectItemEventClosedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger ProjectItemEventReopened",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2ItemEvent: map[string][]ProjectV2ItemEventHandleFunc{
						ProjectV2ItemEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectItemEventReopenedAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Logf("%s action called", ProjectItemEventReopenedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: ptrString(ProjectItemEventReopenedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail ProjectItemEventReopened with empty action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2ItemEvent: map[string][]ProjectV2ItemEventHandleFunc{
						ProjectV2ItemEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectItemEventReopenedAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Logf("%s action called", ProjectItemEventReopenedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail ProjectItemEventReopened with nil action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2ItemEvent: map[string][]ProjectV2ItemEventHandleFunc{
						ProjectV2ItemEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectItemEventReopenedAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Logf("%s action called", ProjectItemEventReopenedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger ProjectItemEventDeleted",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2ItemEvent: map[string][]ProjectV2ItemEventHandleFunc{
						ProjectV2ItemEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectItemEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Logf("%s action called", ProjectItemEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: ptrString(ProjectItemEventDeletedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail ProjectItemEventDeleted with empty action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2ItemEvent: map[string][]ProjectV2ItemEventHandleFunc{
						ProjectV2ItemEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectItemEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Logf("%s action called", ProjectItemEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail ProjectItemEventDeleted with nil action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2ItemEvent: map[string][]ProjectV2ItemEventHandleFunc{
						ProjectV2ItemEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectItemEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Logf("%s action called", ProjectItemEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger ProjectItemEventConverted",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2ItemEvent: map[string][]ProjectV2ItemEventHandleFunc{
						ProjectV2ItemEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectItemEventConvertedAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Logf("%s action called", ProjectItemEventConvertedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: ptrString(ProjectItemEventConvertedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail ProjectItemEventConverted with empty action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2ItemEvent: map[string][]ProjectV2ItemEventHandleFunc{
						ProjectV2ItemEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectItemEventConvertedAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Logf("%s action called", ProjectItemEventConvertedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail ProjectItemEventConverted with nil action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2ItemEvent: map[string][]ProjectV2ItemEventHandleFunc{
						ProjectV2ItemEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectItemEventConvertedAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Logf("%s action called", ProjectItemEventConvertedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger ProjectItemEventRestored",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2ItemEvent: map[string][]ProjectV2ItemEventHandleFunc{
						ProjectV2ItemEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectItemEventRestoredAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Logf("%s action called", ProjectItemEventRestoredAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: ptrString(ProjectItemEventRestoredAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail ProjectItemEventRestored with empty action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2ItemEvent: map[string][]ProjectV2ItemEventHandleFunc{
						ProjectV2ItemEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectItemEventRestoredAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Logf("%s action called", ProjectItemEventRestoredAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail ProjectItemEventRestored with nil action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2ItemEvent: map[string][]ProjectV2ItemEventHandleFunc{
						ProjectV2ItemEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectItemEventRestoredAction: {
							func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
								t.Logf("%s action called", ProjectItemEventRestoredAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2_item",
				event:      &github.ProjectV2ItemEvent{Action: nil},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &EventHandler{
				WebhookSecret: "fake",
				mu:            sync.RWMutex{},
			}
			if err := g.ProjectV2ItemEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("ProjectV2ItemEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
