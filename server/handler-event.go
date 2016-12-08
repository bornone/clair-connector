/*
# IBM Confidential
# OCO Source Materials
# cfc
# @ Copyright IBM Corp. 2016 All Rights Reserved
# The source code for this program is not published or otherwise divested of its trade secrets, irrespective of what has been deposited with the U.S. Copyright Office.
*/

package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
	pretty "github.com/tonnerre/golang-pretty"

	"github.com/bornone/clair-connector/clair"
	"github.com/bornone/clair-connector/common"
	"github.com/bornone/clair-connector/registry"
	"github.com/docker/distribution/notifications"
	"golang.org/x/net/context"
)

var priorities = []string{"Defcon1", "Critical", "High", "Medium", "Low", "Negligible", "Unknown"}
var store = make(map[string][]clair.Vulnerability)

func event(ctx context.Context, w http.ResponseWriter, r *http.Request) *HTTPError {
	all_event := notifications.Envelope{}
	json_decoder := json.NewDecoder(r.Body)
	err := json_decoder.Decode(&all_event)
	if err != nil {
		common.LOG(log.ErrorLevel, "Decode Registry Event Error: {0}", err)
	}

	for _, event := range all_event.Events {
		if event.Action == "push" {
			// clean cache
			store = make(map[string][]clair.Vulnerability, 0)
			var repo, tag, url, summary string
			repo = event.Target.Repository
			tag = event.Target.Tag
			url = event.Target.URL
			if strings.Contains(url, "manifests") {
				common.LOG(log.DebugLevel, "Receive image {0} push completed notication.", fmt.Sprintf("%# v", pretty.Formatter(event)))
				image, err := registry.GetImage(strings.Split(url, "v2")[0]+"v2", repo, tag)
				if err != nil {
					summary = fmt.Sprintf("Failed to get image %s:%s from registry due to: %s", repo, tag, err)
					common.LOG(log.ErrorLevel, summary)
				} else {
					common.LOG(log.DebugLevel, "Get image {0} info", fmt.Sprintf("%# v", pretty.Formatter(image)))
					c := clair.NewClair(ctx.Value("clair-url").(string))
					vs := c.Analyse(image)
					var vs2 = make([]clair.Vulnerability, 0)
					if len(vs) == 0 {
						summary = fmt.Sprintf("No vulnerability found in image %s", repo)
					} else {
						groupBySeverity(vs)
						iteratePriorities(func(sev string) {
							summary += fmt.Sprintf(" %d %s |", len(store[sev]), sev)
							for _, v := range store[sev] {
								vs2 = append(vs2, v)
							}
						})
					}
					report := &clair.Report{Summary: summary, Vulnerabilities: vs2}
					common.LOG(log.DebugLevel, "Get image {0} info", fmt.Sprintf("%# v", pretty.Formatter(report)))
				}
			}
		}
	}
	return nil
}

func iteratePriorities(f func(sev string)) {
	for _, sev := range priorities {
		if len(store[sev]) != 0 {
			f(sev)
		}
	}
}

func groupBySeverity(vs []clair.Vulnerability) {
	for _, v := range vs {
		sevRow := vulnsBy(v.Severity, store)
		store[v.Severity] = append(sevRow, v)
	}
}

func vulnsBy(sev string, store map[string][]clair.Vulnerability) []clair.Vulnerability {
	items, found := store[sev]
	if !found {
		items = make([]clair.Vulnerability, 0)
		store[sev] = items
	}
	return items
}
