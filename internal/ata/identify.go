package ata

import "github.com/sungup/smartgo/pkg"

type DevIdentify struct {
	// part 1 of 19 (p140)
	generic  pkg.Word    // 0        General configuration (see 7.12.6.2)
	_        pkg.Word    // 1        Obsolete
	specific pkg.Word    // 2        Specific configuration (see 7.12.6.4)
	_        pkg.Word    // 3        Obsolete
	_        [2]pkg.Word // 4..5     Retired
	_        pkg.Word    // 6        Obsolete
	cfa1     [2]pkg.Word // 7..8     Reserved for CFA (see 7.12.6.8)
	_        pkg.Word    // 9        Retired
	serial   [20]byte    // 10..19   Serial number (see 7.12.6.10)
	_        [2]pkg.Word // 20..21   Retired
	_        pkg.Word    // 22       Obsolete
	firmware [8]byte     // 23..26   Firmware revision (see 7.12.6.13)
	model    [40]byte    // 27..46   Model number (see 7.12.6.14)

	// part 2 of 19 (p141)
	_            pkg.Word    // 47       Obsolete
	trusted      pkg.Word    // 48       Trusted Computing feature set options (see 7.12.6.16)
	capabilities [2]pkg.Word // 49..50   Compatibility (see 7.12.6.17)
	_            [2]pkg.Word // 51..52   Obsolete

	// part 3 of 19 (p142)
	word53      pkg.Word    // 53       See 7.12.6.19
	_           [5]pkg.Word // 54..58   Obsolete
	word59      pkg.Word    // 59       See 7.12.6.21
	pATASectors pkg.DWord   // 60..61   Total number of user addressable logical sector for 28-bit commands (see 7.12.6.22)
	_           pkg.Word    // 62       Obsolete
	word63      pkg.Word    // 63       See 7.12.6.24

	// part 4 of 19 (p143)
	word64            pkg.Word    // 64       See 7.12.6.25
	multiWordTransfer pkg.Word    // 65       Minimum Multiword DMA transfer cycle time per Word (see 7.12.6.26)
	vendorMultiWord   pkg.Word    // 66       Manufacturer's recommended Multiword DMA transfer cycle time (see 7.12.6.27)
	minPIOTransfer    pkg.Word    // 67       Minimum PIO transfer cycle time without flow control (see 7.12.6.28)
	minPIOWithIORDY   pkg.Word    // 68       Minimum PIO transfer cycle time with IORDY (see ATA8-APT) flow control (see 7.12.6.29)
	additional        pkg.Word    // 69       Additional Supported (see 7.12.6.30)
	_                 pkg.Word    // 70       Reserved
	identifyPktDev    [4]pkg.Word // 71..74   Reserved for the IDENTIFY PACKET DEVICE command (see ACS-3)
	queueDepth        pkg.Word    // 75       Queue Depth (see 7.12.6.33)

	// part 5 of 19 (p144)
	sataCapability pkg.Word // 76       Serial ATA Capabilities (see 7.12.6.34)
	sataAdditional pkg.Word // 77       Serial ATA Additional Capabilities (see 7.12.6.35)

	// part 6 of 19 (p145)
	sataSupported pkg.Word // 78       Serial ATA features supported (see 7.12.6.36)
	sataEnabled   pkg.Word // 79       Serial ATA features enabled (see 7.12.6.37)

	// part 7 of 19 (p146)
	majorVersion pkg.Word // 80       Major version number (see 7.12.6.38)
	minorVersion pkg.Word // 81       Minor version number (see 7.12.6.39)

	// part 8 of 19  (p147)
	// part 9 of 19  (p148)
	// part 10 of 19 (p149)
	// part 11 of 19 (p150)
	featureSet1 [6]pkg.Word // 82..87   Commands and feature set supported and enabled (see 7.12.6.40 and 7.12.6.41)

	// part 12 of 19 (p151)
	ultraDMAMode pkg.Word // 88       Ultra DMA modes (see 7.12.6.42)
	word89       pkg.Word // 89       See 7.12.6.43

	// part 13 of 19 (p152)
	word90         pkg.Word // 90       See 7.12.6.44
	apmLevel       pkg.Word // 91       Current APM level value (see 7.12.6.45)
	masterPassword pkg.Word // 92       Master Password Identifier (see 7.12.6.46)

	// part 14 of 19 (p153)
	hwResetResult pkg.Word // 93       Hardware reset result (see 7.12.6.47)
	_             pkg.Word // 94       Obsolete

	// part 15 of 19 (p154)
	streamMinReqSz  pkg.Word    // 95       Stream Minimum Request Size (see 7.12.6.49)
	streamDMATrTime pkg.Word    // 96       Streaming Transfer Time - DMA (see 7.12.6.50)
	streamAccLat    pkg.Word    // 97       Streaming Access Latency - DMA and PIO (see 7.12.6.51)
	streamPerfGran  pkg.DWord   // 98..99   Streaming Performance Granularity (DWORD) (see 7.12.6.52)
	sataSectors     pkg.QWord   // 100..103 Number of User Addressable Logical Sectors (QWord) (see 7.12.6.53)
	streamPIOTrTime pkg.Word    // 104      Streaming Transfer Time - PIO (see 7.12.6.54)
	maxDTSetMngCmd  pkg.Word    // 105      Maximum number of 512-byte blocks per DATA SET MANAGEMENT command (see 7.5)
	physPerLogSecs  pkg.Word    // 106      Physical sector size / logical sector size (see 7.12.6.56)
	interSeekDelay  pkg.Word    // 107      Inter-seek delay for ISO/IEC 7779 standard acoustic testing
	wwName          [4]pkg.Word // 108..111 World wide name (see 7.12.6.58)
	_               [4]pkg.Word // 112..115 Reserved
	_               pkg.Word    // 116      Obsolete
	sectorSize      pkg.DWord   // 117..118 Logical sector size (DWord) (see 7.12.6.61)

	// part 16 of 19 (p155)
	featureSet2 [2]pkg.Word // 119..120 Command and feature sets supported and enabled (Continued from words 82..84) (see 7.12.6.40)
	_           [6]pkg.Word // 121..126 Reserved for expanded supported and enabled settings

	// part 17 of 19 (p156)
	_               pkg.Word     // 127      Obsolete
	securityStat    pkg.Word     // 128      Security status (see 7.12.6.66)
	vendor          [31]pkg.Word // 129..159 Vendor specific
	cfa2            [8]pkg.Word  // 160..167 Reserved for CFA (see 7.12.6.68)
	formFactor      pkg.Word     // 168      See 7.12.6.69
	dtSetMngSupport pkg.Word     // 169      DATA SET MANAGEMENT command support (see 7.12.6.70)
	additionalIdent [4]pkg.Word  // 170..173 Additional Product Identifier (see 7.12.6.71)
	_               [2]pkg.Word  // 174..175 Reserved
	mediaSerial     [60]byte     // 176..205 Current media serial number (see 7.12.6.73)

	// part 18 of 19 (p157)
	sctCmdTransport pkg.Word    // 206      SCT Command Transport (see 7.12.6.74)
	_               [2]pkg.Word // 207..208 Reserved
	alignmentLSecs  pkg.Word    // 209      Alignment of logical sectors within a physical sector (see 7.12.6.75)
	wrVerifyMod3Cnt pkg.DWord   // 210..211 Write-Read-Verify Sector Mode 3 Count (DWord) (see 7.12.6.76)
	wrVerifyMod2Cnt pkg.DWord   // 212..213 Write-Read-Verify Sector Mode 2 Count (DWord) (see 7.12.6.76)
	_               [3]pkg.Word // 214..216 Obsolete
	rotationRate    pkg.Word    // 217      Nominal media rotation rate (see 7.12.6.79)
	_               pkg.Word    // 218      Reserved
	_               pkg.Word    // 219      Obsolete
	wrVerifyMode    pkg.Word    // 220      See 7.12.6.82
	_               pkg.Word    // 221      Reserved

	// part 19 of 19 (p158)
	transportMajor pkg.Word     // 222      Transport major version number (see 7.12.6.84)
	transportMinor pkg.Word     // 223      Transport minor version number (see 7.12.6.85)
	_              [6]pkg.Word  // 224..229 Obsolete
	extAddrSectors pkg.QWord    // 230..233 Extended Number of User Addressable Sectors (QWord) (see 7.12.6.87)
	minDLMicroOp   pkg.Word     // 234      Minimum number of 512-byte data blocks per Download Microcode operation (see 7.12.6.88)
	maxDLMicroOp   pkg.Word     // 235      Maximum number of 512-byte data blocks per Download Microcode operation (see 7.12.6.88)
	_              [19]pkg.Word // 236..254 Reserved
	integrity      pkg.Word     // 255      Integrity Word (see 7.12.6.91)
}