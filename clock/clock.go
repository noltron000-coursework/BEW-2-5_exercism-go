package clock

import (
	"strconv"
)

type Clock struct {
	time int
}

func New(hrs, min int) Clock {
	/////////////////////////////////////////
	// hrs×60 = hours to minutes           //
	// hrs×60+min = only minutes           //
	// 24 = number of hours in a day       //
	// 24×60 = number of minutes in a day  //
	//-------------------------------------//
	// modulo division(%) will help remove //
	// extra hours and minutes for example //
	//        25:74 = 26:14 = 02:14        //
	/////////////////////////////////////////
	var time int
	time = ((hrs * 60) + min) % (24 * 60)
	// need this if statement; % can give (-) result
	if time < 0 {
		time += 1440
	}
	return Clock{time: time}
}

func (timepiece Clock) Add(min int) Clock {
	var time int
	time = (timepiece.time + min) % (24 * 60)
	// need this if statement; % can give (-) result
	if time < 0 {
		time += 1440
	}
	return Clock{time: time}
}

func (timepiece Clock) Subtract(min int) Clock {
	var time int
	time = (timepiece.time - min) % (24 * 60)
	// need this if statement; % can give (-) result
	if time < 0 {
		time += 1440
	}
	return Clock{time: time}
}

func (timepiece Clock) String() string {
	// Digital() function takes in a time
	// converts it to digital time format
	// then returns that formatted string

	var time int
	var min int
	var hrs int
	var digital string

	time = timepiece.time % (24 * 60)
	// need this if statement; % can give (-) result
	if time < 0 {
		time += 1440
	}

	min = time % 60
	hrs = time / 60

	if hrs < 10 {
		digital += "0" + strconv.Itoa(hrs) + ":"
	} else {
		digital += strconv.Itoa(hrs) + ":"
	}

	if min < 10 {
		digital += "0" + strconv.Itoa(min)
	} else {
		digital += strconv.Itoa(min)
	}

	return digital
}
