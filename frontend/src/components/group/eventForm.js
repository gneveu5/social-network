import { groupService } from "../../_services/group.service"
import { useState } from "react"

export const EventForm = ({ handleEventSubmitReturn }) => {

    const urlParams = window.location.href

    const [eventTitle, setEventTitle] = useState("")
    const [eventContent, setEventContent] = useState("")
    const [eventDate, setEventDate] = useState("")

    const handleEventSubmit = async (e) => {

        e.preventDefault()
        let payload
        payload = {
            Title: eventTitle,
            Description: eventContent,
            EventDate: eventDate,
            groupId: urlParams.split('/')[4],
        }

        await groupService.eventing(payload)
        .then((r) => {
            let tmp = r.data
            for (let i in tmp) {
                tmp[i]["Opened"] = "See attendees"
            }
            handleEventSubmitReturn(tmp)
        })
        .catch((e) => {
            if (e.response.data !== undefined) {
                alert(e.response.data.error)
            } else {
                alert("problem server side")
            }
        })
    }
    
    return (
        <form className="event-form" onSubmit={(event) => handleEventSubmit(event)}>
            <div>
                <label htmlFor={"event-name"}>Event name</label>
                <input
                    id="event-name"
                    type="text"
                    value={eventTitle}
                    onChange={(e) => setEventTitle(e.target.value)}
                />
            </div>
            <div>
                <label htmlFor={"event-content"}>Description</label>
                <input
                    id="event-content"
                    type="text"
                    value={eventContent}
                    onChange={(e) => setEventContent(e.target.value)}
                />
            </div>
            <div>
                <label htmlFor="event-date">Event date:</label>
                <input
                    id="event-date"
                    type="date"
                    value={eventDate}
                    onChange={(e) => setEventDate(e.target.value)}
                />
            </div>
            <button type="submit">Submit</button>
        </form>
    )
}

export default EventForm