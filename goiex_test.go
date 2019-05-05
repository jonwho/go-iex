package goiex

import (
// "testing"
//
// "github.com/jonwho/go-iex/mock-iex"
)

// // TODO: use during hours reponse stub
// func TestEarningsToday(t *testing.T) {
//   mockServer := mockiex.Server()
//   defer mockServer.Close()
//   client, err := NewClient(SetBaseURL(mockServer.URL))
//   if err != nil {
//     t.Error(err)
//   }
//
//   earningsToday, _ := client.EarningsToday()
//
//   if len(earningsToday.BTO) != 0 {
//     t.Errorf("expected 0 but got %v", earningsToday.BTO)
//   }
//
//   if len(earningsToday.AMC) != 0 {
//     t.Errorf("expected 0 but got %v", earningsToday.AMC)
//   }
// }
//
// func TestEarnings(t *testing.T) {
//   mockServer := mockiex.Server()
//   defer mockServer.Close()
//   client, err := NewClient(SetBaseURL(mockServer.URL))
//   if err != nil {
//     t.Error(err)
//   }
//
//   earnings, _ := client.Earnings("aapl")
//
//   if earnings.Symbol != "AAPL" {
//     t.Errorf("expected AAPL but got %v", earnings.Symbol)
//   }
//
//   if earnings.Earnings[0].SymbolId != 11 {
//     t.Errorf("expected 11 but got %v", earnings.Earnings[0].SymbolId)
//   }
// }
//
// func TestQuote(t *testing.T) {
//   mockServer := mockiex.Server()
//   defer mockServer.Close()
//   client, err := NewClient(SetBaseURL(mockServer.URL))
//   if err != nil {
//     t.Error(err)
//   }
//
//   quote, _ := client.Quote("aapl", false)
//
//   if quote.Symbol != "AAPL" {
//     t.Errorf("expected AAPL but got %v", quote.Symbol)
//   }
//
//   if quote.CompanyName != "Apple Inc." {
//     t.Errorf("expected Apple Inc. but got %v", quote.CompanyName)
//   }
//
//   if quote.ChangePercent != -0.01592 {
//     t.Errorf("expected -0.01592 but got %v", quote.ChangePercent)
//   }
//
//   _, err = client.Quote("fakesymbol", false)
//
//   if err == nil {
//     t.Error("expected err but got nil")
//   }
//
//   quote, _ = client.Quote("aapl", true)
//
//   if quote.ChangePercent != -0.01592*100 {
//     t.Errorf("expected -1.592 but got %v", quote.ChangePercent)
//   }
// }
//
// func TestChart(t *testing.T) {
//   mockServer := mockiex.Server()
//   defer mockServer.Close()
//   client, err := NewClient(SetBaseURL(mockServer.URL))
//   if err != nil {
//     t.Error(err)
//   }
//
//   _, err = client.Chart("aapl", "6y")
//
//   if err == nil {
//     t.Error("expected err but got nil")
//   }
//
//   chart, _ := client.Chart("aapl", "1d")
//
//   if len(chart.Charts) == 0 {
//     t.Error("charts shouldn't be empty")
//   }
//
//   if chart.Charts[0].Minute == "" {
//     t.Error("minute should be non-empty string for 1d range")
//   }
// }
//
// func TestRefDataSymbols(t *testing.T) {
//   mockServer := mockiex.Server()
//   defer mockServer.Close()
//   client, err := NewClient(SetBaseURL(mockServer.URL))
//   if err != nil {
//     t.Error(err)
//   }
//
//   rds, _ := client.RefDataSymbols()
//
//   firstSymbol := rds.Symbols[0]
//
//   if firstSymbol.Symbol != "A" {
//     t.Errorf("expected A but got %v", firstSymbol.Symbol)
//   }
//
//   if firstSymbol.Date != "2018-10-26" {
//     t.Errorf("expected 2018-10-26 but got %v", firstSymbol.Date)
//   }
//
//   if firstSymbol.Name != "Agilent Technologies Inc." {
//     t.Errorf("expected Agilent Technologies Inc. but got %v", firstSymbol.Name)
//   }
//
//   if firstSymbol.IsEnabled != true {
//     t.Errorf("expected true but got %v", firstSymbol.IsEnabled)
//   }
//
//   if firstSymbol.Type != "cs" {
//     t.Errorf("expected cs but got %v", firstSymbol.Type)
//   }
//
//   if firstSymbol.IexId != 2 {
//     t.Errorf("expected 2 but got %v", firstSymbol.IexId)
//   }
// }
//
// func TestRefDataCorporateActions(t *testing.T) {
//   mockServer := mockiex.Server()
//   defer mockServer.Close()
//   client, err := NewClient(SetBaseURL(mockServer.URL))
//   if err != nil {
//     t.Error(err)
//   }
//
//   // without date
//   rdca, err := client.RefDataCorporateActions("")
//   if err != nil {
//     t.Errorf("Expected nil but got %s\n", err.Error())
//   }
//   ca := rdca.CorporateActions[0]
//   if ca.RecordID != "CA20171108153808144" {
//     t.Errorf("Expected CA20171108153808144 but got %s", ca.RecordID)
//   }
//
//   // with date - TODO: need to run during day to get value
//   // rdca, err = client.RefDataCorporateActions("20171210")
//   // if err != nil {
//   //   t.Errorf("Expected nil but got %s\n", err.Error())
//   // }
//   // ca = rdca.CorporateActions[0]
//   // if ca.RecordID != "<updatehere>" {
//   //   t.Errorf("Expected <updatehere> but got %s", ca.RecordID)
//   // }
//
//   // with sample
//   rdca, err = client.RefDataCorporateActions("sample")
//   if err != nil {
//     t.Errorf("Expected nil but got %s\n", err.Error())
//   }
//   ca = rdca.CorporateActions[0]
//   if ca.RecordID != "CA20171108153808144" {
//     t.Errorf("Expected CA20171108153808144 but got %s", ca.RecordID)
//   }
// }
//
// func TestRefDataDividends(t *testing.T) {
//   mockServer := mockiex.Server()
//   defer mockServer.Close()
//   client, err := NewClient(SetBaseURL(mockServer.URL))
//   if err != nil {
//     t.Error(err)
//   }
//
//   // without date
//   divs, err := client.RefDataDividends("")
//   if err != nil {
//     t.Errorf("Expected nil but got %s\n", err.Error())
//   }
//   div := divs.Dividends[0]
//   if div.RecordID != "DV20171108154436478" {
//     t.Errorf("Expected DV20171108154436478 but got %s", div.RecordID)
//   }
//
//   // with date - TODO: need to run during day to get value
//   // divs, err = client.Dividends("20171210")
//   // if err != nil {
//   //   t.Errorf("Expected nil but got %s\n", err.Error())
//   // }
//   // div = divs.Dividends[0]
//   // if div.RecordID != "<updatehere>" {
//   //   t.Errorf("Expected <updatehere> but got %s", div.RecordID)
//   // }
//
//   // with sample
//   divs, err = client.RefDataDividends("sample")
//   if err != nil {
//     t.Errorf("Expected nil but got %s\n", err.Error())
//   }
//   div = divs.Dividends[0]
//   if div.RecordID != "DV20171108154436478" {
//     t.Errorf("Expected DV20171108154436478 but got %s", div.RecordID)
//   }
// }
//
// func TestRefDataNextDayExDates(t *testing.T) {
//   mockServer := mockiex.Server()
//   defer mockServer.Close()
//   client, err := NewClient(SetBaseURL(mockServer.URL))
//   if err != nil {
//     t.Error(err)
//   }
//
//   // without date
//   rdnded, err := client.RefDataNextDayExDates("")
//   if err != nil {
//     t.Errorf("Expected nil but got %s\n", err.Error())
//   }
//   nded := rdnded.NextDayExDates[0]
//   if nded.RecordID != "DV20171108154436478" {
//     t.Errorf("Expected DV20171108154436478 but got %s", nded.RecordID)
//   }
//
//   // with date - TODO: need to run during day to get value
//   // rdnded, err = client.Dividends("20171210")
//   // if err != nil {
//   //   t.Errorf("Expected nil but got %s\n", err.Error())
//   // }
//   // nded = rdnded.Dividends[0]
//   // if nded.RecordID != "<updatehere>" {
//   //   t.Errorf("Expected <updatehere> but got %s", nded.RecordID)
//   // }
//
//   // with sample
//   rdnded, err = client.RefDataNextDayExDates("sample")
//   if err != nil {
//     t.Errorf("Expected nil but got %s\n", err.Error())
//   }
//   nded = rdnded.NextDayExDates[0]
//   if nded.RecordID != "DV20171108154436478" {
//     t.Errorf("Expected DV20171108154436478 but got %s", nded.RecordID)
//   }
// }
//
// func TestRefDataSymbolDirectories(t *testing.T) {
//   mockServer := mockiex.Server()
//   defer mockServer.Close()
//   client, err := NewClient(SetBaseURL(mockServer.URL))
//   if err != nil {
//     t.Error(err)
//   }
//
//   // without date
//   rdsd, err := client.RefDataSymbolDirectories("")
//   if err != nil {
//     t.Errorf("Expected nil but got %s\n", err.Error())
//   }
//   sd := rdsd.SymbolDirectories[0]
//   if sd.RecordID != "SD20171020161150890" {
//     t.Errorf("Expected SD20171020161150890 but got %s", sd.RecordID)
//   }
//
//   // with date - TODO: need to run during day to get value
//   // rdsd, err = client.SymbolDirectories("20171210")
//   // if err != nil {
//   //   t.Errorf("Expected nil but got %s\n", err.Error())
//   // }
//   // sd = rdsd.SymbolDirectories[0]
//   // if sd.RecordID != "<updatehere>" {
//   //   t.Errorf("Expected <updatehere> but got %s", sd.RecordID)
//   // }
//
//   // with sample
//   rdsd, err = client.RefDataSymbolDirectories("sample")
//   if err != nil {
//     t.Errorf("Expected nil but got %s\n", err.Error())
//   }
//   sd = rdsd.SymbolDirectories[0]
//   if sd.RecordID != "SD20171020161150890" {
//     t.Errorf("Expected SD20171020161150890 but got %s", sd.RecordID)
//   }
// }
//
// func TestKeyStat(t *testing.T) {
//   mockServer := mockiex.Server()
//   defer mockServer.Close()
//   client, err := NewClient(SetBaseURL(mockServer.URL))
//   if err != nil {
//     t.Error(err)
//   }
//
//   keyStat, err := client.KeyStat("aapl")
//
//   if keyStat.Symbol != "AAPL" {
//     t.Errorf("expected AAPL but got %v", keyStat.Symbol)
//   }
//
//   if keyStat.CompanyName != "Apple Inc." {
//     t.Errorf("expected Apple Inc. but got %v", keyStat.CompanyName)
//   }
//
//   if keyStat.Beta != 1.122636 {
//     t.Errorf("expected 1.122636 but got %v", keyStat.Beta)
//   }
//
//   _, err = client.KeyStat("fakesymbol")
//
//   if err == nil {
//     t.Error("expected err but got nil")
//   }
// }
//
// func TestBook(t *testing.T) {
//   mockServer := mockiex.Server()
//   defer mockServer.Close()
//   client, err := NewClient(SetBaseURL(mockServer.URL))
//   if err != nil {
//     t.Error(err)
//   }
//
//   book, _ := client.Book("aapl")
//
//   if book.Quote.Symbol != "AAPL" {
//     t.Errorf("expected AAPL but got %v", book.Quote.Symbol)
//   }
//
//   if book.SystemEvent.SystemEvent != "R" {
//     t.Errorf("expected R but got %v", book.SystemEvent.SystemEvent)
//   }
//
//   ask := book.Asks[0]
//   if ask.Price != 174.98 {
//     t.Errorf("expected 174.98 but got %v", ask.Price)
//   }
//   if ask.Size != 111 {
//     t.Errorf("expected 174.98 but got %v", ask.Size)
//
//   }
//   if ask.Timestamp != 1551296788138 {
//     t.Errorf("expected 174.98 but got %v", ask.Timestamp)
//   }
//
//   trade := book.Trades[0]
//   if trade.Price != 174.89 {
//     t.Errorf("expected 174.89 but got %v", trade.Price)
//   }
//   if trade.Size != 20 {
//     t.Errorf("expected 20 but got %v", trade.Size)
//   }
//   if trade.TradeID != 726951527 {
//     t.Errorf("expected 726951527 but got %v", trade.TradeID)
//   }
//   if trade.IsISO != true {
//     t.Errorf("expected true but got %v", trade.IsISO)
//   }
//   if trade.IsOddLot != true {
//     t.Errorf("expected true but got %v", trade.IsOddLot)
//   }
//   if trade.IsOutsideRegularHours != false {
//     t.Errorf("expected false but got %v", trade.IsOutsideRegularHours)
//   }
//   if trade.IsSinglePriceCross != false {
//     t.Errorf("expected false but got %v", trade.IsSinglePriceCross)
//   }
//   if trade.IsTradeThroughExempt != false {
//     t.Errorf("expected false but got %v", trade.IsTradeThroughExempt)
//   }
//   if trade.Timestamp != 1551298665713 {
//     t.Errorf("expected false but got %v", trade.Timestamp)
//   }
// }
//
// func TestNews(t *testing.T) {
//   mockServer := mockiex.Server()
//   defer mockServer.Close()
//   client, err := NewClient(SetBaseURL(mockServer.URL))
//   if err != nil {
//     t.Error(err)
//   }
//
//   news, _ := client.News("aapl", 1)
//   if len(news.News) != 1 {
//     t.Errorf("expected 1 but got %v", len(news.News))
//   }
//
//   // can request last 10 pieces of news but not guaranteed to get 10 pieces
//   news, _ = client.News("aapl", 10)
//   if len(news.News) != 4 {
//     t.Errorf("expected 4 but got %v", len(news.News))
//   }
// }
//
// func TestBatch(t *testing.T) {
//   mockServer := mockiex.Server()
//   defer mockServer.Close()
//   client, err := NewClient(SetBaseURL(mockServer.URL))
//   if err != nil {
//     t.Error(err)
//   }
//
//   params := map[string]string{"types": "quote"}
//   batch, _ := client.Batch("aapl", params)
//   if batch.Quote.Symbol != "AAPL" {
//     t.Errorf("expected AAPL but got %v", batch.Quote.Symbol)
//   }
//
//   params = map[string]string{
//     "types": "quote,news,chart",
//     "range": "1m",
//     "last":  "5",
//   }
//   batch, err = client.Batch("aapl", params)
//   if err != nil {
//     t.Error(err.Error())
//   }
//   if batch.Quote.Symbol != "AAPL" {
//     t.Errorf("expected AAPL but got %v", batch.Quote.Symbol)
//   }
//   if len(batch.News.News) != 5 {
//     t.Errorf("expected 5 but got %v", len(batch.News.News))
//   }
//   if len(batch.Chart.Charts) != 23 {
//     t.Errorf("expected 23 but got %v", len(batch.Chart.Charts))
//   }
// }
