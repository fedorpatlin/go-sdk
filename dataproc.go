// Copyright (c) 2019 YANDEX LLC.

package ycsdk

import "github.com/fedorpatlin/go-sdk/gen/dataproc"

const DataProcServiceID = "dataproc"

func (sdk *SDK) Dataproc() *dataproc.Dataproc {
	return dataproc.NewDataproc(sdk.getConn(DataProcServiceID))
}
