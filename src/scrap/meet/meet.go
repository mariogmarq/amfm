package meet

import (
	"sync"
	"time"

	"github.com/go-rod/rod"
)

//MeetSession repressents a google meet session of a given user
type MeetSession struct {
	browser          *rod.Browser
	page             *rod.Page
	inCall           bool
	enterTime        time.Time
	expectedDuration time.Duration
	expectedExitTime time.Time
	url              string
	sincronize       sync.Mutex
}

//NewSession creates a new meetSession with an opened browser
func NewSession(browser *rod.Browser) *MeetSession {
	session := MeetSession{browser, nil, false, time.Now(), time.Second, time.Now(), "", sync.Mutex{}}
	return &session

}

//ConnectMeet connects a session to a give meet page, giving a wrong url may block the actual
//coroutine
func (m *MeetSession) ConnectMeet(url string, duration time.Duration) *MeetSession {
	//Connect to meet
	m.page = m.browser.MustPage(url)

	//Update things
	m.url = url
	m.expectedDuration = duration

	return m
}

//JoinMeet enter the meet with a given meet session
func (m *MeetSession) JoinMeet() *MeetSession {

	if m.page != nil && !m.inCall {

		//disable audio
		m.page.MustSearch("div.IYwVEf.HotEze.uB7U9e.nAZzG").MustClick()
		m.page.MustSearch("div.IYwVEf.HotEze.nAZzG").MustClick()

		//Join meet
		m.page.MustSearch("span.NPEfkd.RveJvd.snByac").MustClick()

		//Update session data
		m.inCall = true
		m.enterTime = time.Now()
		m.expectedExitTime = m.enterTime.Add(m.expectedDuration)
	}

	return m
}

//Wait launchs a coroutine for exiting the meet call after reaching the expected time by closing the
//page
func (m *MeetSession) Wait() {
	if m.inCall {
		m.sincronize.Lock()
		go func() {
			time.Sleep(m.expectedExitTime.Sub(m.enterTime))
			m.browser.MustClose()
			m.page.MustClose()
			m.page = nil
			m.inCall = false
			m.sincronize.Unlock()
		}()
	}
}
