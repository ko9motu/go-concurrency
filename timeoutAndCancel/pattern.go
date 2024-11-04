package main

// func main() {

// 	go func() {
// 		var value interface{}
// 		select {
// 		case <-done:
// 			return
// 		case value = <-valueCh:
// 		}
// 		result := reallyLongCalc(value)
// 		select {
// 		case <-done:
// 			return

// 		case resultCh <- result:
// 		}
// 	}()

// }

// func reallyLongCalc(done <-chan interface{}, value interface{}) interface{} {
// 	intermediateResult := longCalc(done, value)
// 	select {
// 	case <-done:
// 		return nil
// 	default:
// 	}
// 	return longCalc(done,intermediateResult)
// }
