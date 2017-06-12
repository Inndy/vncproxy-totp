package vnc

import "io"

// RawEncoding is raw pixel data sent by the server.
//
// See RFC 6143 Section 7.7.1
type RawEncoding struct {
	Colors []Color
}

func (*RawEncoding) Type() int32 {
	return 0
}



func (*RawEncoding) Read(conn *ClientConn, rect *Rectangle, r io.Reader) (Encoding, error) {
	//conn := &DataSource{conn: conn.c, PixelFormat: conn.PixelFormat}

	bytesPerPixel := int(conn.PixelFormat.BPP / 8)
	//pixelBytes := make([]uint8, bytesPerPixel)

	// var byteOrder binary.ByteOrder = binary.LittleEndian
	// if conn.PixelFormat.BigEndian {
	// 	byteOrder = binary.BigEndian
	// }

	colors := make([]Color, int(rect.Height)*int(rect.Width))

	for y := uint16(0); y < rect.Height; y++ {
		for x := uint16(0); x < rect.Width; x++ {
			if _, err := conn.readBytes(bytesPerPixel); err != nil {
				return nil, err
			}

			// var rawPixel uint32
			// if conn.PixelFormat.BPP == 8 {
			// 	rawPixel = uint32(pixelBytes[0])
			// } else if conn.PixelFormat.BPP == 16 {
			// 	rawPixel = uint32(byteOrder.Uint16(pixelBytes))
			// } else if conn.PixelFormat.BPP == 32 {
			// 	rawPixel = byteOrder.Uint32(pixelBytes)
			// }

			// color := &colors[int(y)*int(rect.Width)+int(x)]
			// if conn.PixelFormat.TrueColor {
			// 	color.R = uint16((rawPixel >> conn.PixelFormat.RedShift) & uint32(conn.PixelFormat.RedMax))
			// 	color.G = uint16((rawPixel >> conn.PixelFormat.GreenShift) & uint32(conn.PixelFormat.GreenMax))
			// 	color.B = uint16((rawPixel >> conn.PixelFormat.BlueShift) & uint32(conn.PixelFormat.BlueMax))
			// } else {
			// 	*color = conn.ColorMap[rawPixel]
			// }
		}
	}

	return &RawEncoding{colors}, nil
}