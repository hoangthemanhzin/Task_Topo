package topo

import (
	"etrib5gc/sbi/models"
	"strings"
	"testing"
)

type TestCase struct {
	dnn  string
	an   string
	path []string
}

func (tc *TestCase) pass(path DataPath) bool {
	if len(path) != len(tc.path) {
		return false
	}
	for i, n := range path {
		if strings.Compare(n.id, tc.path[i]) != 0 {
			return false
		}
	}
	return true
}

func Test_Topo(t *testing.T) {
	if upftopo, err := Load("topo.json"); err != nil {
		t.Errorf("Fail to parse config file: %s", err.Error())
	} else {
		query := &PathQuery{
			Snssai: models.Snssai{
				Sd:  "12345",
				Sst: 0,
			},
		}
		cases := []TestCase{
			TestCase{
				dnn:  "internet",
				an:   "an1",
				path: []string{"upf1", "upf3"},
			},
			TestCase{
				dnn:  "internet",
				an:   "an2",
				path: []string{"upf2", "upf3"},
			},
			TestCase{
				dnn:  "e1",
				an:   "an1",
				path: []string{"upf1"},
			},
			TestCase{
				dnn:  "e2",
				an:   "an2",
				path: []string{"upf2"},
			},
			TestCase{
				dnn:  "e1",
				an:   "an2",
				path: []string{"upf2", "upf3", "upf1"},
			},
			TestCase{
				dnn:  "e2",
				an:   "an1",
				path: []string{"upf1", "upf3", "upf2"},
			},
		}
		for i, tc := range cases {
			query.Dnn = tc.dnn
			query.Nets = []string{tc.an}
			if path, ip := upftopo.FindPath(query); ip == nil || len(path) == 0 {
				t.Errorf("test case %d: fail to find path", i)
			} else if !tc.pass(path) {
				t.Errorf("test case %d: found path %s, expect %v", i, path, tc.path)
			} else {
				log.Infof("test case %d: found path %s, expect %v", i, path, tc.path)
			}
		}
	}
}
