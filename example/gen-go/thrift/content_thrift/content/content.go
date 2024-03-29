// Code generated by Thrift Compiler (0.15.0). DO NOT EDIT.

package content

import (
	"bytes"
	"context"
	"fmt"
	"time"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"github.com/YLeseclaireurs/icafe/example/gen-go/thrift/content_thrift/exception"
	"github.com/YLeseclaireurs/icafe/example/gen-go/thrift/content_thrift/base"

)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal

var _ = exception.GoUnusedProtection__
var _ = base.GoUnusedProtection__
// Attributes:
//  - ContentID
type GetContentParam struct {
  ContentID int64 `thrift:"content_id,1,required" db:"content_id" json:"content_id"`
}

func NewGetContentParam() *GetContentParam {
  return &GetContentParam{}
}


func (p *GetContentParam) GetContentID() int64 {
  return p.ContentID
}
func (p *GetContentParam) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetContentID bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
        issetContentID = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  if !issetContentID{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field ContentID is not set"));
  }
  return nil
}

func (p *GetContentParam)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.ContentID = v
}
  return nil
}

func (p *GetContentParam) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "GetContentParam"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *GetContentParam) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "content_id", thrift.I64, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:content_id: ", p), err) }
  if err := oprot.WriteI64(ctx, int64(p.ContentID)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.content_id (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:content_id: ", p), err) }
  return err
}

func (p *GetContentParam) Equals(other *GetContentParam) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.ContentID != other.ContentID { return false }
  return true
}

func (p *GetContentParam) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("GetContentParam(%+v)", *p)
}

// Attributes:
//  - Content
type GetContentResponse struct {
  Content *base.Content `thrift:"content,1,required" db:"content" json:"content"`
}

func NewGetContentResponse() *GetContentResponse {
  return &GetContentResponse{}
}

var GetContentResponse_Content_DEFAULT *base.Content
func (p *GetContentResponse) GetContent() *base.Content {
  if !p.IsSetContent() {
    return GetContentResponse_Content_DEFAULT
  }
return p.Content
}
func (p *GetContentResponse) IsSetContent() bool {
  return p.Content != nil
}

func (p *GetContentResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetContent bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
        issetContent = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  if !issetContent{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Content is not set"));
  }
  return nil
}

func (p *GetContentResponse)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  p.Content = &base.Content{}
  if err := p.Content.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Content), err)
  }
  return nil
}

func (p *GetContentResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "GetContentResponse"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *GetContentResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "content", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:content: ", p), err) }
  if err := p.Content.Write(ctx, oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Content), err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:content: ", p), err) }
  return err
}

func (p *GetContentResponse) Equals(other *GetContentResponse) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if !p.Content.Equals(other.Content) { return false }
  return true
}

func (p *GetContentResponse) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("GetContentResponse(%+v)", *p)
}

type ContentService interface {
  // Parameters:
  //  - Param
  GetContent(ctx context.Context, param *GetContentParam) (_r *GetContentResponse, _err error)
}

type ContentServiceClient struct {
  c thrift.TClient
  meta thrift.ResponseMeta
}

func NewContentServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *ContentServiceClient {
  return &ContentServiceClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

func NewContentServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *ContentServiceClient {
  return &ContentServiceClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewContentServiceClient(c thrift.TClient) *ContentServiceClient {
  return &ContentServiceClient{
    c: c,
  }
}

func (p *ContentServiceClient) Client_() thrift.TClient {
  return p.c
}

func (p *ContentServiceClient) LastResponseMeta_() thrift.ResponseMeta {
  return p.meta
}

func (p *ContentServiceClient) SetLastResponseMeta_(meta thrift.ResponseMeta) {
  p.meta = meta
}

// Parameters:
//  - Param
func (p *ContentServiceClient) GetContent(ctx context.Context, param *GetContentParam) (_r *GetContentResponse, _err error) {
  var _args0 ContentServiceGetContentArgs
  _args0.Param = param
  var _result2 ContentServiceGetContentResult
  var _meta1 thrift.ResponseMeta
  _meta1, _err = p.Client_().Call(ctx, "get_content", &_args0, &_result2)
  p.SetLastResponseMeta_(_meta1)
  if _err != nil {
    return
  }
  if _ret3 := _result2.GetSuccess(); _ret3 != nil {
    return _ret3, nil
  }
  return nil, thrift.NewTApplicationException(thrift.MISSING_RESULT, "get_content failed: unknown result")
}

type ContentServiceProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler ContentService
}

func (p *ContentServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *ContentServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *ContentServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewContentServiceProcessor(handler ContentService) *ContentServiceProcessor {

  self4 := &ContentServiceProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self4.processorMap["get_content"] = &contentServiceProcessorGetContent{handler:handler}
return self4
}

func (p *ContentServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err2 := iprot.ReadMessageBegin(ctx)
  if err2 != nil { return false, thrift.WrapTException(err2) }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(ctx, thrift.STRUCT)
  iprot.ReadMessageEnd(ctx)
  x5 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(ctx, name, thrift.EXCEPTION, seqId)
  x5.Write(ctx, oprot)
  oprot.WriteMessageEnd(ctx)
  oprot.Flush(ctx)
  return false, x5

}

type contentServiceProcessorGetContent struct {
  handler ContentService
}

func (p *contentServiceProcessorGetContent) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := ContentServiceGetContentArgs{}
  var err2 error
  if err2 = args.Read(ctx, iprot); err2 != nil {
    iprot.ReadMessageEnd(ctx)
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
    oprot.WriteMessageBegin(ctx, "get_content", thrift.EXCEPTION, seqId)
    x.Write(ctx, oprot)
    oprot.WriteMessageEnd(ctx)
    oprot.Flush(ctx)
    return false, thrift.WrapTException(err2)
  }
  iprot.ReadMessageEnd(ctx)

  tickerCancel := func() {}
  // Start a goroutine to do server side connectivity check.
  if thrift.ServerConnectivityCheckInterval > 0 {
    var cancel context.CancelFunc
    ctx, cancel = context.WithCancel(ctx)
    defer cancel()
    var tickerCtx context.Context
    tickerCtx, tickerCancel = context.WithCancel(context.Background())
    defer tickerCancel()
    go func(ctx context.Context, cancel context.CancelFunc) {
      ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
      defer ticker.Stop()
      for {
        select {
        case <-ctx.Done():
          return
        case <-ticker.C:
          if !iprot.Transport().IsOpen() {
            cancel()
            return
          }
        }
      }
    }(tickerCtx, cancel)
  }

  result := ContentServiceGetContentResult{}
  var retval *GetContentResponse
  if retval, err2 = p.handler.GetContent(ctx, args.Param); err2 != nil {
    tickerCancel()
    if err2 == thrift.ErrAbandonRequest {
      return false, thrift.WrapTException(err2)
    }
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing get_content: " + err2.Error())
    oprot.WriteMessageBegin(ctx, "get_content", thrift.EXCEPTION, seqId)
    x.Write(ctx, oprot)
    oprot.WriteMessageEnd(ctx)
    oprot.Flush(ctx)
    return true, thrift.WrapTException(err2)
  } else {
    result.Success = retval
  }
  tickerCancel()
  if err2 = oprot.WriteMessageBegin(ctx, "get_content", thrift.REPLY, seqId); err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err2 = result.Write(ctx, oprot); err == nil && err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err2 = oprot.WriteMessageEnd(ctx); err == nil && err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Param
type ContentServiceGetContentArgs struct {
  Param *GetContentParam `thrift:"param,1" db:"param" json:"param"`
}

func NewContentServiceGetContentArgs() *ContentServiceGetContentArgs {
  return &ContentServiceGetContentArgs{}
}

var ContentServiceGetContentArgs_Param_DEFAULT *GetContentParam
func (p *ContentServiceGetContentArgs) GetParam() *GetContentParam {
  if !p.IsSetParam() {
    return ContentServiceGetContentArgs_Param_DEFAULT
  }
return p.Param
}
func (p *ContentServiceGetContentArgs) IsSetParam() bool {
  return p.Param != nil
}

func (p *ContentServiceGetContentArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *ContentServiceGetContentArgs)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  p.Param = &GetContentParam{}
  if err := p.Param.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Param), err)
  }
  return nil
}

func (p *ContentServiceGetContentArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "get_content_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ContentServiceGetContentArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "param", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:param: ", p), err) }
  if err := p.Param.Write(ctx, oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Param), err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:param: ", p), err) }
  return err
}

func (p *ContentServiceGetContentArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ContentServiceGetContentArgs(%+v)", *p)
}

// Attributes:
//  - Success
type ContentServiceGetContentResult struct {
  Success *GetContentResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewContentServiceGetContentResult() *ContentServiceGetContentResult {
  return &ContentServiceGetContentResult{}
}

var ContentServiceGetContentResult_Success_DEFAULT *GetContentResponse
func (p *ContentServiceGetContentResult) GetSuccess() *GetContentResponse {
  if !p.IsSetSuccess() {
    return ContentServiceGetContentResult_Success_DEFAULT
  }
return p.Success
}
func (p *ContentServiceGetContentResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *ContentServiceGetContentResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField0(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *ContentServiceGetContentResult)  ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
  p.Success = &GetContentResponse{}
  if err := p.Success.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *ContentServiceGetContentResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "get_content_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ContentServiceGetContentResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(ctx, oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(ctx); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *ContentServiceGetContentResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ContentServiceGetContentResult(%+v)", *p)
}


