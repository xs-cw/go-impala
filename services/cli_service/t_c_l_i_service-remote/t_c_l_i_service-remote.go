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
	"github.com/bippio/go-impala/services/cli_service"
)

var _ = cli_service.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
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
  client := cli_service.NewTCLIServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "OpenSession":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "OpenSession requires 1 args")
      flag.Usage()
    }
    arg110 := flag.Arg(1)
    mbTrans111 := thrift.NewTMemoryBufferLen(len(arg110))
    defer mbTrans111.Close()
    _, err112 := mbTrans111.WriteString(arg110)
    if err112 != nil {
      Usage()
      return
    }
    factory113 := thrift.NewTJSONProtocolFactory()
    jsProt114 := factory113.GetProtocol(mbTrans111)
    argvalue0 := cli_service.NewTOpenSessionReq()
    err115 := argvalue0.Read(context.Background(), jsProt114)
    if err115 != nil {
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
    arg116 := flag.Arg(1)
    mbTrans117 := thrift.NewTMemoryBufferLen(len(arg116))
    defer mbTrans117.Close()
    _, err118 := mbTrans117.WriteString(arg116)
    if err118 != nil {
      Usage()
      return
    }
    factory119 := thrift.NewTJSONProtocolFactory()
    jsProt120 := factory119.GetProtocol(mbTrans117)
    argvalue0 := cli_service.NewTCloseSessionReq()
    err121 := argvalue0.Read(context.Background(), jsProt120)
    if err121 != nil {
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
    arg122 := flag.Arg(1)
    mbTrans123 := thrift.NewTMemoryBufferLen(len(arg122))
    defer mbTrans123.Close()
    _, err124 := mbTrans123.WriteString(arg122)
    if err124 != nil {
      Usage()
      return
    }
    factory125 := thrift.NewTJSONProtocolFactory()
    jsProt126 := factory125.GetProtocol(mbTrans123)
    argvalue0 := cli_service.NewTGetInfoReq()
    err127 := argvalue0.Read(context.Background(), jsProt126)
    if err127 != nil {
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
    arg128 := flag.Arg(1)
    mbTrans129 := thrift.NewTMemoryBufferLen(len(arg128))
    defer mbTrans129.Close()
    _, err130 := mbTrans129.WriteString(arg128)
    if err130 != nil {
      Usage()
      return
    }
    factory131 := thrift.NewTJSONProtocolFactory()
    jsProt132 := factory131.GetProtocol(mbTrans129)
    argvalue0 := cli_service.NewTExecuteStatementReq()
    err133 := argvalue0.Read(context.Background(), jsProt132)
    if err133 != nil {
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
    arg134 := flag.Arg(1)
    mbTrans135 := thrift.NewTMemoryBufferLen(len(arg134))
    defer mbTrans135.Close()
    _, err136 := mbTrans135.WriteString(arg134)
    if err136 != nil {
      Usage()
      return
    }
    factory137 := thrift.NewTJSONProtocolFactory()
    jsProt138 := factory137.GetProtocol(mbTrans135)
    argvalue0 := cli_service.NewTGetTypeInfoReq()
    err139 := argvalue0.Read(context.Background(), jsProt138)
    if err139 != nil {
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
    arg140 := flag.Arg(1)
    mbTrans141 := thrift.NewTMemoryBufferLen(len(arg140))
    defer mbTrans141.Close()
    _, err142 := mbTrans141.WriteString(arg140)
    if err142 != nil {
      Usage()
      return
    }
    factory143 := thrift.NewTJSONProtocolFactory()
    jsProt144 := factory143.GetProtocol(mbTrans141)
    argvalue0 := cli_service.NewTGetCatalogsReq()
    err145 := argvalue0.Read(context.Background(), jsProt144)
    if err145 != nil {
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
    arg146 := flag.Arg(1)
    mbTrans147 := thrift.NewTMemoryBufferLen(len(arg146))
    defer mbTrans147.Close()
    _, err148 := mbTrans147.WriteString(arg146)
    if err148 != nil {
      Usage()
      return
    }
    factory149 := thrift.NewTJSONProtocolFactory()
    jsProt150 := factory149.GetProtocol(mbTrans147)
    argvalue0 := cli_service.NewTGetSchemasReq()
    err151 := argvalue0.Read(context.Background(), jsProt150)
    if err151 != nil {
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
    arg152 := flag.Arg(1)
    mbTrans153 := thrift.NewTMemoryBufferLen(len(arg152))
    defer mbTrans153.Close()
    _, err154 := mbTrans153.WriteString(arg152)
    if err154 != nil {
      Usage()
      return
    }
    factory155 := thrift.NewTJSONProtocolFactory()
    jsProt156 := factory155.GetProtocol(mbTrans153)
    argvalue0 := cli_service.NewTGetTablesReq()
    err157 := argvalue0.Read(context.Background(), jsProt156)
    if err157 != nil {
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
    arg158 := flag.Arg(1)
    mbTrans159 := thrift.NewTMemoryBufferLen(len(arg158))
    defer mbTrans159.Close()
    _, err160 := mbTrans159.WriteString(arg158)
    if err160 != nil {
      Usage()
      return
    }
    factory161 := thrift.NewTJSONProtocolFactory()
    jsProt162 := factory161.GetProtocol(mbTrans159)
    argvalue0 := cli_service.NewTGetTableTypesReq()
    err163 := argvalue0.Read(context.Background(), jsProt162)
    if err163 != nil {
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
    arg164 := flag.Arg(1)
    mbTrans165 := thrift.NewTMemoryBufferLen(len(arg164))
    defer mbTrans165.Close()
    _, err166 := mbTrans165.WriteString(arg164)
    if err166 != nil {
      Usage()
      return
    }
    factory167 := thrift.NewTJSONProtocolFactory()
    jsProt168 := factory167.GetProtocol(mbTrans165)
    argvalue0 := cli_service.NewTGetColumnsReq()
    err169 := argvalue0.Read(context.Background(), jsProt168)
    if err169 != nil {
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
    arg170 := flag.Arg(1)
    mbTrans171 := thrift.NewTMemoryBufferLen(len(arg170))
    defer mbTrans171.Close()
    _, err172 := mbTrans171.WriteString(arg170)
    if err172 != nil {
      Usage()
      return
    }
    factory173 := thrift.NewTJSONProtocolFactory()
    jsProt174 := factory173.GetProtocol(mbTrans171)
    argvalue0 := cli_service.NewTGetFunctionsReq()
    err175 := argvalue0.Read(context.Background(), jsProt174)
    if err175 != nil {
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
    arg176 := flag.Arg(1)
    mbTrans177 := thrift.NewTMemoryBufferLen(len(arg176))
    defer mbTrans177.Close()
    _, err178 := mbTrans177.WriteString(arg176)
    if err178 != nil {
      Usage()
      return
    }
    factory179 := thrift.NewTJSONProtocolFactory()
    jsProt180 := factory179.GetProtocol(mbTrans177)
    argvalue0 := cli_service.NewTGetOperationStatusReq()
    err181 := argvalue0.Read(context.Background(), jsProt180)
    if err181 != nil {
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
    arg182 := flag.Arg(1)
    mbTrans183 := thrift.NewTMemoryBufferLen(len(arg182))
    defer mbTrans183.Close()
    _, err184 := mbTrans183.WriteString(arg182)
    if err184 != nil {
      Usage()
      return
    }
    factory185 := thrift.NewTJSONProtocolFactory()
    jsProt186 := factory185.GetProtocol(mbTrans183)
    argvalue0 := cli_service.NewTCancelOperationReq()
    err187 := argvalue0.Read(context.Background(), jsProt186)
    if err187 != nil {
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
    arg188 := flag.Arg(1)
    mbTrans189 := thrift.NewTMemoryBufferLen(len(arg188))
    defer mbTrans189.Close()
    _, err190 := mbTrans189.WriteString(arg188)
    if err190 != nil {
      Usage()
      return
    }
    factory191 := thrift.NewTJSONProtocolFactory()
    jsProt192 := factory191.GetProtocol(mbTrans189)
    argvalue0 := cli_service.NewTCloseOperationReq()
    err193 := argvalue0.Read(context.Background(), jsProt192)
    if err193 != nil {
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
    arg194 := flag.Arg(1)
    mbTrans195 := thrift.NewTMemoryBufferLen(len(arg194))
    defer mbTrans195.Close()
    _, err196 := mbTrans195.WriteString(arg194)
    if err196 != nil {
      Usage()
      return
    }
    factory197 := thrift.NewTJSONProtocolFactory()
    jsProt198 := factory197.GetProtocol(mbTrans195)
    argvalue0 := cli_service.NewTGetResultSetMetadataReq()
    err199 := argvalue0.Read(context.Background(), jsProt198)
    if err199 != nil {
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
    arg200 := flag.Arg(1)
    mbTrans201 := thrift.NewTMemoryBufferLen(len(arg200))
    defer mbTrans201.Close()
    _, err202 := mbTrans201.WriteString(arg200)
    if err202 != nil {
      Usage()
      return
    }
    factory203 := thrift.NewTJSONProtocolFactory()
    jsProt204 := factory203.GetProtocol(mbTrans201)
    argvalue0 := cli_service.NewTFetchResultsReq()
    err205 := argvalue0.Read(context.Background(), jsProt204)
    if err205 != nil {
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
    arg206 := flag.Arg(1)
    mbTrans207 := thrift.NewTMemoryBufferLen(len(arg206))
    defer mbTrans207.Close()
    _, err208 := mbTrans207.WriteString(arg206)
    if err208 != nil {
      Usage()
      return
    }
    factory209 := thrift.NewTJSONProtocolFactory()
    jsProt210 := factory209.GetProtocol(mbTrans207)
    argvalue0 := cli_service.NewTGetDelegationTokenReq()
    err211 := argvalue0.Read(context.Background(), jsProt210)
    if err211 != nil {
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
    arg212 := flag.Arg(1)
    mbTrans213 := thrift.NewTMemoryBufferLen(len(arg212))
    defer mbTrans213.Close()
    _, err214 := mbTrans213.WriteString(arg212)
    if err214 != nil {
      Usage()
      return
    }
    factory215 := thrift.NewTJSONProtocolFactory()
    jsProt216 := factory215.GetProtocol(mbTrans213)
    argvalue0 := cli_service.NewTCancelDelegationTokenReq()
    err217 := argvalue0.Read(context.Background(), jsProt216)
    if err217 != nil {
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
    arg218 := flag.Arg(1)
    mbTrans219 := thrift.NewTMemoryBufferLen(len(arg218))
    defer mbTrans219.Close()
    _, err220 := mbTrans219.WriteString(arg218)
    if err220 != nil {
      Usage()
      return
    }
    factory221 := thrift.NewTJSONProtocolFactory()
    jsProt222 := factory221.GetProtocol(mbTrans219)
    argvalue0 := cli_service.NewTRenewDelegationTokenReq()
    err223 := argvalue0.Read(context.Background(), jsProt222)
    if err223 != nil {
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
    arg224 := flag.Arg(1)
    mbTrans225 := thrift.NewTMemoryBufferLen(len(arg224))
    defer mbTrans225.Close()
    _, err226 := mbTrans225.WriteString(arg224)
    if err226 != nil {
      Usage()
      return
    }
    factory227 := thrift.NewTJSONProtocolFactory()
    jsProt228 := factory227.GetProtocol(mbTrans225)
    argvalue0 := cli_service.NewTGetLogReq()
    err229 := argvalue0.Read(context.Background(), jsProt228)
    if err229 != nil {
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
