package utils

import "github.com/cloudwego/kitex/pkg/klog"

func MustHandleErr(err error) {
	if err != nil {
		klog.Error(err)
	}
}
