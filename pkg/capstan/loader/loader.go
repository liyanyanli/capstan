/*
Copyright (c) 2018 The ZJU-SEL Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package loader

import (
	"fmt"

	"github.com/golang/glog"

	"github.com/ZJU-SEL/capstan/pkg/workload"
	"github.com/ZJU-SEL/capstan/pkg/workload/iperf"
	"github.com/ZJU-SEL/capstan/pkg/workload/nginx"
)

// LoadAllWorkloads loads all workloads by parsing workloads section config,
// return all of workloads which are supported in the capstan.
func LoadAllWorkloads(workloads []workload.Workload) (ret []workload.Interface, err error) {
	for _, wl := range workloads {
		find := false
		for _, wlDef := range workload.DefWorkloads {
			if wl.Name == wlDef {
				w, err := loadWorkload(wl)
				if err != nil {
					glog.Errorf("Failed load the testing workload %q :%v", wl.Name, err)
					return ret, err
				}
				ret = append(ret, w)
				find = true
			}
		}
		if !find {
			return ret, fmt.Errorf("The testing workload %q has not defined in capstan", wl.Name)
		}
	}
	return ret, nil
}

func loadWorkload(wl workload.Workload) (workload.Interface, error) {
	glog.V(1).Infof("Load a testing workload with config:%v", wl)
	switch wl.Name {
	case "nginx":
		return nginx.NewWorkload(wl)
	case "iperf":
		return iperf.NewWorkload(wl)
	default:
		return nil, fmt.Errorf("Unknown workload %v", wl.Name)
	}
}
