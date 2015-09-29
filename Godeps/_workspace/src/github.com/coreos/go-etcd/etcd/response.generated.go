// ************************************************************
// DO NOT EDIT.
// THIS FILE IS AUTO-GENERATED BY codecgen.
// ************************************************************

package etcd

import (
	"errors"
	"fmt"
	codec1978 "github.com/FoxComm/vulcand/Godeps/_workspace/src/github.com/ugorji/go/codec"
	"net/http"
	"reflect"
	"runtime"
	"time"
)

const (
	codecSelferC_UTF84402         = 1
	codecSelferC_RAW4402          = 0
	codecSelverValueTypeArray4402 = 10
	codecSelverValueTypeMap4402   = 9
)

var (
	codecSelferBitsize4402                         = uint8(reflect.TypeOf(uint(0)).Bits())
	codecSelferOnlyMapOrArrayEncodeToStructErr4402 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer4402 struct{}

func init() {
	if codec1978.GenVersion != 2 {
		_, file, _, _ := runtime.Caller(0)
		err := fmt.Errorf("codecgen version mismatch: current: %v, need %v. Re-generate file: %v",
			2, codec1978.GenVersion, file)
		panic(err)
	}
	if false { // reference the types, but skip this branch at build/run time
		var v0 http.Header
		var v1 time.Time
		_, _ = v0, v1
	}
}

func (x responseType) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer4402
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.EncodeInt(int64(x))
}

func (x *responseType) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer4402
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	*((*int)(x)) = int(r.DecodeInt(codecSelferBitsize4402))
}

func (x *RawResponse) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer4402
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yysep1 := !z.EncBinary()
		yy2arr1 := z.EncBasicHandle().StructToArray
		var yyfirst1 bool
		var yyq1 [3]bool
		_, _, _, _ = yysep1, yyfirst1, yyq1, yy2arr1
		const yyr1 bool = false
		if yyr1 || yy2arr1 {
			r.EncodeArrayStart(3)
		} else {
			var yynn1 int = 3
			for _, b := range yyq1 {
				if b {
					yynn1++
				}
			}
			r.EncodeMapStart(yynn1)
		}
		if yyr1 || yy2arr1 {
			r.EncodeInt(int64(x.StatusCode))
		} else {
			yyfirst1 = true
			r.EncodeString(codecSelferC_UTF84402, string("StatusCode"))
			if yysep1 {
				r.EncodeMapKVSeparator()
			}
			r.EncodeInt(int64(x.StatusCode))
		}
		if yyr1 || yy2arr1 {
			if yysep1 {
				r.EncodeArrayEntrySeparator()
			}
			if x.Body == nil {
				r.EncodeNil()
			} else {
				r.EncodeStringBytes(codecSelferC_RAW4402, []byte(x.Body))
			}
		} else {
			if yyfirst1 {
				r.EncodeMapEntrySeparator()
			} else {
				yyfirst1 = true
			}
			r.EncodeString(codecSelferC_UTF84402, string("Body"))
			if yysep1 {
				r.EncodeMapKVSeparator()
			}
			if x.Body == nil {
				r.EncodeNil()
			} else {
				r.EncodeStringBytes(codecSelferC_RAW4402, []byte(x.Body))
			}
		}
		if yyr1 || yy2arr1 {
			if yysep1 {
				r.EncodeArrayEntrySeparator()
			}
			if x.Header == nil {
				r.EncodeNil()
			} else {
				h.enchttp_Header(http.Header(x.Header), e)
			}
		} else {
			if yyfirst1 {
				r.EncodeMapEntrySeparator()
			} else {
				yyfirst1 = true
			}
			r.EncodeString(codecSelferC_UTF84402, string("Header"))
			if yysep1 {
				r.EncodeMapKVSeparator()
			}
			if x.Header == nil {
				r.EncodeNil()
			} else {
				h.enchttp_Header(http.Header(x.Header), e)
			}
		}
		if yysep1 {
			if yyr1 || yy2arr1 {
				r.EncodeArrayEnd()
			} else {
				r.EncodeMapEnd()
			}
		}
	}
}

func (x *RawResponse) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer4402
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if r.IsContainerType(codecSelverValueTypeMap4402) {
		yyl5 := r.ReadMapStart()
		if yyl5 == 0 {
			r.ReadMapEnd()
		} else {
			x.codecDecodeSelfFromMap(yyl5, d)
		}
	} else if r.IsContainerType(codecSelverValueTypeArray4402) {
		yyl5 := r.ReadArrayStart()
		if yyl5 == 0 {
			r.ReadArrayEnd()
		} else {
			x.codecDecodeSelfFromArray(yyl5, d)
		}
	} else {
		panic(codecSelferOnlyMapOrArrayEncodeToStructErr4402)
	}
}

func (x *RawResponse) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer4402
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys6Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys6Slc
	var yyhl6 bool = l >= 0
	for yyj6 := 0; ; yyj6++ {
		if yyhl6 {
			if yyj6 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
			if yyj6 > 0 {
				r.ReadMapEntrySeparator()
			}
		}
		yys6Slc = r.DecodeBytes(yys6Slc, true, true)
		yys6 := string(yys6Slc)
		if !yyhl6 {
			r.ReadMapKVSeparator()
		}
		switch yys6 {
		case "StatusCode":
			if r.TryDecodeAsNil() {
				x.StatusCode = 0
			} else {
				x.StatusCode = int(r.DecodeInt(codecSelferBitsize4402))
			}
		case "Body":
			if r.TryDecodeAsNil() {
				x.Body = nil
			} else {
				yyv8 := &x.Body
				*yyv8 = r.DecodeBytes(*(*[]byte)(yyv8), false, false)
			}
		case "Header":
			if r.TryDecodeAsNil() {
				x.Header = nil
			} else {
				yyv9 := &x.Header
				h.dechttp_Header((*http.Header)(yyv9), d)
			}
		default:
			z.DecStructFieldNotFound(-1, yys6)
		} // end switch yys6
	} // end for yyj6
	if !yyhl6 {
		r.ReadMapEnd()
	}
}

func (x *RawResponse) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer4402
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj10 int
	var yyb10 bool
	var yyhl10 bool = l >= 0
	yyj10++
	if yyhl10 {
		yyb10 = yyj10 > l
	} else {
		yyb10 = r.CheckBreak()
	}
	if yyb10 {
		r.ReadArrayEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.StatusCode = 0
	} else {
		x.StatusCode = int(r.DecodeInt(codecSelferBitsize4402))
	}
	yyj10++
	if yyhl10 {
		yyb10 = yyj10 > l
	} else {
		yyb10 = r.CheckBreak()
	}
	if yyb10 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayEntrySeparator()
	if r.TryDecodeAsNil() {
		x.Body = nil
	} else {
		yyv12 := &x.Body
		*yyv12 = r.DecodeBytes(*(*[]byte)(yyv12), false, false)
	}
	yyj10++
	if yyhl10 {
		yyb10 = yyj10 > l
	} else {
		yyb10 = r.CheckBreak()
	}
	if yyb10 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayEntrySeparator()
	if r.TryDecodeAsNil() {
		x.Header = nil
	} else {
		yyv13 := &x.Header
		h.dechttp_Header((*http.Header)(yyv13), d)
	}
	for {
		yyj10++
		if yyhl10 {
			yyb10 = yyj10 > l
		} else {
			yyb10 = r.CheckBreak()
		}
		if yyb10 {
			break
		}
		if yyj10 > 1 {
			r.ReadArrayEntrySeparator()
		}
		z.DecStructFieldNotFound(yyj10-1, "")
	}
	r.ReadArrayEnd()
}

func (x *Response) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer4402
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yysep14 := !z.EncBinary()
		yy2arr14 := z.EncBasicHandle().StructToArray
		var yyfirst14 bool
		var yyq14 [6]bool
		_, _, _, _ = yysep14, yyfirst14, yyq14, yy2arr14
		const yyr14 bool = false
		yyq14[2] = x.PrevNode != nil
		if yyr14 || yy2arr14 {
			r.EncodeArrayStart(6)
		} else {
			var yynn14 int = 5
			for _, b := range yyq14 {
				if b {
					yynn14++
				}
			}
			r.EncodeMapStart(yynn14)
		}
		if yyr14 || yy2arr14 {
			r.EncodeString(codecSelferC_UTF84402, string(x.Action))
		} else {
			yyfirst14 = true
			r.EncodeString(codecSelferC_UTF84402, string("action"))
			if yysep14 {
				r.EncodeMapKVSeparator()
			}
			r.EncodeString(codecSelferC_UTF84402, string(x.Action))
		}
		if yyr14 || yy2arr14 {
			if yysep14 {
				r.EncodeArrayEntrySeparator()
			}
			if x.Node == nil {
				r.EncodeNil()
			} else {
				x.Node.CodecEncodeSelf(e)
			}
		} else {
			if yyfirst14 {
				r.EncodeMapEntrySeparator()
			} else {
				yyfirst14 = true
			}
			r.EncodeString(codecSelferC_UTF84402, string("node"))
			if yysep14 {
				r.EncodeMapKVSeparator()
			}
			if x.Node == nil {
				r.EncodeNil()
			} else {
				x.Node.CodecEncodeSelf(e)
			}
		}
		if yyr14 || yy2arr14 {
			if yysep14 {
				r.EncodeArrayEntrySeparator()
			}
			if yyq14[2] {
				if x.PrevNode == nil {
					r.EncodeNil()
				} else {
					x.PrevNode.CodecEncodeSelf(e)
				}
			} else {
				r.EncodeNil()
			}
		} else {
			if yyq14[2] {
				if yyfirst14 {
					r.EncodeMapEntrySeparator()
				} else {
					yyfirst14 = true
				}
				r.EncodeString(codecSelferC_UTF84402, string("prevNode"))
				if yysep14 {
					r.EncodeMapKVSeparator()
				}
				if x.PrevNode == nil {
					r.EncodeNil()
				} else {
					x.PrevNode.CodecEncodeSelf(e)
				}
			}
		}
		if yyr14 || yy2arr14 {
			if yysep14 {
				r.EncodeArrayEntrySeparator()
			}
			r.EncodeUint(uint64(x.EtcdIndex))
		} else {
			if yyfirst14 {
				r.EncodeMapEntrySeparator()
			} else {
				yyfirst14 = true
			}
			r.EncodeString(codecSelferC_UTF84402, string("etcdIndex"))
			if yysep14 {
				r.EncodeMapKVSeparator()
			}
			r.EncodeUint(uint64(x.EtcdIndex))
		}
		if yyr14 || yy2arr14 {
			if yysep14 {
				r.EncodeArrayEntrySeparator()
			}
			r.EncodeUint(uint64(x.RaftIndex))
		} else {
			if yyfirst14 {
				r.EncodeMapEntrySeparator()
			} else {
				yyfirst14 = true
			}
			r.EncodeString(codecSelferC_UTF84402, string("raftIndex"))
			if yysep14 {
				r.EncodeMapKVSeparator()
			}
			r.EncodeUint(uint64(x.RaftIndex))
		}
		if yyr14 || yy2arr14 {
			if yysep14 {
				r.EncodeArrayEntrySeparator()
			}
			r.EncodeUint(uint64(x.RaftTerm))
		} else {
			if yyfirst14 {
				r.EncodeMapEntrySeparator()
			} else {
				yyfirst14 = true
			}
			r.EncodeString(codecSelferC_UTF84402, string("raftTerm"))
			if yysep14 {
				r.EncodeMapKVSeparator()
			}
			r.EncodeUint(uint64(x.RaftTerm))
		}
		if yysep14 {
			if yyr14 || yy2arr14 {
				r.EncodeArrayEnd()
			} else {
				r.EncodeMapEnd()
			}
		}
	}
}

func (x *Response) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer4402
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if r.IsContainerType(codecSelverValueTypeMap4402) {
		yyl21 := r.ReadMapStart()
		if yyl21 == 0 {
			r.ReadMapEnd()
		} else {
			x.codecDecodeSelfFromMap(yyl21, d)
		}
	} else if r.IsContainerType(codecSelverValueTypeArray4402) {
		yyl21 := r.ReadArrayStart()
		if yyl21 == 0 {
			r.ReadArrayEnd()
		} else {
			x.codecDecodeSelfFromArray(yyl21, d)
		}
	} else {
		panic(codecSelferOnlyMapOrArrayEncodeToStructErr4402)
	}
}

func (x *Response) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer4402
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys22Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys22Slc
	var yyhl22 bool = l >= 0
	for yyj22 := 0; ; yyj22++ {
		if yyhl22 {
			if yyj22 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
			if yyj22 > 0 {
				r.ReadMapEntrySeparator()
			}
		}
		yys22Slc = r.DecodeBytes(yys22Slc, true, true)
		yys22 := string(yys22Slc)
		if !yyhl22 {
			r.ReadMapKVSeparator()
		}
		switch yys22 {
		case "action":
			if r.TryDecodeAsNil() {
				x.Action = ""
			} else {
				x.Action = string(r.DecodeString())
			}
		case "node":
			if r.TryDecodeAsNil() {
				if x.Node != nil {
					x.Node = nil
				}
			} else {
				if x.Node == nil {
					x.Node = new(Node)
				}
				x.Node.CodecDecodeSelf(d)
			}
		case "prevNode":
			if r.TryDecodeAsNil() {
				if x.PrevNode != nil {
					x.PrevNode = nil
				}
			} else {
				if x.PrevNode == nil {
					x.PrevNode = new(Node)
				}
				x.PrevNode.CodecDecodeSelf(d)
			}
		case "etcdIndex":
			if r.TryDecodeAsNil() {
				x.EtcdIndex = 0
			} else {
				x.EtcdIndex = uint64(r.DecodeUint(64))
			}
		case "raftIndex":
			if r.TryDecodeAsNil() {
				x.RaftIndex = 0
			} else {
				x.RaftIndex = uint64(r.DecodeUint(64))
			}
		case "raftTerm":
			if r.TryDecodeAsNil() {
				x.RaftTerm = 0
			} else {
				x.RaftTerm = uint64(r.DecodeUint(64))
			}
		default:
			z.DecStructFieldNotFound(-1, yys22)
		} // end switch yys22
	} // end for yyj22
	if !yyhl22 {
		r.ReadMapEnd()
	}
}

func (x *Response) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer4402
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj29 int
	var yyb29 bool
	var yyhl29 bool = l >= 0
	yyj29++
	if yyhl29 {
		yyb29 = yyj29 > l
	} else {
		yyb29 = r.CheckBreak()
	}
	if yyb29 {
		r.ReadArrayEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.Action = ""
	} else {
		x.Action = string(r.DecodeString())
	}
	yyj29++
	if yyhl29 {
		yyb29 = yyj29 > l
	} else {
		yyb29 = r.CheckBreak()
	}
	if yyb29 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayEntrySeparator()
	if r.TryDecodeAsNil() {
		if x.Node != nil {
			x.Node = nil
		}
	} else {
		if x.Node == nil {
			x.Node = new(Node)
		}
		x.Node.CodecDecodeSelf(d)
	}
	yyj29++
	if yyhl29 {
		yyb29 = yyj29 > l
	} else {
		yyb29 = r.CheckBreak()
	}
	if yyb29 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayEntrySeparator()
	if r.TryDecodeAsNil() {
		if x.PrevNode != nil {
			x.PrevNode = nil
		}
	} else {
		if x.PrevNode == nil {
			x.PrevNode = new(Node)
		}
		x.PrevNode.CodecDecodeSelf(d)
	}
	yyj29++
	if yyhl29 {
		yyb29 = yyj29 > l
	} else {
		yyb29 = r.CheckBreak()
	}
	if yyb29 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayEntrySeparator()
	if r.TryDecodeAsNil() {
		x.EtcdIndex = 0
	} else {
		x.EtcdIndex = uint64(r.DecodeUint(64))
	}
	yyj29++
	if yyhl29 {
		yyb29 = yyj29 > l
	} else {
		yyb29 = r.CheckBreak()
	}
	if yyb29 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayEntrySeparator()
	if r.TryDecodeAsNil() {
		x.RaftIndex = 0
	} else {
		x.RaftIndex = uint64(r.DecodeUint(64))
	}
	yyj29++
	if yyhl29 {
		yyb29 = yyj29 > l
	} else {
		yyb29 = r.CheckBreak()
	}
	if yyb29 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayEntrySeparator()
	if r.TryDecodeAsNil() {
		x.RaftTerm = 0
	} else {
		x.RaftTerm = uint64(r.DecodeUint(64))
	}
	for {
		yyj29++
		if yyhl29 {
			yyb29 = yyj29 > l
		} else {
			yyb29 = r.CheckBreak()
		}
		if yyb29 {
			break
		}
		if yyj29 > 1 {
			r.ReadArrayEntrySeparator()
		}
		z.DecStructFieldNotFound(yyj29-1, "")
	}
	r.ReadArrayEnd()
}

func (x *Node) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer4402
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yysep36 := !z.EncBinary()
		yy2arr36 := z.EncBasicHandle().StructToArray
		var yyfirst36 bool
		var yyq36 [8]bool
		_, _, _, _ = yysep36, yyfirst36, yyq36, yy2arr36
		const yyr36 bool = false
		yyq36[1] = x.Value != ""
		yyq36[2] = x.Dir != false
		yyq36[3] = x.Expiration != nil
		yyq36[4] = x.TTL != 0
		yyq36[5] = len(x.Nodes) != 0
		yyq36[6] = x.ModifiedIndex != 0
		yyq36[7] = x.CreatedIndex != 0
		if yyr36 || yy2arr36 {
			r.EncodeArrayStart(8)
		} else {
			var yynn36 int = 1
			for _, b := range yyq36 {
				if b {
					yynn36++
				}
			}
			r.EncodeMapStart(yynn36)
		}
		if yyr36 || yy2arr36 {
			r.EncodeString(codecSelferC_UTF84402, string(x.Key))
		} else {
			yyfirst36 = true
			r.EncodeString(codecSelferC_UTF84402, string("key"))
			if yysep36 {
				r.EncodeMapKVSeparator()
			}
			r.EncodeString(codecSelferC_UTF84402, string(x.Key))
		}
		if yyr36 || yy2arr36 {
			if yysep36 {
				r.EncodeArrayEntrySeparator()
			}
			if yyq36[1] {
				r.EncodeString(codecSelferC_UTF84402, string(x.Value))
			} else {
				r.EncodeString(codecSelferC_UTF84402, "")
			}
		} else {
			if yyq36[1] {
				if yyfirst36 {
					r.EncodeMapEntrySeparator()
				} else {
					yyfirst36 = true
				}
				r.EncodeString(codecSelferC_UTF84402, string("value"))
				if yysep36 {
					r.EncodeMapKVSeparator()
				}
				r.EncodeString(codecSelferC_UTF84402, string(x.Value))
			}
		}
		if yyr36 || yy2arr36 {
			if yysep36 {
				r.EncodeArrayEntrySeparator()
			}
			if yyq36[2] {
				r.EncodeBool(bool(x.Dir))
			} else {
				r.EncodeBool(false)
			}
		} else {
			if yyq36[2] {
				if yyfirst36 {
					r.EncodeMapEntrySeparator()
				} else {
					yyfirst36 = true
				}
				r.EncodeString(codecSelferC_UTF84402, string("dir"))
				if yysep36 {
					r.EncodeMapKVSeparator()
				}
				r.EncodeBool(bool(x.Dir))
			}
		}
		if yyr36 || yy2arr36 {
			if yysep36 {
				r.EncodeArrayEntrySeparator()
			}
			if yyq36[3] {
				if x.Expiration == nil {
					r.EncodeNil()
				} else {
					z.EncFallback(x.Expiration)
				}
			} else {
				r.EncodeNil()
			}
		} else {
			if yyq36[3] {
				if yyfirst36 {
					r.EncodeMapEntrySeparator()
				} else {
					yyfirst36 = true
				}
				r.EncodeString(codecSelferC_UTF84402, string("expiration"))
				if yysep36 {
					r.EncodeMapKVSeparator()
				}
				if x.Expiration == nil {
					r.EncodeNil()
				} else {
					z.EncFallback(x.Expiration)
				}
			}
		}
		if yyr36 || yy2arr36 {
			if yysep36 {
				r.EncodeArrayEntrySeparator()
			}
			if yyq36[4] {
				r.EncodeInt(int64(x.TTL))
			} else {
				r.EncodeInt(0)
			}
		} else {
			if yyq36[4] {
				if yyfirst36 {
					r.EncodeMapEntrySeparator()
				} else {
					yyfirst36 = true
				}
				r.EncodeString(codecSelferC_UTF84402, string("ttl"))
				if yysep36 {
					r.EncodeMapKVSeparator()
				}
				r.EncodeInt(int64(x.TTL))
			}
		}
		if yyr36 || yy2arr36 {
			if yysep36 {
				r.EncodeArrayEntrySeparator()
			}
			if yyq36[5] {
				if x.Nodes == nil {
					r.EncodeNil()
				} else {
					x.Nodes.CodecEncodeSelf(e)
				}
			} else {
				r.EncodeNil()
			}
		} else {
			if yyq36[5] {
				if yyfirst36 {
					r.EncodeMapEntrySeparator()
				} else {
					yyfirst36 = true
				}
				r.EncodeString(codecSelferC_UTF84402, string("nodes"))
				if yysep36 {
					r.EncodeMapKVSeparator()
				}
				if x.Nodes == nil {
					r.EncodeNil()
				} else {
					x.Nodes.CodecEncodeSelf(e)
				}
			}
		}
		if yyr36 || yy2arr36 {
			if yysep36 {
				r.EncodeArrayEntrySeparator()
			}
			if yyq36[6] {
				r.EncodeUint(uint64(x.ModifiedIndex))
			} else {
				r.EncodeUint(0)
			}
		} else {
			if yyq36[6] {
				if yyfirst36 {
					r.EncodeMapEntrySeparator()
				} else {
					yyfirst36 = true
				}
				r.EncodeString(codecSelferC_UTF84402, string("modifiedIndex"))
				if yysep36 {
					r.EncodeMapKVSeparator()
				}
				r.EncodeUint(uint64(x.ModifiedIndex))
			}
		}
		if yyr36 || yy2arr36 {
			if yysep36 {
				r.EncodeArrayEntrySeparator()
			}
			if yyq36[7] {
				r.EncodeUint(uint64(x.CreatedIndex))
			} else {
				r.EncodeUint(0)
			}
		} else {
			if yyq36[7] {
				if yyfirst36 {
					r.EncodeMapEntrySeparator()
				} else {
					yyfirst36 = true
				}
				r.EncodeString(codecSelferC_UTF84402, string("createdIndex"))
				if yysep36 {
					r.EncodeMapKVSeparator()
				}
				r.EncodeUint(uint64(x.CreatedIndex))
			}
		}
		if yysep36 {
			if yyr36 || yy2arr36 {
				r.EncodeArrayEnd()
			} else {
				r.EncodeMapEnd()
			}
		}
	}
}

func (x *Node) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer4402
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if r.IsContainerType(codecSelverValueTypeMap4402) {
		yyl45 := r.ReadMapStart()
		if yyl45 == 0 {
			r.ReadMapEnd()
		} else {
			x.codecDecodeSelfFromMap(yyl45, d)
		}
	} else if r.IsContainerType(codecSelverValueTypeArray4402) {
		yyl45 := r.ReadArrayStart()
		if yyl45 == 0 {
			r.ReadArrayEnd()
		} else {
			x.codecDecodeSelfFromArray(yyl45, d)
		}
	} else {
		panic(codecSelferOnlyMapOrArrayEncodeToStructErr4402)
	}
}

func (x *Node) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer4402
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys46Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys46Slc
	var yyhl46 bool = l >= 0
	for yyj46 := 0; ; yyj46++ {
		if yyhl46 {
			if yyj46 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
			if yyj46 > 0 {
				r.ReadMapEntrySeparator()
			}
		}
		yys46Slc = r.DecodeBytes(yys46Slc, true, true)
		yys46 := string(yys46Slc)
		if !yyhl46 {
			r.ReadMapKVSeparator()
		}
		switch yys46 {
		case "key":
			if r.TryDecodeAsNil() {
				x.Key = ""
			} else {
				x.Key = string(r.DecodeString())
			}
		case "value":
			if r.TryDecodeAsNil() {
				x.Value = ""
			} else {
				x.Value = string(r.DecodeString())
			}
		case "dir":
			if r.TryDecodeAsNil() {
				x.Dir = false
			} else {
				x.Dir = bool(r.DecodeBool())
			}
		case "expiration":
			if r.TryDecodeAsNil() {
				if x.Expiration != nil {
					x.Expiration = nil
				}
			} else {
				if x.Expiration == nil {
					x.Expiration = new(time.Time)
				}
				z.DecFallback(x.Expiration, false)
			}
		case "ttl":
			if r.TryDecodeAsNil() {
				x.TTL = 0
			} else {
				x.TTL = int64(r.DecodeInt(64))
			}
		case "nodes":
			if r.TryDecodeAsNil() {
				x.Nodes = nil
			} else {
				yyv52 := &x.Nodes
				yyv52.CodecDecodeSelf(d)
			}
		case "modifiedIndex":
			if r.TryDecodeAsNil() {
				x.ModifiedIndex = 0
			} else {
				x.ModifiedIndex = uint64(r.DecodeUint(64))
			}
		case "createdIndex":
			if r.TryDecodeAsNil() {
				x.CreatedIndex = 0
			} else {
				x.CreatedIndex = uint64(r.DecodeUint(64))
			}
		default:
			z.DecStructFieldNotFound(-1, yys46)
		} // end switch yys46
	} // end for yyj46
	if !yyhl46 {
		r.ReadMapEnd()
	}
}

func (x *Node) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer4402
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj55 int
	var yyb55 bool
	var yyhl55 bool = l >= 0
	yyj55++
	if yyhl55 {
		yyb55 = yyj55 > l
	} else {
		yyb55 = r.CheckBreak()
	}
	if yyb55 {
		r.ReadArrayEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.Key = ""
	} else {
		x.Key = string(r.DecodeString())
	}
	yyj55++
	if yyhl55 {
		yyb55 = yyj55 > l
	} else {
		yyb55 = r.CheckBreak()
	}
	if yyb55 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayEntrySeparator()
	if r.TryDecodeAsNil() {
		x.Value = ""
	} else {
		x.Value = string(r.DecodeString())
	}
	yyj55++
	if yyhl55 {
		yyb55 = yyj55 > l
	} else {
		yyb55 = r.CheckBreak()
	}
	if yyb55 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayEntrySeparator()
	if r.TryDecodeAsNil() {
		x.Dir = false
	} else {
		x.Dir = bool(r.DecodeBool())
	}
	yyj55++
	if yyhl55 {
		yyb55 = yyj55 > l
	} else {
		yyb55 = r.CheckBreak()
	}
	if yyb55 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayEntrySeparator()
	if r.TryDecodeAsNil() {
		if x.Expiration != nil {
			x.Expiration = nil
		}
	} else {
		if x.Expiration == nil {
			x.Expiration = new(time.Time)
		}
		z.DecFallback(x.Expiration, false)
	}
	yyj55++
	if yyhl55 {
		yyb55 = yyj55 > l
	} else {
		yyb55 = r.CheckBreak()
	}
	if yyb55 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayEntrySeparator()
	if r.TryDecodeAsNil() {
		x.TTL = 0
	} else {
		x.TTL = int64(r.DecodeInt(64))
	}
	yyj55++
	if yyhl55 {
		yyb55 = yyj55 > l
	} else {
		yyb55 = r.CheckBreak()
	}
	if yyb55 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayEntrySeparator()
	if r.TryDecodeAsNil() {
		x.Nodes = nil
	} else {
		yyv61 := &x.Nodes
		yyv61.CodecDecodeSelf(d)
	}
	yyj55++
	if yyhl55 {
		yyb55 = yyj55 > l
	} else {
		yyb55 = r.CheckBreak()
	}
	if yyb55 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayEntrySeparator()
	if r.TryDecodeAsNil() {
		x.ModifiedIndex = 0
	} else {
		x.ModifiedIndex = uint64(r.DecodeUint(64))
	}
	yyj55++
	if yyhl55 {
		yyb55 = yyj55 > l
	} else {
		yyb55 = r.CheckBreak()
	}
	if yyb55 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayEntrySeparator()
	if r.TryDecodeAsNil() {
		x.CreatedIndex = 0
	} else {
		x.CreatedIndex = uint64(r.DecodeUint(64))
	}
	for {
		yyj55++
		if yyhl55 {
			yyb55 = yyj55 > l
		} else {
			yyb55 = r.CheckBreak()
		}
		if yyb55 {
			break
		}
		if yyj55 > 1 {
			r.ReadArrayEntrySeparator()
		}
		z.DecStructFieldNotFound(yyj55-1, "")
	}
	r.ReadArrayEnd()
}

func (x Nodes) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer4402
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		h.encNodes(Nodes(x), e)
	}
}

func (x *Nodes) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer4402
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	h.decNodes((*Nodes)(x), d)
}

func (x codecSelfer4402) enchttp_Header(v http.Header, e *codec1978.Encoder) {
	var h codecSelfer4402
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.EncodeMapStart(len(v))
	yys64 := !z.EncBinary()
	yyj64 := 0
	if yys64 {
		for yyk64, yyv64 := range v {
			if yyj64 > 0 {
				r.EncodeMapEntrySeparator()
			}
			r.EncodeString(codecSelferC_UTF84402, string(yyk64))
			r.EncodeMapKVSeparator()
			if yyv64 == nil {
				r.EncodeNil()
			} else {
				z.F.EncSliceStringV(yyv64, false, e)
			}
			yyj64++
		}
		r.EncodeMapEnd()
	} else {
		for yyk64, yyv64 := range v {
			r.EncodeString(codecSelferC_UTF84402, string(yyk64))
			if yyv64 == nil {
				r.EncodeNil()
			} else {
				z.F.EncSliceStringV(yyv64, false, e)
			}
		}
	}
}

func (x codecSelfer4402) dechttp_Header(v *http.Header, d *codec1978.Decoder) {
	var h codecSelfer4402
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv65 := *v
	yyl65 := r.ReadMapStart()
	if yyv65 == nil {
		if yyl65 > 0 {
			yyv65 = make(map[string][]string, yyl65)
		} else {
			yyv65 = make(map[string][]string) // supports indefinite-length, etc
		}
		*v = yyv65
	}
	if yyl65 > 0 {
		for yyj65 := 0; yyj65 < yyl65; yyj65++ {
			var yymk65 string
			if r.TryDecodeAsNil() {
				yymk65 = ""
			} else {
				yymk65 = string(r.DecodeString())
			}

			yymv65 := yyv65[yymk65]
			if r.TryDecodeAsNil() {
				yymv65 = nil
			} else {
				yyv67 := &yymv65
				z.F.DecSliceStringX(yyv67, false, d)
			}

			if yyv65 != nil {
				yyv65[yymk65] = yymv65
			}
		}
	} else if yyl65 < 0 {
		for yyj65 := 0; !r.CheckBreak(); yyj65++ {
			if yyj65 > 0 {
				r.ReadMapEntrySeparator()
			}
			var yymk65 string
			if r.TryDecodeAsNil() {
				yymk65 = ""
			} else {
				yymk65 = string(r.DecodeString())
			}

			r.ReadMapKVSeparator()
			yymv65 := yyv65[yymk65]
			if r.TryDecodeAsNil() {
				yymv65 = nil
			} else {
				yyv69 := &yymv65
				z.F.DecSliceStringX(yyv69, false, d)
			}

			if yyv65 != nil {
				yyv65[yymk65] = yymv65
			}
		}
		r.ReadMapEnd()
	} // else len==0: TODO: Should we clear map entries?
}

func (x codecSelfer4402) encNodes(v Nodes, e *codec1978.Encoder) {
	var h codecSelfer4402
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.EncodeArrayStart(len(v))
	yys70 := !z.EncBinary()
	if yys70 {
		for yyi70, yyv70 := range v {
			if yyi70 > 0 {
				r.EncodeArrayEntrySeparator()
			}
			if yyv70 == nil {
				r.EncodeNil()
			} else {
				yyv70.CodecEncodeSelf(e)
			}
		}
		r.EncodeArrayEnd()
	} else {
		for _, yyv70 := range v {
			if yyv70 == nil {
				r.EncodeNil()
			} else {
				yyv70.CodecEncodeSelf(e)
			}
		}
	}
}

func (x codecSelfer4402) decNodes(v *Nodes, d *codec1978.Decoder) {
	var h codecSelfer4402
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv71 := *v
	yyh71, yyl71 := z.DecSliceHelperStart()

	var yyc71 bool
	if yyv71 == nil {
		if yyl71 <= 0 {
			yyv71 = make(Nodes, 0)
		} else {
			yyv71 = make(Nodes, yyl71)
		}
		yyc71 = true
	}

	if yyl71 == 0 {
		if len(yyv71) != 0 {
			yyv71 = yyv71[:0]
			yyc71 = true
		}
	} else if yyl71 > 0 {

		yyn71 := yyl71
		if yyl71 > cap(yyv71) {
			yyv71 = make([]*Node, yyl71, yyl71)
			yyc71 = true

		} else if yyl71 != len(yyv71) {
			yyv71 = yyv71[:yyl71]
			yyc71 = true
		}
		yyj71 := 0
		for ; yyj71 < yyn71; yyj71++ {
			if r.TryDecodeAsNil() {
				if yyv71[yyj71] != nil {
					*yyv71[yyj71] = Node{}
				}
			} else {
				if yyv71[yyj71] == nil {
					yyv71[yyj71] = new(Node)
				}
				yyw72 := yyv71[yyj71]
				yyw72.CodecDecodeSelf(d)
			}

		}

	} else {
		for yyj71 := 0; !r.CheckBreak(); yyj71++ {
			if yyj71 >= len(yyv71) {
				yyv71 = append(yyv71, nil) // var yyz71 *Node
				yyc71 = true
			}
			if yyj71 > 0 {
				yyh71.Sep(yyj71)
			}

			if yyj71 < len(yyv71) {
				if r.TryDecodeAsNil() {
					if yyv71[yyj71] != nil {
						*yyv71[yyj71] = Node{}
					}
				} else {
					if yyv71[yyj71] == nil {
						yyv71[yyj71] = new(Node)
					}
					yyw73 := yyv71[yyj71]
					yyw73.CodecDecodeSelf(d)
				}

			} else {
				z.DecSwallow()
			}

		}
		yyh71.End()
	}
	if yyc71 {
		*v = yyv71
	}
}
