package opus

import "io"

type OpusData struct {
	OggMagic              [4]byte
	OggVersion            [1]byte
	HeaderTypeFlag        [1]byte
	GranulePosition       [8]byte
	BitstreamSerialNumber [4]byte
	PageSeqNumber         [4]byte
	CrcCheck              [4]byte
	NumberPageSegments    [1]byte
	SegmentTable          []byte
	Data                  [][]byte
}

func (p *OpusData) Parse(reader io.Reader) {
	io.ReadFull(reader, p.OggMagic[:])
	io.ReadFull(reader, p.OggVersion[:])
	io.ReadFull(reader, p.HeaderTypeFlag[:])
	io.ReadFull(reader, p.GranulePosition[:])
	io.ReadFull(reader, p.BitstreamSerialNumber[:])
	io.ReadFull(reader, p.PageSeqNumber[:])
	io.ReadFull(reader, p.CrcCheck[:])
	io.ReadFull(reader, p.NumberPageSegments[:])
	segmentTableLen := int(p.NumberPageSegments[0])
	p.SegmentTable = make([]byte, segmentTableLen)
	_, err := io.ReadFull(reader, p.SegmentTable)
	if err != nil {
		panic(err)
	}
	var segLen uint8 = 0
	for i := 0; i < len(p.SegmentTable); i++ {
		segLen = p.SegmentTable[i]
		if segLen == 0xff {
			continue
		}
		if i > 1 && p.SegmentTable[i-1] == 0xff {
			segLen = 0xff + p.SegmentTable[i]
		}
		data := make([]byte, segLen)
		io.ReadFull(reader, data)
		p.Data = append(p.Data, data)
	}
}
