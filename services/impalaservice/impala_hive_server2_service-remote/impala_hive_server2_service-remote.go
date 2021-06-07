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
	"github.com/bippio/go-impala/services/status"
	"github.com/bippio/go-impala/services/beeswax"
	"github.com/bippio/go-impala/services/cli_service"
	"github.com/bippio/go-impala/services/impalaservice"
)

var _ = status.GoUnusedProtection__
var _ = beeswax.GoUnusedProtection__
var _ = cli_service.GoUnusedProtection__
var _ = impalaservice.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  TStatus ResetCatalog()")
  fmt.Fprintln(os.Stderr, "  TOpenSessionResp OpenSession(TOpenSessionReq req)")
  fmt.Fprintln(os.Stderr, "  TCloseSessionResp CloseSession(TCloseSessionReq req)")
  fmt.Fprintln(os.Stderr, "  TGetInfoResp GetInfo(TGetInfoReq req)")
  fmt.Fprintln(os.Stderr, "  TExecuteStatementResp ExecuteStatement(TExecuteStatementReq req)")
  fmt.Fprintln(os.Stderr, "  TGetTypeInfoResp GetTypeInfo(TGetTypeInfoReq req)")
  fmt.Fprintln(os.Stderr, "  TGetCatalogsResp GetCatalogs(TGetCatalogsReq req)")
  fmt.Fprintln(os.Stderr, "  TGetSchemasResp GetSchemas(TGetSchemasReq req)")
  fmt.Fprintln(os.Stderr, "  TGetTablesResp GetTables(TGetTablesReq req)")
  fmt.Fprintln(os.Stderr, "  TGetTableTypesResp GetTableTypes(TGetTableTypesReq req)")
  fmt.Fprintln(os.Stderr, "  TGetColumnsResp GetColumns(TGetColumnsReq req)")
  fmt.Fprintln(os.Stderr, "  TGetFunctionsResp GetFunctions(TGetFunctionsReq req)")
  fmt.Fprintln(os.Stderr, "  TGetOperationStatusResp GetOperationStatus(TGetOperationStatusReq req)")
  fmt.Fprintln(os.Stderr, "  TCancelOperationResp CancelOperation(TCancelOperationReq req)")
  fmt.Fprintln(os.Stderr, "  TCloseOperationResp CloseOperation(TCloseOperationReq req)")
  fmt.Fprintln(os.Stderr, "  TGetResultSetMetadataResp GetResultSetMetadata(TGetResultSetMetadataReq req)")
  fmt.Fprintln(os.Stderr, "  TFetchResultsResp FetchResults(TFetchResultsReq req)")
  fmt.Fprintln(os.Stderr, "  TGetDelegationTokenResp GetDelegationToken(TGetDelegationTokenReq req)")
  fmt.Fprintln(os.Stderr, "  TCancelDelegationTokenResp CancelDelegationToken(TCancelDelegationTokenReq req)")
  fmt.Fprintln(os.Stderr, "  TRenewDelegationTokenResp RenewDelegationToken(TRenewDelegationTokenReq req)")
  fmt.Fprintln(os.Stderr, "  TGetLogResp GetLog(TGetLogReq req)")
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
  client := impalaservice.NewImpalaHiveServer2ServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "ResetCatalog":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "ResetCatalog requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.ResetCatalog(context.Background()))
    fmt.Print("\n")
    break
  case "OpenSession":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "OpenSession requires 1 args")
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
    argvalue0 := cli_service.NewTOpenSessionReq()
    err83 := argvalue0.Read(context.Background(), jsProt82)
    if err83 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.OpenSession(context.Background(), value0))
    fmt.Print("\n")
    break
  case "CloseSession":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CloseSession requires 1 args")
      flag.Usage()
    }
    arg84 := flag.Arg(1)
    mbTrans85 := thrift.NewTMemoryBufferLen(len(arg84))
    defer mbTrans85.Close()
    _, err86 := mbTrans85.WriteString(arg84)
    if err86 != nil {
      Usage()
      return
    }
    factory87 := thrift.NewTJSONProtocolFactory()
    jsProt88 := factory87.GetProtocol(mbTrans85)
    argvalue0 := cli_service.NewTCloseSessionReq()
    err89 := argvalue0.Read(context.Background(), jsProt88)
    if err89 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CloseSession(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetInfo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetInfo requires 1 args")
      flag.Usage()
    }
    arg90 := flag.Arg(1)
    mbTrans91 := thrift.NewTMemoryBufferLen(len(arg90))
    defer mbTrans91.Close()
    _, err92 := mbTrans91.WriteString(arg90)
    if err92 != nil {
      Usage()
      return
    }
    factory93 := thrift.NewTJSONProtocolFactory()
    jsProt94 := factory93.GetProtocol(mbTrans91)
    argvalue0 := cli_service.NewTGetInfoReq()
    err95 := argvalue0.Read(context.Background(), jsProt94)
    if err95 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetInfo(context.Background(), value0))
    fmt.Print("\n")
    break
  case "ExecuteStatement":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "ExecuteStatement requires 1 args")
      flag.Usage()
    }
    arg96 := flag.Arg(1)
    mbTrans97 := thrift.NewTMemoryBufferLen(len(arg96))
    defer mbTrans97.Close()
    _, err98 := mbTrans97.WriteString(arg96)
    if err98 != nil {
      Usage()
      return
    }
    factory99 := thrift.NewTJSONProtocolFactory()
    jsProt100 := factory99.GetProtocol(mbTrans97)
    argvalue0 := cli_service.NewTExecuteStatementReq()
    err101 := argvalue0.Read(context.Background(), jsProt100)
    if err101 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.ExecuteStatement(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetTypeInfo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTypeInfo requires 1 args")
      flag.Usage()
    }
    arg102 := flag.Arg(1)
    mbTrans103 := thrift.NewTMemoryBufferLen(len(arg102))
    defer mbTrans103.Close()
    _, err104 := mbTrans103.WriteString(arg102)
    if err104 != nil {
      Usage()
      return
    }
    factory105 := thrift.NewTJSONProtocolFactory()
    jsProt106 := factory105.GetProtocol(mbTrans103)
    argvalue0 := cli_service.NewTGetTypeInfoReq()
    err107 := argvalue0.Read(context.Background(), jsProt106)
    if err107 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetTypeInfo(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetCatalogs":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetCatalogs requires 1 args")
      flag.Usage()
    }
    arg108 := flag.Arg(1)
    mbTrans109 := thrift.NewTMemoryBufferLen(len(arg108))
    defer mbTrans109.Close()
    _, err110 := mbTrans109.WriteString(arg108)
    if err110 != nil {
      Usage()
      return
    }
    factory111 := thrift.NewTJSONProtocolFactory()
    jsProt112 := factory111.GetProtocol(mbTrans109)
    argvalue0 := cli_service.NewTGetCatalogsReq()
    err113 := argvalue0.Read(context.Background(), jsProt112)
    if err113 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetCatalogs(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetSchemas":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetSchemas requires 1 args")
      flag.Usage()
    }
    arg114 := flag.Arg(1)
    mbTrans115 := thrift.NewTMemoryBufferLen(len(arg114))
    defer mbTrans115.Close()
    _, err116 := mbTrans115.WriteString(arg114)
    if err116 != nil {
      Usage()
      return
    }
    factory117 := thrift.NewTJSONProtocolFactory()
    jsProt118 := factory117.GetProtocol(mbTrans115)
    argvalue0 := cli_service.NewTGetSchemasReq()
    err119 := argvalue0.Read(context.Background(), jsProt118)
    if err119 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetSchemas(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetTables":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTables requires 1 args")
      flag.Usage()
    }
    arg120 := flag.Arg(1)
    mbTrans121 := thrift.NewTMemoryBufferLen(len(arg120))
    defer mbTrans121.Close()
    _, err122 := mbTrans121.WriteString(arg120)
    if err122 != nil {
      Usage()
      return
    }
    factory123 := thrift.NewTJSONProtocolFactory()
    jsProt124 := factory123.GetProtocol(mbTrans121)
    argvalue0 := cli_service.NewTGetTablesReq()
    err125 := argvalue0.Read(context.Background(), jsProt124)
    if err125 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetTables(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetTableTypes":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTableTypes requires 1 args")
      flag.Usage()
    }
    arg126 := flag.Arg(1)
    mbTrans127 := thrift.NewTMemoryBufferLen(len(arg126))
    defer mbTrans127.Close()
    _, err128 := mbTrans127.WriteString(arg126)
    if err128 != nil {
      Usage()
      return
    }
    factory129 := thrift.NewTJSONProtocolFactory()
    jsProt130 := factory129.GetProtocol(mbTrans127)
    argvalue0 := cli_service.NewTGetTableTypesReq()
    err131 := argvalue0.Read(context.Background(), jsProt130)
    if err131 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetTableTypes(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetColumns":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetColumns requires 1 args")
      flag.Usage()
    }
    arg132 := flag.Arg(1)
    mbTrans133 := thrift.NewTMemoryBufferLen(len(arg132))
    defer mbTrans133.Close()
    _, err134 := mbTrans133.WriteString(arg132)
    if err134 != nil {
      Usage()
      return
    }
    factory135 := thrift.NewTJSONProtocolFactory()
    jsProt136 := factory135.GetProtocol(mbTrans133)
    argvalue0 := cli_service.NewTGetColumnsReq()
    err137 := argvalue0.Read(context.Background(), jsProt136)
    if err137 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetColumns(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetFunctions":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetFunctions requires 1 args")
      flag.Usage()
    }
    arg138 := flag.Arg(1)
    mbTrans139 := thrift.NewTMemoryBufferLen(len(arg138))
    defer mbTrans139.Close()
    _, err140 := mbTrans139.WriteString(arg138)
    if err140 != nil {
      Usage()
      return
    }
    factory141 := thrift.NewTJSONProtocolFactory()
    jsProt142 := factory141.GetProtocol(mbTrans139)
    argvalue0 := cli_service.NewTGetFunctionsReq()
    err143 := argvalue0.Read(context.Background(), jsProt142)
    if err143 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetFunctions(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetOperationStatus":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetOperationStatus requires 1 args")
      flag.Usage()
    }
    arg144 := flag.Arg(1)
    mbTrans145 := thrift.NewTMemoryBufferLen(len(arg144))
    defer mbTrans145.Close()
    _, err146 := mbTrans145.WriteString(arg144)
    if err146 != nil {
      Usage()
      return
    }
    factory147 := thrift.NewTJSONProtocolFactory()
    jsProt148 := factory147.GetProtocol(mbTrans145)
    argvalue0 := cli_service.NewTGetOperationStatusReq()
    err149 := argvalue0.Read(context.Background(), jsProt148)
    if err149 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetOperationStatus(context.Background(), value0))
    fmt.Print("\n")
    break
  case "CancelOperation":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CancelOperation requires 1 args")
      flag.Usage()
    }
    arg150 := flag.Arg(1)
    mbTrans151 := thrift.NewTMemoryBufferLen(len(arg150))
    defer mbTrans151.Close()
    _, err152 := mbTrans151.WriteString(arg150)
    if err152 != nil {
      Usage()
      return
    }
    factory153 := thrift.NewTJSONProtocolFactory()
    jsProt154 := factory153.GetProtocol(mbTrans151)
    argvalue0 := cli_service.NewTCancelOperationReq()
    err155 := argvalue0.Read(context.Background(), jsProt154)
    if err155 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CancelOperation(context.Background(), value0))
    fmt.Print("\n")
    break
  case "CloseOperation":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CloseOperation requires 1 args")
      flag.Usage()
    }
    arg156 := flag.Arg(1)
    mbTrans157 := thrift.NewTMemoryBufferLen(len(arg156))
    defer mbTrans157.Close()
    _, err158 := mbTrans157.WriteString(arg156)
    if err158 != nil {
      Usage()
      return
    }
    factory159 := thrift.NewTJSONProtocolFactory()
    jsProt160 := factory159.GetProtocol(mbTrans157)
    argvalue0 := cli_service.NewTCloseOperationReq()
    err161 := argvalue0.Read(context.Background(), jsProt160)
    if err161 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CloseOperation(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetResultSetMetadata":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetResultSetMetadata requires 1 args")
      flag.Usage()
    }
    arg162 := flag.Arg(1)
    mbTrans163 := thrift.NewTMemoryBufferLen(len(arg162))
    defer mbTrans163.Close()
    _, err164 := mbTrans163.WriteString(arg162)
    if err164 != nil {
      Usage()
      return
    }
    factory165 := thrift.NewTJSONProtocolFactory()
    jsProt166 := factory165.GetProtocol(mbTrans163)
    argvalue0 := cli_service.NewTGetResultSetMetadataReq()
    err167 := argvalue0.Read(context.Background(), jsProt166)
    if err167 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetResultSetMetadata(context.Background(), value0))
    fmt.Print("\n")
    break
  case "FetchResults":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "FetchResults requires 1 args")
      flag.Usage()
    }
    arg168 := flag.Arg(1)
    mbTrans169 := thrift.NewTMemoryBufferLen(len(arg168))
    defer mbTrans169.Close()
    _, err170 := mbTrans169.WriteString(arg168)
    if err170 != nil {
      Usage()
      return
    }
    factory171 := thrift.NewTJSONProtocolFactory()
    jsProt172 := factory171.GetProtocol(mbTrans169)
    argvalue0 := cli_service.NewTFetchResultsReq()
    err173 := argvalue0.Read(context.Background(), jsProt172)
    if err173 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.FetchResults(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetDelegationToken":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetDelegationToken requires 1 args")
      flag.Usage()
    }
    arg174 := flag.Arg(1)
    mbTrans175 := thrift.NewTMemoryBufferLen(len(arg174))
    defer mbTrans175.Close()
    _, err176 := mbTrans175.WriteString(arg174)
    if err176 != nil {
      Usage()
      return
    }
    factory177 := thrift.NewTJSONProtocolFactory()
    jsProt178 := factory177.GetProtocol(mbTrans175)
    argvalue0 := cli_service.NewTGetDelegationTokenReq()
    err179 := argvalue0.Read(context.Background(), jsProt178)
    if err179 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetDelegationToken(context.Background(), value0))
    fmt.Print("\n")
    break
  case "CancelDelegationToken":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CancelDelegationToken requires 1 args")
      flag.Usage()
    }
    arg180 := flag.Arg(1)
    mbTrans181 := thrift.NewTMemoryBufferLen(len(arg180))
    defer mbTrans181.Close()
    _, err182 := mbTrans181.WriteString(arg180)
    if err182 != nil {
      Usage()
      return
    }
    factory183 := thrift.NewTJSONProtocolFactory()
    jsProt184 := factory183.GetProtocol(mbTrans181)
    argvalue0 := cli_service.NewTCancelDelegationTokenReq()
    err185 := argvalue0.Read(context.Background(), jsProt184)
    if err185 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CancelDelegationToken(context.Background(), value0))
    fmt.Print("\n")
    break
  case "RenewDelegationToken":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "RenewDelegationToken requires 1 args")
      flag.Usage()
    }
    arg186 := flag.Arg(1)
    mbTrans187 := thrift.NewTMemoryBufferLen(len(arg186))
    defer mbTrans187.Close()
    _, err188 := mbTrans187.WriteString(arg186)
    if err188 != nil {
      Usage()
      return
    }
    factory189 := thrift.NewTJSONProtocolFactory()
    jsProt190 := factory189.GetProtocol(mbTrans187)
    argvalue0 := cli_service.NewTRenewDelegationTokenReq()
    err191 := argvalue0.Read(context.Background(), jsProt190)
    if err191 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.RenewDelegationToken(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetLog":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetLog requires 1 args")
      flag.Usage()
    }
    arg192 := flag.Arg(1)
    mbTrans193 := thrift.NewTMemoryBufferLen(len(arg192))
    defer mbTrans193.Close()
    _, err194 := mbTrans193.WriteString(arg192)
    if err194 != nil {
      Usage()
      return
    }
    factory195 := thrift.NewTJSONProtocolFactory()
    jsProt196 := factory195.GetProtocol(mbTrans193)
    argvalue0 := cli_service.NewTGetLogReq()
    err197 := argvalue0.Read(context.Background(), jsProt196)
    if err197 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetLog(context.Background(), value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
