// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package pipe

// ChannelWriter passes each object to the callback function
type ChannelWriter struct {
	channel chan interface{}
}

func NewChannelWriter(channel chan interface{}) *ChannelWriter {
	return &ChannelWriter{
		channel: channel,
	}
}

func (cw *ChannelWriter) WriteObject(object interface{}) error {
	cw.channel <- object
	return nil
}

func (cw *ChannelWriter) Flush() error {
	return nil
}
