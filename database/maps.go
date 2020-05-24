 package database

var (
	DKMaps = map[int16][]int16{
		18: {18, 193, 200}, 19: {19, 194, 201}, 25: {25, 195, 202}, 26: {26, 196, 203}, 27: {27, 197, 204}, 29: {29, 198, 205}, 30: {30, 199, 206}, // Normal Maps
		193: {18, 193, 200}, 194: {19, 194, 201}, 195: {25, 195, 202}, 196: {26, 196, 203}, 197: {27, 197, 204}, 198: {29, 198, 205}, 199: {30, 199, 206}, // DK Maps
		200: {18, 193, 200}, 201: {19, 194, 201}, 202: {25, 195, 202}, 203: {26, 196, 203}, 204: {27, 197, 204}, 205: {29, 198, 205}, 206: {30, 199, 206}, // Normal Maps
	}

	sharedMaps = []int16{1, 2, 3, 14, 15, 20, 21, 22, 23, 24, 26, 27, 32, 33, 34, 36, 37, 38, 42, 43, 44, 45, 46, 47,
		89, 93, 94, 95, 100, 101, 102, 108, 109, 110, 111, 112, 120, 121, 122, 123, 160, 161, 162, 163, 164, 165, 166, 167, 168, 169, 170,
		213, 214, 221, 222, 223, 224, 225, 226, 227, 228, 233, 236, 237, 238, 239, 240, 244, 252, 254}

	PvPZones     = []int16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 199, 206, 194, 200, 204, 198, 196, 197, 233, 236, 237, 238, 239, 240, 254}
	unlockedMaps = []int16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 199, 206, 194, 200, 204, 198, 196, 197, 233, 236, 237, 238, 239, 240, 254}
		
)
