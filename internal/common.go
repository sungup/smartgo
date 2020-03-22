package internal

// 7.12.6 Input from the Device to the Host Data Structure
// 7.12.6.1 Overview
type word uint16
type dword [2]word //uint32
type qword [4]word //uint64

type DevIdentify struct {
	// part 1 of 19 (p140)
	generic  word     // 0        General configuration (see 7.12.6.2)
	_        word     // 1        Obsolete
	specific word     // 2        Specific configuration (see 7.12.6.4)
	_        word     // 3        Obsolete
	_        [2]word  // 4..5     Retired
	_        word     // 6        Obsolete
	cfa1     [2]word  // 7..8     Reserved for CFA (see 7.12.6.8)
	_        word     // 9        Retired
	serial   [20]byte // 10..19   Serial number (see 7.12.6.10)
	_        [2]word  // 20..21   Retired
	_        word     // 22       Obsolete
	firmware [8]byte  // 23..26   Firmware revision (see 7.12.6.13)
	model    [40]byte // 27..46   Model number (see 7.12.6.14)

	// part 2 of 19 (p141)
	_            word    // 47       Obsolete
	trusted      word    // 48       Trusted Computing feature set options (see 7.12.6.16)
	capabilities [2]word // 49..50   Compatibility (see 7.12.6.17)
	_            [2]word // 51..52   Obsolete

	// part 3 of 19 (p142)
	word53      word    // 53       See 7.12.6.19
	_           [5]word // 54..58   Obsolete
	word59      word    // 59       See 7.12.6.21
	pATASectors dword   // 60..61   Total number of user addressable logical sector for 28-bit commands (see 7.12.6.22)
	_           word    // 62       Obsolete
	word63      word    // 63       See 7.12.6.24

	// part 4 of 19 (p143)
	word64            word    // 64       See 7.12.6.25
	multiWordTransfer word    // 65       Minimum Multiword DMA transfer cycle time per word (see 7.12.6.26)
	vendorMultiWord   word    // 66       Manufacturer's recommended Multiword DMA transfer cycle time (see 7.12.6.27)
	minPIOTransfer    word    // 67       Minimum PIO transfer cycle time without flow control (see 7.12.6.28)
	minPIOWithIORDY   word    // 68       Minimum PIO transfer cycle time with IORDY (see ATA8-APT) flow control (see 7.12.6.29)
	additional        word    // 69       Additional Supported (see 7.12.6.30)
	_                 word    // 70       Reserved
	identifyPktDev    [4]word // 71..74   Reserved for the IDENTIFY PACKET DEVICE command (see ACS-3)
	queueDepth        word    // 75       Queue Depth (see 7.12.6.33)

	// part 5 of 19 (p144)
	sataCapability word // 76       Serial ATA Capabilities (see 7.12.6.34)
	sataAdditional word // 77       Serial ATA Additional Capabilities (see 7.12.6.35)

	// part 6 of 19 (p145)
	sataSupported word // 78       Serial ATA features supported (see 7.12.6.36)
	sataEnabled   word // 79       Serial ATA features enabled (see 7.12.6.37)

	// part 7 of 19 (p146)
	majorVersion word // 80       Major version number (see 7.12.6.38)
	minorVersion word // 81       Minor version number (see 7.12.6.39)

	// part 8 of 19  (p147)
	// part 9 of 19  (p148)
	// part 10 of 19 (p149)
	// part 11 of 19 (p150)
	featureSet1 [6]word // 82..87   Commands and feature set supported and enabled (see 7.12.6.40 and 7.12.6.41)

	// part 12 of 19 (p151)
	ultraDMAMode word // 88       Ultra DMA modes (see 7.12.6.42)
	word89       word // 89       See 7.12.6.43

	// part 13 of 19 (p152)
	word90         word // 90       See 7.12.6.44
	apmLevel       word // 91       Current APM level value (see 7.12.6.45)
	masterPassword word // 92       Master Password Identifier (see 7.12.6.46)

	// part 14 of 19 (p153)
	hwResetResult word // 93       Hardware reset result (see 7.12.6.47)
	_             word // 94       Obsolete

	// part 15 of 19 (p154)
	streamMinReqSz  word    // 95       Stream Minimum Request Size (see 7.12.6.49)
	streamDMATrTime word    // 96       Streaming Transfer Time - DMA (see 7.12.6.50)
	streamAccLat    word    // 97       Streaming Access Latency - DMA and PIO (see 7.12.6.51)
	streamPerfGran  dword   // 98..99   Streaming Performance Granularity (DWORD) (see 7.12.6.52)
	sataSectors     qword   // 100..103 Number of User Addressable Logical Sectors (QWord) (see 7.12.6.53)
	streamPIOTrTime word    // 104      Streaming Transfer Time - PIO (see 7.12.6.54)
	maxDTSetMngCmd  word    // 105      Maximum number of 512-byte blocks per DATA SET MANAGEMENT command (see 7.5)
	physPerLogSecs  word    // 106      Physical sector size / logical sector size (see 7.12.6.56)
	interSeekDelay  word    // 107      Inter-seek delay for ISO/IEC 7779 standard acoustic testing
	wwName          [4]word // 108..111 World wide name (see 7.12.6.58)
	_               [4]word // 112..115 Reserved
	_               word    // 116      Obsolete
	sectorSize      dword   // 117..118 Logical sector size (DWord) (see 7.12.6.61)

	// part 16 of 19 (p155)
	featureSet2 [2]word // 119..120 Command and feature sets supported and enabled (Continued from words 82..84) (see 7.12.6.40)
	_           [6]word // 121..126 Reserved for expanded supported and enabled settings

	// part 17 of 19 (p156)
	_               word     // 127      Obsolete
	securityStat    word     // 128      Security status (see 7.12.6.66)
	vendor          [31]word // 129..159 Vendor specific
	cfa2            [8]word  // 160..167 Reserved for CFA (see 7.12.6.68)
	formFactor      word     // 168      See 7.12.6.69
	dtSetMngSupport word     // 169      DATA SET MANAGEMENT command support (see 7.12.6.70)
	additionalIdent [4]word  // 170..173 Additional Product Identifier (see 7.12.6.71)
	_               [2]word  // 174..175 Reserved
	mediaSerial     [60]byte // 176..205 Current media serial number (see 7.12.6.73)

	// part 18 of 19 (p157)
	sctCmdTransport word    // 206      SCT Command Transport (see 7.12.6.74)
	_               [2]word // 207..208 Reserved
	alignmentLSecs  word    // 209      Alignment of logical sectors within a physical sector (see 7.12.6.75)
	wrVerifyMod3Cnt dword   // 210..211 Write-Read-Verify Sector Mode 3 Count (DWord) (see 7.12.6.76)
	wrVerifyMod2Cnt dword   // 212..213 Write-Read-Verify Sector Mode 2 Count (DWord) (see 7.12.6.76)
	_               [3]word // 214..216 Obsolete
	rotationRate    word    // 217      Nominal media rotation rate (see 7.12.6.79)
	_               word    // 218      Reserved
	_               word    // 219      Obsolete
	wrVerifyMode    word    // 220      See 7.12.6.82
	_               word    // 221      Reserved

	// part 19 of 19 (p158)
	transportMajor word     // 222      Transport major version number (see 7.12.6.84)
	transportMinor word     // 223      Transport minor version number (see 7.12.6.85)
	_              [6]word  // 224..229 Obsolete
	extAddrSectors qword    // 230..233 Extended Number of User Addressable Sectors (QWord) (see 7.12.6.87)
	minDLMicroOp   word     // 234      Minimum number of 512-byte data blocks per Download Microcode operation (see 7.12.6.88)
	maxDLMicroOp   word     // 235      Maximum number of 512-byte data blocks per Download Microcode operation (see 7.12.6.88)
	_              [19]word // 236..254 Reserved
	integrity      word     // 255      Integrity word (see 7.12.6.91)
}
