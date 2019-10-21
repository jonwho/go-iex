package goiextest

// import (
//   "testing"
// )
//
// func TestSymbols(t *testing.T) {
//   symbols, err := iexSandboxClient.Symbols()
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(symbols) == 0
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestIEXSymbols(t *testing.T) {
//   symbols, err := iexSandboxClient.IEXSymbols()
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(symbols) == 0
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestRegionSymbols(t *testing.T) {
//   symbols, err := iexSandboxClient.RegionSymbols("ca")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(symbols) == 0
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestExchangeSymbols(t *testing.T) {
//   symbols, err := iexSandboxClient.ExchangeSymbols("tse")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(symbols) == 0
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestInternationalExchanges(t *testing.T) {
//   exchanges, err := iexSandboxClient.InternationalExchanges()
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(exchanges) == 0
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestUSExchanges(t *testing.T) {
//   exchanges, err := iexSandboxClient.USExchanges()
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(exchanges) == 0
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestUSHolidaysAndTradingDates(t *testing.T) {
//   dates, err := iexSandboxClient.USHolidaysAndTradingDates("trade", "next")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(dates) == 0
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
//
//   dates, err = iexSandboxClient.USHolidaysAndTradingDates("trade", "last")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(dates) == 0
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
//
//   dates, err = iexSandboxClient.USHolidaysAndTradingDates("holiday", "next")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(dates) == 0
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
//
//   dates, err = iexSandboxClient.USHolidaysAndTradingDates("holiday", "last")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(dates) == 0
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
//
//   dates, err = iexSandboxClient.USHolidaysAndTradingDates("trade", "next", 1)
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(dates) == 0
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
//
//   dates, err = iexSandboxClient.USHolidaysAndTradingDates("holiday", "last", 1, "20190101")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(dates) == 0
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestSectors(t *testing.T) {
//   sectors, err := iexSandboxClient.Sectors()
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(sectors) == 0
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestTags(t *testing.T) {
//   tags, err := iexSandboxClient.Tags()
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(tags) == 0
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestMutualFundSymbols(t *testing.T) {
//   funds, err := iexSandboxClient.MutualFundSymbols()
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(funds) == 0
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestOTCSymbols(t *testing.T) {
//   otc, err := iexSandboxClient.OTCSymbols()
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(otc) == 0
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// // TODO: @mock
// func TestFXSymbols(t *testing.T) {
//   fx, err := iexSandboxClient.FXSymbols()
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(fx.Currencies) == 0
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// // TODO: @mock
// // func TestOptionsSymbols(t *testing.T) {
// //   os, err := iexSandboxClient.OptionsSymbols()
// //   if err != nil {
// //     t.Error(err)
// //   }
// // }
