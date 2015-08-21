package render_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/weaveworks/scope/render"
	"github.com/weaveworks/scope/test"
)

func TestOriginTable(t *testing.T) {
	if _, ok := render.OriginTable(test.Report, "not-found"); ok {
		t.Errorf("unknown origin ID gave unexpected success")
	}
	for originID, want := range map[string]render.Table{
		test.ServerProcessNodeID: {
			Title:   "Origin Process",
			Numeric: false,
			Rank:    2,
			Rows: []render.Row{
				{"Name", "apache", "", false},
				{"PID", test.ServerPID, "", false},
			},
		},
		test.ServerHostNodeID: {
			Title:   "Origin Host",
			Numeric: false,
			Rank:    1,
			Rows: []render.Row{
				{"Host name", test.ServerHostName, "", false},
				{"Load", "0.01 0.01 0.01", "", false},
				{"Operating system", "Linux", "", false},
			},
		},
	} {
		have, ok := render.OriginTable(test.Report, originID)
		if !ok {
			t.Errorf("%q: not OK", originID)
			continue
		}
		if !reflect.DeepEqual(want, have) {
			t.Errorf("%q: %s", originID, test.Diff(want, have))
		}
	}
}

func TestMakeDetailedHostNode(t *testing.T) {
	renderableNode := render.HostRenderer.Render(test.Report)[render.MakeHostID(test.ClientHostID)]
	have := render.MakeDetailedNode(test.Report, renderableNode)
	want := render.DetailedNode{
		ID:         render.MakeHostID(test.ClientHostID),
		LabelMajor: "client",
		LabelMinor: "hostname.com",
		Pseudo:     false,
		Tables: []render.Table{
			{
				Title:   "Origin Host",
				Numeric: false,
				Rank:    1,
				Rows: []render.Row{
					{
						Key:        "Host name",
						ValueMajor: "client.hostname.com",
						ValueMinor: "",
					},
					{
						Key:        "Load",
						ValueMajor: "0.01 0.01 0.01",
						ValueMinor: "",
					},
					{
						Key:        "Operating system",
						ValueMajor: "Linux",
						ValueMinor: "",
					},
				},
			},
			{
				Title:   "Connections",
				Numeric: true,
				Rank:    0,
				Rows: []render.Row{
					{
						Key:        "TCP connections",
						ValueMajor: "3",
						ValueMinor: "",
					},
					{
						Key:        "Client",
						ValueMajor: "Server",
						ValueMinor: "",
						Expandable: true,
					},
					{
						Key:        "10.10.10.20",
						ValueMajor: "192.168.1.1",
						ValueMinor: "",
						Expandable: true,
					},
				},
			},
		},
	}
	if !reflect.DeepEqual(want, have) {
		t.Errorf("%s", test.Diff(want, have))
	}
}

func TestMakeDetailedContainerNode(t *testing.T) {
	renderableNode := render.ContainerRenderer.Render(test.Report)[test.ServerContainerID]
	have := render.MakeDetailedNode(test.Report, renderableNode)
	want := render.DetailedNode{
		ID:         test.ServerContainerID,
		LabelMajor: "server",
		LabelMinor: test.ServerHostName,
		Pseudo:     false,
		Tables: []render.Table{
			{
				Title:   "Origin Container",
				Numeric: false,
				Rank:    3,
				Rows: []render.Row{
					{"ID", test.ServerContainerID, "", false},
					{"Name", "server", "", false},
					{"Image ID", test.ServerContainerImageID, "", false},
				},
			},
			{
				Title:   "Origin Process",
				Numeric: false,
				Rank:    2,
				Rows: []render.Row{
					{"Name", "apache", "", false},
					{"PID", test.ServerPID, "", false},
				},
			},
			{
				Title:   "Origin Host",
				Numeric: false,
				Rank:    1,
				Rows: []render.Row{
					{"Host name", test.ServerHostName, "", false},
					{"Load", "0.01 0.01 0.01", "", false},
					{"Operating system", "Linux", "", false},
				},
			},
			{
				Title:   "Connections",
				Numeric: true,
				Rank:    0,
				Rows: []render.Row{
					{"Egress packet rate", "105", "packets/sec", false},
					{"Egress byte rate", "1.0", "KBps", false},
					{"Client", "Server", "", true},
					{
						fmt.Sprintf("%s:%s", test.UnknownClient1IP, test.ClientPort54010),
						fmt.Sprintf("%s:%s", test.ServerIP, test.ServerPort),
						"",
						true,
					},
					{
						fmt.Sprintf("%s:%s", test.UnknownClient1IP, test.ClientPort54020),
						fmt.Sprintf("%s:%s", test.ServerIP, test.ServerPort),
						"",
						true,
					},
					{
						fmt.Sprintf("%s:%s", test.UnknownClient3IP, test.ClientPort54020),
						fmt.Sprintf("%s:%s", test.ServerIP, test.ServerPort),
						"",
						true,
					},
					{
						fmt.Sprintf("%s:%s", test.ClientIP, test.ClientPort54001),
						fmt.Sprintf("%s:%s", test.ServerIP, test.ServerPort),
						"",
						true,
					},
					{
						fmt.Sprintf("%s:%s", test.ClientIP, test.ClientPort54002),
						fmt.Sprintf("%s:%s", test.ServerIP, test.ServerPort),
						"",
						true,
					},
					{
						fmt.Sprintf("%s:%s", test.RandomClientIP, test.ClientPort12345),
						fmt.Sprintf("%s:%s", test.ServerIP, test.ServerPort),
						"",
						true,
					},
				},
			},
		},
	}
	if !reflect.DeepEqual(want, have) {
		t.Errorf("%s", test.Diff(want, have))
	}
}