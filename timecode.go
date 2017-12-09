package timecode

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var regexpObject *regexp.Regexp

// FrameRate ...
type FrameRate int

const (
	// SMPTE2398 23.98 frame rate. Also known as Film Sync.
	SMPTE2398 = iota
	// SMPTE24 24 fps frame rate.
	SMPTE24
	// SMPTE25 25 fps frame rate. Also known as PAL.
	SMPTE25
	// SMPTE2997DROP 29.97 fps Drop Frame timecode. Used in the NTSC television system.
	SMPTE2997DROP
	// SMPTE2997NONDROP 29.97 fps Non Drop Fram timecode. Used in the NTSC television system.
	SMPTE2997NONDROP
	// SMPTE30 30 fps frame rate.
	SMPTE30
)

func (fr FrameRate) String() string {
	switch fr {
	case SMPTE2398:
		return "SMPTE2398"
	case SMPTE24:
		return "SMPTE24"
	case SMPTE25:
		return "SMPTE25"
	case SMPTE2997DROP:
		return "SMPTE2997DROP"
	case SMPTE2997NONDROP:
		return "SMPTE2997NONDROP"
	case SMPTE30:
		return "SMPTE30"
	default:
		return "UnKnown"
	}
}

// TimeCode ...
type TimeCode struct {
	hh int
	mm int
	ss int
	ff int
}

// GetHours ...
func (tm *TimeCode) GetHours() int {
	return tm.hh
}

// GetMinutes ...
func (tm *TimeCode) GetMinutes() int {
	return tm.mm
}

// GetSeconds ...
func (tm *TimeCode) GetSeconds() int {
	return tm.ss
}

// GetFrames ...
func (tm *TimeCode) GetFrames() int {
	return tm.ff
}

// ParseIntSmpte2398 ...
func (tm *TimeCode) ParseIntSmpte2398() int {
	hours := tm.hh
	minutes := tm.mm
	seconds := tm.ss
	frames := tm.ff
	return frames + (24 * seconds) + (1440 * minutes) + (86400 * hours)
}

// ParseIntSmpte24 ...
func (tm *TimeCode) ParseIntSmpte24() int {
	hours := tm.hh
	minutes := tm.mm
	seconds := tm.ss
	frames := tm.ff
	return frames + (24 * seconds) + (1440 * minutes) + (86400 * hours)
}

// ParseIntSmpte25 ...
func (tm *TimeCode) ParseIntSmpte25() int {
	hours := tm.hh
	minutes := tm.mm
	seconds := tm.ss
	frames := tm.ff
	return frames + (25 * seconds) + (1500 * minutes) + (90000 * hours)
}

// ParseIntSmpte2997Drop ...
func (tm *TimeCode) ParseIntSmpte2997Drop() int {
	hours := tm.hh
	minutes := tm.mm
	seconds := tm.ss
	frames := tm.ff
	// drop frame
	if minutes%10 != 0 && seconds == 0 && (frames == 0 || frames == 1) {
		frames = 2
	}

	return frames + (30 * seconds) + (1798 * minutes) + ((2 * (minutes / 10)) + (107892 * hours))
}

// ParseIntSmpte2997NonDrop ...
func (tm *TimeCode) ParseIntSmpte2997NonDrop() int {
	hours := tm.hh
	minutes := tm.mm
	seconds := tm.ss
	frames := tm.ff
	return frames + (30 * seconds) + (1800 * minutes) + (108000 * hours)
}

// ParseIntSmpte30 ...
func (tm *TimeCode) ParseIntSmpte30() int {
	hours := tm.hh
	minutes := tm.mm
	seconds := tm.ss
	frames := tm.ff
	return frames + (30 * seconds) + (1800 * minutes) + (108000 * hours)
}

// IsValidSmpte2398 ...
func (tm *TimeCode) IsValidSmpte2398() bool {
	if (tm.hh >= 0 && tm.hh <= 23) && (tm.mm >= 0 && tm.mm <= 59) && (tm.ss >= 0 && tm.ss <= 59) && (tm.ff >= 0 && tm.ff <= 23) {
		return true
	}
	return false
}

// IsValidSmpte24 ...
func (tm *TimeCode) IsValidSmpte24() bool {
	if (tm.hh >= 0 && tm.hh <= 23) && (tm.mm >= 0 && tm.mm <= 59) && (tm.ss >= 0 && tm.ss <= 59) && (tm.ff >= 0 && tm.ff <= 23) {
		return true
	}
	return false
}

// IsValidSmpte25 ...
func (tm *TimeCode) IsValidSmpte25() bool {
	if (tm.hh >= 0 && tm.hh <= 23) && (tm.mm >= 0 && tm.mm <= 59) && (tm.ss >= 0 && tm.ss <= 59) && (tm.ff >= 0 && tm.ff <= 24) {
		return true
	}
	return false
}

// IsValidSmpte2997Drop ...
func (tm *TimeCode) IsValidSmpte2997Drop() bool {
	if (tm.hh >= 0 && tm.hh <= 23) && (tm.mm >= 0 && tm.mm <= 59) && (tm.ss >= 0 && tm.ss <= 59) && (tm.ff >= 0 && tm.ff <= 29) {
		return true
	}
	return false
}

// IsValidSmpte2997NonDrop ...
func (tm *TimeCode) IsValidSmpte2997NonDrop() bool {
	if (tm.hh >= 0 && tm.hh <= 23) && (tm.mm >= 0 && tm.mm <= 59) && (tm.ss >= 0 && tm.ss <= 59) && (tm.ff >= 0 && tm.ff <= 29) {
		return true
	}
	return false
}

// IsValidSmpte30 ...
func (tm *TimeCode) IsValidSmpte30() bool {
	if (tm.hh >= 0 && tm.hh <= 23) && (tm.mm >= 0 && tm.mm <= 59) && (tm.ss >= 0 && tm.ss <= 59) && (tm.ff >= 0 && tm.ff <= 29) {
		return true
	}
	return false
}

// IsValid ...
func (tm *TimeCode) IsValid(frameRate FrameRate) bool {
	if frameRate == SMPTE2398 {
		if tm.IsValidSmpte2398() {
			return true
		}
		return false
	}

	if frameRate == SMPTE24 {
		if tm.IsValidSmpte24() {
			return true
		}
		return false
	}

	if frameRate == SMPTE25 {
		if tm.IsValidSmpte25() {
			return true
		}
		return false
	}

	if frameRate == SMPTE2997DROP {
		if tm.IsValidSmpte2997Drop() {
			return true
		}
		return false
	}

	if frameRate == SMPTE2997NONDROP {
		if tm.IsValidSmpte2997NonDrop() {
			return true
		}
		return false
	}

	if frameRate == SMPTE30 {
		if tm.IsValidSmpte30() {
			return true
		}
		return false
	}

	return false
}

func init() {
	expr := "^(?P<hours>\\d{1,2}):(?P<minutes>\\d{1,2}):(?P<seconds>\\d{1,2})(?::|;|\\.)(?P<frames>\\d{1,2})$"
	regexpObject, _ = regexp.Compile(expr)
}

// IsValid ...
func IsValid(s string) bool {
	_, err := parseSmpte(s)
	if err != nil {
		return false
	}

	return true
}

// ParseTimeCode ...
func ParseTimeCode(s string, frameRate FrameRate) (*TimeCode, error) {
	return parseTimeCode(s, frameRate)
}

func parseTimeCode(s string, frameRate FrameRate) (*TimeCode, error) {
	tm, err := parseSmpte(s)
	if err != nil {
		return nil, err
	}

	ok := tm.IsValid(frameRate)
	if ok {
		return tm, nil
	}

	return nil, errors.New("value out of range")
}

func parseSmpte(s string) (*TimeCode, error) {
	match := regexpObject.FindStringSubmatch(s)
	if match == nil {
		return nil, errors.New("invalid value")
	}

	captures := make(map[string]string)
	//for i, name := range r.SubexpNames() {
	for i, name := range regexpObject.SubexpNames() {
		if i > 0 && i <= len(match) {
			captures[name] = match[i]
		}
	}

	var err error
	var hh int
	if hh, err = strconv.Atoi(captures["hours"]); err != nil {
		return nil, err
	}

	var mm int
	if mm, err = strconv.Atoi(captures["minutes"]); err != nil {
		return nil, err
	}

	var ss int
	if ss, err = strconv.Atoi(captures["seconds"]); err != nil {
		return nil, err
	}

	var ff int
	if ff, err = strconv.Atoi(captures["frames"]); err != nil {
		return nil, err
	}

	return &TimeCode{hh, mm, ss, ff}, nil
}

// FormatString ...
func FormatString(framecount int, frameRate FrameRate) string {
	if frameRate == SMPTE2398 {
		return formatTimeCodeSmpte2398(framecount)
	}

	if frameRate == SMPTE24 {
		return formatTimeCodeSmpte24(framecount)
	}

	if frameRate == SMPTE25 {
		return formatTimeCodeSmpte25(framecount)
	}

	if frameRate == SMPTE2997DROP {
		return formatTimeCodeSmpte2997Drop(framecount)
	}

	if frameRate == SMPTE2997NONDROP {
		return formatTimeCodeSmpte2997NonDrop(framecount)
	}

	if frameRate == SMPTE30 {
		return formatTimeCodeSmpte30(framecount)
	}

	return ""
}

func formatTimeCodeSmpte2398(framecount int) string {
	hours := (framecount / 86400) % 24
	minutes := ((framecount - (86400 * hours)) / 1440) % 60
	seconds := ((framecount - (1440 * minutes) - (86400 * hours)) / 24) % 3600
	frames := (framecount - (24 * seconds) - (1440 * minutes) - (86400 * hours)) % 24
	return formatTimeCodeString(hours, minutes, seconds, frames)
}

func formatTimeCodeSmpte24(framecount int) string {
	hours := (framecount / 86400) % 24
	minutes := ((framecount - (86400 * hours)) / 1440) % 60
	seconds := ((framecount - (1440 * minutes) - (86400 * hours)) / 24) % 3600
	frames := (framecount - (24 * seconds) - (1440 * minutes) - (86400 * hours)) % 24
	return formatTimeCodeString(hours, minutes, seconds, frames)
}

func formatTimeCodeSmpte25(framecount int) string {
	hours := (framecount / 90000) % 24
	minutes := ((framecount - (90000 * hours)) / 1500) % 60
	seconds := ((framecount - (1500 * minutes) - (90000 * hours)) / 25) % 3600
	frames := (framecount - (25 * seconds) - (1500 * minutes) - (90000 * hours)) % 25
	return formatTimeCodeString(hours, minutes, seconds, frames)
}

func formatTimeCodeSmpte2997Drop(framecount int) string {
	hours := (framecount / 107892) % 24
	minutes := (framecount + (2 * ((framecount - (107892 * hours)) / 1800)) - (2 * ((framecount - (107892 * hours)) / 18000)) - (107892 * hours)) / 1800
	seconds := (framecount - (1798 * minutes) - (2 * (minutes / 10)) - (107892 * hours)) / 30
	frames := framecount - (30 * seconds) - (1798 * minutes) - (2 * (minutes / 10)) - (107892 * hours)
	return fmt.Sprintf("%02d:%02d:%02d;%02d", hours, minutes, seconds, frames)
}

func formatTimeCodeSmpte2997NonDrop(framecount int) string {
	hours := (framecount / 108000) % 24
	minutes := ((framecount - (108000 * hours)) / 1800) % 60
	seconds := ((framecount - (1800 * minutes) - (108000 * hours)) / 30) % 3600
	frames := (framecount - (30 * seconds) - (1800 * minutes) - (108000 * hours)) % 30
	return formatTimeCodeString(hours, minutes, seconds, frames)
}

func formatTimeCodeSmpte30(framecount int) string {
	hours := (framecount / 108000) % 24
	minutes := ((framecount - (108000 * hours)) / 1800) % 60
	seconds := ((framecount - (1800 * minutes) - (108000 * hours)) / 30) % 3600
	frames := (framecount - (30 * seconds) - (1800 * minutes) - (108000 * hours)) % 30
	return formatTimeCodeString(hours, minutes, seconds, frames)
}

func formatTimeCodeString(hours, minutes, seconds, frames int) string {
	return fmt.Sprintf("%02d:%02d:%02d:%02d", hours, minutes, seconds, frames)
}

// ParseInt ...
func ParseInt(s string, frameRate FrameRate) (int, error) {
	tm, err := parseSmpte(s)
	if err != nil {
		return 0, err
	}

	if frameRate == SMPTE2398 {
		ok := tm.IsValidSmpte2398()
		if ok {
			return tm.ParseIntSmpte2398(), nil
		}
		return 0, errors.New("value out of range")
	}

	if frameRate == SMPTE24 {
		ok := tm.IsValidSmpte24()
		if ok {
			return tm.ParseIntSmpte24(), nil
		}
		return 0, errors.New("value out of range")
	}

	if frameRate == SMPTE25 {
		ok := tm.IsValidSmpte25()
		if ok {
			return tm.ParseIntSmpte25(), nil
		}
		return 0, errors.New("value out of range")
	}

	if frameRate == SMPTE2997DROP {
		ok := tm.IsValidSmpte2997Drop()
		if ok {
			return tm.ParseIntSmpte2997Drop(), nil
		}
		return 0, errors.New("value out of range")
	}

	if frameRate == SMPTE2997NONDROP {
		ok := tm.IsValidSmpte2997NonDrop()
		if ok {
			return tm.ParseIntSmpte2997NonDrop(), nil
		}
		return 0, errors.New("value out of range")
	}

	if frameRate == SMPTE30 {
		ok := tm.IsValidSmpte30()
		if ok {
			return tm.ParseIntSmpte30(), nil
		}
		return 0, errors.New("value out of range")
	}

	return 0, errors.New("invalid frame rate")
}

// ParseDuration ...
func ParseDuration(in, out string, frameRate FrameRate) (int, error) {
	tmIn, err := parseSmpte(in)
	if err != nil {
		return 0, err
	}

	tmOut, err := parseSmpte(out)
	if err != nil {
		return 0, err
	}

	if frameRate == SMPTE2398 {
		return parseDurationSmpte2398(tmIn, tmOut)
	}

	if frameRate == SMPTE24 {
		return parseDurationSmpte24(tmIn, tmOut)
	}

	if frameRate == SMPTE25 {
		return parseDurationSmpte25(tmIn, tmOut)
	}

	if frameRate == SMPTE2997DROP {
		return parseDurationSmpte2997Drop(tmIn, tmOut)
	}

	if frameRate == SMPTE2997NONDROP {
		return parseDurationSmpte2997NonDrop(tmIn, tmOut)
	}

	if frameRate == SMPTE30 {
		return parseDurationSmpte30(tmIn, tmOut)
	}

	return 0, errors.New("invalid frame rate")
}

func parseDurationSmpte2398(in, out *TimeCode) (int, error) {
	if ok := in.IsValidSmpte2398(); !ok {
		return 0, errors.New("value out of range")
	}

	if ok := out.IsValidSmpte2398(); !ok {
		return 0, errors.New("value out of range")
	}

	inVal := in.ParseIntSmpte2398()
	outVal := out.ParseIntSmpte2398()

	if inVal > outVal {
		return (2073600 + (outVal % 2073600)) - inVal, nil
	}

	return outVal - inVal, nil
}

func parseDurationSmpte24(in, out *TimeCode) (int, error) {
	if ok := in.IsValidSmpte24(); !ok {
		return 0, errors.New("value out of range")
	}

	if ok := out.IsValidSmpte24(); !ok {
		return 0, errors.New("value out of range")
	}

	inVal := in.ParseIntSmpte24()
	outVal := out.ParseIntSmpte24()

	if inVal > outVal {
		return (2073600 + (outVal % 2073600)) - inVal, nil
	}

	return outVal - inVal, nil
}

func parseDurationSmpte25(in, out *TimeCode) (int, error) {
	if ok := in.IsValidSmpte25(); !ok {
		return 0, errors.New("value out of range")
	}

	if ok := out.IsValidSmpte25(); !ok {
		return 0, errors.New("value out of range")
	}

	inVal := in.ParseIntSmpte25()
	outVal := out.ParseIntSmpte25()

	if inVal > outVal {
		return (2160000 + (outVal % 2160000)) - inVal, nil
	}

	return outVal - inVal, nil
}

func parseDurationSmpte2997Drop(in, out *TimeCode) (int, error) {
	if ok := in.IsValidSmpte2997Drop(); !ok {
		return 0, errors.New("value out of range")
	}

	if ok := out.IsValidSmpte2997Drop(); !ok {
		return 0, errors.New("value out of range")
	}

	inVal := in.ParseIntSmpte2997Drop()
	outVal := out.ParseIntSmpte2997Drop()

	if inVal > outVal {
		return (2589408 + (outVal % 2589408)) - inVal, nil
	}

	return outVal - inVal, nil
}

func parseDurationSmpte2997NonDrop(in, out *TimeCode) (int, error) {
	if ok := in.IsValidSmpte2997NonDrop(); ok == false {
		return 0, errors.New("value out of range")
	}

	if ok := out.IsValidSmpte2997NonDrop(); ok == false {
		return 0, errors.New("value out of range")
	}

	inVal := in.ParseIntSmpte2997NonDrop()
	outVal := out.ParseIntSmpte2997NonDrop()

	if inVal > outVal {
		return (2592000 + (outVal % 2592000)) - inVal, nil
	}

	return outVal - inVal, nil
}

func parseDurationSmpte30(in, out *TimeCode) (int, error) {
	if ok := in.IsValidSmpte30(); !ok {
		return 0, errors.New("value out of range")
	}

	if ok := out.IsValidSmpte30(); !ok {
		return 0, errors.New("value out of range")
	}

	inVal := in.ParseIntSmpte30()
	outVal := out.ParseIntSmpte30()

	if inVal > outVal {
		return (2592000 + (outVal % 2592000)) - inVal, nil
	}

	return outVal - inVal, nil
}
