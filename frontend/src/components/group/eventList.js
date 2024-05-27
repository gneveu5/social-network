import FormEvent from './eventForm'
import EventAttendeeForm from "./eventAttendeeForm"
import { groupService } from "../../_services/group.service"
import { useState, useEffect} from "react"
import "../../css/post.css"

export const EventList = () => {

    const urlParams = window.location.href
    const [events, setEvents] = useState([])
    const [attendees, setAttendees] = useState([])

    const handleAttendeeFetch = async (e) => {
        const attendeesExists = attendees.some((attendee) => attendee.id === e.Id)
        if (!attendeesExists) {
            e.Opened = "Close"
            await groupService.eventAttendeesFetch(e.Id)
            .then((r) => {
                setAttendees(attendees => [...attendees, { id: e.Id, data: r.data }])
            })
            .catch((err) => {
                if (err.response.data !== undefined) {
                    alert(err.response.data.error)
                } else {
                    alert("problem server side")
                }
            })
        } else {
            e.Opened = "See attendees"
            setAttendees(attendees.filter((attendee) => attendee.id !== e.Id))
        }
    }

    const handleEventSubmitReturn = async (e) => {
        setEvents(e)
    }

    useEffect(() => {
        const fetchData = async () => {
            try {
                const r = await groupService.eventFetch(urlParams.split('/')[4])
                let tmp = r.data
                for (let i in tmp) {
                    tmp[i]["Opened"] = "See attendees"
                }
                setEvents(tmp)
            } catch (e) {
                console.error("Error fetching data:", e)
            }
        }
        fetchData()
    }, [])

    return (
        <div>
            <FormEvent handleEventSubmitReturn={handleEventSubmitReturn}/>
            {events && (events.map((item) => (
                <div key={"event" + item.Id} id={item.Id} className="posts">
                    <div className="post-header">
                        <div className="header-subdiv">
                            <div className="header-title">{item.Title}</div>
                            <div className="header-div-name">{"event created by " + item.CreatorName}</div>
                        </div>
                        <div className="header-date">{item.CreatedAt.replace(/[a-z]/gi, ' ')}</div>
                    </div>
                    <div>
                        <div className="post-body">Description: {item.Description}</div>
                        <div className="post-body">When: {item.EventDate}</div>
                    </div>
                    <button className="see-attendees" onClick={() => handleAttendeeFetch(item)}>{item.Opened}</button>
                    <EventAttendeeForm eventId={item.Id}/>
                    {attendees && (attendees.map((attendee) => attendee.id === item.Id && (
                        <div key={"event-attendees" + item.Id}>
                            {attendee.data && attendee.data.length > 0 ? (
                                attendee.data.map((person) => (
                                    <div key={"person" + person.Id} className="comments">{person.AttendeeName}</div>
                                ))
                            ) : (
                                <div className="comments">No attendee for this event.</div>
                            )}
                        </div>
                    )))}
                </div>
            )))}
        </div>
    )

}
  
export default EventList