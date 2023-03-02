package helpers

type ReturnOnInvestment struct {
	N500ROIs   []string `json:"n500ROIs"`
	N1000ROIs  []string `json:"n1000ROIs"`
	N2000ROIs  []string `json:"n2000ROIs"`
	N5000ROIs  []string `json:"n5000ROIs"`
	N10000ROIs []string `json:"n10000ROIs"`
}

var Level_fractions = []float64{0.05, 0.04, 0.03, 0.02, 0.01, 0.005, 0.0025}
var Level_Contributors = []int{3, 9, 27, 81, 243, 729, 2187}
var Level_Contributors2 = []int{3, 12, 39, 120, 363, 1092, 3279}
var Entitlement_RoomOne = []float64{1500, 3500, 7500, 10400, 15000, 25000, 50000}
var Streams = []string{
	"Category",
	"Category",
	"Category",
	"Category",
	"Category",
}

var LevelCounts = []int{7, 8, 9, 10, 11}

var Level_fractions_ROOMONE = []float64{0.04, 0.03, 0.02, 0.01, 0.005, 0.0025, 0.00125, 0.0006}
var Level_Contributors_ROOMONE = []float64{3, 9, 27, 81, 243, 729, 2187, 6561}

var ROIs = [][]string{
	{
		"150.00",
		"450.00",
		"1,350.00",
		"4,050.00",
		"12,150.00",
		"36,450.00",
		"109,350.00",
	},
	{
		"300.00",
		"900.00",
		"2,700.00",
		"8,100.00",
		"24,300.00",
		"72,900.00",
		"218,700.00",
	},
	{
		"600.00",
		"1,800.00",
		"5,400.00",
		"16,200.00",
		"48,600.00",
		"145,800.00",
		"437,400.00",
	},
	{
		"1,500.00",
		"4,500.00",
		"13,500.00",
		"40,500.00",
		"121,500.00",
		"364,500.00",
		"1,093,500.00",
	},
	{
		"3,000.00",
		"9,000.00",
		"27,000.00",
		"81,000.00",
		"243,000.00",
		"729,000.00",
		"2,187,000.00",
	},
}

const ContibutionAmount = 10000

var CategoryAmount = []float64{500, 1000, 2000, 5000, 10000}

const PaymentSuccessful = "success"
const PaymentFailed = "failed"
const NROOMONE = 3280
const NROOMTWO = 9841
const ROOM_ZERO_REFUGEE_CENTER = "ROOM ZERO REFUGEE CENTER"
const ROOM_ONE = "ROOM ONE"
const ROOM_ONE_REFUGEE_CENTER = "PISHON REFUGEE CENTER"
const ROOM_TWO = "ROOM TWO"
const Paystackkey = "pk_live_5bd6919a85c60cc0d682061c08fd987fff5dc963"

var CategoryBankName = []string{"FCMB", "FCMB", "FCMB", "FCMB", "FCMB"}
var CategoryAccountName = []string{"The Humphrey Empire", "The Humphrey Empire", "The Humphrey Empire",
	"The Humphrey Empire", "The Humphrey Empire"}
var CategoryAccountNumber = []string{"9229852011", "9229852011", "9229852011",
	"9229852011", "9229852011"}
