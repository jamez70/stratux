package main

import (
	"fmt"
	"log"
	"os"
	//	"reflect"
	"time"
)

const (
	TRACKLOG_TIMESTAMP_RESOLUTION = 250 * time.Millisecond
	TRACKLOG_DIRECTORY            = "/root/tracklog"
)

var trackLogWriter chan string
var trackLogRun chan bool
var shutdownTrackLog chan bool
var runTrackLog bool
var trackLogFilename string

// Set up the structs for track logging
func initTrackLog() {

	runTrackLog = false
	trackLogWriter = make(chan string)
	trackLogRun = make(chan bool)
	shutdownTrackLog = make(chan bool)
	// Lets wait until trackLogRun is set, then while it is true, we will record until shutdownTrackLog is set
	for {
		select {
		case s := <-trackLogWriter:
			if runTrackLog {
				_, err := fp.WriteString(s)
				if err != nil {
					fmt.Printf("Error writing file\n")
				}
			}
		case <-trackLogRun:
			if runTrackLog == false {
				runTrackLog = true
			} else {
				fmt.Printf("Track already running?\n")
			}

		case <-shutdownTrackLog:
			log.Printf("TRACKLOG: Shutdown\n")
			runTrackLog = false
		}
	}

}

func trackHeaderLine() string {
	var str string
	str = "FixUTC,Latitude,Longitude,Quality,Satellites,SatellitesTracked,SatellitesSeen,Altitude,AccuracyVert,GPSVertVel,TrueCourse,GroundSpeed,PressureAlt,Pitch,Roll,GyroHeading\n"
	return str

}

func trackSituation() {
	var trackStr string
	if globalStatus.TrackIsRecording == true {
		trackStr = fmt.Sprintf("%8.2f,%f,%f,%d,%d,%d,%d,%3.1f,%3.2f,%f,%f,%d,%f,%f,%f,%f\n",
			mySituation.LastFixSinceMidnightUTC,
			mySituation.Lat, mySituation.Lng, mySituation.Quality,
			mySituation.Satellites, mySituation.SatellitesTracked, mySituation.SatellitesSeen,
			mySituation.Alt, mySituation.AccuracyVert, mySituation.GPSVertVel, mySituation.TrueCourse,
			mySituation.GroundSpeed, mySituation.Pressure_alt, mySituation.Pitch, mySituation.Roll, mySituation.Gyro_heading)
		trackLogWriter <- trackStr
	}
}

func stopTrackLog() {
	runTrackLog = false
	shutdownTrackLog <- true
	globalStatus.TrackRecordingStatus = "Stopped"
	globalStatus.TrackIsRecording = false
	fp.Close()
}

var fp *os.File

func startTrackLog(t time.Time) {
	var fName string
	fName = "TrackLog" + t.Format("20060102150405") + ".csv"
	trackLogFilename = TRACKLOG_DIRECTORY + "/" + fName
	log.Printf("startTrackLog: %s\n", trackLogFilename)
	direrr := os.MkdirAll(TRACKLOG_DIRECTORY, 0755)
	if direrr != nil {
		fmt.Printf("Create dir error\n")
	}
	fplocal, err := os.OpenFile(trackLogFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Cannot open %s for writing\n", trackLogFilename)
		return
	}
	fp = fplocal
	fp.WriteString(trackHeaderLine())
	trackLogRun <- true
	globalStatus.TrackIsRecording = true
	globalStatus.TrackRecordingStatus = "Recording " + fName
}
