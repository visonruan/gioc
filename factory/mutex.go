// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package factory

import (
	"github.com/vlorc/gioc/types"
)

func (mf *MutexFactory) Instance(provider types.Provider) (instance interface{}, err error) {
	mf.m.Lock()
	defer mf.m.Unlock()

	instance, err = mf.factory.Instance(provider)
	return
}
