package goiextest

// IEX has deprecated AlternativeData
// TODO: add implementation for new endpoints

// import (
//   "testing"
// )
//
// func TestCrypto(t *testing.T) {
//   crypto, err := iexSandboxClient.Crypto("btcusdt")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = "BTCUSDT"
//   actual = crypto.Symbol
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// // TODO: @mock
// func TestSocialSentiment(t *testing.T) {
//   ss, err := iexSandboxClient.SocialSentiment("aapl")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = true
//   actual = ss.Sentiment > 0
//   if !actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// // TODO: @mock
// func TestCEOCompensation(t *testing.T) {
//   cc, err := iexSandboxClient.CEOCompensation("aapl")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = "AAPL"
//   actual = cc.Symbol
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
