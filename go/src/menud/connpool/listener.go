package connpool

import (
	"menud/db"
)

type pooledConnection struct {
	dbConn db.Connection
}

func (this *pooledConnection) setUp() {
	this.dbConn = db.GetConnection()
}

func (this *pooledConnection) listen() {
	for {
		select {
		case req := <-getAttendeeChan:
			this.handleGetAttendee(req)
		case req := <-getAttendeeByKeyChan:
			this.handleGetAttendeeByKey(req)
		case req := <-getUserChan:
			this.handleGetUser(req)
		case req := <-getUserByEmailPasswordChan:
			this.handleGetUserByEmailPassword(req)
		case req := <-getEventChan:
			this.handleGetEvent(req)
		case req := <-getEventsChan:
			this.handleGetEventsForUser(req)
		case req := <-getAttendeesChan:
			this.handleGetAttendees(req)
		case req := <-getCoursesChan:
			this.handleGetCourses(req)
		case req := <-getOptionsChan:
			this.handleGetOptions(req)
		case req := <-getSelectionChan:
			this.handleGetSelection(req)
		case req := <-shutDownChan:
			shutDownChan <- req
			break
		}
	}
}

func (this *pooledConnection) handleGetUser(req getUserRequest) {
	var res getUserResponse
	res.user, res.err = this.dbConn.GetUser(req.userId)
	req.retChan <- res
}
func (this *pooledConnection) handleGetUserByEmailPassword(req getUserByEmailPasswordRequest) {
	var res getUserResponse
	res.user, res.err = this.dbConn.GetUserByEmailPassword(req.email, req.password)
	req.retChan <- res
}
func (this *pooledConnection) handleGetAttendee(req getAttendeeRequest) {
	var res getAttendeeResponse
	res.attendee, res.err = this.dbConn.GetAttendee(req.attendeeId)
	req.retChan <- res
}
func (this *pooledConnection) handleGetAttendees(req getAttendeesRequest) {
	var res getAttendeesResponse
	res.attendees, res.err = this.dbConn.GetAttendees(req.eventId)
	req.retChan <- res
}
func (this *pooledConnection) handleGetAttendeeByKey(req getAttendeeByKeyRequest) {
	var res getAttendeeResponse
	res.attendee, res.err = this.dbConn.GetAttendeeByKey(req.token)
	req.retChan <- res
}
func (this *pooledConnection) handleGetCourses(req getCoursesRequest) {
	var res getCoursesResponse
	res.crses, res.err = this.dbConn.GetCourses(req.eventId)
	req.retChan <- res
}
func (this *pooledConnection) handleGetOptions(req getOptionsRequest) {
	var res getOptionsResponse
	res.opts, res.err = this.dbConn.GetOptions(req.courseId)
	req.retChan <- res
}
func (this *pooledConnection) handleGetEvent(req getEventRequest) {
	var res getEventResponse
	res.event, res.err = this.dbConn.GetEvent(req.eventId)
	req.retChan <- res
}
func (this *pooledConnection) handleGetEventsForUser(req getEventsRequest) {
	var res getEventsResponse
	res.event, res.err = this.dbConn.GetEventsForUser(req.userId)
	req.retChan <- res
}

func (this *pooledConnection) handleGetSelection(req getSelectionRequest) {
	var res getSelectionResponse
	res.optionId, res.err = this.dbConn.GetSelection(req.attendeeId, req.courseId)
	req.retChan <- res
}