// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.2.0
// source: Configuration.proto

package main

import (
	fmt "fmt"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	bits "math/bits"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *Configuration) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Configuration) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Configuration) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.Product) > 0 {
		i -= len(m.Product)
		copy(dAtA[i:], m.Product)
		i = encodeVarint(dAtA, i, uint64(len(m.Product)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xca
	}
	if m.SdkVersion != 0 {
		i = encodeVarint(dAtA, i, uint64(m.SdkVersion))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xc0
	}
	if m.Navigation != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Navigation))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xb8
	}
	if m.NavHidden != 0 {
		i = encodeVarint(dAtA, i, uint64(m.NavHidden))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xb0
	}
	if m.Keyboard != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Keyboard))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa8
	}
	if m.KeysHidden != 0 {
		i = encodeVarint(dAtA, i, uint64(m.KeysHidden))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa0
	}
	if m.Touchscreen != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Touchscreen))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x98
	}
	if m.Density != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Density))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x90
	}
	if m.UiModeNight != 0 {
		i = encodeVarint(dAtA, i, uint64(m.UiModeNight))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x88
	}
	if m.UiModeType != 0 {
		i = encodeVarint(dAtA, i, uint64(m.UiModeType))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.Orientation != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Orientation))
		i--
		dAtA[i] = 0x78
	}
	if m.Hdr != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Hdr))
		i--
		dAtA[i] = 0x70
	}
	if m.WideColorGamut != 0 {
		i = encodeVarint(dAtA, i, uint64(m.WideColorGamut))
		i--
		dAtA[i] = 0x68
	}
	if m.ScreenRound != 0 {
		i = encodeVarint(dAtA, i, uint64(m.ScreenRound))
		i--
		dAtA[i] = 0x60
	}
	if m.ScreenLayoutLong != 0 {
		i = encodeVarint(dAtA, i, uint64(m.ScreenLayoutLong))
		i--
		dAtA[i] = 0x58
	}
	if m.ScreenLayoutSize != 0 {
		i = encodeVarint(dAtA, i, uint64(m.ScreenLayoutSize))
		i--
		dAtA[i] = 0x50
	}
	if m.SmallestScreenWidthDp != 0 {
		i = encodeVarint(dAtA, i, uint64(m.SmallestScreenWidthDp))
		i--
		dAtA[i] = 0x48
	}
	if m.ScreenHeightDp != 0 {
		i = encodeVarint(dAtA, i, uint64(m.ScreenHeightDp))
		i--
		dAtA[i] = 0x40
	}
	if m.ScreenWidthDp != 0 {
		i = encodeVarint(dAtA, i, uint64(m.ScreenWidthDp))
		i--
		dAtA[i] = 0x38
	}
	if m.ScreenHeight != 0 {
		i = encodeVarint(dAtA, i, uint64(m.ScreenHeight))
		i--
		dAtA[i] = 0x30
	}
	if m.ScreenWidth != 0 {
		i = encodeVarint(dAtA, i, uint64(m.ScreenWidth))
		i--
		dAtA[i] = 0x28
	}
	if m.LayoutDirection != 0 {
		i = encodeVarint(dAtA, i, uint64(m.LayoutDirection))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Locale) > 0 {
		i -= len(m.Locale)
		copy(dAtA[i:], m.Locale)
		i = encodeVarint(dAtA, i, uint64(len(m.Locale)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Mnc != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Mnc))
		i--
		dAtA[i] = 0x10
	}
	if m.Mcc != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Mcc))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarint(dAtA []byte, offset int, v uint64) int {
	offset -= sov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Configuration) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Mcc != 0 {
		n += 1 + sov(uint64(m.Mcc))
	}
	if m.Mnc != 0 {
		n += 1 + sov(uint64(m.Mnc))
	}
	l = len(m.Locale)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.LayoutDirection != 0 {
		n += 1 + sov(uint64(m.LayoutDirection))
	}
	if m.ScreenWidth != 0 {
		n += 1 + sov(uint64(m.ScreenWidth))
	}
	if m.ScreenHeight != 0 {
		n += 1 + sov(uint64(m.ScreenHeight))
	}
	if m.ScreenWidthDp != 0 {
		n += 1 + sov(uint64(m.ScreenWidthDp))
	}
	if m.ScreenHeightDp != 0 {
		n += 1 + sov(uint64(m.ScreenHeightDp))
	}
	if m.SmallestScreenWidthDp != 0 {
		n += 1 + sov(uint64(m.SmallestScreenWidthDp))
	}
	if m.ScreenLayoutSize != 0 {
		n += 1 + sov(uint64(m.ScreenLayoutSize))
	}
	if m.ScreenLayoutLong != 0 {
		n += 1 + sov(uint64(m.ScreenLayoutLong))
	}
	if m.ScreenRound != 0 {
		n += 1 + sov(uint64(m.ScreenRound))
	}
	if m.WideColorGamut != 0 {
		n += 1 + sov(uint64(m.WideColorGamut))
	}
	if m.Hdr != 0 {
		n += 1 + sov(uint64(m.Hdr))
	}
	if m.Orientation != 0 {
		n += 1 + sov(uint64(m.Orientation))
	}
	if m.UiModeType != 0 {
		n += 2 + sov(uint64(m.UiModeType))
	}
	if m.UiModeNight != 0 {
		n += 2 + sov(uint64(m.UiModeNight))
	}
	if m.Density != 0 {
		n += 2 + sov(uint64(m.Density))
	}
	if m.Touchscreen != 0 {
		n += 2 + sov(uint64(m.Touchscreen))
	}
	if m.KeysHidden != 0 {
		n += 2 + sov(uint64(m.KeysHidden))
	}
	if m.Keyboard != 0 {
		n += 2 + sov(uint64(m.Keyboard))
	}
	if m.NavHidden != 0 {
		n += 2 + sov(uint64(m.NavHidden))
	}
	if m.Navigation != 0 {
		n += 2 + sov(uint64(m.Navigation))
	}
	if m.SdkVersion != 0 {
		n += 2 + sov(uint64(m.SdkVersion))
	}
	l = len(m.Product)
	if l > 0 {
		n += 2 + l + sov(uint64(l))
	}
	if m.unknownFields != nil {
		n += len(m.unknownFields)
	}
	return n
}

func sov(x uint64) (n int) {
	return (bits.Len64(x|1) + 6) / 7
}
func soz(x uint64) (n int) {
	return sov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Configuration) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Configuration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Configuration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mcc", wireType)
			}
			m.Mcc = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mcc |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mnc", wireType)
			}
			m.Mnc = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mnc |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Locale", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Locale = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LayoutDirection", wireType)
			}
			m.LayoutDirection = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LayoutDirection |= Configuration_LayoutDirection(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScreenWidth", wireType)
			}
			m.ScreenWidth = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ScreenWidth |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScreenHeight", wireType)
			}
			m.ScreenHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ScreenHeight |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScreenWidthDp", wireType)
			}
			m.ScreenWidthDp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ScreenWidthDp |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScreenHeightDp", wireType)
			}
			m.ScreenHeightDp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ScreenHeightDp |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SmallestScreenWidthDp", wireType)
			}
			m.SmallestScreenWidthDp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SmallestScreenWidthDp |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScreenLayoutSize", wireType)
			}
			m.ScreenLayoutSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ScreenLayoutSize |= Configuration_ScreenLayoutSize(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScreenLayoutLong", wireType)
			}
			m.ScreenLayoutLong = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ScreenLayoutLong |= Configuration_ScreenLayoutLong(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScreenRound", wireType)
			}
			m.ScreenRound = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ScreenRound |= Configuration_ScreenRound(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WideColorGamut", wireType)
			}
			m.WideColorGamut = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WideColorGamut |= Configuration_WideColorGamut(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hdr", wireType)
			}
			m.Hdr = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Hdr |= Configuration_Hdr(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Orientation", wireType)
			}
			m.Orientation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Orientation |= Configuration_Orientation(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UiModeType", wireType)
			}
			m.UiModeType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UiModeType |= Configuration_UiModeType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UiModeNight", wireType)
			}
			m.UiModeNight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UiModeNight |= Configuration_UiModeNight(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 18:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Density", wireType)
			}
			m.Density = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Density |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 19:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Touchscreen", wireType)
			}
			m.Touchscreen = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Touchscreen |= Configuration_Touchscreen(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 20:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeysHidden", wireType)
			}
			m.KeysHidden = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeysHidden |= Configuration_KeysHidden(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 21:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keyboard", wireType)
			}
			m.Keyboard = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Keyboard |= Configuration_Keyboard(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 22:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NavHidden", wireType)
			}
			m.NavHidden = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NavHidden |= Configuration_NavHidden(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 23:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Navigation", wireType)
			}
			m.Navigation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Navigation |= Configuration_Navigation(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 24:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SdkVersion", wireType)
			}
			m.SdkVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SdkVersion |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 25:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Product", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Product = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skip(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflow
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflow
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflow
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLength
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroup
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLength
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLength        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflow          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroup = fmt.Errorf("proto: unexpected end of group")
)