import { groupService } from "../../_services/group.service"
import { useState, useEffect } from "react"

export const EventAttendeeForm = ({ eventId }) => {

    const [eventRegistrationStatus, setEventRegistrationStatus] = useState("")

    const eventRegistrationStatusFetch = async () => {

        let payload = {
            EventId: eventId,
        }

        await groupService.eventRegister(payload)
        .then((r) => {
            setEventRegistrationStatus(r.data)
        })
        .catch((e) => {
            if (e.response.data !== undefined) {
                alert(e.response.data.error)
            } else {
                alert("problem server side")
            }
        })
    }

    const handleEventRegistration = async (e) => {

        console.log(e)
        let payload = {
            EventId: e,
            Status: eventRegistrationStatus,
        }

        await groupService.eventRegister(payload)
        .then((r) => {
            setEventRegistrationStatus(r.data)
        })
        .catch((e) => {
            if (e.response.data !== undefined) {
                alert(e.response.data.error)
            } else {
                alert("problem server side")
            }
        })
    }

    useEffect(() => {
        eventRegistrationStatusFetch()
    }, [])
    
    return (
        <button onClick={() => handleEventRegistration(eventId)}>{eventRegistrationStatus}</button>
    )
}

export default EventAttendeeForm