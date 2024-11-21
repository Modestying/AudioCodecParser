package opus

import "fmt"

type InfoOpt func(*OpusData) string

var OptRecord map[string]InfoOpt

func init() {
	OptRecord = make(map[string]InfoOpt, 10)
	OptRecord["ver"] = WithOpusVersion
	OptRecord["tf"] = WithTypeFlag
	OptRecord["gp"] = WithGranulePosition
	OptRecord["bsn"] = WithBitStreamSerialNumber
	OptRecord["psn"] = WithPageSeqNumber
	OptRecord["crc"] = WithCrcCheck
	OptRecord["sc"] = WithNumberPageSegments
	OptRecord["st"] = WithSegmentTable
	OptRecord["data"] = WithData
}

func WithOpusVersion(d *OpusData) string {
	return fmt.Sprintf("\tVersion:%X\n", d.OggVersion)
}
func WithTypeFlag(d *OpusData) string {
	return fmt.Sprintf("\tTypeFlag:%X\n", d.HeaderTypeFlag)
}
func WithGranulePosition(d *OpusData) string {
	return fmt.Sprintf("\tGranulePosition:%x\n", d.GranulePosition)
}

func WithBitStreamSerialNumber(d *OpusData) string {
	return fmt.Sprintf("\tBitstreamSerialNumber:%x\n", d.BitstreamSerialNumber)
}

func WithPageSeqNumber(d *OpusData) string {
	return fmt.Sprintf("\tPageSeqNumber:%x\n", d.PageSeqNumber)
}

func WithCrcCheck(d *OpusData) string {
	return fmt.Sprintf("\tCrcCheck:%x\n", d.CrcCheck)
}

func WithNumberPageSegments(d *OpusData) string {
	return fmt.Sprintf("\tNumberPageSegments:%X\n", d.NumberPageSegments)
}

func WithSegmentTable(d *OpusData) string {
	return fmt.Sprintf("\tSegmentTable:%x\n", d.SegmentTable)
}

func WithData(d *OpusData) string {
	str := ""
	for i, data := range d.Data {
		str += fmt.Sprintf("\t\tSeg(%d):%x\n", i, data)
	}
	return str
}
func (p *OpusData) Info(opts ...InfoOpt) string {
	str := ""
	for _, opt := range opts {
		str += opt(p)
	}
	return str
}
func (p *OpusData) String() string {
	str := fmt.Sprintf(
		"Magic:%x\n"+
			"\tVersion:%x\n"+
			"\tTypeFlag:%x\n"+
			"\tGranulePosition:%x\n"+
			"\tBitstreamSerialNumber:%x\n"+
			"\tPageSeqNumber:%x\n"+
			"\tCrcCheck:%x\n"+
			"\tNumberPageSegments:%x\n"+
			"\tSegmentTable:%x\n",
		p.OggMagic,
		p.OggVersion,
		p.HeaderTypeFlag,
		p.GranulePosition,
		p.BitstreamSerialNumber,
		p.PageSeqNumber,
		p.CrcCheck,
		p.NumberPageSegments,
		p.SegmentTable,
	)
	for i, data := range p.Data {
		str += fmt.Sprintf("\t\tSeg(%d):%x\n", i, data)
	}
	return str
}
