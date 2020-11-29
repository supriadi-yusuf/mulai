package simhelper

import (
	"errors"
	"reflect"
	"sync"
)

// MergeChanels is function to merge several receiver channels into one receiver channel.
func MergeChanels(chanels ...interface{}) (chn interface{}, err error) {

	defer GetErrorOnPanic(&err)

	totChan := len(chanels) //total channel element that we want to join
	if totChan == 0 {
		return nil, errors.New("number of joined channel may not be zero")
	}

	refVal0 := reflect.ValueOf(chanels[0])

	if refVal0.Kind() != reflect.Chan {
		return nil, errors.New("input parameter must be channel")
	}

	elType := refVal0.Type().Elem()                      //type of channel
	chType := reflect.ChanOf(reflect.BothDir, elType)    //channel type
	chRcvType := reflect.ChanOf(reflect.RecvDir, elType) //channel type : receiver channel

	outChan := reflect.MakeChan(chType, 0)      // create channel
	outRcvChan := reflect.New(chRcvType).Elem() //create empty receiver channel
	outRcvChan.Set(outChan)                     //convert channel into receiver channel

	wg := new(sync.WaitGroup)
	wg.Add(totChan)

	for _, chanel := range chanels {

		go func(ch reflect.Value) {

			for {

				val, ok := ch.Recv() //receive value from input channel
				if ok {
					outChan.Send(val) //send value into output channel
				} else {
					// chanel is closed
					break
				}

			}

			wg.Done()
		}(reflect.ValueOf(chanel))

	}

	go func() {
		wg.Wait()
		outChan.Close() //close channel
	}()

	return outRcvChan.Interface(), nil
}
