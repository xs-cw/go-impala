// Code generated by Thrift Compiler (0.14.1). DO NOT EDIT.

package main

import (
	"context"
	"flag"
	"fmt"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/bippio/go-impala/services/hive_metastore"
	"github.com/bippio/go-impala/services/beeswax"
)

var _ = hive_metastore.GoUnusedProtection__
var _ = beeswax.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  QueryHandle query(Query query)")
  fmt.Fprintln(os.Stderr, "  QueryHandle executeAndWait(Query query, LogContextId clientCtx)")
  fmt.Fprintln(os.Stderr, "  QueryExplanation explain(Query query)")
  fmt.Fprintln(os.Stderr, "  Results fetch(QueryHandle query_id, bool start_over, i32 fetch_size)")
  fmt.Fprintln(os.Stderr, "  QueryState get_state(QueryHandle handle)")
  fmt.Fprintln(os.Stderr, "  ResultsMetadata get_results_metadata(QueryHandle handle)")
  fmt.Fprintln(os.Stderr, "  string echo(string s)")
  fmt.Fprintln(os.Stderr, "  string dump_config()")
  fmt.Fprintln(os.Stderr, "  string get_log(LogContextId context)")
  fmt.Fprintln(os.Stderr, "   get_default_configuration(bool include_hadoop)")
  fmt.Fprintln(os.Stderr, "  void close(QueryHandle handle)")
  fmt.Fprintln(os.Stderr, "  void clean(LogContextId log_context)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
  var m map[string]string = h
  return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
  parts := strings.Split(value, ": ")
  if len(parts) != 2 {
    return fmt.Errorf("header should be of format 'Key: Value'")
  }
  h[parts[0]] = parts[1]
  return nil
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  headers := make(httpHeaders)
  var parsedUrl *url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
    if len(headers) > 0 {
      httptrans := trans.(*thrift.THttpClient)
      for key, value := range headers {
        httptrans.SetHeader(key, value)
      }
    }
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  iprot := protocolFactory.GetProtocol(trans)
  oprot := protocolFactory.GetProtocol(trans)
  client := beeswax.NewBeeswaxServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "query":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Query requires 1 args")
      flag.Usage()
    }
    arg45 := flag.Arg(1)
    mbTrans46 := thrift.NewTMemoryBufferLen(len(arg45))
    defer mbTrans46.Close()
    _, err47 := mbTrans46.WriteString(arg45)
    if err47 != nil {
      Usage()
      return
    }
    factory48 := thrift.NewTJSONProtocolFactory()
    jsProt49 := factory48.GetProtocol(mbTrans46)
    argvalue0 := beeswax.NewQuery()
    err50 := argvalue0.Read(context.Background(), jsProt49)
    if err50 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.Query(context.Background(), value0))
    fmt.Print("\n")
    break
  case "executeAndWait":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "ExecuteAndWait requires 2 args")
      flag.Usage()
    }
    arg51 := flag.Arg(1)
    mbTrans52 := thrift.NewTMemoryBufferLen(len(arg51))
    defer mbTrans52.Close()
    _, err53 := mbTrans52.WriteString(arg51)
    if err53 != nil {
      Usage()
      return
    }
    factory54 := thrift.NewTJSONProtocolFactory()
    jsProt55 := factory54.GetProtocol(mbTrans52)
    argvalue0 := beeswax.NewQuery()
    err56 := argvalue0.Read(context.Background(), jsProt55)
    if err56 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := beeswax.LogContextId(argvalue1)
    fmt.Print(client.ExecuteAndWait(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "explain":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Explain requires 1 args")
      flag.Usage()
    }
    arg58 := flag.Arg(1)
    mbTrans59 := thrift.NewTMemoryBufferLen(len(arg58))
    defer mbTrans59.Close()
    _, err60 := mbTrans59.WriteString(arg58)
    if err60 != nil {
      Usage()
      return
    }
    factory61 := thrift.NewTJSONProtocolFactory()
    jsProt62 := factory61.GetProtocol(mbTrans59)
    argvalue0 := beeswax.NewQuery()
    err63 := argvalue0.Read(context.Background(), jsProt62)
    if err63 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.Explain(context.Background(), value0))
    fmt.Print("\n")
    break
  case "fetch":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "Fetch requires 3 args")
      flag.Usage()
    }
    arg64 := flag.Arg(1)
    mbTrans65 := thrift.NewTMemoryBufferLen(len(arg64))
    defer mbTrans65.Close()
    _, err66 := mbTrans65.WriteString(arg64)
    if err66 != nil {
      Usage()
      return
    }
    factory67 := thrift.NewTJSONProtocolFactory()
    jsProt68 := factory67.GetProtocol(mbTrans65)
    argvalue0 := beeswax.NewQueryHandle()
    err69 := argvalue0.Read(context.Background(), jsProt68)
    if err69 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2) == "true"
    value1 := argvalue1
    tmp2, err71 := (strconv.Atoi(flag.Arg(3)))
    if err71 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    fmt.Print(client.Fetch(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "get_state":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetState requires 1 args")
      flag.Usage()
    }
    arg72 := flag.Arg(1)
    mbTrans73 := thrift.NewTMemoryBufferLen(len(arg72))
    defer mbTrans73.Close()
    _, err74 := mbTrans73.WriteString(arg72)
    if err74 != nil {
      Usage()
      return
    }
    factory75 := thrift.NewTJSONProtocolFactory()
    jsProt76 := factory75.GetProtocol(mbTrans73)
    argvalue0 := beeswax.NewQueryHandle()
    err77 := argvalue0.Read(context.Background(), jsProt76)
    if err77 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetState(context.Background(), value0))
    fmt.Print("\n")
    break
  case "get_results_metadata":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetResultsMetadata requires 1 args")
      flag.Usage()
    }
    arg78 := flag.Arg(1)
    mbTrans79 := thrift.NewTMemoryBufferLen(len(arg78))
    defer mbTrans79.Close()
    _, err80 := mbTrans79.WriteString(arg78)
    if err80 != nil {
      Usage()
      return
    }
    factory81 := thrift.NewTJSONProtocolFactory()
    jsProt82 := factory81.GetProtocol(mbTrans79)
    argvalue0 := beeswax.NewQueryHandle()
    err83 := argvalue0.Read(context.Background(), jsProt82)
    if err83 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetResultsMetadata(context.Background(), value0))
    fmt.Print("\n")
    break
  case "echo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Echo requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.Echo(context.Background(), value0))
    fmt.Print("\n")
    break
  case "dump_config":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "DumpConfig requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.DumpConfig(context.Background()))
    fmt.Print("\n")
    break
  case "get_log":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetLog requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := beeswax.LogContextId(argvalue0)
    fmt.Print(client.GetLog(context.Background(), value0))
    fmt.Print("\n")
    break
  case "get_default_configuration":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetDefaultConfiguration requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1) == "true"
    value0 := argvalue0
    fmt.Print(client.GetDefaultConfiguration(context.Background(), value0))
    fmt.Print("\n")
    break
  case "close":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Close requires 1 args")
      flag.Usage()
    }
    arg87 := flag.Arg(1)
    mbTrans88 := thrift.NewTMemoryBufferLen(len(arg87))
    defer mbTrans88.Close()
    _, err89 := mbTrans88.WriteString(arg87)
    if err89 != nil {
      Usage()
      return
    }
    factory90 := thrift.NewTJSONProtocolFactory()
    jsProt91 := factory90.GetProtocol(mbTrans88)
    argvalue0 := beeswax.NewQueryHandle()
    err92 := argvalue0.Read(context.Background(), jsProt91)
    if err92 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.Close(context.Background(), value0))
    fmt.Print("\n")
    break
  case "clean":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Clean requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := beeswax.LogContextId(argvalue0)
    fmt.Print(client.Clean(context.Background(), value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
